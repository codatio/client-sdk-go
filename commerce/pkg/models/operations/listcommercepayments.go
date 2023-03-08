package operations

import (
	"net/http"
	"time"
)

type ListCommercePaymentsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListCommercePaymentsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCommercePaymentsRequest struct {
	PathParams  ListCommercePaymentsPathParams
	QueryParams ListCommercePaymentsQueryParams
}

type ListCommercePaymentsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommercePaymentsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommercePaymentsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommercePaymentsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommercePaymentsLinksLinks struct {
	Current  ListCommercePaymentsLinksLinksCurrent   `json:"current"`
	Next     *ListCommercePaymentsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommercePaymentsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommercePaymentsLinksLinksSelf      `json:"self"`
}

// ListCommercePaymentsLinksSourceModifiedDateNameRef
// The payment method the payment is linked to in the commerce platform.
type ListCommercePaymentsLinksSourceModifiedDateNameRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ListCommercePaymentsLinksSourceModifiedDateStatusEnum string

const (
	ListCommercePaymentsLinksSourceModifiedDateStatusEnumPending    ListCommercePaymentsLinksSourceModifiedDateStatusEnum = "Pending"
	ListCommercePaymentsLinksSourceModifiedDateStatusEnumAuthorized ListCommercePaymentsLinksSourceModifiedDateStatusEnum = "Authorized"
	ListCommercePaymentsLinksSourceModifiedDateStatusEnumPaid       ListCommercePaymentsLinksSourceModifiedDateStatusEnum = "Paid"
	ListCommercePaymentsLinksSourceModifiedDateStatusEnumFailed     ListCommercePaymentsLinksSourceModifiedDateStatusEnum = "Failed"
	ListCommercePaymentsLinksSourceModifiedDateStatusEnumCancelled  ListCommercePaymentsLinksSourceModifiedDateStatusEnum = "Cancelled"
	ListCommercePaymentsLinksSourceModifiedDateStatusEnumUnknown    ListCommercePaymentsLinksSourceModifiedDateStatusEnum = "Unknown"
)

// ListCommercePaymentsLinksSourceModifiedDate
// A payment made in a commerce platform
type ListCommercePaymentsLinksSourceModifiedDate struct {
	Amount             *float64                                               `json:"amount,omitempty"`
	CreatedDate        *time.Time                                             `json:"createdDate,omitempty"`
	Currency           *string                                                `json:"currency,omitempty"`
	DueDate            *time.Time                                             `json:"dueDate,omitempty"`
	ID                 string                                                 `json:"id"`
	ModifiedDate       *time.Time                                             `json:"modifiedDate,omitempty"`
	PaymentMethodRef   *ListCommercePaymentsLinksSourceModifiedDateNameRef    `json:"paymentMethodRef,omitempty"`
	PaymentProvider    *string                                                `json:"paymentProvider,omitempty"`
	SourceModifiedDate *time.Time                                             `json:"sourceModifiedDate,omitempty"`
	Status             *ListCommercePaymentsLinksSourceModifiedDateStatusEnum `json:"status,omitempty"`
}

// ListCommercePaymentsLinks
// Codat's Paging Model
type ListCommercePaymentsLinks struct {
	Links        ListCommercePaymentsLinksLinks                `json:"_links"`
	PageNumber   int64                                         `json:"pageNumber"`
	PageSize     int64                                         `json:"pageSize"`
	Results      []ListCommercePaymentsLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                         `json:"totalResults"`
}

type ListCommercePaymentsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListCommercePaymentsLinks
}
