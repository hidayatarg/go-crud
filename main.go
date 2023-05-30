package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hidayatarg/go-crud/controllers"
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
	r.GET("/", controllers.Ping)
	r.POST("/create", controllers.PostCreate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
