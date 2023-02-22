package shared

import (
	"time"
)

// ProfitAndLossReport
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
// Our [Enhanced Financials](https://docs.codat.io/docs/assess-enhanced-financials) endpoints provide the same report under standardized headings, allowing you to pull it in the same format for all of your business customers.
type ProfitAndLossReport struct {
	CostOfSales        *Assets    `json:"costOfSales,omitempty"`
	Expenses           *Assets    `json:"expenses,omitempty"`
	FromDate           *time.Time `json:"fromDate,omitempty"`
	GrossProfit        float64    `json:"grossProfit"`
	Income             *Assets    `json:"income,omitempty"`
	NetOperatingProfit float64    `json:"netOperatingProfit"`
	NetOtherIncome     float64    `json:"netOtherIncome"`
	NetProfit          float64    `json:"netProfit"`
	OtherExpenses      *Assets    `json:"otherExpenses,omitempty"`
	OtherIncome        *Assets    `json:"otherIncome,omitempty"`
	ToDate             *time.Time `json:"toDate,omitempty"`
}
