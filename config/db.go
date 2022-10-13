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

	USERNAME := os.Getenv("DB_USER")
	PASSWORD := os.Getenv("DB_PASSWORD")
	HOSTNAME := os.Getenv("DB_HOSTNAME")
	PORT := os.Getenv("DB_PORT")
	DATABASE := os.Getenv("DB_NAME")
	
	dataSourceName := USERNAME + ":" + PASSWORD + "@tcp(" + HOSTNAME + ":" + PORT + ")/" + DATABASE + "?charset=utf8&parseTime=true&loc=Local"

	fmt.Println("dsn: " + dataSourceName)
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		log.Println("Error connecting to database.")
		return nil
	}
	return db
}
