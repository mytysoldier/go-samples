package repository

import "unittest/model"

type BookRepository interface {
	GetBookById(id int) model.Book
}
