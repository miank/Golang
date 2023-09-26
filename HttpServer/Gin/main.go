package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!!",
		})
	})
	fmt.Println()
	r.Run()
}
