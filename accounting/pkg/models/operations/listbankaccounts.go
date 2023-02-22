package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListBankAccountsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListBankAccountsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListBankAccountsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListBankAccountsRequest struct {
	PathParams  ListBankAccountsPathParams
	QueryParams ListBankAccountsQueryParams
	Security    ListBankAccountsSecurity
}

type ListBankAccountsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListBankAccountsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListBankAccountsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListBankAccountsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListBankAccountsLinksLinks struct {
	Current  ListBankAccountsLinksLinksCurrent   `json:"current"`
	Next     *ListBankAccountsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListBankAccountsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListBankAccountsLinksLinksSelf      `json:"self"`
}

// ListBankAccountsLinks
// Codat's Paging Model
type ListBankAccountsLinks struct {
	Links        ListBankAccountsLinksLinks `json:"_links"`
	PageNumber   int64                      `json:"pageNumber"`
	PageSize     int64                      `json:"pageSize"`
	Results      []shared.BankAccount       `json:"results,omitempty"`
	TotalResults int64                      `json:"totalResults"`
}

type ListBankAccountsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListBankAccountsLinks
}
