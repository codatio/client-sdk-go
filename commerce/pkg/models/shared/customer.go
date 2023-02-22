package shared

import (
	"time"
)

// Customer
// Represents a customer who has placed an order in the commerce system"
type Customer struct {
	Addresses          []Items                `json:"addresses,omitempty"`
	CreatedDate        *time.Time             `json:"createdDate,omitempty"`
	CustomerName       *string                `json:"customerName,omitempty"`
	DefaultCurrency    map[string]interface{} `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                `json:"emailAddress,omitempty"`
	ID                 string                 `json:"id"`
	ModifiedDate       *time.Time             `json:"modifiedDate,omitempty"`
	Note               *string                `json:"note,omitempty"`
	Phone              *string                `json:"phone,omitempty"`
	SourceModifiedDate *time.Time             `json:"sourceModifiedDate,omitempty"`
}
