// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type AccountBalance struct {
	// The account's current balance
	Available *decimal.Big `decimal:"number" json:"available,omitempty"`
	// The currency of the account
	Currency *string `json:"currency,omitempty"`
	// Funds that are not yet available in the balance
	Pending *decimal.Big `decimal:"number" json:"pending,omitempty"`
	// Funds reserved as holdings
	Reserved interface{} `json:"reserved,omitempty"`
}

func (a AccountBalance) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountBalance) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountBalance) GetAvailable() *decimal.Big {
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

func (o *AccountBalance) GetPending() *decimal.Big {
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
