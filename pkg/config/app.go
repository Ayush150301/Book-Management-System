package config

import (
	// "fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {

	// username := "ayush1503"
	// password := "ayush1503@12"
	// host := "127.0.0.1"
	// port := "3306"
	// dbName := "simplerest"

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	username, password, host, port, dbName)

	dsn := "ayush:ayush1503@tcp(hostname:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

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
