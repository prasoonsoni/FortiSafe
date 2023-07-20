package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID          uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `gorm:"not null" json:"description"`
	Permissions pq.StringArray `gorm:"type:text[]" json:"permissions"`
}
