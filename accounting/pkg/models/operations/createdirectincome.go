package operations

import (
	"net/http"
	"time"
)

type CreateDirectIncomePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type CreateDirectIncomeQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// CreateDirectIncomeSourceModifiedDateContactRef
// A customer or supplier associated with the direct income.
type CreateDirectIncomeSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// CreateDirectIncomeSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type CreateDirectIncomeSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateDirectIncomeSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the direct cost is linked.
type CreateDirectIncomeSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateDirectIncomeSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type CreateDirectIncomeSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type CreateDirectIncomeSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateDirectIncomeSourceModifiedDateLineItems struct {
	AccountRef           *CreateDirectIncomeSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                             `json:"description,omitempty"`
	DiscountAmount       *float64                                                            `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                            `json:"discountPercentage,omitempty"`
	ItemRef              *CreateDirectIncomeSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                             `json:"quantity"`
	SubTotal             *float64                                                            `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                            `json:"taxAmount,omitempty"`
	TaxRateRef           *CreateDirectIncomeSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                            `json:"totalAmount,omitempty"`
	TrackingCategoryRefs []CreateDirectIncomeSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                             `json:"unitAmount"`
}

type CreateDirectIncomeSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateDirectIncomeSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// CreateDirectIncomeSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type CreateDirectIncomeSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CreateDirectIncomeSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *CreateDirectIncomeSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                  `json:"currency,omitempty"`
	CurrencyRate *float64                                                                 `json:"currencyRate,omitempty"`
	ID           *string                                                                  `json:"id,omitempty"`
	Note         *string                                                                  `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                               `json:"paidOnDate,omitempty"`
	Reference    *string                                                                  `json:"reference,omitempty"`
	TotalAmount  *float64                                                                 `json:"totalAmount,omitempty"`
}

type CreateDirectIncomeSourceModifiedDatePaymentAllocations struct {
	Allocation CreateDirectIncomeSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    CreateDirectIncomeSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

// CreateDirectIncomeSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateDirectIncomeSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateDirectIncomeSourceModifiedDate
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
type CreateDirectIncomeSourceModifiedDate struct {
	ContactRef         *CreateDirectIncomeSourceModifiedDateContactRef          `json:"contactRef,omitempty"`
	Currency           string                                                   `json:"currency"`
	CurrencyRate       *float64                                                 `json:"currencyRate,omitempty"`
	ID                 *string                                                  `json:"id,omitempty"`
	IssueDate          time.Time                                                `json:"issueDate"`
	LineItems          []CreateDirectIncomeSourceModifiedDateLineItems          `json:"lineItems"`
	Metadata           *CreateDirectIncomeSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                               `json:"modifiedDate,omitempty"`
	Note               *string                                                  `json:"note,omitempty"`
	PaymentAllocations []CreateDirectIncomeSourceModifiedDatePaymentAllocations `json:"paymentAllocations"`
	Reference          *string                                                  `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                               `json:"sourceModifiedDate,omitempty"`
	SubTotal           float64                                                  `json:"subTotal"`
	SupplementalData   *CreateDirectIncomeSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TaxAmount          float64                                                  `json:"taxAmount"`
	TotalAmount        float64                                                  `json:"totalAmount"`
}

type CreateDirectIncomeRequest struct {
	PathParams  CreateDirectIncomePathParams
	QueryParams CreateDirectIncomeQueryParams
	Request     *CreateDirectIncomeSourceModifiedDate `request:"mediaType=application/json"`
}

type CreateDirectIncome200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateDirectIncome200ApplicationJSONChangesTypeEnum string

