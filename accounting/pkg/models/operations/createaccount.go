package operations

import (
	"net/http"
	"time"
)

type CreateAccountSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateAccountSourceModifiedDateStatusEnum string

const (
	CreateAccountSourceModifiedDateStatusEnumUnknown  CreateAccountSourceModifiedDateStatusEnum = "Unknown"
	CreateAccountSourceModifiedDateStatusEnumActive   CreateAccountSourceModifiedDateStatusEnum = "Active"
	CreateAccountSourceModifiedDateStatusEnumArchived CreateAccountSourceModifiedDateStatusEnum = "Archived"
	CreateAccountSourceModifiedDateStatusEnumPending  CreateAccountSourceModifiedDateStatusEnum = "Pending"
)

type CreateAccountSourceModifiedDateTypeEnum string

const (
	CreateAccountSourceModifiedDateTypeEnumUnknown   CreateAccountSourceModifiedDateTypeEnum = "Unknown"
	CreateAccountSourceModifiedDateTypeEnumAsset     CreateAccountSourceModifiedDateTypeEnum = "Asset"
	CreateAccountSourceModifiedDateTypeEnumExpense   CreateAccountSourceModifiedDateTypeEnum = "Expense"
	CreateAccountSourceModifiedDateTypeEnumIncome    CreateAccountSourceModifiedDateTypeEnum = "Income"
	CreateAccountSourceModifiedDateTypeEnumLiability CreateAccountSourceModifiedDateTypeEnum = "Liability"
	CreateAccountSourceModifiedDateTypeEnumEquity    CreateAccountSourceModifiedDateTypeEnum = "Equity"
)

// CreateAccountSourceModifiedDateValidDataTypeLinks
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
type CreateAccountSourceModifiedDateValidDataTypeLinks struct {
	Links    []string `json:"links,omitempty"`
	Property *string  `json:"property,omitempty"`
}

// CreateAccountSourceModifiedDate
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
type CreateAccountSourceModifiedDate struct {
	Currency               *string                                             `json:"currency,omitempty"`
	CurrentBalance         *float64                                            `json:"currentBalance,omitempty"`
	Description            *string                                             `json:"description,omitempty"`
	FullyQualifiedCategory *string                                             `json:"fullyQualifiedCategory,omitempty"`
	FullyQualifiedName     *string                                             `json:"fullyQualifiedName,omitempty"`
	ID                     *string                                             `json:"id,omitempty"`
	IsBankAccount          bool                                                `json:"isBankAccount"`
	Metadata               *CreateAccountSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate           *time.Time                                          `json:"modifiedDate,omitempty"`
	Name                   *string                                             `json:"name,omitempty"`
	NominalCode            *string                                             `json:"nominalCode,omitempty"`
	SourceModifiedDate     *time.Time                                          `json:"sourceModifiedDate,omitempty"`
	Status                 CreateAccountSourceModifiedDateStatusEnum           `json:"status"`
	Type                   CreateAccountSourceModifiedDateTypeEnum             `json:"type"`
	ValidDatatypeLinks     []CreateAccountSourceModifiedDateValidDataTypeLinks `json:"validDatatypeLinks,omitempty"`
}

