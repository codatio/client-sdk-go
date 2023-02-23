package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type GetDirectIncomePathParams struct {
	CompanyID      string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID   string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectIncomeID string `pathParam:"style=simple,explode=false,name=directIncomeId"`
}

type GetDirectIncomeSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetDirectIncomeRequest struct {
	PathParams GetDirectIncomePathParams
	Security   GetDirectIncomeSecurity
}

// GetDirectIncomeSourceModifiedDateContactRef
// A customer or supplier associated with the direct income.
type GetDirectIncomeSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// GetDirectIncomeSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type GetDirectIncomeSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetDirectIncomeSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the direct cost is linked.
type GetDirectIncomeSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// GetDirectIncomeSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type GetDirectIncomeSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type GetDirectIncomeSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type GetDirectIncomeSourceModifiedDateLineItems struct {
	AccountRef           *GetDirectIncomeSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                          `json:"description,omitempty"`
	DiscountAmount       *float64                                                         `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                         `json:"discountPercentage,omitempty"`
	ItemRef              *GetDirectIncomeSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                          `json:"quantity"`
	SubTotal             *float64                                                         `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                         `json:"taxAmount,omitempty"`
	TaxRateRef           *GetDirectIncomeSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                         `json:"totalAmount,omitempty"`
	TrackingCategoryRefs []GetDirectIncomeSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                          `json:"unitAmount"`
}

type GetDirectIncomeSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetDirectIncomeSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// GetDirectIncomeSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type GetDirectIncomeSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetDirectIncomeSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *GetDirectIncomeSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                               `json:"currency,omitempty"`
	CurrencyRate *float64                                                              `json:"currencyRate,omitempty"`
	ID           *string                                                               `json:"id,omitempty"`
	Note         *string                                                               `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                            `json:"paidOnDate,omitempty"`
	Reference    *string                                                               `json:"reference,omitempty"`
	TotalAmount  *float64                                                              `json:"totalAmount,omitempty"`
}

type GetDirectIncomeSourceModifiedDatePaymentAllocations struct {
	Allocation GetDirectIncomeSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    GetDirectIncomeSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type GetDirectIncomeSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// GetDirectIncomeSourceModifiedDate
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
type GetDirectIncomeSourceModifiedDate struct {
	ContactRef         *GetDirectIncomeSourceModifiedDateContactRef          `json:"contactRef,omitempty"`
	Currency           string                                                `json:"currency"`
	CurrencyRate       *float64                                              `json:"currencyRate,omitempty"`
	ID                 *string                                               `json:"id,omitempty"`
	IssueDate          time.Time                                             `json:"issueDate"`
	LineItems          []GetDirectIncomeSourceModifiedDateLineItems          `json:"lineItems"`
	Metadata           *GetDirectIncomeSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                            `json:"modifiedDate,omitempty"`
	Note               *string                                               `json:"note,omitempty"`
	PaymentAllocations []GetDirectIncomeSourceModifiedDatePaymentAllocations `json:"paymentAllocations"`
	Reference          *string                                               `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                            `json:"sourceModifiedDate,omitempty"`
	SubTotal           float64                                               `json:"subTotal"`
	SupplementalData   *GetDirectIncomeSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TaxAmount          float64                                               `json:"taxAmount"`
	TotalAmount        float64                                               `json:"totalAmount"`
}

type GetDirectIncomeResponse struct {
	ContentType        string
	SourceModifiedDate *GetDirectIncomeSourceModifiedDate
	StatusCode         int
}
