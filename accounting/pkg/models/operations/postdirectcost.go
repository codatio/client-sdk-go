package operations

import (
	"net/http"
	"time"
)

type PostDirectCostPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostDirectCostQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// PostDirectCostSourceModifiedDateContactRef
// A customer or supplier associated with the direct cost.
type PostDirectCostSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// PostDirectCostSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type PostDirectCostSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostDirectCostSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the direct cost is linked.
type PostDirectCostSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PostDirectCostSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the the line item is linked.
type PostDirectCostSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type PostDirectCostSourceModifiedDateLineItemsTracking struct {
	InvoiceTo  *string  `json:"invoiceTo,omitempty"`
	RecordRefs []string `json:"recordRefs"`
}

type PostDirectCostSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostDirectCostSourceModifiedDateLineItems struct {
	AccountRef           *PostDirectCostSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                         `json:"description,omitempty"`
	DiscountAmount       *float64                                                        `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                        `json:"discountPercentage,omitempty"`
	ItemRef              *PostDirectCostSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                         `json:"quantity"`
	SubTotal             *float64                                                        `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                        `json:"taxAmount,omitempty"`
	TaxRateRef           *PostDirectCostSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                        `json:"totalAmount,omitempty"`
	Tracking             *PostDirectCostSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []PostDirectCostSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                         `json:"unitAmount"`
}

type PostDirectCostSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostDirectCostSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// PostDirectCostSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type PostDirectCostSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PostDirectCostSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *PostDirectCostSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                              `json:"currency,omitempty"`
	CurrencyRate *float64                                                             `json:"currencyRate,omitempty"`
	ID           *string                                                              `json:"id,omitempty"`
	Note         *string                                                              `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                           `json:"paidOnDate,omitempty"`
	Reference    *string                                                              `json:"reference,omitempty"`
	TotalAmount  *float64                                                             `json:"totalAmount,omitempty"`
}

type PostDirectCostSourceModifiedDatePaymentAllocations struct {
	Allocation PostDirectCostSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    PostDirectCostSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type PostDirectCostSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// PostDirectCostSourceModifiedDate
// > **Language tip: ** Direct costs may also be referred to as **Spend transactions**, **Spend money transactions**, or **Payments** in various accounting platforms.
//
// > View the coverage for direct costs in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Direct costs include:
//   - Purchasing an item and paying it off at the point of the purchase
//   - Receiving cash from a refunded item if the refund is made by the supplier
//   - Withdrawing money from a bank account
//   - Writing a cheque
//
// From the Direct Costs endpoints, you can:
//
//   - [Get a list of all direct costs for a specific company ](https://api.codat.io/swagger/index.html#/DirectCosts/get_companies__companyId__connections__connectionId__data_directCosts)
//   - [Get a single direct cost for a specific company ](https://api.codat.io/swagger/index.html#/DirectCosts/get_companies__companyId__connections__connectionId__data_directCosts__directCostId_)
//   - [Add a new direct cost to a specific company's accounting package](https://api.codat.io/swagger/index.html#/DirectCosts/post_companies__companyId__connections__connectionId__push_directCosts)
//
// Direct costs is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type PostDirectCostSourceModifiedDate struct {
	ContactRef         *PostDirectCostSourceModifiedDateContactRef          `json:"contactRef,omitempty"`
	Currency           string                                               `json:"currency"`
	CurrencyRate       *float64                                             `json:"currencyRate,omitempty"`
	ID                 *string                                              `json:"id,omitempty"`
	IssueDate          time.Time                                            `json:"issueDate"`
	LineItems          []PostDirectCostSourceModifiedDateLineItems          `json:"lineItems"`
	Metadata           *PostDirectCostSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                           `json:"modifiedDate,omitempty"`
	Note               *string                                              `json:"note,omitempty"`
	PaymentAllocations []PostDirectCostSourceModifiedDatePaymentAllocations `json:"paymentAllocations"`
	Reference          *string                                              `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                           `json:"sourceModifiedDate,omitempty"`
	SubTotal           float64                                              `json:"subTotal"`
	SupplementalData   *PostDirectCostSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TaxAmount          float64                                              `json:"taxAmount"`
	TotalAmount        float64                                              `json:"totalAmount"`
}

type PostDirectCostRequest struct {
	PathParams  PostDirectCostPathParams
	QueryParams PostDirectCostQueryParams
	Request     *PostDirectCostSourceModifiedDate `request:"mediaType=application/json"`
}

type PostDirectCost200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostDirectCost200ApplicationJSONChangesTypeEnum string

