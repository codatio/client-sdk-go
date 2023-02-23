package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostCreditNotePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	CreditNoteID string `pathParam:"style=simple,explode=false,name=creditNoteId"`
}

type PostCreditNoteQueryParams struct {
	ForceUpdate      *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// PostCreditNoteSourceModifiedDateCustomerRef
// Reference to the customer the credit note has been issued to.
type PostCreditNoteSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

// PostCreditNoteSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type PostCreditNoteSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostCreditNoteSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type PostCreditNoteSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PostCreditNoteSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type PostCreditNoteSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type PostCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostCreditNoteSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type PostCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	PostCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       PostCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	PostCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable PostCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	PostCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       PostCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type PostCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	PostCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       PostCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	PostCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable PostCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	PostCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       PostCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type PostCreditNoteSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostCreditNoteSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []PostCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *PostCreditNoteSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   PostCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo PostCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *PostCreditNoteSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type PostCreditNoteSourceModifiedDateLineItems struct {
	AccountRef           *PostCreditNoteSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                         `json:"description,omitempty"`
	DiscountAmount       *float64                                                        `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                        `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                                                           `json:"isDirectIncome,omitempty"`
	ItemRef              *PostCreditNoteSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                         `json:"quantity"`
	SubTotal             *float64                                                        `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                        `json:"taxAmount,omitempty"`
	TaxRateRef           *PostCreditNoteSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                        `json:"totalAmount,omitempty"`
	Tracking             *PostCreditNoteSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []PostCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                         `json:"unitAmount"`
}

type PostCreditNoteSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostCreditNoteSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// PostCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type PostCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PostCreditNoteSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *PostCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                              `json:"currency,omitempty"`
	CurrencyRate *float64                                                             `json:"currencyRate,omitempty"`
	ID           *string                                                              `json:"id,omitempty"`
	Note         *string                                                              `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                           `json:"paidOnDate,omitempty"`
	Reference    *string                                                              `json:"reference,omitempty"`
	TotalAmount  *float64                                                             `json:"totalAmount,omitempty"`
}

