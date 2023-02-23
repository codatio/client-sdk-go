package operations

import (
	"time"
)

type ListCompanyConnectionsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListCompanyConnectionsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCompanyConnectionsRequest struct {
	PathParams  ListCompanyConnectionsPathParams
	QueryParams ListCompanyConnectionsQueryParams
}

type ListCompanyConnections401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type ListCompanyConnections400ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type ListCompanyConnectionsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCompanyConnectionsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCompanyConnectionsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCompanyConnectionsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCompanyConnectionsLinksLinks struct {
	Current  ListCompanyConnectionsLinksLinksCurrent   `json:"current"`
	Next     *ListCompanyConnectionsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCompanyConnectionsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCompanyConnectionsLinksLinksSelf      `json:"self"`
}

type ListCompanyConnectionsLinksConnectionConnectionInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

type ListCompanyConnectionsLinksConnectionDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type ListCompanyConnectionsLinksConnectionSourceTypeEnum string

const (
	ListCompanyConnectionsLinksConnectionSourceTypeEnumAccounting ListCompanyConnectionsLinksConnectionSourceTypeEnum = "Accounting"
	ListCompanyConnectionsLinksConnectionSourceTypeEnumBanking    ListCompanyConnectionsLinksConnectionSourceTypeEnum = "Banking"
	ListCompanyConnectionsLinksConnectionSourceTypeEnumCommerce   ListCompanyConnectionsLinksConnectionSourceTypeEnum = "Commerce"
	ListCompanyConnectionsLinksConnectionSourceTypeEnumOther      ListCompanyConnectionsLinksConnectionSourceTypeEnum = "Other"
	ListCompanyConnectionsLinksConnectionSourceTypeEnumUnknown    ListCompanyConnectionsLinksConnectionSourceTypeEnum = "Unknown"
)

type ListCompanyConnectionsLinksConnectionStatusEnum string

const (
	ListCompanyConnectionsLinksConnectionStatusEnumPendingAuth  ListCompanyConnectionsLinksConnectionStatusEnum = "PendingAuth"
	ListCompanyConnectionsLinksConnectionStatusEnumLinked       ListCompanyConnectionsLinksConnectionStatusEnum = "Linked"
	ListCompanyConnectionsLinksConnectionStatusEnumUnlinked     ListCompanyConnectionsLinksConnectionStatusEnum = "Unlinked"
	ListCompanyConnectionsLinksConnectionStatusEnumDeauthorized ListCompanyConnectionsLinksConnectionStatusEnum = "Deauthorized"
)

// ListCompanyConnectionsLinksConnection
// A connection represents the link between a `company` and a source of data.
type ListCompanyConnectionsLinksConnection struct {
	ConnectionInfo       *ListCompanyConnectionsLinksConnectionConnectionInfo        `json:"connectionInfo,omitempty"`
	Created              time.Time                                                   `json:"created"`
	DataConnectionErrors []ListCompanyConnectionsLinksConnectionDataConnectionErrors `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                      `json:"id"`
	IntegrationID        string                                                      `json:"integrationId"`
	IntegrationKey       string                                                      `json:"integrationKey"`
	LastSync             *time.Time                                                  `json:"lastSync,omitempty"`
	LinkURL              string                                                      `json:"linkUrl"`
	PlatformName         string                                                      `json:"platformName"`
	SourceID             string                                                      `json:"sourceId"`
	SourceType           ListCompanyConnectionsLinksConnectionSourceTypeEnum         `json:"sourceType"`
	Status               ListCompanyConnectionsLinksConnectionStatusEnum             `json:"status"`
}

// ListCompanyConnectionsLinks
// Codat's Paging Model
type ListCompanyConnectionsLinks struct {
	Links        ListCompanyConnectionsLinksLinks        `json:"_links"`
	PageNumber   int64                                   `json:"pageNumber"`
	PageSize     int64                                   `json:"pageSize"`
	Results      []ListCompanyConnectionsLinksConnection `json:"results,omitempty"`
	TotalResults int64                                   `json:"totalResults"`
}

type ListCompanyConnectionsResponse struct {
	ContentType                                    string
	StatusCode                                     int
	Links                                          *ListCompanyConnectionsLinks
	ListCompanyConnections400ApplicationJSONObject *ListCompanyConnections400ApplicationJSON
	ListCompanyConnections401ApplicationJSONObject *ListCompanyConnections401ApplicationJSON
}
