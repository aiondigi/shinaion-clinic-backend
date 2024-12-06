package main

import (
	"log"

	"github.com/aiondigi/shinaion-clinic-backend/config"
	"github.com/aiondigi/shinaion-clinic-backend/database"
	"github.com/aiondigi/shinaion-clinic-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Connect to the database
	database.InitDB()

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupAppointmentRoutes(router)
	routes.SetupMedicalRecordRoutes(router)
	routes.SetupPatientRoutes(router)
	routes.SetupDoctorRoutes(router)

	// Seed the database
	database.Seed(database.DB)

	// Start the server
	if err := router.Run(":8088"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
