package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"search/pkg"
)

func SearchHandler(c *gin.Context) {
	keyword := c.Param("keyword")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No keyword provided"})
		return
	}

	filePaths, err := pkg.GetFilePaths("examples")
	if err != nil {
		log.Println("Error searching for files:", err)
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
}
