package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusVedoy/go-pods-poc.git/service-one/src/controller"
)

func StartRouter() {
	PATH := "localhost:8080"
	router := gin.Default()

	router.GET("/pods/up/:amount", controller.IncreasePods)
	router.GET("/pods/down/:amount", controller.DecreasePods)

	router.Run(PATH)
}
