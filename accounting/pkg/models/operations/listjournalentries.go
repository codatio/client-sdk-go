package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListJournalEntriesPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListJournalEntriesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListJournalEntriesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListJournalEntriesRequest struct {
	PathParams  ListJournalEntriesPathParams
	QueryParams ListJournalEntriesQueryParams
	Security    ListJournalEntriesSecurity
}

type ListJournalEntriesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListJournalEntriesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListJournalEntriesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListJournalEntriesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListJournalEntriesLinksLinks struct {
	Current  ListJournalEntriesLinksLinksCurrent   `json:"current"`
	Next     *ListJournalEntriesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListJournalEntriesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListJournalEntriesLinksLinksSelf      `json:"self"`
}

// ListJournalEntriesLinks
// Codat's Paging Model
type ListJournalEntriesLinks struct {
	Links        ListJournalEntriesLinksLinks `json:"_links"`
	PageNumber   int64                        `json:"pageNumber"`
	PageSize     int64                        `json:"pageSize"`
	Results      []shared.JournalEntry        `json:"results,omitempty"`
	TotalResults int64                        `json:"totalResults"`
}

type ListJournalEntriesResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListJournalEntriesLinks
}
