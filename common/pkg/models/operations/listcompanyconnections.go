package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
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

// ListCompanyConnectionsLinks
// Codat's Paging Model
type ListCompanyConnectionsLinks struct {
	Links        ListCompanyConnectionsLinksLinks `json:"_links"`
	PageNumber   int64                            `json:"pageNumber"`
	PageSize     int64                            `json:"pageSize"`
	Results      []shared.Connection              `json:"results,omitempty"`
	TotalResults int64                            `json:"totalResults"`
}

type ListCompanyConnectionsResponse struct {
	ContentType                                    string
	StatusCode                                     int
	Links                                          *ListCompanyConnectionsLinks
	ListCompanyConnections400ApplicationJSONObject *ListCompanyConnections400ApplicationJSON
	ListCompanyConnections401ApplicationJSONObject *ListCompanyConnections401ApplicationJSON
}
