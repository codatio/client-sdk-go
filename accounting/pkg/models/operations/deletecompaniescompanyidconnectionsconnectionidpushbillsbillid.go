package operations

import (
	"net/http"
	"time"
)

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillIDRequest struct {
	BillID       string `pathParam:"style=simple,explode=false,name=billId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONChangesTypeEnum string

const (
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONChangesTypeEnumUnknown            DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONChangesTypeEnum = "Unknown"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONChangesTypeEnumCreated            DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONChangesTypeEnum = "Created"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONChangesTypeEnumModified           DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONChangesTypeEnum = "Modified"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONChangesTypeEnumDeleted            DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONChangesTypeEnum = "Deleted"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONChangesTypeEnumAttachmentUploaded DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONChanges struct {
	AttachmentID *string                                                                                                        `json:"attachmentId,omitempty"`
	RecordRef    *DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONStatusEnum string

const (
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONStatusEnumPending  DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONStatusEnum = "Pending"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONStatusEnumFailed   DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONStatusEnum = "Failed"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONStatusEnumSuccess  DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONStatusEnum = "Success"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONStatusEnumTimedOut DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONStatusEnum = "TimedOut"
)

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONValidation struct {
	Errors   []DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSON struct {
	Changes           []DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                                                                      `json:"companyId"`
	CompletedOnUtc    *time.Time                                                                                  `json:"completedOnUtc,omitempty"`
	DataConnectionKey string                                                                                      `json:"dataConnectionKey"`
	DataType          *string                                                                                     `json:"dataType,omitempty"`
	ErrorMessage      *string                                                                                     `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                                                      `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                                                                   `json:"requestedOnUtc"`
	Status            DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                                                                         `json:"statusCode"`
	TimeoutInMinutes  *int                                                                                        `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                                                        `json:"timeoutInSeconds,omitempty"`
	Validation        *DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONValidation `json:"validation,omitempty"`
}

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillIDResponse struct {
	Body                                                                                   []byte
	ContentType                                                                            string
	StatusCode                                                                             int
	RawResponse                                                                            *http.Response
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSONObject *DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID200ApplicationJSON
}
