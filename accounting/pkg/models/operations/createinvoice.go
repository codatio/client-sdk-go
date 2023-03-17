package operations

import (
	"net/http"
	"time"
)

// CreateInvoiceSourceModifiedDateCustomerRef
// Reference to the customer the invoice has been issued to.
type CreateInvoiceSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

// CreateInvoiceSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type CreateInvoiceSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateInvoiceSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type CreateInvoiceSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateInvoiceSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type CreateInvoiceSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// CreateInvoiceSourceModifiedDateLineItemsTrackingCategoryRefs
// References a category against which the item is tracked.
type CreateInvoiceSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateInvoiceSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type CreateInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	CreateInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       CreateInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	CreateInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable CreateInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	CreateInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       CreateInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type CreateInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	CreateInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       CreateInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	CreateInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable CreateInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	CreateInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       CreateInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type CreateInvoiceSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateInvoiceSourceModifiedDateLineItemsTracking
// Categories, and a project and customer, against which the item is tracked.
type CreateInvoiceSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []CreateInvoiceSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *CreateInvoiceSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   CreateInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo CreateInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *CreateInvoiceSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type CreateInvoiceSourceModifiedDateLineItems struct {
	AccountRef           *CreateInvoiceSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                        `json:"description,omitempty"`
	DiscountAmount       *float64                                                       `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                       `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                                                          `json:"isDirectIncome,omitempty"`
	ItemRef              *CreateInvoiceSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                        `json:"quantity"`
	SubTotal             *float64                                                       `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                       `json:"taxAmount,omitempty"`
	TaxRateRef           *CreateInvoiceSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                       `json:"totalAmount,omitempty"`
	Tracking             *CreateInvoiceSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []CreateInvoiceSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                        `json:"unitAmount"`
}

type CreateInvoiceSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateInvoiceSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// CreateInvoiceSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type CreateInvoiceSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CreateInvoiceSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *CreateInvoiceSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                             `json:"currency,omitempty"`
	CurrencyRate *float64                                                            `json:"currencyRate,omitempty"`
	ID           *string                                                             `json:"id,omitempty"`
	Note         *string                                                             `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                          `json:"paidOnDate,omitempty"`
	Reference    *string                                                             `json:"reference,omitempty"`
	TotalAmount  *float64                                                            `json:"totalAmount,omitempty"`
}

