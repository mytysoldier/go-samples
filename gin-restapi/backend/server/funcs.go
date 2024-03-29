package server

import (
	"backend/model"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetBookByID(db *sql.DB, id int) (model.Book, error) {
	row := db.QueryRow("SELECT id, name, price FROM book WHERE id = $1", id)

	var book model.Book
	err := row.Scan(&book.Id, &book.Name, &book.Price)
	if err != nil {
		log.Fatal("error scanning book")
		return model.Book{}, err
	}

	return book, nil
}

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

func InsertBook(db *sql.DB, book model.Book) (model.Book, error) {
	var insertedBook model.Book

	err := db.QueryRow("INSERT INTO book (name, price) VALUES ($1, $2) RETURNING id, name, price", book.Name, book.Price).Scan(&insertedBook.Id, &insertedBook.Name, &insertedBook.Price)

	if err != nil {
		log.Fatal("error inserting book")
		fmt.Println(err)
		return model.Book{}, err
	}

	fmt.Println("inserted book")

	return insertedBook, nil
}

func UpdateBook(db *sql.DB, book model.Book) (model.Book, error) {
	var updatedBook model.Book

	err := db.QueryRow("UPDATE book SET name=$1, price=$2 WHERE id=$3 RETURNING id, name, price", book.Name, book.Price, book.Id).Scan(&updatedBook.Id, &updatedBook.Name, &updatedBook.Price)

	if err != nil {
		log.Fatal("error updating book")
		fmt.Println(err)
		return model.Book{}, err
	}

	fmt.Println("updated book")

	return updatedBook, nil
}

func DeleteBook(db *sql.DB, id int) error {
	existingBook, err := GetBookByID(db, id)
	if err != nil {
		log.Fatal("error deleting book")
		return err
	}

	// 本が存在しない場合はエラーを返す
	if existingBook.Id == 0 {
		return errors.New("book not found")
	}

	// 本が存在する場合は削除を実行
	_, err = db.Exec("DELETE FROM BOOK WHERE id = $1", id)
	if err != nil {
		log.Fatal("error deleting book")
		return err
	}

	return nil
}

func DeleteAllBook(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM BOOK")
	if err != nil {
		log.Fatal("error deleting all book")
		return err
	}

	return nil
}
