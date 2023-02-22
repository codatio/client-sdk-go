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

type PostBankAccountSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostBankAccountRequest struct {
	PathParams  PostBankAccountPathParams
	QueryParams PostBankAccountQueryParams
	Request     *shared.BankAccount `request:"mediaType=application/json"`
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
	Errors   []PostBankAccount200ApplicationJSONValidationValidationItem                                                                                                                                         `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostBankAccount200ApplicationJSON struct {
	Changes           []PostBankAccount200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                       `json:"companyId"`
	CompletedOnUtc    *time.Time                                   `json:"completedOnUtc,omitempty"`
	Data              *shared.BankAccount                          `json:"data,omitempty"`
	DataConnectionKey string                                       `json:"dataConnectionKey"`
	DataType          *string                                      `json:"dataType,omitempty"`
	ErrorMessage      *string                                      `json:"errorMessage,omitempty"`
	PushOperationKey  string                                       `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                    `json:"requestedOnUtc"`
	Status            PostBankAccount200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                          `json:"statusCode"`
	TimeoutInMinutes  *int                                         `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                         `json:"timeoutInSeconds,omitempty"`
	Validation        *PostBankAccount200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostBankAccountResponse struct {
	ContentType                             string
	StatusCode                              int
	PostBankAccount200ApplicationJSONObject *PostBankAccount200ApplicationJSON
}
