package operations

import (
	"net/http"
	"time"
)

type ListSalesOrdersRequest struct {
	CompanyID string  `pathParam:"style=simple,explode=false,name=companyId"`
	OrderBy   *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page      int     `queryParam:"style=form,explode=true,name=page"`
	PageSize  *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query     *string `queryParam:"style=form,explode=true,name=query"`
}

type ListSalesOrdersLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListSalesOrdersLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListSalesOrdersLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListSalesOrdersLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListSalesOrdersLinksLinks struct {
	Current  ListSalesOrdersLinksLinksCurrent   `json:"current"`
	Next     *ListSalesOrdersLinksLinksNext     `json:"next,omitempty"`
	Previous *ListSalesOrdersLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListSalesOrdersLinksLinksSelf      `json:"self"`
}

// ListSalesOrdersLinksSourceModifiedDateCustomerRef
// The customer that the sales order is recorded against in the accounting system.
type ListSalesOrdersLinksSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type ListSalesOrdersLinksSourceModifiedDateInvoicingStatusEnum string

const (
	ListSalesOrdersLinksSourceModifiedDateInvoicingStatusEnumUnknown           ListSalesOrdersLinksSourceModifiedDateInvoicingStatusEnum = "Unknown"
	ListSalesOrdersLinksSourceModifiedDateInvoicingStatusEnumPartiallyInvoiced ListSalesOrdersLinksSourceModifiedDateInvoicingStatusEnum = "PartiallyInvoiced"
	ListSalesOrdersLinksSourceModifiedDateInvoicingStatusEnumInvoiced          ListSalesOrdersLinksSourceModifiedDateInvoicingStatusEnum = "Invoiced"
	ListSalesOrdersLinksSourceModifiedDateInvoicingStatusEnumNotInvoiced       ListSalesOrdersLinksSourceModifiedDateInvoicingStatusEnum = "NotInvoiced"
)

// ListSalesOrdersLinksSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type ListSalesOrdersLinksSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ListSalesOrdersLinksSourceModifiedDateLineItemsItemRef
// Reference to the product or inventory item to which the line item is linked.
type ListSalesOrdersLinksSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// ListSalesOrdersLinksSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type ListSalesOrdersLinksSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type ListSalesOrdersLinksSourceModifiedDateLineItemsTracking struct {
	RecordRefs []string `json:"recordRefs,omitempty"`
}

type ListSalesOrdersLinksSourceModifiedDateLineItems struct {
	AccountRef         *ListSalesOrdersLinksSourceModifiedDateLineItemsAccountRef `json:"accountRef,omitempty"`
	Description        *string                                                    `json:"description,omitempty"`
	DiscountAmount     *float64                                                   `json:"discountAmount,omitempty"`
	DiscountPercentage *float64                                                   `json:"discountPercentage,omitempty"`
	ItemRef            *ListSalesOrdersLinksSourceModifiedDateLineItemsItemRef    `json:"itemRef,omitempty"`
	Quantity           *float64                                                   `json:"quantity,omitempty"`
	SubTotal           *float64                                                   `json:"subTotal,omitempty"`
	TaxAmount          *float64                                                   `json:"taxAmount,omitempty"`
	TaxRateRef         *ListSalesOrdersLinksSourceModifiedDateLineItemsTaxRateRef `json:"taxRateRef,omitempty"`
	TotalAmount        *float64                                                   `json:"totalAmount,omitempty"`
	Tracking           *ListSalesOrdersLinksSourceModifiedDateLineItemsTracking   `json:"tracking,omitempty"`
	UnitAmount         *float64                                                   `json:"unitAmount,omitempty"`
}

type ListSalesOrdersLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type ListSalesOrdersLinksSourceModifiedDateShipToAddressTypeEnum string

const (
	ListSalesOrdersLinksSourceModifiedDateShipToAddressTypeEnumUnknown  ListSalesOrdersLinksSourceModifiedDateShipToAddressTypeEnum = "Unknown"
	ListSalesOrdersLinksSourceModifiedDateShipToAddressTypeEnumBilling  ListSalesOrdersLinksSourceModifiedDateShipToAddressTypeEnum = "Billing"
	ListSalesOrdersLinksSourceModifiedDateShipToAddressTypeEnumDelivery ListSalesOrdersLinksSourceModifiedDateShipToAddressTypeEnum = "Delivery"
)

