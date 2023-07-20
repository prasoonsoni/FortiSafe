package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"not null" json:"description"`
}
