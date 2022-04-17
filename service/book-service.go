package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/dto"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/entity"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/repository"
)

// interface BookService sebagai contract mengenai service yang dibuat pada data book
type BookService interface {
	Insert(b dto.BookCreateDTO) entity.Book
	Update(b dto.BookUpdateDTO) entity.Book
	Delete(b entity.Book)
	All() []entity.Book
	FIndById(bookID uint64) entity.Book
	IsAllowedToEdit(userID string, bookID uint64) bool
}


// struct bookService untuk menginstansiasi interface dari BookRepository dari package repository
type bookService struct {
	bookRepository repository.BookRepository
}

// fungsi NewBookService untuk memanggil struct BookService dengan menginisialisasi variabel address userRepository
func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}

// Implementasi interface pada method Insert untuk menambahkan data book baru
func (service *bookService) Insert(b dto.BookCreateDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.bookRepository.InsertBook(book)
	return res
}

// Implementasi interface pada method Update untuk mengubah data book
func (service *bookService) Update(b dto.BookUpdateDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.bookRepository.UpdateBook(book)
	return res
}


// Implementasi interface pada method Delete untuk menghapus data book 
func (service *bookService) Delete(b entity.Book) {
	service.bookRepository.DeleteBook(b)
}

// Implementasi interface pada method All untuk mendapatkan semua data book 
func (service *bookService) All() []entity.Book {
	return service.bookRepository.AllBook()
}


// Implementasi interface pada method FindById untuk mecari data book berdasarkan id pada entity book
func (service *bookService) FIndById(bookID uint64) entity.Book {
	return service.bookRepository.FindBookID(bookID)
}


// Implementasi interface pada method IsAllowedToEdit untuk memberi hak data book diubah berdasarkan id
func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	b := service.bookRepository.FindBookID(bookID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}






