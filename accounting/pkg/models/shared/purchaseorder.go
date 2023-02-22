package shared

import (
	"time"
)

type PurchaseOrderLineItems struct {
	AccountRef           *AccountRef `json:"accountRef,omitempty"`
	Description          *string     `json:"description,omitempty"`
	DiscountAmount       *float64    `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64    `json:"discountPercentage,omitempty"`
	ItemRef              *ItemRef    `json:"itemRef,omitempty"`
	Quantity             *float64    `json:"quantity,omitempty"`
	SubTotal             *float64    `json:"subTotal,omitempty"`
	TaxAmount            *float64    `json:"taxAmount,omitempty"`
	TaxRateRef           *TaxRateRef `json:"taxRateRef,omitempty"`
	TotalAmount          *float64    `json:"totalAmount,omitempty"`
	TrackingCategoryRefs []Items     `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           *float64    `json:"unitAmount,omitempty"`
}

type PurchaseOrderStatusEnum string

const (
	PurchaseOrderStatusEnumUnknown PurchaseOrderStatusEnum = "Unknown"
	PurchaseOrderStatusEnumDraft   PurchaseOrderStatusEnum = "Draft"
	PurchaseOrderStatusEnumOpen    PurchaseOrderStatusEnum = "Open"
	PurchaseOrderStatusEnumClosed  PurchaseOrderStatusEnum = "Closed"
	PurchaseOrderStatusEnumVoid    PurchaseOrderStatusEnum = "Void"
)

// PurchaseOrder
// > View the coverage for purchase orders in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Purchase orders represent a business's intent to purchase goods or services from a supplier and normally include information such as expected delivery dates and shipping details.
//
// This information can be used to provide visibility on a business's expected payables and to track a purchase through the full procurement process.
type PurchaseOrder struct {
	Currency             *string                  `json:"currency,omitempty"`
	CurrencyRate         *float64                 `json:"currencyRate,omitempty"`
	DeliveryDate         *time.Time               `json:"deliveryDate,omitempty"`
	ExpectedDeliveryDate *time.Time               `json:"expectedDeliveryDate,omitempty"`
	ID                   *string                  `json:"id,omitempty"`
	IssueDate            *time.Time               `json:"issueDate,omitempty"`
	LineItems            []PurchaseOrderLineItems `json:"lineItems,omitempty"`
	Metadata             *Metadata                `json:"metadata,omitempty"`
	ModifiedDate         *time.Time               `json:"modifiedDate,omitempty"`
	Note                 *string                  `json:"note,omitempty"`
	PaymentDueDate       *time.Time               `json:"paymentDueDate,omitempty"`
	PurchaseOrderNumber  *string                  `json:"purchaseOrderNumber,omitempty"`
	ShipTo               *ShipTo                  `json:"shipTo,omitempty"`
	SourceModifiedDate   *time.Time               `json:"sourceModifiedDate,omitempty"`
	Status               *PurchaseOrderStatusEnum `json:"status,omitempty"`
	SubTotal             *float64                 `json:"subTotal,omitempty"`
	SupplierRef          *SupplierRef             `json:"supplierRef,omitempty"`
	TotalAmount          *float64                 `json:"totalAmount,omitempty"`
	TotalDiscount        *float64                 `json:"totalDiscount,omitempty"`
	TotalTaxAmount       *float64                 `json:"totalTaxAmount,omitempty"`
}
