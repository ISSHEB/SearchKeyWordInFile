package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"search/pkg"
)

func main() {
	router := gin.Default()

	router.GET("/search/:keyword", func(c *gin.Context) {
		keyword := c.Param("keyword")
		if keyword == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Нет ключего слова"})
			return
		}

		filePaths, err := pkg.GetFilePaths("examples")
		if err != nil {
			log.Println("Error searching files:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		jsonData := pkg.Search(filePaths, keyword)

		data, err := json.Marshal(jsonData)
		if err != nil {
			log.Println("Error marshaling JSON:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": string(data)})
	})

	router.Run(":8080")
}
