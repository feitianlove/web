package main

import (
	"fmt"
	"github.com/feitianlove/web/api/middleware"
	"github.com/feitianlove/web/auth"
	"github.com/feitianlove/web/config"
	"github.com/feitianlove/web/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%+v\n", conf.CasBin)
	err = auth.Init(*conf.CasBin)
	if err != nil {
		panic(err)
	}
	err = logger.InitLog(conf)
	if err != nil {
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
