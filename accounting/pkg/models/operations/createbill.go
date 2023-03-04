package operations

import (
	"net/http"
	"time"
)

type CreateBillPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type CreateBillQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// CreateBillSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type CreateBillSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateBillSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the line item is linked.
type CreateBillSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateBillSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type CreateBillSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type CreateBillSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateBillSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type CreateBillSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	CreateBillSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       CreateBillSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	CreateBillSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable CreateBillSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	CreateBillSourceModifiedDateLineItemsTrackingIsBilledToEnumCustomer      CreateBillSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Customer"
	CreateBillSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       CreateBillSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type CreateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	CreateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       CreateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	CreateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable CreateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	CreateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnumCustomer      CreateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Customer"
	CreateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       CreateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type CreateBillSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateBillSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []CreateBillSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *CreateBillSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   CreateBillSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo CreateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *CreateBillSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type CreateBillSourceModifiedDateLineItems struct {
	AccountRef           *CreateBillSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                     `json:"description,omitempty"`
	DiscountAmount       *float64                                                    `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                    `json:"discountPercentage,omitempty"`
	IsDirectCost         *bool                                                       `json:"isDirectCost,omitempty"`
	ItemRef              *CreateBillSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                     `json:"quantity"`
	SubTotal             *float64                                                    `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                    `json:"taxAmount,omitempty"`
	TaxRateRef           *CreateBillSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                    `json:"totalAmount,omitempty"`
	Tracking             *CreateBillSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []CreateBillSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                     `json:"unitAmount"`
}

type CreateBillSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateBillSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// CreateBillSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type CreateBillSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CreateBillSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *CreateBillSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                          `json:"currency,omitempty"`
	CurrencyRate *float64                                                         `json:"currencyRate,omitempty"`
	ID           *string                                                          `json:"id,omitempty"`
	Note         *string                                                          `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                       `json:"paidOnDate,omitempty"`
	Reference    *string                                                          `json:"reference,omitempty"`
	TotalAmount  *float64                                                         `json:"totalAmount,omitempty"`
}

type CreateBillSourceModifiedDatePaymentAllocations struct {
	Allocation CreateBillSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    CreateBillSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type CreateBillSourceModifiedDatePurchaseOrderRefs struct {
	ID                  *string `json:"id,omitempty"`
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
}

type CreateBillSourceModifiedDateStatusEnum string

const (
	CreateBillSourceModifiedDateStatusEnumUnknown       CreateBillSourceModifiedDateStatusEnum = "Unknown"
	CreateBillSourceModifiedDateStatusEnumOpen          CreateBillSourceModifiedDateStatusEnum = "Open"
	CreateBillSourceModifiedDateStatusEnumPartiallyPaid CreateBillSourceModifiedDateStatusEnum = "PartiallyPaid"
	CreateBillSourceModifiedDateStatusEnumPaid          CreateBillSourceModifiedDateStatusEnum = "Paid"
	CreateBillSourceModifiedDateStatusEnumVoid          CreateBillSourceModifiedDateStatusEnum = "Void"
	CreateBillSourceModifiedDateStatusEnumDraft         CreateBillSourceModifiedDateStatusEnum = "Draft"
)

type CreateBillSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateBillSourceModifiedDateSupplierRef
// Reference to the supplier the bill was received from.
type CreateBillSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type CreateBillSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// CreateBillSourceModifiedDate
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
type CreateBillSourceModifiedDate struct {
	AmountDue          *float64                                         `json:"amountDue,omitempty"`
	Currency           *string                                          `json:"currency,omitempty"`
	CurrencyRate       *float64                                         `json:"currencyRate,omitempty"`
	DueDate            *time.Time                                       `json:"dueDate,omitempty"`
	ID                 *string                                          `json:"id,omitempty"`
	IssueDate          time.Time                                        `json:"issueDate"`
	LineItems          []CreateBillSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata           *CreateBillSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                       `json:"modifiedDate,omitempty"`
	Note               *string                                          `json:"note,omitempty"`
	PaymentAllocations []CreateBillSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	PurchaseOrderRefs  []CreateBillSourceModifiedDatePurchaseOrderRefs  `json:"purchaseOrderRefs,omitempty"`
	Reference          *string                                          `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                       `json:"sourceModifiedDate,omitempty"`
	Status             CreateBillSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal           float64                                          `json:"subTotal"`
	SupplementalData   *CreateBillSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef        *CreateBillSourceModifiedDateSupplierRef         `json:"supplierRef,omitempty"`
	TaxAmount          float64                                          `json:"taxAmount"`
	TotalAmount        float64                                          `json:"totalAmount"`
	WithholdingTax     []CreateBillSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type CreateBillRequest struct {
	PathParams  CreateBillPathParams
	QueryParams CreateBillQueryParams
	Request     *CreateBillSourceModifiedDate `request:"mediaType=application/json"`
}

