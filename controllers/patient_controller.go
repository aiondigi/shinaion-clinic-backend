package controllers

import (
	"net/http"

	"github.com/aiondigi/shinaion-clinic-backend/database"
	"github.com/aiondigi/shinaion-clinic-backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllPatients(c *gin.Context) {
	var patients []models.Patient
	result := database.DB.Find(&patients)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, patients)
}

func CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&patient)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, patient)
}

func GetPatientByID(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	result := database.DB.First(&patient, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	// Check if patient exists
	if err := database.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// Bind updated data
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save updates
	database.DB.Save(&patient)
	c.JSON(http.StatusOK, patient)
}

func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	// Check if patient exists
	if err := database.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	database.DB.Delete(&patient)
	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}
