package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type UpdateCustomerPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	CustomerID   string `pathParam:"style=simple,explode=false,name=customerId"`
}

type UpdateCustomerQueryParams struct {
	ForceUpdate      *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type UpdateCustomerSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateCustomerRequest struct {
	PathParams  UpdateCustomerPathParams
	QueryParams UpdateCustomerQueryParams
	Request     *shared.Customer `request:"mediaType=application/json"`
	Security    UpdateCustomerSecurity
}

type UpdateCustomer200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type UpdateCustomer200ApplicationJSONChangesTypeEnum string

const (
	UpdateCustomer200ApplicationJSONChangesTypeEnumUnknown            UpdateCustomer200ApplicationJSONChangesTypeEnum = "Unknown"
	UpdateCustomer200ApplicationJSONChangesTypeEnumCreated            UpdateCustomer200ApplicationJSONChangesTypeEnum = "Created"
	UpdateCustomer200ApplicationJSONChangesTypeEnumModified           UpdateCustomer200ApplicationJSONChangesTypeEnum = "Modified"
	UpdateCustomer200ApplicationJSONChangesTypeEnumDeleted            UpdateCustomer200ApplicationJSONChangesTypeEnum = "Deleted"
	UpdateCustomer200ApplicationJSONChangesTypeEnumAttachmentUploaded UpdateCustomer200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type UpdateCustomer200ApplicationJSONChanges struct {
	AttachmentID *string                                                        `json:"attachmentId,omitempty"`
	RecordRef    *UpdateCustomer200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *UpdateCustomer200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type UpdateCustomer200ApplicationJSONStatusEnum string

const (
	UpdateCustomer200ApplicationJSONStatusEnumPending  UpdateCustomer200ApplicationJSONStatusEnum = "Pending"
	UpdateCustomer200ApplicationJSONStatusEnumFailed   UpdateCustomer200ApplicationJSONStatusEnum = "Failed"
	UpdateCustomer200ApplicationJSONStatusEnumSuccess  UpdateCustomer200ApplicationJSONStatusEnum = "Success"
	UpdateCustomer200ApplicationJSONStatusEnumTimedOut UpdateCustomer200ApplicationJSONStatusEnum = "TimedOut"
)

type UpdateCustomer200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// UpdateCustomer200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type UpdateCustomer200ApplicationJSONValidation struct {
	Errors   []UpdateCustomer200ApplicationJSONValidationValidationItem                                                                                                                                          `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type UpdateCustomer200ApplicationJSON struct {
	Changes           []UpdateCustomer200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                      `json:"companyId"`
	CompletedOnUtc    *time.Time                                  `json:"completedOnUtc,omitempty"`
	Data              *shared.Customer                            `json:"data,omitempty"`
	DataConnectionKey string                                      `json:"dataConnectionKey"`
	DataType          *string                                     `json:"dataType,omitempty"`
	ErrorMessage      *string                                     `json:"errorMessage,omitempty"`
	PushOperationKey  string                                      `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                   `json:"requestedOnUtc"`
	Status            UpdateCustomer200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                         `json:"statusCode"`
	TimeoutInMinutes  *int                                        `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                        `json:"timeoutInSeconds,omitempty"`
	Validation        *UpdateCustomer200ApplicationJSONValidation `json:"validation,omitempty"`
}

type UpdateCustomerResponse struct {
	ContentType                            string
	StatusCode                             int
	UpdateCustomer200ApplicationJSONObject *UpdateCustomer200ApplicationJSON
}
