package controllers

import (
	"btpn-go/app/models"
	"btpn-go/config"
	"btpn-go/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	var input struct {
		URL string `json:"url" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		helpers.JSONError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	photo := models.Photo{UserID: userId.(uint), URL: input.URL}
	config.DB.Create(&photo)

	helpers.JSONResponse(c, http.StatusOK, "Photo created successfully", photo)
}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	config.DB.Find(&photos)

	helpers.JSONResponse(c, http.StatusOK, "Photos retrieved successfully", photos)
}

func GetPhoto(c *gin.Context) {
	var photo models.Photo
	if err := config.DB.Where("id = ?", c.Param("id")).First(&photo).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Photo not found")
		return
	}

	helpers.JSONResponse(c, http.StatusOK, "Photo retrieved successfully", photo)
}

func UpdatePhoto(c *gin.Context) {
	var photo models.Photo
	if err := config.DB.Where("id = ?", c.Param("id")).First(&photo).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Photo not found")
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		helpers.JSONError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if photo.UserID != userId.(uint) {
		helpers.JSONError(c, http.StatusForbidden, "You are not authorized to update this photo")
		return
	}

	var input struct {
		URL     string `json:"url" binding:"required"`
		Caption string `json:"caption"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	photo.URL = input.URL
	config.DB.Save(&photo)

	helpers.JSONResponse(c, http.StatusOK, "Photo updated successfully", photo)
}

func DeletePhoto(c *gin.Context) {
	var photo models.Photo
	if err := config.DB.Where("id = ?", c.Param("id")).First(&photo).Error; err != nil {
		helpers.JSONError(c, http.StatusNotFound, "Photo not found")
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		helpers.JSONError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if photo.UserID != userId.(uint) {
		helpers.JSONError(c, http.StatusForbidden, "You are not authorized to delete this photo")
		return
	}

	config.DB.Delete(&photo)

	c.Status(http.StatusNoContent)
}
