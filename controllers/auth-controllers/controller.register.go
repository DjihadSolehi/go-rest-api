package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello wordl register",
	})
}