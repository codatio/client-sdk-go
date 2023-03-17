package operations

import (
	"net/http"
	"time"
)

// UpdateBillCreditNoteSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type UpdateBillCreditNoteSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// UpdateBillCreditNoteSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type UpdateBillCreditNoteSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// UpdateBillCreditNoteSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type UpdateBillCreditNoteSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs
// References a category against which the item is tracked.
type UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumCustomer      UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Customer"
	UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumCustomer      UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Customer"
	UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// UpdateBillCreditNoteSourceModifiedDateLineItemsTracking
// Categories, and a project and customer, against which the item is tracked.
type UpdateBillCreditNoteSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type UpdateBillCreditNoteSourceModifiedDateLineItems struct {
	AccountRef           *UpdateBillCreditNoteSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                               `json:"description,omitempty"`
	DiscountAmount       *float64                                                              `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                              `json:"discountPercentage,omitempty"`
	ItemRef              *UpdateBillCreditNoteSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                               `json:"quantity"`
	SubTotal             *float64                                                              `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                              `json:"taxAmount,omitempty"`
	TaxRateRef           *UpdateBillCreditNoteSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                              `json:"totalAmount,omitempty"`
	Tracking             *UpdateBillCreditNoteSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []UpdateBillCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                               `json:"unitAmount"`
}

type UpdateBillCreditNoteSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type UpdateBillCreditNoteSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// UpdateBillCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type UpdateBillCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type UpdateBillCreditNoteSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *UpdateBillCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                    `json:"currency,omitempty"`
	CurrencyRate *float64                                                                   `json:"currencyRate,omitempty"`
	ID           *string                                                                    `json:"id,omitempty"`
	Note         *string                                                                    `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                 `json:"paidOnDate,omitempty"`
	Reference    *string                                                                    `json:"reference,omitempty"`
	TotalAmount  *float64                                                                   `json:"totalAmount,omitempty"`
}

type UpdateBillCreditNoteSourceModifiedDatePaymentAllocations struct {
	Allocation UpdateBillCreditNoteSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    UpdateBillCreditNoteSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type UpdateBillCreditNoteSourceModifiedDateStatusEnum string

const (
	UpdateBillCreditNoteSourceModifiedDateStatusEnumUnknown       UpdateBillCreditNoteSourceModifiedDateStatusEnum = "Unknown"
	UpdateBillCreditNoteSourceModifiedDateStatusEnumDraft         UpdateBillCreditNoteSourceModifiedDateStatusEnum = "Draft"
	UpdateBillCreditNoteSourceModifiedDateStatusEnumSubmitted     UpdateBillCreditNoteSourceModifiedDateStatusEnum = "Submitted"
	UpdateBillCreditNoteSourceModifiedDateStatusEnumPaid          UpdateBillCreditNoteSourceModifiedDateStatusEnum = "Paid"
	UpdateBillCreditNoteSourceModifiedDateStatusEnumVoid          UpdateBillCreditNoteSourceModifiedDateStatusEnum = "Void"
	UpdateBillCreditNoteSourceModifiedDateStatusEnumPartiallyPaid UpdateBillCreditNoteSourceModifiedDateStatusEnum = "PartiallyPaid"
)

// UpdateBillCreditNoteSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type UpdateBillCreditNoteSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// UpdateBillCreditNoteSourceModifiedDateSupplierRef
// Supplier that issued the bill credit note.
type UpdateBillCreditNoteSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type UpdateBillCreditNoteSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// UpdateBillCreditNoteSourceModifiedDate
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
type UpdateBillCreditNoteSourceModifiedDate struct {
	AllocatedOnDate      *time.Time                                                 `json:"allocatedOnDate,omitempty"`
	BillCreditNoteNumber *string                                                    `json:"billCreditNoteNumber,omitempty"`
	Currency             *string                                                    `json:"currency,omitempty"`
	CurrencyRate         *float64                                                   `json:"currencyRate,omitempty"`
	DiscountPercentage   float64                                                    `json:"discountPercentage"`
	ID                   *string                                                    `json:"id,omitempty"`
	IssueDate            *time.Time                                                 `json:"issueDate,omitempty"`
	LineItems            []UpdateBillCreditNoteSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata             *UpdateBillCreditNoteSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                                 `json:"modifiedDate,omitempty"`
	Note                 *string                                                    `json:"note,omitempty"`
	PaymentAllocations   []UpdateBillCreditNoteSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	RemainingCredit      float64                                                    `json:"remainingCredit"`
	SourceModifiedDate   *time.Time                                                 `json:"sourceModifiedDate,omitempty"`
	Status               UpdateBillCreditNoteSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal             float64                                                    `json:"subTotal"`
	SupplementalData     *UpdateBillCreditNoteSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef          *UpdateBillCreditNoteSourceModifiedDateSupplierRef         `json:"supplierRef,omitempty"`
	TotalAmount          float64                                                    `json:"totalAmount"`
	TotalDiscount        float64                                                    `json:"totalDiscount"`
	TotalTaxAmount       float64                                                    `json:"totalTaxAmount"`
	WithholdingTax       []UpdateBillCreditNoteSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type UpdateBillCreditNoteRequest struct {
	RequestBody      *UpdateBillCreditNoteSourceModifiedDate `request:"mediaType=application/json"`
	BillCreditNoteID string                                  `pathParam:"style=simple,explode=false,name=billCreditNoteId"`
	CompanyID        string                                  `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string                                  `pathParam:"style=simple,explode=false,name=connectionId"`
	ForceUpdate      *bool                                   `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int                                    `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type UpdateBillCreditNote200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type UpdateBillCreditNote200ApplicationJSONChangesTypeEnum string

