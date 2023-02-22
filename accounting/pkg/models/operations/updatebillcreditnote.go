package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type UpdateBillCreditNotePathParams struct {
	BillCreditNoteID string `pathParam:"style=simple,explode=false,name=billCreditNoteId"`
	CompanyID        string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type UpdateBillCreditNoteQueryParams struct {
	ForceUpdate      *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type UpdateBillCreditNoteSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateBillCreditNoteRequest struct {
	PathParams  UpdateBillCreditNotePathParams
	QueryParams UpdateBillCreditNoteQueryParams
	Request     *shared.BillCreditNote `request:"mediaType=application/json"`
	Security    UpdateBillCreditNoteSecurity
}

type UpdateBillCreditNote200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type UpdateBillCreditNote200ApplicationJSONChangesTypeEnum string

const (
	UpdateBillCreditNote200ApplicationJSONChangesTypeEnumUnknown            UpdateBillCreditNote200ApplicationJSONChangesTypeEnum = "Unknown"
	UpdateBillCreditNote200ApplicationJSONChangesTypeEnumCreated            UpdateBillCreditNote200ApplicationJSONChangesTypeEnum = "Created"
	UpdateBillCreditNote200ApplicationJSONChangesTypeEnumModified           UpdateBillCreditNote200ApplicationJSONChangesTypeEnum = "Modified"
	UpdateBillCreditNote200ApplicationJSONChangesTypeEnumDeleted            UpdateBillCreditNote200ApplicationJSONChangesTypeEnum = "Deleted"
	UpdateBillCreditNote200ApplicationJSONChangesTypeEnumAttachmentUploaded UpdateBillCreditNote200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type UpdateBillCreditNote200ApplicationJSONChanges struct {
	AttachmentID *string                                                              `json:"attachmentId,omitempty"`
	RecordRef    *UpdateBillCreditNote200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *UpdateBillCreditNote200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type UpdateBillCreditNote200ApplicationJSONStatusEnum string

const (
	UpdateBillCreditNote200ApplicationJSONStatusEnumPending  UpdateBillCreditNote200ApplicationJSONStatusEnum = "Pending"
	UpdateBillCreditNote200ApplicationJSONStatusEnumFailed   UpdateBillCreditNote200ApplicationJSONStatusEnum = "Failed"
	UpdateBillCreditNote200ApplicationJSONStatusEnumSuccess  UpdateBillCreditNote200ApplicationJSONStatusEnum = "Success"
	UpdateBillCreditNote200ApplicationJSONStatusEnumTimedOut UpdateBillCreditNote200ApplicationJSONStatusEnum = "TimedOut"
)

type UpdateBillCreditNote200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// UpdateBillCreditNote200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type UpdateBillCreditNote200ApplicationJSONValidation struct {
	Errors   []UpdateBillCreditNote200ApplicationJSONValidationValidationItem                                                                                                                                    `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type UpdateBillCreditNote200ApplicationJSON struct {
	Changes           []UpdateBillCreditNote200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                            `json:"companyId"`
	CompletedOnUtc    *time.Time                                        `json:"completedOnUtc,omitempty"`
	Data              *shared.BillCreditNote                            `json:"data,omitempty"`
	DataConnectionKey string                                            `json:"dataConnectionKey"`
	DataType          *string                                           `json:"dataType,omitempty"`
	ErrorMessage      *string                                           `json:"errorMessage,omitempty"`
	PushOperationKey  string                                            `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                         `json:"requestedOnUtc"`
	Status            UpdateBillCreditNote200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                               `json:"statusCode"`
	TimeoutInMinutes  *int                                              `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                              `json:"timeoutInSeconds,omitempty"`
	Validation        *UpdateBillCreditNote200ApplicationJSONValidation `json:"validation,omitempty"`
}

type UpdateBillCreditNoteResponse struct {
	ContentType                                  string
	StatusCode                                   int
	UpdateBillCreditNote200ApplicationJSONObject *UpdateBillCreditNote200ApplicationJSON
}
