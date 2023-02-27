package operations

import (
	"time"
)

type ListCommerceCustomersPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListCommerceCustomersQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCommerceCustomersRequest struct {
	PathParams  ListCommerceCustomersPathParams
	QueryParams ListCommerceCustomersQueryParams
}

type ListCommerceCustomersLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommerceCustomersLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceCustomersLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceCustomersLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommerceCustomersLinksLinks struct {
	Current  ListCommerceCustomersLinksLinksCurrent   `json:"current"`
	Next     *ListCommerceCustomersLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommerceCustomersLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommerceCustomersLinksLinksSelf      `json:"self"`
}

type ListCommerceCustomersLinksSourceModifiedDateAddress struct {
	City       *string      `json:"city,omitempty"`
	Country    *string      `json:"country,omitempty"`
	Line1      *string      `json:"line1,omitempty"`
	Line2      *string      `json:"line2,omitempty"`
	PostalCode *string      `json:"postalCode,omitempty"`
	Region     *string      `json:"region,omitempty"`
	Type       *interface{} `json:"type,omitempty"`
}

// ListCommerceCustomersLinksSourceModifiedDate
// Represents a customer who has placed an order in the commerce system"
type ListCommerceCustomersLinksSourceModifiedDate struct {
	Addresses          []ListCommerceCustomersLinksSourceModifiedDateAddress `json:"addresses,omitempty"`
	CreatedDate        *time.Time                                            `json:"createdDate,omitempty"`
	CustomerName       *string                                               `json:"customerName,omitempty"`
	DefaultCurrency    map[string]interface{}                                `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                               `json:"emailAddress,omitempty"`
	ID                 string                                                `json:"id"`
	ModifiedDate       *time.Time                                            `json:"modifiedDate,omitempty"`
	Note               *string                                               `json:"note,omitempty"`
	Phone              *string                                               `json:"phone,omitempty"`
	SourceModifiedDate *time.Time                                            `json:"sourceModifiedDate,omitempty"`
}

// ListCommerceCustomersLinks
// Codat's Paging Model
type ListCommerceCustomersLinks struct {
	Links        ListCommerceCustomersLinksLinks                `json:"_links"`
	PageNumber   int64                                          `json:"pageNumber"`
	PageSize     int64                                          `json:"pageSize"`
	Results      []ListCommerceCustomersLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                          `json:"totalResults"`
}

type ListCommerceCustomersResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListCommerceCustomersLinks
}
