package main

import (
	"fmt"
	"github.com/feitianlove/web/api/web"
	"github.com/feitianlove/web/auth"
	"github.com/feitianlove/web/config"
	"github.com/feitianlove/web/logger"
	"github.com/feitianlove/web/master"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%+v\n", conf.CasBin)
	//fmt.Printf("%+v\n", conf.Master)

	err = auth.Init(*conf.CasBin)
	if err != nil {
		panic(err)
	}

	err = logger.InitLog(conf)
	logger.CtrlLog.WithFields(logrus.Fields{}).Info("init log success")
	if err != nil {
		logger.CtrlLog.WithFields(logrus.Fields{
			"err": err,
		}).Error("init log fail")
		panic(err)
	}
	webClient, srv := InitWeb(conf)

	defer func() {
		//关闭srv
		_ = srv.Close()
		//TODO 关闭webClient的开启的数据库等
		fmt.Println(webClient)
	}()
	// 启动master
	m := master.NewMaster(conf)
	err = master.Run(m, conf)
	if err != nil {
		logger.CtrlLog.WithFields(logrus.Fields{
			"err": err,
		}).Error("init master fail")
		panic(err)
	}
	logger.CtrlLog.WithFields(logrus.Fields{}).Info("init master success")
	wg.Add(1)
	wg.Wait()
}

func InitWeb(config *config.Config) (*web.ClientWeb, *http.Server) {
	client, err := web.NewWebClient(config)
	if err != nil {
		logger.CtrlLog.WithFields(logrus.Fields{
			"error": err,
		}).Error("NewStore err")
		panic(err)
	}
	logger.CtrlLog.WithFields(logrus.Fields{}).Info("NewStore success")
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	server.Use(static.Serve("/", static.LocalFile(config.Web.StaticDir, false)))
	//server.Use(middleware.Permission())
	server.NoRoute(func(c *gin.Context) {
		// index.html no cache
		c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
		c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
		c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
		c.File(fmt.Sprintf("%s/index.html", config.Web.StaticDir))
	})
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Web.Domain, config.Web.ListenPort),
		Handler: server,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			logger.CtrlLog.WithFields(logrus.Fields{
				"error": err,
			}).Error("listen: %s\n")
			panic(err)
		}
	}()
	return client, srv
}
