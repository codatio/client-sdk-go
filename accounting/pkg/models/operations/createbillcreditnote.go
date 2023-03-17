package operations

import (
	"net/http"
	"time"
)

type CreateBillCreditNotePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type CreateBillCreditNoteQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// CreateBillCreditNoteSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type CreateBillCreditNoteSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateBillCreditNoteSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type CreateBillCreditNoteSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateBillCreditNoteSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type CreateBillCreditNoteSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// CreateBillCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs
// References a category against which the item is tracked.
type CreateBillCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateBillCreditNoteSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumCustomer      CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Customer"
	CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumCustomer      CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Customer"
	CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type CreateBillCreditNoteSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateBillCreditNoteSourceModifiedDateLineItemsTracking
// Categories, and a project and customer, against which the item is tracked.
type CreateBillCreditNoteSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []CreateBillCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *CreateBillCreditNoteSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo CreateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *CreateBillCreditNoteSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type CreateBillCreditNoteSourceModifiedDateLineItems struct {
	AccountRef           *CreateBillCreditNoteSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                               `json:"description,omitempty"`
	DiscountAmount       *float64                                                              `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                              `json:"discountPercentage,omitempty"`
	ItemRef              *CreateBillCreditNoteSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                               `json:"quantity"`
	SubTotal             *float64                                                              `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                              `json:"taxAmount,omitempty"`
	TaxRateRef           *CreateBillCreditNoteSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                              `json:"totalAmount,omitempty"`
	Tracking             *CreateBillCreditNoteSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []CreateBillCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                               `json:"unitAmount"`
}

type CreateBillCreditNoteSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateBillCreditNoteSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// CreateBillCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type CreateBillCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CreateBillCreditNoteSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *CreateBillCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                    `json:"currency,omitempty"`
	CurrencyRate *float64                                                                   `json:"currencyRate,omitempty"`
	ID           *string                                                                    `json:"id,omitempty"`
	Note         *string                                                                    `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                 `json:"paidOnDate,omitempty"`
	Reference    *string                                                                    `json:"reference,omitempty"`
	TotalAmount  *float64                                                                   `json:"totalAmount,omitempty"`
}

type CreateBillCreditNoteSourceModifiedDatePaymentAllocations struct {
	Allocation CreateBillCreditNoteSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    CreateBillCreditNoteSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type CreateBillCreditNoteSourceModifiedDateStatusEnum string

const (
	CreateBillCreditNoteSourceModifiedDateStatusEnumUnknown       CreateBillCreditNoteSourceModifiedDateStatusEnum = "Unknown"
	CreateBillCreditNoteSourceModifiedDateStatusEnumDraft         CreateBillCreditNoteSourceModifiedDateStatusEnum = "Draft"
	CreateBillCreditNoteSourceModifiedDateStatusEnumSubmitted     CreateBillCreditNoteSourceModifiedDateStatusEnum = "Submitted"
	CreateBillCreditNoteSourceModifiedDateStatusEnumPaid          CreateBillCreditNoteSourceModifiedDateStatusEnum = "Paid"
	CreateBillCreditNoteSourceModifiedDateStatusEnumVoid          CreateBillCreditNoteSourceModifiedDateStatusEnum = "Void"
	CreateBillCreditNoteSourceModifiedDateStatusEnumPartiallyPaid CreateBillCreditNoteSourceModifiedDateStatusEnum = "PartiallyPaid"
)

// CreateBillCreditNoteSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateBillCreditNoteSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateBillCreditNoteSourceModifiedDateSupplierRef
// Supplier that issued the bill credit note.
type CreateBillCreditNoteSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type CreateBillCreditNoteSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// CreateBillCreditNoteSourceModifiedDate
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
type CreateBillCreditNoteSourceModifiedDate struct {
	AllocatedOnDate      *time.Time                                                 `json:"allocatedOnDate,omitempty"`
	BillCreditNoteNumber *string                                                    `json:"billCreditNoteNumber,omitempty"`
	Currency             *string                                                    `json:"currency,omitempty"`
	CurrencyRate         *float64                                                   `json:"currencyRate,omitempty"`
	DiscountPercentage   float64                                                    `json:"discountPercentage"`
	ID                   *string                                                    `json:"id,omitempty"`
	IssueDate            *time.Time                                                 `json:"issueDate,omitempty"`
	LineItems            []CreateBillCreditNoteSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata             *CreateBillCreditNoteSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                                 `json:"modifiedDate,omitempty"`
	Note                 *string                                                    `json:"note,omitempty"`
	PaymentAllocations   []CreateBillCreditNoteSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	RemainingCredit      float64                                                    `json:"remainingCredit"`
	SourceModifiedDate   *time.Time                                                 `json:"sourceModifiedDate,omitempty"`
	Status               CreateBillCreditNoteSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal             float64                                                    `json:"subTotal"`
	SupplementalData     *CreateBillCreditNoteSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef          *CreateBillCreditNoteSourceModifiedDateSupplierRef         `json:"supplierRef,omitempty"`
	TotalAmount          float64                                                    `json:"totalAmount"`
	TotalDiscount        float64                                                    `json:"totalDiscount"`
	TotalTaxAmount       float64                                                    `json:"totalTaxAmount"`
	WithholdingTax       []CreateBillCreditNoteSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type CreateBillCreditNoteRequest struct {
	PathParams  CreateBillCreditNotePathParams
	QueryParams CreateBillCreditNoteQueryParams
	Request     *CreateBillCreditNoteSourceModifiedDate `request:"mediaType=application/json"`
}

type CreateBillCreditNote200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateBillCreditNote200ApplicationJSONChangesTypeEnum string

const (
	CreateBillCreditNote200ApplicationJSONChangesTypeEnumUnknown            CreateBillCreditNote200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateBillCreditNote200ApplicationJSONChangesTypeEnumCreated            CreateBillCreditNote200ApplicationJSONChangesTypeEnum = "Created"
	CreateBillCreditNote200ApplicationJSONChangesTypeEnumModified           CreateBillCreditNote200ApplicationJSONChangesTypeEnum = "Modified"
	CreateBillCreditNote200ApplicationJSONChangesTypeEnumDeleted            CreateBillCreditNote200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateBillCreditNote200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateBillCreditNote200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type CreateBillCreditNote200ApplicationJSONChanges struct {
	AttachmentID *string                                                              `json:"attachmentId,omitempty"`
	RecordRef    *CreateBillCreditNote200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateBillCreditNote200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs
// References a category against which the item is tracked.
type CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumCustomer      CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Customer"
	CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumCustomer      CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Customer"
	CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTracking
// Categories, and a project and customer, against which the item is tracked.
type CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                                 `json:"description,omitempty"`
	DiscountAmount       *float64                                                                                `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                                `json:"discountPercentage,omitempty"`
	ItemRef              *CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                                 `json:"quantity"`
	SubTotal             *float64                                                                                `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                                `json:"taxAmount,omitempty"`
	TaxRateRef           *CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                                `json:"totalAmount,omitempty"`
	Tracking             *CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                                 `json:"unitAmount"`
}

type CreateBillCreditNote200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// CreateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type CreateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CreateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *CreateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                                      `json:"currency,omitempty"`
	CurrencyRate *float64                                                                                     `json:"currencyRate,omitempty"`
	ID           *string                                                                                      `json:"id,omitempty"`
	Note         *string                                                                                      `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                                   `json:"paidOnDate,omitempty"`
	Reference    *string                                                                                      `json:"reference,omitempty"`
	TotalAmount  *float64                                                                                     `json:"totalAmount,omitempty"`
}

type CreateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation CreateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    CreateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type CreateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	CreateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumUnknown       CreateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	CreateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumDraft         CreateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Draft"
	CreateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumSubmitted     CreateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Submitted"
	CreateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumPaid          CreateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Paid"
	CreateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumVoid          CreateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Void"
	CreateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumPartiallyPaid CreateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "PartiallyPaid"
)

