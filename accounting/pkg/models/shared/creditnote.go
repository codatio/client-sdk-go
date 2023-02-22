package shared

import (
	"time"
)

type CreditNoteLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type CreditNoteLineItemsTrackingIsBilledToEnum string

const (
	CreditNoteLineItemsTrackingIsBilledToEnumUnknown       CreditNoteLineItemsTrackingIsBilledToEnum = "Unknown"
	CreditNoteLineItemsTrackingIsBilledToEnumNotApplicable CreditNoteLineItemsTrackingIsBilledToEnum = "NotApplicable"
	CreditNoteLineItemsTrackingIsBilledToEnumProject       CreditNoteLineItemsTrackingIsBilledToEnum = "Project"
)

type CreditNoteLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreditNoteLineItemsTracking struct {
	CategoryRefs []Items                                   `json:"categoryRefs"`
	CustomerRef  *CreditNoteLineItemsTrackingCustomerRef   `json:"customerRef,omitempty"`
	IsBilledTo   CreditNoteLineItemsTrackingIsBilledToEnum `json:"isBilledTo"`
	IsRebilledTo IsBilledToEnum1                           `json:"isRebilledTo"`
	ProjectRef   *CreditNoteLineItemsTrackingProjectRef    `json:"projectRef,omitempty"`
}

type CreditNoteLineItems struct {
	AccountRef           *AccountRef                  `json:"accountRef,omitempty"`
	Description          *string                      `json:"description,omitempty"`
	DiscountAmount       *float64                     `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                     `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                        `json:"isDirectIncome,omitempty"`
	ItemRef              *ItemRef                     `json:"itemRef,omitempty"`
	Quantity             float64                      `json:"quantity"`
	SubTotal             *float64                     `json:"subTotal,omitempty"`
	TaxAmount            *float64                     `json:"taxAmount,omitempty"`
	TaxRateRef           *TaxRateRef                  `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                     `json:"totalAmount,omitempty"`
	Tracking             *CreditNoteLineItemsTracking `json:"tracking,omitempty"`
	TrackingCategoryRefs []Items                      `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                      `json:"unitAmount"`
}

type CreditNoteStatusEnum string

const (
	CreditNoteStatusEnumUnknown       CreditNoteStatusEnum = "Unknown"
	CreditNoteStatusEnumDraft         CreditNoteStatusEnum = "Draft"
	CreditNoteStatusEnumSubmitted     CreditNoteStatusEnum = "Submitted"
	CreditNoteStatusEnumPaid          CreditNoteStatusEnum = "Paid"
	CreditNoteStatusEnumVoid          CreditNoteStatusEnum = "Void"
	CreditNoteStatusEnumPartiallyPaid CreditNoteStatusEnum = "PartiallyPaid"
)

// CreditNote
// > View the coverage for credit notes in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=creditNotes" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Think of a credit note as a voucher issued to a customer. It is a reduction that can be applied against one or multiple invoices. A credit note can either reduce the amount owed or cancel out an invoice entirely.
//
// In the Codat system a credit note is issued to a [customer's](https://docs.codat.io/docs/datamodel-accounting-customers) accounts receivable.
//
// It contains details of:
// * The amount of credit remaining and its status.
// * Payment allocations against the payments type, in this case an invoice.
// * Which customers the credit notes have been issued to.
type CreditNote struct {
	AdditionalTaxAmount     *float64                  `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                  `json:"additionalTaxPercentage,omitempty"`
	AllocatedOnDate         *time.Time                `json:"allocatedOnDate,omitempty"`
	CreditNoteNumber        *string                   `json:"creditNoteNumber,omitempty"`
	Currency                *string                   `json:"currency,omitempty"`
	CurrencyRate            *float64                  `json:"currencyRate,omitempty"`
	CustomerRef             *CustomerRef              `json:"customerRef,omitempty"`
	DiscountPercentage      float64                   `json:"discountPercentage"`
	ID                      *string                   `json:"id,omitempty"`
	IssueDate               *time.Time                `json:"issueDate,omitempty"`
	LineItems               []CreditNoteLineItems     `json:"lineItems,omitempty"`
	Metadata                *Metadata                 `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                `json:"modifiedDate,omitempty"`
	Note                    *string                   `json:"note,omitempty"`
	PaymentAllocations      []PaymentAllocationsitems `json:"paymentAllocations,omitempty"`
	RemainingCredit         float64                   `json:"remainingCredit"`
	SourceModifiedDate      *time.Time                `json:"sourceModifiedDate,omitempty"`
	Status                  CreditNoteStatusEnum      `json:"status"`
	SubTotal                float64                   `json:"subTotal"`
	SupplementalData        *SupplementalData         `json:"supplementalData,omitempty"`
	TotalAmount             float64                   `json:"totalAmount"`
	TotalDiscount           float64                   `json:"totalDiscount"`
	TotalTaxAmount          float64                   `json:"totalTaxAmount"`
	WithholdingTax          []WithholdingTaxitems     `json:"withholdingTax,omitempty"`
}
