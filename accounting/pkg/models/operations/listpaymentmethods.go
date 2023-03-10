package operations

import (
	"net/http"
	"time"
)

type ListPaymentMethodsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListPaymentMethodsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListPaymentMethodsRequest struct {
	PathParams  ListPaymentMethodsPathParams
	QueryParams ListPaymentMethodsQueryParams
}

type ListPaymentMethodsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListPaymentMethodsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListPaymentMethodsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListPaymentMethodsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListPaymentMethodsLinksLinks struct {
	Current  ListPaymentMethodsLinksLinksCurrent   `json:"current"`
	Next     *ListPaymentMethodsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListPaymentMethodsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListPaymentMethodsLinksLinksSelf      `json:"self"`
}

type ListPaymentMethodsLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type ListPaymentMethodsLinksSourceModifiedDateStatusEnum string

const (
	ListPaymentMethodsLinksSourceModifiedDateStatusEnumUnknown  ListPaymentMethodsLinksSourceModifiedDateStatusEnum = "Unknown"
	ListPaymentMethodsLinksSourceModifiedDateStatusEnumActive   ListPaymentMethodsLinksSourceModifiedDateStatusEnum = "Active"
	ListPaymentMethodsLinksSourceModifiedDateStatusEnumArchived ListPaymentMethodsLinksSourceModifiedDateStatusEnum = "Archived"
)

type ListPaymentMethodsLinksSourceModifiedDateTypeEnum string

const (
	ListPaymentMethodsLinksSourceModifiedDateTypeEnumUnknown      ListPaymentMethodsLinksSourceModifiedDateTypeEnum = "Unknown"
	ListPaymentMethodsLinksSourceModifiedDateTypeEnumCash         ListPaymentMethodsLinksSourceModifiedDateTypeEnum = "Cash"
	ListPaymentMethodsLinksSourceModifiedDateTypeEnumCheck        ListPaymentMethodsLinksSourceModifiedDateTypeEnum = "Check"
	ListPaymentMethodsLinksSourceModifiedDateTypeEnumCreditCard   ListPaymentMethodsLinksSourceModifiedDateTypeEnum = "CreditCard"
	ListPaymentMethodsLinksSourceModifiedDateTypeEnumDebitCard    ListPaymentMethodsLinksSourceModifiedDateTypeEnum = "DebitCard"
	ListPaymentMethodsLinksSourceModifiedDateTypeEnumBankTransfer ListPaymentMethodsLinksSourceModifiedDateTypeEnum = "BankTransfer"
	ListPaymentMethodsLinksSourceModifiedDateTypeEnumOther        ListPaymentMethodsLinksSourceModifiedDateTypeEnum = "Other"
)

// ListPaymentMethodsLinksSourceModifiedDate
// > View the coverage for payment methods in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=paymentMethods" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A Payment Method represents the payment method(s) used to pay a Bill. Payment Methods are referenced on [Bill Payments](https://docs.codat.io/accounting-api#/schemas/BillPayment) and [Payments](https://docs.codat.io/accounting-api#/schemas/Payment).
//
// From the Payment Methods endpoints you can retrieve:
//
//   - A list of all the Payment Methods used by a company: `GET/companies/{companyId}/data/paymentMethods`.
//   - The details of an individual Payment Method:
//     `GET /companies/{companyId}/data/paymentMethods/{paymentMethodId}`.
type ListPaymentMethodsLinksSourceModifiedDate struct {
	ID                 *string                                              `json:"id,omitempty"`
	Metadata           *ListPaymentMethodsLinksSourceModifiedDateMetadata   `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                           `json:"modifiedDate,omitempty"`
	Name               *string                                              `json:"name,omitempty"`
	SourceModifiedDate *time.Time                                           `json:"sourceModifiedDate,omitempty"`
	Status             *ListPaymentMethodsLinksSourceModifiedDateStatusEnum `json:"status,omitempty"`
	Type               *ListPaymentMethodsLinksSourceModifiedDateTypeEnum   `json:"type,omitempty"`
}

// ListPaymentMethodsLinks
// Codat's Paging Model
type ListPaymentMethodsLinks struct {
	Links        ListPaymentMethodsLinksLinks                `json:"_links"`
	PageNumber   int64                                       `json:"pageNumber"`
	PageSize     int64                                       `json:"pageSize"`
	Results      []ListPaymentMethodsLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                       `json:"totalResults"`
}

type ListPaymentMethodsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListPaymentMethodsLinks
}
