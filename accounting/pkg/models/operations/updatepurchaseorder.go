package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type UpdatePurchaseOrderPathParams struct {
	CompanyID       string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID    string `pathParam:"style=simple,explode=false,name=connectionId"`
	PurchaseOrderID string `pathParam:"style=simple,explode=false,name=purchaseOrderId"`
}

type UpdatePurchaseOrderQueryParams struct {
	ForceUpdate      *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type UpdatePurchaseOrderSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdatePurchaseOrderRequest struct {
	PathParams  UpdatePurchaseOrderPathParams
	QueryParams UpdatePurchaseOrderQueryParams
	Request     *shared.PurchaseOrder `request:"mediaType=application/json"`
	Security    UpdatePurchaseOrderSecurity
}

type UpdatePurchaseOrder200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type UpdatePurchaseOrder200ApplicationJSONChangesTypeEnum string

const (
	UpdatePurchaseOrder200ApplicationJSONChangesTypeEnumUnknown            UpdatePurchaseOrder200ApplicationJSONChangesTypeEnum = "Unknown"
	UpdatePurchaseOrder200ApplicationJSONChangesTypeEnumCreated            UpdatePurchaseOrder200ApplicationJSONChangesTypeEnum = "Created"
	UpdatePurchaseOrder200ApplicationJSONChangesTypeEnumModified           UpdatePurchaseOrder200ApplicationJSONChangesTypeEnum = "Modified"
	UpdatePurchaseOrder200ApplicationJSONChangesTypeEnumDeleted            UpdatePurchaseOrder200ApplicationJSONChangesTypeEnum = "Deleted"
	UpdatePurchaseOrder200ApplicationJSONChangesTypeEnumAttachmentUploaded UpdatePurchaseOrder200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type UpdatePurchaseOrder200ApplicationJSONChanges struct {
	AttachmentID *string                                                             `json:"attachmentId,omitempty"`
	RecordRef    *UpdatePurchaseOrder200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *UpdatePurchaseOrder200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type UpdatePurchaseOrder200ApplicationJSONStatusEnum string

const (
	UpdatePurchaseOrder200ApplicationJSONStatusEnumPending  UpdatePurchaseOrder200ApplicationJSONStatusEnum = "Pending"
	UpdatePurchaseOrder200ApplicationJSONStatusEnumFailed   UpdatePurchaseOrder200ApplicationJSONStatusEnum = "Failed"
	UpdatePurchaseOrder200ApplicationJSONStatusEnumSuccess  UpdatePurchaseOrder200ApplicationJSONStatusEnum = "Success"
	UpdatePurchaseOrder200ApplicationJSONStatusEnumTimedOut UpdatePurchaseOrder200ApplicationJSONStatusEnum = "TimedOut"
)

type UpdatePurchaseOrder200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// UpdatePurchaseOrder200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type UpdatePurchaseOrder200ApplicationJSONValidation struct {
	Errors   []UpdatePurchaseOrder200ApplicationJSONValidationValidationItem                                                                                                                                     `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type UpdatePurchaseOrder200ApplicationJSON struct {
	Changes           []UpdatePurchaseOrder200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                           `json:"companyId"`
	CompletedOnUtc    *time.Time                                       `json:"completedOnUtc,omitempty"`
	Data              *shared.PurchaseOrder                            `json:"data,omitempty"`
	DataConnectionKey string                                           `json:"dataConnectionKey"`
	DataType          *string                                          `json:"dataType,omitempty"`
	ErrorMessage      *string                                          `json:"errorMessage,omitempty"`
	PushOperationKey  string                                           `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                        `json:"requestedOnUtc"`
	Status            UpdatePurchaseOrder200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                              `json:"statusCode"`
	TimeoutInMinutes  *int                                             `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                             `json:"timeoutInSeconds,omitempty"`
	Validation        *UpdatePurchaseOrder200ApplicationJSONValidation `json:"validation,omitempty"`
}

type UpdatePurchaseOrderResponse struct {
	ContentType                                 string
	StatusCode                                  int
	UpdatePurchaseOrder200ApplicationJSONObject *UpdatePurchaseOrder200ApplicationJSON
}
