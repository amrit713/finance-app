package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID     uuid.UUID  ` gorm:"type:uuid;primaryKey" json:"id"`
	Name   string     `json:"name"`
	Icon   string     `json:"icon"`
	Color  string     `json:"color"`
	UserID *uuid.UUID `json:"user_id"`

	User *User `gorm:"constraint:OnUpdade:CASCADE,OnDelete:SET NULL;"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
