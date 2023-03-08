package operations

import (
	"net/http"
	"time"
)

type UpdateBillPathParams struct {
	BillID       string `pathParam:"style=simple,explode=false,name=billId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type UpdateBillQueryParams struct {
	ForceUpdate      *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// UpdateBillSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type UpdateBillSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// UpdateBillSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the line item is linked.
type UpdateBillSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// UpdateBillSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type UpdateBillSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type UpdateBillSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type UpdateBillSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type UpdateBillSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	UpdateBillSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       UpdateBillSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	UpdateBillSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable UpdateBillSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	UpdateBillSourceModifiedDateLineItemsTrackingIsBilledToEnumCustomer      UpdateBillSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Customer"
	UpdateBillSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       UpdateBillSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type UpdateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	UpdateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       UpdateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	UpdateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable UpdateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	UpdateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnumCustomer      UpdateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Customer"
	UpdateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       UpdateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type UpdateBillSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type UpdateBillSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []UpdateBillSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *UpdateBillSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   UpdateBillSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo UpdateBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *UpdateBillSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type UpdateBillSourceModifiedDateLineItems struct {
	AccountRef           *UpdateBillSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                     `json:"description,omitempty"`
	DiscountAmount       *float64                                                    `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                    `json:"discountPercentage,omitempty"`
	IsDirectCost         *bool                                                       `json:"isDirectCost,omitempty"`
	ItemRef              *UpdateBillSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                     `json:"quantity"`
	SubTotal             *float64                                                    `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                    `json:"taxAmount,omitempty"`
	TaxRateRef           *UpdateBillSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                    `json:"totalAmount,omitempty"`
	Tracking             *UpdateBillSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []UpdateBillSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                     `json:"unitAmount"`
}

type UpdateBillSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type UpdateBillSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// UpdateBillSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type UpdateBillSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type UpdateBillSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *UpdateBillSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                          `json:"currency,omitempty"`
	CurrencyRate *float64                                                         `json:"currencyRate,omitempty"`
	ID           *string                                                          `json:"id,omitempty"`
	Note         *string                                                          `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                       `json:"paidOnDate,omitempty"`
	Reference    *string                                                          `json:"reference,omitempty"`
	TotalAmount  *float64                                                         `json:"totalAmount,omitempty"`
}

