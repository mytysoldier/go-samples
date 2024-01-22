package server

import "database/sql"

func GetBooks(db *sql.DB) {
	row := db.QueryRow("SELECT ")
}
