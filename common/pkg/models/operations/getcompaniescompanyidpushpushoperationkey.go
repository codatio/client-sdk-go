package operations

import (
	"net/http"
	"time"
)

type GetCompaniesCompanyIDPushPushOperationKeyPathParams struct {
	CompanyID        string `pathParam:"style=simple,explode=false,name=companyId"`
	PushOperationKey string `pathParam:"style=simple,explode=false,name=pushOperationKey"`
}

type GetCompaniesCompanyIDPushPushOperationKeyRequest struct {
	PathParams GetCompaniesCompanyIDPushPushOperationKeyPathParams
}

type GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONChangesTypeEnum string

const (
	GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONChangesTypeEnumUnknown            GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONChangesTypeEnum = "Unknown"
	GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONChangesTypeEnumCreated            GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONChangesTypeEnum = "Created"
	GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONChangesTypeEnumModified           GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONChangesTypeEnum = "Modified"
	GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONChangesTypeEnumDeleted            GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONChangesTypeEnum = "Deleted"
	GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONChangesTypeEnumAttachmentUploaded GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONChanges struct {
	AttachmentID *string                                                                                   `json:"attachmentId,omitempty"`
	RecordRef    *GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONStatusEnum string

const (
	GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONStatusEnumPending  GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONStatusEnum = "Pending"
	GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONStatusEnumFailed   GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONStatusEnum = "Failed"
	GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONStatusEnumSuccess  GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONStatusEnum = "Success"
	GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONStatusEnumTimedOut GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONStatusEnum = "TimedOut"
)

type GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONValidation struct {
	Errors   []GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSON struct {
	Changes           []GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                                                 `json:"companyId"`
	CompletedOnUtc    *time.Time                                                             `json:"completedOnUtc,omitempty"`
	Data              map[string]interface{}                                                 `json:"data,omitempty"`
	DataConnectionKey string                                                                 `json:"dataConnectionKey"`
	DataType          *string                                                                `json:"dataType,omitempty"`
	ErrorMessage      *string                                                                `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                                 `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                                              `json:"requestedOnUtc"`
	Status            GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                                                    `json:"statusCode"`
	TimeoutInMinutes  *int                                                                   `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                                   `json:"timeoutInSeconds,omitempty"`
	Validation        *GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONValidation `json:"validation,omitempty"`
}

type GetCompaniesCompanyIDPushPushOperationKeyResponse struct {
	ContentType                                                       string
	StatusCode                                                        int
	RawResponse                                                       *http.Response
	GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSONObject *GetCompaniesCompanyIDPushPushOperationKey200ApplicationJSON
}
