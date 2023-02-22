package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
)

type ListCommercePaymentsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListCommercePaymentsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCommercePaymentsRequest struct {
	PathParams  ListCommercePaymentsPathParams
	QueryParams ListCommercePaymentsQueryParams
}

type ListCommercePaymentsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommercePaymentsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommercePaymentsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommercePaymentsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommercePaymentsLinksLinks struct {
	Current  ListCommercePaymentsLinksLinksCurrent   `json:"current"`
	Next     *ListCommercePaymentsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommercePaymentsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommercePaymentsLinksLinksSelf      `json:"self"`
}

// ListCommercePaymentsLinks
// Codat's Paging Model
type ListCommercePaymentsLinks struct {
	Links        ListCommercePaymentsLinksLinks `json:"_links"`
	PageNumber   int64                          `json:"pageNumber"`
	PageSize     int64                          `json:"pageSize"`
	Results      []shared.Payment               `json:"results,omitempty"`
	TotalResults int64                          `json:"totalResults"`
}

type ListCommercePaymentsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListCommercePaymentsLinks
}
