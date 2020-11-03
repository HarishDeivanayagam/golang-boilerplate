package services

import (
	"clean-architecture/entities"
)

// BooksService Queries BookRepository
type BooksService interface {
	GetAllBooks() []entities.Book
	GetBookByID(id string) entities.Book
}
