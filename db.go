package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

// DB returns a database connection
func DB() *gorm.DB {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@%s/%s",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
		),
	)

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

// MigrateDB creates the necessary tables for our models
func MigrateDB() {
	DB().AutoMigrate(&Relay{})
}