type UpdateBillSourceModifiedDatePaymentAllocations struct {
	Allocation UpdateBillSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    UpdateBillSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type UpdateBillSourceModifiedDatePurchaseOrderRefs struct {
	ID                  *string `json:"id,omitempty"`
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
}

type UpdateBillSourceModifiedDateStatusEnum string

const (
	UpdateBillSourceModifiedDateStatusEnumUnknown       UpdateBillSourceModifiedDateStatusEnum = "Unknown"
	UpdateBillSourceModifiedDateStatusEnumOpen          UpdateBillSourceModifiedDateStatusEnum = "Open"
	UpdateBillSourceModifiedDateStatusEnumPartiallyPaid UpdateBillSourceModifiedDateStatusEnum = "PartiallyPaid"
	UpdateBillSourceModifiedDateStatusEnumPaid          UpdateBillSourceModifiedDateStatusEnum = "Paid"
	UpdateBillSourceModifiedDateStatusEnumVoid          UpdateBillSourceModifiedDateStatusEnum = "Void"
	UpdateBillSourceModifiedDateStatusEnumDraft         UpdateBillSourceModifiedDateStatusEnum = "Draft"
)

// UpdateBillSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type UpdateBillSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// UpdateBillSourceModifiedDateSupplierRef
// Reference to the supplier the bill was received from.
type UpdateBillSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type UpdateBillSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// UpdateBillSourceModifiedDate
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
type UpdateBillSourceModifiedDate struct {
	AmountDue          *float64                                         `json:"amountDue,omitempty"`
	Currency           *string                                          `json:"currency,omitempty"`
	CurrencyRate       *float64                                         `json:"currencyRate,omitempty"`
	DueDate            *time.Time                                       `json:"dueDate,omitempty"`
	ID                 *string                                          `json:"id,omitempty"`
	IssueDate          time.Time                                        `json:"issueDate"`
	LineItems          []UpdateBillSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata           *UpdateBillSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                       `json:"modifiedDate,omitempty"`
	Note               *string                                          `json:"note,omitempty"`
	PaymentAllocations []UpdateBillSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	PurchaseOrderRefs  []UpdateBillSourceModifiedDatePurchaseOrderRefs  `json:"purchaseOrderRefs,omitempty"`
	Reference          *string                                          `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                       `json:"sourceModifiedDate,omitempty"`
	Status             UpdateBillSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal           float64                                          `json:"subTotal"`
	SupplementalData   *UpdateBillSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef        *UpdateBillSourceModifiedDateSupplierRef         `json:"supplierRef,omitempty"`
	TaxAmount          float64                                          `json:"taxAmount"`
	TotalAmount        float64                                          `json:"totalAmount"`
	WithholdingTax     []UpdateBillSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type UpdateBillRequest struct {
	PathParams  UpdateBillPathParams
	QueryParams UpdateBillQueryParams
	Request     *UpdateBillSourceModifiedDate `request:"mediaType=application/json"`
}

type UpdateBill200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type UpdateBill200ApplicationJSONChangesTypeEnum string

const (
	UpdateBill200ApplicationJSONChangesTypeEnumUnknown            UpdateBill200ApplicationJSONChangesTypeEnum = "Unknown"
	UpdateBill200ApplicationJSONChangesTypeEnumCreated            UpdateBill200ApplicationJSONChangesTypeEnum = "Created"
	UpdateBill200ApplicationJSONChangesTypeEnumModified           UpdateBill200ApplicationJSONChangesTypeEnum = "Modified"
	UpdateBill200ApplicationJSONChangesTypeEnumDeleted            UpdateBill200ApplicationJSONChangesTypeEnum = "Deleted"
	UpdateBill200ApplicationJSONChangesTypeEnumAttachmentUploaded UpdateBill200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type UpdateBill200ApplicationJSONChanges struct {
	AttachmentID *string                                                    `json:"attachmentId,omitempty"`
	RecordRef    *UpdateBill200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *UpdateBill200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// UpdateBill200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type UpdateBill200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// UpdateBill200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the line item is linked.
type UpdateBill200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumCustomer      UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Customer"
	UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumCustomer      UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Customer"
	UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type UpdateBill200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *UpdateBill200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                       `json:"description,omitempty"`
	DiscountAmount       *float64                                                                      `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                      `json:"discountPercentage,omitempty"`
	IsDirectCost         *bool                                                                         `json:"isDirectCost,omitempty"`
	ItemRef              *UpdateBill200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                       `json:"quantity"`
	SubTotal             *float64                                                                      `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                      `json:"taxAmount,omitempty"`
	TaxRateRef           *UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                      `json:"totalAmount,omitempty"`
	Tracking             *UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []UpdateBill200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                       `json:"unitAmount"`
}

type UpdateBill200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type UpdateBill200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// UpdateBill200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type UpdateBill200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type UpdateBill200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *UpdateBill200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                            `json:"currency,omitempty"`
	CurrencyRate *float64                                                                           `json:"currencyRate,omitempty"`
	ID           *string                                                                            `json:"id,omitempty"`
	Note         *string                                                                            `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                         `json:"paidOnDate,omitempty"`
	Reference    *string                                                                            `json:"reference,omitempty"`
	TotalAmount  *float64                                                                           `json:"totalAmount,omitempty"`
}

type UpdateBill200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation UpdateBill200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    UpdateBill200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type UpdateBill200ApplicationJSONSourceModifiedDatePurchaseOrderRefs struct {
	ID                  *string `json:"id,omitempty"`
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
}

type UpdateBill200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	UpdateBill200ApplicationJSONSourceModifiedDateStatusEnumUnknown       UpdateBill200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	UpdateBill200ApplicationJSONSourceModifiedDateStatusEnumOpen          UpdateBill200ApplicationJSONSourceModifiedDateStatusEnum = "Open"
	UpdateBill200ApplicationJSONSourceModifiedDateStatusEnumPartiallyPaid UpdateBill200ApplicationJSONSourceModifiedDateStatusEnum = "PartiallyPaid"
	UpdateBill200ApplicationJSONSourceModifiedDateStatusEnumPaid          UpdateBill200ApplicationJSONSourceModifiedDateStatusEnum = "Paid"
	UpdateBill200ApplicationJSONSourceModifiedDateStatusEnumVoid          UpdateBill200ApplicationJSONSourceModifiedDateStatusEnum = "Void"
	UpdateBill200ApplicationJSONSourceModifiedDateStatusEnumDraft         UpdateBill200ApplicationJSONSourceModifiedDateStatusEnum = "Draft"
)

