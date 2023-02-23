package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type UpdateInvoicePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	InvoiceID    string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type UpdateInvoiceQueryParams struct {
	ForceUpdate      *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// UpdateInvoiceSourceModifiedDateCustomerRef
// Reference to the customer the invoice has been issued to.
type UpdateInvoiceSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

// UpdateInvoiceSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type UpdateInvoiceSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// UpdateInvoiceSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type UpdateInvoiceSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// UpdateInvoiceSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type UpdateInvoiceSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type UpdateInvoiceSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type UpdateInvoiceSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type UpdateInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	UpdateInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       UpdateInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	UpdateInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable UpdateInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	UpdateInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       UpdateInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type UpdateInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	UpdateInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       UpdateInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	UpdateInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable UpdateInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	UpdateInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       UpdateInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type UpdateInvoiceSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type UpdateInvoiceSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []UpdateInvoiceSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *UpdateInvoiceSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   UpdateInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo UpdateInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *UpdateInvoiceSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type UpdateInvoiceSourceModifiedDateLineItems struct {
	AccountRef           *UpdateInvoiceSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                        `json:"description,omitempty"`
	DiscountAmount       *float64                                                       `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                       `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                                                          `json:"isDirectIncome,omitempty"`
	ItemRef              *UpdateInvoiceSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                        `json:"quantity"`
	SubTotal             *float64                                                       `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                       `json:"taxAmount,omitempty"`
	TaxRateRef           *UpdateInvoiceSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                       `json:"totalAmount,omitempty"`
	Tracking             *UpdateInvoiceSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []UpdateInvoiceSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                        `json:"unitAmount"`
}

type UpdateInvoiceSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type UpdateInvoiceSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// UpdateInvoiceSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type UpdateInvoiceSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type UpdateInvoiceSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *UpdateInvoiceSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                             `json:"currency,omitempty"`
	CurrencyRate *float64                                                            `json:"currencyRate,omitempty"`
	ID           *string                                                             `json:"id,omitempty"`
	Note         *string                                                             `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                          `json:"paidOnDate,omitempty"`
	Reference    *string                                                             `json:"reference,omitempty"`
	TotalAmount  *float64                                                            `json:"totalAmount,omitempty"`
}

