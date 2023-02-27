package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type GetDirectCostPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectCostID string `pathParam:"style=simple,explode=false,name=directCostId"`
}

type GetDirectCostSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetDirectCostRequest struct {
	PathParams GetDirectCostPathParams
	Security   GetDirectCostSecurity
}

// GetDirectCostSourceModifiedDateContactRef
// A customer or supplier associated with the direct cost.
type GetDirectCostSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// GetDirectCostSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type GetDirectCostSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetDirectCostSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the direct cost is linked.
type GetDirectCostSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// GetDirectCostSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the the line item is linked.
type GetDirectCostSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type GetDirectCostSourceModifiedDateLineItemsTracking struct {
	InvoiceTo  *string  `json:"invoiceTo,omitempty"`
	RecordRefs []string `json:"recordRefs"`
}

type GetDirectCostSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type GetDirectCostSourceModifiedDateLineItems struct {
	AccountRef           *GetDirectCostSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                        `json:"description,omitempty"`
	DiscountAmount       *float64                                                       `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                       `json:"discountPercentage,omitempty"`
	ItemRef              *GetDirectCostSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                        `json:"quantity"`
	SubTotal             *float64                                                       `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                       `json:"taxAmount,omitempty"`
	TaxRateRef           *GetDirectCostSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                       `json:"totalAmount,omitempty"`
	Tracking             *GetDirectCostSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []GetDirectCostSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                        `json:"unitAmount"`
}

type GetDirectCostSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetDirectCostSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// GetDirectCostSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type GetDirectCostSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetDirectCostSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *GetDirectCostSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                             `json:"currency,omitempty"`
	CurrencyRate *float64                                                            `json:"currencyRate,omitempty"`
	ID           *string                                                             `json:"id,omitempty"`
	Note         *string                                                             `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                          `json:"paidOnDate,omitempty"`
	Reference    *string                                                             `json:"reference,omitempty"`
	TotalAmount  *float64                                                            `json:"totalAmount,omitempty"`
}

type GetDirectCostSourceModifiedDatePaymentAllocations struct {
	Allocation GetDirectCostSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    GetDirectCostSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type GetDirectCostSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// GetDirectCostSourceModifiedDate
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
type GetDirectCostSourceModifiedDate struct {
	ContactRef         *GetDirectCostSourceModifiedDateContactRef          `json:"contactRef,omitempty"`
	Currency           string                                              `json:"currency"`
	CurrencyRate       *float64                                            `json:"currencyRate,omitempty"`
	ID                 *string                                             `json:"id,omitempty"`
	IssueDate          time.Time                                           `json:"issueDate"`
	LineItems          []GetDirectCostSourceModifiedDateLineItems          `json:"lineItems"`
	Metadata           *GetDirectCostSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                          `json:"modifiedDate,omitempty"`
	Note               *string                                             `json:"note,omitempty"`
	PaymentAllocations []GetDirectCostSourceModifiedDatePaymentAllocations `json:"paymentAllocations"`
	Reference          *string                                             `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                          `json:"sourceModifiedDate,omitempty"`
	SubTotal           float64                                             `json:"subTotal"`
	SupplementalData   *GetDirectCostSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TaxAmount          float64                                             `json:"taxAmount"`
	TotalAmount        float64                                             `json:"totalAmount"`
}

type GetDirectCostResponse struct {
	ContentType        string
	SourceModifiedDate *GetDirectCostSourceModifiedDate
	StatusCode         int
}
