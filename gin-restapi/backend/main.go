package main

import (
	"backend/db"
	"backend/server"
)

func main() {
	server.Server()
	db.InitDB()
}
