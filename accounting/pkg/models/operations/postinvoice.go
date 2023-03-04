package operations

import (
	"net/http"
	"time"
)

type PostInvoicePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostInvoiceQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// PostInvoiceSourceModifiedDateCustomerRef
// Reference to the customer the invoice has been issued to.
type PostInvoiceSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

// PostInvoiceSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type PostInvoiceSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostInvoiceSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type PostInvoiceSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PostInvoiceSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type PostInvoiceSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type PostInvoiceSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostInvoiceSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type PostInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	PostInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       PostInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	PostInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable PostInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	PostInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       PostInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type PostInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	PostInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       PostInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	PostInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable PostInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	PostInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       PostInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type PostInvoiceSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostInvoiceSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []PostInvoiceSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *PostInvoiceSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   PostInvoiceSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo PostInvoiceSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *PostInvoiceSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type PostInvoiceSourceModifiedDateLineItems struct {
	AccountRef           *PostInvoiceSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                      `json:"description,omitempty"`
	DiscountAmount       *float64                                                     `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                     `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                                                        `json:"isDirectIncome,omitempty"`
	ItemRef              *PostInvoiceSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                      `json:"quantity"`
	SubTotal             *float64                                                     `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                     `json:"taxAmount,omitempty"`
	TaxRateRef           *PostInvoiceSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                     `json:"totalAmount,omitempty"`
	Tracking             *PostInvoiceSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []PostInvoiceSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                      `json:"unitAmount"`
}

type PostInvoiceSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostInvoiceSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// PostInvoiceSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type PostInvoiceSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PostInvoiceSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *PostInvoiceSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                           `json:"currency,omitempty"`
	CurrencyRate *float64                                                          `json:"currencyRate,omitempty"`
	ID           *string                                                           `json:"id,omitempty"`
	Note         *string                                                           `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                        `json:"paidOnDate,omitempty"`
	Reference    *string                                                           `json:"reference,omitempty"`
	TotalAmount  *float64                                                          `json:"totalAmount,omitempty"`
}

