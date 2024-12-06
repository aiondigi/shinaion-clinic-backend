package controllers

import (
	"net/http"

	"github.com/aiondigi/shinaion-clinic-backend/database"
	"github.com/aiondigi/shinaion-clinic-backend/models"
	"github.com/gin-gonic/gin"
)

func GetAllAppointments(c *gin.Context) {
	var appointments []models.Appointment
	result := database.DB.Find(&appointments)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, appointments)
}

func CreateAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&appointment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, appointment)
}

func GetAppointmentByID(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment

	result := database.DB.First(&appointment, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func UpdateAppointment(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment

	// Check if appointment exists
	if err := database.DB.First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	// Bind updated data
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save updates
	database.DB.Save(&appointment)
	c.JSON(http.StatusOK, appointment)
}

func DeleteAppointment(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment

	// Check if appointment exists
	if err := database.DB.First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	database.DB.Delete(&appointment)
	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted successfully"})
}
