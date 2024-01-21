package server

import (
	"gin-restapi/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server() {
	r := gin.Default()
	r.GET("/", handleHome)
	r.Run()
}

func handleHome(ctx *gin.Context) {
	books := model.GetRandomBooks()
	ctx.JSON(http.StatusOK, gin.H{"books": books})
}
