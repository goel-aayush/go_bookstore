package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// MySQL connection string
	dsn := "aayush:manugoel@tcp(localhost:3306)/simplerest?charset=utf8&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to the database: %v", err))
	}
	fmt.Println("Database connected successfully!")
}

func GetDB() *gorm.DB {
	return db
}
