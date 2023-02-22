package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListJournalsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListJournalsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListJournalsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListJournalsRequest struct {
	PathParams  ListJournalsPathParams
	QueryParams ListJournalsQueryParams
	Security    ListJournalsSecurity
}

type ListJournalsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListJournalsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListJournalsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListJournalsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListJournalsLinksLinks struct {
	Current  ListJournalsLinksLinksCurrent   `json:"current"`
	Next     *ListJournalsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListJournalsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListJournalsLinksLinksSelf      `json:"self"`
}

// ListJournalsLinks
// Codat's Paging Model
type ListJournalsLinks struct {
	Links        ListJournalsLinksLinks `json:"_links"`
	PageNumber   int64                  `json:"pageNumber"`
	PageSize     int64                  `json:"pageSize"`
	Results      []shared.Journal       `json:"results,omitempty"`
	TotalResults int64                  `json:"totalResults"`
}

type ListJournalsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListJournalsLinks
}
