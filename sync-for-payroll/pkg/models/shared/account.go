// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/sync-for-payroll/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type AccountMetadata struct {
	// Indicates whether the record has been deleted in the third-party system this record originated from.
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

func (o *AccountMetadata) GetIsDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.IsDeleted
}

// AccountValidDataTypeLinks - When querying Codat's data model, some data types return `validDatatypeLinks` metadata in the JSON response. This indicates where that object can be used as a reference—a _valid link_—when creating or updating other data.
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
	// Supported `dataTypes` that the record can be linked to.
	Links []string `json:"links,omitempty"`
	// The property from the account that can be linked.
	Property *string `json:"property,omitempty"`
}

func (o *AccountValidDataTypeLinks) GetLinks() []string {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *AccountValidDataTypeLinks) GetProperty() *string {
	if o == nil {
		return nil
	}
	return o.Property
}

// Account - > **Language tip:** Accounts are also referred to as **chart of accounts**, **nominal accounts**, and **general ledger**.
//
// View the coverage for accounts in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=chartOfAccounts" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Accounts are the categories a business uses to record accounting transactions. From the Accounts endpoints, you can retrieve a list of all accounts for a specified company.
//
// The categories for an account include:
//   - Asset
//   - Expense
//   - Income
//   - Liability
//   - Equity.
//
// The same account may have a different category based on the integration it is used in. For example, a current account (known as checking in the US) should be categorized as `Asset.Current` for Xero, and `Asset.Bank.Checking` for QuickBooks Online.
//
// At the same time, each integration may have its own requirements to the categories. For example, a Paypal account in Xero is of the `Asset.Bank` category and therefore requires additional properties to be provided.
//
// To determine the list of allowed categories for a specific integration, you can:
// - Follow our [Create, update, delete data](https://docs.codat.io/using-the-api/push) guide and use the [Get create account model](https://docs.codat.io/sync-for-payroll-api#/operations/get-create-chartOfAccounts-model).
// - Refer to the integration's own documentation.
//
// > **Accounts with no category**
// >
// > If an account is pulled from the chart of accounts and its nominal code does not lie within the category layout for the company's accounts, then the **type** is `Unknown`. The **fullyQualifiedCategory** and **fullyQualifiedName** fields return `null`.
// >
// > This approach gives a true representation of the company's accounts whilst preventing distorting financials such as a company's profit and loss and balance sheet reports.
type Account struct {
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// Current balance in the account.
	CurrentBalance *decimal.Big `decimal:"number" json:"currentBalance,omitempty"`
	// Description for the account.
	Description *string `json:"description,omitempty"`
	// Full category of the account.
	//
	// For example, `Liability.Current` or `Income.Revenue`. To determine a list of possible categories for each integration, see our examples, follow our [Create, update, delete data](https://docs.codat.io/using-the-api/push) guide, or refer to the integration's own documentation.
	FullyQualifiedCategory *string `json:"fullyQualifiedCategory,omitempty"`
	// Full name of the account, for example:
	// - `Cash On Hand`
	// - `Rents Held In Trust`
	// - `Fixed Asset`
	FullyQualifiedName *string `json:"fullyQualifiedName,omitempty"`
	// Identifier for the account, unique for the company.
	ID *string `json:"id,omitempty"`
	// Confirms whether the account is a bank account or not.
	IsBankAccount *bool            `json:"isBankAccount,omitempty"`
	Metadata      *AccountMetadata `json:"metadata,omitempty"`
	ModifiedDate  *string          `json:"modifiedDate,omitempty"`
	// Name of the account.
	Name *string `json:"name,omitempty"`
	// Reference given to each nominal account for a business. It ensures money is allocated to the correct account. This code isn't a unique identifier in the Codat system.
	NominalCode        *string `json:"nominalCode,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Status of the account
	Status *AccountStatus `json:"status,omitempty"`
	// Type of account
	Type *AccountType `json:"type,omitempty"`
	// The validDatatypeLinks can be used to determine whether an account can be correctly mapped to another object; for example, accounts with a `type` of `income` might only support being used on an Invoice and Direct Income. For more information, see [Valid Data Type Links](/sync-for-payroll-api#/schemas/ValidDataTypeLinks).
	ValidDatatypeLinks []AccountValidDataTypeLinks `json:"validDatatypeLinks,omitempty"`
}

func (a Account) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Account) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Account) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *Account) GetCurrentBalance() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrentBalance
}

func (o *Account) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Account) GetFullyQualifiedCategory() *string {
	if o == nil {
		return nil
	}
	return o.FullyQualifiedCategory
}

func (o *Account) GetFullyQualifiedName() *string {
	if o == nil {
		return nil
	}
	return o.FullyQualifiedName
}

func (o *Account) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Account) GetIsBankAccount() *bool {
	if o == nil {
		return nil
	}
	return o.IsBankAccount
}

func (o *Account) GetMetadata() *AccountMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *Account) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *Account) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Account) GetNominalCode() *string {
	if o == nil {
		return nil
	}
	return o.NominalCode
}

func (o *Account) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *Account) GetStatus() *AccountStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Account) GetType() *AccountType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Account) GetValidDatatypeLinks() []AccountValidDataTypeLinks {
	if o == nil {
		return nil
	}
	return o.ValidDatatypeLinks
}
