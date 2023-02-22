package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PushJournalPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PushJournalQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PushJournalSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PushJournalRequest struct {
	PathParams  PushJournalPathParams
	QueryParams PushJournalQueryParams
	Request     *shared.JournalInput `request:"mediaType=application/json"`
	Security    PushJournalSecurity
}

type PushJournal200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PushJournal200ApplicationJSONChangesTypeEnum string

const (
	PushJournal200ApplicationJSONChangesTypeEnumUnknown            PushJournal200ApplicationJSONChangesTypeEnum = "Unknown"
	PushJournal200ApplicationJSONChangesTypeEnumCreated            PushJournal200ApplicationJSONChangesTypeEnum = "Created"
	PushJournal200ApplicationJSONChangesTypeEnumModified           PushJournal200ApplicationJSONChangesTypeEnum = "Modified"
	PushJournal200ApplicationJSONChangesTypeEnumDeleted            PushJournal200ApplicationJSONChangesTypeEnum = "Deleted"
	PushJournal200ApplicationJSONChangesTypeEnumAttachmentUploaded PushJournal200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PushJournal200ApplicationJSONChanges struct {
	AttachmentID *string                                                     `json:"attachmentId,omitempty"`
	RecordRef    *PushJournal200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PushJournal200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PushJournal200ApplicationJSONStatusEnum string

const (
	PushJournal200ApplicationJSONStatusEnumPending  PushJournal200ApplicationJSONStatusEnum = "Pending"
	PushJournal200ApplicationJSONStatusEnumFailed   PushJournal200ApplicationJSONStatusEnum = "Failed"
	PushJournal200ApplicationJSONStatusEnumSuccess  PushJournal200ApplicationJSONStatusEnum = "Success"
	PushJournal200ApplicationJSONStatusEnumTimedOut PushJournal200ApplicationJSONStatusEnum = "TimedOut"
)

type PushJournal200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PushJournal200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PushJournal200ApplicationJSONValidation struct {
	Errors   []PushJournal200ApplicationJSONValidationValidationItem                                                                                                                                             `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PushJournal200ApplicationJSON struct {
	Changes           []PushJournal200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                   `json:"companyId"`
	CompletedOnUtc    *time.Time                               `json:"completedOnUtc,omitempty"`
	Data              *shared.Journal                          `json:"data,omitempty"`
	DataConnectionKey string                                   `json:"dataConnectionKey"`
	DataType          *string                                  `json:"dataType,omitempty"`
	ErrorMessage      *string                                  `json:"errorMessage,omitempty"`
	PushOperationKey  string                                   `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                `json:"requestedOnUtc"`
	Status            PushJournal200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                      `json:"statusCode"`
	TimeoutInMinutes  *int                                     `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                     `json:"timeoutInSeconds,omitempty"`
	Validation        *PushJournal200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PushJournalResponse struct {
	ContentType                         string
	StatusCode                          int
	PushJournal200ApplicationJSONObject *PushJournal200ApplicationJSON
}