type PostCreditNoteSourceModifiedDatePaymentAllocations struct {
	Allocation PostCreditNoteSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    PostCreditNoteSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type PostCreditNoteSourceModifiedDateStatusEnum string

const (
	PostCreditNoteSourceModifiedDateStatusEnumUnknown       PostCreditNoteSourceModifiedDateStatusEnum = "Unknown"
	PostCreditNoteSourceModifiedDateStatusEnumDraft         PostCreditNoteSourceModifiedDateStatusEnum = "Draft"
	PostCreditNoteSourceModifiedDateStatusEnumSubmitted     PostCreditNoteSourceModifiedDateStatusEnum = "Submitted"
	PostCreditNoteSourceModifiedDateStatusEnumPaid          PostCreditNoteSourceModifiedDateStatusEnum = "Paid"
	PostCreditNoteSourceModifiedDateStatusEnumVoid          PostCreditNoteSourceModifiedDateStatusEnum = "Void"
	PostCreditNoteSourceModifiedDateStatusEnumPartiallyPaid PostCreditNoteSourceModifiedDateStatusEnum = "PartiallyPaid"
)

type PostCreditNoteSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type PostCreditNoteSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// PostCreditNoteSourceModifiedDate
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
type PostCreditNoteSourceModifiedDate struct {
	AdditionalTaxAmount     *float64                                             `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                                             `json:"additionalTaxPercentage,omitempty"`
	AllocatedOnDate         *time.Time                                           `json:"allocatedOnDate,omitempty"`
	CreditNoteNumber        *string                                              `json:"creditNoteNumber,omitempty"`
	Currency                *string                                              `json:"currency,omitempty"`
	CurrencyRate            *float64                                             `json:"currencyRate,omitempty"`
	CustomerRef             *PostCreditNoteSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	DiscountPercentage      float64                                              `json:"discountPercentage"`
	ID                      *string                                              `json:"id,omitempty"`
	IssueDate               *time.Time                                           `json:"issueDate,omitempty"`
	LineItems               []PostCreditNoteSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                *PostCreditNoteSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                                           `json:"modifiedDate,omitempty"`
	Note                    *string                                              `json:"note,omitempty"`
	PaymentAllocations      []PostCreditNoteSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	RemainingCredit         float64                                              `json:"remainingCredit"`
	SourceModifiedDate      *time.Time                                           `json:"sourceModifiedDate,omitempty"`
	Status                  PostCreditNoteSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal                float64                                              `json:"subTotal"`
	SupplementalData        *PostCreditNoteSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TotalAmount             float64                                              `json:"totalAmount"`
	TotalDiscount           float64                                              `json:"totalDiscount"`
	TotalTaxAmount          float64                                              `json:"totalTaxAmount"`
	WithholdingTax          []PostCreditNoteSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type PostCreditNoteSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostCreditNoteRequest struct {
	PathParams  PostCreditNotePathParams
	QueryParams PostCreditNoteQueryParams
	Request     *PostCreditNoteSourceModifiedDate `request:"mediaType=application/json"`
	Security    PostCreditNoteSecurity
}

type PostCreditNote200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostCreditNote200ApplicationJSONChangesTypeEnum string

const (
	PostCreditNote200ApplicationJSONChangesTypeEnumUnknown            PostCreditNote200ApplicationJSONChangesTypeEnum = "Unknown"
	PostCreditNote200ApplicationJSONChangesTypeEnumCreated            PostCreditNote200ApplicationJSONChangesTypeEnum = "Created"
	PostCreditNote200ApplicationJSONChangesTypeEnumModified           PostCreditNote200ApplicationJSONChangesTypeEnum = "Modified"
	PostCreditNote200ApplicationJSONChangesTypeEnumDeleted            PostCreditNote200ApplicationJSONChangesTypeEnum = "Deleted"
	PostCreditNote200ApplicationJSONChangesTypeEnumAttachmentUploaded PostCreditNote200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostCreditNote200ApplicationJSONChanges struct {
	AttachmentID *string                                                        `json:"attachmentId,omitempty"`
	RecordRef    *PostCreditNote200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostCreditNote200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// PostCreditNote200ApplicationJSONSourceModifiedDateCustomerRef
// Reference to the customer the credit note has been issued to.
type PostCreditNote200ApplicationJSONSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

// PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type PostCreditNote200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                           `json:"description,omitempty"`
	DiscountAmount       *float64                                                                          `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                          `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                                                                             `json:"isDirectIncome,omitempty"`
	ItemRef              *PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                           `json:"quantity"`
	SubTotal             *float64                                                                          `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                          `json:"taxAmount,omitempty"`
	TaxRateRef           *PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                          `json:"totalAmount,omitempty"`
	Tracking             *PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []PostCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                           `json:"unitAmount"`
}

type PostCreditNote200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// PostCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type PostCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PostCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *PostCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                                `json:"currency,omitempty"`
	CurrencyRate *float64                                                                               `json:"currencyRate,omitempty"`
	ID           *string                                                                                `json:"id,omitempty"`
	Note         *string                                                                                `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                             `json:"paidOnDate,omitempty"`
	Reference    *string                                                                                `json:"reference,omitempty"`
	TotalAmount  *float64                                                                               `json:"totalAmount,omitempty"`
}

type PostCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation PostCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    PostCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type PostCreditNote200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	PostCreditNote200ApplicationJSONSourceModifiedDateStatusEnumUnknown       PostCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	PostCreditNote200ApplicationJSONSourceModifiedDateStatusEnumDraft         PostCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Draft"
	PostCreditNote200ApplicationJSONSourceModifiedDateStatusEnumSubmitted     PostCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Submitted"
	PostCreditNote200ApplicationJSONSourceModifiedDateStatusEnumPaid          PostCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Paid"
	PostCreditNote200ApplicationJSONSourceModifiedDateStatusEnumVoid          PostCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Void"
	PostCreditNote200ApplicationJSONSourceModifiedDateStatusEnumPartiallyPaid PostCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "PartiallyPaid"
)

type PostCreditNote200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type PostCreditNote200ApplicationJSONSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// PostCreditNote200ApplicationJSONSourceModifiedDate
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
type PostCreditNote200ApplicationJSONSourceModifiedDate struct {
	AdditionalTaxAmount     *float64                                                               `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                                                               `json:"additionalTaxPercentage,omitempty"`
	AllocatedOnDate         *time.Time                                                             `json:"allocatedOnDate,omitempty"`
	CreditNoteNumber        *string                                                                `json:"creditNoteNumber,omitempty"`
	Currency                *string                                                                `json:"currency,omitempty"`
	CurrencyRate            *float64                                                               `json:"currencyRate,omitempty"`
	CustomerRef             *PostCreditNote200ApplicationJSONSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	DiscountPercentage      float64                                                                `json:"discountPercentage"`
	ID                      *string                                                                `json:"id,omitempty"`
	IssueDate               *time.Time                                                             `json:"issueDate,omitempty"`
	LineItems               []PostCreditNote200ApplicationJSONSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                *PostCreditNote200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                                                             `json:"modifiedDate,omitempty"`
	Note                    *string                                                                `json:"note,omitempty"`
	PaymentAllocations      []PostCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	RemainingCredit         float64                                                                `json:"remainingCredit"`
	SourceModifiedDate      *time.Time                                                             `json:"sourceModifiedDate,omitempty"`
	Status                  PostCreditNote200ApplicationJSONSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal                float64                                                                `json:"subTotal"`
	SupplementalData        *PostCreditNote200ApplicationJSONSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TotalAmount             float64                                                                `json:"totalAmount"`
	TotalDiscount           float64                                                                `json:"totalDiscount"`
	TotalTaxAmount          float64                                                                `json:"totalTaxAmount"`
	WithholdingTax          []PostCreditNote200ApplicationJSONSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type PostCreditNote200ApplicationJSONStatusEnum string

const (
	PostCreditNote200ApplicationJSONStatusEnumPending  PostCreditNote200ApplicationJSONStatusEnum = "Pending"
	PostCreditNote200ApplicationJSONStatusEnumFailed   PostCreditNote200ApplicationJSONStatusEnum = "Failed"
	PostCreditNote200ApplicationJSONStatusEnumSuccess  PostCreditNote200ApplicationJSONStatusEnum = "Success"
	PostCreditNote200ApplicationJSONStatusEnumTimedOut PostCreditNote200ApplicationJSONStatusEnum = "TimedOut"
)

type PostCreditNote200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostCreditNote200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostCreditNote200ApplicationJSONValidation struct {
	Errors   []PostCreditNote200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PostCreditNote200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PostCreditNote200ApplicationJSON struct {
	Changes           []PostCreditNote200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                              `json:"companyId"`
	CompletedOnUtc    *time.Time                                          `json:"completedOnUtc,omitempty"`
	Data              *PostCreditNote200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                              `json:"dataConnectionKey"`
	DataType          *string                                             `json:"dataType,omitempty"`
	ErrorMessage      *string                                             `json:"errorMessage,omitempty"`
	PushOperationKey  string                                              `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                           `json:"requestedOnUtc"`
	Status            PostCreditNote200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                 `json:"statusCode"`
	TimeoutInMinutes  *int                                                `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                `json:"timeoutInSeconds,omitempty"`
	Validation        *PostCreditNote200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PostCreditNoteResponse struct {
	ContentType                            string
	StatusCode                             int
	PostCreditNote200ApplicationJSONObject *PostCreditNote200ApplicationJSON
}
