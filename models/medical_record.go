package models

import (
	"time"

	"gorm.io/gorm"
)

type MedicalRecord struct {
	ID           string         `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PatientID    string         `gorm:"type:uuid;not null"`
	DoctorID     string         `gorm:"type:uuid;not null"`
	RecordNumber string         `gorm:"type:varchar(50);not null;unique"`
	Subjective   string         `gorm:"type:text;not null"`
	Objective    string         `gorm:"type:text;not null"`
	Assessment   string         `gorm:"type:text;not null"`
	Plan         string         `gorm:"type:text;not null"`
	Attachments  *string        `gorm:"type:text"` // Store as comma-separated URLs or paths
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"` // Soft delete

	// Relationships
	Patient Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE"`
	Doctor  Doctor  `gorm:"foreignKey:DoctorID;constraint:OnDelete:CASCADE"`
}
