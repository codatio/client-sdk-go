package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetDirectIncomesPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetDirectIncomesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type GetDirectIncomesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetDirectIncomesRequest struct {
	PathParams  GetDirectIncomesPathParams
	QueryParams GetDirectIncomesQueryParams
	Security    GetDirectIncomesSecurity
}

type GetDirectIncomesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type GetDirectIncomesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type GetDirectIncomesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type GetDirectIncomesLinksLinksSelf struct {
	Href string `json:"href"`
}

type GetDirectIncomesLinksLinks struct {
	Current  GetDirectIncomesLinksLinksCurrent   `json:"current"`
	Next     *GetDirectIncomesLinksLinksNext     `json:"next,omitempty"`
	Previous *GetDirectIncomesLinksLinksPrevious `json:"previous,omitempty"`
	Self     GetDirectIncomesLinksLinksSelf      `json:"self"`
}

// GetDirectIncomesLinks
// Codat's Paging Model
type GetDirectIncomesLinks struct {
	Links        GetDirectIncomesLinksLinks `json:"_links"`
	PageNumber   int64                      `json:"pageNumber"`
	PageSize     int64                      `json:"pageSize"`
	Results      []shared.DirectIncome      `json:"results,omitempty"`
	TotalResults int64                      `json:"totalResults"`
}

type GetDirectIncomesResponse struct {
	ContentType string
	StatusCode  int
	Links       *GetDirectIncomesLinks
}
