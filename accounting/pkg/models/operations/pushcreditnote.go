package operations

import (
	"net/http"
	"time"
)

type PushCreditNotePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PushCreditNoteQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// PushCreditNoteSourceModifiedDateCustomerRef
// Reference to the customer the credit note has been issued to.
type PushCreditNoteSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

// PushCreditNoteSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type PushCreditNoteSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PushCreditNoteSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type PushCreditNoteSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PushCreditNoteSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type PushCreditNoteSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type PushCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PushCreditNoteSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type PushCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	PushCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       PushCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	PushCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable PushCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	PushCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       PushCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type PushCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	PushCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       PushCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	PushCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable PushCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	PushCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       PushCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type PushCreditNoteSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PushCreditNoteSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []PushCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *PushCreditNoteSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   PushCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo PushCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *PushCreditNoteSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type PushCreditNoteSourceModifiedDateLineItems struct {
	AccountRef           *PushCreditNoteSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                         `json:"description,omitempty"`
	DiscountAmount       *float64                                                        `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                        `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                                                           `json:"isDirectIncome,omitempty"`
	ItemRef              *PushCreditNoteSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                         `json:"quantity"`
	SubTotal             *float64                                                        `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                        `json:"taxAmount,omitempty"`
	TaxRateRef           *PushCreditNoteSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                        `json:"totalAmount,omitempty"`
	Tracking             *PushCreditNoteSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []PushCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                         `json:"unitAmount"`
}

type PushCreditNoteSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PushCreditNoteSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// PushCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type PushCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PushCreditNoteSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *PushCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                              `json:"currency,omitempty"`
	CurrencyRate *float64                                                             `json:"currencyRate,omitempty"`
	ID           *string                                                              `json:"id,omitempty"`
	Note         *string                                                              `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                           `json:"paidOnDate,omitempty"`
	Reference    *string                                                              `json:"reference,omitempty"`
	TotalAmount  *float64                                                             `json:"totalAmount,omitempty"`
}

