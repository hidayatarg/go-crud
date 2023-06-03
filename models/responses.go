package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnGenericBadResponse(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":   err,
		"success": false,
	})
}

func ReturnGenericSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"success": true,
	})
}

func ReturnGenericSuccessWithMessageResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"success": true,
		"message": message,
	})
}

func ReturnGenericSuccessWithNoMessageResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
