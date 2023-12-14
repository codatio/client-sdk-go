// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type LoanSummaryReportItem struct {
	// The loan outstanding balance.  This may not equal totalDrawdowns - totalRepayments due to interest which has been accrued.
	Balance *decimal.Big `decimal:"number" json:"balance,omitempty"`
	// The description of the object being referred to. E.g. the account.
	Description *string `json:"description,omitempty"`
	// The name of lender providing the loan.
	LenderName *string               `json:"lenderName,omitempty"`
	RecordRef  *LoanSummaryRecordRef `json:"recordRef,omitempty"`
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
	StartDate *string `json:"startDate,omitempty"`
	// The total loan drawdowns.
	TotalDrawdowns *decimal.Big `decimal:"number" json:"totalDrawdowns,omitempty"`
	// The total loan repayments which includes capital plus any interest.
	TotalRepayments *decimal.Big `decimal:"number" json:"totalRepayments,omitempty"`
}

func (l LoanSummaryReportItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoanSummaryReportItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoanSummaryReportItem) GetBalance() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *LoanSummaryReportItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *LoanSummaryReportItem) GetLenderName() *string {
	if o == nil {
		return nil
	}
	return o.LenderName
}

func (o *LoanSummaryReportItem) GetRecordRef() *LoanSummaryRecordRef {
	if o == nil {
		return nil
	}
	return o.RecordRef
}

func (o *LoanSummaryReportItem) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *LoanSummaryReportItem) GetTotalDrawdowns() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalDrawdowns
}

func (o *LoanSummaryReportItem) GetTotalRepayments() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalRepayments
}
