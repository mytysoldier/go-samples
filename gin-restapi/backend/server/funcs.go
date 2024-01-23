package server

import (
	"backend/model"
	"database/sql"
	"log"
)

func GetBooks(db *sql.DB) ([]model.Book, error) {
	rows, err := db.Query("SELECT id, name, price FROM book")
	if err != nil {
		log.Fatal("error fetching books")
		return nil, err
	}
	defer rows.Close()

	var books []model.Book

	for rows.Next() {
		var book model.Book
		book = model.Book{}

		if err := rows.Scan(&book.Id, &book.Name, &book.Price); err != nil {
			log.Fatal("error scaning books")
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
