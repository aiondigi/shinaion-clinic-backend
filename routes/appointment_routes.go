package routes

import (
	"github.com/aiondigi/shinaion-clinic-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupAppointmentRoutes(router *gin.Engine) {
	appointments := router.Group("/api/appointments")
	{
		appointments.GET("/", controllers.GetAllAppointments)
		appointments.POST("/", controllers.CreateAppointment)
		appointments.GET("/:id", controllers.GetAppointmentByID)
		appointments.PUT("/:id", controllers.UpdateAppointment)
		appointments.DELETE("/:id", controllers.DeleteAppointment)
	}
}
