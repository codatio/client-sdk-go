package operations

import (
	"github.com/codatio/client-sdk-go/banking/pkg/models/shared"
)

type ListBankingAccountBalancesPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListBankingAccountBalancesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListBankingAccountBalancesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListBankingAccountBalancesRequest struct {
	PathParams  ListBankingAccountBalancesPathParams
	QueryParams ListBankingAccountBalancesQueryParams
	Security    ListBankingAccountBalancesSecurity
}

type ListBankingAccountBalancesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListBankingAccountBalancesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListBankingAccountBalancesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListBankingAccountBalancesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListBankingAccountBalancesLinksLinks struct {
	Current  ListBankingAccountBalancesLinksLinksCurrent   `json:"current"`
	Next     *ListBankingAccountBalancesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListBankingAccountBalancesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListBankingAccountBalancesLinksLinksSelf      `json:"self"`
}

// ListBankingAccountBalancesLinks
// Codat's Paging Model
type ListBankingAccountBalancesLinks struct {
	Links        ListBankingAccountBalancesLinksLinks `json:"_links"`
	PageNumber   int64                                `json:"pageNumber"`
	PageSize     int64                                `json:"pageSize"`
	Results      *shared.AccountBalance               `json:"results,omitempty"`
	TotalResults int64                                `json:"totalResults"`
}

type ListBankingAccountBalancesResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListBankingAccountBalancesLinks
}
