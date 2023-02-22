package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
)

type ListCommerceOrdersPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListCommerceOrdersQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCommerceOrdersRequest struct {
	PathParams  ListCommerceOrdersPathParams
	QueryParams ListCommerceOrdersQueryParams
}

type ListCommerceOrdersLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommerceOrdersLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceOrdersLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceOrdersLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommerceOrdersLinksLinks struct {
	Current  ListCommerceOrdersLinksLinksCurrent   `json:"current"`
	Next     *ListCommerceOrdersLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommerceOrdersLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommerceOrdersLinksLinksSelf      `json:"self"`
}

// ListCommerceOrdersLinks
// Codat's Paging Model
type ListCommerceOrdersLinks struct {
	Links        ListCommerceOrdersLinksLinks `json:"_links"`
	PageNumber   int64                        `json:"pageNumber"`
	PageSize     int64                        `json:"pageSize"`
	Results      []shared.Order               `json:"results,omitempty"`
	TotalResults int64                        `json:"totalResults"`
}

type ListCommerceOrdersResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListCommerceOrdersLinks
}
