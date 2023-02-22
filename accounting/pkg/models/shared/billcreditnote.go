package shared

import (
	"time"
)

type BillCreditNoteLineItems struct {
	AccountRef           *AccountRef `json:"accountRef,omitempty"`
	Description          *string     `json:"description,omitempty"`
	DiscountAmount       *float64    `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64    `json:"discountPercentage,omitempty"`
	ItemRef              *ItemRef    `json:"itemRef,omitempty"`
	Quantity             float64     `json:"quantity"`
	SubTotal             *float64    `json:"subTotal,omitempty"`
	TaxAmount            *float64    `json:"taxAmount,omitempty"`
	TaxRateRef           *TaxRateRef `json:"taxRateRef,omitempty"`
	TotalAmount          *float64    `json:"totalAmount,omitempty"`
	Tracking             *Tracking   `json:"tracking,omitempty"`
	TrackingCategoryRefs []Items     `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64     `json:"unitAmount"`
}

type BillCreditNoteStatusEnum string

const (
	BillCreditNoteStatusEnumUnknown       BillCreditNoteStatusEnum = "Unknown"
	BillCreditNoteStatusEnumDraft         BillCreditNoteStatusEnum = "Draft"
	BillCreditNoteStatusEnumSubmitted     BillCreditNoteStatusEnum = "Submitted"
	BillCreditNoteStatusEnumPaid          BillCreditNoteStatusEnum = "Paid"
	BillCreditNoteStatusEnumVoid          BillCreditNoteStatusEnum = "Void"
	BillCreditNoteStatusEnumPartiallyPaid BillCreditNoteStatusEnum = "PartiallyPaid"
)

// BillCreditNote
// > **Bill credit notes or credit notes?**
// >
// > In Codat, bill credit notes represent accounts payable only. For accounts receivable, see [Credit notes](https://docs.codat.io/docs/datamodel-accounting-creditnotes).
//
// View the coverage for bill credit notes in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A bill credit note is issued by a supplier for the purpose of recording credit. For example, if a supplier was unable to fulfil an order that was placed by a business, or delivered damaged goods, they would issue a bill credit note. A bill credit note reduces the amount a business owes to the supplier. It can be refunded to the business or used to pay off future bills.
//
// In the Codat API, a bill credit note is an accounts payable record issued by a [supplier](https://docs.codat.io/docs/datamodel-accounting-suppliers).
//
// A bill credit note includes details of:
// * The original and remaining credit.
// * Any allocations of the credit against other records, such as [bills](https://docs.codat.io/docs/datamodel-accounting-bills).
// * The supplier that issued the bill credit note.
type BillCreditNote struct {
	AllocatedOnDate      *time.Time                `json:"allocatedOnDate,omitempty"`
	BillCreditNoteNumber *string                   `json:"billCreditNoteNumber,omitempty"`
	Currency             *string                   `json:"currency,omitempty"`
	CurrencyRate         *float64                  `json:"currencyRate,omitempty"`
	DiscountPercentage   float64                   `json:"discountPercentage"`
	ID                   *string                   `json:"id,omitempty"`
	IssueDate            *time.Time                `json:"issueDate,omitempty"`
	LineItems            []BillCreditNoteLineItems `json:"lineItems,omitempty"`
	Metadata             *Metadata                 `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                `json:"modifiedDate,omitempty"`
	Note                 *string                   `json:"note,omitempty"`
	PaymentAllocations   []PaymentAllocationsitems `json:"paymentAllocations,omitempty"`
	RemainingCredit      float64                   `json:"remainingCredit"`
	SourceModifiedDate   *time.Time                `json:"sourceModifiedDate,omitempty"`
	Status               BillCreditNoteStatusEnum  `json:"status"`
	SubTotal             float64                   `json:"subTotal"`
	SupplementalData     *SupplementalData         `json:"supplementalData,omitempty"`
	SupplierRef          *SupplierRef              `json:"supplierRef,omitempty"`
	TotalAmount          float64                   `json:"totalAmount"`
	TotalDiscount        float64                   `json:"totalDiscount"`
	TotalTaxAmount       float64                   `json:"totalTaxAmount"`
	WithholdingTax       []WithholdingTaxitems     `json:"withholdingTax,omitempty"`
}
