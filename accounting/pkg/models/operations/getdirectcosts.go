package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetDirectCostsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetDirectCostsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type GetDirectCostsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetDirectCostsRequest struct {
	PathParams  GetDirectCostsPathParams
	QueryParams GetDirectCostsQueryParams
	Security    GetDirectCostsSecurity
}

type GetDirectCostsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type GetDirectCostsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type GetDirectCostsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type GetDirectCostsLinksLinksSelf struct {
	Href string `json:"href"`
}

type GetDirectCostsLinksLinks struct {
	Current  GetDirectCostsLinksLinksCurrent   `json:"current"`
	Next     *GetDirectCostsLinksLinksNext     `json:"next,omitempty"`
	Previous *GetDirectCostsLinksLinksPrevious `json:"previous,omitempty"`
	Self     GetDirectCostsLinksLinksSelf      `json:"self"`
}

// GetDirectCostsLinks
// Codat's Paging Model
type GetDirectCostsLinks struct {
	Links        GetDirectCostsLinksLinks `json:"_links"`
	PageNumber   int64                    `json:"pageNumber"`
	PageSize     int64                    `json:"pageSize"`
	Results      []shared.DirectCost      `json:"results,omitempty"`
	TotalResults int64                    `json:"totalResults"`
}

type GetDirectCostsResponse struct {
	ContentType string
	StatusCode  int
	Links       *GetDirectCostsLinks
}
