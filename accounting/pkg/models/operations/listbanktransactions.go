package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListBankTransactionsPathParams struct {
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListBankTransactionsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListBankTransactionsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListBankTransactionsRequest struct {
	PathParams  ListBankTransactionsPathParams
	QueryParams ListBankTransactionsQueryParams
	Security    ListBankTransactionsSecurity
}

type ListBankTransactionsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListBankTransactionsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListBankTransactionsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListBankTransactionsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListBankTransactionsLinksLinks struct {
	Current  ListBankTransactionsLinksLinksCurrent   `json:"current"`
	Next     *ListBankTransactionsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListBankTransactionsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListBankTransactionsLinksLinksSelf      `json:"self"`
}

// ListBankTransactionsLinks
// Codat's Paging Model
type ListBankTransactionsLinks struct {
	Links        ListBankTransactionsLinksLinks `json:"_links"`
	PageNumber   int64                          `json:"pageNumber"`
	PageSize     int64                          `json:"pageSize"`
	Results      []shared.BankStatementLine     `json:"results,omitempty"`
	TotalResults int64                          `json:"totalResults"`
}

type ListBankTransactionsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListBankTransactionsLinks
}
