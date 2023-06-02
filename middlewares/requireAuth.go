package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hidayatarg/go-crud/initalizers"
	"github.com/hidayatarg/go-crud/models"
)

func RequireAuth(c *gin.Context) {
	//fmt.Println("in middleware")

	// Get Cookie of Request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode validate it

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			// if time now is greater than expiration
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// find the user with token subject

		// attach to request
		// find user by id
		var user models.User
		initalizers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			// no user found
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// attach to request
		c.Set("user", user)

		// continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}

func RequireAdmin(c *gin.Context) {
	user, _ := c.Get("user")

	if user.(models.User).Role != "admin" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}

func RequireMember(c *gin.Context) {
	user, _ := c.Get("user")

	if user.(models.User).Role != "member" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}

func RequireAdminOrMember(c *gin.Context) {
	user, _ := c.Get("user")

	if user.(models.User).Role != "member" && user.(models.User).Role != "admin" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}
