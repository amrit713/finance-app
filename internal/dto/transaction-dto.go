package dto

import (
	"time"

	"github.com/amirt713/finance-app/internal/models"
	"github.com/google/uuid"
)

type TransactionInput struct {
	CategoryID  uuid.UUID `json:"category_id"`
	AccountID   uuid.UUID `json:"account_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`

	Type              models.TransactionType        `json:"type"`
	ReceiptURL        *string                       `json:"receipt_url"`
	IsRecurring       *bool                         ` json:"is_recurring"`
	RecurringInterval *models.RecurringIntervalType `json:"recurring_interval"`
	NextRecurringDate *time.Time                    `json:"next_recurring_date"`
	LastProcessed     *time.Time                    `json:"last_processed"`
	Status            *models.StatusType            ` json:"status"`
}

type UpdateTransactionInput struct {
	CategoryID  *uuid.UUID `json:"category_id"` //preload is required
	Amount      *float64   `json:"amount"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	Date        *time.Time `json:"date"`

	Type              *models.TransactionType       `json:"type"`
	ReceiptURL        *string                       `json:"receipt_url"`
	IsRecurring       *bool                         ` json:"is_recurring"`
	RecurringInterval *models.RecurringIntervalType `json:"recurring_interval"`
	NextRecurringDate *time.Time                    `json:"next_recurring_date"`
	LastProcessed     *time.Time                    `json:"last_processed"`
	Status            *models.StatusType            ` json:"status"`
}
