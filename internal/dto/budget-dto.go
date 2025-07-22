package dto

import "time"

type BudgetInput struct {
	Amount float64 `json:"amount"`
}

type UpdateBudgetInput struct {
	Amount        *float64   `json:"amount"`
	LastAlertSend *time.Time `json:"last_alert_send"`
}
