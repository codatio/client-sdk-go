package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListSalesOrdersPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListSalesOrdersQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListSalesOrdersSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListSalesOrdersRequest struct {
	PathParams  ListSalesOrdersPathParams
	QueryParams ListSalesOrdersQueryParams
	Security    ListSalesOrdersSecurity
}

type ListSalesOrdersLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListSalesOrdersLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListSalesOrdersLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListSalesOrdersLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListSalesOrdersLinksLinks struct {
	Current  ListSalesOrdersLinksLinksCurrent   `json:"current"`
	Next     *ListSalesOrdersLinksLinksNext     `json:"next,omitempty"`
	Previous *ListSalesOrdersLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListSalesOrdersLinksLinksSelf      `json:"self"`
}

// ListSalesOrdersLinks
// Codat's Paging Model
type ListSalesOrdersLinks struct {
	Links        ListSalesOrdersLinksLinks `json:"_links"`
	PageNumber   int64                     `json:"pageNumber"`
	PageSize     int64                     `json:"pageSize"`
	Results      []shared.SalesOrder       `json:"results,omitempty"`
	TotalResults int64                     `json:"totalResults"`
}

type ListSalesOrdersResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListSalesOrdersLinks
}
