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
			c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Нет ключего слова"})
			return
		}

		filePaths, err := pkg.GetFilePaths("examples")
		if err != nil {
			log.Println("Ошибка поиска файлов:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
			return
		}

		jsonData := pkg.Search(filePaths, keyword)

		data, err := json.Marshal(jsonData)
		if err != nil {
			log.Println("Error marshaling JSON:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": string(data)})
	})

	router.Run(":8080")
}
