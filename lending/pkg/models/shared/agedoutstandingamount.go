// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v4/pkg/types"
)

type AgedOutstandingAmount struct {
	// The amount outstanding.
	Amount *types.Decimal `json:"amount,omitempty"`
	// Array of details.
	Details []AgedOutstandingAmountDetail `json:"details,omitempty"`
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
	FromDate *string `json:"fromDate,omitempty"`
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
	ToDate *string `json:"toDate,omitempty"`
}

func (o *AgedOutstandingAmount) GetAmount() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *AgedOutstandingAmount) GetDetails() []AgedOutstandingAmountDetail {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *AgedOutstandingAmount) GetFromDate() *string {
	if o == nil {
		return nil
	}
	return o.FromDate
}

func (o *AgedOutstandingAmount) GetToDate() *string {
	if o == nil {
		return nil
	}
	return o.ToDate
}
