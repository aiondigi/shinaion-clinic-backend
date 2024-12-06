// backend/database/seed.go
package database

import (
	"log"
	"time"

	"github.com/aiondigi/shinaion-clinic-backend/models"
	"gorm.io/gorm"
)

// Seed function to populate the database with initial data
func Seed(db *gorm.DB) {
	// Check if the patients table is empty
	var count int64
	db.Model(&models.Patient{}).Count(&count)
	if count > 0 {
		log.Println("Database already seeded with patients.")
		return
	}

	// Create seed data
	patients := []models.Patient{
		{FullName: "John Doe", DateOfBirth: time.Now().AddDate(-30, 0, 0), Gender: "Male"},
		{FullName: "Jane Smith", DateOfBirth: time.Now().AddDate(-25, 0, 0), Gender: "Female"},
		{FullName: "Alice Johnson", DateOfBirth: time.Now().AddDate(-40, 0, 0), Gender: "Female"},
		{FullName: "Bob Brown", DateOfBirth: time.Now().AddDate(-35, 0, 0), Gender: "Male"},
	}

	// Insert seed data into the database
	for _, patient := range patients {
		if err := db.Create(&patient).Error; err != nil {
			log.Printf("Error seeding patient: %v", err)
		} else {
			log.Printf("Seeded patient: %v", patient)
		}
	}

	// Create seed data for doctors
	var doctorCount int64
	db.Model(&models.Doctor{}).Count(&doctorCount)

	var doctors []models.Doctor // Declare doctors variable here

	if doctorCount == 0 {
		doctors = []models.Doctor{
			{FullName: "Dr. ENT", Specialization: ptr("ENT")},
		}

		// Insert seed data into the database
		for _, doctor := range doctors {
			if err := db.Create(&doctor).Error; err != nil {
				log.Printf("Error seeding doctor: %v", err)
			} else {
				log.Printf("Seeded doctor: %v", doctor)
			}
		}
	}

	// Create seed data for appointments
	var appointmentCount int64
	db.Model(&models.Appointment{}).Count(&appointmentCount)
	if appointmentCount == 0 {
		appointments := []models.Appointment{
			{DoctorID: doctors[0].ID, PatientID: patients[0].ID, AppointmentDate: time.Now(), Status: "Scheduled", Notes: ""},
			{DoctorID: doctors[0].ID, PatientID: patients[1].ID, AppointmentDate: time.Now(), Status: "Scheduled", Notes: ""},
			{DoctorID: doctors[0].ID, PatientID: patients[2].ID, AppointmentDate: time.Now(), Status: "Scheduled", Notes: ""},
			{DoctorID: doctors[0].ID, PatientID: patients[3].ID, AppointmentDate: time.Now(), Status: "Scheduled", Notes: ""},
		}

		// Insert seed data into the database
		for _, appointment := range appointments {
			if err := db.Create(&appointment).Error; err != nil {
				log.Printf("Error seeding appointment: %v", err)
			} else {
				log.Printf("Seeded appointment: %v", appointment)
			}
		}
	}

	// Create seed data for medical records
	medicalRecords := []models.MedicalRecord{
		{PatientID: patients[0].ID, DoctorID: doctors[0].ID, RecordNumber: "MR-001", Subjective: "Ear infection", Objective: "", Assessment: "", Plan: "Prescribed antibiotics", Attachments: nil},
		{PatientID: patients[1].ID, DoctorID: doctors[0].ID, RecordNumber: "MR-002", Subjective: "Sinusitis", Objective: "", Assessment: "", Plan: "Prescribed decongestants", Attachments: nil},
		{PatientID: patients[2].ID, DoctorID: doctors[0].ID, RecordNumber: "MR-003", Subjective: "Tonsillitis", Objective: "", Assessment: "", Plan: "Prescribed painkillers", Attachments: nil},
		{PatientID: patients[3].ID, DoctorID: doctors[0].ID, RecordNumber: "MR-004", Subjective: "Hearing loss", Objective: "", Assessment: "", Plan: "Referred for hearing test", Attachments: nil},
	}

	// Insert seed data into the database
	for _, medicalRecord := range medicalRecords {
		if err := db.Create(&medicalRecord).Error; err != nil {
			log.Printf("Error seeding medical record: %v", err)
		} else {
			log.Printf("Seeded medical record: %v", medicalRecord)
		}
	}
}

// ResetSeed function to clear existing data and reseed the database
func ResetSeed(db *gorm.DB) {
	// Clear existing data
	if err := db.Exec("DELETE FROM medical_records").Error; err != nil {
		log.Printf("Error clearing medical records: %v", err)
	}
	if err := db.Exec("DELETE FROM appointments").Error; err != nil {
		log.Printf("Error clearing appointments: %v", err)
	}
	if err := db.Exec("DELETE FROM doctors").Error; err != nil {
		log.Printf("Error clearing doctors: %v", err)
	}
	if err := db.Exec("DELETE FROM patients").Error; err != nil {
		log.Printf("Error clearing patients: %v", err)
	}

	// Call the Seed function to insert initial data
	Seed(db)
}

// Helper function to convert string to *string
func ptr(s string) *string {
	return &s
}
