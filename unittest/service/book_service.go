package service

import (
	"unittest/model"
	"unittest/repository"
)

type BookService struct {
	Repository repository.BookRepository
}

func (service *BookService) GetBookById(id int) model.Book {
	return service.Repository.GetBookById(id)
}
