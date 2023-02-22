package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
)

type ListCommerceTransactionsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListCommerceTransactionsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCommerceTransactionsRequest struct {
	PathParams  ListCommerceTransactionsPathParams
	QueryParams ListCommerceTransactionsQueryParams
}

type ListCommerceTransactionsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommerceTransactionsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceTransactionsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceTransactionsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommerceTransactionsLinksLinks struct {
	Current  ListCommerceTransactionsLinksLinksCurrent   `json:"current"`
	Next     *ListCommerceTransactionsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommerceTransactionsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommerceTransactionsLinksLinksSelf      `json:"self"`
}

// ListCommerceTransactionsLinks
// Codat's Paging Model
type ListCommerceTransactionsLinks struct {
	Links        ListCommerceTransactionsLinksLinks `json:"_links"`
	PageNumber   int64                              `json:"pageNumber"`
	PageSize     int64                              `json:"pageSize"`
	Results      []shared.Transaction               `json:"results,omitempty"`
	TotalResults int64                              `json:"totalResults"`
}

type ListCommerceTransactionsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListCommerceTransactionsLinks
}
