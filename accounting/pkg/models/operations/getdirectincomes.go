package operations

import (
	"net/http"
	"time"
)

type GetDirectIncomesPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetDirectIncomesQueryParams struct {
	OrderBy  *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page     int     `queryParam:"style=form,explode=true,name=page"`
	PageSize *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string `queryParam:"style=form,explode=true,name=query"`
}

type GetDirectIncomesRequest struct {
	PathParams  GetDirectIncomesPathParams
	QueryParams GetDirectIncomesQueryParams
}

type GetDirectIncomesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type GetDirectIncomesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type GetDirectIncomesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type GetDirectIncomesLinksLinksSelf struct {
	Href string `json:"href"`
}

type GetDirectIncomesLinksLinks struct {
	Current  GetDirectIncomesLinksLinksCurrent   `json:"current"`
	Next     *GetDirectIncomesLinksLinksNext     `json:"next,omitempty"`
	Previous *GetDirectIncomesLinksLinksPrevious `json:"previous,omitempty"`
	Self     GetDirectIncomesLinksLinksSelf      `json:"self"`
}

// GetDirectIncomesLinksSourceModifiedDateContactRef
// A customer or supplier associated with the direct income.
type GetDirectIncomesLinksSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// GetDirectIncomesLinksSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type GetDirectIncomesLinksSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetDirectIncomesLinksSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the direct cost is linked.
type GetDirectIncomesLinksSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// GetDirectIncomesLinksSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type GetDirectIncomesLinksSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// GetDirectIncomesLinksSourceModifiedDateLineItemsTrackingCategoryRefs
// References a category against which the item is tracked.
type GetDirectIncomesLinksSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type GetDirectIncomesLinksSourceModifiedDateLineItems struct {
	AccountRef           *GetDirectIncomesLinksSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                `json:"description,omitempty"`
	DiscountAmount       *float64                                                               `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                               `json:"discountPercentage,omitempty"`
	ItemRef              *GetDirectIncomesLinksSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                                `json:"quantity"`
	SubTotal             *float64                                                               `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                               `json:"taxAmount,omitempty"`
	TaxRateRef           *GetDirectIncomesLinksSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                               `json:"totalAmount,omitempty"`
	TrackingCategoryRefs []GetDirectIncomesLinksSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                                `json:"unitAmount"`
}

type GetDirectIncomesLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetDirectIncomesLinksSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// GetDirectIncomesLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type GetDirectIncomesLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetDirectIncomesLinksSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *GetDirectIncomesLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                     `json:"currency,omitempty"`
	CurrencyRate *float64                                                                    `json:"currencyRate,omitempty"`
	ID           *string                                                                     `json:"id,omitempty"`
	Note         *string                                                                     `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                  `json:"paidOnDate,omitempty"`
	Reference    *string                                                                     `json:"reference,omitempty"`
	TotalAmount  *float64                                                                    `json:"totalAmount,omitempty"`
}

type GetDirectIncomesLinksSourceModifiedDatePaymentAllocations struct {
	Allocation GetDirectIncomesLinksSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    GetDirectIncomesLinksSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

// GetDirectIncomesLinksSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type GetDirectIncomesLinksSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// GetDirectIncomesLinksSourceModifiedDate
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
type GetDirectIncomesLinksSourceModifiedDate struct {
	ContactRef         *GetDirectIncomesLinksSourceModifiedDateContactRef          `json:"contactRef,omitempty"`
	Currency           string                                                      `json:"currency"`
	CurrencyRate       *float64                                                    `json:"currencyRate,omitempty"`
	ID                 *string                                                     `json:"id,omitempty"`
	IssueDate          time.Time                                                   `json:"issueDate"`
	LineItems          []GetDirectIncomesLinksSourceModifiedDateLineItems          `json:"lineItems"`
	Metadata           *GetDirectIncomesLinksSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                  `json:"modifiedDate,omitempty"`
	Note               *string                                                     `json:"note,omitempty"`
	PaymentAllocations []GetDirectIncomesLinksSourceModifiedDatePaymentAllocations `json:"paymentAllocations"`
	Reference          *string                                                     `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                                  `json:"sourceModifiedDate,omitempty"`
	SubTotal           float64                                                     `json:"subTotal"`
	SupplementalData   *GetDirectIncomesLinksSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TaxAmount          float64                                                     `json:"taxAmount"`
	TotalAmount        float64                                                     `json:"totalAmount"`
}

// GetDirectIncomesLinks
// Codat's Paging Model
type GetDirectIncomesLinks struct {
	Links        GetDirectIncomesLinksLinks                `json:"_links"`
	PageNumber   int64                                     `json:"pageNumber"`
	PageSize     int64                                     `json:"pageSize"`
	Results      []GetDirectIncomesLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                     `json:"totalResults"`
}

type GetDirectIncomesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *GetDirectIncomesLinks
}
