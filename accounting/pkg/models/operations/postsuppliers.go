package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostSuppliersPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostSuppliersQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostSuppliersSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostSuppliersRequest struct {
	PathParams  PostSuppliersPathParams
	QueryParams PostSuppliersQueryParams
	Request     *shared.Supplier `request:"mediaType=application/json"`
	Security    PostSuppliersSecurity
}

type PostSuppliers200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostSuppliers200ApplicationJSONChangesTypeEnum string

const (
	PostSuppliers200ApplicationJSONChangesTypeEnumUnknown            PostSuppliers200ApplicationJSONChangesTypeEnum = "Unknown"
	PostSuppliers200ApplicationJSONChangesTypeEnumCreated            PostSuppliers200ApplicationJSONChangesTypeEnum = "Created"
	PostSuppliers200ApplicationJSONChangesTypeEnumModified           PostSuppliers200ApplicationJSONChangesTypeEnum = "Modified"
	PostSuppliers200ApplicationJSONChangesTypeEnumDeleted            PostSuppliers200ApplicationJSONChangesTypeEnum = "Deleted"
	PostSuppliers200ApplicationJSONChangesTypeEnumAttachmentUploaded PostSuppliers200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostSuppliers200ApplicationJSONChanges struct {
	AttachmentID *string                                                       `json:"attachmentId,omitempty"`
	RecordRef    *PostSuppliers200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostSuppliers200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostSuppliers200ApplicationJSONStatusEnum string

const (
	PostSuppliers200ApplicationJSONStatusEnumPending  PostSuppliers200ApplicationJSONStatusEnum = "Pending"
	PostSuppliers200ApplicationJSONStatusEnumFailed   PostSuppliers200ApplicationJSONStatusEnum = "Failed"
	PostSuppliers200ApplicationJSONStatusEnumSuccess  PostSuppliers200ApplicationJSONStatusEnum = "Success"
	PostSuppliers200ApplicationJSONStatusEnumTimedOut PostSuppliers200ApplicationJSONStatusEnum = "TimedOut"
)

type PostSuppliers200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostSuppliers200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostSuppliers200ApplicationJSONValidation struct {
	Errors   []PostSuppliers200ApplicationJSONValidationValidationItem                                                                                                                                           `json:"errors,omitempty"`
	Warnings []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1push1billsPostResponses200ContentApplication1jsonSchemaAllOf1PropertiesValidationPropertiesErrorsItems `json:"warnings,omitempty"`
}

type PostSuppliers200ApplicationJSON struct {
	Changes           []PostSuppliers200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                     `json:"companyId"`
	CompletedOnUtc    *time.Time                                 `json:"completedOnUtc,omitempty"`
	Data              *shared.Supplier                           `json:"data,omitempty"`
	DataConnectionKey string                                     `json:"dataConnectionKey"`
	DataType          *string                                    `json:"dataType,omitempty"`
	ErrorMessage      *string                                    `json:"errorMessage,omitempty"`
	PushOperationKey  string                                     `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                  `json:"requestedOnUtc"`
	Status            PostSuppliers200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                        `json:"statusCode"`
	TimeoutInMinutes  *int                                       `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                       `json:"timeoutInSeconds,omitempty"`
	Validation        *PostSuppliers200ApplicationJSONValidation `json:"validation,omitempty"`
}

type PostSuppliersResponse struct {
	ContentType                           string
	StatusCode                            int
	PostSuppliers200ApplicationJSONObject *PostSuppliers200ApplicationJSON
}
