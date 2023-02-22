package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostInvoicePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostInvoiceQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostInvoiceSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostInvoiceRequest struct {
	PathParams  PostInvoicePathParams
	QueryParams PostInvoiceQueryParams
	Request     *shared.Invoice `request:"mediaType=application/json"`
	Security    PostInvoiceSecurity
}

type PostInvoice200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostInvoice200ApplicationJSONChangesTypeEnum string

const (
	PostInvoice200ApplicationJSONChangesTypeEnumUnknown            PostInvoice200ApplicationJSONChangesTypeEnum = "Unknown"
	PostInvoice200ApplicationJSONChangesTypeEnumCreated            PostInvoice200ApplicationJSONChangesTypeEnum = "Created"
	PostInvoice200ApplicationJSONChangesTypeEnumModified           PostInvoice200ApplicationJSONChangesTypeEnum = "Modified"
	PostInvoice200ApplicationJSONChangesTypeEnumDeleted            PostInvoice200ApplicationJSONChangesTypeEnum = "Deleted"
	PostInvoice200ApplicationJSONChangesTypeEnumAttachmentUploaded PostInvoice200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostInvoice200ApplicationJSONChanges struct {
	AttachmentID *string                                                     `json:"attachmentId,omitempty"`
	RecordRef    *PostInvoice200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostInvoice200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostInvoice200ApplicationJSONStatusEnum string

const (
	PostInvoice200ApplicationJSONStatusEnumPending  PostInvoice200ApplicationJSONStatusEnum = "Pending"
	PostInvoice200ApplicationJSONStatusEnumFailed   PostInvoice200ApplicationJSONStatusEnum = "Failed"
	PostInvoice200ApplicationJSONStatusEnumSuccess  PostInvoice200ApplicationJSONStatusEnum = "Success"
	PostInvoice200ApplicationJSONStatusEnumTimedOut PostInvoice200ApplicationJSONStatusEnum = "TimedOut"
)

type PostInvoice200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostInvoice200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostInvoice200ApplicationJSONValidation struct {
	Errors   []PostInvoice200ApplicationJSONValidationValidationItem                                                                                                                                             `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostInvoice200ApplicationJSON struct {
	Changes           []PostInvoice200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                   `json:"companyId"`
	CompletedOnUtc    *time.Time                               `json:"completedOnUtc,omitempty"`
	Data              *shared.Invoice                          `json:"data,omitempty"`
	DataConnectionKey string                                   `json:"dataConnectionKey"`
	DataType          *string                                  `json:"dataType,omitempty"`
	ErrorMessage      *string                                  `json:"errorMessage,omitempty"`
	PushOperationKey  string                                   `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                `json:"requestedOnUtc"`
	Status            PostInvoice200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                      `json:"statusCode"`
	TimeoutInMinutes  *int                                     `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                     `json:"timeoutInSeconds,omitempty"`
	Validation        *PostInvoice200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostInvoiceResponse struct {
	ContentType                         string
	StatusCode                          int
	PostInvoice200ApplicationJSONObject *PostInvoice200ApplicationJSON
}
