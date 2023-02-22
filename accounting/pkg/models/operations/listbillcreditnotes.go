package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListBillCreditNotesPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListBillCreditNotesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListBillCreditNotesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListBillCreditNotesRequest struct {
	PathParams  ListBillCreditNotesPathParams
	QueryParams ListBillCreditNotesQueryParams
	Security    ListBillCreditNotesSecurity
}

type ListBillCreditNotesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListBillCreditNotesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListBillCreditNotesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListBillCreditNotesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListBillCreditNotesLinksLinks struct {
	Current  ListBillCreditNotesLinksLinksCurrent   `json:"current"`
	Next     *ListBillCreditNotesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListBillCreditNotesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListBillCreditNotesLinksLinksSelf      `json:"self"`
}

// ListBillCreditNotesLinks
// Codat's Paging Model
type ListBillCreditNotesLinks struct {
	Links        ListBillCreditNotesLinksLinks `json:"_links"`
	PageNumber   int64                         `json:"pageNumber"`
	PageSize     int64                         `json:"pageSize"`
	Results      []shared.BillCreditNote       `json:"results,omitempty"`
	TotalResults int64                         `json:"totalResults"`
}

type ListBillCreditNotesResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListBillCreditNotesLinks
}
