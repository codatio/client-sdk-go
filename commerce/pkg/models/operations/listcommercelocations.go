package operations

import (
	"time"
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

// ListCommerceLocationsLinksSourceModifiedDateAddress
// Address associated with the location
type ListCommerceLocationsLinksSourceModifiedDateAddress struct {
	City       *string     `json:"city,omitempty"`
	Country    *string     `json:"country,omitempty"`
	Line1      *string     `json:"line1,omitempty"`
	Line2      *string     `json:"line2,omitempty"`
	PostalCode *string     `json:"postalCode,omitempty"`
	Region     *string     `json:"region,omitempty"`
	Type       interface{} `json:"type,omitempty"`
}

// ListCommerceLocationsLinksSourceModifiedDate
// The Locations datatype holds information on geographic locations at which stocks of products may be held, as referenced in the Products data type.
//
// Locations also holds information on geographic locations where orders were placed, as referenced in the Orders data type.
//
// From the Locations endpoints you can retrieve:
//
// A list of all the Locations of a commerce company: `GET /companies/{companyId}/connections/{connectionId}/data/commerce-locations`.
// The details of an individual location: `GET /companies/{companyId}/connections/{connectionId}/data/commerce-locations/{locationId}`.
type ListCommerceLocationsLinksSourceModifiedDate struct {
	Address            *ListCommerceLocationsLinksSourceModifiedDateAddress `json:"address,omitempty"`
	ID                 string                                               `json:"id"`
	ModifiedDate       *time.Time                                           `json:"modifiedDate,omitempty"`
	Name               *string                                              `json:"name,omitempty"`
	SourceModifiedDate *time.Time                                           `json:"sourceModifiedDate,omitempty"`
}

// ListCommerceLocationsLinks
// Codat's Paging Model
type ListCommerceLocationsLinks struct {
	Links        ListCommerceLocationsLinksLinks                `json:"_links"`
	PageNumber   int64                                          `json:"pageNumber"`
	PageSize     int64                                          `json:"pageSize"`
	Results      []ListCommerceLocationsLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                          `json:"totalResults"`
}

type ListCommerceLocationsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListCommerceLocationsLinks
}
