package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostItemPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostItemQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostItemSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostItemRequest struct {
	PathParams  PostItemPathParams
	QueryParams PostItemQueryParams
	Request     *shared.Item `request:"mediaType=application/json"`
	Security    PostItemSecurity
}

type PostItem200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostItem200ApplicationJSONChangesTypeEnum string

const (
	PostItem200ApplicationJSONChangesTypeEnumUnknown            PostItem200ApplicationJSONChangesTypeEnum = "Unknown"
	PostItem200ApplicationJSONChangesTypeEnumCreated            PostItem200ApplicationJSONChangesTypeEnum = "Created"
	PostItem200ApplicationJSONChangesTypeEnumModified           PostItem200ApplicationJSONChangesTypeEnum = "Modified"
	PostItem200ApplicationJSONChangesTypeEnumDeleted            PostItem200ApplicationJSONChangesTypeEnum = "Deleted"
	PostItem200ApplicationJSONChangesTypeEnumAttachmentUploaded PostItem200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostItem200ApplicationJSONChanges struct {
	AttachmentID *string                                                  `json:"attachmentId,omitempty"`
	RecordRef    *PostItem200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostItem200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostItem200ApplicationJSONStatusEnum string

const (
	PostItem200ApplicationJSONStatusEnumPending  PostItem200ApplicationJSONStatusEnum = "Pending"
	PostItem200ApplicationJSONStatusEnumFailed   PostItem200ApplicationJSONStatusEnum = "Failed"
	PostItem200ApplicationJSONStatusEnumSuccess  PostItem200ApplicationJSONStatusEnum = "Success"
	PostItem200ApplicationJSONStatusEnumTimedOut PostItem200ApplicationJSONStatusEnum = "TimedOut"
)

type PostItem200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostItem200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostItem200ApplicationJSONValidation struct {
	Errors   []PostItem200ApplicationJSONValidationValidationItem                                                                                                                                                `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostItem200ApplicationJSON struct {
	Changes           []PostItem200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                `json:"companyId"`
	CompletedOnUtc    *time.Time                            `json:"completedOnUtc,omitempty"`
	Data              *shared.Item                          `json:"data,omitempty"`
	DataConnectionKey string                                `json:"dataConnectionKey"`
	DataType          *string                               `json:"dataType,omitempty"`
	ErrorMessage      *string                               `json:"errorMessage,omitempty"`
	PushOperationKey  string                                `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                             `json:"requestedOnUtc"`
	Status            PostItem200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                   `json:"statusCode"`
	TimeoutInMinutes  *int                                  `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                  `json:"timeoutInSeconds,omitempty"`
	Validation        *PostItem200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostItemResponse struct {
	ContentType                      string
	StatusCode                       int
	PostItem200ApplicationJSONObject *PostItem200ApplicationJSON
}
