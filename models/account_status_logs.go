package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountStatusLogs struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;" json:"user_id"`
	Status string    `gorm:"not null" json:"status"`
}
