package infrastructure

import (
	"clean-architecture/entities"
	"errors"

	"gorm.io/gorm"
)

//BooksRepositoryImpl implements bookRepository Interface using gorm
type BooksRepositoryImpl struct {
	DB *gorm.DB
}

// GetAllBooks returns all books in database
func (b *BooksRepositoryImpl) GetAllBooks() ([]entities.Book, error) {
	var books []entities.Book
	dbRes := b.DB.Find(&books)
	if dbRes.Error != nil {
		return nil, errors.New("Unable to fetch data")
	}
	return books, nil
}

// GetBookByID returns the book for the given ID
func (b *BooksRepositoryImpl) GetBookByID(id string) (entities.Book, error) {
	var book entities.Book
	dbRes := b.DB.Where("id = ?", id).First(&book)
	if dbRes.Error != nil {
		emptyBook := entities.Book{}
		return emptyBook, errors.New("Unable to fetch data")
	}
	return book, nil
}
