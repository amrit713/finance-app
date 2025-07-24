package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountType string

const (
	CURRENT AccountType = "CURRENT"
	SAVING  AccountType = "SAVING"
)

type Account struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`

	Name      string      `gorm:"not null;" json:"name"`
	Type      AccountType `gorm:"not null;" json:"type"`
	Balance   float64     `gorm:"default:0;" json:"balance"`
	IsDefault bool        `gorm:"default:false;" json:"is_default"`

	UserID uuid.UUID ` gorm:"index not null" json:"user_id"`
	User   *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user "`

	Transactions []Transaction `json:"transactions"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
