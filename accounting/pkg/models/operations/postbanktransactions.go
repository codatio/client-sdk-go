package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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

type PostBankTransactionsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostBankTransactionsRequest struct {
	PathParams  PostBankTransactionsPathParams
	QueryParams PostBankTransactionsQueryParams
	Request     *shared.BankTransactions `request:"mediaType=application/json"`
	Security    PostBankTransactionsSecurity
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
	Errors   []PostBankTransactions200ApplicationJSONValidationValidationItem                                                                                                                                    `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostBankTransactions200ApplicationJSON struct {
	Changes           []PostBankTransactions200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                            `json:"companyId"`
	CompletedOnUtc    *time.Time                                        `json:"completedOnUtc,omitempty"`
	Data              *shared.BankTransactions                          `json:"data,omitempty"`
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
	PostBankTransactions200ApplicationJSONObject *PostBankTransactions200ApplicationJSON
}
