package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetCustomersPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCustomersQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type GetCustomersSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetCustomersRequest struct {
	PathParams  GetCustomersPathParams
	QueryParams GetCustomersQueryParams
	Security    GetCustomersSecurity
}

type GetCustomersLinksLinksCurrent struct {
	Href string `json:"href"`
}

type GetCustomersLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type GetCustomersLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type GetCustomersLinksLinksSelf struct {
	Href string `json:"href"`
}

type GetCustomersLinksLinks struct {
	Current  GetCustomersLinksLinksCurrent   `json:"current"`
	Next     *GetCustomersLinksLinksNext     `json:"next,omitempty"`
	Previous *GetCustomersLinksLinksPrevious `json:"previous,omitempty"`
	Self     GetCustomersLinksLinksSelf      `json:"self"`
}

// GetCustomersLinks
// Codat's Paging Model
type GetCustomersLinks struct {
	Links        GetCustomersLinksLinks `json:"_links"`
	PageNumber   int64                  `json:"pageNumber"`
	PageSize     int64                  `json:"pageSize"`
	Results      []shared.Customer      `json:"results,omitempty"`
	TotalResults int64                  `json:"totalResults"`
}

type GetCustomersResponse struct {
	ContentType string
	StatusCode  int
	Links       *GetCustomersLinks
}
