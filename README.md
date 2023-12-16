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
- - Gorm modules provide us various functions to convert the code into mysql querys --> `go get "github.com/jinzhu/gorm"`
- - Used for database Integration `go get "github.com/jinzhu/gorm/dialects/mysql"`
- - install marshal and unmarshall `go get "github.com/gorilla/mux"`
6. Inside the 'cmd' folder create a main folder in which main.go file will be created.
7. Create a few folders inside pkg folders for better project structuring 
- - config -> app.go(connection of database)
- - controllers --> book-controller.go(for CRUD operations)
- - model --> book.go(for creating books model)
- - routes--> bookstore-routes.go(routes of all the functions)
- - utils --> utils.go(use of marshal to store inside the database)

## Run the Application
1. Go in the directory where main.go file is present
2. To start the server, execute the below commands from the terminal:
- - go build
- - go run main.go
