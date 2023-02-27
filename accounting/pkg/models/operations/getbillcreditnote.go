package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type GetBillCreditNotePathParams struct {
	BillCreditNoteID string `pathParam:"style=simple,explode=false,name=billCreditNoteId"`
	CompanyID        string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetBillCreditNoteSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetBillCreditNoteRequest struct {
	PathParams GetBillCreditNotePathParams
	Security   GetBillCreditNoteSecurity
}

// GetBillCreditNoteSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type GetBillCreditNoteSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetBillCreditNoteSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type GetBillCreditNoteSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// GetBillCreditNoteSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type GetBillCreditNoteSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type GetBillCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type GetBillCreditNoteSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumCustomer      GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Customer"
	GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumCustomer      GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Customer"
	GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type GetBillCreditNoteSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type GetBillCreditNoteSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []GetBillCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *GetBillCreditNoteSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo GetBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *GetBillCreditNoteSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type GetBillCreditNoteSourceModifiedDateLineItems struct {
	AccountRef           *GetBillCreditNoteSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                            `json:"description,omitempty"`
	DiscountAmount       *float64                                                           `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                           `json:"discountPercentage,omitempty"`
	ItemRef              *GetBillCreditNoteSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                            `json:"quantity"`
	SubTotal             *float64                                                           `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                           `json:"taxAmount,omitempty"`
	TaxRateRef           *GetBillCreditNoteSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                           `json:"totalAmount,omitempty"`
	Tracking             *GetBillCreditNoteSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []GetBillCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                            `json:"unitAmount"`
}

type GetBillCreditNoteSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetBillCreditNoteSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// GetBillCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type GetBillCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetBillCreditNoteSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *GetBillCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                 `json:"currency,omitempty"`
	CurrencyRate *float64                                                                `json:"currencyRate,omitempty"`
	ID           *string                                                                 `json:"id,omitempty"`
	Note         *string                                                                 `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                              `json:"paidOnDate,omitempty"`
	Reference    *string                                                                 `json:"reference,omitempty"`
	TotalAmount  *float64                                                                `json:"totalAmount,omitempty"`
}

type GetBillCreditNoteSourceModifiedDatePaymentAllocations struct {
	Allocation GetBillCreditNoteSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    GetBillCreditNoteSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type GetBillCreditNoteSourceModifiedDateStatusEnum string

const (
	GetBillCreditNoteSourceModifiedDateStatusEnumUnknown       GetBillCreditNoteSourceModifiedDateStatusEnum = "Unknown"
	GetBillCreditNoteSourceModifiedDateStatusEnumDraft         GetBillCreditNoteSourceModifiedDateStatusEnum = "Draft"
	GetBillCreditNoteSourceModifiedDateStatusEnumSubmitted     GetBillCreditNoteSourceModifiedDateStatusEnum = "Submitted"
	GetBillCreditNoteSourceModifiedDateStatusEnumPaid          GetBillCreditNoteSourceModifiedDateStatusEnum = "Paid"
	GetBillCreditNoteSourceModifiedDateStatusEnumVoid          GetBillCreditNoteSourceModifiedDateStatusEnum = "Void"
	GetBillCreditNoteSourceModifiedDateStatusEnumPartiallyPaid GetBillCreditNoteSourceModifiedDateStatusEnum = "PartiallyPaid"
)

type GetBillCreditNoteSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// GetBillCreditNoteSourceModifiedDateSupplierRef
// Supplier that issued the bill credit note.
type GetBillCreditNoteSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type GetBillCreditNoteSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// GetBillCreditNoteSourceModifiedDate
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
type GetBillCreditNoteSourceModifiedDate struct {
	AllocatedOnDate      *time.Time                                              `json:"allocatedOnDate,omitempty"`
	BillCreditNoteNumber *string                                                 `json:"billCreditNoteNumber,omitempty"`
	Currency             *string                                                 `json:"currency,omitempty"`
	CurrencyRate         *float64                                                `json:"currencyRate,omitempty"`
	DiscountPercentage   float64                                                 `json:"discountPercentage"`
	ID                   *string                                                 `json:"id,omitempty"`
	IssueDate            *time.Time                                              `json:"issueDate,omitempty"`
	LineItems            []GetBillCreditNoteSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata             *GetBillCreditNoteSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                              `json:"modifiedDate,omitempty"`
	Note                 *string                                                 `json:"note,omitempty"`
	PaymentAllocations   []GetBillCreditNoteSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	RemainingCredit      float64                                                 `json:"remainingCredit"`
	SourceModifiedDate   *time.Time                                              `json:"sourceModifiedDate,omitempty"`
	Status               GetBillCreditNoteSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal             float64                                                 `json:"subTotal"`
	SupplementalData     *GetBillCreditNoteSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef          *GetBillCreditNoteSourceModifiedDateSupplierRef         `json:"supplierRef,omitempty"`
	TotalAmount          float64                                                 `json:"totalAmount"`
	TotalDiscount        float64                                                 `json:"totalDiscount"`
	TotalTaxAmount       float64                                                 `json:"totalTaxAmount"`
	WithholdingTax       []GetBillCreditNoteSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type GetBillCreditNoteResponse struct {
	ContentType        string
	SourceModifiedDate *GetBillCreditNoteSourceModifiedDate
	StatusCode         int
}
