package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostPaymentPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostPaymentQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostPaymentRequest struct {
	PathParams  PostPaymentPathParams
	QueryParams PostPaymentQueryParams
	Request     *shared.Payment `request:"mediaType=application/json"`
	Security    PostPaymentSecurity
}

type PostPayment200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostPayment200ApplicationJSONChangesTypeEnum string

const (
	PostPayment200ApplicationJSONChangesTypeEnumUnknown            PostPayment200ApplicationJSONChangesTypeEnum = "Unknown"
	PostPayment200ApplicationJSONChangesTypeEnumCreated            PostPayment200ApplicationJSONChangesTypeEnum = "Created"
	PostPayment200ApplicationJSONChangesTypeEnumModified           PostPayment200ApplicationJSONChangesTypeEnum = "Modified"
	PostPayment200ApplicationJSONChangesTypeEnumDeleted            PostPayment200ApplicationJSONChangesTypeEnum = "Deleted"
	PostPayment200ApplicationJSONChangesTypeEnumAttachmentUploaded PostPayment200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostPayment200ApplicationJSONChanges struct {
	AttachmentID *string                                                     `json:"attachmentId,omitempty"`
	RecordRef    *PostPayment200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostPayment200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostPayment200ApplicationJSONStatusEnum string

const (
	PostPayment200ApplicationJSONStatusEnumPending  PostPayment200ApplicationJSONStatusEnum = "Pending"
	PostPayment200ApplicationJSONStatusEnumFailed   PostPayment200ApplicationJSONStatusEnum = "Failed"
	PostPayment200ApplicationJSONStatusEnumSuccess  PostPayment200ApplicationJSONStatusEnum = "Success"
	PostPayment200ApplicationJSONStatusEnumTimedOut PostPayment200ApplicationJSONStatusEnum = "TimedOut"
)

type PostPayment200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostPayment200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostPayment200ApplicationJSONValidation struct {
	Errors   []PostPayment200ApplicationJSONValidationValidationItem                                                                                                                                             `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostPayment200ApplicationJSON struct {
	Changes           []PostPayment200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                   `json:"companyId"`
	CompletedOnUtc    *time.Time                               `json:"completedOnUtc,omitempty"`
	Data              *shared.Payment                          `json:"data,omitempty"`
	DataConnectionKey string                                   `json:"dataConnectionKey"`
	DataType          *string                                  `json:"dataType,omitempty"`
	ErrorMessage      *string                                  `json:"errorMessage,omitempty"`
	PushOperationKey  string                                   `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                `json:"requestedOnUtc"`
	Status            PostPayment200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                      `json:"statusCode"`
	TimeoutInMinutes  *int                                     `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                     `json:"timeoutInSeconds,omitempty"`
	Validation        *PostPayment200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostPaymentResponse struct {
	ContentType                         string
	StatusCode                          int
	PostPayment200ApplicationJSONObject *PostPayment200ApplicationJSON
}
