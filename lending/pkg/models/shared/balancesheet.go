// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v6/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type BalanceSheet struct {
	Assets *ReportLine `json:"assets,omitempty"`
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
	Date        *string     `json:"date,omitempty"`
	Equity      *ReportLine `json:"equity,omitempty"`
	Liabilities *ReportLine `json:"liabilities,omitempty"`
	// Value of net assets for a company in their base currency.
	NetAssets *decimal.Big `decimal:"number" json:"netAssets"`
}

func (b BalanceSheet) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BalanceSheet) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BalanceSheet) GetAssets() *ReportLine {
	if o == nil {
		return nil
	}
	return o.Assets
}

func (o *BalanceSheet) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *BalanceSheet) GetEquity() *ReportLine {
	if o == nil {
		return nil
	}
	return o.Equity
}

func (o *BalanceSheet) GetLiabilities() *ReportLine {
	if o == nil {
		return nil
	}
	return o.Liabilities
}

func (o *BalanceSheet) GetNetAssets() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.NetAssets
}
