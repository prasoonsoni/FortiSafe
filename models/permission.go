package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:"uuid-ossp";primary_key" json:"id"`
	Name        string    `gorm:"not null" json:"Name"`
	Description string    `gorm:"not null" json:"description"`
}
