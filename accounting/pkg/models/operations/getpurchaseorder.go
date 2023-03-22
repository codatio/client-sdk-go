// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"time"
)

type GetPurchaseOrderRequest struct {
	CompanyID       string `pathParam:"style=simple,explode=false,name=companyId"`
	PurchaseOrderID string `pathParam:"style=simple,explode=false,name=purchaseOrderId"`
}

// GetPurchaseOrderSourceModifiedDateLineItemsAccountRef - Reference to the account to which the line item is linked.
type GetPurchaseOrderSourceModifiedDateLineItemsAccountRef struct {
	// 'id' from the Accounts data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the Accounts data type.
	Name *string `json:"name,omitempty"`
}

// GetPurchaseOrderSourceModifiedDateLineItemsItemRef - Reference to the product or inventory item to which the line item is linked.
type GetPurchaseOrderSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// GetPurchaseOrderSourceModifiedDateLineItemsTaxRateRef - Reference to the tax rate to which the line item is linked.
type GetPurchaseOrderSourceModifiedDateLineItemsTaxRateRef struct {
	// Applicable tax rate.
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	// 'id' from the 'taxRates' data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the 'taxRates' data type.
	Name *string `json:"name,omitempty"`
}

// GetPurchaseOrderSourceModifiedDateLineItemsTrackingCategoryRefs - References a category against which the item is tracked.
type GetPurchaseOrderSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type GetPurchaseOrderSourceModifiedDateLineItems struct {
	// Reference to the account to which the line item is linked.
	AccountRef *GetPurchaseOrderSourceModifiedDateLineItemsAccountRef `json:"accountRef,omitempty"`
	// Description of the goods / services that have been ordered.
	Description *string `json:"description,omitempty"`
	// Value of any discounts applied.
	DiscountAmount *float64 `json:"discountAmount,omitempty"`
	// Percentage rate (from 0 to 100) of any discounts applied to the unit amount.
	DiscountPercentage *float64 `json:"discountPercentage,omitempty"`
	// Reference to the product or inventory item to which the line item is linked.
	ItemRef *GetPurchaseOrderSourceModifiedDateLineItemsItemRef `json:"itemRef,omitempty"`
	// Number of units that have been ordered.
	Quantity *float64 `json:"quantity,omitempty"`
	// Amount of the line, inclusive of discounts but exclusive of tax.
	SubTotal *float64 `json:"subTotal,omitempty"`
	// Amount of tax for the line.
	TaxAmount *float64 `json:"taxAmount,omitempty"`
	// Reference to the tax rate to which the line item is linked.
	TaxRateRef *GetPurchaseOrderSourceModifiedDateLineItemsTaxRateRef `json:"taxRateRef,omitempty"`
	// Total amount of the line, inclusive of discounts and tax.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	// Reference to the tracking categories to which the line item is linked.
	TrackingCategoryRefs []GetPurchaseOrderSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	// Price of each unit.
	UnitAmount *float64 `json:"unitAmount,omitempty"`
}

type GetPurchaseOrderSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnum - Type of the address.
type GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnum string

const (
	GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnumUnknown  GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnum = "Unknown"
	GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnumBilling  GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnum = "Billing"
	GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnumDelivery GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnum = "Delivery"
)

// GetPurchaseOrderSourceModifiedDateShipToAddress - Delivery address for any goods that have been ordered.
type GetPurchaseOrderSourceModifiedDateShipToAddress struct {
	// City of the customer address.
	City *string `json:"city,omitempty"`
	// Country of the customer address.
	Country *string `json:"country,omitempty"`
	// Line 1 of the customer address.
	Line1 *string `json:"line1,omitempty"`
	// Line 2 of the customer address.
	Line2 *string `json:"line2,omitempty"`
	// Postal code or zip code.
	PostalCode *string `json:"postalCode,omitempty"`
	// Region of the customer address.
	Region *string `json:"region,omitempty"`
	// Type of the address.
	Type GetPurchaseOrderSourceModifiedDateShipToAddressTypeEnum `json:"type"`
}

// GetPurchaseOrderSourceModifiedDateShipToContact - Details of the named contact at the delivery address.
type GetPurchaseOrderSourceModifiedDateShipToContact struct {
	// Email address of the contact at the delivery address.
	Email *string `json:"email,omitempty"`
	// Name of the contact at the delivery address.
	Name *string `json:"name,omitempty"`
	// Phone number of the contact at the delivery address.
	Phone *string `json:"phone,omitempty"`
}

// GetPurchaseOrderSourceModifiedDateShipTo - Delivery details for any goods that have been ordered.
type GetPurchaseOrderSourceModifiedDateShipTo struct {
	// Delivery address for any goods that have been ordered.
	Address *GetPurchaseOrderSourceModifiedDateShipToAddress `json:"address,omitempty"`
	// Details of the named contact at the delivery address.
	Contact *GetPurchaseOrderSourceModifiedDateShipToContact `json:"contact,omitempty"`
}

