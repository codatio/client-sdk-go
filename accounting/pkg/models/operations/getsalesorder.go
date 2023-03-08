package operations

import (
	"net/http"
	"time"
)

type GetSalesOrderPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	SalesOrderID string `pathParam:"style=simple,explode=false,name=salesOrderId"`
}

type GetSalesOrderRequest struct {
	PathParams GetSalesOrderPathParams
}

// GetSalesOrderSourceModifiedDateCustomerRef
// The customer that the sales order is recorded against in the accounting system.
type GetSalesOrderSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type GetSalesOrderSourceModifiedDateInvoicingStatusEnum string

const (
	GetSalesOrderSourceModifiedDateInvoicingStatusEnumUnknown           GetSalesOrderSourceModifiedDateInvoicingStatusEnum = "Unknown"
	GetSalesOrderSourceModifiedDateInvoicingStatusEnumPartiallyInvoiced GetSalesOrderSourceModifiedDateInvoicingStatusEnum = "PartiallyInvoiced"
	GetSalesOrderSourceModifiedDateInvoicingStatusEnumInvoiced          GetSalesOrderSourceModifiedDateInvoicingStatusEnum = "Invoiced"
	GetSalesOrderSourceModifiedDateInvoicingStatusEnumNotInvoiced       GetSalesOrderSourceModifiedDateInvoicingStatusEnum = "NotInvoiced"
)

// GetSalesOrderSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type GetSalesOrderSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetSalesOrderSourceModifiedDateLineItemsItemRef
// Reference to the product or inventory item to which the line item is linked.
type GetSalesOrderSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// GetSalesOrderSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type GetSalesOrderSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type GetSalesOrderSourceModifiedDateLineItemsTracking struct {
	RecordRefs []string `json:"recordRefs,omitempty"`
}

type GetSalesOrderSourceModifiedDateLineItems struct {
	AccountRef         *GetSalesOrderSourceModifiedDateLineItemsAccountRef `json:"accountRef,omitempty"`
	Description        *string                                             `json:"description,omitempty"`
	DiscountAmount     *float64                                            `json:"discountAmount,omitempty"`
	DiscountPercentage *float64                                            `json:"discountPercentage,omitempty"`
	ItemRef            *GetSalesOrderSourceModifiedDateLineItemsItemRef    `json:"itemRef,omitempty"`
	Quantity           *float64                                            `json:"quantity,omitempty"`
	SubTotal           *float64                                            `json:"subTotal,omitempty"`
	TaxAmount          *float64                                            `json:"taxAmount,omitempty"`
	TaxRateRef         *GetSalesOrderSourceModifiedDateLineItemsTaxRateRef `json:"taxRateRef,omitempty"`
	TotalAmount        *float64                                            `json:"totalAmount,omitempty"`
	Tracking           *GetSalesOrderSourceModifiedDateLineItemsTracking   `json:"tracking,omitempty"`
	UnitAmount         *float64                                            `json:"unitAmount,omitempty"`
}

type GetSalesOrderSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetSalesOrderSourceModifiedDateShipToAddressTypeEnum string

const (
	GetSalesOrderSourceModifiedDateShipToAddressTypeEnumUnknown  GetSalesOrderSourceModifiedDateShipToAddressTypeEnum = "Unknown"
	GetSalesOrderSourceModifiedDateShipToAddressTypeEnumBilling  GetSalesOrderSourceModifiedDateShipToAddressTypeEnum = "Billing"
	GetSalesOrderSourceModifiedDateShipToAddressTypeEnumDelivery GetSalesOrderSourceModifiedDateShipToAddressTypeEnum = "Delivery"
)

