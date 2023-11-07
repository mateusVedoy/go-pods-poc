package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	PATH := "localhost:8081"
	router := gin.Default()

	router.GET("/", SayHello)
	router.GET("/health", GetHealth)

	router.Run(PATH)
}

func SayHello(context *gin.Context) {
	context.IndentedJSON(
		http.StatusOK,
		"Hello. I'm Up on port 8081",
	)
}

func GetHealth(context *gin.Context) {
	context.IndentedJSON(
		http.StatusOK,
		"OK, I'M UP",
	)
}
