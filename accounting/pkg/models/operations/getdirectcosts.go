package operations

import (
	"net/http"
	"time"
)

type GetDirectCostsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetDirectCostsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type GetDirectCostsRequest struct {
	PathParams  GetDirectCostsPathParams
	QueryParams GetDirectCostsQueryParams
}

type GetDirectCostsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type GetDirectCostsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type GetDirectCostsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type GetDirectCostsLinksLinksSelf struct {
	Href string `json:"href"`
}

type GetDirectCostsLinksLinks struct {
	Current  GetDirectCostsLinksLinksCurrent   `json:"current"`
	Next     *GetDirectCostsLinksLinksNext     `json:"next,omitempty"`
	Previous *GetDirectCostsLinksLinksPrevious `json:"previous,omitempty"`
	Self     GetDirectCostsLinksLinksSelf      `json:"self"`
}

// GetDirectCostsLinksSourceModifiedDateContactRef
// A customer or supplier associated with the direct cost.
type GetDirectCostsLinksSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// GetDirectCostsLinksSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type GetDirectCostsLinksSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetDirectCostsLinksSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the direct cost is linked.
type GetDirectCostsLinksSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// GetDirectCostsLinksSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the the line item is linked.
type GetDirectCostsLinksSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type GetDirectCostsLinksSourceModifiedDateLineItemsTracking struct {
	InvoiceTo  *string  `json:"invoiceTo,omitempty"`
	RecordRefs []string `json:"recordRefs"`
}

type GetDirectCostsLinksSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type GetDirectCostsLinksSourceModifiedDateLineItems struct {
	AccountRef           *GetDirectCostsLinksSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                              `json:"description,omitempty"`
	DiscountAmount       *float64                                                             `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                             `json:"discountPercentage,omitempty"`
	ItemRef              *GetDirectCostsLinksSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                              `json:"quantity"`
	SubTotal             *float64                                                             `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                             `json:"taxAmount,omitempty"`
	TaxRateRef           *GetDirectCostsLinksSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                             `json:"totalAmount,omitempty"`
	Tracking             *GetDirectCostsLinksSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []GetDirectCostsLinksSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                              `json:"unitAmount"`
}

type GetDirectCostsLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetDirectCostsLinksSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// GetDirectCostsLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type GetDirectCostsLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetDirectCostsLinksSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *GetDirectCostsLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                                   `json:"currency,omitempty"`
	CurrencyRate *float64                                                                  `json:"currencyRate,omitempty"`
	ID           *string                                                                   `json:"id,omitempty"`
	Note         *string                                                                   `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                                `json:"paidOnDate,omitempty"`
	Reference    *string                                                                   `json:"reference,omitempty"`
	TotalAmount  *float64                                                                  `json:"totalAmount,omitempty"`
}

type GetDirectCostsLinksSourceModifiedDatePaymentAllocations struct {
	Allocation GetDirectCostsLinksSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    GetDirectCostsLinksSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

// GetDirectCostsLinksSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type GetDirectCostsLinksSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// GetDirectCostsLinksSourceModifiedDate
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
type GetDirectCostsLinksSourceModifiedDate struct {
	ContactRef         *GetDirectCostsLinksSourceModifiedDateContactRef          `json:"contactRef,omitempty"`
	Currency           string                                                    `json:"currency"`
	CurrencyRate       *float64                                                  `json:"currencyRate,omitempty"`
	ID                 *string                                                   `json:"id,omitempty"`
	IssueDate          time.Time                                                 `json:"issueDate"`
	LineItems          []GetDirectCostsLinksSourceModifiedDateLineItems          `json:"lineItems"`
	Metadata           *GetDirectCostsLinksSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                `json:"modifiedDate,omitempty"`
	Note               *string                                                   `json:"note,omitempty"`
	PaymentAllocations []GetDirectCostsLinksSourceModifiedDatePaymentAllocations `json:"paymentAllocations"`
	Reference          *string                                                   `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                                `json:"sourceModifiedDate,omitempty"`
	SubTotal           float64                                                   `json:"subTotal"`
	SupplementalData   *GetDirectCostsLinksSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	TaxAmount          float64                                                   `json:"taxAmount"`
	TotalAmount        float64                                                   `json:"totalAmount"`
}

// GetDirectCostsLinks
// Codat's Paging Model
type GetDirectCostsLinks struct {
	Links        GetDirectCostsLinksLinks                `json:"_links"`
	PageNumber   int64                                   `json:"pageNumber"`
	PageSize     int64                                   `json:"pageSize"`
	Results      []GetDirectCostsLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                   `json:"totalResults"`
}

type GetDirectCostsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *GetDirectCostsLinks
}
