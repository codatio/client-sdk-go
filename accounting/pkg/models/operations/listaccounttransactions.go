package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListAccountTransactionsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListAccountTransactionsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListAccountTransactionsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListAccountTransactionsRequest struct {
	PathParams  ListAccountTransactionsPathParams
	QueryParams ListAccountTransactionsQueryParams
	Security    ListAccountTransactionsSecurity
}

type ListAccountTransactionsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListAccountTransactionsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListAccountTransactionsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListAccountTransactionsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListAccountTransactionsLinksLinks struct {
	Current  ListAccountTransactionsLinksLinksCurrent   `json:"current"`
	Next     *ListAccountTransactionsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListAccountTransactionsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListAccountTransactionsLinksLinksSelf      `json:"self"`
}

// ListAccountTransactionsLinks
// Codat's Paging Model
type ListAccountTransactionsLinks struct {
	Links        ListAccountTransactionsLinksLinks `json:"_links"`
	PageNumber   int64                             `json:"pageNumber"`
	PageSize     int64                             `json:"pageSize"`
	Results      []shared.AccountTransaction       `json:"results,omitempty"`
	TotalResults int64                             `json:"totalResults"`
}

type ListAccountTransactionsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListAccountTransactionsLinks
}
