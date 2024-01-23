package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// 外部からアクセス可能なDB変数を定義
var DB *sql.DB

func InitDB() {
	connStr := "user=user password=password dbname=sample sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
		panic("failed to connect DB")
	}

	DB = db

	log.Println("DB initialized")
	fmt.Println("DB initialized")
}
