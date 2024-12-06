package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	ID              string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	PatientID       string    `gorm:"not null"`
	DoctorID        string    `gorm:"not null"`
	AppointmentDate time.Time `gorm:"not null"`
	Status          string    `gorm:"type:varchar(20);not null"`
	Notes           string    `gorm:"type:text"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Patient         Patient        `gorm:"foreignKey:PatientID"`
	Doctor          Doctor         `gorm:"foreignKey:DoctorID"`
}
