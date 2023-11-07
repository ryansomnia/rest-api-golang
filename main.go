package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Heriyanto Sitorus",
			"bio":  "A software engineer",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Heriyanto Sitorus",
			"say":  "Hello gaisss Im a software engineer",
		})
	})
	router.Run(":2371")
}
