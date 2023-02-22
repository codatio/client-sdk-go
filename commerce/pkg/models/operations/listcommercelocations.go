package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
)

type ListCommerceLocationsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListCommerceLocationsRequest struct {
	PathParams ListCommerceLocationsPathParams
}

type ListCommerceLocationsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommerceLocationsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceLocationsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceLocationsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommerceLocationsLinksLinks struct {
	Current  ListCommerceLocationsLinksLinksCurrent   `json:"current"`
	Next     *ListCommerceLocationsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommerceLocationsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommerceLocationsLinksLinksSelf      `json:"self"`
}

// ListCommerceLocationsLinks
// Codat's Paging Model
type ListCommerceLocationsLinks struct {
	Links        ListCommerceLocationsLinksLinks `json:"_links"`
	PageNumber   int64                           `json:"pageNumber"`
	PageSize     int64                           `json:"pageSize"`
	Results      []shared.Locations              `json:"results,omitempty"`
	TotalResults int64                           `json:"totalResults"`
}

type ListCommerceLocationsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListCommerceLocationsLinks
}