// GetPurchaseOrderSourceModifiedDateStatusEnum - Current state of the purchase order
type GetPurchaseOrderSourceModifiedDateStatusEnum string

const (
	GetPurchaseOrderSourceModifiedDateStatusEnumUnknown GetPurchaseOrderSourceModifiedDateStatusEnum = "Unknown"
	GetPurchaseOrderSourceModifiedDateStatusEnumDraft   GetPurchaseOrderSourceModifiedDateStatusEnum = "Draft"
	GetPurchaseOrderSourceModifiedDateStatusEnumOpen    GetPurchaseOrderSourceModifiedDateStatusEnum = "Open"
	GetPurchaseOrderSourceModifiedDateStatusEnumClosed  GetPurchaseOrderSourceModifiedDateStatusEnum = "Closed"
	GetPurchaseOrderSourceModifiedDateStatusEnumVoid    GetPurchaseOrderSourceModifiedDateStatusEnum = "Void"
)

// GetPurchaseOrderSourceModifiedDateSupplierRef - Supplier that the purchase order is recorded against in the accounting system.
type GetPurchaseOrderSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

// GetPurchaseOrderSourceModifiedDate - > View the coverage for purchase orders in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Purchase orders represent a business's intent to purchase goods or services from a supplier and normally include information such as expected delivery dates and shipping details.
//
// This information can be used to provide visibility on a business's expected payables and to track a purchase through the full procurement process.
type GetPurchaseOrderSourceModifiedDate struct {
	// Currency of the purchase order.
	Currency *string `json:"currency,omitempty"`
	// Rate to convert the total amount of the payment into the base currency for the company at the time of the payment.
	//
	// Currency rates in Codat are implemented as the multiple of foreign currency units to each base currency unit.
	//
	// Where the currency rate is provided by the underlying accounting platform, it will be available from Codat with the same precision (up to a maximum of 9 decimal places).
	//
	// For accounting platforms which do not provide an explicit currency rate, it is calculated as `baseCurrency / foreignCurrency` and will be returned to 9 decimal places.
	//
	// ## Examples with base currency of GBP
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (GBP) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **USD**          | $20            | 0.781         | £15.62                     |
	// | **EUR**          | €20            | 0.885         | £17.70                     |
	// | **RUB**          | ₽20            | 0.011         | £0.22                      |
	//
	// ## Examples with base currency of USD
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (USD) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **GBP**          | £20            | 1.277         | $25.54                     |
	// | **EUR**          | €20            | 1.134         | $22.68                     |
	// | **RUB**          | ₽20            | 0.015         | $0.30                      |
	CurrencyRate *float64 `json:"currencyRate,omitempty"`
	// Actual delivery date for any goods that have been ordered.
	DeliveryDate *time.Time `json:"deliveryDate,omitempty"`
	// Expected delivery date for any goods that have been ordered.
	ExpectedDeliveryDate *time.Time `json:"expectedDeliveryDate,omitempty"`
	// Identifier for the purchase order, unique for the company in the accounting platform.
	ID *string `json:"id,omitempty"`
	// Date of the purchase order as recorded in the accounting platform.
	IssueDate *time.Time `json:"issueDate,omitempty"`
	// Array of line items.
	LineItems []GetPurchaseOrderSourceModifiedDateLineItems `json:"lineItems,omitempty"`
	Metadata  *GetPurchaseOrderSourceModifiedDateMetadata   `json:"metadata,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	// Any additional information associated with the purchase order.
	Note *string `json:"note,omitempty"`
	// Date the supplier is due to be paid.
	PaymentDueDate *time.Time `json:"paymentDueDate,omitempty"`
	// Friendly reference for the purchase order, commonly generated by the accounting platform.
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
	// Delivery details for any goods that have been ordered.
	ShipTo *GetPurchaseOrderSourceModifiedDateShipTo `json:"shipTo,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
	// Current state of the purchase order
	Status *GetPurchaseOrderSourceModifiedDateStatusEnum `json:"status,omitempty"`
	// Total amount of the purchase order, including discounts but excluding tax.
	SubTotal *float64 `json:"subTotal,omitempty"`
	// Supplier that the purchase order is recorded against in the accounting system.
	SupplierRef *GetPurchaseOrderSourceModifiedDateSupplierRef `json:"supplierRef,omitempty"`
	// Total amount of the purchase order, including discounts and tax.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	// Total value of any discounts applied to the purchase order.
	TotalDiscount *float64 `json:"totalDiscount,omitempty"`
	//
	// Total amount of tax included in the purchase order.
	TotalTaxAmount *float64 `json:"totalTaxAmount,omitempty"`
}

type GetPurchaseOrderResponse struct {
	ContentType string
	// Success
	SourceModifiedDate *GetPurchaseOrderSourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
