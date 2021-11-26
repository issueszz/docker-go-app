package main

import (
	"github.com/gin-gonic/gin"
	"goapp/services"
)

func main() {
	router := gin.Default()

	// err := model.InitMysql()
	//if err != nil {
	//	panic(err)
	//}

	router.GET("/ping", services.PingHandle())

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
