package repository

import (
	"github.com/Wordyka/golang_gin_gorm_JWT.git/entity"
	"gorm.io/gorm"
)


// Berisi modul untuk memanipulasi database

// interface BookRepository yang berisi method dengan parameter berupa nilai dari entity Buku
type BookRepository interface {
	InsertBook(b entity.Book) entity.Book
	UpdateBook(b entity.Book) entity.Book
	DeleteBook(b entity.Book)
	AllBook() []entity.Book
	FindBookID(bookID uint64) entity.Book
}

// struct bookConnection yang berisi koneksi dari database
type bookConnection struct {
	connection *gorm.DB
}

// fungsi NewBookRepository untuk membuat repository baru pada database yang telah terkoneksi
func NewBookRepository(dbConn *gorm.DB) BookRepository {
	return &bookConnection{
		connection: dbConn,
	}
}

// Implementasi interface pada method InsertBook untuk menambahakan data buku baru
func (db *bookConnection) InsertBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

// Implementasi interface pada method UpdateBook untuk mengubah data buku berdasarkan id nya
func (db *bookConnection) UpdateBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

// Implementasi interface pada method DeleteBook untuk menghapus data buku berdasarkan id nya
func (db *bookConnection) DeleteBook(b entity.Book) {
	db.connection.Delete(&b)
}

// Implementasi interface pada method FindBookID untuk mencari data buku berdasarkan id nya
func (db *bookConnection) FindBookID(bookID uint64) entity.Book {
	var book entity.Book
	db.connection.Preload("User").Find(&book, bookID)
	return book
}

// Implementasi interface pada method AllBook untuk menampilkan semua data buku
func (db *bookConnection) AllBook() []entity.Book {
	var books []entity.Book
	db.connection.Preload("User").Find(&books)
	return books
}


