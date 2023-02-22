package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
)

type ListCommerceDisputesPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListCommerceDisputesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCommerceDisputesRequest struct {
	PathParams  ListCommerceDisputesPathParams
	QueryParams ListCommerceDisputesQueryParams
}

type ListCommerceDisputesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommerceDisputesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceDisputesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceDisputesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommerceDisputesLinksLinks struct {
	Current  ListCommerceDisputesLinksLinksCurrent   `json:"current"`
	Next     *ListCommerceDisputesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommerceDisputesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommerceDisputesLinksLinksSelf      `json:"self"`
}

// ListCommerceDisputesLinks
// Codat's Paging Model
type ListCommerceDisputesLinks struct {
	Links        ListCommerceDisputesLinksLinks `json:"_links"`
	PageNumber   int64                          `json:"pageNumber"`
	PageSize     int64                          `json:"pageSize"`
	Results      []shared.Dispute               `json:"results,omitempty"`
	TotalResults int64                          `json:"totalResults"`
}

type ListCommerceDisputesResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListCommerceDisputesLinks
}
