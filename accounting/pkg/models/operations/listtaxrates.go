package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListTaxRatesPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListTaxRatesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListTaxRatesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListTaxRatesRequest struct {
	PathParams  ListTaxRatesPathParams
	QueryParams ListTaxRatesQueryParams
	Security    ListTaxRatesSecurity
}

type ListTaxRatesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListTaxRatesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListTaxRatesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListTaxRatesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListTaxRatesLinksLinks struct {
	Current  ListTaxRatesLinksLinksCurrent   `json:"current"`
	Next     *ListTaxRatesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListTaxRatesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListTaxRatesLinksLinksSelf      `json:"self"`
}

// ListTaxRatesLinks
// Codat's Paging Model
type ListTaxRatesLinks struct {
	Links        ListTaxRatesLinksLinks `json:"_links"`
	PageNumber   int64                  `json:"pageNumber"`
	PageSize     int64                  `json:"pageSize"`
	Results      []shared.TaxRate       `json:"results,omitempty"`
	TotalResults int64                  `json:"totalResults"`
}

type ListTaxRatesResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListTaxRatesLinks
}
