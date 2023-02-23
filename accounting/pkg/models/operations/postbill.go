package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostBillPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostBillQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// PostBillSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type PostBillSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostBillSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the line item is linked.
type PostBillSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PostBillSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type PostBillSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type PostBillSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostBillSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type PostBillSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	PostBillSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       PostBillSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	PostBillSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable PostBillSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	PostBillSourceModifiedDateLineItemsTrackingIsBilledToEnumCustomer      PostBillSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Customer"
	PostBillSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       PostBillSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type PostBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	PostBillSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       PostBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	PostBillSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable PostBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	PostBillSourceModifiedDateLineItemsTrackingIsRebilledToEnumCustomer      PostBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Customer"
	PostBillSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       PostBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type PostBillSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostBillSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []PostBillSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *PostBillSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   PostBillSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo PostBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *PostBillSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type PostBillSourceModifiedDateLineItems struct {
	AccountRef           *PostBillSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                   `json:"description,omitempty"`
	DiscountAmount       *float64                                                  `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                  `json:"discountPercentage,omitempty"`
	IsDirectCost         *bool                                                     `json:"isDirectCost,omitempty"`
	ItemRef              *PostBillSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                   `json:"quantity"`
	SubTotal             *float64                                                  `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                  `json:"taxAmount,omitempty"`
	TaxRateRef           *PostBillSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                  `json:"totalAmount,omitempty"`
	Tracking             *PostBillSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []PostBillSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                   `json:"unitAmount"`
}

type PostBillSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostBillSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// PostBillSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type PostBillSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PostBillSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *PostBillSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                        `json:"currency,omitempty"`
	CurrencyRate *float64                                                       `json:"currencyRate,omitempty"`
	ID           *string                                                        `json:"id,omitempty"`
	Note         *string                                                        `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                     `json:"paidOnDate,omitempty"`
	Reference    *string                                                        `json:"reference,omitempty"`
	TotalAmount  *float64                                                       `json:"totalAmount,omitempty"`
}

type PostBillSourceModifiedDatePaymentAllocations struct {
	Allocation PostBillSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    PostBillSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type PostBillSourceModifiedDatePurchaseOrderRefs struct {
	ID                  *string `json:"id,omitempty"`
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
}

type PostBillSourceModifiedDateStatusEnum string

const (
	PostBillSourceModifiedDateStatusEnumUnknown       PostBillSourceModifiedDateStatusEnum = "Unknown"
	PostBillSourceModifiedDateStatusEnumOpen          PostBillSourceModifiedDateStatusEnum = "Open"
	PostBillSourceModifiedDateStatusEnumPartiallyPaid PostBillSourceModifiedDateStatusEnum = "PartiallyPaid"
	PostBillSourceModifiedDateStatusEnumPaid          PostBillSourceModifiedDateStatusEnum = "Paid"
	PostBillSourceModifiedDateStatusEnumVoid          PostBillSourceModifiedDateStatusEnum = "Void"
	PostBillSourceModifiedDateStatusEnumDraft         PostBillSourceModifiedDateStatusEnum = "Draft"
)

type PostBillSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// PostBillSourceModifiedDateSupplierRef
// Reference to the supplier the bill was received from.
type PostBillSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type PostBillSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// PostBillSourceModifiedDate
// > **Invoices or bills?**
// >
// > In Codat, bills are for accounts payable only. For the accounts receivable equivalent of bills, see [Invoices](https://docs.codat.io/accounting-api#/schemas/Invoice).
//
// View the coverage for bills in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// In Codat, a bill contains details of:
// * When the bill was recorded in the accounting system.
// * How much the bill is for and the currency of the amount.
// * Who the bill was received from — the *supplier*.
// * What the bill is for — the *line items*.
//
// Some accounting platforms give a separate name to purchases where the payment is made immediately, such as something bought with a credit card or online payment. One example of this would be QuickBooks Online's *expenses*.
//
// You can find these types of transactions in our [Direct costs](https://docs.codat.io/accounting-api#/schemas/DirectCost) data model.
type PostBillSourceModifiedDate struct {
	AmountDue          *float64                                       `json:"amountDue,omitempty"`
	Currency           *string                                        `json:"currency,omitempty"`
	CurrencyRate       *float64                                       `json:"currencyRate,omitempty"`
	DueDate            *time.Time                                     `json:"dueDate,omitempty"`
	ID                 *string                                        `json:"id,omitempty"`
	IssueDate          time.Time                                      `json:"issueDate"`
	LineItems          []PostBillSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata           *PostBillSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                     `json:"modifiedDate,omitempty"`
	Note               *string                                        `json:"note,omitempty"`
	PaymentAllocations []PostBillSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	PurchaseOrderRefs  []PostBillSourceModifiedDatePurchaseOrderRefs  `json:"purchaseOrderRefs,omitempty"`
	Reference          *string                                        `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                     `json:"sourceModifiedDate,omitempty"`
	Status             PostBillSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal           float64                                        `json:"subTotal"`
	SupplementalData   *PostBillSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef        *PostBillSourceModifiedDateSupplierRef         `json:"supplierRef,omitempty"`
	TaxAmount          float64                                        `json:"taxAmount"`
	TotalAmount        float64                                        `json:"totalAmount"`
	WithholdingTax     []PostBillSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type PostBillSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostBillRequest struct {
	PathParams  PostBillPathParams
	QueryParams PostBillQueryParams
	Request     *PostBillSourceModifiedDate `request:"mediaType=application/json"`
	Security    PostBillSecurity
}

