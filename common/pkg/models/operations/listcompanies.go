package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type ListCompaniesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCompaniesRequest struct {
	QueryParams ListCompaniesQueryParams
}

type ListCompanies401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type ListCompanies400ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type ListCompaniesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCompaniesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCompaniesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCompaniesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCompaniesLinksLinks struct {
	Current  ListCompaniesLinksLinksCurrent   `json:"current"`
	Next     *ListCompaniesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCompaniesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCompaniesLinksLinksSelf      `json:"self"`
}

// ListCompaniesLinks
// Codat's Paging Model
type ListCompaniesLinks struct {
	Links        ListCompaniesLinksLinks `json:"_links"`
	PageNumber   int64                   `json:"pageNumber"`
	PageSize     int64                   `json:"pageSize"`
	Results      []shared.Company        `json:"results,omitempty"`
	TotalResults int64                   `json:"totalResults"`
}

type ListCompaniesResponse struct {
	ContentType                           string
	StatusCode                            int
	Links                                 *ListCompaniesLinks
	ListCompanies400ApplicationJSONObject *ListCompanies400ApplicationJSON
	ListCompanies401ApplicationJSONObject *ListCompanies401ApplicationJSON
}
