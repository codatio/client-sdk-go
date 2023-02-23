package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type ListBillCreditNotesPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListBillCreditNotesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListBillCreditNotesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListBillCreditNotesRequest struct {
	PathParams  ListBillCreditNotesPathParams
	QueryParams ListBillCreditNotesQueryParams
	Security    ListBillCreditNotesSecurity
}

type ListBillCreditNotesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListBillCreditNotesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListBillCreditNotesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListBillCreditNotesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListBillCreditNotesLinksLinks struct {
	Current  ListBillCreditNotesLinksLinksCurrent   `json:"current"`
	Next     *ListBillCreditNotesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListBillCreditNotesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListBillCreditNotesLinksLinksSelf      `json:"self"`
}

// ListBillCreditNotesLinksSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type ListBillCreditNotesLinksSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ListBillCreditNotesLinksSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type ListBillCreditNotesLinksSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// ListBillCreditNotesLinksSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type ListBillCreditNotesLinksSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnumCustomer      ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Customer"
	ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnumCustomer      ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Customer"
	ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type ListBillCreditNotesLinksSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type ListBillCreditNotesLinksSourceModifiedDateLineItems struct {
	AccountRef           *ListBillCreditNotesLinksSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                   `json:"description,omitempty"`
	DiscountAmount       *float64                                                                  `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                  `json:"discountPercentage,omitempty"`
	ItemRef              *ListBillCreditNotesLinksSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                   `json:"quantity"`
	SubTotal             *float64                                                                  `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                  `json:"taxAmount,omitempty"`
	TaxRateRef           *ListBillCreditNotesLinksSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                  `json:"totalAmount,omitempty"`
	Tracking             *ListBillCreditNotesLinksSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []ListBillCreditNotesLinksSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                   `json:"unitAmount"`
}

type ListBillCreditNotesLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type ListBillCreditNotesLinksSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// ListBillCreditNotesLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type ListBillCreditNotesLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ListBillCreditNotesLinksSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *ListBillCreditNotesLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                        `json:"currency,omitempty"`
	CurrencyRate *float64                                                                       `json:"currencyRate,omitempty"`
	ID           *string                                                                        `json:"id,omitempty"`
	Note         *string                                                                        `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                     `json:"paidOnDate,omitempty"`
	Reference    *string                                                                        `json:"reference,omitempty"`
	TotalAmount  *float64                                                                       `json:"totalAmount,omitempty"`
}

type ListBillCreditNotesLinksSourceModifiedDatePaymentAllocations struct {
	Allocation ListBillCreditNotesLinksSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    ListBillCreditNotesLinksSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type ListBillCreditNotesLinksSourceModifiedDateStatusEnum string

const (
	ListBillCreditNotesLinksSourceModifiedDateStatusEnumUnknown       ListBillCreditNotesLinksSourceModifiedDateStatusEnum = "Unknown"
	ListBillCreditNotesLinksSourceModifiedDateStatusEnumDraft         ListBillCreditNotesLinksSourceModifiedDateStatusEnum = "Draft"
	ListBillCreditNotesLinksSourceModifiedDateStatusEnumSubmitted     ListBillCreditNotesLinksSourceModifiedDateStatusEnum = "Submitted"
	ListBillCreditNotesLinksSourceModifiedDateStatusEnumPaid          ListBillCreditNotesLinksSourceModifiedDateStatusEnum = "Paid"
	ListBillCreditNotesLinksSourceModifiedDateStatusEnumVoid          ListBillCreditNotesLinksSourceModifiedDateStatusEnum = "Void"
	ListBillCreditNotesLinksSourceModifiedDateStatusEnumPartiallyPaid ListBillCreditNotesLinksSourceModifiedDateStatusEnum = "PartiallyPaid"
)

type ListBillCreditNotesLinksSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// ListBillCreditNotesLinksSourceModifiedDateSupplierRef
// Supplier that issued the bill credit note.
type ListBillCreditNotesLinksSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type ListBillCreditNotesLinksSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// ListBillCreditNotesLinksSourceModifiedDate
// > **Bill credit notes or credit notes?**
// >
// > In Codat, bill credit notes represent accounts payable only. For accounts receivable, see [Credit notes](https://docs.codat.io/accounting-api#/schemas/CreditNote).
//
// View the coverage for bill credit notes in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A bill credit note is issued by a supplier for the purpose of recording credit. For example, if a supplier was unable to fulfil an order that was placed by a business, or delivered damaged goods, they would issue a bill credit note. A bill credit note reduces the amount a business owes to the supplier. It can be refunded to the business or used to pay off future bills.
//
// In the Codat API, a bill credit note is an accounts payable record issued by a [supplier](https://docs.codat.io/accounting-api#/schemas/Supplier).
//
// A bill credit note includes details of:
// * The original and remaining credit.
// * Any allocations of the credit against other records, such as [bills](https://docs.codat.io/accounting-api#/schemas/Bill).
// * The supplier that issued the bill credit note.
type ListBillCreditNotesLinksSourceModifiedDate struct {
	AllocatedOnDate      *time.Time                                                     `json:"allocatedOnDate,omitempty"`
	BillCreditNoteNumber *string                                                        `json:"billCreditNoteNumber,omitempty"`
	Currency             *string                                                        `json:"currency,omitempty"`
	CurrencyRate         *float64                                                       `json:"currencyRate,omitempty"`
	DiscountPercentage   float64                                                        `json:"discountPercentage"`
	ID                   *string                                                        `json:"id,omitempty"`
	IssueDate            *time.Time                                                     `json:"issueDate,omitempty"`
	LineItems            []ListBillCreditNotesLinksSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata             *ListBillCreditNotesLinksSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                                     `json:"modifiedDate,omitempty"`
	Note                 *string                                                        `json:"note,omitempty"`
	PaymentAllocations   []ListBillCreditNotesLinksSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	RemainingCredit      float64                                                        `json:"remainingCredit"`
	SourceModifiedDate   *time.Time                                                     `json:"sourceModifiedDate,omitempty"`
	Status               ListBillCreditNotesLinksSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal             float64                                                        `json:"subTotal"`
	SupplementalData     *ListBillCreditNotesLinksSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef          *ListBillCreditNotesLinksSourceModifiedDateSupplierRef         `json:"supplierRef,omitempty"`
	TotalAmount          float64                                                        `json:"totalAmount"`
	TotalDiscount        float64                                                        `json:"totalDiscount"`
	TotalTaxAmount       float64                                                        `json:"totalTaxAmount"`
	WithholdingTax       []ListBillCreditNotesLinksSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

// ListBillCreditNotesLinks
// Codat's Paging Model
type ListBillCreditNotesLinks struct {
	Links        ListBillCreditNotesLinksLinks                `json:"_links"`
	PageNumber   int64                                        `json:"pageNumber"`
	PageSize     int64                                        `json:"pageSize"`
	Results      []ListBillCreditNotesLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                        `json:"totalResults"`
}

type ListBillCreditNotesResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListBillCreditNotesLinks
}
