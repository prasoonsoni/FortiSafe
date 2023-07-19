package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID          uuid.UUID    `gorm:"type:uuid;default:"uuid-ossp";primary_key" json:"id"`
	Name        string       `gorm:"not null" json:"status"`
	Description string       `gorm:"not null" json:"status"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}
