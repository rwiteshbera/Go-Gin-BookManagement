package models

import (
	"github.com/rwiteshbera/gin_golang/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `gorm:""json:"author"`
	Publication string `gorm:""json:"publication"`
}

func init() {
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func CreateBook(b *Book) *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book // Type slice Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}