// ListSalesOrdersLinksSourceModifiedDateShipToAddress
// Delivery address for any goods that have been ordered.
type ListSalesOrdersLinksSourceModifiedDateShipToAddress struct {
	City       *string                                                     `json:"city,omitempty"`
	Country    *string                                                     `json:"country,omitempty"`
	Line1      *string                                                     `json:"line1,omitempty"`
	Line2      *string                                                     `json:"line2,omitempty"`
	PostalCode *string                                                     `json:"postalCode,omitempty"`
	Region     *string                                                     `json:"region,omitempty"`
	Type       ListSalesOrdersLinksSourceModifiedDateShipToAddressTypeEnum `json:"type"`
}

// ListSalesOrdersLinksSourceModifiedDateShipToContact
// Details of the named contact at the delivery address.
type ListSalesOrdersLinksSourceModifiedDateShipToContact struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// ListSalesOrdersLinksSourceModifiedDateShipTo
// Delivery details for any goods that have been ordered.
type ListSalesOrdersLinksSourceModifiedDateShipTo struct {
	Address *ListSalesOrdersLinksSourceModifiedDateShipToAddress `json:"address,omitempty"`
	Contact *ListSalesOrdersLinksSourceModifiedDateShipToContact `json:"contact,omitempty"`
}

type ListSalesOrdersLinksSourceModifiedDateStatusEnum string

const (
	ListSalesOrdersLinksSourceModifiedDateStatusEnumUnknown ListSalesOrdersLinksSourceModifiedDateStatusEnum = "Unknown"
	ListSalesOrdersLinksSourceModifiedDateStatusEnumDraft   ListSalesOrdersLinksSourceModifiedDateStatusEnum = "Draft"
	ListSalesOrdersLinksSourceModifiedDateStatusEnumOpen    ListSalesOrdersLinksSourceModifiedDateStatusEnum = "Open"
	ListSalesOrdersLinksSourceModifiedDateStatusEnumClosed  ListSalesOrdersLinksSourceModifiedDateStatusEnum = "Closed"
	ListSalesOrdersLinksSourceModifiedDateStatusEnumVoid    ListSalesOrdersLinksSourceModifiedDateStatusEnum = "Void"
)

// ListSalesOrdersLinksSourceModifiedDate
// > View the coverage for sales orders in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=salesOrders" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A sales order represents a customer's intention to purchase goods or services from a seller and usually includes information such as the expected delivery date and shipping details. This information can be used to provide visibility on a business's expected receivables and to track sales through the full procurement process.
//
// A sales order is typically converted to an [invoice](https://docs.codat.io/accounting-api#/schemas/Invoice) after approval.
type ListSalesOrdersLinksSourceModifiedDate struct {
	Currency                    *string                                                    `json:"currency,omitempty"`
	CurrencyRate                *float64                                                   `json:"currencyRate,omitempty"`
	CustomerPurchaseOrderNumber *string                                                    `json:"customerPurchaseOrderNumber,omitempty"`
	CustomerRef                 *ListSalesOrdersLinksSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	ExpectedDeliveryDate        *time.Time                                                 `json:"expectedDeliveryDate,omitempty"`
	ID                          *string                                                    `json:"id,omitempty"`
	InvoicingStatus             *ListSalesOrdersLinksSourceModifiedDateInvoicingStatusEnum `json:"invoicingStatus,omitempty"`
	IssueDate                   *time.Time                                                 `json:"issueDate,omitempty"`
	LineItems                   []ListSalesOrdersLinksSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                    *ListSalesOrdersLinksSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate                *time.Time                                                 `json:"modifiedDate,omitempty"`
	Note                        *string                                                    `json:"note,omitempty"`
	SalesOrderNumber            *string                                                    `json:"salesOrderNumber,omitempty"`
	ShipTo                      *ListSalesOrdersLinksSourceModifiedDateShipTo              `json:"shipTo,omitempty"`
	SourceModifiedDate          *time.Time                                                 `json:"sourceModifiedDate,omitempty"`
	Status                      *ListSalesOrdersLinksSourceModifiedDateStatusEnum          `json:"status,omitempty"`
	SubTotal                    *float64                                                   `json:"subTotal,omitempty"`
	TotalAmount                 *float64                                                   `json:"totalAmount,omitempty"`
	TotalDiscount               *float64                                                   `json:"totalDiscount,omitempty"`
	TotalTaxAmount              *float64                                                   `json:"totalTaxAmount,omitempty"`
}

// ListSalesOrdersLinks
// Codat's Paging Model
type ListSalesOrdersLinks struct {
	Links        ListSalesOrdersLinksLinks                `json:"_links"`
	PageNumber   int64                                    `json:"pageNumber"`
	PageSize     int64                                    `json:"pageSize"`
	Results      []ListSalesOrdersLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                    `json:"totalResults"`
}

type ListSalesOrdersResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListSalesOrdersLinks
}
