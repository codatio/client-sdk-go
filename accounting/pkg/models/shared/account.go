package shared

import (
	"time"
)

type AccountStatusEnum string

const (
	AccountStatusEnumUnknown  AccountStatusEnum = "Unknown"
	AccountStatusEnumActive   AccountStatusEnum = "Active"
	AccountStatusEnumArchived AccountStatusEnum = "Archived"
	AccountStatusEnumPending  AccountStatusEnum = "Pending"
)

type AccountTypeEnum string

const (
	AccountTypeEnumUnknown   AccountTypeEnum = "Unknown"
	AccountTypeEnumAsset     AccountTypeEnum = "Asset"
	AccountTypeEnumExpense   AccountTypeEnum = "Expense"
	AccountTypeEnumIncome    AccountTypeEnum = "Income"
	AccountTypeEnumLiability AccountTypeEnum = "Liability"
	AccountTypeEnumEquity    AccountTypeEnum = "Equity"
)

// AccountValidDataTypeLinks
// When querying Codat's data model, some data types return `validDatatypeLinks` metadata in the JSON response. This indicates where that object can be used as a reference—a _valid link_—when creating or updating other data.
//
// For example, `validDatatypeLinks` might indicate the following references:
//
// - Which tax rates are valid to use on the line item of a bill.
// - Which items can be used when creating an invoice.
//
// You can use `validDatatypeLinks` to present your SMB customers with only valid choices when selecting objects from a list, for example.
//
// ## `validDatatypeLinks` example
//
// The following example uses the `Accounting.Accounts` data type. It shows that, on the linked integration, this account is valid as the account on a payment or bill payment; and as the account referenced on the line item of a direct income or direct cost. Because there is no valid link to Invoices or Bills, using this account on those data types will result in an error.
//
// ```json validDatatypeLinks for an account
//
//	{
//	            "id": "bd9e85e0-0478-433d-ae9f-0b3c4f04bfe4",
//	            "nominalCode": "090",
//	            "name": "Business Bank Account",
//	            #...
//	            "validDatatypeLinks": [
//	                {
//	                    "property": "Id",
//	                    "links": [
//	                        "Payment.AccountRef.Id",
//	                        "BillPayment.AccountRef.Id",
//	                        "DirectIncome.LineItems.AccountRef.Id",
//	                        "DirectCost.LineItems.AccountRef.Id"
//	                    ]
//	                }
//	            ]
//	        }
//
// ```
//
// ## Support for `validDatatypeLinks`
//
// Codat currently supports `validDatatypeLinks` for some data types on our Xero, QuickBooks Online, QuickBooks Desktop, Exact (NL), and Sage Business Cloud integrations.
//
// If you'd like us to extend support to more data types or integrations, suggest or vote for this on our <a href="https://portal.productboard.com/codat/5-product-roadmap">Product Roadmap</a>.
type AccountValidDataTypeLinks struct {
	Links    []string `json:"links,omitempty"`
	Property *string  `json:"property,omitempty"`
}

// Account
// > **Language tip:** Accounts are also referred to as **chart of accounts**, **nominal accounts**, and **general ledger**.
//
// Explore the <a className="external" href="https://api.codat.io/swagger/index.html#/Accounts" target="_blank">Accounts</a> endpoints in Swagger.
//
// View the coverage for accounts in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=chartOfAccounts" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Accounts are the categories a business uses to record accounting transactions. From the Accounts endpoints, you can retrieve [a list of all accounts for a specified company](https://api.codat.io/swagger/index.html#/Accounts/Accounts_List).
//
// The categories for an account include:
//   - Asset
//   - Expense
//   - Income
//   - Liability
//   - Equity.
//
// > **Accounts with no category**
// >
// > If an account is pulled from the chart of accounts and its nominal code does not lie within the category layout for the company's accounts, then the **type** is `Unknown`. The **fullyQualifiedCategory** and **fullyQualifiedName** fields return `null`.
// >
// > This approach gives a true representation of the company's accounts whilst preventing distorting financials such as a company's profit and loss and balance sheet reports.
type Account struct {
	Currency               *string                     `json:"currency,omitempty"`
	CurrentBalance         *float64                    `json:"currentBalance,omitempty"`
	Description            *string                     `json:"description,omitempty"`
	FullyQualifiedCategory *string                     `json:"fullyQualifiedCategory,omitempty"`
	FullyQualifiedName     *string                     `json:"fullyQualifiedName,omitempty"`
	ID                     *string                     `json:"id,omitempty"`
	IsBankAccount          bool                        `json:"isBankAccount"`
	Metadata               *Metadata                   `json:"metadata,omitempty"`
	ModifiedDate           *time.Time                  `json:"modifiedDate,omitempty"`
	Name                   *string                     `json:"name,omitempty"`
	NominalCode            *string                     `json:"nominalCode,omitempty"`
	SourceModifiedDate     *time.Time                  `json:"sourceModifiedDate,omitempty"`
	Status                 AccountStatusEnum           `json:"status"`
	Type                   AccountTypeEnum             `json:"type"`
	ValidDatatypeLinks     []AccountValidDataTypeLinks `json:"validDatatypeLinks,omitempty"`
}
