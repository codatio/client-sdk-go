package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostJournalEntryPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostJournalEntryQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostJournalEntrySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostJournalEntryRequest struct {
	PathParams  PostJournalEntryPathParams
	QueryParams PostJournalEntryQueryParams
	Request     *shared.JournalEntry `request:"mediaType=application/json"`
	Security    PostJournalEntrySecurity
}

type PostJournalEntry200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostJournalEntry200ApplicationJSONChangesTypeEnum string

const (
	PostJournalEntry200ApplicationJSONChangesTypeEnumUnknown            PostJournalEntry200ApplicationJSONChangesTypeEnum = "Unknown"
	PostJournalEntry200ApplicationJSONChangesTypeEnumCreated            PostJournalEntry200ApplicationJSONChangesTypeEnum = "Created"
	PostJournalEntry200ApplicationJSONChangesTypeEnumModified           PostJournalEntry200ApplicationJSONChangesTypeEnum = "Modified"
	PostJournalEntry200ApplicationJSONChangesTypeEnumDeleted            PostJournalEntry200ApplicationJSONChangesTypeEnum = "Deleted"
	PostJournalEntry200ApplicationJSONChangesTypeEnumAttachmentUploaded PostJournalEntry200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostJournalEntry200ApplicationJSONChanges struct {
	AttachmentID *string                                                          `json:"attachmentId,omitempty"`
	RecordRef    *PostJournalEntry200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostJournalEntry200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostJournalEntry200ApplicationJSONStatusEnum string

const (
	PostJournalEntry200ApplicationJSONStatusEnumPending  PostJournalEntry200ApplicationJSONStatusEnum = "Pending"
	PostJournalEntry200ApplicationJSONStatusEnumFailed   PostJournalEntry200ApplicationJSONStatusEnum = "Failed"
	PostJournalEntry200ApplicationJSONStatusEnumSuccess  PostJournalEntry200ApplicationJSONStatusEnum = "Success"
	PostJournalEntry200ApplicationJSONStatusEnumTimedOut PostJournalEntry200ApplicationJSONStatusEnum = "TimedOut"
)

type PostJournalEntry200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostJournalEntry200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostJournalEntry200ApplicationJSONValidation struct {
	Errors   []PostJournalEntry200ApplicationJSONValidationValidationItem                                                                                                                                        `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostJournalEntry200ApplicationJSON struct {
	Changes           []PostJournalEntry200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                        `json:"companyId"`
	CompletedOnUtc    *time.Time                                    `json:"completedOnUtc,omitempty"`
	Data              *shared.JournalEntry                          `json:"data,omitempty"`
	DataConnectionKey string                                        `json:"dataConnectionKey"`
	DataType          *string                                       `json:"dataType,omitempty"`
	ErrorMessage      *string                                       `json:"errorMessage,omitempty"`
	PushOperationKey  string                                        `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                     `json:"requestedOnUtc"`
	Status            PostJournalEntry200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                           `json:"statusCode"`
	TimeoutInMinutes  *int                                          `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                          `json:"timeoutInSeconds,omitempty"`
	Validation        *PostJournalEntry200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostJournalEntryResponse struct {
	ContentType                              string
	StatusCode                               int
	PostJournalEntry200ApplicationJSONObject *PostJournalEntry200ApplicationJSON
}
