package operations

import (
	"net/http"
	"time"
)

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentIDPathParams struct {
	BillPaymentID string `pathParam:"style=simple,explode=false,name=billPaymentId"`
	CompanyID     string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID  string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentIDRequest struct {
	PathParams DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentIDPathParams
}

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum string

const (
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnumUnknown            DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum = "Unknown"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnumCreated            DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum = "Created"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnumModified           DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum = "Modified"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnumDeleted            DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum = "Deleted"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnumAttachmentUploaded DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChanges struct {
	AttachmentID *string                                                                                                                      `json:"attachmentId,omitempty"`
	RecordRef    *DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnum string

const (
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnumPending  DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnum = "Pending"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnumFailed   DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnum = "Failed"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnumSuccess  DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnum = "Success"
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnumTimedOut DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnum = "TimedOut"
)

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONValidation struct {
	Errors   []DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSON struct {
	Changes           []DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONChanges   `json:"changes,omitempty"`
	CompanyID         string                                                                                                    `json:"companyId"`
	CompletedOnUtc    *time.Time                                                                                                `json:"completedOnUtc,omitempty"`
	DataConnectionKey string                                                                                                    `json:"dataConnectionKey"`
	DataType          *string                                                                                                   `json:"dataType,omitempty"`
	ErrorMessage      *string                                                                                                   `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                                                                    `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                                                                                 `json:"requestedOnUtc"`
	Status            DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONStatusEnum  `json:"status"`
	StatusCode        int                                                                                                       `json:"statusCode"`
	TimeoutInMinutes  *int                                                                                                      `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                                                                      `json:"timeoutInSeconds,omitempty"`
	Validation        *DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONValidation `json:"validation,omitempty"`
}

type DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentIDResponse struct {
	ContentType                                                                                          string
	StatusCode                                                                                           int
	RawResponse                                                                                          *http.Response
	DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSONObject *DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID200ApplicationJSON
}
