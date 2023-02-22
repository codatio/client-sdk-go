package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type UpdateInvoicePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	InvoiceID    string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type UpdateInvoiceQueryParams struct {
	ForceUpdate      *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type UpdateInvoiceSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateInvoiceRequest struct {
	PathParams  UpdateInvoicePathParams
	QueryParams UpdateInvoiceQueryParams
	Request     *shared.Invoice `request:"mediaType=application/json"`
	Security    UpdateInvoiceSecurity
}

type UpdateInvoice200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type UpdateInvoice200ApplicationJSONChangesTypeEnum string

const (
	UpdateInvoice200ApplicationJSONChangesTypeEnumUnknown            UpdateInvoice200ApplicationJSONChangesTypeEnum = "Unknown"
	UpdateInvoice200ApplicationJSONChangesTypeEnumCreated            UpdateInvoice200ApplicationJSONChangesTypeEnum = "Created"
	UpdateInvoice200ApplicationJSONChangesTypeEnumModified           UpdateInvoice200ApplicationJSONChangesTypeEnum = "Modified"
	UpdateInvoice200ApplicationJSONChangesTypeEnumDeleted            UpdateInvoice200ApplicationJSONChangesTypeEnum = "Deleted"
	UpdateInvoice200ApplicationJSONChangesTypeEnumAttachmentUploaded UpdateInvoice200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type UpdateInvoice200ApplicationJSONChanges struct {
	AttachmentID *string                                                       `json:"attachmentId,omitempty"`
	RecordRef    *UpdateInvoice200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *UpdateInvoice200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type UpdateInvoice200ApplicationJSONStatusEnum string

const (
	UpdateInvoice200ApplicationJSONStatusEnumPending  UpdateInvoice200ApplicationJSONStatusEnum = "Pending"
	UpdateInvoice200ApplicationJSONStatusEnumFailed   UpdateInvoice200ApplicationJSONStatusEnum = "Failed"
	UpdateInvoice200ApplicationJSONStatusEnumSuccess  UpdateInvoice200ApplicationJSONStatusEnum = "Success"
	UpdateInvoice200ApplicationJSONStatusEnumTimedOut UpdateInvoice200ApplicationJSONStatusEnum = "TimedOut"
)

type UpdateInvoice200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// UpdateInvoice200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type UpdateInvoice200ApplicationJSONValidation struct {
	Errors   []UpdateInvoice200ApplicationJSONValidationValidationItem                                                                                                                                           `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type UpdateInvoice200ApplicationJSON struct {
	Changes           []UpdateInvoice200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                     `json:"companyId"`
	CompletedOnUtc    *time.Time                                 `json:"completedOnUtc,omitempty"`
	Data              *shared.Invoice                            `json:"data,omitempty"`
	DataConnectionKey string                                     `json:"dataConnectionKey"`
	DataType          *string                                    `json:"dataType,omitempty"`
	ErrorMessage      *string                                    `json:"errorMessage,omitempty"`
	PushOperationKey  string                                     `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                  `json:"requestedOnUtc"`
	Status            UpdateInvoice200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                        `json:"statusCode"`
	TimeoutInMinutes  *int                                       `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                       `json:"timeoutInSeconds,omitempty"`
	Validation        *UpdateInvoice200ApplicationJSONValidation `json:"validation,omitempty"`
}

type UpdateInvoiceResponse struct {
	ContentType                           string
	StatusCode                            int
	UpdateInvoice200ApplicationJSONObject *UpdateInvoice200ApplicationJSON
}
