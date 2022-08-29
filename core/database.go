package core

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// https://gorm.io/docs/

func DB() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	db, err := gorm.Open(sqlite.Open(os.Getenv("DATABASE_FILE")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
