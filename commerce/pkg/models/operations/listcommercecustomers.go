package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
)

type ListCommerceCustomersPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListCommerceCustomersQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCommerceCustomersRequest struct {
	PathParams  ListCommerceCustomersPathParams
	QueryParams ListCommerceCustomersQueryParams
}

type ListCommerceCustomersLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommerceCustomersLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceCustomersLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceCustomersLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommerceCustomersLinksLinks struct {
	Current  ListCommerceCustomersLinksLinksCurrent   `json:"current"`
	Next     *ListCommerceCustomersLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommerceCustomersLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommerceCustomersLinksLinksSelf      `json:"self"`
}

// ListCommerceCustomersLinks
// Codat's Paging Model
type ListCommerceCustomersLinks struct {
	Links        ListCommerceCustomersLinksLinks `json:"_links"`
	PageNumber   int64                           `json:"pageNumber"`
	PageSize     int64                           `json:"pageSize"`
	Results      []shared.Customer               `json:"results,omitempty"`
	TotalResults int64                           `json:"totalResults"`
}

type ListCommerceCustomersResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListCommerceCustomersLinks
}
