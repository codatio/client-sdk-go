package operations

import (
	"net/http"
	"time"
)

type GetProfitAndLossPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetProfitAndLossQueryParams struct {
	PeriodLength     int        `queryParam:"style=form,explode=true,name=periodLength"`
	PeriodsToCompare int        `queryParam:"style=form,explode=true,name=periodsToCompare"`
	StartMonth       *time.Time `queryParam:"style=form,explode=true,name=startMonth"`
}

type GetProfitAndLossRequest struct {
	PathParams  GetProfitAndLossPathParams
	QueryParams GetProfitAndLossQueryParams
}

type GetProfitAndLoss200ApplicationJSONReportBasisEnum string

const (
	GetProfitAndLoss200ApplicationJSONReportBasisEnumUnknown GetProfitAndLoss200ApplicationJSONReportBasisEnum = "Unknown"
	GetProfitAndLoss200ApplicationJSONReportBasisEnumAccrual GetProfitAndLoss200ApplicationJSONReportBasisEnum = "Accrual"
	GetProfitAndLoss200ApplicationJSONReportBasisEnumCash    GetProfitAndLoss200ApplicationJSONReportBasisEnum = "Cash"
)

type GetProfitAndLoss200ApplicationJSONProfitAndLossReportReportLineReportLineReportLineReportLine struct {
	AccountID *string `json:"accountId,omitempty"`
	Name      *string `json:"name,omitempty"`
	Value     float64 `json:"value"`
}

type GetProfitAndLoss200ApplicationJSONProfitAndLossReportReportLineReportLineReportLine struct {
	AccountID *string                                                                                         `json:"accountId,omitempty"`
	Items     []GetProfitAndLoss200ApplicationJSONProfitAndLossReportReportLineReportLineReportLineReportLine `json:"items,omitempty"`
	Name      *string                                                                                         `json:"name,omitempty"`
	Value     float64                                                                                         `json:"value"`
}

type GetProfitAndLoss200ApplicationJSONProfitAndLossReportReportLineReportLine struct {
	AccountID *string                                                                               `json:"accountId,omitempty"`
	Items     []GetProfitAndLoss200ApplicationJSONProfitAndLossReportReportLineReportLineReportLine `json:"items,omitempty"`
	Name      *string                                                                               `json:"name,omitempty"`
	Value     float64                                                                               `json:"value"`
}

// GetProfitAndLoss200ApplicationJSONProfitAndLossReportReportLine
// ReportLine items for cost of sales in the given date range.
type GetProfitAndLoss200ApplicationJSONProfitAndLossReportReportLine struct {
	AccountID *string                                                                     `json:"accountId,omitempty"`
	Items     []GetProfitAndLoss200ApplicationJSONProfitAndLossReportReportLineReportLine `json:"items,omitempty"`
	Name      *string                                                                     `json:"name,omitempty"`
	Value     float64                                                                     `json:"value"`
}

// GetProfitAndLoss200ApplicationJSONProfitAndLossReport
// > **Language tip:** Profit and loss statement is also referred to as **income statement** under US GAAP (Generally Accepted Accounting Principles).
//
// > View the coverage for profit and loss in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=profitAndLoss" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// The purpose of a profit and loss report is to present the financial performance of a company over a specified time period.
//
// A profit and loss report shows a company's total income and expenses for a specified period of time and whether a profit or loss has been made.
//
// > **Profit and loss or balance sheet?**
// > Profit and loss reports summarise the total revenue, expenses, and profit or loss over a specified time period. A balance sheet report presents all assets, liability, and equity for a given date.
//
// **Structure of this report**
// This report will reflect the structure and line descriptions that the business has set in their own accounting platform.
//
// **History**
// By default, Codat pulls (up to) 24 months of profit and loss history for a company. You can adjust this to fetch more history, where available, by updating the `monthsToSync` value for `profitAndLoss` on the [data type settings endpoint](https://docs.codat.io/codat-api#/operations/post-profile-syncSettings).
//
// **Want to pull this in a standardised structure?**
// Our [Enhanced Financials](https://docs.codat.io/assess/reports/enhanced-financials/financials) endpoints provide the same report under standardized headings, allowing you to pull it in the same format for all of your business customers.
type GetProfitAndLoss200ApplicationJSONProfitAndLossReport struct {
	CostOfSales        *GetProfitAndLoss200ApplicationJSONProfitAndLossReportReportLine `json:"costOfSales,omitempty"`
	Expenses           *GetProfitAndLoss200ApplicationJSONProfitAndLossReportReportLine `json:"expenses,omitempty"`
	FromDate           *time.Time                                                       `json:"fromDate,omitempty"`
	GrossProfit        float64                                                          `json:"grossProfit"`
	Income             *GetProfitAndLoss200ApplicationJSONProfitAndLossReportReportLine `json:"income,omitempty"`
	NetOperatingProfit float64                                                          `json:"netOperatingProfit"`
	NetOtherIncome     float64                                                          `json:"netOtherIncome"`
	NetProfit          float64                                                          `json:"netProfit"`
	OtherExpenses      *GetProfitAndLoss200ApplicationJSONProfitAndLossReportReportLine `json:"otherExpenses,omitempty"`
	OtherIncome        *GetProfitAndLoss200ApplicationJSONProfitAndLossReportReportLine `json:"otherIncome,omitempty"`
	ToDate             *time.Time                                                       `json:"toDate,omitempty"`
}

type GetProfitAndLoss200ApplicationJSON struct {
	Currency                 string                                                  `json:"currency"`
	EarliestAvailableMonth   *time.Time                                              `json:"earliestAvailableMonth,omitempty"`
	MostRecentAvailableMonth *time.Time                                              `json:"mostRecentAvailableMonth,omitempty"`
	ReportBasis              GetProfitAndLoss200ApplicationJSONReportBasisEnum       `json:"reportBasis"`
	Reports                  []GetProfitAndLoss200ApplicationJSONProfitAndLossReport `json:"reports"`
}

type GetProfitAndLossResponse struct {
	ContentType                              string
	StatusCode                               int
	RawResponse                              *http.Response
	GetProfitAndLoss200ApplicationJSONObject *GetProfitAndLoss200ApplicationJSON
}
