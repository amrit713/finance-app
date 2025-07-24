package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Budget struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`

	UserID uuid.UUID ` gorm:"unique;not null;" json:"user_id"`
	User   *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Amount        float64    `gorm:"not null;" json:"amount"`
	LastAlertSend *time.Time `json:"last_alert_send"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (b *Budget) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}
