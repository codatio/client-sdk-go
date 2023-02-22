package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PushCreditNotePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PushCreditNoteQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PushCreditNoteSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PushCreditNoteRequest struct {
	PathParams  PushCreditNotePathParams
	QueryParams PushCreditNoteQueryParams
	Request     *shared.CreditNote `request:"mediaType=application/json"`
	Security    PushCreditNoteSecurity
}

type PushCreditNote200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PushCreditNote200ApplicationJSONChangesTypeEnum string

const (
	PushCreditNote200ApplicationJSONChangesTypeEnumUnknown            PushCreditNote200ApplicationJSONChangesTypeEnum = "Unknown"
	PushCreditNote200ApplicationJSONChangesTypeEnumCreated            PushCreditNote200ApplicationJSONChangesTypeEnum = "Created"
	PushCreditNote200ApplicationJSONChangesTypeEnumModified           PushCreditNote200ApplicationJSONChangesTypeEnum = "Modified"
	PushCreditNote200ApplicationJSONChangesTypeEnumDeleted            PushCreditNote200ApplicationJSONChangesTypeEnum = "Deleted"
	PushCreditNote200ApplicationJSONChangesTypeEnumAttachmentUploaded PushCreditNote200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PushCreditNote200ApplicationJSONChanges struct {
	AttachmentID *string                                                        `json:"attachmentId,omitempty"`
	RecordRef    *PushCreditNote200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PushCreditNote200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PushCreditNote200ApplicationJSONStatusEnum string

const (
	PushCreditNote200ApplicationJSONStatusEnumPending  PushCreditNote200ApplicationJSONStatusEnum = "Pending"
	PushCreditNote200ApplicationJSONStatusEnumFailed   PushCreditNote200ApplicationJSONStatusEnum = "Failed"
	PushCreditNote200ApplicationJSONStatusEnumSuccess  PushCreditNote200ApplicationJSONStatusEnum = "Success"
	PushCreditNote200ApplicationJSONStatusEnumTimedOut PushCreditNote200ApplicationJSONStatusEnum = "TimedOut"
)

type PushCreditNote200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PushCreditNote200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PushCreditNote200ApplicationJSONValidation struct {
	Errors   []PushCreditNote200ApplicationJSONValidationValidationItem                                                                                                                                          `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PushCreditNote200ApplicationJSON struct {
	Changes           []PushCreditNote200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                      `json:"companyId"`
	CompletedOnUtc    *time.Time                                  `json:"completedOnUtc,omitempty"`
	Data              *shared.CreditNote                          `json:"data,omitempty"`
	DataConnectionKey string                                      `json:"dataConnectionKey"`
	DataType          *string                                     `json:"dataType,omitempty"`
	ErrorMessage      *string                                     `json:"errorMessage,omitempty"`
	PushOperationKey  string                                      `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                   `json:"requestedOnUtc"`
	Status            PushCreditNote200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                         `json:"statusCode"`
	TimeoutInMinutes  *int                                        `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                        `json:"timeoutInSeconds,omitempty"`
	Validation        *PushCreditNote200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PushCreditNoteResponse struct {
	ContentType                            string
	StatusCode                             int
	PushCreditNote200ApplicationJSONObject *PushCreditNote200ApplicationJSON
}
