package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	Name          string     `gorm:"type:varchar(255);not null" json:"name"`
	Email         string     `gorm:"uniqueIndex;not null" json:"email"`
	Password      string     `gorm:"not null" json:"password"`
	IsDeleted     bool       `gorm:"default:false" json:"is_deleted"`
	IsDeactivated bool       `gorm:"default:false" json:"is_deactivated"`
	DeletedAt     *time.Time `json:"deleted_at"`
	DeactivatedAt *time.Time `json:"deactivated_at"`
	RoleID        uuid.UUID  `gorm:"type:uuid" json:"role_id"`
}
