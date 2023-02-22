package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PutBankAccountPathParams struct {
	BankAccountID string `pathParam:"style=simple,explode=false,name=bankAccountId"`
	CompanyID     string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID  string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PutBankAccountQueryParams struct {
	ForceUpdate      *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PutBankAccountSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PutBankAccountRequest struct {
	PathParams  PutBankAccountPathParams
	QueryParams PutBankAccountQueryParams
	Request     *shared.BankAccount `request:"mediaType=application/json"`
	Security    PutBankAccountSecurity
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
	Errors   []PutBankAccount200ApplicationJSONValidationValidationItem                                                                                                                                          `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PutBankAccount200ApplicationJSON struct {
	Changes           []PutBankAccount200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                      `json:"companyId"`
	CompletedOnUtc    *time.Time                                  `json:"completedOnUtc,omitempty"`
	Data              *shared.BankAccount                         `json:"data,omitempty"`
	DataConnectionKey string                                      `json:"dataConnectionKey"`
	DataType          *string                                     `json:"dataType,omitempty"`
	ErrorMessage      *string                                     `json:"errorMessage,omitempty"`
	PushOperationKey  string                                      `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                   `json:"requestedOnUtc"`
	Status            PutBankAccount200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                         `json:"statusCode"`
	TimeoutInMinutes  *int                                        `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                        `json:"timeoutInSeconds,omitempty"`
	Validation        *PutBankAccount200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PutBankAccountResponse struct {
	ContentType                            string
	StatusCode                             int
	PutBankAccount200ApplicationJSONObject *PutBankAccount200ApplicationJSON
}
