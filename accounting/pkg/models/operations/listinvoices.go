package operations

import (
	"net/http"
	"time"
)

type ListInvoicesPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListInvoicesQueryParams struct {
	OrderBy  *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page     int     `queryParam:"style=form,explode=true,name=page"`
	PageSize *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string `queryParam:"style=form,explode=true,name=query"`
}

type ListInvoicesRequest struct {
	PathParams  ListInvoicesPathParams
	QueryParams ListInvoicesQueryParams
}

type ListInvoicesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListInvoicesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListInvoicesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListInvoicesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListInvoicesLinksLinks struct {
	Current  ListInvoicesLinksLinksCurrent   `json:"current"`
	Next     *ListInvoicesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListInvoicesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListInvoicesLinksLinksSelf      `json:"self"`
}

// ListInvoicesLinksSourceModifiedDateCustomerRef
// Reference to the customer the invoice has been issued to.
type ListInvoicesLinksSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

// ListInvoicesLinksSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type ListInvoicesLinksSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ListInvoicesLinksSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type ListInvoicesLinksSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// ListInvoicesLinksSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type ListInvoicesLinksSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// ListInvoicesLinksSourceModifiedDateLineItemsTrackingCategoryRefs
// References a category against which the item is tracked.
type ListInvoicesLinksSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type ListInvoicesLinksSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type ListInvoicesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	ListInvoicesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       ListInvoicesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	ListInvoicesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable ListInvoicesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	ListInvoicesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       ListInvoicesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type ListInvoicesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	ListInvoicesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       ListInvoicesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	ListInvoicesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable ListInvoicesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	ListInvoicesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       ListInvoicesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type ListInvoicesLinksSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// ListInvoicesLinksSourceModifiedDateLineItemsTracking
// Categories, and a project and customer, against which the item is tracked.
type ListInvoicesLinksSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []ListInvoicesLinksSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *ListInvoicesLinksSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   ListInvoicesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo ListInvoicesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *ListInvoicesLinksSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type ListInvoicesLinksSourceModifiedDateLineItems struct {
	AccountRef           *ListInvoicesLinksSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                            `json:"description,omitempty"`
	DiscountAmount       *float64                                                           `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                           `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                                                              `json:"isDirectIncome,omitempty"`
	ItemRef              *ListInvoicesLinksSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                            `json:"quantity"`
	SubTotal             *float64                                                           `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                           `json:"taxAmount,omitempty"`
	TaxRateRef           *ListInvoicesLinksSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                           `json:"totalAmount,omitempty"`
	Tracking             *ListInvoicesLinksSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []ListInvoicesLinksSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                            `json:"unitAmount"`
}

type ListInvoicesLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type ListInvoicesLinksSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// ListInvoicesLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type ListInvoicesLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ListInvoicesLinksSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *ListInvoicesLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                 `json:"currency,omitempty"`
	CurrencyRate *float64                                                                `json:"currencyRate,omitempty"`
	ID           *string                                                                 `json:"id,omitempty"`
	Note         *string                                                                 `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                              `json:"paidOnDate,omitempty"`
	Reference    *string                                                                 `json:"reference,omitempty"`
	TotalAmount  *float64                                                                `json:"totalAmount,omitempty"`
}

type ListInvoicesLinksSourceModifiedDatePaymentAllocations struct {
	Allocation ListInvoicesLinksSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    ListInvoicesLinksSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type ListInvoicesLinksSourceModifiedDateStatusEnum string

const (
	ListInvoicesLinksSourceModifiedDateStatusEnumUnknown       ListInvoicesLinksSourceModifiedDateStatusEnum = "Unknown"
	ListInvoicesLinksSourceModifiedDateStatusEnumDraft         ListInvoicesLinksSourceModifiedDateStatusEnum = "Draft"
	ListInvoicesLinksSourceModifiedDateStatusEnumSubmitted     ListInvoicesLinksSourceModifiedDateStatusEnum = "Submitted"
	ListInvoicesLinksSourceModifiedDateStatusEnumPartiallyPaid ListInvoicesLinksSourceModifiedDateStatusEnum = "PartiallyPaid"
	ListInvoicesLinksSourceModifiedDateStatusEnumPaid          ListInvoicesLinksSourceModifiedDateStatusEnum = "Paid"
	ListInvoicesLinksSourceModifiedDateStatusEnumVoid          ListInvoicesLinksSourceModifiedDateStatusEnum = "Void"
)

// ListInvoicesLinksSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type ListInvoicesLinksSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type ListInvoicesLinksSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// ListInvoicesLinksSourceModifiedDate
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
type ListInvoicesLinksSourceModifiedDate struct {
	AdditionalTaxAmount     *float64                                                `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                                                `json:"additionalTaxPercentage,omitempty"`
	AmountDue               float64                                                 `json:"amountDue"`
	Currency                *string                                                 `json:"currency,omitempty"`
	CurrencyRate            *float64                                                `json:"currencyRate,omitempty"`
	CustomerRef             *ListInvoicesLinksSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	DiscountPercentage      *float64                                                `json:"discountPercentage,omitempty"`
	DueDate                 *time.Time                                              `json:"dueDate,omitempty"`
	ID                      *string                                                 `json:"id,omitempty"`
	InvoiceNumber           *string                                                 `json:"invoiceNumber,omitempty"`
	IssueDate               time.Time                                               `json:"issueDate"`
	LineItems               []ListInvoicesLinksSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                *ListInvoicesLinksSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                                              `json:"modifiedDate,omitempty"`
	Note                    *string                                                 `json:"note,omitempty"`
	PaidOnDate              *time.Time                                              `json:"paidOnDate,omitempty"`
	PaymentAllocations      []ListInvoicesLinksSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	SalesOrderRefs          []string                                                `json:"salesOrderRefs,omitempty"`
	SourceModifiedDate      *time.Time                                              `json:"sourceModifiedDate,omitempty"`
	Status                  ListInvoicesLinksSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal                *float64                                                `json:"subTotal,omitempty"`
	SupplementalData        *ListInvoicesLinksSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TotalAmount             float64                                                 `json:"totalAmount"`
	TotalDiscount           *float64                                                `json:"totalDiscount,omitempty"`
	TotalTaxAmount          float64                                                 `json:"totalTaxAmount"`
	WithholdingTax          []ListInvoicesLinksSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

// ListInvoicesLinks
// Codat's Paging Model
type ListInvoicesLinks struct {
	Links        ListInvoicesLinksLinks                `json:"_links"`
	PageNumber   int64                                 `json:"pageNumber"`
	PageSize     int64                                 `json:"pageSize"`
	Results      []ListInvoicesLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                 `json:"totalResults"`
}

type ListInvoicesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListInvoicesLinks
}
