package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListBillsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListBillsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListBillsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListBillsRequest struct {
	PathParams  ListBillsPathParams
	QueryParams ListBillsQueryParams
	Security    ListBillsSecurity
}

type ListBillsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListBillsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListBillsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListBillsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListBillsLinksLinks struct {
	Current  ListBillsLinksLinksCurrent   `json:"current"`
	Next     *ListBillsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListBillsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListBillsLinksLinksSelf      `json:"self"`
}

// ListBillsLinks
// Codat's Paging Model
type ListBillsLinks struct {
	Links        ListBillsLinksLinks `json:"_links"`
	PageNumber   int64               `json:"pageNumber"`
	PageSize     int64               `json:"pageSize"`
	Results      []shared.Bill       `json:"results,omitempty"`
	TotalResults int64               `json:"totalResults"`
}

type ListBillsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListBillsLinks
}