type PostBill200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostBill200ApplicationJSONChangesTypeEnum string

const (
	PostBill200ApplicationJSONChangesTypeEnumUnknown            PostBill200ApplicationJSONChangesTypeEnum = "Unknown"
	PostBill200ApplicationJSONChangesTypeEnumCreated            PostBill200ApplicationJSONChangesTypeEnum = "Created"
	PostBill200ApplicationJSONChangesTypeEnumModified           PostBill200ApplicationJSONChangesTypeEnum = "Modified"
	PostBill200ApplicationJSONChangesTypeEnumDeleted            PostBill200ApplicationJSONChangesTypeEnum = "Deleted"
	PostBill200ApplicationJSONChangesTypeEnumAttachmentUploaded PostBill200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostBill200ApplicationJSONChanges struct {
	AttachmentID *string                                                  `json:"attachmentId,omitempty"`
	RecordRef    *PostBill200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostBill200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// PostBill200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type PostBill200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostBill200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the line item is linked.
type PostBill200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PostBill200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type PostBill200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumCustomer      PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Customer"
	PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumCustomer      PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Customer"
	PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostBill200ApplicationJSONSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type PostBill200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *PostBill200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                     `json:"description,omitempty"`
	DiscountAmount       *float64                                                                    `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                    `json:"discountPercentage,omitempty"`
	IsDirectCost         *bool                                                                       `json:"isDirectCost,omitempty"`
	ItemRef              *PostBill200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                     `json:"quantity"`
	SubTotal             *float64                                                                    `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                    `json:"taxAmount,omitempty"`
	TaxRateRef           *PostBill200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                    `json:"totalAmount,omitempty"`
	Tracking             *PostBill200ApplicationJSONSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []PostBill200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                     `json:"unitAmount"`
}

type PostBill200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostBill200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// PostBill200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type PostBill200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PostBill200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *PostBill200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                          `json:"currency,omitempty"`
	CurrencyRate *float64                                                                         `json:"currencyRate,omitempty"`
	ID           *string                                                                          `json:"id,omitempty"`
	Note         *string                                                                          `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                       `json:"paidOnDate,omitempty"`
	Reference    *string                                                                          `json:"reference,omitempty"`
	TotalAmount  *float64                                                                         `json:"totalAmount,omitempty"`
}

type PostBill200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation PostBill200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    PostBill200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type PostBill200ApplicationJSONSourceModifiedDatePurchaseOrderRefs struct {
	ID                  *string `json:"id,omitempty"`
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
}

type PostBill200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	PostBill200ApplicationJSONSourceModifiedDateStatusEnumUnknown       PostBill200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	PostBill200ApplicationJSONSourceModifiedDateStatusEnumOpen          PostBill200ApplicationJSONSourceModifiedDateStatusEnum = "Open"
	PostBill200ApplicationJSONSourceModifiedDateStatusEnumPartiallyPaid PostBill200ApplicationJSONSourceModifiedDateStatusEnum = "PartiallyPaid"
	PostBill200ApplicationJSONSourceModifiedDateStatusEnumPaid          PostBill200ApplicationJSONSourceModifiedDateStatusEnum = "Paid"
	PostBill200ApplicationJSONSourceModifiedDateStatusEnumVoid          PostBill200ApplicationJSONSourceModifiedDateStatusEnum = "Void"
	PostBill200ApplicationJSONSourceModifiedDateStatusEnumDraft         PostBill200ApplicationJSONSourceModifiedDateStatusEnum = "Draft"
)

type PostBill200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// PostBill200ApplicationJSONSourceModifiedDateSupplierRef
// Reference to the supplier the bill was received from.
type PostBill200ApplicationJSONSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type PostBill200ApplicationJSONSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// PostBill200ApplicationJSONSourceModifiedDate
// > **Invoices or bills?**
// >
// > In Codat, bills are for accounts payable only. For the accounts receivable equivalent of bills, see [Invoices](https://docs.codat.io/accounting-api#/schemas/Invoice).
//
// View the coverage for bills in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// In Codat, a bill contains details of:
// * When the bill was recorded in the accounting system.
// * How much the bill is for and the currency of the amount.
// * Who the bill was received from — the *supplier*.
// * What the bill is for — the *line items*.
//
// Some accounting platforms give a separate name to purchases where the payment is made immediately, such as something bought with a credit card or online payment. One example of this would be QuickBooks Online's *expenses*.
//
// You can find these types of transactions in our [Direct costs](https://docs.codat.io/accounting-api#/schemas/DirectCost) data model.
type PostBill200ApplicationJSONSourceModifiedDate struct {
	AmountDue          *float64                                                         `json:"amountDue,omitempty"`
	Currency           *string                                                          `json:"currency,omitempty"`
	CurrencyRate       *float64                                                         `json:"currencyRate,omitempty"`
	DueDate            *time.Time                                                       `json:"dueDate,omitempty"`
	ID                 *string                                                          `json:"id,omitempty"`
	IssueDate          time.Time                                                        `json:"issueDate"`
	LineItems          []PostBill200ApplicationJSONSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata           *PostBill200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                       `json:"modifiedDate,omitempty"`
	Note               *string                                                          `json:"note,omitempty"`
	PaymentAllocations []PostBill200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	PurchaseOrderRefs  []PostBill200ApplicationJSONSourceModifiedDatePurchaseOrderRefs  `json:"purchaseOrderRefs,omitempty"`
	Reference          *string                                                          `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                                       `json:"sourceModifiedDate,omitempty"`
	Status             PostBill200ApplicationJSONSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal           float64                                                          `json:"subTotal"`
	SupplementalData   *PostBill200ApplicationJSONSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef        *PostBill200ApplicationJSONSourceModifiedDateSupplierRef         `json:"supplierRef,omitempty"`
	TaxAmount          float64                                                          `json:"taxAmount"`
	TotalAmount        float64                                                          `json:"totalAmount"`
	WithholdingTax     []PostBill200ApplicationJSONSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type PostBill200ApplicationJSONStatusEnum string

const (
	PostBill200ApplicationJSONStatusEnumPending  PostBill200ApplicationJSONStatusEnum = "Pending"
	PostBill200ApplicationJSONStatusEnumFailed   PostBill200ApplicationJSONStatusEnum = "Failed"
	PostBill200ApplicationJSONStatusEnumSuccess  PostBill200ApplicationJSONStatusEnum = "Success"
	PostBill200ApplicationJSONStatusEnumTimedOut PostBill200ApplicationJSONStatusEnum = "TimedOut"
)

type PostBill200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostBill200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostBill200ApplicationJSONValidation struct {
	Errors   []PostBill200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PostBill200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PostBill200ApplicationJSON struct {
	Changes           []PostBill200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                        `json:"companyId"`
	CompletedOnUtc    *time.Time                                    `json:"completedOnUtc,omitempty"`
	Data              *PostBill200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                        `json:"dataConnectionKey"`
	DataType          *string                                       `json:"dataType,omitempty"`
	ErrorMessage      *string                                       `json:"errorMessage,omitempty"`
	PushOperationKey  string                                        `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                     `json:"requestedOnUtc"`
	Status            PostBill200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                           `json:"statusCode"`
	TimeoutInMinutes  *int                                          `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                          `json:"timeoutInSeconds,omitempty"`
	Validation        *PostBill200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PostBillResponse struct {
	ContentType                      string
	StatusCode                       int
	PostBill200ApplicationJSONObject *PostBill200ApplicationJSON
}
