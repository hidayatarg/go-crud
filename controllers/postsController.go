package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hidayatarg/go-crud/initalizers"
	"github.com/hidayatarg/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	// get data from request body

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// creat post
	post := models.Post{Title: body.Title, Body: body.Body}

	// add to table
	result := initalizers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return response
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
