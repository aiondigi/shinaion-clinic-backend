package routes

import (
	"github.com/aiondigi/shinaion-clinic-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupMedicalRecordRoutes(router *gin.Engine) {
	medicalRecords := router.Group("/api/medical-records")
	{
		medicalRecords.GET("/", controllers.GetAllMedicalRecords)
		medicalRecords.POST("/", controllers.CreateMedicalRecord)
		medicalRecords.GET("/:id", controllers.GetMedicalRecordByID)
		medicalRecords.PUT("/:id", controllers.UpdateMedicalRecord)
		medicalRecords.DELETE("/:id", controllers.DeleteMedicalRecord)
	}
}
