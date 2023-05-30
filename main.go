package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hidayatarg/go-crud/initalizers"
)

// init -> run before main
func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDatabase()
}

func main() {
	fmt.Println("Hello Go")
	fmt.Println("It is working")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
