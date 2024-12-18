// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v8/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type DataIntegrityByAmount struct {
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// The percentage of the absolute value of transactions of the type specified in the route which have a match.
	MatchPercentage *decimal.Big `decimal:"number" json:"matchPercentage,omitempty"`
	// The sum of the absolute value of transactions of the type specified in the route which have a match.
	Matched *decimal.Big `decimal:"number" json:"matched,omitempty"`
	// The total of unmatched and matched.
	Total *decimal.Big `decimal:"number" json:"total,omitempty"`
	// The sum of the absolute value of transactions of the type specified in the route which don't have a match.
	Unmatched *decimal.Big `decimal:"number" json:"unmatched,omitempty"`
}

func (d DataIntegrityByAmount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DataIntegrityByAmount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DataIntegrityByAmount) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *DataIntegrityByAmount) GetMatchPercentage() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.MatchPercentage
}

func (o *DataIntegrityByAmount) GetMatched() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Matched
}

func (o *DataIntegrityByAmount) GetTotal() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *DataIntegrityByAmount) GetUnmatched() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Unmatched
}
