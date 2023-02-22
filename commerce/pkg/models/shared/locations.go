package shared

import (
	"time"
)

// LocationsAddress
// Address associated with the location
type LocationsAddress struct {
	City       *string      `json:"city,omitempty"`
	Country    *string      `json:"country,omitempty"`
	Line1      *string      `json:"line1,omitempty"`
	Line2      *string      `json:"line2,omitempty"`
	PostalCode *string      `json:"postalCode,omitempty"`
	Region     *string      `json:"region,omitempty"`
	Type       *interface{} `json:"type,omitempty"`
}

// Locations
// The Locations datatype holds information on geographic locations at which stocks of products may be held, as referenced in the Products data type.
//
// Locations also holds information on geographic locations where orders were placed, as referenced in the Orders data type.
//
// From the Locations endpoints you can retrieve:
//
// A list of all the Locations of a commerce company: `GET /companies/{companyId}/connections/{connectionId}/data/commerce-locations`.
// The details of an individual location: `GET /companies/{companyId}/connections/{connectionId}/data/commerce-locations/{locationId}`.
type Locations struct {
	Address            *LocationsAddress `json:"address,omitempty"`
	ID                 string            `json:"id"`
	ModifiedDate       *time.Time        `json:"modifiedDate,omitempty"`
	Name               *string           `json:"name,omitempty"`
	SourceModifiedDate *time.Time        `json:"sourceModifiedDate,omitempty"`
}
