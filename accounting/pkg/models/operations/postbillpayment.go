package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostBillPaymentPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostBillPaymentQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostBillPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostBillPaymentRequest struct {
	PathParams  PostBillPaymentPathParams
	QueryParams PostBillPaymentQueryParams
	Request     *shared.BillPayment `request:"mediaType=application/json"`
	Security    PostBillPaymentSecurity
}

type PostBillPayment200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostBillPayment200ApplicationJSONChangesTypeEnum string

const (
	PostBillPayment200ApplicationJSONChangesTypeEnumUnknown            PostBillPayment200ApplicationJSONChangesTypeEnum = "Unknown"
	PostBillPayment200ApplicationJSONChangesTypeEnumCreated            PostBillPayment200ApplicationJSONChangesTypeEnum = "Created"
	PostBillPayment200ApplicationJSONChangesTypeEnumModified           PostBillPayment200ApplicationJSONChangesTypeEnum = "Modified"
	PostBillPayment200ApplicationJSONChangesTypeEnumDeleted            PostBillPayment200ApplicationJSONChangesTypeEnum = "Deleted"
	PostBillPayment200ApplicationJSONChangesTypeEnumAttachmentUploaded PostBillPayment200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostBillPayment200ApplicationJSONChanges struct {
	AttachmentID *string                                                         `json:"attachmentId,omitempty"`
	RecordRef    *PostBillPayment200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostBillPayment200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostBillPayment200ApplicationJSONStatusEnum string

const (
	PostBillPayment200ApplicationJSONStatusEnumPending  PostBillPayment200ApplicationJSONStatusEnum = "Pending"
	PostBillPayment200ApplicationJSONStatusEnumFailed   PostBillPayment200ApplicationJSONStatusEnum = "Failed"
	PostBillPayment200ApplicationJSONStatusEnumSuccess  PostBillPayment200ApplicationJSONStatusEnum = "Success"
	PostBillPayment200ApplicationJSONStatusEnumTimedOut PostBillPayment200ApplicationJSONStatusEnum = "TimedOut"
)

type PostBillPayment200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostBillPayment200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostBillPayment200ApplicationJSONValidation struct {
	Errors   []PostBillPayment200ApplicationJSONValidationValidationItem                                                                                                                                         `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostBillPayment200ApplicationJSON struct {
	Changes           []PostBillPayment200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                       `json:"companyId"`
	CompletedOnUtc    *time.Time                                   `json:"completedOnUtc,omitempty"`
	Data              *shared.BillPayment                          `json:"data,omitempty"`
	DataConnectionKey string                                       `json:"dataConnectionKey"`
	DataType          *string                                      `json:"dataType,omitempty"`
	ErrorMessage      *string                                      `json:"errorMessage,omitempty"`
	PushOperationKey  string                                       `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                    `json:"requestedOnUtc"`
	Status            PostBillPayment200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                          `json:"statusCode"`
	TimeoutInMinutes  *int                                         `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                         `json:"timeoutInSeconds,omitempty"`
	Validation        *PostBillPayment200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostBillPaymentResponse struct {
	ContentType                             string
	StatusCode                              int
	PostBillPayment200ApplicationJSONObject *PostBillPayment200ApplicationJSON
}