const (
	UpdateBillCreditNote200ApplicationJSONChangesTypeEnumUnknown            UpdateBillCreditNote200ApplicationJSONChangesTypeEnum = "Unknown"
	UpdateBillCreditNote200ApplicationJSONChangesTypeEnumCreated            UpdateBillCreditNote200ApplicationJSONChangesTypeEnum = "Created"
	UpdateBillCreditNote200ApplicationJSONChangesTypeEnumModified           UpdateBillCreditNote200ApplicationJSONChangesTypeEnum = "Modified"
	UpdateBillCreditNote200ApplicationJSONChangesTypeEnumDeleted            UpdateBillCreditNote200ApplicationJSONChangesTypeEnum = "Deleted"
	UpdateBillCreditNote200ApplicationJSONChangesTypeEnumAttachmentUploaded UpdateBillCreditNote200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type UpdateBillCreditNote200ApplicationJSONChanges struct {
	AttachmentID *string                                                              `json:"attachmentId,omitempty"`
	RecordRef    *UpdateBillCreditNote200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *UpdateBillCreditNote200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs
// References a category against which the item is tracked.
type UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumCustomer      UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Customer"
	UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumCustomer      UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Customer"
	UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTracking
// Categories, and a project and customer, against which the item is tracked.
type UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                                 `json:"description,omitempty"`
	DiscountAmount       *float64                                                                                `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                                `json:"discountPercentage,omitempty"`
	ItemRef              *UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                                 `json:"quantity"`
	SubTotal             *float64                                                                                `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                                `json:"taxAmount,omitempty"`
	TaxRateRef           *UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                                `json:"totalAmount,omitempty"`
	Tracking             *UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                                 `json:"unitAmount"`
}

type UpdateBillCreditNote200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type UpdateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// UpdateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type UpdateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type UpdateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *UpdateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                                      `json:"currency,omitempty"`
	CurrencyRate *float64                                                                                     `json:"currencyRate,omitempty"`
	ID           *string                                                                                      `json:"id,omitempty"`
	Note         *string                                                                                      `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                                   `json:"paidOnDate,omitempty"`
	Reference    *string                                                                                      `json:"reference,omitempty"`
	TotalAmount  *float64                                                                                     `json:"totalAmount,omitempty"`
}

type UpdateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation UpdateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    UpdateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type UpdateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	UpdateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumUnknown       UpdateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	UpdateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumDraft         UpdateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Draft"
	UpdateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumSubmitted     UpdateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Submitted"
	UpdateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumPaid          UpdateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Paid"
	UpdateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumVoid          UpdateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Void"
	UpdateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnumPartiallyPaid UpdateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "PartiallyPaid"
)

// UpdateBillCreditNote200ApplicationJSONSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type UpdateBillCreditNote200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// UpdateBillCreditNote200ApplicationJSONSourceModifiedDateSupplierRef
// Supplier that issued the bill credit note.
type UpdateBillCreditNote200ApplicationJSONSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type UpdateBillCreditNote200ApplicationJSONSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// UpdateBillCreditNote200ApplicationJSONSourceModifiedDate
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
type UpdateBillCreditNote200ApplicationJSONSourceModifiedDate struct {
	AllocatedOnDate      *time.Time                                                                   `json:"allocatedOnDate,omitempty"`
	BillCreditNoteNumber *string                                                                      `json:"billCreditNoteNumber,omitempty"`
	Currency             *string                                                                      `json:"currency,omitempty"`
	CurrencyRate         *float64                                                                     `json:"currencyRate,omitempty"`
	DiscountPercentage   float64                                                                      `json:"discountPercentage"`
	ID                   *string                                                                      `json:"id,omitempty"`
	IssueDate            *time.Time                                                                   `json:"issueDate,omitempty"`
	LineItems            []UpdateBillCreditNote200ApplicationJSONSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata             *UpdateBillCreditNote200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                                                   `json:"modifiedDate,omitempty"`
	Note                 *string                                                                      `json:"note,omitempty"`
	PaymentAllocations   []UpdateBillCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	RemainingCredit      float64                                                                      `json:"remainingCredit"`
	SourceModifiedDate   *time.Time                                                                   `json:"sourceModifiedDate,omitempty"`
	Status               UpdateBillCreditNote200ApplicationJSONSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal             float64                                                                      `json:"subTotal"`
	SupplementalData     *UpdateBillCreditNote200ApplicationJSONSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef          *UpdateBillCreditNote200ApplicationJSONSourceModifiedDateSupplierRef         `json:"supplierRef,omitempty"`
	TotalAmount          float64                                                                      `json:"totalAmount"`
	TotalDiscount        float64                                                                      `json:"totalDiscount"`
	TotalTaxAmount       float64                                                                      `json:"totalTaxAmount"`
	WithholdingTax       []UpdateBillCreditNote200ApplicationJSONSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type UpdateBillCreditNote200ApplicationJSONStatusEnum string

const (
	UpdateBillCreditNote200ApplicationJSONStatusEnumPending  UpdateBillCreditNote200ApplicationJSONStatusEnum = "Pending"
	UpdateBillCreditNote200ApplicationJSONStatusEnumFailed   UpdateBillCreditNote200ApplicationJSONStatusEnum = "Failed"
	UpdateBillCreditNote200ApplicationJSONStatusEnumSuccess  UpdateBillCreditNote200ApplicationJSONStatusEnum = "Success"
	UpdateBillCreditNote200ApplicationJSONStatusEnumTimedOut UpdateBillCreditNote200ApplicationJSONStatusEnum = "TimedOut"
)

type UpdateBillCreditNote200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// UpdateBillCreditNote200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type UpdateBillCreditNote200ApplicationJSONValidation struct {
	Errors   []UpdateBillCreditNote200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []UpdateBillCreditNote200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type UpdateBillCreditNote200ApplicationJSON struct {
	Changes           []UpdateBillCreditNote200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                                    `json:"companyId"`
	CompletedOnUtc    *time.Time                                                `json:"completedOnUtc,omitempty"`
	Data              *UpdateBillCreditNote200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                                    `json:"dataConnectionKey"`
	DataType          *string                                                   `json:"dataType,omitempty"`
	ErrorMessage      *string                                                   `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                    `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                                 `json:"requestedOnUtc"`
	Status            UpdateBillCreditNote200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                       `json:"statusCode"`
	TimeoutInMinutes  *int                                                      `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                      `json:"timeoutInSeconds,omitempty"`
	Validation        *UpdateBillCreditNote200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type UpdateBillCreditNoteResponse struct {
	ContentType                                  string
	StatusCode                                   int
	RawResponse                                  *http.Response
	UpdateBillCreditNote200ApplicationJSONObject *UpdateBillCreditNote200ApplicationJSON
}
