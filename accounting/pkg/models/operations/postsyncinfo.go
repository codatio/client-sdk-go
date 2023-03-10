package operations

import (
	"net/http"
	"time"
)

type PostSyncInfoPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type PostSyncInfoRequest struct {
	PathParams PostSyncInfoPathParams
}

type PostSyncInfo200ApplicationJSONStatusEnum string

const (
	PostSyncInfo200ApplicationJSONStatusEnumInitial            PostSyncInfo200ApplicationJSONStatusEnum = "Initial"
	PostSyncInfo200ApplicationJSONStatusEnumQueued             PostSyncInfo200ApplicationJSONStatusEnum = "Queued"
	PostSyncInfo200ApplicationJSONStatusEnumFetching           PostSyncInfo200ApplicationJSONStatusEnum = "Fetching"
	PostSyncInfo200ApplicationJSONStatusEnumMapQueued          PostSyncInfo200ApplicationJSONStatusEnum = "MapQueued"
	PostSyncInfo200ApplicationJSONStatusEnumMapping            PostSyncInfo200ApplicationJSONStatusEnum = "Mapping"
	PostSyncInfo200ApplicationJSONStatusEnumComplete           PostSyncInfo200ApplicationJSONStatusEnum = "Complete"
	PostSyncInfo200ApplicationJSONStatusEnumFetchError         PostSyncInfo200ApplicationJSONStatusEnum = "FetchError"
	PostSyncInfo200ApplicationJSONStatusEnumMapError           PostSyncInfo200ApplicationJSONStatusEnum = "MapError"
	PostSyncInfo200ApplicationJSONStatusEnumInternalError      PostSyncInfo200ApplicationJSONStatusEnum = "InternalError"
	PostSyncInfo200ApplicationJSONStatusEnumProcessingQueued   PostSyncInfo200ApplicationJSONStatusEnum = "ProcessingQueued"
	PostSyncInfo200ApplicationJSONStatusEnumProcessing         PostSyncInfo200ApplicationJSONStatusEnum = "Processing"
	PostSyncInfo200ApplicationJSONStatusEnumProcessingError    PostSyncInfo200ApplicationJSONStatusEnum = "ProcessingError"
	PostSyncInfo200ApplicationJSONStatusEnumValidationQueued   PostSyncInfo200ApplicationJSONStatusEnum = "ValidationQueued"
	PostSyncInfo200ApplicationJSONStatusEnumValidating         PostSyncInfo200ApplicationJSONStatusEnum = "Validating"
	PostSyncInfo200ApplicationJSONStatusEnumValidationError    PostSyncInfo200ApplicationJSONStatusEnum = "ValidationError"
	PostSyncInfo200ApplicationJSONStatusEnumAuthError          PostSyncInfo200ApplicationJSONStatusEnum = "AuthError"
	PostSyncInfo200ApplicationJSONStatusEnumCancelled          PostSyncInfo200ApplicationJSONStatusEnum = "Cancelled"
	PostSyncInfo200ApplicationJSONStatusEnumNotSupported       PostSyncInfo200ApplicationJSONStatusEnum = "NotSupported"
	PostSyncInfo200ApplicationJSONStatusEnumRateLimitError     PostSyncInfo200ApplicationJSONStatusEnum = "RateLimitError"
	PostSyncInfo200ApplicationJSONStatusEnumPermissionsError   PostSyncInfo200ApplicationJSONStatusEnum = "PermissionsError"
	PostSyncInfo200ApplicationJSONStatusEnumPrerequisiteNotMet PostSyncInfo200ApplicationJSONStatusEnum = "PrerequisiteNotMet"
)

type PostSyncInfo200ApplicationJSON struct {
	CompanyID                string                                   `json:"companyId"`
	Completed                *time.Time                               `json:"completed,omitempty"`
	ConnectionID             string                                   `json:"connectionId"`
	DataType                 *string                                  `json:"dataType,omitempty"`
	DatasetLogsURL           *string                                  `json:"datasetLogsUrl,omitempty"`
	ErrorMessage             *string                                  `json:"errorMessage,omitempty"`
	ID                       string                                   `json:"id"`
	IsCompleted              bool                                     `json:"isCompleted"`
	IsErrored                bool                                     `json:"isErrored"`
	Progress                 int                                      `json:"progress"`
	Requested                time.Time                                `json:"requested"`
	Status                   PostSyncInfo200ApplicationJSONStatusEnum `json:"status"`
	ValidationinformationURL *string                                  `json:"validationinformationUrl,omitempty"`
}

type PostSyncInfoResponse struct {
	ContentType                          string
	StatusCode                           int
	RawResponse                          *http.Response
	PostSyncInfo200ApplicationJSONObject *PostSyncInfo200ApplicationJSON
}
