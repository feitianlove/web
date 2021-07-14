package main

import (
	"fmt"
	"github.com/feitianlove/web/api/middleware"
	"github.com/feitianlove/web/auth"
	"github.com/feitianlove/web/config"
	"github.com/feitianlove/web/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	conf, err := config.InitConfig()
	logger.CtrlLog.WithFields(logrus.Fields{}).Info("init  config success")
	if err != nil {
		logger.CtrlLog.WithFields(logrus.Fields{
			"error": err,
		}).Error("init config  fail")
		panic(err)
	}
	//fmt.Printf("%+v\n", conf.CasBin)
	err = auth.Init(*conf.CasBin)
	logger.CtrlLog.WithFields(logrus.Fields{}).Info("init  auth success")
	if err != nil {
		logger.CtrlLog.WithFields(logrus.Fields{
			"error": err,
		}).Error("auth init  fail")
		panic(err)
	}

	err = logger.InitLog(conf)
	logger.CtrlLog.WithFields(logrus.Fields{}).Info("init log  success")
	if err != nil {
		logger.CtrlLog.WithFields(logrus.Fields{
			"err": err,
		}).Error("init log  fail")
		panic(err)
	}
	InitWeb(conf)
}
func InitWeb(config *config.Config) {
	server := gin.New()
	server.Use(middleware.Permission())
	server.GET("/test", func(c *gin.Context) {
		c.JSON(200, "success")
	})
	err := server.Run(":8080")
	fmt.Println(err)
}
