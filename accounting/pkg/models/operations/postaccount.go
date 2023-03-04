package operations

import (
	"net/http"
	"time"
)

type PostAccountPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostAccountQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostAccountSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostAccountSourceModifiedDateStatusEnum string

const (
	PostAccountSourceModifiedDateStatusEnumUnknown  PostAccountSourceModifiedDateStatusEnum = "Unknown"
	PostAccountSourceModifiedDateStatusEnumActive   PostAccountSourceModifiedDateStatusEnum = "Active"
	PostAccountSourceModifiedDateStatusEnumArchived PostAccountSourceModifiedDateStatusEnum = "Archived"
	PostAccountSourceModifiedDateStatusEnumPending  PostAccountSourceModifiedDateStatusEnum = "Pending"
)

type PostAccountSourceModifiedDateTypeEnum string

const (
	PostAccountSourceModifiedDateTypeEnumUnknown   PostAccountSourceModifiedDateTypeEnum = "Unknown"
	PostAccountSourceModifiedDateTypeEnumAsset     PostAccountSourceModifiedDateTypeEnum = "Asset"
	PostAccountSourceModifiedDateTypeEnumExpense   PostAccountSourceModifiedDateTypeEnum = "Expense"
	PostAccountSourceModifiedDateTypeEnumIncome    PostAccountSourceModifiedDateTypeEnum = "Income"
	PostAccountSourceModifiedDateTypeEnumLiability PostAccountSourceModifiedDateTypeEnum = "Liability"
	PostAccountSourceModifiedDateTypeEnumEquity    PostAccountSourceModifiedDateTypeEnum = "Equity"
)

// PostAccountSourceModifiedDateValidDataTypeLinks
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
type PostAccountSourceModifiedDateValidDataTypeLinks struct {
	Links    []string `json:"links,omitempty"`
	Property *string  `json:"property,omitempty"`
}

// PostAccountSourceModifiedDate
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
type PostAccountSourceModifiedDate struct {
	Currency               *string                                           `json:"currency,omitempty"`
	CurrentBalance         *float64                                          `json:"currentBalance,omitempty"`
	Description            *string                                           `json:"description,omitempty"`
	FullyQualifiedCategory *string                                           `json:"fullyQualifiedCategory,omitempty"`
	FullyQualifiedName     *string                                           `json:"fullyQualifiedName,omitempty"`
	ID                     *string                                           `json:"id,omitempty"`
	IsBankAccount          bool                                              `json:"isBankAccount"`
	Metadata               *PostAccountSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate           *time.Time                                        `json:"modifiedDate,omitempty"`
	Name                   *string                                           `json:"name,omitempty"`
	NominalCode            *string                                           `json:"nominalCode,omitempty"`
	SourceModifiedDate     *time.Time                                        `json:"sourceModifiedDate,omitempty"`
	Status                 PostAccountSourceModifiedDateStatusEnum           `json:"status"`
	Type                   PostAccountSourceModifiedDateTypeEnum             `json:"type"`
	ValidDatatypeLinks     []PostAccountSourceModifiedDateValidDataTypeLinks `json:"validDatatypeLinks,omitempty"`
}

type PostAccountRequest struct {
	PathParams  PostAccountPathParams
	QueryParams PostAccountQueryParams
	Request     *PostAccountSourceModifiedDate `request:"mediaType=application/json"`
}

type PostAccount200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostAccount200ApplicationJSONChangesTypeEnum string

