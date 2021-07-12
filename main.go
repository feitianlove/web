package main

import (
	"fmt"
	"github.com/feitianlove/web/auth"
	"github.com/feitianlove/web/config"
)

func main() {
	//engin := gin.New()
	//engin.GET("/", func(c *gin.Context) {
	//	c.JSON(200, "success")
	//})
	//err := engin.Run(":8080")
	//if err != nil {
	//	panic(err)
	//}
	tests()
}
func tests() {
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", conf.CasBin)
	err = auth.Init(*conf.CasBin)
	if err != nil {
		panic(err)
	}
	//
	t, err := auth.AddPolicy("ftfeng", "/test/*", "POST|GET")
	fmt.Println(err, t)
	fmt.Println(auth.CheckPolicy("ftfeng", "/test/", "POST"))
	fmt.Println(auth.CheckPolicy("ftfeng", "/test/", "GET"))

}
