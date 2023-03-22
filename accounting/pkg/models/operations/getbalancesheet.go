// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"time"
)

type GetBalanceSheetRequest struct {
	CompanyID        string     `pathParam:"style=simple,explode=false,name=companyId"`
	PeriodLength     int        `queryParam:"style=form,explode=true,name=periodLength"`
	PeriodsToCompare int        `queryParam:"style=form,explode=true,name=periodsToCompare"`
	StartMonth       *time.Time `queryParam:"style=form,explode=true,name=startMonth"`
}

type GetBalanceSheet200ApplicationJSONBalanceSheetReportLineReportLineReportLineReportLine struct {
	// Identifier for the account, unique for the company in the accounting platform.
	AccountID *string `json:"accountId,omitempty"`
	// Name of the report line item.
	Name *string `json:"name,omitempty"`
	// Numerical value of the line item.
	Value float64 `json:"value"`
}

type GetBalanceSheet200ApplicationJSONBalanceSheetReportLineReportLineReportLine struct {
	// Identifier for the account, unique for the company in the accounting platform.
	AccountID *string `json:"accountId,omitempty"`
	// An array of ReportLine items.
	Items []GetBalanceSheet200ApplicationJSONBalanceSheetReportLineReportLineReportLineReportLine `json:"items,omitempty"`
	// Name of the report line item.
	Name *string `json:"name,omitempty"`
	// Numerical value of the line item.
	Value float64 `json:"value"`
}

type GetBalanceSheet200ApplicationJSONBalanceSheetReportLineReportLine struct {
	// Identifier for the account, unique for the company in the accounting platform.
	AccountID *string `json:"accountId,omitempty"`
	// An array of ReportLine items.
	Items []GetBalanceSheet200ApplicationJSONBalanceSheetReportLineReportLineReportLine `json:"items,omitempty"`
	// Name of the report line item.
	Name *string `json:"name,omitempty"`
	// Numerical value of the line item.
	Value float64 `json:"value"`
}

// GetBalanceSheet200ApplicationJSONBalanceSheetReportLine - ReportLines for assets. For example, fixed and current assets.
type GetBalanceSheet200ApplicationJSONBalanceSheetReportLine struct {
	// Identifier for the account, unique for the company in the accounting platform.
	AccountID *string `json:"accountId,omitempty"`
	// An array of ReportLine items.
	Items []GetBalanceSheet200ApplicationJSONBalanceSheetReportLineReportLine `json:"items,omitempty"`
	// Name of the report line item.
	Name *string `json:"name,omitempty"`
	// Numerical value of the line item.
	Value float64 `json:"value"`
}

// GetBalanceSheet200ApplicationJSONBalanceSheet - > View the coverage for balance sheet in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=balanceSheet" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// The balance sheet gives interested parties an idea of the company's financial position, in addition to displaying what the company owns and owes.
//
// > **Balance sheet or profit and loss report?**
// >
// > A profit and loss report summarises the total revenue, expenses, and profit or loss during a specified time period. A balance sheet report shows the financial position of a company at a specific moment in time.
//
// **Structure of this report**
// This report will reflect the structure and line descriptions that the business has set in their own accounting platform.
//
// **History**
// By default, Codat pulls (up to) 24 months of balance sheets for a company. You can adjust this to fetch more history, where available, by updating the `monthsToSync` value for `balanceSheet` on the [data type settings endpoint](https://docs.codat.io/codat-api#/operations/post-profile-syncSettings).
//
// **Want to pull this in a standardised structure?**
// Our [Enhanced Financials](https://docs.codat.io/assess/reports/enhanced-financials/financials) endpoints provide the same report under standardized headings, allowing you to pull it in the same format for all of your business customers.
type GetBalanceSheet200ApplicationJSONBalanceSheet struct {
	// ReportLines for assets. For example, fixed and current assets.
	Assets *GetBalanceSheet200ApplicationJSONBalanceSheetReportLine `json:"assets,omitempty"`
	// Point in time when a snapshot of a company's financial position is taken.
	Date *time.Time `json:"date,omitempty"`
	// ReportLines for equities. For example, retained and current year earnings. See below.
	Equity *GetBalanceSheet200ApplicationJSONBalanceSheetReportLine `json:"equity,omitempty"`
	// ReportLines for liabilities. For example, current liabilities.
	Liabilities *GetBalanceSheet200ApplicationJSONBalanceSheetReportLine `json:"liabilities,omitempty"`
	// Value of net assets for a company in their base currency.
	NetAssets float64 `json:"netAssets"`
}

// GetBalanceSheet200ApplicationJSON - Success
type GetBalanceSheet200ApplicationJSON struct {
	// Currency of the balance sheet.
	Currency string `json:"currency"`
	// Earliest available monthly report data.
	EarliestAvailableMonth *time.Time `json:"earliestAvailableMonth,omitempty"`
	// Most recent available monthly report data.
	MostRecentAvailableMonth *time.Time `json:"mostRecentAvailableMonth,omitempty"`
	// An array of BalanceSheet reports.
	Reports []GetBalanceSheet200ApplicationJSONBalanceSheet `json:"reports"`
}

type GetBalanceSheetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	GetBalanceSheet200ApplicationJSONObject *GetBalanceSheet200ApplicationJSON
}