// CreateBillCreditNote200ApplicationJSONSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateBillCreditNote200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateBillCreditNote200ApplicationJSONSourceModifiedDateSupplierRef
// Supplier that issued the bill credit note.
type CreateBillCreditNote200ApplicationJSONSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type CreateBillCreditNote200ApplicationJSONSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// CreateBillCreditNote200ApplicationJSONSourceModifiedDate
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
type CreateBillCreditNote200ApplicationJSONSourceModifiedDate struct {
	AllocatedOnDate      *time.Time                                                                   `json:"allocatedOnDate,omitempty"`
	BillCreditNoteNumber *string                                                                      `json:"billCreditNoteNumber,omitempty"`
	Currency             *string                                                                      `json:"currency,omitempty"`
	CurrencyRate         *float64                                                                     `json:"currencyRate,omitempty"`
	DiscountPercentage   float64                                                                      `json:"discountPercentage"`
	ID                   *string                                                                      `json:"id,omitempty"`
	IssueDate            *time.Time                                                                   `json:"issueDate,omitempty"`
	LineItems            []CreateBillCreditNote200ApplicationJSONSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata             *CreateBillCreditNote200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                                                   `json:"modifiedDate,omitempty"`
	Note                 *string                                                                      `json:"note,omitempty"`
	PaymentAllocations   []CreateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	RemainingCredit      float64                                                                      `json:"remainingCredit"`
	SourceModifiedDate   *time.Time                                                                   `json:"sourceModifiedDate,omitempty"`
	Status               CreateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal             float64                                                                      `json:"subTotal"`
	SupplementalData     *CreateBillCreditNote200ApplicationJSONSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef          *CreateBillCreditNote200ApplicationJSONSourceModifiedDateSupplierRef         `json:"supplierRef,omitempty"`
	TotalAmount          float64                                                                      `json:"totalAmount"`
	TotalDiscount        float64                                                                      `json:"totalDiscount"`
	TotalTaxAmount       float64                                                                      `json:"totalTaxAmount"`
	WithholdingTax       []CreateBillCreditNote200ApplicationJSONSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type CreateBillCreditNote200ApplicationJSONStatusEnum string

const (
	CreateBillCreditNote200ApplicationJSONStatusEnumPending  CreateBillCreditNote200ApplicationJSONStatusEnum = "Pending"
	CreateBillCreditNote200ApplicationJSONStatusEnumFailed   CreateBillCreditNote200ApplicationJSONStatusEnum = "Failed"
	CreateBillCreditNote200ApplicationJSONStatusEnumSuccess  CreateBillCreditNote200ApplicationJSONStatusEnum = "Success"
	CreateBillCreditNote200ApplicationJSONStatusEnumTimedOut CreateBillCreditNote200ApplicationJSONStatusEnum = "TimedOut"
)

type CreateBillCreditNote200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateBillCreditNote200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateBillCreditNote200ApplicationJSONValidation struct {
	Errors   []CreateBillCreditNote200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateBillCreditNote200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type CreateBillCreditNote200ApplicationJSON struct {
	Changes           []CreateBillCreditNote200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                                    `json:"companyId"`
	CompletedOnUtc    *time.Time                                                `json:"completedOnUtc,omitempty"`
	Data              *CreateBillCreditNote200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                                    `json:"dataConnectionKey"`
	DataType          *string                                                   `json:"dataType,omitempty"`
	ErrorMessage      *string                                                   `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                    `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                                 `json:"requestedOnUtc"`
	Status            CreateBillCreditNote200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                       `json:"statusCode"`
	TimeoutInMinutes  *int                                                      `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                      `json:"timeoutInSeconds,omitempty"`
	Validation        *CreateBillCreditNote200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type CreateBillCreditNoteResponse struct {
	ContentType                                  string
	StatusCode                                   int
	RawResponse                                  *http.Response
	CreateBillCreditNote200ApplicationJSONObject *CreateBillCreditNote200ApplicationJSON
}
