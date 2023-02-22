package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListInvoicesPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListInvoicesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListInvoicesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListInvoicesRequest struct {
	PathParams  ListInvoicesPathParams
	QueryParams ListInvoicesQueryParams
	Security    ListInvoicesSecurity
}

type ListInvoicesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListInvoicesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListInvoicesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListInvoicesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListInvoicesLinksLinks struct {
	Current  ListInvoicesLinksLinksCurrent   `json:"current"`
	Next     *ListInvoicesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListInvoicesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListInvoicesLinksLinksSelf      `json:"self"`
}

// ListInvoicesLinks
// Codat's Paging Model
type ListInvoicesLinks struct {
	Links        ListInvoicesLinksLinks `json:"_links"`
	PageNumber   int64                  `json:"pageNumber"`
	PageSize     int64                  `json:"pageSize"`
	Results      []shared.Invoice       `json:"results,omitempty"`
	TotalResults int64                  `json:"totalResults"`
}

type ListInvoicesResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListInvoicesLinks
}
