package shared

import (
	"time"
)

type PullOperationStatusEnum string

const (
	PullOperationStatusEnumInitial            PullOperationStatusEnum = "Initial"
	PullOperationStatusEnumQueued             PullOperationStatusEnum = "Queued"
	PullOperationStatusEnumFetching           PullOperationStatusEnum = "Fetching"
	PullOperationStatusEnumMapQueued          PullOperationStatusEnum = "MapQueued"
	PullOperationStatusEnumMapping            PullOperationStatusEnum = "Mapping"
	PullOperationStatusEnumComplete           PullOperationStatusEnum = "Complete"
	PullOperationStatusEnumFetchError         PullOperationStatusEnum = "FetchError"
	PullOperationStatusEnumMapError           PullOperationStatusEnum = "MapError"
	PullOperationStatusEnumInternalError      PullOperationStatusEnum = "InternalError"
	PullOperationStatusEnumProcessingQueued   PullOperationStatusEnum = "ProcessingQueued"
	PullOperationStatusEnumProcessing         PullOperationStatusEnum = "Processing"
	PullOperationStatusEnumProcessingError    PullOperationStatusEnum = "ProcessingError"
	PullOperationStatusEnumValidationQueued   PullOperationStatusEnum = "ValidationQueued"
	PullOperationStatusEnumValidating         PullOperationStatusEnum = "Validating"
	PullOperationStatusEnumValidationError    PullOperationStatusEnum = "ValidationError"
	PullOperationStatusEnumAuthError          PullOperationStatusEnum = "AuthError"
	PullOperationStatusEnumCancelled          PullOperationStatusEnum = "Cancelled"
	PullOperationStatusEnumRouting            PullOperationStatusEnum = "Routing"
	PullOperationStatusEnumRoutingError       PullOperationStatusEnum = "RoutingError"
	PullOperationStatusEnumNotSupported       PullOperationStatusEnum = "NotSupported"
	PullOperationStatusEnumRateLimitError     PullOperationStatusEnum = "RateLimitError"
	PullOperationStatusEnumPermissionsError   PullOperationStatusEnum = "PermissionsError"
	PullOperationStatusEnumPrerequisiteNotMet PullOperationStatusEnum = "PrerequisiteNotMet"
)

// PullOperation
// Information about a queued, in progress or completed pull operation.
// *Formally called `dataset`*
type PullOperation struct {
	CompanyID    string                  `json:"companyId"`
	ConnectionID string                  `json:"connectionId"`
	DataType     string                  `json:"dataType"`
	ID           string                  `json:"id"`
	IsCompleted  bool                    `json:"isCompleted"`
	IsErrored    bool                    `json:"isErrored"`
	Progress     int64                   `json:"progress"`
	Requested    time.Time               `json:"requested"`
	Status       PullOperationStatusEnum `json:"status"`
}
