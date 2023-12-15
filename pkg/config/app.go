package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {

	// Replace placeholders with your actual MySQL details
	username := "ayush1503"
	password := "ayush1503@12"
	host := "127.0.0.1"
	port := "3306"
	dbName := "simplerest"

	// Create a connection string with proper encoding
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, dbName)

	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
	defer d.Close()
}

func GetDB() *gorm.DB {
	return db
}
