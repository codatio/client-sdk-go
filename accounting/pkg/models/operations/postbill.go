package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostBillPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostBillQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostBillSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostBillRequest struct {
	PathParams  PostBillPathParams
	QueryParams PostBillQueryParams
	Request     *shared.Bill `request:"mediaType=application/json"`
	Security    PostBillSecurity
}

type PostBill200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostBill200ApplicationJSONChangesTypeEnum string

const (
	PostBill200ApplicationJSONChangesTypeEnumUnknown            PostBill200ApplicationJSONChangesTypeEnum = "Unknown"
	PostBill200ApplicationJSONChangesTypeEnumCreated            PostBill200ApplicationJSONChangesTypeEnum = "Created"
	PostBill200ApplicationJSONChangesTypeEnumModified           PostBill200ApplicationJSONChangesTypeEnum = "Modified"
	PostBill200ApplicationJSONChangesTypeEnumDeleted            PostBill200ApplicationJSONChangesTypeEnum = "Deleted"
	PostBill200ApplicationJSONChangesTypeEnumAttachmentUploaded PostBill200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostBill200ApplicationJSONChanges struct {
	AttachmentID *string                                                  `json:"attachmentId,omitempty"`
	RecordRef    *PostBill200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostBill200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostBill200ApplicationJSONStatusEnum string

const (
	PostBill200ApplicationJSONStatusEnumPending  PostBill200ApplicationJSONStatusEnum = "Pending"
	PostBill200ApplicationJSONStatusEnumFailed   PostBill200ApplicationJSONStatusEnum = "Failed"
	PostBill200ApplicationJSONStatusEnumSuccess  PostBill200ApplicationJSONStatusEnum = "Success"
	PostBill200ApplicationJSONStatusEnumTimedOut PostBill200ApplicationJSONStatusEnum = "TimedOut"
)

type PostBill200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostBill200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostBill200ApplicationJSONValidation struct {
	Errors   []PostBill200ApplicationJSONValidationValidationItem                                                                                                                                                `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostBill200ApplicationJSON struct {
	Changes           []PostBill200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                `json:"companyId"`
	CompletedOnUtc    *time.Time                            `json:"completedOnUtc,omitempty"`
	Data              *shared.Bill                          `json:"data,omitempty"`
	DataConnectionKey string                                `json:"dataConnectionKey"`
	DataType          *string                               `json:"dataType,omitempty"`
	ErrorMessage      *string                               `json:"errorMessage,omitempty"`
	PushOperationKey  string                                `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                             `json:"requestedOnUtc"`
	Status            PostBill200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                   `json:"statusCode"`
	TimeoutInMinutes  *int                                  `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                  `json:"timeoutInSeconds,omitempty"`
	Validation        *PostBill200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostBillResponse struct {
	ContentType                      string
	StatusCode                       int
	PostBill200ApplicationJSONObject *PostBill200ApplicationJSON
}
