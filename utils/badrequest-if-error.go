package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequestIfError(err error, ctx *gin.Context) {
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func BadRequestIfConditionTruly(condition bool, message string, ctx *gin.Context) {
	if condition {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}
}
