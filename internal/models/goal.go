package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Goal struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`

	UserID uuid.UUID `json:"user_id"`
	User   User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Name          string    `json:"name"`
	TargetAmount  float64   `json:"target_amount"`
	CurrentAmount float64   `json:"current_amount"`
	Deadline      time.Time `json:"deadline"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}

func (g *Goal) BeforeCreate(tx *gorm.DB) (err error) {
	g.ID = uuid.New()
	return
}
