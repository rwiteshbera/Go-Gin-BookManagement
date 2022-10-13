package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rwiteshbera/gin_golang/controllers"
)

var BookStoreRoutes = func(router *gin.Engine) {
	router.GET("/book", controllers.GetBook)
	router.GET("/book/:id", controllers.GetBookById)
	router.POST("/book", controllers.CreateBook)
	router.PUT("/book/:id", controllers.UpdateBook)
	router.DELETE("/book/:id", controllers.DeleteBook)
}
