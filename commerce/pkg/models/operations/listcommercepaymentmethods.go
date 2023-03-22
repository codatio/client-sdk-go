// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"time"
)

type ListCommercePaymentMethodsRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Field to order results by. [Read more](https://docs.codat.io/using-the-api/ordering-results).
	OrderBy *string `queryParam:"style=form,explode=true,name=orderBy"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page int `queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `queryParam:"style=form,explode=true,name=pageSize"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

type ListCommercePaymentMethodsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommercePaymentMethodsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommercePaymentMethodsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommercePaymentMethodsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommercePaymentMethodsLinksLinks struct {
	Current  ListCommercePaymentMethodsLinksLinksCurrent   `json:"current"`
	Next     *ListCommercePaymentMethodsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommercePaymentMethodsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommercePaymentMethodsLinksLinksSelf      `json:"self"`
}

// ListCommercePaymentMethodsLinksSourceModifiedDateStatusEnum - Status of the Payment Method
type ListCommercePaymentMethodsLinksSourceModifiedDateStatusEnum string

const (
	ListCommercePaymentMethodsLinksSourceModifiedDateStatusEnumActive   ListCommercePaymentMethodsLinksSourceModifiedDateStatusEnum = "Active"
	ListCommercePaymentMethodsLinksSourceModifiedDateStatusEnumArchived ListCommercePaymentMethodsLinksSourceModifiedDateStatusEnum = "Archived"
	ListCommercePaymentMethodsLinksSourceModifiedDateStatusEnumUnknown  ListCommercePaymentMethodsLinksSourceModifiedDateStatusEnum = "Unknown"
)

type ListCommercePaymentMethodsLinksSourceModifiedDate struct {
	// A unique, persistent identifier for this record
	ID string `json:"id"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	// The name of the PaymentMethod
	Name *string `json:"name,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
	// Status of the Payment Method
	Status *ListCommercePaymentMethodsLinksSourceModifiedDateStatusEnum `json:"status,omitempty"`
}

// ListCommercePaymentMethodsLinks - Codat's Paging Model
type ListCommercePaymentMethodsLinks struct {
	Links        ListCommercePaymentMethodsLinksLinks                `json:"_links"`
	PageNumber   int64                                               `json:"pageNumber"`
	PageSize     int64                                               `json:"pageSize"`
	Results      []ListCommercePaymentMethodsLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                               `json:"totalResults"`
}

type ListCommercePaymentMethodsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Links *ListCommercePaymentMethodsLinks
}
