package operations

import (
	"net/http"
	"time"
)

type GetInvoiceRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	InvoiceID string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

// GetInvoiceSourceModifiedDateCustomerRef
// Reference to the customer the invoice has been issued to.
type GetInvoiceSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

// GetInvoiceSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type GetInvoiceSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetInvoiceSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type GetInvoiceSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// GetInvoiceSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type GetInvoiceSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// GetInvoiceSourceModifiedDateLineItemsTrackingCategoryRefs
// References a category against which the item is tracked.
type GetInvoiceSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type GetInvoiceSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type GetInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	GetInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       GetInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	GetInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable GetInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	GetInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       GetInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type GetInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	GetInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       GetInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	GetInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable GetInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	GetInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       GetInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type GetInvoiceSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// GetInvoiceSourceModifiedDateLineItemsTracking
// Categories, and a project and customer, against which the item is tracked.
type GetInvoiceSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []GetInvoiceSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *GetInvoiceSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   GetInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo GetInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *GetInvoiceSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type GetInvoiceSourceModifiedDateLineItems struct {
	AccountRef           *GetInvoiceSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                     `json:"description,omitempty"`
	DiscountAmount       *float64                                                    `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                    `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                                                       `json:"isDirectIncome,omitempty"`
	ItemRef              *GetInvoiceSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                     `json:"quantity"`
	SubTotal             *float64                                                    `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                    `json:"taxAmount,omitempty"`
	TaxRateRef           *GetInvoiceSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                    `json:"totalAmount,omitempty"`
	Tracking             *GetInvoiceSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []GetInvoiceSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                     `json:"unitAmount"`
}

type GetInvoiceSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetInvoiceSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// GetInvoiceSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type GetInvoiceSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetInvoiceSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *GetInvoiceSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                          `json:"currency,omitempty"`
	CurrencyRate *float64                                                         `json:"currencyRate,omitempty"`
	ID           *string                                                          `json:"id,omitempty"`
	Note         *string                                                          `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                       `json:"paidOnDate,omitempty"`
	Reference    *string                                                          `json:"reference,omitempty"`
	TotalAmount  *float64                                                         `json:"totalAmount,omitempty"`
}

type GetInvoiceSourceModifiedDatePaymentAllocations struct {
	Allocation GetInvoiceSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    GetInvoiceSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type GetInvoiceSourceModifiedDateStatusEnum string

const (
	GetInvoiceSourceModifiedDateStatusEnumUnknown       GetInvoiceSourceModifiedDateStatusEnum = "Unknown"
	GetInvoiceSourceModifiedDateStatusEnumDraft         GetInvoiceSourceModifiedDateStatusEnum = "Draft"
	GetInvoiceSourceModifiedDateStatusEnumSubmitted     GetInvoiceSourceModifiedDateStatusEnum = "Submitted"
	GetInvoiceSourceModifiedDateStatusEnumPartiallyPaid GetInvoiceSourceModifiedDateStatusEnum = "PartiallyPaid"
	GetInvoiceSourceModifiedDateStatusEnumPaid          GetInvoiceSourceModifiedDateStatusEnum = "Paid"
	GetInvoiceSourceModifiedDateStatusEnumVoid          GetInvoiceSourceModifiedDateStatusEnum = "Void"
)

// GetInvoiceSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type GetInvoiceSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type GetInvoiceSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// GetInvoiceSourceModifiedDate
// > View the coverage for invoices in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// An invoice is an itemized record of goods sold or services provided to a [customer](https://docs.codat.io/accounting-api#/schemas/Customer).
//
// In Codat, an invoice contains details of:
//
// - The timeline of the invoice—when it was raised, marked as paid, last edited, and so on.
// - How much the invoice is for, what portion of the invoice is tax or discounts, and what currency the amounts are represented in.
// - Who the invoice has been raised to; the _customer_.
// - The breakdown of what the invoice is for; the _line items_.
// - Any [payments](https://docs.codat.io/accounting-api#/schemas/Payment) assigned to the invoice; the _payment allocations_.
//
// > **Invoices or bills?**
// >
// > In Codat, invoices represent accounts receivable only. For accounts payable invoices, see [Bills](https://docs.codat.io/accounting-api#/schemas/Bill).
//
// > **Invoice PDF downloads**
// >
// > You can <a className="external" href="https://api.codat.io/swagger/index.html#/Invoices/get_companies__companyId__data_invoices__invoiceId__pdf" target="_blank">download a PDF version</a> of an invoice for supported integrations.
// >
// > The filename will be invoice-{number}.pdf.
type GetInvoiceSourceModifiedDate struct {
	AdditionalTaxAmount     *float64                                         `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                                         `json:"additionalTaxPercentage,omitempty"`
	AmountDue               float64                                          `json:"amountDue"`
	Currency                *string                                          `json:"currency,omitempty"`
	CurrencyRate            *float64                                         `json:"currencyRate,omitempty"`
	CustomerRef             *GetInvoiceSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	DiscountPercentage      *float64                                         `json:"discountPercentage,omitempty"`
	DueDate                 *time.Time                                       `json:"dueDate,omitempty"`
	ID                      *string                                          `json:"id,omitempty"`
	InvoiceNumber           *string                                          `json:"invoiceNumber,omitempty"`
	IssueDate               time.Time                                        `json:"issueDate"`
	LineItems               []GetInvoiceSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                *GetInvoiceSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                                       `json:"modifiedDate,omitempty"`
	Note                    *string                                          `json:"note,omitempty"`
	PaidOnDate              *time.Time                                       `json:"paidOnDate,omitempty"`
	PaymentAllocations      []GetInvoiceSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	SalesOrderRefs          []string                                         `json:"salesOrderRefs,omitempty"`
	SourceModifiedDate      *time.Time                                       `json:"sourceModifiedDate,omitempty"`
	Status                  GetInvoiceSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal                *float64                                         `json:"subTotal,omitempty"`
	SupplementalData        *GetInvoiceSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TotalAmount             float64                                          `json:"totalAmount"`
	TotalDiscount           *float64                                         `json:"totalDiscount,omitempty"`
	TotalTaxAmount          float64                                          `json:"totalTaxAmount"`
	WithholdingTax          []GetInvoiceSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type GetInvoiceResponse struct {
	ContentType        string
	SourceModifiedDate *GetInvoiceSourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
