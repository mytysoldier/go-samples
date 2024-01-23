package server

import (
	"backend/db"
	"backend/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server() {
	r := gin.Default()
	r.Use(corsMiddleware())
	r.GET("/", handleHome)
	r.GET("/books", handleSearchBooks)
	r.Run()
}

// CORSミドルウェアの定義
func corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Range")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.Next()
	}
}

func handleHome(ctx *gin.Context) {
	books := model.GetRandomBooks()
	ctx.JSON(http.StatusOK, gin.H{"books": books})
}

func handleSearchBooks(ctx *gin.Context) {
	books, err := GetBooks(db.DB)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"books": books})
}