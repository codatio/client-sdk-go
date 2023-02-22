package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostCustomerPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostCustomerQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostCustomerSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostCustomerRequest struct {
	PathParams  PostCustomerPathParams
	QueryParams PostCustomerQueryParams
	Request     *shared.Customer `request:"mediaType=application/json"`
	Security    PostCustomerSecurity
}

type PostCustomer200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostCustomer200ApplicationJSONChangesTypeEnum string

const (
	PostCustomer200ApplicationJSONChangesTypeEnumUnknown            PostCustomer200ApplicationJSONChangesTypeEnum = "Unknown"
	PostCustomer200ApplicationJSONChangesTypeEnumCreated            PostCustomer200ApplicationJSONChangesTypeEnum = "Created"
	PostCustomer200ApplicationJSONChangesTypeEnumModified           PostCustomer200ApplicationJSONChangesTypeEnum = "Modified"
	PostCustomer200ApplicationJSONChangesTypeEnumDeleted            PostCustomer200ApplicationJSONChangesTypeEnum = "Deleted"
	PostCustomer200ApplicationJSONChangesTypeEnumAttachmentUploaded PostCustomer200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostCustomer200ApplicationJSONChanges struct {
	AttachmentID *string                                                      `json:"attachmentId,omitempty"`
	RecordRef    *PostCustomer200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostCustomer200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostCustomer200ApplicationJSONStatusEnum string

const (
	PostCustomer200ApplicationJSONStatusEnumPending  PostCustomer200ApplicationJSONStatusEnum = "Pending"
	PostCustomer200ApplicationJSONStatusEnumFailed   PostCustomer200ApplicationJSONStatusEnum = "Failed"
	PostCustomer200ApplicationJSONStatusEnumSuccess  PostCustomer200ApplicationJSONStatusEnum = "Success"
	PostCustomer200ApplicationJSONStatusEnumTimedOut PostCustomer200ApplicationJSONStatusEnum = "TimedOut"
)

type PostCustomer200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostCustomer200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostCustomer200ApplicationJSONValidation struct {
	Errors   []PostCustomer200ApplicationJSONValidationValidationItem                                                                                                                                            `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostCustomer200ApplicationJSON struct {
	Changes           []PostCustomer200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                    `json:"companyId"`
	CompletedOnUtc    *time.Time                                `json:"completedOnUtc,omitempty"`
	Data              *shared.Customer                          `json:"data,omitempty"`
	DataConnectionKey string                                    `json:"dataConnectionKey"`
	DataType          *string                                   `json:"dataType,omitempty"`
	ErrorMessage      *string                                   `json:"errorMessage,omitempty"`
	PushOperationKey  string                                    `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                 `json:"requestedOnUtc"`
	Status            PostCustomer200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                       `json:"statusCode"`
	TimeoutInMinutes  *int                                      `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                      `json:"timeoutInSeconds,omitempty"`
	Validation        *PostCustomer200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostCustomerResponse struct {
	ContentType                          string
	StatusCode                           int
	PostCustomer200ApplicationJSONObject *PostCustomer200ApplicationJSON
}
