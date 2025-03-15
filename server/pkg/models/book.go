package models

import (
	"github.com/goel-aayush/go_bookstore/server/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	// Initialize the database connection and auto-migrate the Book struct
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// CreateBook inserts a new book record into the database
func (b *Book) CreateBook() *Book {
	// Create the record in the database
	if err := db.Create(&b).Error; err != nil {
		// Handle any errors that occur during the creation
		return nil
	}
	return b
}

// GetAllBooks fetches all the books from the database
func GetAllBooks() []Book {
	var books []Book
	// Fetch all books from the database
	if err := db.Find(&books).Error; err != nil {
		// Handle error if necessary
		return nil
	}
	return books
}

// GetBookById retrieves a book by its ID from the database
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	// Find a book by its ID
	result := db.Where("ID = ?", Id).First(&book)
	if result.Error != nil {
		// If no record is found or any error occurs, return nil
		return nil, result
	}
	return &book, result
}

// DeleteBook deletes a book by its ID from the database
func DeleteBook(ID int64) (*Book, error) {
	var book Book
	// Try to find the book by its ID
	result := db.Where("ID = ?", ID).First(&book)

	if result.Error != nil {
		// If no book is found or another error occurs, return an error
		return nil, result.Error
	}

	// Attempt to delete the book
	result = db.Delete(&book)

	// Check if there was an error while deleting
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}
