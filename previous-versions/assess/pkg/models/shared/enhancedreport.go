// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type ReportItem struct {
	AccountCategory *EnhancedReportAccountCategory `json:"accountCategory,omitempty"`
	// The unique account ID.
	AccountID *string `json:"accountId,omitempty"`
	// Name of the account.
	AccountName *string `json:"accountName,omitempty"`
	// Balance of the account as reported on the profit and loss or Balance sheet.
	Balance *decimal.Big `decimal:"number" json:"balance,omitempty"`
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
}

func (r ReportItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ReportItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ReportItem) GetAccountCategory() *EnhancedReportAccountCategory {
	if o == nil {
		return nil
	}
	return o.AccountCategory
}

func (o *ReportItem) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *ReportItem) GetAccountName() *string {
	if o == nil {
		return nil
	}
	return o.AccountName
}

func (o *ReportItem) GetBalance() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *ReportItem) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

type EnhancedReport struct {
	ReportInfo *EnhancedReportInfo `json:"reportInfo,omitempty"`
	// An array of report items.
	ReportItems []ReportItem `json:"reportItems,omitempty"`
}

func (o *EnhancedReport) GetReportInfo() *EnhancedReportInfo {
	if o == nil {
		return nil
	}
	return o.ReportInfo
}

func (o *EnhancedReport) GetReportItems() []ReportItem {
	if o == nil {
		return nil
	}
	return o.ReportItems
}