type CreateInvoiceSourceModifiedDatePaymentAllocations struct {
	Allocation CreateInvoiceSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    CreateInvoiceSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type CreateInvoiceSourceModifiedDateStatusEnum string

const (
	CreateInvoiceSourceModifiedDateStatusEnumUnknown       CreateInvoiceSourceModifiedDateStatusEnum = "Unknown"
	CreateInvoiceSourceModifiedDateStatusEnumDraft         CreateInvoiceSourceModifiedDateStatusEnum = "Draft"
	CreateInvoiceSourceModifiedDateStatusEnumSubmitted     CreateInvoiceSourceModifiedDateStatusEnum = "Submitted"
	CreateInvoiceSourceModifiedDateStatusEnumPartiallyPaid CreateInvoiceSourceModifiedDateStatusEnum = "PartiallyPaid"
	CreateInvoiceSourceModifiedDateStatusEnumPaid          CreateInvoiceSourceModifiedDateStatusEnum = "Paid"
	CreateInvoiceSourceModifiedDateStatusEnumVoid          CreateInvoiceSourceModifiedDateStatusEnum = "Void"
)

// CreateInvoiceSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateInvoiceSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type CreateInvoiceSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// CreateInvoiceSourceModifiedDate
// > View the coverage for invoices in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// An invoice is an itemized record of goods sold or services provided to a [customer](https://docs.codat.io/accounting-api#/schemas/Customer).
//
// In Codat, an invoice contains details of:
//
// - The timeline of the invoice—when it was raised, marked as paid, last edited, and so on.
// - How much the invoice is for, what portion of the invoice is tax or discounts, and what currency the amounts are represented in.
// - Who the invoice has been raised to; the _customer_.
// - The breakdown of what the invoice is for; the _line items_.
// - Any [payments](https://docs.codat.io/accounting-api#/schemas/Payment) assigned to the invoice; the _payment allocations_.
//
// > **Invoices or bills?**
// >
// > In Codat, invoices represent accounts receivable only. For accounts payable invoices, see [Bills](https://docs.codat.io/accounting-api#/schemas/Bill).
//
// > **Invoice PDF downloads**
// >
// > You can <a className="external" href="https://api.codat.io/swagger/index.html#/Invoices/get_companies__companyId__data_invoices__invoiceId__pdf" target="_blank">download a PDF version</a> of an invoice for supported integrations.
// >
// > The filename will be invoice-{number}.pdf.
type CreateInvoiceSourceModifiedDate struct {
	AdditionalTaxAmount     *float64                                            `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                                            `json:"additionalTaxPercentage,omitempty"`
	AmountDue               float64                                             `json:"amountDue"`
	Currency                *string                                             `json:"currency,omitempty"`
	CurrencyRate            *float64                                            `json:"currencyRate,omitempty"`
	CustomerRef             *CreateInvoiceSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	DiscountPercentage      *float64                                            `json:"discountPercentage,omitempty"`
	DueDate                 *time.Time                                          `json:"dueDate,omitempty"`
	ID                      *string                                             `json:"id,omitempty"`
	InvoiceNumber           *string                                             `json:"invoiceNumber,omitempty"`
	IssueDate               time.Time                                           `json:"issueDate"`
	LineItems               []CreateInvoiceSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                *CreateInvoiceSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                                          `json:"modifiedDate,omitempty"`
	Note                    *string                                             `json:"note,omitempty"`
	PaidOnDate              *time.Time                                          `json:"paidOnDate,omitempty"`
	PaymentAllocations      []CreateInvoiceSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	SalesOrderRefs          []string                                            `json:"salesOrderRefs,omitempty"`
	SourceModifiedDate      *time.Time                                          `json:"sourceModifiedDate,omitempty"`
	Status                  CreateInvoiceSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal                *float64                                            `json:"subTotal,omitempty"`
	SupplementalData        *CreateInvoiceSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TotalAmount             float64                                             `json:"totalAmount"`
	TotalDiscount           *float64                                            `json:"totalDiscount,omitempty"`
	TotalTaxAmount          float64                                             `json:"totalTaxAmount"`
	WithholdingTax          []CreateInvoiceSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type CreateInvoiceRequest struct {
	RequestBody      *CreateInvoiceSourceModifiedDate `request:"mediaType=application/json"`
	CompanyID        string                           `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string                           `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes *int                             `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreateInvoice200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateInvoice200ApplicationJSONChangesTypeEnum string

const (
	CreateInvoice200ApplicationJSONChangesTypeEnumUnknown            CreateInvoice200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateInvoice200ApplicationJSONChangesTypeEnumCreated            CreateInvoice200ApplicationJSONChangesTypeEnum = "Created"
	CreateInvoice200ApplicationJSONChangesTypeEnumModified           CreateInvoice200ApplicationJSONChangesTypeEnum = "Modified"
	CreateInvoice200ApplicationJSONChangesTypeEnumDeleted            CreateInvoice200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateInvoice200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateInvoice200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type CreateInvoice200ApplicationJSONChanges struct {
	AttachmentID *string                                                       `json:"attachmentId,omitempty"`
	RecordRef    *CreateInvoice200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateInvoice200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// CreateInvoice200ApplicationJSONSourceModifiedDateCustomerRef
// Reference to the customer the invoice has been issued to.
type CreateInvoice200ApplicationJSONSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

// CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs
// References a category against which the item is tracked.
type CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTracking
// Categories, and a project and customer, against which the item is tracked.
type CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type CreateInvoice200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                          `json:"description,omitempty"`
	DiscountAmount       *float64                                                                         `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                         `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                                                                            `json:"isDirectIncome,omitempty"`
	ItemRef              *CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                          `json:"quantity"`
	SubTotal             *float64                                                                         `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                         `json:"taxAmount,omitempty"`
	TaxRateRef           *CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                         `json:"totalAmount,omitempty"`
	Tracking             *CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []CreateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                          `json:"unitAmount"`
}

type CreateInvoice200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// CreateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type CreateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CreateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *CreateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                               `json:"currency,omitempty"`
	CurrencyRate *float64                                                                              `json:"currencyRate,omitempty"`
	ID           *string                                                                               `json:"id,omitempty"`
	Note         *string                                                                               `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                            `json:"paidOnDate,omitempty"`
	Reference    *string                                                                               `json:"reference,omitempty"`
	TotalAmount  *float64                                                                              `json:"totalAmount,omitempty"`
}

type CreateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation CreateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    CreateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type CreateInvoice200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	CreateInvoice200ApplicationJSONSourceModifiedDateStatusEnumUnknown       CreateInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	CreateInvoice200ApplicationJSONSourceModifiedDateStatusEnumDraft         CreateInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "Draft"
	CreateInvoice200ApplicationJSONSourceModifiedDateStatusEnumSubmitted     CreateInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "Submitted"
	CreateInvoice200ApplicationJSONSourceModifiedDateStatusEnumPartiallyPaid CreateInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "PartiallyPaid"
	CreateInvoice200ApplicationJSONSourceModifiedDateStatusEnumPaid          CreateInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "Paid"
	CreateInvoice200ApplicationJSONSourceModifiedDateStatusEnumVoid          CreateInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "Void"
)

// CreateInvoice200ApplicationJSONSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateInvoice200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type CreateInvoice200ApplicationJSONSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// CreateInvoice200ApplicationJSONSourceModifiedDate
// > View the coverage for invoices in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// An invoice is an itemized record of goods sold or services provided to a [customer](https://docs.codat.io/accounting-api#/schemas/Customer).
//
// In Codat, an invoice contains details of:
//
// - The timeline of the invoice—when it was raised, marked as paid, last edited, and so on.
// - How much the invoice is for, what portion of the invoice is tax or discounts, and what currency the amounts are represented in.
// - Who the invoice has been raised to; the _customer_.
// - The breakdown of what the invoice is for; the _line items_.
// - Any [payments](https://docs.codat.io/accounting-api#/schemas/Payment) assigned to the invoice; the _payment allocations_.
//
// > **Invoices or bills?**
// >
// > In Codat, invoices represent accounts receivable only. For accounts payable invoices, see [Bills](https://docs.codat.io/accounting-api#/schemas/Bill).
//
// > **Invoice PDF downloads**
// >
// > You can <a className="external" href="https://api.codat.io/swagger/index.html#/Invoices/get_companies__companyId__data_invoices__invoiceId__pdf" target="_blank">download a PDF version</a> of an invoice for supported integrations.
// >
// > The filename will be invoice-{number}.pdf.
type CreateInvoice200ApplicationJSONSourceModifiedDate struct {
	AdditionalTaxAmount     *float64                                                              `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                                                              `json:"additionalTaxPercentage,omitempty"`
	AmountDue               float64                                                               `json:"amountDue"`
	Currency                *string                                                               `json:"currency,omitempty"`
	CurrencyRate            *float64                                                              `json:"currencyRate,omitempty"`
	CustomerRef             *CreateInvoice200ApplicationJSONSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	DiscountPercentage      *float64                                                              `json:"discountPercentage,omitempty"`
	DueDate                 *time.Time                                                            `json:"dueDate,omitempty"`
	ID                      *string                                                               `json:"id,omitempty"`
	InvoiceNumber           *string                                                               `json:"invoiceNumber,omitempty"`
	IssueDate               time.Time                                                             `json:"issueDate"`
	LineItems               []CreateInvoice200ApplicationJSONSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                *CreateInvoice200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                                                            `json:"modifiedDate,omitempty"`
	Note                    *string                                                               `json:"note,omitempty"`
	PaidOnDate              *time.Time                                                            `json:"paidOnDate,omitempty"`
	PaymentAllocations      []CreateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	SalesOrderRefs          []string                                                              `json:"salesOrderRefs,omitempty"`
	SourceModifiedDate      *time.Time                                                            `json:"sourceModifiedDate,omitempty"`
	Status                  CreateInvoice200ApplicationJSONSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal                *float64                                                              `json:"subTotal,omitempty"`
	SupplementalData        *CreateInvoice200ApplicationJSONSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TotalAmount             float64                                                               `json:"totalAmount"`
	TotalDiscount           *float64                                                              `json:"totalDiscount,omitempty"`
	TotalTaxAmount          float64                                                               `json:"totalTaxAmount"`
	WithholdingTax          []CreateInvoice200ApplicationJSONSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type CreateInvoice200ApplicationJSONStatusEnum string

const (
	CreateInvoice200ApplicationJSONStatusEnumPending  CreateInvoice200ApplicationJSONStatusEnum = "Pending"
	CreateInvoice200ApplicationJSONStatusEnumFailed   CreateInvoice200ApplicationJSONStatusEnum = "Failed"
	CreateInvoice200ApplicationJSONStatusEnumSuccess  CreateInvoice200ApplicationJSONStatusEnum = "Success"
	CreateInvoice200ApplicationJSONStatusEnumTimedOut CreateInvoice200ApplicationJSONStatusEnum = "TimedOut"
)

type CreateInvoice200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateInvoice200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateInvoice200ApplicationJSONValidation struct {
	Errors   []CreateInvoice200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateInvoice200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type CreateInvoice200ApplicationJSON struct {
	Changes           []CreateInvoice200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                             `json:"companyId"`
	CompletedOnUtc    *time.Time                                         `json:"completedOnUtc,omitempty"`
	Data              *CreateInvoice200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                             `json:"dataConnectionKey"`
	DataType          *string                                            `json:"dataType,omitempty"`
	ErrorMessage      *string                                            `json:"errorMessage,omitempty"`
	PushOperationKey  string                                             `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                          `json:"requestedOnUtc"`
	Status            CreateInvoice200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                `json:"statusCode"`
	TimeoutInMinutes  *int                                               `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                               `json:"timeoutInSeconds,omitempty"`
	Validation        *CreateInvoice200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type CreateInvoiceResponse struct {
	ContentType                           string
	StatusCode                            int
	RawResponse                           *http.Response
	CreateInvoice200ApplicationJSONObject *CreateInvoice200ApplicationJSON
}
