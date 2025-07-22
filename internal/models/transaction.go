package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionType string

const (
	Expense TransactionType = "EXPENSE"
	Income  TransactionType = "INCOME"
)

type RecurringIntervalType string

const (
	DAILY   RecurringIntervalType = "DAILY"
	MONTHLY RecurringIntervalType = "MONTHLY"
	YEARLY  RecurringIntervalType = "YEARLY"
)

type StatusType string

const (
	COMPLETED StatusType = "COMPLETED"
	PENDING   StatusType = "PENDING"
	FAILED    StatusType = "FAILED"
)

type Transaction struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`

	UserID uuid.UUID `gorm:"index" json:"user_id"`
	User   User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CategoryID uuid.UUID `gorm:"index" json:"category_id"`
	Category   Category  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	AccountID uuid.UUID `gorm:"index" json:"account_id"`
	Account   Account   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Name              string                 `gorm:"index; not-null" json:"name"`
	Description       string                 `gorm:"type:text" json:"description"`
	Amount            float64                `json:"amount"`
	Date              time.Time              `json:"date"`
	Type              TransactionType        `json:"type"`
	ReciptURL         *string                `json:"recipt_url"`
	IsRecurring       bool                   `gorm:"default:false;" json:"is_recurring"`
	RecurringInterval *RecurringIntervalType `json:"recurring_interval"`
	NextRecurringDate *time.Time             `json:"next_recurring_date"`
	LastProcessed     *time.Time             `json:"last_processed"`
	Status            StatusType             `gorm:"default:COMPLETED" json:"status"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
