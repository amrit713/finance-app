package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	UserID uuid.UUID `json:"user_id"`
	User   User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Type    string `json:"type"`
	Message string `json:"message"`
	Read    bool   `gorm:"type:boolean;defult:false;" json:"read" `

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (n *Notification) BeforeCreate(tx *gorm.DB) (err error) {
	n.ID = uuid.New()
	return
}
