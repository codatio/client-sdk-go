package operations

import (
	"net/http"
	"time"
)

type ListBankingAccountBalancesPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListBankingAccountBalancesQueryParams struct {
	OrderBy  *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page     int     `queryParam:"style=form,explode=true,name=page"`
	PageSize *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string `queryParam:"style=form,explode=true,name=query"`
}

type ListBankingAccountBalancesRequest struct {
	PathParams  ListBankingAccountBalancesPathParams
	QueryParams ListBankingAccountBalancesQueryParams
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

// ListBankingAccountBalancesLinksSourceModifiedDateBalance
// An object containing bank balance data.
type ListBankingAccountBalancesLinksSourceModifiedDateBalance struct {
	Available *float64 `json:"available,omitempty"`
	Current   float64  `json:"current"`
	Limit     *float64 `json:"limit,omitempty"`
}

// ListBankingAccountBalancesLinksSourceModifiedDate
// The Banking Account Balances data type provides a list of balances for a bank account including end-of-day batch balance or running balances per transaction.
type ListBankingAccountBalancesLinksSourceModifiedDate struct {
	AccountID          string                                                   `json:"accountId"`
	Balance            ListBankingAccountBalancesLinksSourceModifiedDateBalance `json:"balance"`
	Date               time.Time                                                `json:"date"`
	ModifiedDate       *time.Time                                               `json:"modifiedDate,omitempty"`
	SourceModifiedDate *time.Time                                               `json:"sourceModifiedDate,omitempty"`
}

// ListBankingAccountBalancesLinks
// Codat's Paging Model
type ListBankingAccountBalancesLinks struct {
	Links        ListBankingAccountBalancesLinksLinks               `json:"_links"`
	PageNumber   int64                                              `json:"pageNumber"`
	PageSize     int64                                              `json:"pageSize"`
	Results      *ListBankingAccountBalancesLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                              `json:"totalResults"`
}

type ListBankingAccountBalancesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListBankingAccountBalancesLinks
}
