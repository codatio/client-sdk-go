package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListSuppliersPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListSuppliersQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListSuppliersSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListSuppliersRequest struct {
	PathParams  ListSuppliersPathParams
	QueryParams ListSuppliersQueryParams
	Security    ListSuppliersSecurity
}

type ListSuppliersLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListSuppliersLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListSuppliersLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListSuppliersLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListSuppliersLinksLinks struct {
	Current  ListSuppliersLinksLinksCurrent   `json:"current"`
	Next     *ListSuppliersLinksLinksNext     `json:"next,omitempty"`
	Previous *ListSuppliersLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListSuppliersLinksLinksSelf      `json:"self"`
}

// ListSuppliersLinks
// Codat's Paging Model
type ListSuppliersLinks struct {
	Links        ListSuppliersLinksLinks `json:"_links"`
	PageNumber   int64                   `json:"pageNumber"`
	PageSize     int64                   `json:"pageSize"`
	Results      []shared.Supplier       `json:"results,omitempty"`
	TotalResults int64                   `json:"totalResults"`
}

type ListSuppliersResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListSuppliersLinks
}
