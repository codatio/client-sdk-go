package operations

import (
	"net/http"
	"time"
)

type ListCommerceOrdersRequest struct {
	CompanyID    string  `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string  `pathParam:"style=simple,explode=false,name=connectionId"`
	OrderBy      *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page         int     `queryParam:"style=form,explode=true,name=page"`
	PageSize     *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query        *string `queryParam:"style=form,explode=true,name=query"`
}

type ListCommerceOrdersLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommerceOrdersLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceOrdersLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceOrdersLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommerceOrdersLinksLinks struct {
	Current  ListCommerceOrdersLinksLinksCurrent   `json:"current"`
	Next     *ListCommerceOrdersLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommerceOrdersLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommerceOrdersLinksLinksSelf      `json:"self"`
}

type ListCommerceOrdersLinksSourceModifiedDateNameRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ListCommerceOrdersLinksSourceModifiedDateOrderLineItemsDiscountAllocations struct {
	Name        *string  `json:"name,omitempty"`
	TotalAmount *float64 `json:"totalAmount,omitempty"`
}

type ListCommerceOrdersLinksSourceModifiedDateOrderLineItemsNameRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ListCommerceOrdersLinksSourceModifiedDateOrderLineItems struct {
	DiscountAllocations []ListCommerceOrdersLinksSourceModifiedDateOrderLineItemsDiscountAllocations `json:"discountAllocations,omitempty"`
	ID                  string                                                                       `json:"id"`
	ProductRef          *ListCommerceOrdersLinksSourceModifiedDateOrderLineItemsNameRef              `json:"productRef,omitempty"`
	ProductVariantRef   *ListCommerceOrdersLinksSourceModifiedDateOrderLineItemsNameRef              `json:"productVariantRef,omitempty"`
	Quantity            *float64                                                                     `json:"quantity,omitempty"`
	TaxPercentage       *float64                                                                     `json:"taxPercentage,omitempty"`
	Taxes               []interface{}                                                                `json:"taxes,omitempty"`
	TotalAmount         *float64                                                                     `json:"totalAmount,omitempty"`
	TotalTaxAmount      *float64                                                                     `json:"totalTaxAmount,omitempty"`
	UnitPrice           *float64                                                                     `json:"unitPrice,omitempty"`
}

type ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateStatusEnum string

const (
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateStatusEnumPending    ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateStatusEnum = "Pending"
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateStatusEnumAuthorized ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateStatusEnum = "Authorized"
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateStatusEnumPaid       ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateStatusEnum = "Paid"
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateStatusEnumFailed     ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateStatusEnum = "Failed"
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateStatusEnumCancelled  ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateStatusEnum = "Cancelled"
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateStatusEnumUnknown    ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateStatusEnum = "Unknown"
)

type ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnum string

const (
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnumCash        ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnum = "Cash"
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnumCard        ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnum = "Card"
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnumInvoice     ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnum = "Invoice"
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnumOnlineCard  ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnum = "OnlineCard"
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnumSwish       ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnum = "Swish"
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnumVipps       ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnum = "Vipps"
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnumMobile      ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnum = "Mobile"
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnumStoreCredit ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnum = "StoreCredit"
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnumPaypal      ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnum = "Paypal"
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnumCustom      ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnum = "Custom"
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnumPrepaid     ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnum = "Prepaid"
	ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnumUnknown     ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnum = "Unknown"
)

type ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDate struct {
	Amount             *float64                                                               `json:"amount,omitempty"`
	CreatedDate        *time.Time                                                             `json:"createdDate,omitempty"`
	Currency           *string                                                                `json:"currency,omitempty"`
	DueDate            *time.Time                                                             `json:"dueDate,omitempty"`
	ID                 string                                                                 `json:"id"`
	ModifiedDate       *time.Time                                                             `json:"modifiedDate,omitempty"`
	PaymentProvider    *string                                                                `json:"paymentProvider,omitempty"`
	SourceModifiedDate *time.Time                                                             `json:"sourceModifiedDate,omitempty"`
	Status             *ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateStatusEnum `json:"status,omitempty"`
	Type               *ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDateTypeEnum   `json:"type,omitempty"`
}

type ListCommerceOrdersLinksSourceModifiedDateServiceChargesTypeEnum string

const (
	ListCommerceOrdersLinksSourceModifiedDateServiceChargesTypeEnumGeneric     ListCommerceOrdersLinksSourceModifiedDateServiceChargesTypeEnum = "Generic"
	ListCommerceOrdersLinksSourceModifiedDateServiceChargesTypeEnumShipping    ListCommerceOrdersLinksSourceModifiedDateServiceChargesTypeEnum = "Shipping"
	ListCommerceOrdersLinksSourceModifiedDateServiceChargesTypeEnumOverpayment ListCommerceOrdersLinksSourceModifiedDateServiceChargesTypeEnum = "Overpayment"
	ListCommerceOrdersLinksSourceModifiedDateServiceChargesTypeEnumUnknown     ListCommerceOrdersLinksSourceModifiedDateServiceChargesTypeEnum = "Unknown"
)

type ListCommerceOrdersLinksSourceModifiedDateServiceCharges struct {
	Description   *string                                                          `json:"description,omitempty"`
	Quantity      *float64                                                         `json:"quantity,omitempty"`
	TaxAmount     *float64                                                         `json:"taxAmount,omitempty"`
	TaxPercentage *float64                                                         `json:"taxPercentage,omitempty"`
	Taxes         []interface{}                                                    `json:"taxes,omitempty"`
	TotalAmount   *float64                                                         `json:"totalAmount,omitempty"`
	Type          *ListCommerceOrdersLinksSourceModifiedDateServiceChargesTypeEnum `json:"type,omitempty"`
}

// ListCommerceOrdersLinksSourceModifiedDate
// Orders contain the transaction details for all products sold by the company, and include details of any payments, service charges, or refunds related to each order.
//
// From the Orders endpoints you can retrieve:
//
// A list of all the orders for a commerce company:
// `GET /companies/{companyId}/connections/{connectionId}/data/commerce-orders`.
// The details of an individual order:
// `GET /companies/{companyId}/connections/{connectionId}/data/commerce-orders/{orderId}`.
// Note that for refunds `quantity` is a negative value and `unitPrice` is a positive value.
type ListCommerceOrdersLinksSourceModifiedDate struct {
	ClosedDate         *time.Time                                                    `json:"closedDate,omitempty"`
	Country            *string                                                       `json:"country,omitempty"`
	CreatedDate        *time.Time                                                    `json:"createdDate,omitempty"`
	Currency           *string                                                       `json:"currency,omitempty"`
	CustomerRef        *ListCommerceOrdersLinksSourceModifiedDateNameRef             `json:"customerRef,omitempty"`
	ID                 string                                                        `json:"id"`
	LocationRef        *ListCommerceOrdersLinksSourceModifiedDateNameRef             `json:"locationRef,omitempty"`
	ModifiedDate       *time.Time                                                    `json:"modifiedDate,omitempty"`
	OrderLineItems     []ListCommerceOrdersLinksSourceModifiedDateOrderLineItems     `json:"orderLineItems,omitempty"`
	OrderNumber        *string                                                       `json:"orderNumber,omitempty"`
	Payments           []ListCommerceOrdersLinksSourceModifiedDateSourceModifiedDate `json:"payments,omitempty"`
	ServiceCharges     []ListCommerceOrdersLinksSourceModifiedDateServiceCharges     `json:"serviceCharges,omitempty"`
	SourceModifiedDate *time.Time                                                    `json:"sourceModifiedDate,omitempty"`
	TotalAmount        *float64                                                      `json:"totalAmount,omitempty"`
	TotalDiscount      *float64                                                      `json:"totalDiscount,omitempty"`
	TotalGratuity      *float64                                                      `json:"totalGratuity,omitempty"`
	TotalRefund        *float64                                                      `json:"totalRefund,omitempty"`
	TotalTaxAmount     *float64                                                      `json:"totalTaxAmount,omitempty"`
}

// ListCommerceOrdersLinks
// Codat's Paging Model
type ListCommerceOrdersLinks struct {
	Links        ListCommerceOrdersLinksLinks                `json:"_links"`
	PageNumber   int64                                       `json:"pageNumber"`
	PageSize     int64                                       `json:"pageSize"`
	Results      []ListCommerceOrdersLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                       `json:"totalResults"`
}

type ListCommerceOrdersResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListCommerceOrdersLinks
}
