package shared

import (
	"time"
)

type BalanceSheetResponse struct {
	Currency                 string         `json:"currency"`
	EarliestAvailableMonth   *time.Time     `json:"earliestAvailableMonth,omitempty"`
	MostRecentAvailableMonth *time.Time     `json:"mostRecentAvailableMonth,omitempty"`
	Reports                  []BalanceSheet `json:"reports"`
}
