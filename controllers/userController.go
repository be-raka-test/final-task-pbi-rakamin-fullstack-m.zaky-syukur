package controllers

import (
	"btpn-go/app/models"
	"btpn-go/config"
	"btpn-go/helpers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register user
func Register(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := helpers.HashPassword(input.Password)
	if err != nil {
		helpers.JSONError(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := models.User{Email: input.Email, Password: hashedPassword}
	config.DB.Create(&user)

	token := helpers.GenerateToken(user.ID)
	helpers.JSONResponse(c, http.StatusOK, "User registered successfully", gin.H{"token": token})
}

// Login user
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	err := config.DB.Where("email = ?", input.Email).First(&user).Error
	if err != nil {
		log.Printf("Error finding user with email %s: %v", input.Email, err)
		helpers.JSONError(c, http.StatusUnauthorized, "User not found")
		return
	}

	if !helpers.CheckPasswordHash(input.Password, user.Password) {
		helpers.JSONError(c, http.StatusUnauthorized, "b")
		return
	}

	token := helpers.GenerateToken(user.ID)
	helpers.JSONResponse(c, http.StatusOK, "User logged in successfully", gin.H{"token": token})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}
