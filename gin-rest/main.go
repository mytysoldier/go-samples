package main

import (
	"database/sql"
	"gin-rest/funcs"
	"gin-rest/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func main() {
	initDB()
	r := gin.Default()
	r.Use(corsMiddleware())
	r.GET("/", handleHome)
	r.POST("/user", handleCreateUser)
	r.GET("/user/:id", handleGetUserByID)
	r.GET("/users", handleGetUsers)
	r.PUT("/user/:id", handleUpdateUser)
	r.DELETE("/user/:id", handleDeleteUser)
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

func initDB() {
	connStr := "user=user password=password dbname=sample sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
		panic("failed to connect DB")
	}

	DB = db
}

func handleHome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "test"})
}

func handleCreateUser(ctx *gin.Context) {
	var user model.User

	// リクエストボディからデータを取得
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := funcs.InsertUser(DB, user)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func handleGetUserByID(ctx *gin.Context) {
	// パスパラメーター "id" を取得
	id := ctx.Param("id")

	// id を整数に変換
	userID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// GetBookByID 関数を呼び出し
	user, err := funcs.GetUserByID(DB, userID)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": "Failed to retrieve user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func handleGetUsers(ctx *gin.Context) {


	users, err := funcs.GetUsers(DB)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": "Failed to retrieve users"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func handleUpdateUser(ctx *gin.Context) {
	// パスパラメーター "id" を取得
	id := ctx.Param("id")

	// id を整数に変換
	userID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var user model.User

	// リクエストボディからデータを取得
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// リクエストボディーに含まれるIDを更新対象のIDとして設定
	user.Id = userID

	user, err = funcs.UpdateUser(DB, user)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func handleDeleteUser(ctx *gin.Context) {
	// パスパラメーター "id" を取得
	id := ctx.Param("id")

	// id を整数に変換
	userID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = funcs.DeleteUser(DB, userID)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
