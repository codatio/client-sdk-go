package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PutSupplierPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	SupplierID   string `pathParam:"style=simple,explode=false,name=supplierId"`
}

type PutSupplierQueryParams struct {
	ForceUpdate      *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PutSupplierSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PutSupplierRequest struct {
	PathParams  PutSupplierPathParams
	QueryParams PutSupplierQueryParams
	Request     *shared.Supplier `request:"mediaType=application/json"`
	Security    PutSupplierSecurity
}

type PutSupplier200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PutSupplier200ApplicationJSONChangesTypeEnum string

const (
	PutSupplier200ApplicationJSONChangesTypeEnumUnknown            PutSupplier200ApplicationJSONChangesTypeEnum = "Unknown"
	PutSupplier200ApplicationJSONChangesTypeEnumCreated            PutSupplier200ApplicationJSONChangesTypeEnum = "Created"
	PutSupplier200ApplicationJSONChangesTypeEnumModified           PutSupplier200ApplicationJSONChangesTypeEnum = "Modified"
	PutSupplier200ApplicationJSONChangesTypeEnumDeleted            PutSupplier200ApplicationJSONChangesTypeEnum = "Deleted"
	PutSupplier200ApplicationJSONChangesTypeEnumAttachmentUploaded PutSupplier200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PutSupplier200ApplicationJSONChanges struct {
	AttachmentID *string                                                     `json:"attachmentId,omitempty"`
	RecordRef    *PutSupplier200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PutSupplier200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PutSupplier200ApplicationJSONStatusEnum string

const (
	PutSupplier200ApplicationJSONStatusEnumPending  PutSupplier200ApplicationJSONStatusEnum = "Pending"
	PutSupplier200ApplicationJSONStatusEnumFailed   PutSupplier200ApplicationJSONStatusEnum = "Failed"
	PutSupplier200ApplicationJSONStatusEnumSuccess  PutSupplier200ApplicationJSONStatusEnum = "Success"
	PutSupplier200ApplicationJSONStatusEnumTimedOut PutSupplier200ApplicationJSONStatusEnum = "TimedOut"
)

type PutSupplier200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PutSupplier200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PutSupplier200ApplicationJSONValidation struct {
	Errors   []PutSupplier200ApplicationJSONValidationValidationItem                                                                                                                                             `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PutSupplier200ApplicationJSON struct {
	Changes           []PutSupplier200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                   `json:"companyId"`
	CompletedOnUtc    *time.Time                               `json:"completedOnUtc,omitempty"`
	Data              *shared.Supplier                         `json:"data,omitempty"`
	DataConnectionKey string                                   `json:"dataConnectionKey"`
	DataType          *string                                  `json:"dataType,omitempty"`
	ErrorMessage      *string                                  `json:"errorMessage,omitempty"`
	PushOperationKey  string                                   `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                `json:"requestedOnUtc"`
	Status            PutSupplier200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                      `json:"statusCode"`
	TimeoutInMinutes  *int                                     `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                     `json:"timeoutInSeconds,omitempty"`
	Validation        *PutSupplier200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PutSupplierResponse struct {
	ContentType                         string
	StatusCode                          int
	PutSupplier200ApplicationJSONObject *PutSupplier200ApplicationJSON
}
