package models

import (
	"time"

	"github.com/google/uuid"
)

type AccountType string

const (
	CURRENT AccountType = "CURRENT"
	SAVING  AccountType = "SAVING"
)

type Account struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`

	Name      string      `json:"name"`
	Type      AccountType `json:"type"`
	Balance   float64     `gorm:"default:0;" json:"balance"`
	IsDefault bool        `gorm:"default:false;" json:"is_default"`

	UserID uuid.UUID ` gorm:"index" json:"user_id"`
	User   User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Transactions []Transaction `json:"transactions"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
