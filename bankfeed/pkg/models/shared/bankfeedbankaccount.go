package shared

import (
	"time"
)

type BankFeedBankAccount struct {
	AccountName   *string    `json:"accountName,omitempty"`
	AccountNumber *string    `json:"accountNumber,omitempty"`
	AccountType   *string    `json:"accountType,omitempty"`
	Balance       *float64   `json:"balance,omitempty"`
	Currency      *string    `json:"currency,omitempty"`
	FeedStartDate *time.Time `json:"feedStartDate,omitempty"`
	ID            *string    `json:"id,omitempty"`
	ModifiedDate  *time.Time `json:"modifiedDate,omitempty"`
	SortCode      *string    `json:"sortCode,omitempty"`
	Status        *string    `json:"status,omitempty"`
}
