package operations

import (
	"net/http"
	"time"
)

type PutBankAccountSourceModifiedDateAccountTypeEnum string

const (
	PutBankAccountSourceModifiedDateAccountTypeEnumUnknown PutBankAccountSourceModifiedDateAccountTypeEnum = "Unknown"
	PutBankAccountSourceModifiedDateAccountTypeEnumCredit  PutBankAccountSourceModifiedDateAccountTypeEnum = "Credit"
	PutBankAccountSourceModifiedDateAccountTypeEnumDebit   PutBankAccountSourceModifiedDateAccountTypeEnum = "Debit"
)

type PutBankAccountSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// PutBankAccountSourceModifiedDate
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
type PutBankAccountSourceModifiedDate struct {
	AccountName        *string                                          `json:"accountName,omitempty"`
	AccountNumber      *string                                          `json:"accountNumber,omitempty"`
	AccountType        *PutBankAccountSourceModifiedDateAccountTypeEnum `json:"accountType,omitempty"`
	AvailableBalance   *float64                                         `json:"availableBalance,omitempty"`
	Balance            *float64                                         `json:"balance,omitempty"`
	Currency           *string                                          `json:"currency,omitempty"`
	IBan               *string                                          `json:"iBan,omitempty"`
	ID                 *string                                          `json:"id,omitempty"`
	Institution        *string                                          `json:"institution,omitempty"`
	Metadata           *PutBankAccountSourceModifiedDateMetadata        `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                       `json:"modifiedDate,omitempty"`
	NominalCode        *string                                          `json:"nominalCode,omitempty"`
	OverdraftLimit     *float64                                         `json:"overdraftLimit,omitempty"`
	SortCode           *string                                          `json:"sortCode,omitempty"`
	SourceModifiedDate *time.Time                                       `json:"sourceModifiedDate,omitempty"`
}

type PutBankAccountRequest struct {
	RequestBody      *PutBankAccountSourceModifiedDate `request:"mediaType=application/json"`
	BankAccountID    string                            `pathParam:"style=simple,explode=false,name=bankAccountId"`
	CompanyID        string                            `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string                            `pathParam:"style=simple,explode=false,name=connectionId"`
	ForceUpdate      *bool                             `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int                              `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PutBankAccount200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PutBankAccount200ApplicationJSONChangesTypeEnum string

const (
	PutBankAccount200ApplicationJSONChangesTypeEnumUnknown            PutBankAccount200ApplicationJSONChangesTypeEnum = "Unknown"
	PutBankAccount200ApplicationJSONChangesTypeEnumCreated            PutBankAccount200ApplicationJSONChangesTypeEnum = "Created"
	PutBankAccount200ApplicationJSONChangesTypeEnumModified           PutBankAccount200ApplicationJSONChangesTypeEnum = "Modified"
	PutBankAccount200ApplicationJSONChangesTypeEnumDeleted            PutBankAccount200ApplicationJSONChangesTypeEnum = "Deleted"
	PutBankAccount200ApplicationJSONChangesTypeEnumAttachmentUploaded PutBankAccount200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PutBankAccount200ApplicationJSONChanges struct {
	AttachmentID *string                                                        `json:"attachmentId,omitempty"`
	RecordRef    *PutBankAccount200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PutBankAccount200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PutBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum string

const (
	PutBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnumUnknown PutBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum = "Unknown"
	PutBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnumCredit  PutBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum = "Credit"
	PutBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnumDebit   PutBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum = "Debit"
)

type PutBankAccount200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// PutBankAccount200ApplicationJSONSourceModifiedDate
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
type PutBankAccount200ApplicationJSONSourceModifiedDate struct {
	AccountName        *string                                                            `json:"accountName,omitempty"`
	AccountNumber      *string                                                            `json:"accountNumber,omitempty"`
	AccountType        *PutBankAccount200ApplicationJSONSourceModifiedDateAccountTypeEnum `json:"accountType,omitempty"`
	AvailableBalance   *float64                                                           `json:"availableBalance,omitempty"`
	Balance            *float64                                                           `json:"balance,omitempty"`
	Currency           *string                                                            `json:"currency,omitempty"`
	IBan               *string                                                            `json:"iBan,omitempty"`
	ID                 *string                                                            `json:"id,omitempty"`
	Institution        *string                                                            `json:"institution,omitempty"`
	Metadata           *PutBankAccount200ApplicationJSONSourceModifiedDateMetadata        `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                         `json:"modifiedDate,omitempty"`
	NominalCode        *string                                                            `json:"nominalCode,omitempty"`
	OverdraftLimit     *float64                                                           `json:"overdraftLimit,omitempty"`
	SortCode           *string                                                            `json:"sortCode,omitempty"`
	SourceModifiedDate *time.Time                                                         `json:"sourceModifiedDate,omitempty"`
}

type PutBankAccount200ApplicationJSONStatusEnum string

const (
	PutBankAccount200ApplicationJSONStatusEnumPending  PutBankAccount200ApplicationJSONStatusEnum = "Pending"
	PutBankAccount200ApplicationJSONStatusEnumFailed   PutBankAccount200ApplicationJSONStatusEnum = "Failed"
	PutBankAccount200ApplicationJSONStatusEnumSuccess  PutBankAccount200ApplicationJSONStatusEnum = "Success"
	PutBankAccount200ApplicationJSONStatusEnumTimedOut PutBankAccount200ApplicationJSONStatusEnum = "TimedOut"
)

type PutBankAccount200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PutBankAccount200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PutBankAccount200ApplicationJSONValidation struct {
	Errors   []PutBankAccount200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PutBankAccount200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PutBankAccount200ApplicationJSON struct {
	Changes           []PutBankAccount200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                              `json:"companyId"`
	CompletedOnUtc    *time.Time                                          `json:"completedOnUtc,omitempty"`
	Data              *PutBankAccount200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                              `json:"dataConnectionKey"`
	DataType          *string                                             `json:"dataType,omitempty"`
	ErrorMessage      *string                                             `json:"errorMessage,omitempty"`
	PushOperationKey  string                                              `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                           `json:"requestedOnUtc"`
	Status            PutBankAccount200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                 `json:"statusCode"`
	TimeoutInMinutes  *int                                                `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                `json:"timeoutInSeconds,omitempty"`
	Validation        *PutBankAccount200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PutBankAccountResponse struct {
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
	PutBankAccount200ApplicationJSONObject *PutBankAccount200ApplicationJSON
}
