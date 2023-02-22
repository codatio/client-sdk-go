package operations

import (
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
)

type ListAllBankTransactionscountPathParams struct {
	AccountID    string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListAllBankTransactionscountQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListAllBankTransactionscountSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListAllBankTransactionscountRequest struct {
	PathParams  ListAllBankTransactionscountPathParams
	QueryParams ListAllBankTransactionscountQueryParams
	Security    ListAllBankTransactionscountSecurity
}

type ListAllBankTransactionscountLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListAllBankTransactionscountLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListAllBankTransactionscountLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListAllBankTransactionscountLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListAllBankTransactionscountLinksLinks struct {
	Current  ListAllBankTransactionscountLinksLinksCurrent   `json:"current"`
	Next     *ListAllBankTransactionscountLinksLinksNext     `json:"next,omitempty"`
	Previous *ListAllBankTransactionscountLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListAllBankTransactionscountLinksLinksSelf      `json:"self"`
}

// ListAllBankTransactionscountLinks
// Codat's Paging Model
type ListAllBankTransactionscountLinks struct {
	Links        ListAllBankTransactionscountLinksLinks `json:"_links"`
	PageNumber   int64                                  `json:"pageNumber"`
	PageSize     int64                                  `json:"pageSize"`
	Results      []shared.BankTransactions              `json:"results,omitempty"`
	TotalResults int64                                  `json:"totalResults"`
}

type ListAllBankTransactionscountResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListAllBankTransactionscountLinks
}