// UpdateBill200ApplicationJSONSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type UpdateBill200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// UpdateBill200ApplicationJSONSourceModifiedDateSupplierRef
// Reference to the supplier the bill was received from.
type UpdateBill200ApplicationJSONSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type UpdateBill200ApplicationJSONSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// UpdateBill200ApplicationJSONSourceModifiedDate
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
type UpdateBill200ApplicationJSONSourceModifiedDate struct {
	AmountDue          *float64                                                           `json:"amountDue,omitempty"`
	Currency           *string                                                            `json:"currency,omitempty"`
	CurrencyRate       *float64                                                           `json:"currencyRate,omitempty"`
	DueDate            *time.Time                                                         `json:"dueDate,omitempty"`
	ID                 *string                                                            `json:"id,omitempty"`
	IssueDate          time.Time                                                          `json:"issueDate"`
	LineItems          []UpdateBill200ApplicationJSONSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata           *UpdateBill200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                         `json:"modifiedDate,omitempty"`
	Note               *string                                                            `json:"note,omitempty"`
	PaymentAllocations []UpdateBill200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	PurchaseOrderRefs  []UpdateBill200ApplicationJSONSourceModifiedDatePurchaseOrderRefs  `json:"purchaseOrderRefs,omitempty"`
	Reference          *string                                                            `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                                         `json:"sourceModifiedDate,omitempty"`
	Status             UpdateBill200ApplicationJSONSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal           float64                                                            `json:"subTotal"`
	SupplementalData   *UpdateBill200ApplicationJSONSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef        *UpdateBill200ApplicationJSONSourceModifiedDateSupplierRef         `json:"supplierRef,omitempty"`
	TaxAmount          float64                                                            `json:"taxAmount"`
	TotalAmount        float64                                                            `json:"totalAmount"`
	WithholdingTax     []UpdateBill200ApplicationJSONSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type UpdateBill200ApplicationJSONStatusEnum string

const (
	UpdateBill200ApplicationJSONStatusEnumPending  UpdateBill200ApplicationJSONStatusEnum = "Pending"
	UpdateBill200ApplicationJSONStatusEnumFailed   UpdateBill200ApplicationJSONStatusEnum = "Failed"
	UpdateBill200ApplicationJSONStatusEnumSuccess  UpdateBill200ApplicationJSONStatusEnum = "Success"
	UpdateBill200ApplicationJSONStatusEnumTimedOut UpdateBill200ApplicationJSONStatusEnum = "TimedOut"
)

type UpdateBill200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// UpdateBill200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type UpdateBill200ApplicationJSONValidation struct {
	Errors   []UpdateBill200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []UpdateBill200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type UpdateBill200ApplicationJSON struct {
	Changes           []UpdateBill200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                          `json:"companyId"`
	CompletedOnUtc    *time.Time                                      `json:"completedOnUtc,omitempty"`
	Data              *UpdateBill200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                          `json:"dataConnectionKey"`
	DataType          *string                                         `json:"dataType,omitempty"`
	ErrorMessage      *string                                         `json:"errorMessage,omitempty"`
	PushOperationKey  string                                          `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                       `json:"requestedOnUtc"`
	Status            UpdateBill200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                             `json:"statusCode"`
	TimeoutInMinutes  *int                                            `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                            `json:"timeoutInSeconds,omitempty"`
	Validation        *UpdateBill200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type UpdateBillResponse struct {
	ContentType                        string
	StatusCode                         int
	RawResponse                        *http.Response
	UpdateBill200ApplicationJSONObject *UpdateBill200ApplicationJSON
}
