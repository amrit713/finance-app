package dto

import (
	"time"

	"github.com/amirt713/finance-app/internal/models"
)

type TransactionInput struct {
	CategoryID  string  `json:"category_id"`
	AccountID   string  `json:"account_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`

	Type              models.TransactionType        `json:"type"`
	ReceiptURL        *string                       `json:"receipt_url"`
	IsRecurring       *bool                         ` json:"is_recurring"`
	RecurringInterval *models.RecurringIntervalType `json:"recurring_interval"`
	NextRecurringDate *time.Time                    `json:"next_recurring_date"`
	LastProcessed     *time.Time                    `json:"last_processed"`
	Status            *models.StatusType            ` json:"status"`
}

type UpdateTransactionInput struct {
	CategoryID  *string  `json:"category_id"` //preload is required
	Amount      *float64 `json:"amount"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Date        *string  `json:"date"`

	Type              *models.TransactionType       `json:"type"`
	ReceiptURL        *string                       `json:"receipt_url"`
	IsRecurring       *bool                         ` json:"is_recurring"`
	RecurringInterval *models.RecurringIntervalType `json:"recurring_interval"`
	NextRecurringDate *time.Time                    `json:"next_recurring_date"`
	LastProcessed     *time.Time                    `json:"last_processed"`
	Status            *models.StatusType            ` json:"status"`
}
