package shared

import (
	"time"
)

type BalanceSheetReportLine struct {
	AccountID *string  `json:"accountId,omitempty"`
	Items     []Assets `json:"items,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Value     float64  `json:"value"`
}

// BalanceSheet
// > View the coverage for balance sheet in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=balanceSheet" target="_blank">Data coverage explorer</a>.
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
// Our [Enhanced Financials](https://docs.codat.io/docs/assess-enhanced-financials) endpoints provide the same report under standardized headings, allowing you to pull it in the same format for all of your business customers.
type BalanceSheet struct {
	Assets      *BalanceSheetReportLine `json:"assets,omitempty"`
	Date        *time.Time              `json:"date,omitempty"`
	Equity      *Assets                 `json:"equity,omitempty"`
	Liabilities *Assets                 `json:"liabilities,omitempty"`
	NetAssets   float64                 `json:"netAssets"`
}
