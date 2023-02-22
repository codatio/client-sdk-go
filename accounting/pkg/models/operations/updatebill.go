package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type UpdateBillPathParams struct {
	BillID       string `pathParam:"style=simple,explode=false,name=billId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type UpdateBillQueryParams struct {
	ForceUpdate      *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type UpdateBillSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateBillRequest struct {
	PathParams  UpdateBillPathParams
	QueryParams UpdateBillQueryParams
	Request     *shared.Bill `request:"mediaType=application/json"`
	Security    UpdateBillSecurity
}

type UpdateBill200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type UpdateBill200ApplicationJSONChangesTypeEnum string

const (
	UpdateBill200ApplicationJSONChangesTypeEnumUnknown            UpdateBill200ApplicationJSONChangesTypeEnum = "Unknown"
	UpdateBill200ApplicationJSONChangesTypeEnumCreated            UpdateBill200ApplicationJSONChangesTypeEnum = "Created"
	UpdateBill200ApplicationJSONChangesTypeEnumModified           UpdateBill200ApplicationJSONChangesTypeEnum = "Modified"
	UpdateBill200ApplicationJSONChangesTypeEnumDeleted            UpdateBill200ApplicationJSONChangesTypeEnum = "Deleted"
	UpdateBill200ApplicationJSONChangesTypeEnumAttachmentUploaded UpdateBill200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type UpdateBill200ApplicationJSONChanges struct {
	AttachmentID *string                                                    `json:"attachmentId,omitempty"`
	RecordRef    *UpdateBill200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *UpdateBill200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type UpdateBill200ApplicationJSONStatusEnum string

const (
	UpdateBill200ApplicationJSONStatusEnumPending  UpdateBill200ApplicationJSONStatusEnum = "Pending"
	UpdateBill200ApplicationJSONStatusEnumFailed   UpdateBill200ApplicationJSONStatusEnum = "Failed"
	UpdateBill200ApplicationJSONStatusEnumSuccess  UpdateBill200ApplicationJSONStatusEnum = "Success"
	UpdateBill200ApplicationJSONStatusEnumTimedOut UpdateBill200ApplicationJSONStatusEnum = "TimedOut"
)

type UpdateBill200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// UpdateBill200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type UpdateBill200ApplicationJSONValidation struct {
	Errors   []UpdateBill200ApplicationJSONValidationValidationItem                                                                                                                                              `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type UpdateBill200ApplicationJSON struct {
	Changes           []UpdateBill200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                  `json:"companyId"`
	CompletedOnUtc    *time.Time                              `json:"completedOnUtc,omitempty"`
	Data              *shared.Bill                            `json:"data,omitempty"`
	DataConnectionKey string                                  `json:"dataConnectionKey"`
	DataType          *string                                 `json:"dataType,omitempty"`
	ErrorMessage      *string                                 `json:"errorMessage,omitempty"`
	PushOperationKey  string                                  `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                               `json:"requestedOnUtc"`
	Status            UpdateBill200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                     `json:"statusCode"`
	TimeoutInMinutes  *int                                    `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                    `json:"timeoutInSeconds,omitempty"`
	Validation        *UpdateBill200ApplicationJSONValidation `json:"validation,omitempty"`
}

type UpdateBillResponse struct {
	ContentType                        string
	StatusCode                         int
	UpdateBill200ApplicationJSONObject *UpdateBill200ApplicationJSON
}
