package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostBillCreditNotePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostBillCreditNoteQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostBillCreditNoteSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostBillCreditNoteRequest struct {
	PathParams  PostBillCreditNotePathParams
	QueryParams PostBillCreditNoteQueryParams
	Request     *shared.BillCreditNote `request:"mediaType=application/json"`
	Security    PostBillCreditNoteSecurity
}

type PostBillCreditNote200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostBillCreditNote200ApplicationJSONChangesTypeEnum string

const (
	PostBillCreditNote200ApplicationJSONChangesTypeEnumUnknown            PostBillCreditNote200ApplicationJSONChangesTypeEnum = "Unknown"
	PostBillCreditNote200ApplicationJSONChangesTypeEnumCreated            PostBillCreditNote200ApplicationJSONChangesTypeEnum = "Created"
	PostBillCreditNote200ApplicationJSONChangesTypeEnumModified           PostBillCreditNote200ApplicationJSONChangesTypeEnum = "Modified"
	PostBillCreditNote200ApplicationJSONChangesTypeEnumDeleted            PostBillCreditNote200ApplicationJSONChangesTypeEnum = "Deleted"
	PostBillCreditNote200ApplicationJSONChangesTypeEnumAttachmentUploaded PostBillCreditNote200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostBillCreditNote200ApplicationJSONChanges struct {
	AttachmentID *string                                                            `json:"attachmentId,omitempty"`
	RecordRef    *PostBillCreditNote200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostBillCreditNote200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostBillCreditNote200ApplicationJSONStatusEnum string

const (
	PostBillCreditNote200ApplicationJSONStatusEnumPending  PostBillCreditNote200ApplicationJSONStatusEnum = "Pending"
	PostBillCreditNote200ApplicationJSONStatusEnumFailed   PostBillCreditNote200ApplicationJSONStatusEnum = "Failed"
	PostBillCreditNote200ApplicationJSONStatusEnumSuccess  PostBillCreditNote200ApplicationJSONStatusEnum = "Success"
	PostBillCreditNote200ApplicationJSONStatusEnumTimedOut PostBillCreditNote200ApplicationJSONStatusEnum = "TimedOut"
)

type PostBillCreditNote200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostBillCreditNote200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostBillCreditNote200ApplicationJSONValidation struct {
	Errors   []PostBillCreditNote200ApplicationJSONValidationValidationItem                                                                                                                                      `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostBillCreditNote200ApplicationJSON struct {
	Changes           []PostBillCreditNote200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                          `json:"companyId"`
	CompletedOnUtc    *time.Time                                      `json:"completedOnUtc,omitempty"`
	Data              *shared.BillCreditNote                          `json:"data,omitempty"`
	DataConnectionKey string                                          `json:"dataConnectionKey"`
	DataType          *string                                         `json:"dataType,omitempty"`
	ErrorMessage      *string                                         `json:"errorMessage,omitempty"`
	PushOperationKey  string                                          `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                       `json:"requestedOnUtc"`
	Status            PostBillCreditNote200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                             `json:"statusCode"`
	TimeoutInMinutes  *int                                            `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                            `json:"timeoutInSeconds,omitempty"`
	Validation        *PostBillCreditNote200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostBillCreditNoteResponse struct {
	ContentType                                string
	StatusCode                                 int
	PostBillCreditNote200ApplicationJSONObject *PostBillCreditNote200ApplicationJSON
}
