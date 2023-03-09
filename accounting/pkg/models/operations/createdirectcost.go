package operations

import (
	"net/http"
	"time"
)

type CreateDirectCostPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type CreateDirectCostQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// CreateDirectCostSourceModifiedDateContactRef
// A customer or supplier associated with the direct cost.
type CreateDirectCostSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// CreateDirectCostSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type CreateDirectCostSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateDirectCostSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the direct cost is linked.
type CreateDirectCostSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateDirectCostSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the the line item is linked.
type CreateDirectCostSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type CreateDirectCostSourceModifiedDateLineItemsTracking struct {
	InvoiceTo  *string  `json:"invoiceTo,omitempty"`
	RecordRefs []string `json:"recordRefs"`
}

type CreateDirectCostSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateDirectCostSourceModifiedDateLineItems struct {
	AccountRef           *CreateDirectCostSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                           `json:"description,omitempty"`
	DiscountAmount       *float64                                                          `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                          `json:"discountPercentage,omitempty"`
	ItemRef              *CreateDirectCostSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                           `json:"quantity"`
	SubTotal             *float64                                                          `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                          `json:"taxAmount,omitempty"`
	TaxRateRef           *CreateDirectCostSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                          `json:"totalAmount,omitempty"`
	Tracking             *CreateDirectCostSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []CreateDirectCostSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                           `json:"unitAmount"`
}

type CreateDirectCostSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateDirectCostSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// CreateDirectCostSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type CreateDirectCostSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CreateDirectCostSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *CreateDirectCostSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                `json:"currency,omitempty"`
	CurrencyRate *float64                                                               `json:"currencyRate,omitempty"`
	ID           *string                                                                `json:"id,omitempty"`
	Note         *string                                                                `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                             `json:"paidOnDate,omitempty"`
	Reference    *string                                                                `json:"reference,omitempty"`
	TotalAmount  *float64                                                               `json:"totalAmount,omitempty"`
}

type CreateDirectCostSourceModifiedDatePaymentAllocations struct {
	Allocation CreateDirectCostSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    CreateDirectCostSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

// CreateDirectCostSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateDirectCostSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateDirectCostSourceModifiedDate
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
type CreateDirectCostSourceModifiedDate struct {
	ContactRef         *CreateDirectCostSourceModifiedDateContactRef          `json:"contactRef,omitempty"`
	Currency           string                                                 `json:"currency"`
	CurrencyRate       *float64                                               `json:"currencyRate,omitempty"`
	ID                 *string                                                `json:"id,omitempty"`
	IssueDate          time.Time                                              `json:"issueDate"`
	LineItems          []CreateDirectCostSourceModifiedDateLineItems          `json:"lineItems"`
	Metadata           *CreateDirectCostSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                             `json:"modifiedDate,omitempty"`
	Note               *string                                                `json:"note,omitempty"`
	PaymentAllocations []CreateDirectCostSourceModifiedDatePaymentAllocations `json:"paymentAllocations"`
	Reference          *string                                                `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                             `json:"sourceModifiedDate,omitempty"`
	SubTotal           float64                                                `json:"subTotal"`
	SupplementalData   *CreateDirectCostSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TaxAmount          float64                                                `json:"taxAmount"`
	TotalAmount        float64                                                `json:"totalAmount"`
}

type CreateDirectCostRequest struct {
	PathParams  CreateDirectCostPathParams
	QueryParams CreateDirectCostQueryParams
	Request     *CreateDirectCostSourceModifiedDate `request:"mediaType=application/json"`
}

type CreateDirectCost200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateDirectCost200ApplicationJSONChangesTypeEnum string

