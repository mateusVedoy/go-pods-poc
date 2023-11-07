package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IncreasePods(context *gin.Context) {
	amount := context.Param("amount")
	result := "Pods increased to " + amount + " instances"

	context.IndentedJSON(
		http.StatusOK,
		result,
	)
}

func DecreasePods(context *gin.Context) {
	amount := context.Param("amount")
	result := "Pods decreased to " + amount + " instances"

	context.IndentedJSON(
		http.StatusOK,
		result,
	)
}
