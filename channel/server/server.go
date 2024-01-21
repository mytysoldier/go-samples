package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server() {
	r := gin.Default()
	r.GET("/", handleHome)
	r.Run()
}

func handleHome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "test"})
}
