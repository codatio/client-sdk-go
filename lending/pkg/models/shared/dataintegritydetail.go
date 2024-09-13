// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v5/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type DataIntegrityDetail struct {
	// The transaction value.
	Amount *decimal.Big `decimal:"number" json:"amount,omitempty"`
	// ID GUID representing the connection of the accounting or banking platform.
	ConnectionID *string `json:"connectionId,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	Date *string `json:"date,omitempty"`
	// The transaction description.
	Description *string `json:"description,omitempty"`
	// ID GUID of the transaction.
	ID      *string              `json:"id,omitempty"`
	Matches []DataIntegrityMatch `json:"matches,omitempty"`
	// The data type of the record.
	Type *string `json:"type,omitempty"`
}

func (d DataIntegrityDetail) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DataIntegrityDetail) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DataIntegrityDetail) GetAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *DataIntegrityDetail) GetConnectionID() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionID
}

func (o *DataIntegrityDetail) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *DataIntegrityDetail) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *DataIntegrityDetail) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *DataIntegrityDetail) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DataIntegrityDetail) GetMatches() []DataIntegrityMatch {
	if o == nil {
		return nil
	}
	return o.Matches
}

func (o *DataIntegrityDetail) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}
