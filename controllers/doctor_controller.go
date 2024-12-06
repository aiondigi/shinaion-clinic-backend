package controllers

import (
	"net/http"

	"github.com/aiondigi/shinaion-clinic-backend/database"
	"github.com/aiondigi/shinaion-clinic-backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllDoctors(c *gin.Context) {
	var doctors []models.Doctor
	result := database.DB.Find(&doctors)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, doctors)
}

func CreateDoctor(c *gin.Context) {
	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&doctor)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, doctor)
}

func GetDoctorByID(c *gin.Context) {
	id := c.Param("id")
	var doctor models.Doctor

	result := database.DB.First(&doctor, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	c.JSON(http.StatusOK, doctor)
}

func UpdateDoctor(c *gin.Context) {
	id := c.Param("id")
	var doctor models.Doctor

	// Check if doctor exists
	if err := database.DB.First(&doctor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	// Bind updated data
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save updates
	database.DB.Save(&doctor)
	c.JSON(http.StatusOK, doctor)
}

func DeleteDoctor(c *gin.Context) {
	id := c.Param("id")
	var doctor models.Doctor

	// Check if doctor exists
	if err := database.DB.First(&doctor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	database.DB.Delete(&doctor)
	c.JSON(http.StatusOK, gin.H{"message": "Doctor deleted successfully"})
}
