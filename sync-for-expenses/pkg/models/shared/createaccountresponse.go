// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type CreateAccountResponseMetadata struct {
	// Indicates whether the record has been deleted in the third-party system this record originated from.
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

func (o *CreateAccountResponseMetadata) GetIsDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.IsDeleted
}

// CreateAccountResponseValidDataTypeLinks - When querying Codat's data model, some data types return `validDatatypeLinks` metadata in the JSON response. This indicates where that object can be used as a reference—a _valid link_—when creating or updating other data.
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
type CreateAccountResponseValidDataTypeLinks struct {
	// Supported `dataTypes` that the record can be linked to.
	Links []string `json:"links,omitempty"`
	// The property from the account that can be linked.
	Property *string `json:"property,omitempty"`
}

func (o *CreateAccountResponseValidDataTypeLinks) GetLinks() []string {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *CreateAccountResponseValidDataTypeLinks) GetProperty() *string {
	if o == nil {
		return nil
	}
	return o.Property
}

// AccountingAccount - > **Language tip:** Accounts are also referred to as **chart of accounts**, **nominal accounts**, and **general ledger**.
//
// View the coverage for accounts in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=chartOfAccounts" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Accounts are the categories a business uses to record accounting transactions. From the Accounts endpoints, you can retrieve a list of all accounts for a specified company.
//
// The categories for an account include:
// * Asset
// * Expense
// * Income
// * Liability
// * Equity.
//
// The same account may have a different category based on the integration it is used in. For example, a current account (known as checking in the US) should be categorized as `Asset.Current` for Xero, and `Asset.Bank.Checking` for QuickBooks Online.
//
// At the same time, each integration may have its own requirements to the categories. For example, a Paypal account in Xero is of the `Asset.Bank` category and therefore requires additional properties to be provided.
//
// To determine the list of allowed categories for a specific integration, you can:
// - Follow our [Create, update, delete data](https://docs.codat.io/using-the-api/push) guide and use the [Get create account model](https://docs.codat.io/sync-for-expenses-api#/operations/get-create-chartOfAccounts-model).
// - Refer to the integration's own documentation.
//
// > **Accounts with no category**
// >
// > If an account is pulled from the chart of accounts and its nominal code does not lie within the category layout for the company's accounts, then the **type** is `Unknown`. The **fullyQualifiedCategory** and **fullyQualifiedName** fields return `null`.
// >
// > This approach gives a true representation of the company's accounts whilst preventing distorting financials such as a company's profit and loss and balance sheet reports.
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type AccountingAccount struct {
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
	IsBankAccount *bool                          `json:"isBankAccount,omitempty"`
	Metadata      *CreateAccountResponseMetadata `json:"metadata,omitempty"`
	ModifiedDate  *string                        `json:"modifiedDate,omitempty"`
	// Name of the account.
	Name *string `json:"name,omitempty"`
	// Reference given to each nominal account for a business. It ensures money is allocated to the correct account. This code isn't a unique identifier in the Codat system.
	NominalCode        *string `json:"nominalCode,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Status of the account
	Status *AccountStatus `json:"status,omitempty"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// Type of account
	Type *AccountType `json:"type,omitempty"`
	// The validDatatypeLinks can be used to determine whether an account can be correctly mapped to another object; for example, accounts with a `type` of `income` might only support being used on an Invoice and Direct Income. For more information, see [Valid Data Type Links](/sync-for-expenses-api#/schemas/ValidDataTypeLinks).
	ValidDatatypeLinks []CreateAccountResponseValidDataTypeLinks `json:"validDatatypeLinks,omitempty"`
}

func (a AccountingAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingAccount) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountingAccount) GetCurrentBalance() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrentBalance
}

func (o *AccountingAccount) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AccountingAccount) GetFullyQualifiedCategory() *string {
	if o == nil {
		return nil
	}
	return o.FullyQualifiedCategory
}

func (o *AccountingAccount) GetFullyQualifiedName() *string {
	if o == nil {
		return nil
	}
	return o.FullyQualifiedName
}

func (o *AccountingAccount) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingAccount) GetIsBankAccount() *bool {
	if o == nil {
		return nil
	}
	return o.IsBankAccount
}

func (o *AccountingAccount) GetMetadata() *CreateAccountResponseMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AccountingAccount) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *AccountingAccount) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AccountingAccount) GetNominalCode() *string {
	if o == nil {
		return nil
	}
	return o.NominalCode
}

func (o *AccountingAccount) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *AccountingAccount) GetStatus() *AccountStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *AccountingAccount) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *AccountingAccount) GetType() *AccountType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *AccountingAccount) GetValidDatatypeLinks() []CreateAccountResponseValidDataTypeLinks {
	if o == nil {
		return nil
	}
	return o.ValidDatatypeLinks
}

type CreateAccountResponse struct {
	// Contains a single entry that communicates which record has changed and the manner in which it changed.
	Changes []PushOperationChange `json:"changes,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID string `json:"companyId"`
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
	CompletedOnUtc *string            `json:"completedOnUtc,omitempty"`
	Data           *AccountingAccount `json:"data,omitempty"`
	// Unique identifier for a company's data connection.
	DataConnectionKey string `json:"dataConnectionKey"`
	// Available data types
	DataType *DataType `json:"dataType,omitempty"`
	// A message about the error.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// A unique identifier generated by Codat to represent this single push operation. This identifier can be used to track the status of the push, and should be persisted.
	PushOperationKey string `json:"pushOperationKey"`
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
	RequestedOnUtc string `json:"requestedOnUtc"`
	// The current status of the push operation.
	Status PushOperationStatus `json:"status"`
	// Push status code.
	StatusCode int64 `json:"statusCode"`
	// Number of minutes the push operation must complete within before it times out.
	TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
	// Number of seconds the push operation must complete within before it times out.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	TimeoutInSeconds *int `json:"timeoutInSeconds,omitempty"`
	// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
	Validation *Validation `json:"validation,omitempty"`
}

func (o *CreateAccountResponse) GetChanges() []PushOperationChange {
	if o == nil {
		return nil
	}
	return o.Changes
}

func (o *CreateAccountResponse) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateAccountResponse) GetCompletedOnUtc() *string {
	if o == nil {
		return nil
	}
	return o.CompletedOnUtc
}

func (o *CreateAccountResponse) GetData() *AccountingAccount {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *CreateAccountResponse) GetDataConnectionKey() string {
	if o == nil {
		return ""
	}
	return o.DataConnectionKey
}

func (o *CreateAccountResponse) GetDataType() *DataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *CreateAccountResponse) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateAccountResponse) GetPushOperationKey() string {
	if o == nil {
		return ""
	}
	return o.PushOperationKey
}

func (o *CreateAccountResponse) GetRequestedOnUtc() string {
	if o == nil {
		return ""
	}
	return o.RequestedOnUtc
}

func (o *CreateAccountResponse) GetStatus() PushOperationStatus {
	if o == nil {
		return PushOperationStatus("")
	}
	return o.Status
}

func (o *CreateAccountResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAccountResponse) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

func (o *CreateAccountResponse) GetTimeoutInSeconds() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInSeconds
}

func (o *CreateAccountResponse) GetValidation() *Validation {
	if o == nil {
		return nil
	}
	return o.Validation
}