const (
	CreateDirectCost200ApplicationJSONChangesTypeEnumUnknown            CreateDirectCost200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateDirectCost200ApplicationJSONChangesTypeEnumCreated            CreateDirectCost200ApplicationJSONChangesTypeEnum = "Created"
	CreateDirectCost200ApplicationJSONChangesTypeEnumModified           CreateDirectCost200ApplicationJSONChangesTypeEnum = "Modified"
	CreateDirectCost200ApplicationJSONChangesTypeEnumDeleted            CreateDirectCost200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateDirectCost200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateDirectCost200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type CreateDirectCost200ApplicationJSONChanges struct {
	AttachmentID *string                                                          `json:"attachmentId,omitempty"`
	RecordRef    *CreateDirectCost200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateDirectCost200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// CreateDirectCost200ApplicationJSONSourceModifiedDateContactRef
// A customer or supplier associated with the direct cost.
type CreateDirectCost200ApplicationJSONSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// CreateDirectCost200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type CreateDirectCost200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateDirectCost200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the direct cost is linked.
type CreateDirectCost200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateDirectCost200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the the line item is linked.
type CreateDirectCost200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type CreateDirectCost200ApplicationJSONSourceModifiedDateLineItemsTracking struct {
	InvoiceTo  *string  `json:"invoiceTo,omitempty"`
	RecordRefs []string `json:"recordRefs"`
}

type CreateDirectCost200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateDirectCost200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *CreateDirectCost200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                             `json:"description,omitempty"`
	DiscountAmount       *float64                                                                            `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                            `json:"discountPercentage,omitempty"`
	ItemRef              *CreateDirectCost200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                             `json:"quantity"`
	SubTotal             *float64                                                                            `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                            `json:"taxAmount,omitempty"`
	TaxRateRef           *CreateDirectCost200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                            `json:"totalAmount,omitempty"`
	Tracking             *CreateDirectCost200ApplicationJSONSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []CreateDirectCost200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                             `json:"unitAmount"`
}

type CreateDirectCost200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// CreateDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type CreateDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CreateDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *CreateDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                                  `json:"currency,omitempty"`
	CurrencyRate *float64                                                                                 `json:"currencyRate,omitempty"`
	ID           *string                                                                                  `json:"id,omitempty"`
	Note         *string                                                                                  `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                               `json:"paidOnDate,omitempty"`
	Reference    *string                                                                                  `json:"reference,omitempty"`
	TotalAmount  *float64                                                                                 `json:"totalAmount,omitempty"`
}

type CreateDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation CreateDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    CreateDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

// CreateDirectCost200ApplicationJSONSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateDirectCost200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateDirectCost200ApplicationJSONSourceModifiedDate
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
type CreateDirectCost200ApplicationJSONSourceModifiedDate struct {
	ContactRef         *CreateDirectCost200ApplicationJSONSourceModifiedDateContactRef          `json:"contactRef,omitempty"`
	Currency           string                                                                   `json:"currency"`
	CurrencyRate       *float64                                                                 `json:"currencyRate,omitempty"`
	ID                 *string                                                                  `json:"id,omitempty"`
	IssueDate          time.Time                                                                `json:"issueDate"`
	LineItems          []CreateDirectCost200ApplicationJSONSourceModifiedDateLineItems          `json:"lineItems"`
	Metadata           *CreateDirectCost200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                               `json:"modifiedDate,omitempty"`
	Note               *string                                                                  `json:"note,omitempty"`
	PaymentAllocations []CreateDirectCost200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations"`
	Reference          *string                                                                  `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                                               `json:"sourceModifiedDate,omitempty"`
	SubTotal           float64                                                                  `json:"subTotal"`
	SupplementalData   *CreateDirectCost200ApplicationJSONSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TaxAmount          float64                                                                  `json:"taxAmount"`
	TotalAmount        float64                                                                  `json:"totalAmount"`
}

type CreateDirectCost200ApplicationJSONStatusEnum string

const (
	CreateDirectCost200ApplicationJSONStatusEnumPending  CreateDirectCost200ApplicationJSONStatusEnum = "Pending"
	CreateDirectCost200ApplicationJSONStatusEnumFailed   CreateDirectCost200ApplicationJSONStatusEnum = "Failed"
	CreateDirectCost200ApplicationJSONStatusEnumSuccess  CreateDirectCost200ApplicationJSONStatusEnum = "Success"
	CreateDirectCost200ApplicationJSONStatusEnumTimedOut CreateDirectCost200ApplicationJSONStatusEnum = "TimedOut"
)

type CreateDirectCost200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateDirectCost200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateDirectCost200ApplicationJSONValidation struct {
	Errors   []CreateDirectCost200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateDirectCost200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type CreateDirectCost200ApplicationJSON struct {
	Changes           []CreateDirectCost200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                                `json:"companyId"`
	CompletedOnUtc    *time.Time                                            `json:"completedOnUtc,omitempty"`
	Data              *CreateDirectCost200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                                `json:"dataConnectionKey"`
	DataType          *string                                               `json:"dataType,omitempty"`
	ErrorMessage      *string                                               `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                             `json:"requestedOnUtc"`
	Status            CreateDirectCost200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                   `json:"statusCode"`
	TimeoutInMinutes  *int                                                  `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                  `json:"timeoutInSeconds,omitempty"`
	Validation        *CreateDirectCost200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type CreateDirectCostResponse struct {
	ContentType                              string
	StatusCode                               int
	RawResponse                              *http.Response
	CreateDirectCost200ApplicationJSONObject *CreateDirectCost200ApplicationJSON
}
