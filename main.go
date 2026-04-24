package main

/*
import "fmt"

func	main() {
	fmt.Println("Hello, Vaultlink")
}
*/

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default();

	r.GET("/ping", func(c *gin.Context) {
		// JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
