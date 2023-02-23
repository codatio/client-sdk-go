package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostBankAccountPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostBankAccountQueryParams struct {
	AllowSyncOnPushComplete *bool `queryParam:"style=form,explode=true,name=allowSyncOnPushComplete"`
	TimeoutInMinutes        *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostBankAccountSourceModifiedDateAccountTypeEnum string

const (
	PostBankAccountSourceModifiedDateAccountTypeEnumUnknown PostBankAccountSourceModifiedDateAccountTypeEnum = "Unknown"
	PostBankAccountSourceModifiedDateAccountTypeEnumCredit  PostBankAccountSourceModifiedDateAccountTypeEnum = "Credit"
	PostBankAccountSourceModifiedDateAccountTypeEnumDebit   PostBankAccountSourceModifiedDateAccountTypeEnum = "Debit"
)

type PostBankAccountSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// PostBankAccountSourceModifiedDate
// > **Accessing Bank Accounts through Banking API**
// >
// > This datatype was originally used for accessing bank account data both in accounting integrations and open banking aggregators.
// >
// > To view bank account data through the Banking API, please refer to the new datatype [here](https://docs.codat.io/banking-api#/schemas/Account)
//
// > View the coverage for bank accounts in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A list of bank accounts associated with a company and a specific [data connection](https://api.codat.io/swagger/index.html#/Connection/get_companies__companyId__connections__connectionId_).
//
// Bank accounts data includes:
// * The name and ID of the account in the accounting platform.
// * The currency and balance of the account.
// * The sort code and account number.
type PostBankAccountSourceModifiedDate struct {
	AccountName        *string                                           `json:"accountName,omitempty"`
	AccountNumber      *string                                           `json:"accountNumber,omitempty"`
	AccountType        *PostBankAccountSourceModifiedDateAccountTypeEnum `json:"accountType,omitempty"`
	AvailableBalance   *float64                                          `json:"availableBalance,omitempty"`
	Balance            *float64                                          `json:"balance,omitempty"`
	Currency           *string                                           `json:"currency,omitempty"`
	IBan               *string                                           `json:"iBan,omitempty"`
	ID                 *string                                           `json:"id,omitempty"`
	Institution        *string                                           `json:"institution,omitempty"`
	Metadata           *PostBankAccountSourceModifiedDateMetadata        `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                        `json:"modifiedDate,omitempty"`
	NominalCode        *string                                           `json:"nominalCode,omitempty"`
	OverdraftLimit     *float64                                          `json:"overdraftLimit,omitempty"`
	SortCode           *string                                           `json:"sortCode,omitempty"`
	SourceModifiedDate *time.Time                                        `json:"sourceModifiedDate,omitempty"`
}

type PostBankAccountSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostBankAccountRequest struct {
	PathParams  PostBankAccountPathParams
	QueryParams PostBankAccountQueryParams
	Request     *PostBankAccountSourceModifiedDate `request:"mediaType=application/json"`
	Security    PostBankAccountSecurity
}

type PostBankAccount200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostBankAccount200ApplicationJSONChangesTypeEnum string

const (
	PostBankAccount200ApplicationJSONChangesTypeEnumUnknown            PostBankAccount200ApplicationJSONChangesTypeEnum = "Unknown"
	PostBankAccount200ApplicationJSONChangesTypeEnumCreated            PostBankAccount200ApplicationJSONChangesTypeEnum = "Created"
	PostBankAccount200ApplicationJSONChangesTypeEnumModified           PostBankAccount200ApplicationJSONChangesTypeEnum = "Modified"
	PostBankAccount200ApplicationJSONChangesTypeEnumDeleted            PostBankAccount200ApplicationJSONChangesTypeEnum = "Deleted"
	PostBankAccount200ApplicationJSONChangesTypeEnumAttachmentUploaded PostBankAccount200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostBankAccount200ApplicationJSONChanges struct {
	AttachmentID *string                                                         `json:"attachmentId,omitempty"`
	RecordRef    *PostBankAccount200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostBankAccount200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum string

const (
	PostBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnumUnknown PostBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum = "Unknown"
	PostBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnumCredit  PostBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum = "Credit"
	PostBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnumDebit   PostBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum = "Debit"
)

type PostBankAccount200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// PostBankAccount200ApplicationJSONSourceModifiedDate
// > **Accessing Bank Accounts through Banking API**
// >
// > This datatype was originally used for accessing bank account data both in accounting integrations and open banking aggregators.
// >
// > To view bank account data through the Banking API, please refer to the new datatype [here](https://docs.codat.io/banking-api#/schemas/Account)
//
// > View the coverage for bank accounts in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A list of bank accounts associated with a company and a specific [data connection](https://api.codat.io/swagger/index.html#/Connection/get_companies__companyId__connections__connectionId_).
//
// Bank accounts data includes:
// * The name and ID of the account in the accounting platform.
// * The currency and balance of the account.
// * The sort code and account number.
type PostBankAccount200ApplicationJSONSourceModifiedDate struct {
	AccountName        *string                                                             `json:"accountName,omitempty"`
	AccountNumber      *string                                                             `json:"accountNumber,omitempty"`
	AccountType        *PostBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum `json:"accountType,omitempty"`
	AvailableBalance   *float64                                                            `json:"availableBalance,omitempty"`
	Balance            *float64                                                            `json:"balance,omitempty"`
	Currency           *string                                                             `json:"currency,omitempty"`
	IBan               *string                                                             `json:"iBan,omitempty"`
	ID                 *string                                                             `json:"id,omitempty"`
	Institution        *string                                                             `json:"institution,omitempty"`
	Metadata           *PostBankAccount200ApplicationJSONSourceModifiedDateMetadata        `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                          `json:"modifiedDate,omitempty"`
	NominalCode        *string                                                             `json:"nominalCode,omitempty"`
	OverdraftLimit     *float64                                                            `json:"overdraftLimit,omitempty"`
	SortCode           *string                                                             `json:"sortCode,omitempty"`
	SourceModifiedDate *time.Time                                                          `json:"sourceModifiedDate,omitempty"`
}

type PostBankAccount200ApplicationJSONStatusEnum string

const (
	PostBankAccount200ApplicationJSONStatusEnumPending  PostBankAccount200ApplicationJSONStatusEnum = "Pending"
	PostBankAccount200ApplicationJSONStatusEnumFailed   PostBankAccount200ApplicationJSONStatusEnum = "Failed"
	PostBankAccount200ApplicationJSONStatusEnumSuccess  PostBankAccount200ApplicationJSONStatusEnum = "Success"
	PostBankAccount200ApplicationJSONStatusEnumTimedOut PostBankAccount200ApplicationJSONStatusEnum = "TimedOut"
)

type PostBankAccount200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostBankAccount200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostBankAccount200ApplicationJSONValidation struct {
	Errors   []PostBankAccount200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PostBankAccount200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PostBankAccount200ApplicationJSON struct {
	Changes           []PostBankAccount200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                               `json:"companyId"`
	CompletedOnUtc    *time.Time                                           `json:"completedOnUtc,omitempty"`
	Data              *PostBankAccount200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                               `json:"dataConnectionKey"`
	DataType          *string                                              `json:"dataType,omitempty"`
	ErrorMessage      *string                                              `json:"errorMessage,omitempty"`
	PushOperationKey  string                                               `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                            `json:"requestedOnUtc"`
	Status            PostBankAccount200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                  `json:"statusCode"`
	TimeoutInMinutes  *int                                                 `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                 `json:"timeoutInSeconds,omitempty"`
	Validation        *PostBankAccount200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PostBankAccountResponse struct {
	ContentType                             string
	StatusCode                              int
	PostBankAccount200ApplicationJSONObject *PostBankAccount200ApplicationJSON
}
