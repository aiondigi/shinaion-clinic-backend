package routes

import (
	"github.com/aiondigi/shinaion-clinic-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupPatientRoutes(router *gin.Engine) {
	patients := router.Group("/api/patients")
	{
		patients.GET("/", controllers.GetAllPatients)
		patients.POST("/", controllers.CreatePatient)
		patients.GET("/:id", controllers.GetPatientByID)
		patients.PUT("/:id", controllers.UpdatePatient)
		patients.DELETE("/:id", controllers.DeletePatient)
	}
}
