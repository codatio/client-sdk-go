package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostTransferPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostTransferSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostTransferRequest struct {
	PathParams PostTransferPathParams
	Request    *shared.Transfer `request:"mediaType=application/json"`
	Security   PostTransferSecurity
}

type PostTransfer200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostTransfer200ApplicationJSONChangesTypeEnum string

const (
	PostTransfer200ApplicationJSONChangesTypeEnumUnknown            PostTransfer200ApplicationJSONChangesTypeEnum = "Unknown"
	PostTransfer200ApplicationJSONChangesTypeEnumCreated            PostTransfer200ApplicationJSONChangesTypeEnum = "Created"
	PostTransfer200ApplicationJSONChangesTypeEnumModified           PostTransfer200ApplicationJSONChangesTypeEnum = "Modified"
	PostTransfer200ApplicationJSONChangesTypeEnumDeleted            PostTransfer200ApplicationJSONChangesTypeEnum = "Deleted"
	PostTransfer200ApplicationJSONChangesTypeEnumAttachmentUploaded PostTransfer200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostTransfer200ApplicationJSONChanges struct {
	AttachmentID *string                                                      `json:"attachmentId,omitempty"`
	RecordRef    *PostTransfer200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostTransfer200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostTransfer200ApplicationJSONStatusEnum string

const (
	PostTransfer200ApplicationJSONStatusEnumPending  PostTransfer200ApplicationJSONStatusEnum = "Pending"
	PostTransfer200ApplicationJSONStatusEnumFailed   PostTransfer200ApplicationJSONStatusEnum = "Failed"
	PostTransfer200ApplicationJSONStatusEnumSuccess  PostTransfer200ApplicationJSONStatusEnum = "Success"
	PostTransfer200ApplicationJSONStatusEnumTimedOut PostTransfer200ApplicationJSONStatusEnum = "TimedOut"
)

type PostTransfer200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostTransfer200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostTransfer200ApplicationJSONValidation struct {
	Errors   []PostTransfer200ApplicationJSONValidationValidationItem                                                                                                                                            `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostTransfer200ApplicationJSON struct {
	Changes           []PostTransfer200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                    `json:"companyId"`
	CompletedOnUtc    *time.Time                                `json:"completedOnUtc,omitempty"`
	Data              *shared.Transfer                          `json:"data,omitempty"`
	DataConnectionKey string                                    `json:"dataConnectionKey"`
	DataType          *string                                   `json:"dataType,omitempty"`
	ErrorMessage      *string                                   `json:"errorMessage,omitempty"`
	PushOperationKey  string                                    `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                 `json:"requestedOnUtc"`
	Status            PostTransfer200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                       `json:"statusCode"`
	TimeoutInMinutes  *int                                      `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                      `json:"timeoutInSeconds,omitempty"`
	Validation        *PostTransfer200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostTransferResponse struct {
	ContentType                          string
	StatusCode                           int
	PostTransfer200ApplicationJSONObject *PostTransfer200ApplicationJSON
}
