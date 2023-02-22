package operations

import (
	"github.com/codatio/client-sdk-go/banking/pkg/models/shared"
)

type ListBankingAccountsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListBankingAccountsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListBankingAccountsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListBankingAccountsRequest struct {
	PathParams  ListBankingAccountsPathParams
	QueryParams ListBankingAccountsQueryParams
	Security    ListBankingAccountsSecurity
}

type ListBankingAccountsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListBankingAccountsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListBankingAccountsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListBankingAccountsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListBankingAccountsLinksLinks struct {
	Current  ListBankingAccountsLinksLinksCurrent   `json:"current"`
	Next     *ListBankingAccountsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListBankingAccountsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListBankingAccountsLinksLinksSelf      `json:"self"`
}

// ListBankingAccountsLinks
// Codat's Paging Model
type ListBankingAccountsLinks struct {
	Links        ListBankingAccountsLinksLinks `json:"_links"`
	PageNumber   int64                         `json:"pageNumber"`
	PageSize     int64                         `json:"pageSize"`
	Results      *shared.Account               `json:"results,omitempty"`
	TotalResults int64                         `json:"totalResults"`
}

type ListBankingAccountsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListBankingAccountsLinks
}
