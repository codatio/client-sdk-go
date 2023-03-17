package operations

import (
	"net/http"
	"time"
)

// CreateCreditNoteSourceModifiedDateCustomerRef
// Reference to the customer the credit note has been issued to.
type CreateCreditNoteSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

// CreateCreditNoteSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type CreateCreditNoteSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateCreditNoteSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type CreateCreditNoteSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateCreditNoteSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type CreateCreditNoteSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// CreateCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs
// References a category against which the item is tracked.
type CreateCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateCreditNoteSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type CreateCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	CreateCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       CreateCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	CreateCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable CreateCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	CreateCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       CreateCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type CreateCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	CreateCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       CreateCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	CreateCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable CreateCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	CreateCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       CreateCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type CreateCreditNoteSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateCreditNoteSourceModifiedDateLineItemsTracking
// Categories, and a project and customer, against which the item is tracked.
type CreateCreditNoteSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []CreateCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *CreateCreditNoteSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   CreateCreditNoteSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo CreateCreditNoteSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *CreateCreditNoteSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type CreateCreditNoteSourceModifiedDateLineItems struct {
	AccountRef           *CreateCreditNoteSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                           `json:"description,omitempty"`
	DiscountAmount       *float64                                                          `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                          `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                                                             `json:"isDirectIncome,omitempty"`
	ItemRef              *CreateCreditNoteSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                           `json:"quantity"`
	SubTotal             *float64                                                          `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                          `json:"taxAmount,omitempty"`
	TaxRateRef           *CreateCreditNoteSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                          `json:"totalAmount,omitempty"`
	Tracking             *CreateCreditNoteSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []CreateCreditNoteSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                           `json:"unitAmount"`
}

type CreateCreditNoteSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateCreditNoteSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// CreateCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type CreateCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CreateCreditNoteSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *CreateCreditNoteSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                `json:"currency,omitempty"`
	CurrencyRate *float64                                                               `json:"currencyRate,omitempty"`
	ID           *string                                                                `json:"id,omitempty"`
	Note         *string                                                                `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                             `json:"paidOnDate,omitempty"`
	Reference    *string                                                                `json:"reference,omitempty"`
	TotalAmount  *float64                                                               `json:"totalAmount,omitempty"`
}

