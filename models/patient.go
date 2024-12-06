package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	ID            string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FullName      string    `gorm:"type:varchar(100);not null"`
	DateOfBirth   time.Time `gorm:"not null"`
	Gender        string    `gorm:"type:varchar(10);not null"`
	ContactNumber string    `gorm:"type:varchar(15)"`
	Email         string    `gorm:"type:varchar(100)"`
	Address       string    `gorm:"type:text"`
	InsuranceID   string    `gorm:"type:varchar(50)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
