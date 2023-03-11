package operations

import (
	"net/http"
	"time"
)

type CreateBankAccountPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type CreateBankAccountQueryParams struct {
	AllowSyncOnPushComplete *bool `queryParam:"style=form,explode=true,name=allowSyncOnPushComplete"`
	TimeoutInMinutes        *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreateBankAccountSourceModifiedDateAccountTypeEnum string

const (
	CreateBankAccountSourceModifiedDateAccountTypeEnumUnknown CreateBankAccountSourceModifiedDateAccountTypeEnum = "Unknown"
	CreateBankAccountSourceModifiedDateAccountTypeEnumCredit  CreateBankAccountSourceModifiedDateAccountTypeEnum = "Credit"
	CreateBankAccountSourceModifiedDateAccountTypeEnumDebit   CreateBankAccountSourceModifiedDateAccountTypeEnum = "Debit"
)

type CreateBankAccountSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// CreateBankAccountSourceModifiedDate
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
type CreateBankAccountSourceModifiedDate struct {
	AccountName        *string                                             `json:"accountName,omitempty"`
	AccountNumber      *string                                             `json:"accountNumber,omitempty"`
	AccountType        *CreateBankAccountSourceModifiedDateAccountTypeEnum `json:"accountType,omitempty"`
	AvailableBalance   *float64                                            `json:"availableBalance,omitempty"`
	Balance            *float64                                            `json:"balance,omitempty"`
	Currency           *string                                             `json:"currency,omitempty"`
	IBan               *string                                             `json:"iBan,omitempty"`
	ID                 *string                                             `json:"id,omitempty"`
	Institution        *string                                             `json:"institution,omitempty"`
	Metadata           *CreateBankAccountSourceModifiedDateMetadata        `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                          `json:"modifiedDate,omitempty"`
	NominalCode        *string                                             `json:"nominalCode,omitempty"`
	OverdraftLimit     *float64                                            `json:"overdraftLimit,omitempty"`
	SortCode           *string                                             `json:"sortCode,omitempty"`
	SourceModifiedDate *time.Time                                          `json:"sourceModifiedDate,omitempty"`
}

type CreateBankAccountRequest struct {
	PathParams  CreateBankAccountPathParams
	QueryParams CreateBankAccountQueryParams
	Request     *CreateBankAccountSourceModifiedDate `request:"mediaType=application/json"`
}

type CreateBankAccount200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateBankAccount200ApplicationJSONChangesTypeEnum string

const (
	CreateBankAccount200ApplicationJSONChangesTypeEnumUnknown            CreateBankAccount200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateBankAccount200ApplicationJSONChangesTypeEnumCreated            CreateBankAccount200ApplicationJSONChangesTypeEnum = "Created"
	CreateBankAccount200ApplicationJSONChangesTypeEnumModified           CreateBankAccount200ApplicationJSONChangesTypeEnum = "Modified"
	CreateBankAccount200ApplicationJSONChangesTypeEnumDeleted            CreateBankAccount200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateBankAccount200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateBankAccount200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type CreateBankAccount200ApplicationJSONChanges struct {
	AttachmentID *string                                                           `json:"attachmentId,omitempty"`
	RecordRef    *CreateBankAccount200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateBankAccount200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum string

const (
	CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnumUnknown CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum = "Unknown"
	CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnumCredit  CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum = "Credit"
	CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnumDebit   CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum = "Debit"
)

type CreateBankAccount200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// CreateBankAccount200ApplicationJSONSourceModifiedDate
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
type CreateBankAccount200ApplicationJSONSourceModifiedDate struct {
	AccountName        *string                                                               `json:"accountName,omitempty"`
	AccountNumber      *string                                                               `json:"accountNumber,omitempty"`
	AccountType        *CreateBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum `json:"accountType,omitempty"`
	AvailableBalance   *float64                                                              `json:"availableBalance,omitempty"`
	Balance            *float64                                                              `json:"balance,omitempty"`
	Currency           *string                                                               `json:"currency,omitempty"`
	IBan               *string                                                               `json:"iBan,omitempty"`
	ID                 *string                                                               `json:"id,omitempty"`
	Institution        *string                                                               `json:"institution,omitempty"`
	Metadata           *CreateBankAccount200ApplicationJSONSourceModifiedDateMetadata        `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                            `json:"modifiedDate,omitempty"`
	NominalCode        *string                                                               `json:"nominalCode,omitempty"`
	OverdraftLimit     *float64                                                              `json:"overdraftLimit,omitempty"`
	SortCode           *string                                                               `json:"sortCode,omitempty"`
	SourceModifiedDate *time.Time                                                            `json:"sourceModifiedDate,omitempty"`
}

type CreateBankAccount200ApplicationJSONStatusEnum string

const (
	CreateBankAccount200ApplicationJSONStatusEnumPending  CreateBankAccount200ApplicationJSONStatusEnum = "Pending"
	CreateBankAccount200ApplicationJSONStatusEnumFailed   CreateBankAccount200ApplicationJSONStatusEnum = "Failed"
	CreateBankAccount200ApplicationJSONStatusEnumSuccess  CreateBankAccount200ApplicationJSONStatusEnum = "Success"
	CreateBankAccount200ApplicationJSONStatusEnumTimedOut CreateBankAccount200ApplicationJSONStatusEnum = "TimedOut"
)

type CreateBankAccount200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateBankAccount200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateBankAccount200ApplicationJSONValidation struct {
	Errors   []CreateBankAccount200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateBankAccount200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type CreateBankAccount200ApplicationJSON struct {
	Changes           []CreateBankAccount200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                                 `json:"companyId"`
	CompletedOnUtc    *time.Time                                             `json:"completedOnUtc,omitempty"`
	Data              *CreateBankAccount200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                                 `json:"dataConnectionKey"`
	DataType          *string                                                `json:"dataType,omitempty"`
	ErrorMessage      *string                                                `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                 `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                              `json:"requestedOnUtc"`
	Status            CreateBankAccount200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                    `json:"statusCode"`
	TimeoutInMinutes  *int                                                   `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                   `json:"timeoutInSeconds,omitempty"`
	Validation        *CreateBankAccount200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type CreateBankAccountResponse struct {
	ContentType                               string
	StatusCode                                int
	RawResponse                               *http.Response
	CreateBankAccount200ApplicationJSONObject *CreateBankAccount200ApplicationJSON
}
