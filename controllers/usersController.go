package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hidayatarg/go-crud/initalizers"
	"github.com/hidayatarg/go-crud/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// get email and password from request body
	var body models.UserRequest

	if c.Bind(&body) == nil {
		models.ReturnGenericBadResponse(c, "Invalid request body 1")
		return
	}

	if (body.Email == "") || (body.Password == "") {
		models.ReturnGenericBadResponse(c, "Email and Password cannot be empty")
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		models.ReturnGenericBadResponse(c, "Failed to hash password")
		return
	}
	// create user
	user := models.User{Email: body.Email, Password: string(hash)}
	// add to table
	result := initalizers.DB.Create(&user)

	if result.Error != nil {
		models.ReturnGenericBadResponse(c, "Failed to create user")
		return
	}

	// return response
	models.ReturnGenericSuccessWithNoMessageResponse(c)
}

func Login(c *gin.Context) {
	// get email and password from request body
	var body models.UserRequest

	if c.Bind(&body) != nil {
		models.ReturnGenericBadResponse(c, "Invalid request body")
		return
	}

	if (body.Email == "") || (body.Password == "") {
		models.ReturnGenericBadResponse(c, "Email and Password cannot be empty")
		return
	}

	// look up email and password
	var user models.User
	initalizers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		// didnt find user
		models.ReturnGenericBadResponse(c, "Invalid email or password")
		return
	}

	// compare send and save user password hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		models.ReturnGenericBadResponse(c, "Invalid email or password")
		return
	}

	// generate a jwt token and send it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		models.ReturnGenericBadResponse(c, "Failed to create token")
		return
	}

	// send token
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	// c.JSON(http.StatusOK, gin.H{
	// 	// "token": tokenString,
	// })
	models.ReturnGenericSuccessWithNoMessageResponse(c)
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	models.ReturnGenericSuccessWithMessageResponse(c, nil, user.(models.User).Email)
}
