package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type ListPurchaseOrdersPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListPurchaseOrdersQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListPurchaseOrdersSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListPurchaseOrdersRequest struct {
	PathParams  ListPurchaseOrdersPathParams
	QueryParams ListPurchaseOrdersQueryParams
	Security    ListPurchaseOrdersSecurity
}

type ListPurchaseOrdersLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListPurchaseOrdersLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListPurchaseOrdersLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListPurchaseOrdersLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListPurchaseOrdersLinksLinks struct {
	Current  ListPurchaseOrdersLinksLinksCurrent   `json:"current"`
	Next     *ListPurchaseOrdersLinksLinksNext     `json:"next,omitempty"`
	Previous *ListPurchaseOrdersLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListPurchaseOrdersLinksLinksSelf      `json:"self"`
}

// ListPurchaseOrdersLinksSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type ListPurchaseOrdersLinksSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ListPurchaseOrdersLinksSourceModifiedDateLineItemsItemRef
// Reference to the product or inventory item to which the line item is linked.
type ListPurchaseOrdersLinksSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// ListPurchaseOrdersLinksSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type ListPurchaseOrdersLinksSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type ListPurchaseOrdersLinksSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type ListPurchaseOrdersLinksSourceModifiedDateLineItems struct {
	AccountRef           *ListPurchaseOrdersLinksSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                  `json:"description,omitempty"`
	DiscountAmount       *float64                                                                 `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                 `json:"discountPercentage,omitempty"`
	ItemRef              *ListPurchaseOrdersLinksSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             *float64                                                                 `json:"quantity,omitempty"`
	SubTotal             *float64                                                                 `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                 `json:"taxAmount,omitempty"`
	TaxRateRef           *ListPurchaseOrdersLinksSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                 `json:"totalAmount,omitempty"`
	TrackingCategoryRefs []ListPurchaseOrdersLinksSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           *float64                                                                 `json:"unitAmount,omitempty"`
}

type ListPurchaseOrdersLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type ListPurchaseOrdersLinksSourceModifiedDateShipToAddressTypeEnum string

const (
	ListPurchaseOrdersLinksSourceModifiedDateShipToAddressTypeEnumUnknown  ListPurchaseOrdersLinksSourceModifiedDateShipToAddressTypeEnum = "Unknown"
	ListPurchaseOrdersLinksSourceModifiedDateShipToAddressTypeEnumBilling  ListPurchaseOrdersLinksSourceModifiedDateShipToAddressTypeEnum = "Billing"
	ListPurchaseOrdersLinksSourceModifiedDateShipToAddressTypeEnumDelivery ListPurchaseOrdersLinksSourceModifiedDateShipToAddressTypeEnum = "Delivery"
)

// ListPurchaseOrdersLinksSourceModifiedDateShipToAddress
// Delivery address for any goods that have been ordered.
type ListPurchaseOrdersLinksSourceModifiedDateShipToAddress struct {
	City       *string                                                        `json:"city,omitempty"`
	Country    *string                                                        `json:"country,omitempty"`
	Line1      *string                                                        `json:"line1,omitempty"`
	Line2      *string                                                        `json:"line2,omitempty"`
	PostalCode *string                                                        `json:"postalCode,omitempty"`
	Region     *string                                                        `json:"region,omitempty"`
	Type       ListPurchaseOrdersLinksSourceModifiedDateShipToAddressTypeEnum `json:"type"`
}

// ListPurchaseOrdersLinksSourceModifiedDateShipToContact
// Details of the named contact at the delivery address.
type ListPurchaseOrdersLinksSourceModifiedDateShipToContact struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// ListPurchaseOrdersLinksSourceModifiedDateShipTo
// Delivery details for any goods that have been ordered.
type ListPurchaseOrdersLinksSourceModifiedDateShipTo struct {
	Address *ListPurchaseOrdersLinksSourceModifiedDateShipToAddress `json:"address,omitempty"`
	Contact *ListPurchaseOrdersLinksSourceModifiedDateShipToContact `json:"contact,omitempty"`
}

