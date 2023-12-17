# Book-Management-System
## Introduction
Welcome to the Book Management System built with Golang, MySQL, GORM, and Gorilla Mux. This system provides a robust solution for managing a collection of books, including features such as adding, updating, deleting, and retrieving book information from a MySQL database.

## Prerequisites
Before running the application, ensure you have the following installed on your system:

- Golang: Installation Guide
- MySQL: Installation Guide
- GORM: Golang Object Relational Mapper (GitHub)
- Gorilla Mux: A powerful URL router and dispatcher for Golang (GitHub)

## Steps
1. Make a directory on your system -->  `mkdir directory-name`
2. Change the current directory --> `cd directory-name`
3. Now run this command --> `go mod init github.com/User-Name/Book-Management-System`
4. Now make 2 folders named as cmd and pkg for better structure of the projects 
5. Get the required modules for the projects 
    - Gorm modules provide us various functions to convert the code into mysql querys --> `go get "github.com/jinzhu/gorm"`
    - Used for database Integration `go get "github.com/jinzhu/gorm/dialects/mysql"`
    - install marshal and unmarshall `go get "github.com/gorilla/mux"`
6. Inside the 'cmd' folder create a main folder in which main.go file will be created.
7. Create a few folders inside pkg folders for better project structuring 
    - config -> app.go(connection of database)
    - controllers --> book-controller.go(for CRUD operations)
    - model --> book.go(for creating books model)
    - routes--> bookstore-routes.go(routes of all the functions)
    - utils --> utils.go(use of marshal to store inside the database)

## Run the Application
1. Go in the directory where main.go file is present
2. To start the server, execute the below commands from the terminal:
    - `go build`
    - `go run main.go`


## API ENDPOINTS
| Method | Endpoint                   | Description                           | Access Control    |
|--------|----------------------------|---------------------------------------|-------------------|
| GET    | /books                     | Returns list of all books             | Public            |
| POST   | /books                     | Add new book                          | Private (JWT)     |
| GET    | /books/{bookid}            | Return details of a specific book     | Public            |
| PUT    | /books/{bookid}/update     | Update existing book                  | Private (JWT)     |
| DELETE | /books/{bookid}/delete     | Delete an existing book               | Private (JWT)     |
</s>

## Create a Book
- Endpoint
`POST http://localhost:9010/books`
- Method: POST
- Request Body
  ```
  {
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "publication_year": 1925,
  "isbn": "978-3-16-148410-0"
  }
  ```
- Response Body
  ```
  {
  "id": 1,
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "publication_year": 1925,
  "isbn": "978-3-16-148410-0"
  }
  ```
## GetAll Books

- Endpoint
`GET http://localhost:9010/books`
- Method: GETALL
- Response Body
  ```
  [
    {
      "id": 1,
      "title": "The Great Gatsby",
      "author": "F. Scott Fitzgerald",
      "publication_year": 1925,
      "isbn": "978-3-16-148410-0"
    },
  // Other books...
  ]
  ```

## Get A Single Book
- Endpoint
`GET http://localhost:9010/books/:bookId`
- Method: GET
- URL Parameter : `bookId` is the id of the book you want to get.
- No request body needed.
- Response Body
  ```
  {
    "id": 1,
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "publication_year": 1925,
    "isbn": "978-3-16-148410-0"
  }
  ```

## Update a Book
- Endpoint
`PUT http://localhost:9010/books/:bookId`
- Method: PUT
- URL Parameter : `bookId` is the id of the book you want to update.
- Request Body (application/json)
  ```
  {
    "title": "New Title",
    "author": "New Author",
    "publication_year": 2023,
    "isbn": "123-456-789"
  }
  ```
- Response Body
  ```
  {
  "id": 1,
  "title": "New Title",
  "author": "New Author",
  "publication_year": 2023,
  "isbn": "123-456-789"
  }
  ```

## Delete a Book
- Endpoint
`DELETE http://localhost:9010/books/:bookId`
- Method: DELETE
- URL Parameter : `bookId` is the id of the book you want to delete.
- No request body needed.
- Response Code: 204 - No Content</s>


## Known Issues
- Make sure that the datbase server is connected
- Steps to check the database server is connected or not:-
  - On Windows: Use the Services application to check if the MySQL service is running
  - Test Database Connection in Code: If you have a Golang application using GORM to connect to the database, you can add a simple test in your code to check the database connection. For example:
    ```
      package main

      import (
      	"fmt"
      	"log"

      	"gorm.io/driver/mysql"
      	"gorm.io/gorm"
      )

      func main() {
      	dsn := "username:password@tcp(hostname:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
      	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
      	if err != nil {
      		log.Fatal("Database connection failed:", err)
      	}
      	fmt.Println("Connected to the database!")
      }

    ```