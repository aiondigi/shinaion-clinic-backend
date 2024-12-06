package main

import (
	"log"

	"github.com/aiondigi/shinaion-clinic-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=clinic_user password=@Angsoka_2025 dbname=shinaion_clinic port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Enable UUID generation for PostgreSQL
	err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
	if err != nil {
		log.Fatal("Failed to enable UUID extension:", err)
	}

	// Create ENUM type before migrations
	db.Exec(`DO $$ BEGIN
		CREATE TYPE gender AS ENUM ('Male', 'Female', 'Other');
	EXCEPTION
		WHEN duplicate_object THEN null;
	END $$;`)

	// Then run migrations
	err = db.AutoMigrate(&models.Patient{}, &models.Doctor{}, &models.Appointment{}, &models.MedicalRecord{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed successfully!")
}
