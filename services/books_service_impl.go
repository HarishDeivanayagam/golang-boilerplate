package services

import (
	"clean-architecture/di/injector"
	"clean-architecture/entities"
	"clean-architecture/repository"
)

// BooksServiceImpl implements BookService
type BooksServiceImpl struct{}

// GetAllBooks returns all books from database
func (b *BooksServiceImpl) GetAllBooks() []entities.Book {
	br, ok := injector.Inject("BooksRepository").(repository.BooksRepository)
	if !ok {
		return nil
	}
	books, err := br.GetAllBooks()
	if err != nil {
		return nil
	}
	return books
}

// GetBookByID returns the book for the given ID
func (b *BooksServiceImpl) GetBookByID(id string) entities.Book {
	br, ok := injector.Inject("BooksRepository").(repository.BooksRepository)
	emptyBook := entities.Book{}
	if !ok {
		return emptyBook
	}
	book, err := br.GetBookByID(id)
	if err != nil {
		return emptyBook
	}
	return book
}
