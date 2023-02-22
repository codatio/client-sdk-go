package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
)

type ListCommercePaymentMethodsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListCommercePaymentMethodsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCommercePaymentMethodsRequest struct {
	PathParams  ListCommercePaymentMethodsPathParams
	QueryParams ListCommercePaymentMethodsQueryParams
}

type ListCommercePaymentMethodsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommercePaymentMethodsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommercePaymentMethodsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommercePaymentMethodsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommercePaymentMethodsLinksLinks struct {
	Current  ListCommercePaymentMethodsLinksLinksCurrent   `json:"current"`
	Next     *ListCommercePaymentMethodsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommercePaymentMethodsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommercePaymentMethodsLinksLinksSelf      `json:"self"`
}

// ListCommercePaymentMethodsLinks
// Codat's Paging Model
type ListCommercePaymentMethodsLinks struct {
	Links        ListCommercePaymentMethodsLinksLinks `json:"_links"`
	PageNumber   int64                                `json:"pageNumber"`
	PageSize     int64                                `json:"pageSize"`
	Results      []shared.PaymentMethod               `json:"results,omitempty"`
	TotalResults int64                                `json:"totalResults"`
}

type ListCommercePaymentMethodsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListCommercePaymentMethodsLinks
}