type CreateBill200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateBill200ApplicationJSONChangesTypeEnum string

const (
	CreateBill200ApplicationJSONChangesTypeEnumUnknown            CreateBill200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateBill200ApplicationJSONChangesTypeEnumCreated            CreateBill200ApplicationJSONChangesTypeEnum = "Created"
	CreateBill200ApplicationJSONChangesTypeEnumModified           CreateBill200ApplicationJSONChangesTypeEnum = "Modified"
	CreateBill200ApplicationJSONChangesTypeEnumDeleted            CreateBill200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateBill200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateBill200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type CreateBill200ApplicationJSONChanges struct {
	AttachmentID *string                                                    `json:"attachmentId,omitempty"`
	RecordRef    *CreateBill200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateBill200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// CreateBill200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type CreateBill200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateBill200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the line item is linked.
type CreateBill200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateBill200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type CreateBill200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumCustomer      CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Customer"
	CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumCustomer      CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Customer"
	CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateBill200ApplicationJSONSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type CreateBill200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *CreateBill200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                       `json:"description,omitempty"`
	DiscountAmount       *float64                                                                      `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                      `json:"discountPercentage,omitempty"`
	IsDirectCost         *bool                                                                         `json:"isDirectCost,omitempty"`
	ItemRef              *CreateBill200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                       `json:"quantity"`
	SubTotal             *float64                                                                      `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                      `json:"taxAmount,omitempty"`
	TaxRateRef           *CreateBill200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                      `json:"totalAmount,omitempty"`
	Tracking             *CreateBill200ApplicationJSONSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []CreateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                       `json:"unitAmount"`
}

type CreateBill200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateBill200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// CreateBill200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type CreateBill200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CreateBill200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *CreateBill200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                            `json:"currency,omitempty"`
	CurrencyRate *float64                                                                           `json:"currencyRate,omitempty"`
	ID           *string                                                                            `json:"id,omitempty"`
	Note         *string                                                                            `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                         `json:"paidOnDate,omitempty"`
	Reference    *string                                                                            `json:"reference,omitempty"`
	TotalAmount  *float64                                                                           `json:"totalAmount,omitempty"`
}

type CreateBill200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation CreateBill200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    CreateBill200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type CreateBill200ApplicationJSONSourceModifiedDatePurchaseOrderRefs struct {
	ID                  *string `json:"id,omitempty"`
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
}

type CreateBill200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	CreateBill200ApplicationJSONSourceModifiedDateStatusEnumUnknown       CreateBill200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	CreateBill200ApplicationJSONSourceModifiedDateStatusEnumOpen          CreateBill200ApplicationJSONSourceModifiedDateStatusEnum = "Open"
	CreateBill200ApplicationJSONSourceModifiedDateStatusEnumPartiallyPaid CreateBill200ApplicationJSONSourceModifiedDateStatusEnum = "PartiallyPaid"
	CreateBill200ApplicationJSONSourceModifiedDateStatusEnumPaid          CreateBill200ApplicationJSONSourceModifiedDateStatusEnum = "Paid"
	CreateBill200ApplicationJSONSourceModifiedDateStatusEnumVoid          CreateBill200ApplicationJSONSourceModifiedDateStatusEnum = "Void"
	CreateBill200ApplicationJSONSourceModifiedDateStatusEnumDraft         CreateBill200ApplicationJSONSourceModifiedDateStatusEnum = "Draft"
)

type CreateBill200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateBill200ApplicationJSONSourceModifiedDateSupplierRef
// Reference to the supplier the bill was received from.
type CreateBill200ApplicationJSONSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type CreateBill200ApplicationJSONSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// CreateBill200ApplicationJSONSourceModifiedDate
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
type CreateBill200ApplicationJSONSourceModifiedDate struct {
	AmountDue          *float64                                                           `json:"amountDue,omitempty"`
	Currency           *string                                                            `json:"currency,omitempty"`
	CurrencyRate       *float64                                                           `json:"currencyRate,omitempty"`
	DueDate            *time.Time                                                         `json:"dueDate,omitempty"`
	ID                 *string                                                            `json:"id,omitempty"`
	IssueDate          time.Time                                                          `json:"issueDate"`
	LineItems          []CreateBill200ApplicationJSONSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata           *CreateBill200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                         `json:"modifiedDate,omitempty"`
	Note               *string                                                            `json:"note,omitempty"`
	PaymentAllocations []CreateBill200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	PurchaseOrderRefs  []CreateBill200ApplicationJSONSourceModifiedDatePurchaseOrderRefs  `json:"purchaseOrderRefs,omitempty"`
	Reference          *string                                                            `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                                         `json:"sourceModifiedDate,omitempty"`
	Status             CreateBill200ApplicationJSONSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal           float64                                                            `json:"subTotal"`
	SupplementalData   *CreateBill200ApplicationJSONSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef        *CreateBill200ApplicationJSONSourceModifiedDateSupplierRef         `json:"supplierRef,omitempty"`
	TaxAmount          float64                                                            `json:"taxAmount"`
	TotalAmount        float64                                                            `json:"totalAmount"`
	WithholdingTax     []CreateBill200ApplicationJSONSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type CreateBill200ApplicationJSONStatusEnum string

const (
	CreateBill200ApplicationJSONStatusEnumPending  CreateBill200ApplicationJSONStatusEnum = "Pending"
	CreateBill200ApplicationJSONStatusEnumFailed   CreateBill200ApplicationJSONStatusEnum = "Failed"
	CreateBill200ApplicationJSONStatusEnumSuccess  CreateBill200ApplicationJSONStatusEnum = "Success"
	CreateBill200ApplicationJSONStatusEnumTimedOut CreateBill200ApplicationJSONStatusEnum = "TimedOut"
)

type CreateBill200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateBill200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateBill200ApplicationJSONValidation struct {
	Errors   []CreateBill200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateBill200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type CreateBill200ApplicationJSON struct {
	Changes           []CreateBill200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                          `json:"companyId"`
	CompletedOnUtc    *time.Time                                      `json:"completedOnUtc,omitempty"`
	Data              *CreateBill200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                          `json:"dataConnectionKey"`
	DataType          *string                                         `json:"dataType,omitempty"`
	ErrorMessage      *string                                         `json:"errorMessage,omitempty"`
	PushOperationKey  string                                          `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                       `json:"requestedOnUtc"`
	Status            CreateBill200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                             `json:"statusCode"`
	TimeoutInMinutes  *int                                            `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                            `json:"timeoutInSeconds,omitempty"`
	Validation        *CreateBill200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type CreateBillResponse struct {
	ContentType                        string
	StatusCode                         int
	RawResponse                        *http.Response
	CreateBill200ApplicationJSONObject *CreateBill200ApplicationJSON
}
