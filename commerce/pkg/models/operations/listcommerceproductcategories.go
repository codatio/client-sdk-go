package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
)

type ListCommerceProductCategoriesPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListCommerceProductCategoriesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCommerceProductCategoriesRequest struct {
	PathParams  ListCommerceProductCategoriesPathParams
	QueryParams ListCommerceProductCategoriesQueryParams
}

type ListCommerceProductCategoriesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommerceProductCategoriesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceProductCategoriesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceProductCategoriesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommerceProductCategoriesLinksLinks struct {
	Current  ListCommerceProductCategoriesLinksLinksCurrent   `json:"current"`
	Next     *ListCommerceProductCategoriesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommerceProductCategoriesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommerceProductCategoriesLinksLinksSelf      `json:"self"`
}

// ListCommerceProductCategoriesLinks
// Codat's Paging Model
type ListCommerceProductCategoriesLinks struct {
	Links        ListCommerceProductCategoriesLinksLinks `json:"_links"`
	PageNumber   int64                                   `json:"pageNumber"`
	PageSize     int64                                   `json:"pageSize"`
	Results      []shared.ProductCategory                `json:"results,omitempty"`
	TotalResults int64                                   `json:"totalResults"`
}

type ListCommerceProductCategoriesResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListCommerceProductCategoriesLinks
}
