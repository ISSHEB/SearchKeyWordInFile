package main

import (
	"github.com/gin-gonic/gin"
	"search/handlers"
)

func main() {
	router := gin.Default()

	router.GET("/search/:keyword", handlers.SearchHandler)

	router.Run(":8080")
}