type CreateCreditNoteSourceModifiedDatePaymentAllocations struct {
	Allocation CreateCreditNoteSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    CreateCreditNoteSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type CreateCreditNoteSourceModifiedDateStatusEnum string

const (
	CreateCreditNoteSourceModifiedDateStatusEnumUnknown       CreateCreditNoteSourceModifiedDateStatusEnum = "Unknown"
	CreateCreditNoteSourceModifiedDateStatusEnumDraft         CreateCreditNoteSourceModifiedDateStatusEnum = "Draft"
	CreateCreditNoteSourceModifiedDateStatusEnumSubmitted     CreateCreditNoteSourceModifiedDateStatusEnum = "Submitted"
	CreateCreditNoteSourceModifiedDateStatusEnumPaid          CreateCreditNoteSourceModifiedDateStatusEnum = "Paid"
	CreateCreditNoteSourceModifiedDateStatusEnumVoid          CreateCreditNoteSourceModifiedDateStatusEnum = "Void"
	CreateCreditNoteSourceModifiedDateStatusEnumPartiallyPaid CreateCreditNoteSourceModifiedDateStatusEnum = "PartiallyPaid"
)

// CreateCreditNoteSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateCreditNoteSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type CreateCreditNoteSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// CreateCreditNoteSourceModifiedDate
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
type CreateCreditNoteSourceModifiedDate struct {
	AdditionalTaxAmount     *float64                                               `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                                               `json:"additionalTaxPercentage,omitempty"`
	AllocatedOnDate         *time.Time                                             `json:"allocatedOnDate,omitempty"`
	CreditNoteNumber        *string                                                `json:"creditNoteNumber,omitempty"`
	Currency                *string                                                `json:"currency,omitempty"`
	CurrencyRate            *float64                                               `json:"currencyRate,omitempty"`
	CustomerRef             *CreateCreditNoteSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	DiscountPercentage      float64                                                `json:"discountPercentage"`
	ID                      *string                                                `json:"id,omitempty"`
	IssueDate               *time.Time                                             `json:"issueDate,omitempty"`
	LineItems               []CreateCreditNoteSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                *CreateCreditNoteSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                                             `json:"modifiedDate,omitempty"`
	Note                    *string                                                `json:"note,omitempty"`
	PaymentAllocations      []CreateCreditNoteSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	RemainingCredit         float64                                                `json:"remainingCredit"`
	SourceModifiedDate      *time.Time                                             `json:"sourceModifiedDate,omitempty"`
	Status                  CreateCreditNoteSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal                float64                                                `json:"subTotal"`
	SupplementalData        *CreateCreditNoteSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TotalAmount             float64                                                `json:"totalAmount"`
	TotalDiscount           float64                                                `json:"totalDiscount"`
	TotalTaxAmount          float64                                                `json:"totalTaxAmount"`
	WithholdingTax          []CreateCreditNoteSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type CreateCreditNoteRequest struct {
	RequestBody      *CreateCreditNoteSourceModifiedDate `request:"mediaType=application/json"`
	CompanyID        string                              `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string                              `pathParam:"style=simple,explode=false,name=connectionId"`
	CreditNoteID     string                              `pathParam:"style=simple,explode=false,name=creditNoteId"`
	ForceUpdate      *bool                               `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int                                `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreateCreditNote200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateCreditNote200ApplicationJSONChangesTypeEnum string

const (
	CreateCreditNote200ApplicationJSONChangesTypeEnumUnknown            CreateCreditNote200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateCreditNote200ApplicationJSONChangesTypeEnumCreated            CreateCreditNote200ApplicationJSONChangesTypeEnum = "Created"
	CreateCreditNote200ApplicationJSONChangesTypeEnumModified           CreateCreditNote200ApplicationJSONChangesTypeEnum = "Modified"
	CreateCreditNote200ApplicationJSONChangesTypeEnumDeleted            CreateCreditNote200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateCreditNote200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateCreditNote200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type CreateCreditNote200ApplicationJSONChanges struct {
	AttachmentID *string                                                          `json:"attachmentId,omitempty"`
	RecordRef    *CreateCreditNote200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateCreditNote200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// CreateCreditNote200ApplicationJSONSourceModifiedDateCustomerRef
// Reference to the customer the credit note has been issued to.
type CreateCreditNote200ApplicationJSONSourceModifiedDateCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

// CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the item the line is linked to.
type CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs
// References a category against which the item is tracked.
type CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTracking
// Categories, and a project and customer, against which the item is tracked.
type CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type CreateCreditNote200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                             `json:"description,omitempty"`
	DiscountAmount       *float64                                                                            `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                            `json:"discountPercentage,omitempty"`
	IsDirectIncome       *bool                                                                               `json:"isDirectIncome,omitempty"`
	ItemRef              *CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                             `json:"quantity"`
	SubTotal             *float64                                                                            `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                            `json:"taxAmount,omitempty"`
	TaxRateRef           *CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                            `json:"totalAmount,omitempty"`
	Tracking             *CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []CreateCreditNote200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                             `json:"unitAmount"`
}

type CreateCreditNote200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// CreateCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type CreateCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CreateCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *CreateCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                                  `json:"currency,omitempty"`
	CurrencyRate *float64                                                                                 `json:"currencyRate,omitempty"`
	ID           *string                                                                                  `json:"id,omitempty"`
	Note         *string                                                                                  `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                               `json:"paidOnDate,omitempty"`
	Reference    *string                                                                                  `json:"reference,omitempty"`
	TotalAmount  *float64                                                                                 `json:"totalAmount,omitempty"`
}

type CreateCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation CreateCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    CreateCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type CreateCreditNote200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	CreateCreditNote200ApplicationJSONSourceModifiedDateStatusEnumUnknown       CreateCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	CreateCreditNote200ApplicationJSONSourceModifiedDateStatusEnumDraft         CreateCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Draft"
	CreateCreditNote200ApplicationJSONSourceModifiedDateStatusEnumSubmitted     CreateCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Submitted"
	CreateCreditNote200ApplicationJSONSourceModifiedDateStatusEnumPaid          CreateCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Paid"
	CreateCreditNote200ApplicationJSONSourceModifiedDateStatusEnumVoid          CreateCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "Void"
	CreateCreditNote200ApplicationJSONSourceModifiedDateStatusEnumPartiallyPaid CreateCreditNote200ApplicationJSONSourceModifiedDateStatusEnum = "PartiallyPaid"
)

// CreateCreditNote200ApplicationJSONSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateCreditNote200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type CreateCreditNote200ApplicationJSONSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// CreateCreditNote200ApplicationJSONSourceModifiedDate
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
type CreateCreditNote200ApplicationJSONSourceModifiedDate struct {
	AdditionalTaxAmount     *float64                                                                 `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxPercentage *float64                                                                 `json:"additionalTaxPercentage,omitempty"`
	AllocatedOnDate         *time.Time                                                               `json:"allocatedOnDate,omitempty"`
	CreditNoteNumber        *string                                                                  `json:"creditNoteNumber,omitempty"`
	Currency                *string                                                                  `json:"currency,omitempty"`
	CurrencyRate            *float64                                                                 `json:"currencyRate,omitempty"`
	CustomerRef             *CreateCreditNote200ApplicationJSONSourceModifiedDateCustomerRef         `json:"customerRef,omitempty"`
	DiscountPercentage      float64                                                                  `json:"discountPercentage"`
	ID                      *string                                                                  `json:"id,omitempty"`
	IssueDate               *time.Time                                                               `json:"issueDate,omitempty"`
	LineItems               []CreateCreditNote200ApplicationJSONSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata                *CreateCreditNote200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate            *time.Time                                                               `json:"modifiedDate,omitempty"`
	Note                    *string                                                                  `json:"note,omitempty"`
	PaymentAllocations      []CreateCreditNote200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	RemainingCredit         float64                                                                  `json:"remainingCredit"`
	SourceModifiedDate      *time.Time                                                               `json:"sourceModifiedDate,omitempty"`
	Status                  CreateCreditNote200ApplicationJSONSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal                float64                                                                  `json:"subTotal"`
	SupplementalData        *CreateCreditNote200ApplicationJSONSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TotalAmount             float64                                                                  `json:"totalAmount"`
	TotalDiscount           float64                                                                  `json:"totalDiscount"`
	TotalTaxAmount          float64                                                                  `json:"totalTaxAmount"`
	WithholdingTax          []CreateCreditNote200ApplicationJSONSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type CreateCreditNote200ApplicationJSONStatusEnum string

const (
	CreateCreditNote200ApplicationJSONStatusEnumPending  CreateCreditNote200ApplicationJSONStatusEnum = "Pending"
	CreateCreditNote200ApplicationJSONStatusEnumFailed   CreateCreditNote200ApplicationJSONStatusEnum = "Failed"
	CreateCreditNote200ApplicationJSONStatusEnumSuccess  CreateCreditNote200ApplicationJSONStatusEnum = "Success"
	CreateCreditNote200ApplicationJSONStatusEnumTimedOut CreateCreditNote200ApplicationJSONStatusEnum = "TimedOut"
)

type CreateCreditNote200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateCreditNote200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateCreditNote200ApplicationJSONValidation struct {
	Errors   []CreateCreditNote200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateCreditNote200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type CreateCreditNote200ApplicationJSON struct {
	Changes           []CreateCreditNote200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                                `json:"companyId"`
	CompletedOnUtc    *time.Time                                            `json:"completedOnUtc,omitempty"`
	Data              *CreateCreditNote200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                                `json:"dataConnectionKey"`
	DataType          *string                                               `json:"dataType,omitempty"`
	ErrorMessage      *string                                               `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                             `json:"requestedOnUtc"`
	Status            CreateCreditNote200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                   `json:"statusCode"`
	TimeoutInMinutes  *int                                                  `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                  `json:"timeoutInSeconds,omitempty"`
	Validation        *CreateCreditNote200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type CreateCreditNoteResponse struct {
	ContentType                              string
	StatusCode                               int
	RawResponse                              *http.Response
	CreateCreditNote200ApplicationJSONObject *CreateCreditNote200ApplicationJSON
}
