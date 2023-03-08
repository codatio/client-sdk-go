package operations

import (
	"net/http"
	"time"
)

type PostBankTransactionsPathParams struct {
	AccountID    string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostBankTransactionsQueryParams struct {
	AllowSyncOnPushComplete *bool `queryParam:"style=form,explode=true,name=allowSyncOnPushComplete"`
	TimeoutInMinutes        *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum string

const (
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumUnknown     PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "Unknown"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumCredit      PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "Credit"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumDebit       PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "Debit"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumInt         PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "Int"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumDiv         PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "Div"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumFee         PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "Fee"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumSerChg      PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "SerChg"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumDep         PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "Dep"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumAtm         PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "Atm"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumPos         PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "Pos"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumXfer        PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "Xfer"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumCheck       PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "Check"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumPayment     PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "Payment"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumCash        PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "Cash"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumDirectDep   PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "DirectDep"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumDirectDebit PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "DirectDebit"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumRepeatPmt   PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "RepeatPmt"
	PostBankTransactionsRequestBodyTransactionsTransactionTypeEnumOther       PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum = "Other"
)

type PostBankTransactionsRequestBodyTransactions struct {
	Amount             float64                                                        `json:"amount"`
	Balance            float64                                                        `json:"balance"`
	Counterparty       *string                                                        `json:"counterparty,omitempty"`
	Date               time.Time                                                      `json:"date"`
	Description        *string                                                        `json:"description,omitempty"`
	ID                 *string                                                        `json:"id,omitempty"`
	ModifiedDate       *time.Time                                                     `json:"modifiedDate,omitempty"`
	Reconciled         bool                                                           `json:"reconciled"`
	Reference          *string                                                        `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                                     `json:"sourceModifiedDate,omitempty"`
	TransactionType    PostBankTransactionsRequestBodyTransactionsTransactionTypeEnum `json:"transactionType"`
}

// PostBankTransactionsRequestBody
// > **Accessing Bank Accounts through Banking API**
// >
// > This datatype was originally used for accessing bank account data both in accounting integrations and open banking aggregators.
// >
// > To view bank account data through the Banking API, please refer to the new datatype [here](https://docs.codat.io/banking-api#/operations/list-all-banking-transactions)
//
// > View the coverage for bank transactions in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankTransactions" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Transactional banking data for a specific company and account.
//
// Bank transactions include the:
// * Amount of the transaction.
// * Current account balance.
// * Transaction type, for example, credit, debit, or transfer.
type PostBankTransactionsRequestBody struct {
	AccountID       *string                                       `json:"accountId,omitempty"`
	ContractVersion *string                                       `json:"contractVersion,omitempty"`
	Transactions    []PostBankTransactionsRequestBodyTransactions `json:"transactions,omitempty"`
}

type PostBankTransactionsRequest struct {
	PathParams  PostBankTransactionsPathParams
	QueryParams PostBankTransactionsQueryParams
	Request     *PostBankTransactionsRequestBody `request:"mediaType=application/json"`
}

type PostBankTransactions200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostBankTransactions200ApplicationJSONChangesTypeEnum string

const (
	PostBankTransactions200ApplicationJSONChangesTypeEnumUnknown            PostBankTransactions200ApplicationJSONChangesTypeEnum = "Unknown"
	PostBankTransactions200ApplicationJSONChangesTypeEnumCreated            PostBankTransactions200ApplicationJSONChangesTypeEnum = "Created"
	PostBankTransactions200ApplicationJSONChangesTypeEnumModified           PostBankTransactions200ApplicationJSONChangesTypeEnum = "Modified"
	PostBankTransactions200ApplicationJSONChangesTypeEnumDeleted            PostBankTransactions200ApplicationJSONChangesTypeEnum = "Deleted"
	PostBankTransactions200ApplicationJSONChangesTypeEnumAttachmentUploaded PostBankTransactions200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostBankTransactions200ApplicationJSONChanges struct {
	AttachmentID *string                                                              `json:"attachmentId,omitempty"`
	RecordRef    *PostBankTransactions200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostBankTransactions200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum string

const (
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumUnknown     PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "Unknown"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumCredit      PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "Credit"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumDebit       PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "Debit"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumInt         PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "Int"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumDiv         PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "Div"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumFee         PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "Fee"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumSerChg      PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "SerChg"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumDep         PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "Dep"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumAtm         PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "Atm"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumPos         PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "Pos"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumXfer        PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "Xfer"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumCheck       PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "Check"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumPayment     PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "Payment"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumCash        PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "Cash"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumDirectDep   PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "DirectDep"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumDirectDebit PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "DirectDebit"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumRepeatPmt   PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "RepeatPmt"
	PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnumOther       PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum = "Other"
)

type PostBankTransactions200ApplicationJSONDataTransactions struct {
	Amount             float64                                                                   `json:"amount"`
	Balance            float64                                                                   `json:"balance"`
	Counterparty       *string                                                                   `json:"counterparty,omitempty"`
	Date               time.Time                                                                 `json:"date"`
	Description        *string                                                                   `json:"description,omitempty"`
	ID                 *string                                                                   `json:"id,omitempty"`
	ModifiedDate       *time.Time                                                                `json:"modifiedDate,omitempty"`
	Reconciled         bool                                                                      `json:"reconciled"`
	Reference          *string                                                                   `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                                                `json:"sourceModifiedDate,omitempty"`
	TransactionType    PostBankTransactions200ApplicationJSONDataTransactionsTransactionTypeEnum `json:"transactionType"`
}

// PostBankTransactions200ApplicationJSONData
// > **Accessing Bank Accounts through Banking API**
// >
// > This datatype was originally used for accessing bank account data both in accounting integrations and open banking aggregators.
// >
// > To view bank account data through the Banking API, please refer to the new datatype [here](https://docs.codat.io/banking-api#/operations/list-all-banking-transactions)
//
// > View the coverage for bank transactions in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankTransactions" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Transactional banking data for a specific company and account.
//
// Bank transactions include the:
// * Amount of the transaction.
// * Current account balance.
// * Transaction type, for example, credit, debit, or transfer.
type PostBankTransactions200ApplicationJSONData struct {
	AccountID       *string                                                  `json:"accountId,omitempty"`
	ContractVersion *string                                                  `json:"contractVersion,omitempty"`
	Transactions    []PostBankTransactions200ApplicationJSONDataTransactions `json:"transactions,omitempty"`
}

type PostBankTransactions200ApplicationJSONStatusEnum string

const (
	PostBankTransactions200ApplicationJSONStatusEnumPending  PostBankTransactions200ApplicationJSONStatusEnum = "Pending"
	PostBankTransactions200ApplicationJSONStatusEnumFailed   PostBankTransactions200ApplicationJSONStatusEnum = "Failed"
	PostBankTransactions200ApplicationJSONStatusEnumSuccess  PostBankTransactions200ApplicationJSONStatusEnum = "Success"
	PostBankTransactions200ApplicationJSONStatusEnumTimedOut PostBankTransactions200ApplicationJSONStatusEnum = "TimedOut"
)

type PostBankTransactions200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostBankTransactions200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostBankTransactions200ApplicationJSONValidation struct {
	Errors   []PostBankTransactions200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PostBankTransactions200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PostBankTransactions200ApplicationJSON struct {
	Changes           []PostBankTransactions200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                            `json:"companyId"`
	CompletedOnUtc    *time.Time                                        `json:"completedOnUtc,omitempty"`
	Data              *PostBankTransactions200ApplicationJSONData       `json:"data,omitempty"`
	DataConnectionKey string                                            `json:"dataConnectionKey"`
	DataType          *string                                           `json:"dataType,omitempty"`
	ErrorMessage      *string                                           `json:"errorMessage,omitempty"`
	PushOperationKey  string                                            `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                         `json:"requestedOnUtc"`
	Status            PostBankTransactions200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                               `json:"statusCode"`
	TimeoutInMinutes  *int                                              `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                              `json:"timeoutInSeconds,omitempty"`
	Validation        *PostBankTransactions200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostBankTransactionsResponse struct {
	ContentType                                  string
	StatusCode                                   int
	RawResponse                                  *http.Response
	PostBankTransactions200ApplicationJSONObject *PostBankTransactions200ApplicationJSON
}
