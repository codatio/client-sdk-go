package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostBillCreditNotePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostBillCreditNoteQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// PostBillCreditNoteSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type PostBillCreditNoteSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostBillCreditNoteSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type PostBillCreditNoteSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PostBillCreditNoteSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type PostBillCreditNoteSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type PostBillCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostBillCreditNoteSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumCustomer      PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Customer"
	PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumCustomer      PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Customer"
	PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type PostBillCreditNoteSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostBillCreditNoteSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []PostBillCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *PostBillCreditNoteSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo PostBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *PostBillCreditNoteSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type PostBillCreditNoteSourceModifiedDateLineItems struct {
	AccountRef           *PostBillCreditNoteSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                             `json:"description,omitempty"`
	DiscountAmount       *float64                                                            `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                            `json:"discountPercentage,omitempty"`
	ItemRef              *PostBillCreditNoteSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                             `json:"quantity"`
	SubTotal             *float64                                                            `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                            `json:"taxAmount,omitempty"`
	TaxRateRef           *PostBillCreditNoteSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                            `json:"totalAmount,omitempty"`
	Tracking             *PostBillCreditNoteSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []PostBillCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                             `json:"unitAmount"`
}

type PostBillCreditNoteSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostBillCreditNoteSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// PostBillCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type PostBillCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PostBillCreditNoteSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *PostBillCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                  `json:"currency,omitempty"`
	CurrencyRate *float64                                                                 `json:"currencyRate,omitempty"`
	ID           *string                                                                  `json:"id,omitempty"`
	Note         *string                                                                  `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                               `json:"paidOnDate,omitempty"`
	Reference    *string                                                                  `json:"reference,omitempty"`
	TotalAmount  *float64                                                                 `json:"totalAmount,omitempty"`
}

type PostBillCreditNoteSourceModifiedDatePaymentAllocations struct {
	Allocation PostBillCreditNoteSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    PostBillCreditNoteSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type PostBillCreditNoteSourceModifiedDateStatusEnum string

const (
	PostBillCreditNoteSourceModifiedDateStatusEnumUnknown       PostBillCreditNoteSourceModifiedDateStatusEnum = "Unknown"
	PostBillCreditNoteSourceModifiedDateStatusEnumDraft         PostBillCreditNoteSourceModifiedDateStatusEnum = "Draft"
	PostBillCreditNoteSourceModifiedDateStatusEnumSubmitted     PostBillCreditNoteSourceModifiedDateStatusEnum = "Submitted"
	PostBillCreditNoteSourceModifiedDateStatusEnumPaid          PostBillCreditNoteSourceModifiedDateStatusEnum = "Paid"
	PostBillCreditNoteSourceModifiedDateStatusEnumVoid          PostBillCreditNoteSourceModifiedDateStatusEnum = "Void"
	PostBillCreditNoteSourceModifiedDateStatusEnumPartiallyPaid PostBillCreditNoteSourceModifiedDateStatusEnum = "PartiallyPaid"
)

type PostBillCreditNoteSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// PostBillCreditNoteSourceModifiedDateSupplierRef
// Supplier that issued the bill credit note.
type PostBillCreditNoteSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type PostBillCreditNoteSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// PostBillCreditNoteSourceModifiedDate
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
type PostBillCreditNoteSourceModifiedDate struct {
	AllocatedOnDate      *time.Time                                               `json:"allocatedOnDate,omitempty"`
	BillCreditNoteNumber *string                                                  `json:"billCreditNoteNumber,omitempty"`
	Currency             *string                                                  `json:"currency,omitempty"`
	CurrencyRate         *float64                                                 `json:"currencyRate,omitempty"`
	DiscountPercentage   float64                                                  `json:"discountPercentage"`
	ID                   *string                                                  `json:"id,omitempty"`
	IssueDate            *time.Time                                               `json:"issueDate,omitempty"`
	LineItems            []PostBillCreditNoteSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata             *PostBillCreditNoteSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                               `json:"modifiedDate,omitempty"`
	Note                 *string                                                  `json:"note,omitempty"`
	PaymentAllocations   []PostBillCreditNoteSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	RemainingCredit      float64                                                  `json:"remainingCredit"`
	SourceModifiedDate   *time.Time                                               `json:"sourceModifiedDate,omitempty"`
	Status               PostBillCreditNoteSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal             float64                                                  `json:"subTotal"`
	SupplementalData     *PostBillCreditNoteSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef          *PostBillCreditNoteSourceModifiedDateSupplierRef         `json:"supplierRef,omitempty"`
	TotalAmount          float64                                                  `json:"totalAmount"`
	TotalDiscount        float64                                                  `json:"totalDiscount"`
	TotalTaxAmount       float64                                                  `json:"totalTaxAmount"`
	WithholdingTax       []PostBillCreditNoteSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type PostBillCreditNoteSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostBillCreditNoteRequest struct {
	PathParams  PostBillCreditNotePathParams
	QueryParams PostBillCreditNoteQueryParams
	Request     *PostBillCreditNoteSourceModifiedDate `request:"mediaType=application/json"`
	Security    PostBillCreditNoteSecurity
}

type PostBillCreditNote200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostBillCreditNote200ApplicationJSONChangesTypeEnum string

const (
	PostBillCreditNote200ApplicationJSONChangesTypeEnumUnknown            PostBillCreditNote200ApplicationJSONChangesTypeEnum = "Unknown"
	PostBillCreditNote200ApplicationJSONChangesTypeEnumCreated            PostBillCreditNote200ApplicationJSONChangesTypeEnum = "Created"
	PostBillCreditNote200ApplicationJSONChangesTypeEnumModified           PostBillCreditNote200ApplicationJSONChangesTypeEnum = "Modified"
	PostBillCreditNote200ApplicationJSONChangesTypeEnumDeleted            PostBillCreditNote200ApplicationJSONChangesTypeEnum = "Deleted"
	PostBillCreditNote200ApplicationJSONChangesTypeEnumAttachmentUploaded PostBillCreditNote200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostBillCreditNote200ApplicationJSONChanges struct {
	AttachmentID *string                                                            `json:"attachmentId,omitempty"`
	RecordRef    *PostBillCreditNote200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostBillCreditNote200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumCustomer      PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Customer"
	PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumCustomer      PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Customer"
	PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                               `json:"description,omitempty"`
	DiscountAmount       *float64                                                                              `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                              `json:"discountPercentage,omitempty"`
	ItemRef              *PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                               `json:"quantity"`
	SubTotal             *float64                                                                              `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                              `json:"taxAmount,omitempty"`
	TaxRateRef           *PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                              `json:"totalAmount,omitempty"`
	Tracking             *PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                               `json:"unitAmount"`
}

type PostBillCreditNote200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// PostBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type PostBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PostBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *PostBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                                    `json:"currency,omitempty"`
	CurrencyRate *float64                                                                                   `json:"currencyRate,omitempty"`
	ID           *string                                                                                    `json:"id,omitempty"`
	Note         *string                                                                                    `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                                 `json:"paidOnDate,omitempty"`
	Reference    *string                                                                                    `json:"reference,omitempty"`
	TotalAmount  *float64                                                                                   `json:"totalAmount,omitempty"`
}

type PostBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation PostBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    PostBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type PostBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	PostBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumUnknown       PostBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	PostBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumDraft         PostBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Draft"
	PostBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumSubmitted     PostBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Submitted"
	PostBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumPaid          PostBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Paid"
	PostBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumVoid          PostBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Void"
	PostBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumPartiallyPaid PostBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "PartiallyPaid"
)

type PostBillCreditNote200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// PostBillCreditNote200ApplicationJSONSourceModifiedDateSupplierRef
// Supplier that issued the bill credit note.
type PostBillCreditNote200ApplicationJSONSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type PostBillCreditNote200ApplicationJSONSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// PostBillCreditNote200ApplicationJSONSourceModifiedDate
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
type PostBillCreditNote200ApplicationJSONSourceModifiedDate struct {
	AllocatedOnDate      *time.Time                                                                 `json:"allocatedOnDate,omitempty"`
	BillCreditNoteNumber *string                                                                    `json:"billCreditNoteNumber,omitempty"`
	Currency             *string                                                                    `json:"currency,omitempty"`
	CurrencyRate         *float64                                                                   `json:"currencyRate,omitempty"`
	DiscountPercentage   float64                                                                    `json:"discountPercentage"`
	ID                   *string                                                                    `json:"id,omitempty"`
	IssueDate            *time.Time                                                                 `json:"issueDate,omitempty"`
	LineItems            []PostBillCreditNote200ApplicationJSONSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata             *PostBillCreditNote200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                                                 `json:"modifiedDate,omitempty"`
	Note                 *string                                                                    `json:"note,omitempty"`
	PaymentAllocations   []PostBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	RemainingCredit      float64                                                                    `json:"remainingCredit"`
	SourceModifiedDate   *time.Time                                                                 `json:"sourceModifiedDate,omitempty"`
	Status               PostBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal             float64                                                                    `json:"subTotal"`
	SupplementalData     *PostBillCreditNote200ApplicationJSONSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef          *PostBillCreditNote200ApplicationJSONSourceModifiedDateSupplierRef         `json:"supplierRef,omitempty"`
	TotalAmount          float64                                                                    `json:"totalAmount"`
	TotalDiscount        float64                                                                    `json:"totalDiscount"`
	TotalTaxAmount       float64                                                                    `json:"totalTaxAmount"`
	WithholdingTax       []PostBillCreditNote200ApplicationJSONSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type PostBillCreditNote200ApplicationJSONStatusEnum string

const (
	PostBillCreditNote200ApplicationJSONStatusEnumPending  PostBillCreditNote200ApplicationJSONStatusEnum = "Pending"
	PostBillCreditNote200ApplicationJSONStatusEnumFailed   PostBillCreditNote200ApplicationJSONStatusEnum = "Failed"
	PostBillCreditNote200ApplicationJSONStatusEnumSuccess  PostBillCreditNote200ApplicationJSONStatusEnum = "Success"
	PostBillCreditNote200ApplicationJSONStatusEnumTimedOut PostBillCreditNote200ApplicationJSONStatusEnum = "TimedOut"
)

type PostBillCreditNote200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostBillCreditNote200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostBillCreditNote200ApplicationJSONValidation struct {
	Errors   []PostBillCreditNote200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PostBillCreditNote200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PostBillCreditNote200ApplicationJSON struct {
	Changes           []PostBillCreditNote200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                                  `json:"companyId"`
	CompletedOnUtc    *time.Time                                              `json:"completedOnUtc,omitempty"`
	Data              *PostBillCreditNote200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                                  `json:"dataConnectionKey"`
	DataType          *string                                                 `json:"dataType,omitempty"`
	ErrorMessage      *string                                                 `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                  `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                               `json:"requestedOnUtc"`
	Status            PostBillCreditNote200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                     `json:"statusCode"`
	TimeoutInMinutes  *int                                                    `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                    `json:"timeoutInSeconds,omitempty"`
	Validation        *PostBillCreditNote200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PostBillCreditNoteResponse struct {
	ContentType                                string
	StatusCode                                 int
	PostBillCreditNote200ApplicationJSONObject *PostBillCreditNote200ApplicationJSON
}
