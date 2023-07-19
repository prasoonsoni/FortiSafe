package models

import (
	"time"

	"github.com/google/uuid"
)

type AccountStatusLogs struct {
	ID        uuid.UUID `gorm:"type:uuid;default:"uuid-ossp";primary_key" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;" json:"user_id"`
	Status    string    `gorm:"not null" json:"status"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}