type CreateAccountRequest struct {
	RequestBody      *CreateAccountSourceModifiedDate `request:"mediaType=application/json"`
	CompanyID        string                           `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string                           `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes *int                             `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreateAccount200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateAccount200ApplicationJSONChangesTypeEnum string

const (
	CreateAccount200ApplicationJSONChangesTypeEnumUnknown            CreateAccount200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateAccount200ApplicationJSONChangesTypeEnumCreated            CreateAccount200ApplicationJSONChangesTypeEnum = "Created"
	CreateAccount200ApplicationJSONChangesTypeEnumModified           CreateAccount200ApplicationJSONChangesTypeEnum = "Modified"
	CreateAccount200ApplicationJSONChangesTypeEnumDeleted            CreateAccount200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateAccount200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateAccount200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type CreateAccount200ApplicationJSONChanges struct {
	AttachmentID *string                                                       `json:"attachmentId,omitempty"`
	RecordRef    *CreateAccount200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateAccount200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type CreateAccount200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateAccount200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	CreateAccount200ApplicationJSONSourceModifiedDateStatusEnumUnknown  CreateAccount200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	CreateAccount200ApplicationJSONSourceModifiedDateStatusEnumActive   CreateAccount200ApplicationJSONSourceModifiedDateStatusEnum = "Active"
	CreateAccount200ApplicationJSONSourceModifiedDateStatusEnumArchived CreateAccount200ApplicationJSONSourceModifiedDateStatusEnum = "Archived"
	CreateAccount200ApplicationJSONSourceModifiedDateStatusEnumPending  CreateAccount200ApplicationJSONSourceModifiedDateStatusEnum = "Pending"
)

type CreateAccount200ApplicationJSONSourceModifiedDateTypeEnum string

const (
	CreateAccount200ApplicationJSONSourceModifiedDateTypeEnumUnknown   CreateAccount200ApplicationJSONSourceModifiedDateTypeEnum = "Unknown"
	CreateAccount200ApplicationJSONSourceModifiedDateTypeEnumAsset     CreateAccount200ApplicationJSONSourceModifiedDateTypeEnum = "Asset"
	CreateAccount200ApplicationJSONSourceModifiedDateTypeEnumExpense   CreateAccount200ApplicationJSONSourceModifiedDateTypeEnum = "Expense"
	CreateAccount200ApplicationJSONSourceModifiedDateTypeEnumIncome    CreateAccount200ApplicationJSONSourceModifiedDateTypeEnum = "Income"
	CreateAccount200ApplicationJSONSourceModifiedDateTypeEnumLiability CreateAccount200ApplicationJSONSourceModifiedDateTypeEnum = "Liability"
	CreateAccount200ApplicationJSONSourceModifiedDateTypeEnumEquity    CreateAccount200ApplicationJSONSourceModifiedDateTypeEnum = "Equity"
)

// CreateAccount200ApplicationJSONSourceModifiedDateValidDataTypeLinks
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
type CreateAccount200ApplicationJSONSourceModifiedDateValidDataTypeLinks struct {
	Links    []string `json:"links,omitempty"`
	Property *string  `json:"property,omitempty"`
}

// CreateAccount200ApplicationJSONSourceModifiedDate
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
type CreateAccount200ApplicationJSONSourceModifiedDate struct {
	Currency               *string                                                               `json:"currency,omitempty"`
	CurrentBalance         *float64                                                              `json:"currentBalance,omitempty"`
	Description            *string                                                               `json:"description,omitempty"`
	FullyQualifiedCategory *string                                                               `json:"fullyQualifiedCategory,omitempty"`
	FullyQualifiedName     *string                                                               `json:"fullyQualifiedName,omitempty"`
	ID                     *string                                                               `json:"id,omitempty"`
	IsBankAccount          bool                                                                  `json:"isBankAccount"`
	Metadata               *CreateAccount200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate           *time.Time                                                            `json:"modifiedDate,omitempty"`
	Name                   *string                                                               `json:"name,omitempty"`
	NominalCode            *string                                                               `json:"nominalCode,omitempty"`
	SourceModifiedDate     *time.Time                                                            `json:"sourceModifiedDate,omitempty"`
	Status                 CreateAccount200ApplicationJSONSourceModifiedDateStatusEnum           `json:"status"`
	Type                   CreateAccount200ApplicationJSONSourceModifiedDateTypeEnum             `json:"type"`
	ValidDatatypeLinks     []CreateAccount200ApplicationJSONSourceModifiedDateValidDataTypeLinks `json:"validDatatypeLinks,omitempty"`
}

type CreateAccount200ApplicationJSONStatusEnum string

const (
	CreateAccount200ApplicationJSONStatusEnumPending  CreateAccount200ApplicationJSONStatusEnum = "Pending"
	CreateAccount200ApplicationJSONStatusEnumFailed   CreateAccount200ApplicationJSONStatusEnum = "Failed"
	CreateAccount200ApplicationJSONStatusEnumSuccess  CreateAccount200ApplicationJSONStatusEnum = "Success"
	CreateAccount200ApplicationJSONStatusEnumTimedOut CreateAccount200ApplicationJSONStatusEnum = "TimedOut"
)

type CreateAccount200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateAccount200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateAccount200ApplicationJSONValidation struct {
	Errors   []CreateAccount200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateAccount200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type CreateAccount200ApplicationJSON struct {
	Changes           []CreateAccount200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                             `json:"companyId"`
	CompletedOnUtc    *time.Time                                         `json:"completedOnUtc,omitempty"`
	Data              *CreateAccount200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                             `json:"dataConnectionKey"`
	DataType          *string                                            `json:"dataType,omitempty"`
	ErrorMessage      *string                                            `json:"errorMessage,omitempty"`
	PushOperationKey  string                                             `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                          `json:"requestedOnUtc"`
	Status            CreateAccount200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                `json:"statusCode"`
	TimeoutInMinutes  *int                                               `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                               `json:"timeoutInSeconds,omitempty"`
	Validation        *CreateAccount200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type CreateAccountResponse struct {
	ContentType                           string
	StatusCode                            int
	RawResponse                           *http.Response
	CreateAccount200ApplicationJSONObject *CreateAccount200ApplicationJSON
}
