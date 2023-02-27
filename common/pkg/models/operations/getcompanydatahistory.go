package operations

import (
	"time"
)

type GetCompanyDataHistoryPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompanyDataHistoryQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type GetCompanyDataHistoryRequest struct {
	PathParams  GetCompanyDataHistoryPathParams
	QueryParams GetCompanyDataHistoryQueryParams
}

type GetCompanyDataHistory404ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetCompanyDataHistory401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetCompanyDataHistory400ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetCompanyDataHistoryLinksLinksCurrent struct {
	Href string `json:"href"`
}

type GetCompanyDataHistoryLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type GetCompanyDataHistoryLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type GetCompanyDataHistoryLinksLinksSelf struct {
	Href string `json:"href"`
}

type GetCompanyDataHistoryLinksLinks struct {
	Current  GetCompanyDataHistoryLinksLinksCurrent   `json:"current"`
	Next     *GetCompanyDataHistoryLinksLinksNext     `json:"next,omitempty"`
	Previous *GetCompanyDataHistoryLinksLinksPrevious `json:"previous,omitempty"`
	Self     GetCompanyDataHistoryLinksLinksSelf      `json:"self"`
}

type GetCompanyDataHistoryLinksPullOperationStatusEnum string

const (
	GetCompanyDataHistoryLinksPullOperationStatusEnumInitial            GetCompanyDataHistoryLinksPullOperationStatusEnum = "Initial"
	GetCompanyDataHistoryLinksPullOperationStatusEnumQueued             GetCompanyDataHistoryLinksPullOperationStatusEnum = "Queued"
	GetCompanyDataHistoryLinksPullOperationStatusEnumFetching           GetCompanyDataHistoryLinksPullOperationStatusEnum = "Fetching"
	GetCompanyDataHistoryLinksPullOperationStatusEnumMapQueued          GetCompanyDataHistoryLinksPullOperationStatusEnum = "MapQueued"
	GetCompanyDataHistoryLinksPullOperationStatusEnumMapping            GetCompanyDataHistoryLinksPullOperationStatusEnum = "Mapping"
	GetCompanyDataHistoryLinksPullOperationStatusEnumComplete           GetCompanyDataHistoryLinksPullOperationStatusEnum = "Complete"
	GetCompanyDataHistoryLinksPullOperationStatusEnumFetchError         GetCompanyDataHistoryLinksPullOperationStatusEnum = "FetchError"
	GetCompanyDataHistoryLinksPullOperationStatusEnumMapError           GetCompanyDataHistoryLinksPullOperationStatusEnum = "MapError"
	GetCompanyDataHistoryLinksPullOperationStatusEnumInternalError      GetCompanyDataHistoryLinksPullOperationStatusEnum = "InternalError"
	GetCompanyDataHistoryLinksPullOperationStatusEnumProcessingQueued   GetCompanyDataHistoryLinksPullOperationStatusEnum = "ProcessingQueued"
	GetCompanyDataHistoryLinksPullOperationStatusEnumProcessing         GetCompanyDataHistoryLinksPullOperationStatusEnum = "Processing"
	GetCompanyDataHistoryLinksPullOperationStatusEnumProcessingError    GetCompanyDataHistoryLinksPullOperationStatusEnum = "ProcessingError"
	GetCompanyDataHistoryLinksPullOperationStatusEnumValidationQueued   GetCompanyDataHistoryLinksPullOperationStatusEnum = "ValidationQueued"
	GetCompanyDataHistoryLinksPullOperationStatusEnumValidating         GetCompanyDataHistoryLinksPullOperationStatusEnum = "Validating"
	GetCompanyDataHistoryLinksPullOperationStatusEnumValidationError    GetCompanyDataHistoryLinksPullOperationStatusEnum = "ValidationError"
	GetCompanyDataHistoryLinksPullOperationStatusEnumAuthError          GetCompanyDataHistoryLinksPullOperationStatusEnum = "AuthError"
	GetCompanyDataHistoryLinksPullOperationStatusEnumCancelled          GetCompanyDataHistoryLinksPullOperationStatusEnum = "Cancelled"
	GetCompanyDataHistoryLinksPullOperationStatusEnumRouting            GetCompanyDataHistoryLinksPullOperationStatusEnum = "Routing"
	GetCompanyDataHistoryLinksPullOperationStatusEnumRoutingError       GetCompanyDataHistoryLinksPullOperationStatusEnum = "RoutingError"
	GetCompanyDataHistoryLinksPullOperationStatusEnumNotSupported       GetCompanyDataHistoryLinksPullOperationStatusEnum = "NotSupported"
	GetCompanyDataHistoryLinksPullOperationStatusEnumRateLimitError     GetCompanyDataHistoryLinksPullOperationStatusEnum = "RateLimitError"
	GetCompanyDataHistoryLinksPullOperationStatusEnumPermissionsError   GetCompanyDataHistoryLinksPullOperationStatusEnum = "PermissionsError"
	GetCompanyDataHistoryLinksPullOperationStatusEnumPrerequisiteNotMet GetCompanyDataHistoryLinksPullOperationStatusEnum = "PrerequisiteNotMet"
)

// GetCompanyDataHistoryLinksPullOperation
// Information about a queued, in progress or completed pull operation.
// *Formally called `dataset`*
type GetCompanyDataHistoryLinksPullOperation struct {
	CompanyID    string                                            `json:"companyId"`
	ConnectionID string                                            `json:"connectionId"`
	DataType     string                                            `json:"dataType"`
	ID           string                                            `json:"id"`
	IsCompleted  bool                                              `json:"isCompleted"`
	IsErrored    bool                                              `json:"isErrored"`
	Progress     int64                                             `json:"progress"`
	Requested    time.Time                                         `json:"requested"`
	Status       GetCompanyDataHistoryLinksPullOperationStatusEnum `json:"status"`
}

// GetCompanyDataHistoryLinks
// Codat's Paging Model
type GetCompanyDataHistoryLinks struct {
	Links        GetCompanyDataHistoryLinksLinks           `json:"_links"`
	PageNumber   int64                                     `json:"pageNumber"`
	PageSize     int64                                     `json:"pageSize"`
	Results      []GetCompanyDataHistoryLinksPullOperation `json:"results,omitempty"`
	TotalResults int64                                     `json:"totalResults"`
}

type GetCompanyDataHistoryResponse struct {
	ContentType                                   string
	StatusCode                                    int
	Links                                         *GetCompanyDataHistoryLinks
	GetCompanyDataHistory400ApplicationJSONObject *GetCompanyDataHistory400ApplicationJSON
	GetCompanyDataHistory401ApplicationJSONObject *GetCompanyDataHistory401ApplicationJSON
	GetCompanyDataHistory404ApplicationJSONObject *GetCompanyDataHistory404ApplicationJSON
}
