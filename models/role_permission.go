package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RolePermission struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;default:"uuid-ossp";primary_key" json:"id"`
	RoleID       uuid.UUID `gorm:"type:uuid" json:"role_id"`
	PermissionID uuid.UUID `gorm:"type:uuid" json:"role_id"`
}
