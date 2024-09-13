// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// BalanceSheetInput - > View the coverage for balance sheet in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=balanceSheet" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// The balance sheet is a snapshot of a company's accounts at a single point in time that provides a statement of the assets, liabilities and equity of an organization. It gives interested parties an idea of the company's financial position, in addition to displaying what the company owns and owes.
//
// > **Balance sheet or profit and loss report?**
// >
// > A profit and loss report summarises the total revenue, expenses, and profit or loss during a specified time period. A balance sheet report shows the financial position of a company at a specific moment in time.
//
// **Structure of this report**
// This report will reflect the structure and line descriptions that the business has set in their own accounting software.
//
// **History**
// By default, Codat pulls (up to) 24 months of balance sheets for a company. You can adjust this to fetch more history, where available, by updating the `monthsToSync` value for `balanceSheet` on the [data type settings endpoint](https://docs.codat.io/codat-api#/operations/update-sync-settings).
//
// **Want to pull this in a standardised structure?**
// Our [Enhanced Financials](https://docs.codat.io/lending/features/financial-statements-overview) endpoints provide the same report under standardized headings, allowing you to pull it in the same format for all of your business customers.
type BalanceSheetInput struct {
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency string `json:"currency"`
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
	EarliestAvailableMonth *string `json:"earliestAvailableMonth,omitempty"`
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
	MostRecentAvailableMonth *string `json:"mostRecentAvailableMonth,omitempty"`
	// An array of balance sheet reports.
	Reports []BalanceSheet `json:"reports"`
}

func (o *BalanceSheetInput) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *BalanceSheetInput) GetEarliestAvailableMonth() *string {
	if o == nil {
		return nil
	}
	return o.EarliestAvailableMonth
}

func (o *BalanceSheetInput) GetMostRecentAvailableMonth() *string {
	if o == nil {
		return nil
	}
	return o.MostRecentAvailableMonth
}

func (o *BalanceSheetInput) GetReports() []BalanceSheet {
	if o == nil {
		return []BalanceSheet{}
	}
	return o.Reports
}
