package shared

import (
	"time"
)

// AccountBalanceBalance
// An object containing bank balance data.
type AccountBalanceBalance struct {
	Available *float64 `json:"available,omitempty"`
	Current   float64  `json:"current"`
	Limit     *float64 `json:"limit,omitempty"`
}

// AccountBalance
// The Banking Account Balances data type provides a list of balances for a bank account including end-of-day batch balance or running balances per transaction.
type AccountBalance struct {
	AccountID          string                `json:"accountId"`
	Balance            AccountBalanceBalance `json:"balance"`
	Date               time.Time             `json:"date"`
	ModifiedDate       *time.Time            `json:"modifiedDate,omitempty"`
	SourceModifiedDate *time.Time            `json:"sourceModifiedDate,omitempty"`
}
