package controllers

import (
	"net/http"

	"github.com/aiondigi/shinaion-clinic-backend/database"
	"github.com/aiondigi/shinaion-clinic-backend/models"
	"github.com/gin-gonic/gin"
)

func GetAllMedicalRecords(c *gin.Context) {
	var medicalRecords []models.MedicalRecord
	result := database.DB.Find(&medicalRecords)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, medicalRecords)
}

func CreateMedicalRecord(c *gin.Context) {
	var medicalRecord models.MedicalRecord
	if err := c.ShouldBindJSON(&medicalRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&medicalRecord)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, medicalRecord)
}

func GetMedicalRecordByID(c *gin.Context) {
	id := c.Param("id")
	var medicalRecord models.MedicalRecord

	result := database.DB.First(&medicalRecord, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medical record not found"})
		return
	}

	c.JSON(http.StatusOK, medicalRecord)
}

func UpdateMedicalRecord(c *gin.Context) {
	id := c.Param("id")
	var medicalRecord models.MedicalRecord

	// Check if medical record exists
	if err := database.DB.First(&medicalRecord, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medical record not found"})
		return
	}

	// Bind updated data
	if err := c.ShouldBindJSON(&medicalRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save updates
	database.DB.Save(&medicalRecord)
	c.JSON(http.StatusOK, medicalRecord)
}

func DeleteMedicalRecord(c *gin.Context) {
	id := c.Param("id")
	var medicalRecord models.MedicalRecord

	// Check if medical record exists
	if err := database.DB.First(&medicalRecord, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medical record not found"})
		return
	}

	database.DB.Delete(&medicalRecord)
	c.JSON(http.StatusOK, gin.H{"message": "Medical record deleted successfully"})
}
