package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/hidayatarg/go-crud/initalizers"
	"github.com/hidayatarg/go-crud/models"
	"gorm.io/gorm"
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

func PostsIndex(c *gin.Context) {
	// read data from db
	var posts []models.Post
	initalizers.DB.Find(&posts) // finding the posts and assigning them to posts array

	// return response
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsSingleById(c *gin.Context) {

	// Get Id from url
	id := c.Param("id")

	// read data from db
	var post models.Post
	result := initalizers.DB.First(&post, id) // finding the posts and assigning them to posts array
	err := result.Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(200, gin.H{
			"error": "Not found error",
		})
		return
	}

	// return response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdateById(c *gin.Context) {
	// Get Id from url
	id := c.Param("id")

	// get data from request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// find post we want to update
	var post models.Post
	initalizers.DB.First(&post, id)

	// update
	initalizers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// return response
	c.JSON(200, gin.H{
		"post": post,
	})
}