const (
	PostAccount200ApplicationJSONChangesTypeEnumUnknown            PostAccount200ApplicationJSONChangesTypeEnum = "Unknown"
	PostAccount200ApplicationJSONChangesTypeEnumCreated            PostAccount200ApplicationJSONChangesTypeEnum = "Created"
	PostAccount200ApplicationJSONChangesTypeEnumModified           PostAccount200ApplicationJSONChangesTypeEnum = "Modified"
	PostAccount200ApplicationJSONChangesTypeEnumDeleted            PostAccount200ApplicationJSONChangesTypeEnum = "Deleted"
	PostAccount200ApplicationJSONChangesTypeEnumAttachmentUploaded PostAccount200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostAccount200ApplicationJSONChanges struct {
	AttachmentID *string                                                     `json:"attachmentId,omitempty"`
	RecordRef    *PostAccount200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostAccount200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostAccount200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostAccount200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	PostAccount200ApplicationJSONSourceModifiedDateStatusEnumUnknown  PostAccount200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	PostAccount200ApplicationJSONSourceModifiedDateStatusEnumActive   PostAccount200ApplicationJSONSourceModifiedDateStatusEnum = "Active"
	PostAccount200ApplicationJSONSourceModifiedDateStatusEnumArchived PostAccount200ApplicationJSONSourceModifiedDateStatusEnum = "Archived"
	PostAccount200ApplicationJSONSourceModifiedDateStatusEnumPending  PostAccount200ApplicationJSONSourceModifiedDateStatusEnum = "Pending"
)

type PostAccount200ApplicationJSONSourceModifiedDateTypeEnum string

const (
	PostAccount200ApplicationJSONSourceModifiedDateTypeEnumUnknown   PostAccount200ApplicationJSONSourceModifiedDateTypeEnum = "Unknown"
	PostAccount200ApplicationJSONSourceModifiedDateTypeEnumAsset     PostAccount200ApplicationJSONSourceModifiedDateTypeEnum = "Asset"
	PostAccount200ApplicationJSONSourceModifiedDateTypeEnumExpense   PostAccount200ApplicationJSONSourceModifiedDateTypeEnum = "Expense"
	PostAccount200ApplicationJSONSourceModifiedDateTypeEnumIncome    PostAccount200ApplicationJSONSourceModifiedDateTypeEnum = "Income"
	PostAccount200ApplicationJSONSourceModifiedDateTypeEnumLiability PostAccount200ApplicationJSONSourceModifiedDateTypeEnum = "Liability"
	PostAccount200ApplicationJSONSourceModifiedDateTypeEnumEquity    PostAccount200ApplicationJSONSourceModifiedDateTypeEnum = "Equity"
)

// PostAccount200ApplicationJSONSourceModifiedDateValidDataTypeLinks
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
type PostAccount200ApplicationJSONSourceModifiedDateValidDataTypeLinks struct {
	Links    []string `json:"links,omitempty"`
	Property *string  `json:"property,omitempty"`
}

// PostAccount200ApplicationJSONSourceModifiedDate
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
type PostAccount200ApplicationJSONSourceModifiedDate struct {
	Currency               *string                                                             `json:"currency,omitempty"`
	CurrentBalance         *float64                                                            `json:"currentBalance,omitempty"`
	Description            *string                                                             `json:"description,omitempty"`
	FullyQualifiedCategory *string                                                             `json:"fullyQualifiedCategory,omitempty"`
	FullyQualifiedName     *string                                                             `json:"fullyQualifiedName,omitempty"`
	ID                     *string                                                             `json:"id,omitempty"`
	IsBankAccount          bool                                                                `json:"isBankAccount"`
	Metadata               *PostAccount200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate           *time.Time                                                          `json:"modifiedDate,omitempty"`
	Name                   *string                                                             `json:"name,omitempty"`
	NominalCode            *string                                                             `json:"nominalCode,omitempty"`
	SourceModifiedDate     *time.Time                                                          `json:"sourceModifiedDate,omitempty"`
	Status                 PostAccount200ApplicationJSONSourceModifiedDateStatusEnum           `json:"status"`
	Type                   PostAccount200ApplicationJSONSourceModifiedDateTypeEnum             `json:"type"`
	ValidDatatypeLinks     []PostAccount200ApplicationJSONSourceModifiedDateValidDataTypeLinks `json:"validDatatypeLinks,omitempty"`
}

type PostAccount200ApplicationJSONStatusEnum string

const (
	PostAccount200ApplicationJSONStatusEnumPending  PostAccount200ApplicationJSONStatusEnum = "Pending"
	PostAccount200ApplicationJSONStatusEnumFailed   PostAccount200ApplicationJSONStatusEnum = "Failed"
	PostAccount200ApplicationJSONStatusEnumSuccess  PostAccount200ApplicationJSONStatusEnum = "Success"
	PostAccount200ApplicationJSONStatusEnumTimedOut PostAccount200ApplicationJSONStatusEnum = "TimedOut"
)

type PostAccount200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostAccount200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostAccount200ApplicationJSONValidation struct {
	Errors   []PostAccount200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PostAccount200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PostAccount200ApplicationJSON struct {
	Changes           []PostAccount200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                           `json:"companyId"`
	CompletedOnUtc    *time.Time                                       `json:"completedOnUtc,omitempty"`
	Data              *PostAccount200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                           `json:"dataConnectionKey"`
	DataType          *string                                          `json:"dataType,omitempty"`
	ErrorMessage      *string                                          `json:"errorMessage,omitempty"`
	PushOperationKey  string                                           `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                        `json:"requestedOnUtc"`
	Status            PostAccount200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                              `json:"statusCode"`
	TimeoutInMinutes  *int                                             `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                             `json:"timeoutInSeconds,omitempty"`
	Validation        *PostAccount200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PostAccountResponse struct {
	ContentType                         string
	StatusCode                          int
	RawResponse                         *http.Response
	PostAccount200ApplicationJSONObject *PostAccount200ApplicationJSON
}
