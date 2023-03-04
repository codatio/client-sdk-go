package operations

import (
	"net/http"
	"time"
)

type GetPurchaseOrderPathParams struct {
	CompanyID       string `pathParam:"style=simple,explode=false,name=companyId"`
	PurchaseOrderID string `pathParam:"style=simple,explode=false,name=purchaseOrderId"`
}

type GetPurchaseOrderRequest struct {
	PathParams GetPurchaseOrderPathParams
}

// GetPurchaseOrderSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type GetPurchaseOrderSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetPurchaseOrderSourceModifiedDateLineItemsItemRef
// Reference to the product or inventory item to which the line item is linked.
type GetPurchaseOrderSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// GetPurchaseOrderSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type GetPurchaseOrderSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type GetPurchaseOrderSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type GetPurchaseOrderSourceModifiedDateLineItems struct {
	AccountRef           *GetPurchaseOrderSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                           `json:"description,omitempty"`
	DiscountAmount       *float64                                                          `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                          `json:"discountPercentage,omitempty"`
	ItemRef              *GetPurchaseOrderSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             *float64                                                          `json:"quantity,omitempty"`
	SubTotal             *float64                                                          `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                          `json:"taxAmount,omitempty"`
	TaxRateRef           *GetPurchaseOrderSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                          `json:"totalAmount,omitempty"`
	TrackingCategoryRefs []GetPurchaseOrderSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           *float64                                                          `json:"unitAmount,omitempty"`
}

type GetPurchaseOrderSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnum string

const (
	GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnumUnknown  GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnum = "Unknown"
	GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnumBilling  GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnum = "Billing"
	GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnumDelivery GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnum = "Delivery"
)

// GetPurchaseOrderSourceModifiedDateShipToAddress
// Delivery address for any goods that have been ordered.
type GetPurchaseOrderSourceModifiedDateShipToAddress struct {
	City       *string                                                 `json:"city,omitempty"`
	Country    *string                                                 `json:"country,omitempty"`
	Line1      *string                                                 `json:"line1,omitempty"`
	Line2      *string                                                 `json:"line2,omitempty"`
	PostalCode *string                                                 `json:"postalCode,omitempty"`
	Region     *string                                                 `json:"region,omitempty"`
	Type       GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnum `json:"type"`
}

// GetPurchaseOrderSourceModifiedDateShipToContact
// Details of the named contact at the delivery address.
type GetPurchaseOrderSourceModifiedDateShipToContact struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// GetPurchaseOrderSourceModifiedDateShipTo
// Delivery details for any goods that have been ordered.
type GetPurchaseOrderSourceModifiedDateShipTo struct {
	Address *GetPurchaseOrderSourceModifiedDateShipToAddress `json:"address,omitempty"`
	Contact *GetPurchaseOrderSourceModifiedDateShipToContact `json:"contact,omitempty"`
}

type GetPurchaseOrderSourceModifiedDateStatusEnum string

const (
	GetPurchaseOrderSourceModifiedDateStatusEnumUnknown GetPurchaseOrderSourceModifiedDateStatusEnum = "Unknown"
	GetPurchaseOrderSourceModifiedDateStatusEnumDraft   GetPurchaseOrderSourceModifiedDateStatusEnum = "Draft"
	GetPurchaseOrderSourceModifiedDateStatusEnumOpen    GetPurchaseOrderSourceModifiedDateStatusEnum = "Open"
	GetPurchaseOrderSourceModifiedDateStatusEnumClosed  GetPurchaseOrderSourceModifiedDateStatusEnum = "Closed"
	GetPurchaseOrderSourceModifiedDateStatusEnumVoid    GetPurchaseOrderSourceModifiedDateStatusEnum = "Void"
)

// GetPurchaseOrderSourceModifiedDateSupplierRef
// Supplier that the purchase order is recorded against in the accounting system.
type GetPurchaseOrderSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

// GetPurchaseOrderSourceModifiedDate
// > View the coverage for purchase orders in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Purchase orders represent a business's intent to purchase goods or services from a supplier and normally include information such as expected delivery dates and shipping details.
//
// This information can be used to provide visibility on a business's expected payables and to track a purchase through the full procurement process.
type GetPurchaseOrderSourceModifiedDate struct {
	Currency             *string                                        `json:"currency,omitempty"`
	CurrencyRate         *float64                                       `json:"currencyRate,omitempty"`
	DeliveryDate         *time.Time                                     `json:"deliveryDate,omitempty"`
	ExpectedDeliveryDate *time.Time                                     `json:"expectedDeliveryDate,omitempty"`
	ID                   *string                                        `json:"id,omitempty"`
	IssueDate            *time.Time                                     `json:"issueDate,omitempty"`
	LineItems            []GetPurchaseOrderSourceModifiedDateLineItems  `json:"lineItems,omitempty"`
	Metadata             *GetPurchaseOrderSourceModifiedDateMetadata    `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                     `json:"modifiedDate,omitempty"`
	Note                 *string                                        `json:"note,omitempty"`
	PaymentDueDate       *time.Time                                     `json:"paymentDueDate,omitempty"`
	PurchaseOrderNumber  *string                                        `json:"purchaseOrderNumber,omitempty"`
	ShipTo               *GetPurchaseOrderSourceModifiedDateShipTo      `json:"shipTo,omitempty"`
	SourceModifiedDate   *time.Time                                     `json:"sourceModifiedDate,omitempty"`
	Status               *GetPurchaseOrderSourceModifiedDateStatusEnum  `json:"status,omitempty"`
	SubTotal             *float64                                       `json:"subTotal,omitempty"`
	SupplierRef          *GetPurchaseOrderSourceModifiedDateSupplierRef `json:"supplierRef,omitempty"`
	TotalAmount          *float64                                       `json:"totalAmount,omitempty"`
	TotalDiscount        *float64                                       `json:"totalDiscount,omitempty"`
	TotalTaxAmount       *float64                                       `json:"totalTaxAmount,omitempty"`
}

type GetPurchaseOrderResponse struct {
	ContentType        string
	SourceModifiedDate *GetPurchaseOrderSourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
