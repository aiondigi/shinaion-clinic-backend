package models

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	ID             string         `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	FullName       string         `gorm:"type:varchar(100);not null"`
	Specialization *string        `gorm:"type:varchar(100)"`
	LicenseNumber  string         `gorm:"type:varchar(50);not null"`
	SIP            string         `gorm:"type:varchar(50);not null"`
	ContactNumber  *string        `gorm:"type:varchar(15)"`
	Email          *string        `gorm:"type:varchar(100)"`
	Availability   *string        `gorm:"type:json"` // Store weekly availability as JSON
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"index"` // Soft delete
}
