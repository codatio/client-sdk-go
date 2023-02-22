package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type ListIntegrationsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListIntegrationsRequest struct {
	QueryParams ListIntegrationsQueryParams
}

type ListIntegrations401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type ListIntegrations400ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type ListIntegrationsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListIntegrationsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListIntegrationsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListIntegrationsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListIntegrationsLinksLinks struct {
	Current  ListIntegrationsLinksLinksCurrent   `json:"current"`
	Next     *ListIntegrationsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListIntegrationsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListIntegrationsLinksLinksSelf      `json:"self"`
}

// ListIntegrationsLinks
// Codat's Paging Model
type ListIntegrationsLinks struct {
	Links        ListIntegrationsLinksLinks `json:"_links"`
	PageNumber   int64                      `json:"pageNumber"`
	PageSize     int64                      `json:"pageSize"`
	Results      []shared.Integration       `json:"results,omitempty"`
	TotalResults int64                      `json:"totalResults"`
}

type ListIntegrationsResponse struct {
	ContentType                              string
	StatusCode                               int
	Links                                    *ListIntegrationsLinks
	ListIntegrations400ApplicationJSONObject *ListIntegrations400ApplicationJSON
	ListIntegrations401ApplicationJSONObject *ListIntegrations401ApplicationJSON
}
