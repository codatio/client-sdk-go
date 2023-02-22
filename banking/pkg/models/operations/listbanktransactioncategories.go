package operations

import (
	"github.com/codatio/client-sdk-go/banking/pkg/models/shared"
)

type ListBankTransactionCategoriesPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListBankTransactionCategoriesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListBankTransactionCategoriesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListBankTransactionCategoriesRequest struct {
	PathParams  ListBankTransactionCategoriesPathParams
	QueryParams ListBankTransactionCategoriesQueryParams
	Security    ListBankTransactionCategoriesSecurity
}

type ListBankTransactionCategoriesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListBankTransactionCategoriesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListBankTransactionCategoriesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListBankTransactionCategoriesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListBankTransactionCategoriesLinksLinks struct {
	Current  ListBankTransactionCategoriesLinksLinksCurrent   `json:"current"`
	Next     *ListBankTransactionCategoriesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListBankTransactionCategoriesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListBankTransactionCategoriesLinksLinksSelf      `json:"self"`
}

// ListBankTransactionCategoriesLinks
// Codat's Paging Model
type ListBankTransactionCategoriesLinks struct {
	Links        ListBankTransactionCategoriesLinksLinks `json:"_links"`
	PageNumber   int64                                   `json:"pageNumber"`
	PageSize     int64                                   `json:"pageSize"`
	Results      *shared.TransactionCategory             `json:"results,omitempty"`
	TotalResults int64                                   `json:"totalResults"`
}

type ListBankTransactionCategoriesResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListBankTransactionCategoriesLinks
}
