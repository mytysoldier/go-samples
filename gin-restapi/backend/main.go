package main

import (
	"backend/db"
	"backend/server"
)

func main() {
	db.InitDB()
	server.Server()
}
