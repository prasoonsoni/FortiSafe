package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:"uuid-ossp";primary_key" json:"id"`
	Name          string     `gorm:"type:varchar(255);not null" json:"name"`
	Email         string     `gorm:"uniqueIndex;not null" json:"email"`
	Password      string     `gorm:"not null" json:"password"`
	CreatedAt     time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"not null" json:"updated_at"`
	IsDeleted     bool       `gorm:"default:false" json:"is_deleted"`
	IsDeactivated bool       `gorm:"default:false" json:"is_deactivated"`
	DeletedAt     *time.Time `json:"deleted_at"`
	DeactivatedAt *time.Time `json:"deactivated_at"`
}
