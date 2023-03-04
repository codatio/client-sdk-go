package operations

import (
	"net/http"
	"time"
)

type ListCreditNotesPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListCreditNotesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCreditNotesRequest struct {
	PathParams  ListCreditNotesPathParams
	QueryParams ListCreditNotesQueryParams
}

type ListCreditNotesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCreditNotesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCreditNotesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCreditNotesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCreditNotesLinksLinks struct {
	Current  ListCreditNotesLinksLinksCurrent   `json:"current"`
	Next     *ListCreditNotesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCreditNotesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCreditNotesLinksLinksSelf      `json:"self"`
}

// ListCreditNotesLinksSourceModifiedDateCustomerRef
// Reference to the customer the credit note has been issued to.
type ListCreditNotesLinksSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

// ListCreditNotesLinksSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type ListCreditNotesLinksSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ListCreditNotesLinksSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type ListCreditNotesLinksSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// ListCreditNotesLinksSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type ListCreditNotesLinksSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type ListCreditNotesLinksSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type ListCreditNotesLinksSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type ListCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	ListCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       ListCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	ListCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable ListCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	ListCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       ListCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type ListCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	ListCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       ListCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	ListCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable ListCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	ListCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       ListCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type ListCreditNotesLinksSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type ListCreditNotesLinksSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []ListCreditNotesLinksSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *ListCreditNotesLinksSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   ListCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo ListCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *ListCreditNotesLinksSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type ListCreditNotesLinksSourceModifiedDateLineItems struct {
	AccountRef           *ListCreditNotesLinksSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                               `json:"description,omitempty"`
	DiscountAmount       *float64                                                              `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                              `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                                                                 `json:"isDirectIncome,omitempty"`
	ItemRef              *ListCreditNotesLinksSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                               `json:"quantity"`
	SubTotal             *float64                                                              `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                              `json:"taxAmount,omitempty"`
	TaxRateRef           *ListCreditNotesLinksSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                              `json:"totalAmount,omitempty"`
	Tracking             *ListCreditNotesLinksSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []ListCreditNotesLinksSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                               `json:"unitAmount"`
}

type ListCreditNotesLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type ListCreditNotesLinksSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// ListCreditNotesLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type ListCreditNotesLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ListCreditNotesLinksSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *ListCreditNotesLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                    `json:"currency,omitempty"`
	CurrencyRate *float64                                                                   `json:"currencyRate,omitempty"`
	ID           *string                                                                    `json:"id,omitempty"`
	Note         *string                                                                    `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                 `json:"paidOnDate,omitempty"`
	Reference    *string                                                                    `json:"reference,omitempty"`
	TotalAmount  *float64                                                                   `json:"totalAmount,omitempty"`
}

type ListCreditNotesLinksSourceModifiedDatePaymentAllocations struct {
	Allocation ListCreditNotesLinksSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    ListCreditNotesLinksSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type ListCreditNotesLinksSourceModifiedDateStatusEnum string

const (
	ListCreditNotesLinksSourceModifiedDateStatusEnumUnknown       ListCreditNotesLinksSourceModifiedDateStatusEnum = "Unknown"
	ListCreditNotesLinksSourceModifiedDateStatusEnumDraft         ListCreditNotesLinksSourceModifiedDateStatusEnum = "Draft"
	ListCreditNotesLinksSourceModifiedDateStatusEnumSubmitted     ListCreditNotesLinksSourceModifiedDateStatusEnum = "Submitted"
	ListCreditNotesLinksSourceModifiedDateStatusEnumPaid          ListCreditNotesLinksSourceModifiedDateStatusEnum = "Paid"
	ListCreditNotesLinksSourceModifiedDateStatusEnumVoid          ListCreditNotesLinksSourceModifiedDateStatusEnum = "Void"
	ListCreditNotesLinksSourceModifiedDateStatusEnumPartiallyPaid ListCreditNotesLinksSourceModifiedDateStatusEnum = "PartiallyPaid"
)

type ListCreditNotesLinksSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type ListCreditNotesLinksSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// ListCreditNotesLinksSourceModifiedDate
// > View the coverage for credit notes in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Think of a credit note as a voucher issued to a customer. It is a reduction that can be applied against one or multiple invoices. A credit note can either reduce the amount owed or cancel out an invoice entirely.
//
// In the Codat system a credit note is issued to a [customer's](https://docs.codat.io/accounting-api#/schemas/Customer) accounts receivable.
//
// It contains details of:
// * The amount of credit remaining and its status.
// * Payment allocations against the payments type, in this case an invoice.
// * Which customers the credit notes have been issued to.
type ListCreditNotesLinksSourceModifiedDate struct {
	AdditionalTaxAmount     *float64                                                   `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                                                   `json:"additionalTaxPercentage,omitempty"`
	AllocatedOnDate         *time.Time                                                 `json:"allocatedOnDate,omitempty"`
	CreditNoteNumber        *string                                                    `json:"creditNoteNumber,omitempty"`
	Currency                *string                                                    `json:"currency,omitempty"`
	CurrencyRate            *float64                                                   `json:"currencyRate,omitempty"`
	CustomerRef             *ListCreditNotesLinksSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	DiscountPercentage      float64                                                    `json:"discountPercentage"`
	ID                      *string                                                    `json:"id,omitempty"`
	IssueDate               *time.Time                                                 `json:"issueDate,omitempty"`
	LineItems               []ListCreditNotesLinksSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                *ListCreditNotesLinksSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                                                 `json:"modifiedDate,omitempty"`
	Note                    *string                                                    `json:"note,omitempty"`
	PaymentAllocations      []ListCreditNotesLinksSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	RemainingCredit         float64                                                    `json:"remainingCredit"`
	SourceModifiedDate      *time.Time                                                 `json:"sourceModifiedDate,omitempty"`
	Status                  ListCreditNotesLinksSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal                float64                                                    `json:"subTotal"`
	SupplementalData        *ListCreditNotesLinksSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TotalAmount             float64                                                    `json:"totalAmount"`
	TotalDiscount           float64                                                    `json:"totalDiscount"`
	TotalTaxAmount          float64                                                    `json:"totalTaxAmount"`
	WithholdingTax          []ListCreditNotesLinksSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

// ListCreditNotesLinks
// Codat's Paging Model
type ListCreditNotesLinks struct {
	Links        ListCreditNotesLinksLinks                `json:"_links"`
	PageNumber   int64                                    `json:"pageNumber"`
	PageSize     int64                                    `json:"pageSize"`
	Results      []ListCreditNotesLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                    `json:"totalResults"`
}

type ListCreditNotesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListCreditNotesLinks
}
