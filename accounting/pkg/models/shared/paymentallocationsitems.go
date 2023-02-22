package shared

import (
	"time"
)

type ItemsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

type ItemsPayment struct {
	AccountRef   *AccountRef `json:"accountRef,omitempty"`
	Currency     *string     `json:"currency,omitempty"`
	CurrencyRate *float64    `json:"currencyRate,omitempty"`
	ID           *string     `json:"id,omitempty"`
	Note         *string     `json:"note,omitempty"`
	PaidOnDate   *time.Time  `json:"paidOnDate,omitempty"`
	Reference    *string     `json:"reference,omitempty"`
	TotalAmount  *float64    `json:"totalAmount,omitempty"`
}

type PaymentAllocationsitems struct {
	Allocation ItemsAllocation `json:"allocation"`
	Payment    ItemsPayment    `json:"payment"`
}
