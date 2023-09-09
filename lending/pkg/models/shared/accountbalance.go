// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/pkg/types"
)

type AccountBalance struct {
	// The account's current balance
	Available *types.Decimal `json:"available,omitempty"`
	// The currency of the account
	Currency *string `json:"currency,omitempty"`
	// Funds that are not yet available in the balance
	Pending *types.Decimal `json:"pending,omitempty"`
	// Funds reserved as holdings
	Reserved interface{} `json:"reserved,omitempty"`
}

func (o *AccountBalance) GetAvailable() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.Available
}

func (o *AccountBalance) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountBalance) GetPending() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.Pending
}

func (o *AccountBalance) GetReserved() interface{} {
	if o == nil {
		return nil
	}
	return o.Reserved
}
