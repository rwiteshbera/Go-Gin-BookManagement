package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rwiteshbera/gin_golang/routes"
)

func main() {
	var r = gin.Default()
	var PORT = ":5050"

	routes.BookStoreRoutes(r)

	log.Fatal(r.Run(PORT))
}