// GetSalesOrderSourceModifiedDateShipToAddress
// Delivery address for any goods that have been ordered.
type GetSalesOrderSourceModifiedDateShipToAddress struct {
	City       *string                                              `json:"city,omitempty"`
	Country    *string                                              `json:"country,omitempty"`
	Line1      *string                                              `json:"line1,omitempty"`
	Line2      *string                                              `json:"line2,omitempty"`
	PostalCode *string                                              `json:"postalCode,omitempty"`
	Region     *string                                              `json:"region,omitempty"`
	Type       GetSalesOrderSourceModifiedDateShipToAddressTypeEnum `json:"type"`
}

// GetSalesOrderSourceModifiedDateShipToContact
// Details of the named contact at the delivery address.
type GetSalesOrderSourceModifiedDateShipToContact struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// GetSalesOrderSourceModifiedDateShipTo
// Delivery details for any goods that have been ordered.
type GetSalesOrderSourceModifiedDateShipTo struct {
	Address *GetSalesOrderSourceModifiedDateShipToAddress `json:"address,omitempty"`
	Contact *GetSalesOrderSourceModifiedDateShipToContact `json:"contact,omitempty"`
}

type GetSalesOrderSourceModifiedDateStatusEnum string

const (
	GetSalesOrderSourceModifiedDateStatusEnumUnknown GetSalesOrderSourceModifiedDateStatusEnum = "Unknown"
	GetSalesOrderSourceModifiedDateStatusEnumDraft   GetSalesOrderSourceModifiedDateStatusEnum = "Draft"
	GetSalesOrderSourceModifiedDateStatusEnumOpen    GetSalesOrderSourceModifiedDateStatusEnum = "Open"
	GetSalesOrderSourceModifiedDateStatusEnumClosed  GetSalesOrderSourceModifiedDateStatusEnum = "Closed"
	GetSalesOrderSourceModifiedDateStatusEnumVoid    GetSalesOrderSourceModifiedDateStatusEnum = "Void"
)

// GetSalesOrderSourceModifiedDate
// > View the coverage for sales orders in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=salesOrders" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A sales order represents a customer's intention to purchase goods or services from a seller and usually includes information such as the expected delivery date and shipping details. This information can be used to provide visibility on a business's expected receivables and to track sales through the full procurement process.
//
// A sales order is typically converted to an [invoice](https://docs.codat.io/accounting-api#/schemas/Invoice) after approval.
type GetSalesOrderSourceModifiedDate struct {
	Currency                    *string                                             `json:"currency,omitempty"`
	CurrencyRate                *float64                                            `json:"currencyRate,omitempty"`
	CustomerPurchaseOrderNumber *string                                             `json:"customerPurchaseOrderNumber,omitempty"`
	CustomerRef                 *GetSalesOrderSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	ExpectedDeliveryDate        *time.Time                                          `json:"expectedDeliveryDate,omitempty"`
	ID                          *string                                             `json:"id,omitempty"`
	InvoicingStatus             *GetSalesOrderSourceModifiedDateInvoicingStatusEnum `json:"invoicingStatus,omitempty"`
	IssueDate                   *time.Time                                          `json:"issueDate,omitempty"`
	LineItems                   []GetSalesOrderSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                    *GetSalesOrderSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate                *time.Time                                          `json:"modifiedDate,omitempty"`
	Note                        *string                                             `json:"note,omitempty"`
	SalesOrderNumber            *string                                             `json:"salesOrderNumber,omitempty"`
	ShipTo                      *GetSalesOrderSourceModifiedDateShipTo              `json:"shipTo,omitempty"`
	SourceModifiedDate          *time.Time                                          `json:"sourceModifiedDate,omitempty"`
	Status                      *GetSalesOrderSourceModifiedDateStatusEnum          `json:"status,omitempty"`
	SubTotal                    *float64                                            `json:"subTotal,omitempty"`
	TotalAmount                 *float64                                            `json:"totalAmount,omitempty"`
	TotalDiscount               *float64                                            `json:"totalDiscount,omitempty"`
	TotalTaxAmount              *float64                                            `json:"totalTaxAmount,omitempty"`
}

type GetSalesOrderResponse struct {
	ContentType        string
	SourceModifiedDate *GetSalesOrderSourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
