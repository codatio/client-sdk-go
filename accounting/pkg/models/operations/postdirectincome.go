package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostDirectIncomePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostDirectIncomeQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// PostDirectIncomeSourceModifiedDateContactRef
// A customer or supplier associated with the direct income.
type PostDirectIncomeSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// PostDirectIncomeSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type PostDirectIncomeSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostDirectIncomeSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the direct cost is linked.
type PostDirectIncomeSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PostDirectIncomeSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type PostDirectIncomeSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type PostDirectIncomeSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostDirectIncomeSourceModifiedDateLineItems struct {
	AccountRef           *PostDirectIncomeSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                           `json:"description,omitempty"`
	DiscountAmount       *float64                                                          `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                          `json:"discountPercentage,omitempty"`
	ItemRef              *PostDirectIncomeSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                           `json:"quantity"`
	SubTotal             *float64                                                          `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                          `json:"taxAmount,omitempty"`
	TaxRateRef           *PostDirectIncomeSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                          `json:"totalAmount,omitempty"`
	TrackingCategoryRefs []PostDirectIncomeSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                           `json:"unitAmount"`
}

type PostDirectIncomeSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostDirectIncomeSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// PostDirectIncomeSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type PostDirectIncomeSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PostDirectIncomeSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *PostDirectIncomeSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                `json:"currency,omitempty"`
	CurrencyRate *float64                                                               `json:"currencyRate,omitempty"`
	ID           *string                                                                `json:"id,omitempty"`
	Note         *string                                                                `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                             `json:"paidOnDate,omitempty"`
	Reference    *string                                                                `json:"reference,omitempty"`
	TotalAmount  *float64                                                               `json:"totalAmount,omitempty"`
}

type PostDirectIncomeSourceModifiedDatePaymentAllocations struct {
	Allocation PostDirectIncomeSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    PostDirectIncomeSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type PostDirectIncomeSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// PostDirectIncomeSourceModifiedDate
// > **Language tip:**  Direct incomes may also be referred to as **Receive transactions**, **Receive money transactions**, **Sales receipts**, or **Cash sales** in various accounting platforms.
//
// > View the coverage for direct incomes in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Direct incomes include:
//
// - Selling an item directly to a contact, and receiving payment at the point of the sale.
// - Refunding an item in cash to a contact.
// - Depositing money into a bank account.
//
// From the Direct Incomes endpoints, you can:
//
// - [Get a list of all direct incomes for a specific company](https://api-uat.codat.io/swagger/index.html#/DirectIncomes/get_companies__companyId__connections__connectionId__data_directIncomes)
// - [Get a single direct income for a specific company and connection](https://api-uat.codat.io/swagger/index.html#/DirectIncomes/get_companies__companyId__connections__connectionId__data_directIncomes__directIncomeId_)
// - [Add a new direct income to a specific company's accounting package](https://api-uat.codat.io/swagger/index.html#/DirectIncomes/post_companies__companyId__connections__connectionId__push_directIncomes)
//
// Direct incomes is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type PostDirectIncomeSourceModifiedDate struct {
	ContactRef         *PostDirectIncomeSourceModifiedDateContactRef          `json:"contactRef,omitempty"`
	Currency           string                                                 `json:"currency"`
	CurrencyRate       *float64                                               `json:"currencyRate,omitempty"`
	ID                 *string                                                `json:"id,omitempty"`
	IssueDate          time.Time                                              `json:"issueDate"`
	LineItems          []PostDirectIncomeSourceModifiedDateLineItems          `json:"lineItems"`
	Metadata           *PostDirectIncomeSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                             `json:"modifiedDate,omitempty"`
	Note               *string                                                `json:"note,omitempty"`
	PaymentAllocations []PostDirectIncomeSourceModifiedDatePaymentAllocations `json:"paymentAllocations"`
	Reference          *string                                                `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                             `json:"sourceModifiedDate,omitempty"`
	SubTotal           float64                                                `json:"subTotal"`
	SupplementalData   *PostDirectIncomeSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TaxAmount          float64                                                `json:"taxAmount"`
	TotalAmount        float64                                                `json:"totalAmount"`
}

type PostDirectIncomeSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostDirectIncomeRequest struct {
	PathParams  PostDirectIncomePathParams
	QueryParams PostDirectIncomeQueryParams
	Request     *PostDirectIncomeSourceModifiedDate `request:"mediaType=application/json"`
	Security    PostDirectIncomeSecurity
}

type PostDirectIncome200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostDirectIncome200ApplicationJSONChangesTypeEnum string

