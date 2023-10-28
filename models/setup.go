package models

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {

	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	if dbConnectionString == "" {
		log.Fatal("'DB_CONNECTION_STRING' environment variable is not defined.")
	}

	log.Println(dbConnectionString)
	database, err := gorm.Open(mysql.Open(dbConnectionString), &gorm.Config{})

	if err != nil {
		panic("Database connection failed.")
	}

	err = database.AutoMigrate(&Session{})

	if err != nil {
		return
	}

	DB = database

}
