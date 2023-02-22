package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
)

type ListCommerceProductsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListCommerceProductsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCommerceProductsRequest struct {
	PathParams  ListCommerceProductsPathParams
	QueryParams ListCommerceProductsQueryParams
}

type ListCommerceProductsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommerceProductsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceProductsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceProductsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommerceProductsLinksLinks struct {
	Current  ListCommerceProductsLinksLinksCurrent   `json:"current"`
	Next     *ListCommerceProductsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommerceProductsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommerceProductsLinksLinksSelf      `json:"self"`
}

// ListCommerceProductsLinks
// Codat's Paging Model
type ListCommerceProductsLinks struct {
	Links        ListCommerceProductsLinksLinks `json:"_links"`
	PageNumber   int64                          `json:"pageNumber"`
	PageSize     int64                          `json:"pageSize"`
	Results      []shared.Product               `json:"results,omitempty"`
	TotalResults int64                          `json:"totalResults"`
}

type ListCommerceProductsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListCommerceProductsLinks
}