const (
	PostDirectIncome200ApplicationJSONChangesTypeEnumUnknown            PostDirectIncome200ApplicationJSONChangesTypeEnum = "Unknown"
	PostDirectIncome200ApplicationJSONChangesTypeEnumCreated            PostDirectIncome200ApplicationJSONChangesTypeEnum = "Created"
	PostDirectIncome200ApplicationJSONChangesTypeEnumModified           PostDirectIncome200ApplicationJSONChangesTypeEnum = "Modified"
	PostDirectIncome200ApplicationJSONChangesTypeEnumDeleted            PostDirectIncome200ApplicationJSONChangesTypeEnum = "Deleted"
	PostDirectIncome200ApplicationJSONChangesTypeEnumAttachmentUploaded PostDirectIncome200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostDirectIncome200ApplicationJSONChanges struct {
	AttachmentID *string                                                          `json:"attachmentId,omitempty"`
	RecordRef    *PostDirectIncome200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostDirectIncome200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// PostDirectIncome200ApplicationJSONSourceModifiedDateContactRef
// A customer or supplier associated with the direct income.
type PostDirectIncome200ApplicationJSONSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// PostDirectIncome200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type PostDirectIncome200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostDirectIncome200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the direct cost is linked.
type PostDirectIncome200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PostDirectIncome200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type PostDirectIncome200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type PostDirectIncome200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostDirectIncome200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *PostDirectIncome200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                             `json:"description,omitempty"`
	DiscountAmount       *float64                                                                            `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                            `json:"discountPercentage,omitempty"`
	ItemRef              *PostDirectIncome200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                             `json:"quantity"`
	SubTotal             *float64                                                                            `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                            `json:"taxAmount,omitempty"`
	TaxRateRef           *PostDirectIncome200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                            `json:"totalAmount,omitempty"`
	TrackingCategoryRefs []PostDirectIncome200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                             `json:"unitAmount"`
}

type PostDirectIncome200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// PostDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type PostDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PostDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *PostDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                                  `json:"currency,omitempty"`
	CurrencyRate *float64                                                                                 `json:"currencyRate,omitempty"`
	ID           *string                                                                                  `json:"id,omitempty"`
	Note         *string                                                                                  `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                               `json:"paidOnDate,omitempty"`
	Reference    *string                                                                                  `json:"reference,omitempty"`
	TotalAmount  *float64                                                                                 `json:"totalAmount,omitempty"`
}

type PostDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation PostDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    PostDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type PostDirectIncome200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// PostDirectIncome200ApplicationJSONSourceModifiedDate
// > **Language tip:**  Direct incomes may also be referred to as **Receive transactions**, **Receive money transactions**, **Sales receipts**, or **Cash sales** in various accounting platforms.
//
// > View the coverage for direct incomes in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Direct incomes include:
//
// - Selling an item directly to a contact, and receiving payment at the point of the sale.
// - Refunding an item in cash to a contact.
// - Depositing money into a bank account.
//
// From the Direct Incomes endpoints, you can:
//
// - [Get a list of all direct incomes for a specific company](https://api-uat.codat.io/swagger/index.html#/DirectIncomes/get_companies__companyId__connections__connectionId__data_directIncomes)
// - [Get a single direct income for a specific company and connection](https://api-uat.codat.io/swagger/index.html#/DirectIncomes/get_companies__companyId__connections__connectionId__data_directIncomes__directIncomeId_)
// - [Add a new direct income to a specific company's accounting package](https://api-uat.codat.io/swagger/index.html#/DirectIncomes/post_companies__companyId__connections__connectionId__push_directIncomes)
//
// Direct incomes is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type PostDirectIncome200ApplicationJSONSourceModifiedDate struct {
	ContactRef         *PostDirectIncome200ApplicationJSONSourceModifiedDateContactRef          `json:"contactRef,omitempty"`
	Currency           string                                                                   `json:"currency"`
	CurrencyRate       *float64                                                                 `json:"currencyRate,omitempty"`
	ID                 *string                                                                  `json:"id,omitempty"`
	IssueDate          time.Time                                                                `json:"issueDate"`
	LineItems          []PostDirectIncome200ApplicationJSONSourceModifiedDateLineItems          `json:"lineItems"`
	Metadata           *PostDirectIncome200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                               `json:"modifiedDate,omitempty"`
	Note               *string                                                                  `json:"note,omitempty"`
	PaymentAllocations []PostDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations"`
	Reference          *string                                                                  `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                                               `json:"sourceModifiedDate,omitempty"`
	SubTotal           float64                                                                  `json:"subTotal"`
	SupplementalData   *PostDirectIncome200ApplicationJSONSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TaxAmount          float64                                                                  `json:"taxAmount"`
	TotalAmount        float64                                                                  `json:"totalAmount"`
}

type PostDirectIncome200ApplicationJSONStatusEnum string

const (
	PostDirectIncome200ApplicationJSONStatusEnumPending  PostDirectIncome200ApplicationJSONStatusEnum = "Pending"
	PostDirectIncome200ApplicationJSONStatusEnumFailed   PostDirectIncome200ApplicationJSONStatusEnum = "Failed"
	PostDirectIncome200ApplicationJSONStatusEnumSuccess  PostDirectIncome200ApplicationJSONStatusEnum = "Success"
	PostDirectIncome200ApplicationJSONStatusEnumTimedOut PostDirectIncome200ApplicationJSONStatusEnum = "TimedOut"
)

type PostDirectIncome200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostDirectIncome200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostDirectIncome200ApplicationJSONValidation struct {
	Errors   []PostDirectIncome200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PostDirectIncome200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PostDirectIncome200ApplicationJSON struct {
	Changes           []PostDirectIncome200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                                `json:"companyId"`
	CompletedOnUtc    *time.Time                                            `json:"completedOnUtc,omitempty"`
	Data              *PostDirectIncome200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                                `json:"dataConnectionKey"`
	DataType          *string                                               `json:"dataType,omitempty"`
	ErrorMessage      *string                                               `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                             `json:"requestedOnUtc"`
	Status            PostDirectIncome200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                   `json:"statusCode"`
	TimeoutInMinutes  *int                                                  `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                  `json:"timeoutInSeconds,omitempty"`
	Validation        *PostDirectIncome200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PostDirectIncomeResponse struct {
	ContentType                              string
	StatusCode                               int
	PostDirectIncome200ApplicationJSONObject *PostDirectIncome200ApplicationJSON
}
