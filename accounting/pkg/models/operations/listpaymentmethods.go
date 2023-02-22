package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListPaymentMethodsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListPaymentMethodsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListPaymentMethodsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListPaymentMethodsRequest struct {
	PathParams  ListPaymentMethodsPathParams
	QueryParams ListPaymentMethodsQueryParams
	Security    ListPaymentMethodsSecurity
}

type ListPaymentMethodsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListPaymentMethodsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListPaymentMethodsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListPaymentMethodsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListPaymentMethodsLinksLinks struct {
	Current  ListPaymentMethodsLinksLinksCurrent   `json:"current"`
	Next     *ListPaymentMethodsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListPaymentMethodsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListPaymentMethodsLinksLinksSelf      `json:"self"`
}

// ListPaymentMethodsLinks
// Codat's Paging Model
type ListPaymentMethodsLinks struct {
	Links        ListPaymentMethodsLinksLinks `json:"_links"`
	PageNumber   int64                        `json:"pageNumber"`
	PageSize     int64                        `json:"pageSize"`
	Results      []shared.PaymentMethod       `json:"results,omitempty"`
	TotalResults int64                        `json:"totalResults"`
}

type ListPaymentMethodsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListPaymentMethodsLinks
}
