package operations

import (
	"net/http"
	"time"
)

type GetCreditNotePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	CreditNoteID string `pathParam:"style=simple,explode=false,name=creditNoteId"`
}

type GetCreditNoteRequest struct {
	PathParams GetCreditNotePathParams
}

// GetCreditNoteSourceModifiedDateCustomerRef
// Reference to the customer the credit note has been issued to.
type GetCreditNoteSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

// GetCreditNoteSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type GetCreditNoteSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetCreditNoteSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type GetCreditNoteSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// GetCreditNoteSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type GetCreditNoteSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type GetCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type GetCreditNoteSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type GetCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	GetCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       GetCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	GetCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable GetCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	GetCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       GetCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type GetCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	GetCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       GetCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	GetCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable GetCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	GetCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       GetCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type GetCreditNoteSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type GetCreditNoteSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []GetCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *GetCreditNoteSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   GetCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo GetCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *GetCreditNoteSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type GetCreditNoteSourceModifiedDateLineItems struct {
	AccountRef           *GetCreditNoteSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                        `json:"description,omitempty"`
	DiscountAmount       *float64                                                       `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                       `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                                                          `json:"isDirectIncome,omitempty"`
	ItemRef              *GetCreditNoteSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                        `json:"quantity"`
	SubTotal             *float64                                                       `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                       `json:"taxAmount,omitempty"`
	TaxRateRef           *GetCreditNoteSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                       `json:"totalAmount,omitempty"`
	Tracking             *GetCreditNoteSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []GetCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                        `json:"unitAmount"`
}

type GetCreditNoteSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetCreditNoteSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// GetCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type GetCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetCreditNoteSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *GetCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                             `json:"currency,omitempty"`
	CurrencyRate *float64                                                            `json:"currencyRate,omitempty"`
	ID           *string                                                             `json:"id,omitempty"`
	Note         *string                                                             `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                          `json:"paidOnDate,omitempty"`
	Reference    *string                                                             `json:"reference,omitempty"`
	TotalAmount  *float64                                                            `json:"totalAmount,omitempty"`
}

type GetCreditNoteSourceModifiedDatePaymentAllocations struct {
	Allocation GetCreditNoteSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    GetCreditNoteSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type GetCreditNoteSourceModifiedDateStatusEnum string

const (
	GetCreditNoteSourceModifiedDateStatusEnumUnknown       GetCreditNoteSourceModifiedDateStatusEnum = "Unknown"
	GetCreditNoteSourceModifiedDateStatusEnumDraft         GetCreditNoteSourceModifiedDateStatusEnum = "Draft"
	GetCreditNoteSourceModifiedDateStatusEnumSubmitted     GetCreditNoteSourceModifiedDateStatusEnum = "Submitted"
	GetCreditNoteSourceModifiedDateStatusEnumPaid          GetCreditNoteSourceModifiedDateStatusEnum = "Paid"
	GetCreditNoteSourceModifiedDateStatusEnumVoid          GetCreditNoteSourceModifiedDateStatusEnum = "Void"
	GetCreditNoteSourceModifiedDateStatusEnumPartiallyPaid GetCreditNoteSourceModifiedDateStatusEnum = "PartiallyPaid"
)

// GetCreditNoteSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type GetCreditNoteSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type GetCreditNoteSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// GetCreditNoteSourceModifiedDate
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
type GetCreditNoteSourceModifiedDate struct {
	AdditionalTaxAmount     *float64                                            `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                                            `json:"additionalTaxPercentage,omitempty"`
	AllocatedOnDate         *time.Time                                          `json:"allocatedOnDate,omitempty"`
	CreditNoteNumber        *string                                             `json:"creditNoteNumber,omitempty"`
	Currency                *string                                             `json:"currency,omitempty"`
	CurrencyRate            *float64                                            `json:"currencyRate,omitempty"`
	CustomerRef             *GetCreditNoteSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	DiscountPercentage      float64                                             `json:"discountPercentage"`
	ID                      *string                                             `json:"id,omitempty"`
	IssueDate               *time.Time                                          `json:"issueDate,omitempty"`
	LineItems               []GetCreditNoteSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                *GetCreditNoteSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                                          `json:"modifiedDate,omitempty"`
	Note                    *string                                             `json:"note,omitempty"`
	PaymentAllocations      []GetCreditNoteSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	RemainingCredit         float64                                             `json:"remainingCredit"`
	SourceModifiedDate      *time.Time                                          `json:"sourceModifiedDate,omitempty"`
	Status                  GetCreditNoteSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal                float64                                             `json:"subTotal"`
	SupplementalData        *GetCreditNoteSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TotalAmount             float64                                             `json:"totalAmount"`
	TotalDiscount           float64                                             `json:"totalDiscount"`
	TotalTaxAmount          float64                                             `json:"totalTaxAmount"`
	WithholdingTax          []GetCreditNoteSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type GetCreditNoteResponse struct {
	ContentType        string
	SourceModifiedDate *GetCreditNoteSourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
