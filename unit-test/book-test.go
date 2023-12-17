// book_test.go

package main

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// TestCreateBook tests the CreateBook function
func TestCreateBook(t *testing.T) {
	// Set up a connection to a test database

	dsn := "ayush:ayush1503@tcp(hostname:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Run migrations for creating necessary tables
	db.AutoMigrate(&Book{})

	// Test data
	title := "The Great Gatsby"
	author := "F. Scott Fitzgerald"
	publicationYear := 1925
	isbn := "978-3-16-148410-0"

	// Call the function being tested
	createdBook, err := CreateBook(db, title, author, publicationYear, isbn)
	if err != nil {
		t.Fatalf("CreateBook failed: %v", err)
	}

	// Verify that the book was created in the database
	var savedBook Book
	result := db.First(&savedBook, "title = ?", title)
	if result.Error != nil {
		t.Fatalf("Failed to retrieve book from the database: %v", result.Error)
	}

	// Assert that the retrieved book matches the expected values
	if savedBook.Title != title || savedBook.Author != author || savedBook.PublicationYear != publicationYear || savedBook.ISBN != isbn {
		t.Fatalf("Mismatch in book details. Expected: %+v, Got: %+v", Book{
			Title:           title,
			Author:          author,
			PublicationYear: publicationYear,
			ISBN:            isbn,
		}, savedBook)
	}

	// Add more test functions to cover other parts of the codebase...
}

// Example test function for another feature
func TestUpdateBook(t *testing.T) {
	// Similar setup as above...
}
