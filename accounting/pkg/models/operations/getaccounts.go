package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetAccountsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetAccountsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type GetAccountsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetAccountsRequest struct {
	PathParams  GetAccountsPathParams
	QueryParams GetAccountsQueryParams
	Security    GetAccountsSecurity
}

type GetAccountsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type GetAccountsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type GetAccountsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type GetAccountsLinksLinksSelf struct {
	Href string `json:"href"`
}

type GetAccountsLinksLinks struct {
	Current  GetAccountsLinksLinksCurrent   `json:"current"`
	Next     *GetAccountsLinksLinksNext     `json:"next,omitempty"`
	Previous *GetAccountsLinksLinksPrevious `json:"previous,omitempty"`
	Self     GetAccountsLinksLinksSelf      `json:"self"`
}

// GetAccountsLinks
// Codat's Paging Model
type GetAccountsLinks struct {
	Links        GetAccountsLinksLinks `json:"_links"`
	PageNumber   int64                 `json:"pageNumber"`
	PageSize     int64                 `json:"pageSize"`
	Results      []shared.Account      `json:"results,omitempty"`
	TotalResults int64                 `json:"totalResults"`
}

type GetAccountsResponse struct {
	ContentType string
	StatusCode  int
	Links       *GetAccountsLinks
}
