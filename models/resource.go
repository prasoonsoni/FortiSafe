package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	ID               uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	Name             string         `gorm:"not null" json:"name"`
	Description      string         `gorm:"not null" json:"description"`
	AssociatedRoles  pq.StringArray `gorm:"type:text[]" json:"associated_roles"`
	AssociatedGroups pq.StringArray `gorm:"type:text[]" json:"associated_groups"`
	CreatedBy        uuid.UUID      `gorm:"type:uuid" json:"created_by"`
}
