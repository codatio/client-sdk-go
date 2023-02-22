package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostCreditNotePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	CreditNoteID string `pathParam:"style=simple,explode=false,name=creditNoteId"`
}

type PostCreditNoteQueryParams struct {
	ForceUpdate      *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostCreditNoteSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostCreditNoteRequest struct {
	PathParams  PostCreditNotePathParams
	QueryParams PostCreditNoteQueryParams
	Request     *shared.CreditNote `request:"mediaType=application/json"`
	Security    PostCreditNoteSecurity
}

type PostCreditNote200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostCreditNote200ApplicationJSONChangesTypeEnum string

const (
	PostCreditNote200ApplicationJSONChangesTypeEnumUnknown            PostCreditNote200ApplicationJSONChangesTypeEnum = "Unknown"
	PostCreditNote200ApplicationJSONChangesTypeEnumCreated            PostCreditNote200ApplicationJSONChangesTypeEnum = "Created"
	PostCreditNote200ApplicationJSONChangesTypeEnumModified           PostCreditNote200ApplicationJSONChangesTypeEnum = "Modified"
	PostCreditNote200ApplicationJSONChangesTypeEnumDeleted            PostCreditNote200ApplicationJSONChangesTypeEnum = "Deleted"
	PostCreditNote200ApplicationJSONChangesTypeEnumAttachmentUploaded PostCreditNote200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostCreditNote200ApplicationJSONChanges struct {
	AttachmentID *string                                                        `json:"attachmentId,omitempty"`
	RecordRef    *PostCreditNote200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostCreditNote200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostCreditNote200ApplicationJSONStatusEnum string

const (
	PostCreditNote200ApplicationJSONStatusEnumPending  PostCreditNote200ApplicationJSONStatusEnum = "Pending"
	PostCreditNote200ApplicationJSONStatusEnumFailed   PostCreditNote200ApplicationJSONStatusEnum = "Failed"
	PostCreditNote200ApplicationJSONStatusEnumSuccess  PostCreditNote200ApplicationJSONStatusEnum = "Success"
	PostCreditNote200ApplicationJSONStatusEnumTimedOut PostCreditNote200ApplicationJSONStatusEnum = "TimedOut"
)

type PostCreditNote200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostCreditNote200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostCreditNote200ApplicationJSONValidation struct {
	Errors   []PostCreditNote200ApplicationJSONValidationValidationItem                                                                                                                                          `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostCreditNote200ApplicationJSON struct {
	Changes           []PostCreditNote200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                      `json:"companyId"`
	CompletedOnUtc    *time.Time                                  `json:"completedOnUtc,omitempty"`
	Data              *shared.CreditNote                          `json:"data,omitempty"`
	DataConnectionKey string                                      `json:"dataConnectionKey"`
	DataType          *string                                     `json:"dataType,omitempty"`
	ErrorMessage      *string                                     `json:"errorMessage,omitempty"`
	PushOperationKey  string                                      `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                   `json:"requestedOnUtc"`
	Status            PostCreditNote200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                         `json:"statusCode"`
	TimeoutInMinutes  *int                                        `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                        `json:"timeoutInSeconds,omitempty"`
	Validation        *PostCreditNote200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostCreditNoteResponse struct {
	ContentType                            string
	StatusCode                             int
	PostCreditNote200ApplicationJSONObject *PostCreditNote200ApplicationJSON
}
