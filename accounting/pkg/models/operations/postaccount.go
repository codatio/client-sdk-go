package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostAccountPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostAccountQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostAccountSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostAccountRequest struct {
	PathParams  PostAccountPathParams
	QueryParams PostAccountQueryParams
	Request     *shared.Account `request:"mediaType=application/json"`
	Security    PostAccountSecurity
}

type PostAccount200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostAccount200ApplicationJSONChangesTypeEnum string

const (
	PostAccount200ApplicationJSONChangesTypeEnumUnknown            PostAccount200ApplicationJSONChangesTypeEnum = "Unknown"
	PostAccount200ApplicationJSONChangesTypeEnumCreated            PostAccount200ApplicationJSONChangesTypeEnum = "Created"
	PostAccount200ApplicationJSONChangesTypeEnumModified           PostAccount200ApplicationJSONChangesTypeEnum = "Modified"
	PostAccount200ApplicationJSONChangesTypeEnumDeleted            PostAccount200ApplicationJSONChangesTypeEnum = "Deleted"
	PostAccount200ApplicationJSONChangesTypeEnumAttachmentUploaded PostAccount200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostAccount200ApplicationJSONChanges struct {
	AttachmentID *string                                                     `json:"attachmentId,omitempty"`
	RecordRef    *PostAccount200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostAccount200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostAccount200ApplicationJSONStatusEnum string

const (
	PostAccount200ApplicationJSONStatusEnumPending  PostAccount200ApplicationJSONStatusEnum = "Pending"
	PostAccount200ApplicationJSONStatusEnumFailed   PostAccount200ApplicationJSONStatusEnum = "Failed"
	PostAccount200ApplicationJSONStatusEnumSuccess  PostAccount200ApplicationJSONStatusEnum = "Success"
	PostAccount200ApplicationJSONStatusEnumTimedOut PostAccount200ApplicationJSONStatusEnum = "TimedOut"
)

type PostAccount200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostAccount200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostAccount200ApplicationJSONValidation struct {
	Errors   []PostAccount200ApplicationJSONValidationValidationItem                                                                                                                                             `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostAccount200ApplicationJSON struct {
	Changes           []PostAccount200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                   `json:"companyId"`
	CompletedOnUtc    *time.Time                               `json:"completedOnUtc,omitempty"`
	Data              *shared.Account                          `json:"data,omitempty"`
	DataConnectionKey string                                   `json:"dataConnectionKey"`
	DataType          *string                                  `json:"dataType,omitempty"`
	ErrorMessage      *string                                  `json:"errorMessage,omitempty"`
	PushOperationKey  string                                   `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                `json:"requestedOnUtc"`
	Status            PostAccount200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                      `json:"statusCode"`
	TimeoutInMinutes  *int                                     `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                     `json:"timeoutInSeconds,omitempty"`
	Validation        *PostAccount200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostAccountResponse struct {
	ContentType                         string
	StatusCode                          int
	PostAccount200ApplicationJSONObject *PostAccount200ApplicationJSON
}
