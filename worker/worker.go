package main

import (
	"context"
	"fmt"
	wpb "github.com/feitianlove/web/master/master_pb/m_pb"
	"google.golang.org/grpc"
)

type Worker struct {
	Id string
}

func main() {
	grpc.WithInsecure()
	conn, err := grpc.Dial(":8900", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = conn.Close()
	}()
	grpcClient := wpb.NewRegisterClient(conn)
	req := wpb.RegisterRequest{
		Port:  1001,
		Ip:    "127.0.0.1",
		Token: "ftfeng",
	}
	res, err := grpcClient.Register(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Code)
	fmt.Println(res.Message)

}

func RunWorker() {

}
