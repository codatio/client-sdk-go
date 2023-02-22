package shared

import (
	"time"
)

type DataSetStatusEnum string

const (
	DataSetStatusEnumInitial            DataSetStatusEnum = "Initial"
	DataSetStatusEnumQueued             DataSetStatusEnum = "Queued"
	DataSetStatusEnumFetching           DataSetStatusEnum = "Fetching"
	DataSetStatusEnumMapQueued          DataSetStatusEnum = "MapQueued"
	DataSetStatusEnumMapping            DataSetStatusEnum = "Mapping"
	DataSetStatusEnumComplete           DataSetStatusEnum = "Complete"
	DataSetStatusEnumFetchError         DataSetStatusEnum = "FetchError"
	DataSetStatusEnumMapError           DataSetStatusEnum = "MapError"
	DataSetStatusEnumInternalError      DataSetStatusEnum = "InternalError"
	DataSetStatusEnumProcessingQueued   DataSetStatusEnum = "ProcessingQueued"
	DataSetStatusEnumProcessing         DataSetStatusEnum = "Processing"
	DataSetStatusEnumProcessingError    DataSetStatusEnum = "ProcessingError"
	DataSetStatusEnumValidationQueued   DataSetStatusEnum = "ValidationQueued"
	DataSetStatusEnumValidating         DataSetStatusEnum = "Validating"
	DataSetStatusEnumValidationError    DataSetStatusEnum = "ValidationError"
	DataSetStatusEnumAuthError          DataSetStatusEnum = "AuthError"
	DataSetStatusEnumCancelled          DataSetStatusEnum = "Cancelled"
	DataSetStatusEnumNotSupported       DataSetStatusEnum = "NotSupported"
	DataSetStatusEnumRateLimitError     DataSetStatusEnum = "RateLimitError"
	DataSetStatusEnumPermissionsError   DataSetStatusEnum = "PermissionsError"
	DataSetStatusEnumPrerequisiteNotMet DataSetStatusEnum = "PrerequisiteNotMet"
)

type DataSet struct {
	CompanyID                string            `json:"companyId"`
	Completed                *time.Time        `json:"completed,omitempty"`
	ConnectionID             string            `json:"connectionId"`
	DataType                 *string           `json:"dataType,omitempty"`
	DatasetLogsURL           *string           `json:"datasetLogsUrl,omitempty"`
	ErrorMessage             *string           `json:"errorMessage,omitempty"`
	ID                       string            `json:"id"`
	IsCompleted              bool              `json:"isCompleted"`
	IsErrored                bool              `json:"isErrored"`
	Progress                 int               `json:"progress"`
	Requested                time.Time         `json:"requested"`
	Status                   DataSetStatusEnum `json:"status"`
	ValidationinformationURL *string           `json:"validationinformationUrl,omitempty"`
}
