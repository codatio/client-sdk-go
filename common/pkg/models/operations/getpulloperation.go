package operations

import (
	"net/http"
	"time"
)

type GetPullOperationPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	DatasetID string `pathParam:"style=simple,explode=false,name=datasetId"`
}

type GetPullOperationRequest struct {
	PathParams GetPullOperationPathParams
}

type GetPullOperation404ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetPullOperation401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetPullOperationPullOperationStatusEnum string

const (
	GetPullOperationPullOperationStatusEnumInitial            GetPullOperationPullOperationStatusEnum = "Initial"
	GetPullOperationPullOperationStatusEnumQueued             GetPullOperationPullOperationStatusEnum = "Queued"
	GetPullOperationPullOperationStatusEnumFetching           GetPullOperationPullOperationStatusEnum = "Fetching"
	GetPullOperationPullOperationStatusEnumMapQueued          GetPullOperationPullOperationStatusEnum = "MapQueued"
	GetPullOperationPullOperationStatusEnumMapping            GetPullOperationPullOperationStatusEnum = "Mapping"
	GetPullOperationPullOperationStatusEnumComplete           GetPullOperationPullOperationStatusEnum = "Complete"
	GetPullOperationPullOperationStatusEnumFetchError         GetPullOperationPullOperationStatusEnum = "FetchError"
	GetPullOperationPullOperationStatusEnumMapError           GetPullOperationPullOperationStatusEnum = "MapError"
	GetPullOperationPullOperationStatusEnumInternalError      GetPullOperationPullOperationStatusEnum = "InternalError"
	GetPullOperationPullOperationStatusEnumProcessingQueued   GetPullOperationPullOperationStatusEnum = "ProcessingQueued"
	GetPullOperationPullOperationStatusEnumProcessing         GetPullOperationPullOperationStatusEnum = "Processing"
	GetPullOperationPullOperationStatusEnumProcessingError    GetPullOperationPullOperationStatusEnum = "ProcessingError"
	GetPullOperationPullOperationStatusEnumValidationQueued   GetPullOperationPullOperationStatusEnum = "ValidationQueued"
	GetPullOperationPullOperationStatusEnumValidating         GetPullOperationPullOperationStatusEnum = "Validating"
	GetPullOperationPullOperationStatusEnumValidationError    GetPullOperationPullOperationStatusEnum = "ValidationError"
	GetPullOperationPullOperationStatusEnumAuthError          GetPullOperationPullOperationStatusEnum = "AuthError"
	GetPullOperationPullOperationStatusEnumCancelled          GetPullOperationPullOperationStatusEnum = "Cancelled"
	GetPullOperationPullOperationStatusEnumRouting            GetPullOperationPullOperationStatusEnum = "Routing"
	GetPullOperationPullOperationStatusEnumRoutingError       GetPullOperationPullOperationStatusEnum = "RoutingError"
	GetPullOperationPullOperationStatusEnumNotSupported       GetPullOperationPullOperationStatusEnum = "NotSupported"
	GetPullOperationPullOperationStatusEnumRateLimitError     GetPullOperationPullOperationStatusEnum = "RateLimitError"
	GetPullOperationPullOperationStatusEnumPermissionsError   GetPullOperationPullOperationStatusEnum = "PermissionsError"
	GetPullOperationPullOperationStatusEnumPrerequisiteNotMet GetPullOperationPullOperationStatusEnum = "PrerequisiteNotMet"
)

// GetPullOperationPullOperation
// Information about a queued, in progress or completed pull operation.
// *Formally called `dataset`*
type GetPullOperationPullOperation struct {
	CompanyID    string                                  `json:"companyId"`
	ConnectionID string                                  `json:"connectionId"`
	DataType     string                                  `json:"dataType"`
	ID           string                                  `json:"id"`
	IsCompleted  bool                                    `json:"isCompleted"`
	IsErrored    bool                                    `json:"isErrored"`
	Progress     int64                                   `json:"progress"`
	Requested    time.Time                               `json:"requested"`
	Status       GetPullOperationPullOperationStatusEnum `json:"status"`
}

type GetPullOperationResponse struct {
	ContentType                              string
	PullOperation                            *GetPullOperationPullOperation
	StatusCode                               int
	RawResponse                              *http.Response
	GetPullOperation401ApplicationJSONObject *GetPullOperation401ApplicationJSON
	GetPullOperation404ApplicationJSONObject *GetPullOperation404ApplicationJSON
}
