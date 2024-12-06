package routes

import (
	"github.com/aiondigi/shinaion-clinic-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupDoctorRoutes(router *gin.Engine) {
	doctors := router.Group("/api/doctors")
	{
		doctors.GET("/", controllers.GetAllDoctors)
		doctors.POST("/", controllers.CreateDoctor)
		doctors.GET("/:id", controllers.GetDoctorByID)
		doctors.PUT("/:id", controllers.UpdateDoctor)
		doctors.DELETE("/:id", controllers.DeleteDoctor)
	}
}