const (
	CreateDirectIncome200ApplicationJSONChangesTypeEnumUnknown            CreateDirectIncome200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateDirectIncome200ApplicationJSONChangesTypeEnumCreated            CreateDirectIncome200ApplicationJSONChangesTypeEnum = "Created"
	CreateDirectIncome200ApplicationJSONChangesTypeEnumModified           CreateDirectIncome200ApplicationJSONChangesTypeEnum = "Modified"
	CreateDirectIncome200ApplicationJSONChangesTypeEnumDeleted            CreateDirectIncome200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateDirectIncome200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateDirectIncome200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type CreateDirectIncome200ApplicationJSONChanges struct {
	AttachmentID *string                                                            `json:"attachmentId,omitempty"`
	RecordRef    *CreateDirectIncome200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateDirectIncome200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// CreateDirectIncome200ApplicationJSONSourceModifiedDateContactRef
// A customer or supplier associated with the direct income.
type CreateDirectIncome200ApplicationJSONSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the direct cost is linked.
type CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                               `json:"description,omitempty"`
	DiscountAmount       *float64                                                                              `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                              `json:"discountPercentage,omitempty"`
	ItemRef              *CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                               `json:"quantity"`
	SubTotal             *float64                                                                              `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                              `json:"taxAmount,omitempty"`
	TaxRateRef           *CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                              `json:"totalAmount,omitempty"`
	TrackingCategoryRefs []CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                               `json:"unitAmount"`
}

type CreateDirectIncome200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                                    `json:"currency,omitempty"`
	CurrencyRate *float64                                                                                   `json:"currencyRate,omitempty"`
	ID           *string                                                                                    `json:"id,omitempty"`
	Note         *string                                                                                    `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                                 `json:"paidOnDate,omitempty"`
	Reference    *string                                                                                    `json:"reference,omitempty"`
	TotalAmount  *float64                                                                                   `json:"totalAmount,omitempty"`
}

type CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocations struct {
	Allocation CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

// CreateDirectIncome200ApplicationJSONSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateDirectIncome200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateDirectIncome200ApplicationJSONSourceModifiedDate
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
type CreateDirectIncome200ApplicationJSONSourceModifiedDate struct {
	ContactRef         *CreateDirectIncome200ApplicationJSONSourceModifiedDateContactRef          `json:"contactRef,omitempty"`
	Currency           string                                                                     `json:"currency"`
	CurrencyRate       *float64                                                                   `json:"currencyRate,omitempty"`
	ID                 *string                                                                    `json:"id,omitempty"`
	IssueDate          time.Time                                                                  `json:"issueDate"`
	LineItems          []CreateDirectIncome200ApplicationJSONSourceModifiedDateLineItems          `json:"lineItems"`
	Metadata           *CreateDirectIncome200ApplicationJSONSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                                 `json:"modifiedDate,omitempty"`
	Note               *string                                                                    `json:"note,omitempty"`
	PaymentAllocations []CreateDirectIncome200ApplicationJSONSourceModifiedDatePaymentAllocations `json:"paymentAllocations"`
	Reference          *string                                                                    `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                                                 `json:"sourceModifiedDate,omitempty"`
	SubTotal           float64                                                                    `json:"subTotal"`
	SupplementalData   *CreateDirectIncome200ApplicationJSONSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TaxAmount          float64                                                                    `json:"taxAmount"`
	TotalAmount        float64                                                                    `json:"totalAmount"`
}

type CreateDirectIncome200ApplicationJSONStatusEnum string

const (
	CreateDirectIncome200ApplicationJSONStatusEnumPending  CreateDirectIncome200ApplicationJSONStatusEnum = "Pending"
	CreateDirectIncome200ApplicationJSONStatusEnumFailed   CreateDirectIncome200ApplicationJSONStatusEnum = "Failed"
	CreateDirectIncome200ApplicationJSONStatusEnumSuccess  CreateDirectIncome200ApplicationJSONStatusEnum = "Success"
	CreateDirectIncome200ApplicationJSONStatusEnumTimedOut CreateDirectIncome200ApplicationJSONStatusEnum = "TimedOut"
)

type CreateDirectIncome200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateDirectIncome200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateDirectIncome200ApplicationJSONValidation struct {
	Errors   []CreateDirectIncome200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateDirectIncome200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type CreateDirectIncome200ApplicationJSON struct {
	Changes           []CreateDirectIncome200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                                  `json:"companyId"`
	CompletedOnUtc    *time.Time                                              `json:"completedOnUtc,omitempty"`
	Data              *CreateDirectIncome200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                                  `json:"dataConnectionKey"`
	DataType          *string                                                 `json:"dataType,omitempty"`
	ErrorMessage      *string                                                 `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                  `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                               `json:"requestedOnUtc"`
	Status            CreateDirectIncome200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                     `json:"statusCode"`
	TimeoutInMinutes  *int                                                    `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                    `json:"timeoutInSeconds,omitempty"`
	Validation        *CreateDirectIncome200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type CreateDirectIncomeResponse struct {
	ContentType                                string
	StatusCode                                 int
	RawResponse                                *http.Response
	CreateDirectIncome200ApplicationJSONObject *CreateDirectIncome200ApplicationJSON
}
