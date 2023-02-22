package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostDirectCostPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostDirectCostQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostDirectCostSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostDirectCostRequest struct {
	PathParams  PostDirectCostPathParams
	QueryParams PostDirectCostQueryParams
	Request     *shared.DirectCost `request:"mediaType=application/json"`
	Security    PostDirectCostSecurity
}

type PostDirectCost200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostDirectCost200ApplicationJSONChangesTypeEnum string

const (
	PostDirectCost200ApplicationJSONChangesTypeEnumUnknown            PostDirectCost200ApplicationJSONChangesTypeEnum = "Unknown"
	PostDirectCost200ApplicationJSONChangesTypeEnumCreated            PostDirectCost200ApplicationJSONChangesTypeEnum = "Created"
	PostDirectCost200ApplicationJSONChangesTypeEnumModified           PostDirectCost200ApplicationJSONChangesTypeEnum = "Modified"
	PostDirectCost200ApplicationJSONChangesTypeEnumDeleted            PostDirectCost200ApplicationJSONChangesTypeEnum = "Deleted"
	PostDirectCost200ApplicationJSONChangesTypeEnumAttachmentUploaded PostDirectCost200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostDirectCost200ApplicationJSONChanges struct {
	AttachmentID *string                                                        `json:"attachmentId,omitempty"`
	RecordRef    *PostDirectCost200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostDirectCost200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostDirectCost200ApplicationJSONStatusEnum string

const (
	PostDirectCost200ApplicationJSONStatusEnumPending  PostDirectCost200ApplicationJSONStatusEnum = "Pending"
	PostDirectCost200ApplicationJSONStatusEnumFailed   PostDirectCost200ApplicationJSONStatusEnum = "Failed"
	PostDirectCost200ApplicationJSONStatusEnumSuccess  PostDirectCost200ApplicationJSONStatusEnum = "Success"
	PostDirectCost200ApplicationJSONStatusEnumTimedOut PostDirectCost200ApplicationJSONStatusEnum = "TimedOut"
)

type PostDirectCost200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostDirectCost200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostDirectCost200ApplicationJSONValidation struct {
	Errors   []PostDirectCost200ApplicationJSONValidationValidationItem                                                                                                                                          `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostDirectCost200ApplicationJSON struct {
	Changes           []PostDirectCost200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                      `json:"companyId"`
	CompletedOnUtc    *time.Time                                  `json:"completedOnUtc,omitempty"`
	Data              *shared.DirectCost                          `json:"data,omitempty"`
	DataConnectionKey string                                      `json:"dataConnectionKey"`
	DataType          *string                                     `json:"dataType,omitempty"`
	ErrorMessage      *string                                     `json:"errorMessage,omitempty"`
	PushOperationKey  string                                      `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                   `json:"requestedOnUtc"`
	Status            PostDirectCost200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                         `json:"statusCode"`
	TimeoutInMinutes  *int                                        `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                        `json:"timeoutInSeconds,omitempty"`
	Validation        *PostDirectCost200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostDirectCostResponse struct {
	ContentType                            string
	StatusCode                             int
	PostDirectCost200ApplicationJSONObject *PostDirectCost200ApplicationJSON
}
