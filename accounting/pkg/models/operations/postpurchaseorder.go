package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostPurchaseOrderPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostPurchaseOrderQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostPurchaseOrderSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostPurchaseOrderRequest struct {
	PathParams  PostPurchaseOrderPathParams
	QueryParams PostPurchaseOrderQueryParams
	Request     *shared.PurchaseOrder `request:"mediaType=application/json"`
	Security    PostPurchaseOrderSecurity
}

type PostPurchaseOrder200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostPurchaseOrder200ApplicationJSONChangesTypeEnum string

const (
	PostPurchaseOrder200ApplicationJSONChangesTypeEnumUnknown            PostPurchaseOrder200ApplicationJSONChangesTypeEnum = "Unknown"
	PostPurchaseOrder200ApplicationJSONChangesTypeEnumCreated            PostPurchaseOrder200ApplicationJSONChangesTypeEnum = "Created"
	PostPurchaseOrder200ApplicationJSONChangesTypeEnumModified           PostPurchaseOrder200ApplicationJSONChangesTypeEnum = "Modified"
	PostPurchaseOrder200ApplicationJSONChangesTypeEnumDeleted            PostPurchaseOrder200ApplicationJSONChangesTypeEnum = "Deleted"
	PostPurchaseOrder200ApplicationJSONChangesTypeEnumAttachmentUploaded PostPurchaseOrder200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostPurchaseOrder200ApplicationJSONChanges struct {
	AttachmentID *string                                                           `json:"attachmentId,omitempty"`
	RecordRef    *PostPurchaseOrder200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostPurchaseOrder200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostPurchaseOrder200ApplicationJSONStatusEnum string

const (
	PostPurchaseOrder200ApplicationJSONStatusEnumPending  PostPurchaseOrder200ApplicationJSONStatusEnum = "Pending"
	PostPurchaseOrder200ApplicationJSONStatusEnumFailed   PostPurchaseOrder200ApplicationJSONStatusEnum = "Failed"
	PostPurchaseOrder200ApplicationJSONStatusEnumSuccess  PostPurchaseOrder200ApplicationJSONStatusEnum = "Success"
	PostPurchaseOrder200ApplicationJSONStatusEnumTimedOut PostPurchaseOrder200ApplicationJSONStatusEnum = "TimedOut"
)

type PostPurchaseOrder200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostPurchaseOrder200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostPurchaseOrder200ApplicationJSONValidation struct {
	Errors   []PostPurchaseOrder200ApplicationJSONValidationValidationItem                                                                                                                                       `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostPurchaseOrder200ApplicationJSON struct {
	Changes           []PostPurchaseOrder200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                         `json:"companyId"`
	CompletedOnUtc    *time.Time                                     `json:"completedOnUtc,omitempty"`
	Data              *shared.PurchaseOrder                          `json:"data,omitempty"`
	DataConnectionKey string                                         `json:"dataConnectionKey"`
	DataType          *string                                        `json:"dataType,omitempty"`
	ErrorMessage      *string                                        `json:"errorMessage,omitempty"`
	PushOperationKey  string                                         `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                      `json:"requestedOnUtc"`
	Status            PostPurchaseOrder200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                            `json:"statusCode"`
	TimeoutInMinutes  *int                                           `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                           `json:"timeoutInSeconds,omitempty"`
	Validation        *PostPurchaseOrder200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostPurchaseOrderResponse struct {
	ContentType                               string
	StatusCode                                int
	PostPurchaseOrder200ApplicationJSONObject *PostPurchaseOrder200ApplicationJSON
}
