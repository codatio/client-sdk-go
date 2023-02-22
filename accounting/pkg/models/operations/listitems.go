package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListItemsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListItemsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListItemsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListItemsRequest struct {
	PathParams  ListItemsPathParams
	QueryParams ListItemsQueryParams
	Security    ListItemsSecurity
}

type ListItemsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListItemsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListItemsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListItemsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListItemsLinksLinks struct {
	Current  ListItemsLinksLinksCurrent   `json:"current"`
	Next     *ListItemsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListItemsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListItemsLinksLinksSelf      `json:"self"`
}

// ListItemsLinks
// Codat's Paging Model
type ListItemsLinks struct {
	Links        ListItemsLinksLinks `json:"_links"`
	PageNumber   int64               `json:"pageNumber"`
	PageSize     int64               `json:"pageSize"`
	Results      []shared.Item       `json:"results,omitempty"`
	TotalResults int64               `json:"totalResults"`
}

type ListItemsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListItemsLinks
}
