package shared

import (
	"time"
)

type InvoiceLineItemsTracking struct {
	CategoryRefs []Items         `json:"categoryRefs"`
	CustomerRef  *CustomerRef    `json:"customerRef,omitempty"`
	IsBilledTo   IsBilledToEnum1 `json:"isBilledTo"`
	IsRebilledTo IsBilledToEnum1 `json:"isRebilledTo"`
	ProjectRef   *ProjectRef     `json:"projectRef,omitempty"`
}

type InvoiceLineItems struct {
	AccountRef           *AccountRef               `json:"accountRef,omitempty"`
	Description          *string                   `json:"description,omitempty"`
	DiscountAmount       *float64                  `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                  `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                     `json:"isDirectIncome,omitempty"`
	ItemRef              *ItemRef                  `json:"itemRef,omitempty"`
	Quantity             float64                   `json:"quantity"`
	SubTotal             *float64                  `json:"subTotal,omitempty"`
	TaxAmount            *float64                  `json:"taxAmount,omitempty"`
	TaxRateRef           *TaxRateRef               `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                  `json:"totalAmount,omitempty"`
	Tracking             *InvoiceLineItemsTracking `json:"tracking,omitempty"`
	TrackingCategoryRefs []Items                   `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                   `json:"unitAmount"`
}

type InvoiceStatusEnum string

const (
	InvoiceStatusEnumUnknown       InvoiceStatusEnum = "Unknown"
	InvoiceStatusEnumDraft         InvoiceStatusEnum = "Draft"
	InvoiceStatusEnumSubmitted     InvoiceStatusEnum = "Submitted"
	InvoiceStatusEnumPartiallyPaid InvoiceStatusEnum = "PartiallyPaid"
	InvoiceStatusEnumPaid          InvoiceStatusEnum = "Paid"
	InvoiceStatusEnumVoid          InvoiceStatusEnum = "Void"
)

// Invoice
// > View the coverage for invoices in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// An invoice is an itemized record of goods sold or services provided to a [customer](https://docs.codat.io/docs/datamodel-accounting-customers).
//
// In Codat, an invoice contains details of:
//
// - The timeline of the invoice—when it was raised, marked as paid, last edited, and so on.
// - How much the invoice is for, what portion of the invoice is tax or discounts, and what currency the amounts are represented in.
// - Who the invoice has been raised to; the _customer_.
// - The breakdown of what the invoice is for; the _line items_.
// - Any [payments](https://docs.codat.io/docs/datamodel-accounting-payments) assigned to the invoice; the _payment allocations_.
//
// > **Invoices or bills?**
// >
// > In Codat, invoices represent accounts receivable only. For accounts payable invoices, see [Bills](https://docs.codat.io/docs/datamodel-accounting-bills).
//
// > **Invoice PDF downloads**
// >
// > You can <a className="external" href="https://api.codat.io/swagger/index.html#/Invoices/get_companies__companyId__data_invoices__invoiceId__pdf" target="_blank">download a PDF version</a> of an invoice for supported integrations.
// >
// > The filename will be invoice-{number}.pdf.
type Invoice struct {
	AdditionalTaxAmount     *float64                  `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                  `json:"additionalTaxPercentage,omitempty"`
	AmountDue               float64                   `json:"amountDue"`
	Currency                *string                   `json:"currency,omitempty"`
	CurrencyRate            *float64                  `json:"currencyRate,omitempty"`
	CustomerRef             *CustomerRef              `json:"customerRef,omitempty"`
	DiscountPercentage      *float64                  `json:"discountPercentage,omitempty"`
	DueDate                 *time.Time                `json:"dueDate,omitempty"`
	ID                      *string                   `json:"id,omitempty"`
	InvoiceNumber           *string                   `json:"invoiceNumber,omitempty"`
	IssueDate               time.Time                 `json:"issueDate"`
	LineItems               []InvoiceLineItems        `json:"lineItems,omitempty"`
	Metadata                *Metadata                 `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                `json:"modifiedDate,omitempty"`
	Note                    *string                   `json:"note,omitempty"`
	PaidOnDate              *time.Time                `json:"paidOnDate,omitempty"`
	PaymentAllocations      []PaymentAllocationsitems `json:"paymentAllocations,omitempty"`
	SalesOrderRefs          []string                  `json:"salesOrderRefs,omitempty"`
	SourceModifiedDate      *time.Time                `json:"sourceModifiedDate,omitempty"`
	Status                  InvoiceStatusEnum         `json:"status"`
	SubTotal                *float64                  `json:"subTotal,omitempty"`
	SupplementalData        *SupplementalData         `json:"supplementalData,omitempty"`
	TotalAmount             float64                   `json:"totalAmount"`
	TotalDiscount           *float64                  `json:"totalDiscount,omitempty"`
	TotalTaxAmount          float64                   `json:"totalTaxAmount"`
	WithholdingTax          []WithholdingTaxitems     `json:"withholdingTax,omitempty"`
}
