package main

import (
	"github.com/hidayatarg/go-crud/initalizers"
	"github.com/hidayatarg/go-crud/models"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDatabase()
}

func main() {
	initalizers.DB.AutoMigrate(&models.Post{})
}
