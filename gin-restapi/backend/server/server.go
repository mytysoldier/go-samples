package server

import (
	"backend/db"
	"backend/model"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Server() {
	r := gin.Default()
	r.Use(corsMiddleware())
	r.GET("/", handleHome)
	r.GET("/book", handleSearchBook)
	r.GET("/books", handleSearchBooks)
	r.POST("/book", handleInsertBook)
	r.PUT("/book/:id", handleUpdateBook)
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

func handleSearchBook(ctx *gin.Context) {
	// リクエストパラメーター "id" を取得
	id := ctx.Query("id")

	bookID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// GetBookByID 関数を呼び出し
	book, err := GetBookByID(db.DB, bookID)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve book"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"book": book})
}

func handleSearchBooks(ctx *gin.Context) {
	books, err := GetBooks(db.DB)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"books": books})
}

func handleInsertBook(ctx *gin.Context) {
	var book model.Book

	// リクエストボディからデータを取得
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Bookリクエストデータ:", book)

	book, err := InsertBook(db.DB, book)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"book": book})
}

func handleUpdateBook(ctx *gin.Context) {
	// パスパラメーター "id" を取得
	id := ctx.Param("id")

	// id を整数に変換
	bookID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var book model.Book

	// リクエストボディからデータを取得
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// リクエストボディーに含まれるIDを更新対象のIDとして設定
	book.Id = bookID

	fmt.Println("Bookリクエストデータ:", book)

	updatedBook, err := UpdateBook(db.DB, book)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"book": updatedBook})
}

func handleDeleteBook() {}

func handleDeleteAllBook() {}
