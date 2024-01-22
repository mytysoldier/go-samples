package model

import "fmt"

type Book struct {
	Id    int
	Name  string
	Price int
}

func GetRandomBooks() []Book {
	var books []Book

	for i := 1; i < 6; i++ {
		book := Book{
			Id:    i,
			Name:  fmt.Sprintf("本タイトルその%d", i),
			Price: i * 1000,
		}
		books = append(books, book)
	}

	return books
}
