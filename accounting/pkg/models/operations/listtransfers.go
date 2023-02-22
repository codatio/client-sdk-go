package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListTransfersPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListTransfersQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListTransfersSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListTransfersRequest struct {
	PathParams  ListTransfersPathParams
	QueryParams ListTransfersQueryParams
	Security    ListTransfersSecurity
}

type ListTransfersLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListTransfersLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListTransfersLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListTransfersLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListTransfersLinksLinks struct {
	Current  ListTransfersLinksLinksCurrent   `json:"current"`
	Next     *ListTransfersLinksLinksNext     `json:"next,omitempty"`
	Previous *ListTransfersLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListTransfersLinksLinksSelf      `json:"self"`
}

// ListTransfersLinks
// Codat's Paging Model
type ListTransfersLinks struct {
	Links        ListTransfersLinksLinks `json:"_links"`
	PageNumber   int64                   `json:"pageNumber"`
	PageSize     int64                   `json:"pageSize"`
	Results      []shared.Transfer       `json:"results,omitempty"`
	TotalResults int64                   `json:"totalResults"`
}

type ListTransfersResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListTransfersLinks
}
