package master

import (
	"context"
	"errors"
	"fmt"
	"github.com/feitianlove/web/config"
	"github.com/feitianlove/web/logger"
	pb "github.com/feitianlove/web/master/master_pb/m_pb"
	w_pb "github.com/feitianlove/web/master/master_pb/w_pb"
	"github.com/feitianlove/web/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"hash/crc32"
	"math"
	"net"
	"sync"
	"time"
)

// 前端收到上架信息调用注册到master的 worker执行任务

type Master struct {
	Lock          sync.RWMutex
	ListenPort    int64
	Domain        string
	Token         string
	ClientAddress []*WorkerMess
	totalWeight   int32
	md            map[string]uint32
}

//下面
type WorkerMess struct {
	Domain          string
	Port            int64
	Weight          int32 // 配置的权重，即在配置文件或初始化时约定好的每个节点的权重
	currentWeight   int32 //节点当前权重，会一直变化
	EffectiveWeight int32 //有效权重，初始值为weight, 通讯过程中发现节点异常，则-1 ，之后再次选取本节点，调用成功一次则+1，
	// 直达恢复到weight 。 用于健康检查，处理异常节点，降低其权重。
}

func NewMaster(conf *config.Config) *Master {
	return &Master{
		ListenPort:    conf.Master.ListenPort,
		Domain:        conf.Master.Domain,
		Token:         conf.Master.Token,
		ClientAddress: make([]*WorkerMess, 0),
		md:            make(map[string]uint32),
	}
}

func Run(m *Master, conf *config.Config) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", conf.Master.Domain, conf.Master.ListenPort))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRegisterServer(grpcServer, m)
	err = grpcServer.Serve(listener)
	if err != nil {
		return err
	}
	return nil

}

func (master *Master) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	t := time.Now().Unix()
	var response *pb.RegisterResponse = &pb.RegisterResponse{
		Code:    0,
		Message: "success",
	}
	defer func() {
		master.Lock.Unlock()
		logger.CtrlLog.WithFields(logrus.Fields{
			"request":     fmt.Sprintf("%+v\n", request),
			"response":    fmt.Sprintf("%+v\n", response),
			"requestTime": time.Now().Unix() - t,
		}).Info()
	}()
	if request.Token != master.Token {
		response.Code = -1
		response.Message = "Permission denied"
		return response, errors.New("you token is Invalid")
	}
	if len(request.Ip) == 0 || request.Port < 1000 {
		response.Code = -1
		response.Message = fmt.Sprintf("params is invalid. %s:%d", request.Ip, request.Port)
		return response, errors.New("params is Invalid")
	}

	addr := fmt.Sprintf("%s:%d", request.Ip, request.Port)
	crc := crc32.ChecksumIEEE([]byte(addr))
	if _, ok := master.md[addr]; ok {
		response.Code = 4000
		response.Message = fmt.Sprintf("this %s is already register", addr)
		return response, errors.New(fmt.Sprintf("this %s is already register", addr))
	}
	master.Lock.Lock()
	master.totalWeight += request.Weight
	master.md[addr] = crc
	master.ClientAddress = append(master.ClientAddress, &WorkerMess{
		Domain:          request.Ip,
		Port:            request.Port,
		Weight:          request.Weight,
		currentWeight:   request.Weight,
		EffectiveWeight: request.Weight,
	})
	return response, nil
}

// 上架任务结构体
type CosGround struct {
	ModuleName string
	IpSet      []string
}

func (master *Master) Schedule(data interface{}) error {
	if req, ok := data.(CosGround); !ok {
		return errors.New("your param is invalid")
	} else {
		for _, item := range req.IpSet {
			workerAddr := master.WeightedRoundRobin()
			requestId, _ := util.AllocateRequestId()
			// 根据module Name分解每个模块的安装步骤
			task, err := master.DivideByModuleName(item, req.ModuleName)
			if err != nil {
				logger.CtrlLog.WithFields(logrus.Fields{
					"Error": err,
				}).Info("DivideByModuleName err")
			}
			// 将每个步骤发送给worker
			client, err := ConnectWorker(workerAddr)
			if err != nil {
				return err
			}
			req := &w_pb.TaskRequest{
				Data:      task,
				RequestId: requestId,
				Ip:        item,
			}
			ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
			response, err := client.DistributeTask(ctx, req)
			if err != nil {
				logger.CtrlLog.WithFields(logrus.Fields{
					"Message":   response.Message,
					"RequestID": requestId,
					"Code":      response.Code,
				}).Info(err)
				return err
			}
			logger.CtrlLog.WithFields(logrus.Fields{
				"Message":   response.Message,
				"RequestID": requestId,
				"Code":      response.Code,
			}).Info("Schedule task success")
		}
	}
	return nil
}

// 加权轮询
func (master *Master) WeightedRoundRobin() string {
	var maxNode *WorkerMess = &WorkerMess{
		currentWeight: math.MinInt32,
	}
	for _, worker := range master.ClientAddress {
		if worker.currentWeight > maxNode.currentWeight {
			maxNode = worker
		}
		worker.currentWeight += worker.EffectiveWeight
	}
	//找到最大的之后
	maxNode.currentWeight = maxNode.currentWeight - master.totalWeight
	return fmt.Sprintf("%s:%d", maxNode.Domain, maxNode.Port)
}

//通过grpc 发送给worker 任务
func ConnectWorker(addr string) (w_pb.TaskClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	grpcClient := w_pb.NewTaskClient(conn)
	return grpcClient, nil
}

func (master *Master) DivideByModuleName(ip string, moduleName string) ([]string, error) {
	var res []string
	switch moduleName {
	case "CosHttpSvr":
		res = master.DivideCosHttpSvr(ip)
	}
	return res, nil

}

func (master *Master) DivideCosHttpSvr(ip string) []string {
	var res = make([]string, 0)
	// 	创建目录
	res = append(res, "mkdir -p /usr/local/httpsvr/")
	res = append(res, "echo "+ip+" >> /usr/local/httpsvr/test.conf")
	return res
}
