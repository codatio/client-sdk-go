package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostDirectIncomePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostDirectIncomeQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostDirectIncomeSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostDirectIncomeRequest struct {
	PathParams  PostDirectIncomePathParams
	QueryParams PostDirectIncomeQueryParams
	Request     *shared.DirectIncome `request:"mediaType=application/json"`
	Security    PostDirectIncomeSecurity
}

type PostDirectIncome200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostDirectIncome200ApplicationJSONChangesTypeEnum string

const (
	PostDirectIncome200ApplicationJSONChangesTypeEnumUnknown            PostDirectIncome200ApplicationJSONChangesTypeEnum = "Unknown"
	PostDirectIncome200ApplicationJSONChangesTypeEnumCreated            PostDirectIncome200ApplicationJSONChangesTypeEnum = "Created"
	PostDirectIncome200ApplicationJSONChangesTypeEnumModified           PostDirectIncome200ApplicationJSONChangesTypeEnum = "Modified"
	PostDirectIncome200ApplicationJSONChangesTypeEnumDeleted            PostDirectIncome200ApplicationJSONChangesTypeEnum = "Deleted"
	PostDirectIncome200ApplicationJSONChangesTypeEnumAttachmentUploaded PostDirectIncome200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostDirectIncome200ApplicationJSONChanges struct {
	AttachmentID *string                                                          `json:"attachmentId,omitempty"`
	RecordRef    *PostDirectIncome200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostDirectIncome200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostDirectIncome200ApplicationJSONStatusEnum string

const (
	PostDirectIncome200ApplicationJSONStatusEnumPending  PostDirectIncome200ApplicationJSONStatusEnum = "Pending"
	PostDirectIncome200ApplicationJSONStatusEnumFailed   PostDirectIncome200ApplicationJSONStatusEnum = "Failed"
	PostDirectIncome200ApplicationJSONStatusEnumSuccess  PostDirectIncome200ApplicationJSONStatusEnum = "Success"
	PostDirectIncome200ApplicationJSONStatusEnumTimedOut PostDirectIncome200ApplicationJSONStatusEnum = "TimedOut"
)

type PostDirectIncome200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostDirectIncome200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostDirectIncome200ApplicationJSONValidation struct {
	Errors   []PostDirectIncome200ApplicationJSONValidationValidationItem                                                                                                                                        `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostDirectIncome200ApplicationJSON struct {
	Changes           []PostDirectIncome200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                        `json:"companyId"`
	CompletedOnUtc    *time.Time                                    `json:"completedOnUtc,omitempty"`
	Data              *shared.DirectIncome                          `json:"data,omitempty"`
	DataConnectionKey string                                        `json:"dataConnectionKey"`
	DataType          *string                                       `json:"dataType,omitempty"`
	ErrorMessage      *string                                       `json:"errorMessage,omitempty"`
	PushOperationKey  string                                        `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                     `json:"requestedOnUtc"`
	Status            PostDirectIncome200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                           `json:"statusCode"`
	TimeoutInMinutes  *int                                          `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                          `json:"timeoutInSeconds,omitempty"`
	Validation        *PostDirectIncome200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostDirectIncomeResponse struct {
	ContentType                              string
	StatusCode                               int
	PostDirectIncome200ApplicationJSONObject *PostDirectIncome200ApplicationJSON
}
