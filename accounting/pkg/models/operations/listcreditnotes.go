package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListCreditNotesPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListCreditNotesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCreditNotesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListCreditNotesRequest struct {
	PathParams  ListCreditNotesPathParams
	QueryParams ListCreditNotesQueryParams
	Security    ListCreditNotesSecurity
}

type ListCreditNotesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCreditNotesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCreditNotesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCreditNotesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCreditNotesLinksLinks struct {
	Current  ListCreditNotesLinksLinksCurrent   `json:"current"`
	Next     *ListCreditNotesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCreditNotesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCreditNotesLinksLinksSelf      `json:"self"`
}

// ListCreditNotesLinks
// Codat's Paging Model
type ListCreditNotesLinks struct {
	Links        ListCreditNotesLinksLinks `json:"_links"`
	PageNumber   int64                     `json:"pageNumber"`
	PageSize     int64                     `json:"pageSize"`
	Results      []shared.CreditNote       `json:"results,omitempty"`
	TotalResults int64                     `json:"totalResults"`
}

type ListCreditNotesResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListCreditNotesLinks
}
