package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/high-ping-devs/simple-chat-room/backend/auth"
	"github.com/high-ping-devs/simple-chat-room/backend/database"
	"github.com/high-ping-devs/simple-chat-room/backend/models"
)

type UserLogin struct {
	Email    string `json:"email" validate:"nonzero,min=3,max=255,email"`
	Password string `json:"password" validate:"nonzero,min=12,max=255"`
}

func Login(c *gin.Context) {
	var user models.User
	var wl auth.WhiteList
	c.BindJSON(&user)

	// Hack to bypass validation on login (username is not required)
	if user.Username == "" {
		user.Username = "<<temp>>"
	}

	if err := user.Validate(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	password := user.Password

	result := database.DB.Where("email = ?", user.Email).First(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	if !user.PasswordMatch(password) {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	claims := jwt.MapClaims{
		"permissions": "user",
		"id":          user.ID,
		"username":    user.Username,
	}

	token, err := auth.GenerateToken(5, claims)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := auth.GenerateRefreshToken(claims, 60)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ID := strconv.FormatUint(uint64(user.ID), 10)

	wl.Create()
	defer wl.Close()
	wl.HSet(c, ID, "token", token)
	wl.HSet(c, ID, "refreshToken", refreshToken)

	c.JSON(200, gin.H{
		"accessToken":  token,
		"refreshToken": refreshToken,
	})
}

func Register(c *gin.Context) {
	var user models.User

	c.BindJSON(&user)

	if err := user.Validate(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully"})
}

func Logout(c *gin.Context) {
	var wl auth.WhiteList

	// Hopefully will not get an error here (middleware should have already checked)
	ID, exists := c.Get("id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	ID = strconv.FormatUint(uint64(ID.(float64)), 10)

	wl.Create()
	defer wl.Close()
	wl.HDel(c, ID.(string), "token", "refreshToken")
	c.JSON(200, gin.H{"message": "User logged out successfully"})
}

func Refresh(c *gin.Context) {

}
