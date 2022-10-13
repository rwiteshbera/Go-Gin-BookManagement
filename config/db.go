package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	DB = ConnectDB()
	return DB
}

func ConnectDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	dataSourceName := os.Getenv("DATABASE_DSN")

	fmt.Println("dsn: " + dataSourceName)
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		log.Println("Error connecting to database.")
		return nil
	}
	return db
}