const (
	PostDirectCost200ApplicationJSONChangesTypeEnumUnknown            PostDirectCost200ApplicationJSONChangesTypeEnum = "Unknown"
	PostDirectCost200ApplicationJSONChangesTypeEnumCreated            PostDirectCost200ApplicationJSONChangesTypeEnum = "Created"
	PostDirectCost200ApplicationJSONChangesTypeEnumModified           PostDirectCost200ApplicationJSONChangesTypeEnum = "Modified"
	PostDirectCost200ApplicationJSONChangesTypeEnumDeleted            PostDirectCost200ApplicationJSONChangesTypeEnum = "Deleted"
	PostDirectCost200ApplicationJSONChangesTypeEnumAttachmentUploaded PostDirectCost200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostDirectCost200ApplicationJSONChanges struct {
	AttachmentID *string                                                        `json:"attachmentId,omitempty"`
	RecordRef    *PostDirectCost200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostDirectCost200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// PostDirectCost200ApplicationJSONSourceModifiedDateContactRef
// A customer or supplier associated with the direct cost.
type PostDirectCost200ApplicationJSONSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// PostDirectCost200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type PostDirectCost200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostDirectCost200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the direct cost is linked.
type PostDirectCost200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PostDirectCost200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the the line item is linked.
type PostDirectCost200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type PostDirectCost200ApplicationJSONSourceModifiedDateLineItemsTracking struct {
	InvoiceTo  *string  `json:"invoiceTo,omitempty"`
	RecordRefs []string `json:"recordRefs"`
}

type PostDirectCost200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostDirectCost200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *PostDirectCost200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                           `json:"description,omitempty"`
	DiscountAmount       *float64                                                                          `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                          `json:"discountPercentage,omitempty"`
	ItemRef              *PostDirectCost200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                           `json:"quantity"`
	SubTotal             *float64                                                                          `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                          `json:"taxAmount,omitempty"`
	TaxRateRef           *PostDirectCost200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                          `json:"totalAmount,omitempty"`
	Tracking             *PostDirectCost200ApplicationJSONSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []PostDirectCost200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                           `json:"unitAmount"`
}

type PostDirectCost200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// PostDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type PostDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PostDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *PostDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                                `json:"currency,omitempty"`
	CurrencyRate *float64                                                                               `json:"currencyRate,omitempty"`
	ID           *string                                                                                `json:"id,omitempty"`
	Note         *string                                                                                `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                             `json:"paidOnDate,omitempty"`
	Reference    *string                                                                                `json:"reference,omitempty"`
	TotalAmount  *float64                                                                               `json:"totalAmount,omitempty"`
}

type PostDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation PostDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    PostDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type PostDirectCost200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// PostDirectCost200ApplicationJSONSourceModifiedDate
// > **Language tip: ** Direct costs may also be referred to as **Spend transactions**, **Spend money transactions**, or **Payments** in various accounting platforms.
//
// > View the coverage for direct costs in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Direct costs include:
//   - Purchasing an item and paying it off at the point of the purchase
//   - Receiving cash from a refunded item if the refund is made by the supplier
//   - Withdrawing money from a bank account
//   - Writing a cheque
//
// From the Direct Costs endpoints, you can:
//
//   - [Get a list of all direct costs for a specific company ](https://api.codat.io/swagger/index.html#/DirectCosts/get_companies__companyId__connections__connectionId__data_directCosts)
//   - [Get a single direct cost for a specific company ](https://api.codat.io/swagger/index.html#/DirectCosts/get_companies__companyId__connections__connectionId__data_directCosts__directCostId_)
//   - [Add a new direct cost to a specific company's accounting package](https://api.codat.io/swagger/index.html#/DirectCosts/post_companies__companyId__connections__connectionId__push_directCosts)
//
// Direct costs is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type PostDirectCost200ApplicationJSONSourceModifiedDate struct {
	ContactRef         *PostDirectCost200ApplicationJSONSourceModifiedDateContactRef          `json:"contactRef,omitempty"`
	Currency           string                                                                 `json:"currency"`
	CurrencyRate       *float64                                                               `json:"currencyRate,omitempty"`
	ID                 *string                                                                `json:"id,omitempty"`
	IssueDate          time.Time                                                              `json:"issueDate"`
	LineItems          []PostDirectCost200ApplicationJSONSourceModifiedDateLineItems          `json:"lineItems"`
	Metadata           *PostDirectCost200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                             `json:"modifiedDate,omitempty"`
	Note               *string                                                                `json:"note,omitempty"`
	PaymentAllocations []PostDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations"`
	Reference          *string                                                                `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                                             `json:"sourceModifiedDate,omitempty"`
	SubTotal           float64                                                                `json:"subTotal"`
	SupplementalData   *PostDirectCost200ApplicationJSONSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TaxAmount          float64                                                                `json:"taxAmount"`
	TotalAmount        float64                                                                `json:"totalAmount"`
}

type PostDirectCost200ApplicationJSONStatusEnum string

const (
	PostDirectCost200ApplicationJSONStatusEnumPending  PostDirectCost200ApplicationJSONStatusEnum = "Pending"
	PostDirectCost200ApplicationJSONStatusEnumFailed   PostDirectCost200ApplicationJSONStatusEnum = "Failed"
	PostDirectCost200ApplicationJSONStatusEnumSuccess  PostDirectCost200ApplicationJSONStatusEnum = "Success"
	PostDirectCost200ApplicationJSONStatusEnumTimedOut PostDirectCost200ApplicationJSONStatusEnum = "TimedOut"
)

type PostDirectCost200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostDirectCost200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostDirectCost200ApplicationJSONValidation struct {
	Errors   []PostDirectCost200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PostDirectCost200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PostDirectCost200ApplicationJSON struct {
	Changes           []PostDirectCost200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                              `json:"companyId"`
	CompletedOnUtc    *time.Time                                          `json:"completedOnUtc,omitempty"`
	Data              *PostDirectCost200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                              `json:"dataConnectionKey"`
	DataType          *string                                             `json:"dataType,omitempty"`
	ErrorMessage      *string                                             `json:"errorMessage,omitempty"`
	PushOperationKey  string                                              `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                           `json:"requestedOnUtc"`
	Status            PostDirectCost200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                 `json:"statusCode"`
	TimeoutInMinutes  *int                                                `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                `json:"timeoutInSeconds,omitempty"`
	Validation        *PostDirectCost200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PostDirectCostResponse struct {
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
	PostDirectCost200ApplicationJSONObject *PostDirectCost200ApplicationJSON
}
