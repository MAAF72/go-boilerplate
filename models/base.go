package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base base model
type Base struct {
	ID        string          `json:"id" gorm:"primary_key" uri:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// BeforeCreate set ID of model
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewString()
	base.ID = uuid
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
