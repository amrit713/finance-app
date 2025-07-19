package dto

import "github.com/amirt713/finance-app/internal/models"

type AccountInput struct {
	Name      string             `json:"name"`
	Type      models.AccountType `json:"type"`
	Balance   float64            ` json:"balance"`
	IsDefault *bool              ` json:"is_default"`
}

type UpdateAccountInput struct {
	Name      *string             `json:"name"`
	Type      *models.AccountType `json:"type"`
	Balance   *float64            `json:"balance"`
	IsDefault *bool               `json:"is_default"`
}
