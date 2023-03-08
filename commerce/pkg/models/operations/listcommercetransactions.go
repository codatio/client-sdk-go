package operations

import (
	"net/http"
	"time"
)

type ListCommerceTransactionsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListCommerceTransactionsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCommerceTransactionsRequest struct {
	PathParams  ListCommerceTransactionsPathParams
	QueryParams ListCommerceTransactionsQueryParams
}

type ListCommerceTransactionsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommerceTransactionsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceTransactionsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceTransactionsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommerceTransactionsLinksLinks struct {
	Current  ListCommerceTransactionsLinksLinksCurrent   `json:"current"`
	Next     *ListCommerceTransactionsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommerceTransactionsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommerceTransactionsLinksLinksSelf      `json:"self"`
}

type ListCommerceTransactionsLinksSourceModifiedDateRecordRef struct {
	ID   string      `json:"id"`
	Type interface{} `json:"type"`
}

type ListCommerceTransactionsLinksSourceModifiedDateTypeEnum string

const (
	ListCommerceTransactionsLinksSourceModifiedDateTypeEnumPayment          ListCommerceTransactionsLinksSourceModifiedDateTypeEnum = "Payment"
	ListCommerceTransactionsLinksSourceModifiedDateTypeEnumRefund           ListCommerceTransactionsLinksSourceModifiedDateTypeEnum = "Refund"
	ListCommerceTransactionsLinksSourceModifiedDateTypeEnumPayout           ListCommerceTransactionsLinksSourceModifiedDateTypeEnum = "Payout"
	ListCommerceTransactionsLinksSourceModifiedDateTypeEnumFailedPayout     ListCommerceTransactionsLinksSourceModifiedDateTypeEnum = "FailedPayout"
	ListCommerceTransactionsLinksSourceModifiedDateTypeEnumTransfer         ListCommerceTransactionsLinksSourceModifiedDateTypeEnum = "Transfer"
	ListCommerceTransactionsLinksSourceModifiedDateTypeEnumPaymentFee       ListCommerceTransactionsLinksSourceModifiedDateTypeEnum = "PaymentFee"
	ListCommerceTransactionsLinksSourceModifiedDateTypeEnumPaymentFeeRefund ListCommerceTransactionsLinksSourceModifiedDateTypeEnum = "PaymentFeeRefund"
	ListCommerceTransactionsLinksSourceModifiedDateTypeEnumUnknown          ListCommerceTransactionsLinksSourceModifiedDateTypeEnum = "Unknown"
)

// ListCommerceTransactionsLinksSourceModifiedDate
// A financial transaction recorded in the commerce or point of sale system
type ListCommerceTransactionsLinksSourceModifiedDate struct {
	CreatedDate          *time.Time                                                `json:"createdDate,omitempty"`
	Currency             *string                                                   `json:"currency,omitempty"`
	ID                   string                                                    `json:"id"`
	ModifiedDate         *time.Time                                                `json:"modifiedDate,omitempty"`
	SourceCreatedDate    *time.Time                                                `json:"sourceCreatedDate,omitempty"`
	SourceModifiedDate   *time.Time                                                `json:"sourceModifiedDate,omitempty"`
	SubType              *string                                                   `json:"subType,omitempty"`
	TotalAmount          *float64                                                  `json:"totalAmount,omitempty"`
	TransactionSourceRef *ListCommerceTransactionsLinksSourceModifiedDateRecordRef `json:"transactionSourceRef,omitempty"`
	Type                 *ListCommerceTransactionsLinksSourceModifiedDateTypeEnum  `json:"type,omitempty"`
}

// ListCommerceTransactionsLinks
// Codat's Paging Model
type ListCommerceTransactionsLinks struct {
	Links        ListCommerceTransactionsLinksLinks                `json:"_links"`
	PageNumber   int64                                             `json:"pageNumber"`
	PageSize     int64                                             `json:"pageSize"`
	Results      []ListCommerceTransactionsLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                             `json:"totalResults"`
}

type ListCommerceTransactionsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListCommerceTransactionsLinks
}
