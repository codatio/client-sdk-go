// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v5/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// DataIntegrityAmounts - Only returned for transactions. For accounts, there is nothing returned.
type DataIntegrityAmounts struct {
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// Highest value of transaction set.
	Max *decimal.Big `decimal:"number" json:"max,omitempty"`
	// Lowest value of transaction set.
	Min *decimal.Big `decimal:"number" json:"min,omitempty"`
}

func (d DataIntegrityAmounts) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DataIntegrityAmounts) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DataIntegrityAmounts) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *DataIntegrityAmounts) GetMax() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Max
}

func (o *DataIntegrityAmounts) GetMin() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Min
}
