package shared

import (
	"time"
)

type SalesOrderInvoicingStatusEnum string

const (
	SalesOrderInvoicingStatusEnumUnknown           SalesOrderInvoicingStatusEnum = "Unknown"
	SalesOrderInvoicingStatusEnumPartiallyInvoiced SalesOrderInvoicingStatusEnum = "PartiallyInvoiced"
	SalesOrderInvoicingStatusEnumInvoiced          SalesOrderInvoicingStatusEnum = "Invoiced"
	SalesOrderInvoicingStatusEnumNotInvoiced       SalesOrderInvoicingStatusEnum = "NotInvoiced"
)

type SalesOrderLineItemsTracking struct {
	RecordRefs []string `json:"recordRefs,omitempty"`
}

type SalesOrderLineItems struct {
	AccountRef         *AccountRef                  `json:"accountRef,omitempty"`
	Description        *string                      `json:"description,omitempty"`
	DiscountAmount     *float64                     `json:"discountAmount,omitempty"`
	DiscountPercentage *float64                     `json:"discountPercentage,omitempty"`
	ItemRef            *ItemRef                     `json:"itemRef,omitempty"`
	Quantity           *float64                     `json:"quantity,omitempty"`
	SubTotal           *float64                     `json:"subTotal,omitempty"`
	TaxAmount          *float64                     `json:"taxAmount,omitempty"`
	TaxRateRef         *TaxRateRef                  `json:"taxRateRef,omitempty"`
	TotalAmount        *float64                     `json:"totalAmount,omitempty"`
	Tracking           *SalesOrderLineItemsTracking `json:"tracking,omitempty"`
	UnitAmount         *float64                     `json:"unitAmount,omitempty"`
}

// SalesOrderShipToContact
// Details of the named contact at the delivery address.
type SalesOrderShipToContact struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// SalesOrderShipTo
// Delivery details for any goods that have been ordered.
type SalesOrderShipTo struct {
	Address *Addressesitems          `json:"address,omitempty"`
	Contact *SalesOrderShipToContact `json:"contact,omitempty"`
}

type SalesOrderStatusEnum string

const (
	SalesOrderStatusEnumUnknown SalesOrderStatusEnum = "Unknown"
	SalesOrderStatusEnumDraft   SalesOrderStatusEnum = "Draft"
	SalesOrderStatusEnumOpen    SalesOrderStatusEnum = "Open"
	SalesOrderStatusEnumClosed  SalesOrderStatusEnum = "Closed"
	SalesOrderStatusEnumVoid    SalesOrderStatusEnum = "Void"
)

// SalesOrder
// > View the coverage for sales orders in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=salesOrders" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A sales order represents a customer's intention to purchase goods or services from a seller and usually includes information such as the expected delivery date and shipping details. This information can be used to provide visibility on a business's expected receivables and to track sales through the full procurement process.
//
// A sales order is typically converted to an [invoice](https://docs.codat.io/docs/datamodel-accounting-invoices) after approval.
type SalesOrder struct {
	Currency                    *string                        `json:"currency,omitempty"`
	CurrencyRate                *float64                       `json:"currencyRate,omitempty"`
	CustomerPurchaseOrderNumber *string                        `json:"customerPurchaseOrderNumber,omitempty"`
	CustomerRef                 *CustomerRef                   `json:"customerRef,omitempty"`
	ExpectedDeliveryDate        *time.Time                     `json:"expectedDeliveryDate,omitempty"`
	ID                          *string                        `json:"id,omitempty"`
	InvoicingStatus             *SalesOrderInvoicingStatusEnum `json:"invoicingStatus,omitempty"`
	IssueDate                   *time.Time                     `json:"issueDate,omitempty"`
	LineItems                   []SalesOrderLineItems          `json:"lineItems,omitempty"`
	Metadata                    *Metadata                      `json:"metadata,omitempty"`
	ModifiedDate                *time.Time                     `json:"modifiedDate,omitempty"`
	Note                        *string                        `json:"note,omitempty"`
	SalesOrderNumber            *string                        `json:"salesOrderNumber,omitempty"`
	ShipTo                      *SalesOrderShipTo              `json:"shipTo,omitempty"`
	SourceModifiedDate          *time.Time                     `json:"sourceModifiedDate,omitempty"`
	Status                      *SalesOrderStatusEnum          `json:"status,omitempty"`
	SubTotal                    *float64                       `json:"subTotal,omitempty"`
	TotalAmount                 *float64                       `json:"totalAmount,omitempty"`
	TotalDiscount               *float64                       `json:"totalDiscount,omitempty"`
	TotalTaxAmount              *float64                       `json:"totalTaxAmount,omitempty"`
}