type PushCreditNoteSourceModifiedDatePaymentAllocations struct {
	Allocation PushCreditNoteSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    PushCreditNoteSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type PushCreditNoteSourceModifiedDateStatusEnum string

const (
	PushCreditNoteSourceModifiedDateStatusEnumUnknown       PushCreditNoteSourceModifiedDateStatusEnum = "Unknown"
	PushCreditNoteSourceModifiedDateStatusEnumDraft         PushCreditNoteSourceModifiedDateStatusEnum = "Draft"
	PushCreditNoteSourceModifiedDateStatusEnumSubmitted     PushCreditNoteSourceModifiedDateStatusEnum = "Submitted"
	PushCreditNoteSourceModifiedDateStatusEnumPaid          PushCreditNoteSourceModifiedDateStatusEnum = "Paid"
	PushCreditNoteSourceModifiedDateStatusEnumVoid          PushCreditNoteSourceModifiedDateStatusEnum = "Void"
	PushCreditNoteSourceModifiedDateStatusEnumPartiallyPaid PushCreditNoteSourceModifiedDateStatusEnum = "PartiallyPaid"
)

type PushCreditNoteSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type PushCreditNoteSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// PushCreditNoteSourceModifiedDate
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
type PushCreditNoteSourceModifiedDate struct {
	AdditionalTaxAmount     *float64                                             `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                                             `json:"additionalTaxPercentage,omitempty"`
	AllocatedOnDate         *time.Time                                           `json:"allocatedOnDate,omitempty"`
	CreditNoteNumber        *string                                              `json:"creditNoteNumber,omitempty"`
	Currency                *string                                              `json:"currency,omitempty"`
	CurrencyRate            *float64                                             `json:"currencyRate,omitempty"`
	CustomerRef             *PushCreditNoteSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	DiscountPercentage      float64                                              `json:"discountPercentage"`
	ID                      *string                                              `json:"id,omitempty"`
	IssueDate               *time.Time                                           `json:"issueDate,omitempty"`
	LineItems               []PushCreditNoteSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                *PushCreditNoteSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                                           `json:"modifiedDate,omitempty"`
	Note                    *string                                              `json:"note,omitempty"`
	PaymentAllocations      []PushCreditNoteSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	RemainingCredit         float64                                              `json:"remainingCredit"`
	SourceModifiedDate      *time.Time                                           `json:"sourceModifiedDate,omitempty"`
	Status                  PushCreditNoteSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal                float64                                              `json:"subTotal"`
	SupplementalData        *PushCreditNoteSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TotalAmount             float64                                              `json:"totalAmount"`
	TotalDiscount           float64                                              `json:"totalDiscount"`
	TotalTaxAmount          float64                                              `json:"totalTaxAmount"`
	WithholdingTax          []PushCreditNoteSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type PushCreditNoteRequest struct {
	PathParams  PushCreditNotePathParams
	QueryParams PushCreditNoteQueryParams
	Request     *PushCreditNoteSourceModifiedDate `request:"mediaType=application/json"`
}

type PushCreditNote200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PushCreditNote200ApplicationJSONChangesTypeEnum string

const (
	PushCreditNote200ApplicationJSONChangesTypeEnumUnknown            PushCreditNote200ApplicationJSONChangesTypeEnum = "Unknown"
	PushCreditNote200ApplicationJSONChangesTypeEnumCreated            PushCreditNote200ApplicationJSONChangesTypeEnum = "Created"
	PushCreditNote200ApplicationJSONChangesTypeEnumModified           PushCreditNote200ApplicationJSONChangesTypeEnum = "Modified"
	PushCreditNote200ApplicationJSONChangesTypeEnumDeleted            PushCreditNote200ApplicationJSONChangesTypeEnum = "Deleted"
	PushCreditNote200ApplicationJSONChangesTypeEnumAttachmentUploaded PushCreditNote200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PushCreditNote200ApplicationJSONChanges struct {
	AttachmentID *string                                                        `json:"attachmentId,omitempty"`
	RecordRef    *PushCreditNote200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PushCreditNote200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// PushCreditNote200ApplicationJSONSourceModifiedDateCustomerRef
// Reference to the customer the credit note has been issued to.
type PushCreditNote200ApplicationJSONSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

// PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type PushCreditNote200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                           `json:"description,omitempty"`
	DiscountAmount       *float64                                                                          `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                          `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                                                                             `json:"isDirectIncome,omitempty"`
	ItemRef              *PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                           `json:"quantity"`
	SubTotal             *float64                                                                          `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                          `json:"taxAmount,omitempty"`
	TaxRateRef           *PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                          `json:"totalAmount,omitempty"`
	Tracking             *PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []PushCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                           `json:"unitAmount"`
}

type PushCreditNote200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PushCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// PushCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type PushCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PushCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *PushCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                                `json:"currency,omitempty"`
	CurrencyRate *float64                                                                               `json:"currencyRate,omitempty"`
	ID           *string                                                                                `json:"id,omitempty"`
	Note         *string                                                                                `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                             `json:"paidOnDate,omitempty"`
	Reference    *string                                                                                `json:"reference,omitempty"`
	TotalAmount  *float64                                                                               `json:"totalAmount,omitempty"`
}

type PushCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation PushCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    PushCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type PushCreditNote200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	PushCreditNote200ApplicationJSONSourceModifiedDateStatusEnumUnknown       PushCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	PushCreditNote200ApplicationJSONSourceModifiedDateStatusEnumDraft         PushCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Draft"
	PushCreditNote200ApplicationJSONSourceModifiedDateStatusEnumSubmitted     PushCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Submitted"
	PushCreditNote200ApplicationJSONSourceModifiedDateStatusEnumPaid          PushCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Paid"
	PushCreditNote200ApplicationJSONSourceModifiedDateStatusEnumVoid          PushCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Void"
	PushCreditNote200ApplicationJSONSourceModifiedDateStatusEnumPartiallyPaid PushCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "PartiallyPaid"
)

type PushCreditNote200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type PushCreditNote200ApplicationJSONSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// PushCreditNote200ApplicationJSONSourceModifiedDate
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
type PushCreditNote200ApplicationJSONSourceModifiedDate struct {
	AdditionalTaxAmount     *float64                                                               `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                                                               `json:"additionalTaxPercentage,omitempty"`
	AllocatedOnDate         *time.Time                                                             `json:"allocatedOnDate,omitempty"`
	CreditNoteNumber        *string                                                                `json:"creditNoteNumber,omitempty"`
	Currency                *string                                                                `json:"currency,omitempty"`
	CurrencyRate            *float64                                                               `json:"currencyRate,omitempty"`
	CustomerRef             *PushCreditNote200ApplicationJSONSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	DiscountPercentage      float64                                                                `json:"discountPercentage"`
	ID                      *string                                                                `json:"id,omitempty"`
	IssueDate               *time.Time                                                             `json:"issueDate,omitempty"`
	LineItems               []PushCreditNote200ApplicationJSONSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                *PushCreditNote200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                                                             `json:"modifiedDate,omitempty"`
	Note                    *string                                                                `json:"note,omitempty"`
	PaymentAllocations      []PushCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	RemainingCredit         float64                                                                `json:"remainingCredit"`
	SourceModifiedDate      *time.Time                                                             `json:"sourceModifiedDate,omitempty"`
	Status                  PushCreditNote200ApplicationJSONSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal                float64                                                                `json:"subTotal"`
	SupplementalData        *PushCreditNote200ApplicationJSONSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TotalAmount             float64                                                                `json:"totalAmount"`
	TotalDiscount           float64                                                                `json:"totalDiscount"`
	TotalTaxAmount          float64                                                                `json:"totalTaxAmount"`
	WithholdingTax          []PushCreditNote200ApplicationJSONSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type PushCreditNote200ApplicationJSONStatusEnum string

const (
	PushCreditNote200ApplicationJSONStatusEnumPending  PushCreditNote200ApplicationJSONStatusEnum = "Pending"
	PushCreditNote200ApplicationJSONStatusEnumFailed   PushCreditNote200ApplicationJSONStatusEnum = "Failed"
	PushCreditNote200ApplicationJSONStatusEnumSuccess  PushCreditNote200ApplicationJSONStatusEnum = "Success"
	PushCreditNote200ApplicationJSONStatusEnumTimedOut PushCreditNote200ApplicationJSONStatusEnum = "TimedOut"
)

type PushCreditNote200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PushCreditNote200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PushCreditNote200ApplicationJSONValidation struct {
	Errors   []PushCreditNote200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PushCreditNote200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PushCreditNote200ApplicationJSON struct {
	Changes           []PushCreditNote200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                              `json:"companyId"`
	CompletedOnUtc    *time.Time                                          `json:"completedOnUtc,omitempty"`
	Data              *PushCreditNote200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                              `json:"dataConnectionKey"`
	DataType          *string                                             `json:"dataType,omitempty"`
	ErrorMessage      *string                                             `json:"errorMessage,omitempty"`
	PushOperationKey  string                                              `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                           `json:"requestedOnUtc"`
	Status            PushCreditNote200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                 `json:"statusCode"`
	TimeoutInMinutes  *int                                                `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                `json:"timeoutInSeconds,omitempty"`
	Validation        *PushCreditNote200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PushCreditNoteResponse struct {
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
	PushCreditNote200ApplicationJSONObject *PushCreditNote200ApplicationJSON
}