type PostInvoiceSourceModifiedDatePaymentAllocations struct {
	Allocation PostInvoiceSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    PostInvoiceSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type PostInvoiceSourceModifiedDateStatusEnum string

const (
	PostInvoiceSourceModifiedDateStatusEnumUnknown       PostInvoiceSourceModifiedDateStatusEnum = "Unknown"
	PostInvoiceSourceModifiedDateStatusEnumDraft         PostInvoiceSourceModifiedDateStatusEnum = "Draft"
	PostInvoiceSourceModifiedDateStatusEnumSubmitted     PostInvoiceSourceModifiedDateStatusEnum = "Submitted"
	PostInvoiceSourceModifiedDateStatusEnumPartiallyPaid PostInvoiceSourceModifiedDateStatusEnum = "PartiallyPaid"
	PostInvoiceSourceModifiedDateStatusEnumPaid          PostInvoiceSourceModifiedDateStatusEnum = "Paid"
	PostInvoiceSourceModifiedDateStatusEnumVoid          PostInvoiceSourceModifiedDateStatusEnum = "Void"
)

type PostInvoiceSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type PostInvoiceSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// PostInvoiceSourceModifiedDate
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
type PostInvoiceSourceModifiedDate struct {
	AdditionalTaxAmount     *float64                                          `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                                          `json:"additionalTaxPercentage,omitempty"`
	AmountDue               float64                                           `json:"amountDue"`
	Currency                *string                                           `json:"currency,omitempty"`
	CurrencyRate            *float64                                          `json:"currencyRate,omitempty"`
	CustomerRef             *PostInvoiceSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	DiscountPercentage      *float64                                          `json:"discountPercentage,omitempty"`
	DueDate                 *time.Time                                        `json:"dueDate,omitempty"`
	ID                      *string                                           `json:"id,omitempty"`
	InvoiceNumber           *string                                           `json:"invoiceNumber,omitempty"`
	IssueDate               time.Time                                         `json:"issueDate"`
	LineItems               []PostInvoiceSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                *PostInvoiceSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                                        `json:"modifiedDate,omitempty"`
	Note                    *string                                           `json:"note,omitempty"`
	PaidOnDate              *time.Time                                        `json:"paidOnDate,omitempty"`
	PaymentAllocations      []PostInvoiceSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	SalesOrderRefs          []string                                          `json:"salesOrderRefs,omitempty"`
	SourceModifiedDate      *time.Time                                        `json:"sourceModifiedDate,omitempty"`
	Status                  PostInvoiceSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal                *float64                                          `json:"subTotal,omitempty"`
	SupplementalData        *PostInvoiceSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TotalAmount             float64                                           `json:"totalAmount"`
	TotalDiscount           *float64                                          `json:"totalDiscount,omitempty"`
	TotalTaxAmount          float64                                           `json:"totalTaxAmount"`
	WithholdingTax          []PostInvoiceSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type PostInvoiceRequest struct {
	PathParams  PostInvoicePathParams
	QueryParams PostInvoiceQueryParams
	Request     *PostInvoiceSourceModifiedDate `request:"mediaType=application/json"`
}

type PostInvoice200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostInvoice200ApplicationJSONChangesTypeEnum string

const (
	PostInvoice200ApplicationJSONChangesTypeEnumUnknown            PostInvoice200ApplicationJSONChangesTypeEnum = "Unknown"
	PostInvoice200ApplicationJSONChangesTypeEnumCreated            PostInvoice200ApplicationJSONChangesTypeEnum = "Created"
	PostInvoice200ApplicationJSONChangesTypeEnumModified           PostInvoice200ApplicationJSONChangesTypeEnum = "Modified"
	PostInvoice200ApplicationJSONChangesTypeEnumDeleted            PostInvoice200ApplicationJSONChangesTypeEnum = "Deleted"
	PostInvoice200ApplicationJSONChangesTypeEnumAttachmentUploaded PostInvoice200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostInvoice200ApplicationJSONChanges struct {
	AttachmentID *string                                                     `json:"attachmentId,omitempty"`
	RecordRef    *PostInvoice200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostInvoice200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// PostInvoice200ApplicationJSONSourceModifiedDateCustomerRef
// Reference to the customer the invoice has been issued to.
type PostInvoice200ApplicationJSONSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

// PostInvoice200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type PostInvoice200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostInvoice200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type PostInvoice200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type PostInvoice200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *PostInvoice200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                        `json:"description,omitempty"`
	DiscountAmount       *float64                                                                       `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                       `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                                                                          `json:"isDirectIncome,omitempty"`
	ItemRef              *PostInvoice200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                        `json:"quantity"`
	SubTotal             *float64                                                                       `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                       `json:"taxAmount,omitempty"`
	TaxRateRef           *PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                       `json:"totalAmount,omitempty"`
	Tracking             *PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []PostInvoice200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                        `json:"unitAmount"`
}

type PostInvoice200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// PostInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type PostInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PostInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *PostInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                             `json:"currency,omitempty"`
	CurrencyRate *float64                                                                            `json:"currencyRate,omitempty"`
	ID           *string                                                                             `json:"id,omitempty"`
	Note         *string                                                                             `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                          `json:"paidOnDate,omitempty"`
	Reference    *string                                                                             `json:"reference,omitempty"`
	TotalAmount  *float64                                                                            `json:"totalAmount,omitempty"`
}

type PostInvoice200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation PostInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    PostInvoice200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type PostInvoice200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	PostInvoice200ApplicationJSONSourceModifiedDateStatusEnumUnknown       PostInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	PostInvoice200ApplicationJSONSourceModifiedDateStatusEnumDraft         PostInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "Draft"
	PostInvoice200ApplicationJSONSourceModifiedDateStatusEnumSubmitted     PostInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "Submitted"
	PostInvoice200ApplicationJSONSourceModifiedDateStatusEnumPartiallyPaid PostInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "PartiallyPaid"
	PostInvoice200ApplicationJSONSourceModifiedDateStatusEnumPaid          PostInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "Paid"
	PostInvoice200ApplicationJSONSourceModifiedDateStatusEnumVoid          PostInvoice200ApplicationJSONSourceModifiedDateStatusEnum = "Void"
)

type PostInvoice200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type PostInvoice200ApplicationJSONSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// PostInvoice200ApplicationJSONSourceModifiedDate
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
type PostInvoice200ApplicationJSONSourceModifiedDate struct {
	AdditionalTaxAmount     *float64                                                            `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                                                            `json:"additionalTaxPercentage,omitempty"`
	AmountDue               float64                                                             `json:"amountDue"`
	Currency                *string                                                             `json:"currency,omitempty"`
	CurrencyRate            *float64                                                            `json:"currencyRate,omitempty"`
	CustomerRef             *PostInvoice200ApplicationJSONSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	DiscountPercentage      *float64                                                            `json:"discountPercentage,omitempty"`
	DueDate                 *time.Time                                                          `json:"dueDate,omitempty"`
	ID                      *string                                                             `json:"id,omitempty"`
	InvoiceNumber           *string                                                             `json:"invoiceNumber,omitempty"`
	IssueDate               time.Time                                                           `json:"issueDate"`
	LineItems               []PostInvoice200ApplicationJSONSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                *PostInvoice200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                                                          `json:"modifiedDate,omitempty"`
	Note                    *string                                                             `json:"note,omitempty"`
	PaidOnDate              *time.Time                                                          `json:"paidOnDate,omitempty"`
	PaymentAllocations      []PostInvoice200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	SalesOrderRefs          []string                                                            `json:"salesOrderRefs,omitempty"`
	SourceModifiedDate      *time.Time                                                          `json:"sourceModifiedDate,omitempty"`
	Status                  PostInvoice200ApplicationJSONSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal                *float64                                                            `json:"subTotal,omitempty"`
	SupplementalData        *PostInvoice200ApplicationJSONSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TotalAmount             float64                                                             `json:"totalAmount"`
	TotalDiscount           *float64                                                            `json:"totalDiscount,omitempty"`
	TotalTaxAmount          float64                                                             `json:"totalTaxAmount"`
	WithholdingTax          []PostInvoice200ApplicationJSONSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type PostInvoice200ApplicationJSONStatusEnum string

const (
	PostInvoice200ApplicationJSONStatusEnumPending  PostInvoice200ApplicationJSONStatusEnum = "Pending"
	PostInvoice200ApplicationJSONStatusEnumFailed   PostInvoice200ApplicationJSONStatusEnum = "Failed"
	PostInvoice200ApplicationJSONStatusEnumSuccess  PostInvoice200ApplicationJSONStatusEnum = "Success"
	PostInvoice200ApplicationJSONStatusEnumTimedOut PostInvoice200ApplicationJSONStatusEnum = "TimedOut"
)

type PostInvoice200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostInvoice200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostInvoice200ApplicationJSONValidation struct {
	Errors   []PostInvoice200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PostInvoice200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PostInvoice200ApplicationJSON struct {
	Changes           []PostInvoice200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                           `json:"companyId"`
	CompletedOnUtc    *time.Time                                       `json:"completedOnUtc,omitempty"`
	Data              *PostInvoice200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                           `json:"dataConnectionKey"`
	DataType          *string                                          `json:"dataType,omitempty"`
	ErrorMessage      *string                                          `json:"errorMessage,omitempty"`
	PushOperationKey  string                                           `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                        `json:"requestedOnUtc"`
	Status            PostInvoice200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                              `json:"statusCode"`
	TimeoutInMinutes  *int                                             `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                             `json:"timeoutInSeconds,omitempty"`
	Validation        *PostInvoice200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PostInvoiceResponse struct {
	ContentType                         string
	StatusCode                          int
	RawResponse                         *http.Response
	PostInvoice200ApplicationJSONObject *PostInvoice200ApplicationJSON
}
