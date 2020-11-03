package repository

import (
	"clean-architecture/entities"
)

// BooksRepository holds all methods in a bookRepository implementation
type BooksRepository interface {
	GetAllBooks() ([]entities.Book, error)
	GetBookByID(id string) (entities.Book, error)
}