type ListPurchaseOrdersLinksSourceModifiedDateStatusEnum string

const (
	ListPurchaseOrdersLinksSourceModifiedDateStatusEnumUnknown ListPurchaseOrdersLinksSourceModifiedDateStatusEnum = "Unknown"
	ListPurchaseOrdersLinksSourceModifiedDateStatusEnumDraft   ListPurchaseOrdersLinksSourceModifiedDateStatusEnum = "Draft"
	ListPurchaseOrdersLinksSourceModifiedDateStatusEnumOpen    ListPurchaseOrdersLinksSourceModifiedDateStatusEnum = "Open"
	ListPurchaseOrdersLinksSourceModifiedDateStatusEnumClosed  ListPurchaseOrdersLinksSourceModifiedDateStatusEnum = "Closed"
	ListPurchaseOrdersLinksSourceModifiedDateStatusEnumVoid    ListPurchaseOrdersLinksSourceModifiedDateStatusEnum = "Void"
)

// ListPurchaseOrdersLinksSourceModifiedDateSupplierRef
// Supplier that the purchase order is recorded against in the accounting system.
type ListPurchaseOrdersLinksSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

// ListPurchaseOrdersLinksSourceModifiedDate
// > View the coverage for purchase orders in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Purchase orders represent a business's intent to purchase goods or services from a supplier and normally include information such as expected delivery dates and shipping details.
//
// This information can be used to provide visibility on a business's expected payables and to track a purchase through the full procurement process.
type ListPurchaseOrdersLinksSourceModifiedDate struct {
	Currency             *string                                               `json:"currency,omitempty"`
	CurrencyRate         *float64                                              `json:"currencyRate,omitempty"`
	DeliveryDate         *time.Time                                            `json:"deliveryDate,omitempty"`
	ExpectedDeliveryDate *time.Time                                            `json:"expectedDeliveryDate,omitempty"`
	ID                   *string                                               `json:"id,omitempty"`
	IssueDate            *time.Time                                            `json:"issueDate,omitempty"`
	LineItems            []ListPurchaseOrdersLinksSourceModifiedDateLineItems  `json:"lineItems,omitempty"`
	Metadata             *ListPurchaseOrdersLinksSourceModifiedDateMetadata    `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                            `json:"modifiedDate,omitempty"`
	Note                 *string                                               `json:"note,omitempty"`
	PaymentDueDate       *time.Time                                            `json:"paymentDueDate,omitempty"`
	PurchaseOrderNumber  *string                                               `json:"purchaseOrderNumber,omitempty"`
	ShipTo               *ListPurchaseOrdersLinksSourceModifiedDateShipTo      `json:"shipTo,omitempty"`
	SourceModifiedDate   *time.Time                                            `json:"sourceModifiedDate,omitempty"`
	Status               *ListPurchaseOrdersLinksSourceModifiedDateStatusEnum  `json:"status,omitempty"`
	SubTotal             *float64                                              `json:"subTotal,omitempty"`
	SupplierRef          *ListPurchaseOrdersLinksSourceModifiedDateSupplierRef `json:"supplierRef,omitempty"`
	TotalAmount          *float64                                              `json:"totalAmount,omitempty"`
	TotalDiscount        *float64                                              `json:"totalDiscount,omitempty"`
	TotalTaxAmount       *float64                                              `json:"totalTaxAmount,omitempty"`
}

// ListPurchaseOrdersLinks
// Codat's Paging Model
type ListPurchaseOrdersLinks struct {
	Links        ListPurchaseOrdersLinksLinks                `json:"_links"`
	PageNumber   int64                                       `json:"pageNumber"`
	PageSize     int64                                       `json:"pageSize"`
	Results      []ListPurchaseOrdersLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                       `json:"totalResults"`
}

type ListPurchaseOrdersResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListPurchaseOrdersLinks
}
