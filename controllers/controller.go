package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rwiteshbera/gin_golang/models"
)

var NewBook models.Book

func GetBook(ctx *gin.Context) {
	newBooks := models.GetAllBooks()
	ctx.JSON(200, newBooks)
}

func GetBookById(ctx *gin.Context) {
	Id := ctx.Param("id")
	id, err := strconv.ParseInt(Id, 10, 64) // Parsing string type to int64
	if err != nil {
		fmt.Println("Error while parsing url")
	}
	bookDetails, _ := models.GetBookById(id)
	ctx.JSON(200, bookDetails)
}

func CreateBook(ctx *gin.Context) {
	bookData := &models.Book{}
	ctx.Bind(&bookData)
	createBook := models.CreateBook(bookData)
	ctx.JSON(200, createBook)
}

func UpdateBook(ctx *gin.Context) {
	updatedBook := &models.Book{}
	ctx.Bind(&updatedBook)

	Id := ctx.Param("id")
	id, err := strconv.ParseInt(Id, 10, 64) // Parsing string type to int64
	if err != nil {
		fmt.Println("Error while parsing url")
	}
	bookDetails, db := models.GetBookById(id)

	if updatedBook.Name != "" {
		bookDetails.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}
	db.Save(&bookDetails)
	ctx.JSON(200, bookDetails)
}

func DeleteBook(ctx *gin.Context) {
	Id := ctx.Param("id")
	id, err := strconv.ParseInt(Id, 10, 64) // Parsing string type to int64
	if err != nil {
		fmt.Println("Error while parsing url")
	}
	deletedBook := models.DeleteBook(id)
	ctx.JSON(200, deletedBook)
}
