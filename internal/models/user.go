package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Name     string    `gorm:"not null;" json:"name"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Password string    `gorm:"not null;" json:"-"`
	Image    string    `json:"image,omitempty"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Transactions []Transaction `json:"transactions"`
	Accounts     []Account     `json:"accounts"`
	Budgets      []Budget      `json:"budgets"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