type UpdateInvoiceSourceModifiedDatePaymentAllocations struct {
	Allocation UpdateInvoiceSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    UpdateInvoiceSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type UpdateInvoiceSourceModifiedDateStatusEnum string

const (
	UpdateInvoiceSourceModifiedDateStatusEnumUnknown       UpdateInvoiceSourceModifiedDateStatusEnum = "Unknown"
	UpdateInvoiceSourceModifiedDateStatusEnumDraft         UpdateInvoiceSourceModifiedDateStatusEnum = "Draft"
	UpdateInvoiceSourceModifiedDateStatusEnumSubmitted     UpdateInvoiceSourceModifiedDateStatusEnum = "Submitted"
	UpdateInvoiceSourceModifiedDateStatusEnumPartiallyPaid UpdateInvoiceSourceModifiedDateStatusEnum = "PartiallyPaid"
	UpdateInvoiceSourceModifiedDateStatusEnumPaid          UpdateInvoiceSourceModifiedDateStatusEnum = "Paid"
	UpdateInvoiceSourceModifiedDateStatusEnumVoid          UpdateInvoiceSourceModifiedDateStatusEnum = "Void"
)

type UpdateInvoiceSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type UpdateInvoiceSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// UpdateInvoiceSourceModifiedDate
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
type UpdateInvoiceSourceModifiedDate struct {
	AdditionalTaxAmount     *float64                                            `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                                            `json:"additionalTaxPercentage,omitempty"`
	AmountDue               float64                                             `json:"amountDue"`
	Currency                *string                                             `json:"currency,omitempty"`
	CurrencyRate            *float64                                            `json:"currencyRate,omitempty"`
	CustomerRef             *UpdateInvoiceSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	DiscountPercentage      *float64                                            `json:"discountPercentage,omitempty"`
	DueDate                 *time.Time                                          `json:"dueDate,omitempty"`
	ID                      *string                                             `json:"id,omitempty"`
	InvoiceNumber           *string                                             `json:"invoiceNumber,omitempty"`
	IssueDate               time.Time                                           `json:"issueDate"`
	LineItems               []UpdateInvoiceSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                *UpdateInvoiceSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                                          `json:"modifiedDate,omitempty"`
	Note                    *string                                             `json:"note,omitempty"`
	PaidOnDate              *time.Time                                          `json:"paidOnDate,omitempty"`
	PaymentAllocations      []UpdateInvoiceSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	SalesOrderRefs          []string                                            `json:"salesOrderRefs,omitempty"`
	SourceModifiedDate      *time.Time                                          `json:"sourceModifiedDate,omitempty"`
	Status                  UpdateInvoiceSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal                *float64                                            `json:"subTotal,omitempty"`
	SupplementalData        *UpdateInvoiceSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TotalAmount             float64                                             `json:"totalAmount"`
	TotalDiscount           *float64                                            `json:"totalDiscount,omitempty"`
	TotalTaxAmount          float64                                             `json:"totalTaxAmount"`
	WithholdingTax          []UpdateInvoiceSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type UpdateInvoiceSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateInvoiceRequest struct {
	PathParams  UpdateInvoicePathParams
	QueryParams UpdateInvoiceQueryParams
	Request     *UpdateInvoiceSourceModifiedDate `request:"mediaType=application/json"`
	Security    UpdateInvoiceSecurity
}

type UpdateInvoice200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type UpdateInvoice200ApplicationJSONChangesTypeEnum string

const (
	UpdateInvoice200ApplicationJSONChangesTypeEnumUnknown            UpdateInvoice200ApplicationJSONChangesTypeEnum = "Unknown"
	UpdateInvoice200ApplicationJSONChangesTypeEnumCreated            UpdateInvoice200ApplicationJSONChangesTypeEnum = "Created"
	UpdateInvoice200ApplicationJSONChangesTypeEnumModified           UpdateInvoice200ApplicationJSONChangesTypeEnum = "Modified"
	UpdateInvoice200ApplicationJSONChangesTypeEnumDeleted            UpdateInvoice200ApplicationJSONChangesTypeEnum = "Deleted"
	UpdateInvoice200ApplicationJSONChangesTypeEnumAttachmentUploaded UpdateInvoice200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type UpdateInvoice200ApplicationJSONChanges struct {
	AttachmentID *string                                                       `json:"attachmentId,omitempty"`
	RecordRef    *UpdateInvoice200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *UpdateInvoice200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// UpdateInvoice200ApplicationJSONSourceModifiedDateCustomerRef
// Reference to the customer the invoice has been issued to.
type UpdateInvoice200ApplicationJSONSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

// UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type UpdateInvoice200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                          `json:"description,omitempty"`
	DiscountAmount       *float64                                                                         `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                         `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                                                                            `json:"isDirectIncome,omitempty"`
	ItemRef              *UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                          `json:"quantity"`
	SubTotal             *float64                                                                         `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                         `json:"taxAmount,omitempty"`
	TaxRateRef           *UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                         `json:"totalAmount,omitempty"`
	Tracking             *UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []UpdateInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                          `json:"unitAmount"`
}

type UpdateInvoice200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type UpdateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// UpdateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type UpdateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type UpdateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *UpdateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                               `json:"currency,omitempty"`
	CurrencyRate *float64                                                                              `json:"currencyRate,omitempty"`
	ID           *string                                                                               `json:"id,omitempty"`
	Note         *string                                                                               `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                            `json:"paidOnDate,omitempty"`
	Reference    *string                                                                               `json:"reference,omitempty"`
	TotalAmount  *float64                                                                              `json:"totalAmount,omitempty"`
}

type UpdateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation UpdateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    UpdateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type UpdateInvoice200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	UpdateInvoice200ApplicationJSONSourceModifiedDateStatusEnumUnknown       UpdateInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	UpdateInvoice200ApplicationJSONSourceModifiedDateStatusEnumDraft         UpdateInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "Draft"
	UpdateInvoice200ApplicationJSONSourceModifiedDateStatusEnumSubmitted     UpdateInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "Submitted"
	UpdateInvoice200ApplicationJSONSourceModifiedDateStatusEnumPartiallyPaid UpdateInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "PartiallyPaid"
	UpdateInvoice200ApplicationJSONSourceModifiedDateStatusEnumPaid          UpdateInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "Paid"
	UpdateInvoice200ApplicationJSONSourceModifiedDateStatusEnumVoid          UpdateInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "Void"
)

type UpdateInvoice200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type UpdateInvoice200ApplicationJSONSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// UpdateInvoice200ApplicationJSONSourceModifiedDate
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
type UpdateInvoice200ApplicationJSONSourceModifiedDate struct {
	AdditionalTaxAmount     *float64                                                              `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                                                              `json:"additionalTaxPercentage,omitempty"`
	AmountDue               float64                                                               `json:"amountDue"`
	Currency                *string                                                               `json:"currency,omitempty"`
	CurrencyRate            *float64                                                              `json:"currencyRate,omitempty"`
	CustomerRef             *UpdateInvoice200ApplicationJSONSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	DiscountPercentage      *float64                                                              `json:"discountPercentage,omitempty"`
	DueDate                 *time.Time                                                            `json:"dueDate,omitempty"`
	ID                      *string                                                               `json:"id,omitempty"`
	InvoiceNumber           *string                                                               `json:"invoiceNumber,omitempty"`
	IssueDate               time.Time                                                             `json:"issueDate"`
	LineItems               []UpdateInvoice200ApplicationJSONSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                *UpdateInvoice200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                                                            `json:"modifiedDate,omitempty"`
	Note                    *string                                                               `json:"note,omitempty"`
	PaidOnDate              *time.Time                                                            `json:"paidOnDate,omitempty"`
	PaymentAllocations      []UpdateInvoice200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	SalesOrderRefs          []string                                                              `json:"salesOrderRefs,omitempty"`
	SourceModifiedDate      *time.Time                                                            `json:"sourceModifiedDate,omitempty"`
	Status                  UpdateInvoice200ApplicationJSONSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal                *float64                                                              `json:"subTotal,omitempty"`
	SupplementalData        *UpdateInvoice200ApplicationJSONSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TotalAmount             float64                                                               `json:"totalAmount"`
	TotalDiscount           *float64                                                              `json:"totalDiscount,omitempty"`
	TotalTaxAmount          float64                                                               `json:"totalTaxAmount"`
	WithholdingTax          []UpdateInvoice200ApplicationJSONSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type UpdateInvoice200ApplicationJSONStatusEnum string

const (
	UpdateInvoice200ApplicationJSONStatusEnumPending  UpdateInvoice200ApplicationJSONStatusEnum = "Pending"
	UpdateInvoice200ApplicationJSONStatusEnumFailed   UpdateInvoice200ApplicationJSONStatusEnum = "Failed"
	UpdateInvoice200ApplicationJSONStatusEnumSuccess  UpdateInvoice200ApplicationJSONStatusEnum = "Success"
	UpdateInvoice200ApplicationJSONStatusEnumTimedOut UpdateInvoice200ApplicationJSONStatusEnum = "TimedOut"
)

type UpdateInvoice200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// UpdateInvoice200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type UpdateInvoice200ApplicationJSONValidation struct {
	Errors   []UpdateInvoice200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []UpdateInvoice200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type UpdateInvoice200ApplicationJSON struct {
	Changes           []UpdateInvoice200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                             `json:"companyId"`
	CompletedOnUtc    *time.Time                                         `json:"completedOnUtc,omitempty"`
	Data              *UpdateInvoice200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                             `json:"dataConnectionKey"`
	DataType          *string                                            `json:"dataType,omitempty"`
	ErrorMessage      *string                                            `json:"errorMessage,omitempty"`
	PushOperationKey  string                                             `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                          `json:"requestedOnUtc"`
	Status            UpdateInvoice200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                `json:"statusCode"`
	TimeoutInMinutes  *int                                               `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                               `json:"timeoutInSeconds,omitempty"`
	Validation        *UpdateInvoice200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type UpdateInvoiceResponse struct {
	ContentType                           string
	StatusCode                            int
	UpdateInvoice200ApplicationJSONObject *UpdateInvoice200ApplicationJSON
}
