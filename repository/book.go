package repository

import (
	"books-api/entity"

	"github.com/jinzhu/gorm"
)

type BookRepository interface {
	InsertBook(b entity.Book) entity.Book
	UpdateBook(b entity.Book) entity.Book
	DeleteBook(b entity.Book)
	AllBook() []entity.Book
	FindBookByID(bookID uint64) entity.Book
}

type bookConnection struct {
	connection *gorm.DB
}

func NewBookRepository(dbConn *gorm.DB) BookRepository {
	return &bookConnection{
		connection: dbConn,
	}
}

func (db *bookConnection) InsertBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *bookConnection) UpdateBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *bookConnection) DeleteBook(b entity.Book) {
	db.connection.Delete(&b)
}

func (db *bookConnection) FindBookByID(bookID uint64) entity.Book {
	var book entity.Book
	db.connection.Find(&book, bookID)
	return book
}

func (db *bookConnection) AllBook() []entity.Book {
	var books []entity.Book
	db.connection.Find(&books)
	return books
}
