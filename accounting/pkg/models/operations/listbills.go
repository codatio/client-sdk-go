package operations

import (
	"net/http"
	"time"
)

type ListBillsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListBillsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListBillsRequest struct {
	PathParams  ListBillsPathParams
	QueryParams ListBillsQueryParams
}

type ListBillsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListBillsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListBillsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListBillsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListBillsLinksLinks struct {
	Current  ListBillsLinksLinksCurrent   `json:"current"`
	Next     *ListBillsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListBillsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListBillsLinksLinksSelf      `json:"self"`
}

// ListBillsLinksSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type ListBillsLinksSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ListBillsLinksSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the line item is linked.
type ListBillsLinksSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// ListBillsLinksSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type ListBillsLinksSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type ListBillsLinksSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type ListBillsLinksSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type ListBillsLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	ListBillsLinksSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       ListBillsLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	ListBillsLinksSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable ListBillsLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	ListBillsLinksSourceModifiedDateLineItemsTrackingIsBilledToEnumCustomer      ListBillsLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Customer"
	ListBillsLinksSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       ListBillsLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type ListBillsLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	ListBillsLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       ListBillsLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	ListBillsLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable ListBillsLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	ListBillsLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnumCustomer      ListBillsLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Customer"
	ListBillsLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       ListBillsLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type ListBillsLinksSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type ListBillsLinksSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []ListBillsLinksSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *ListBillsLinksSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   ListBillsLinksSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo ListBillsLinksSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *ListBillsLinksSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type ListBillsLinksSourceModifiedDateLineItems struct {
	AccountRef           *ListBillsLinksSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                         `json:"description,omitempty"`
	DiscountAmount       *float64                                                        `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                        `json:"discountPercentage,omitempty"`
	IsDirectCost         *bool                                                           `json:"isDirectCost,omitempty"`
	ItemRef              *ListBillsLinksSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                         `json:"quantity"`
	SubTotal             *float64                                                        `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                        `json:"taxAmount,omitempty"`
	TaxRateRef           *ListBillsLinksSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                        `json:"totalAmount,omitempty"`
	Tracking             *ListBillsLinksSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []ListBillsLinksSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                         `json:"unitAmount"`
}

type ListBillsLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type ListBillsLinksSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// ListBillsLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type ListBillsLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ListBillsLinksSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *ListBillsLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                              `json:"currency,omitempty"`
	CurrencyRate *float64                                                             `json:"currencyRate,omitempty"`
	ID           *string                                                              `json:"id,omitempty"`
	Note         *string                                                              `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                           `json:"paidOnDate,omitempty"`
	Reference    *string                                                              `json:"reference,omitempty"`
	TotalAmount  *float64                                                             `json:"totalAmount,omitempty"`
}

type ListBillsLinksSourceModifiedDatePaymentAllocations struct {
	Allocation ListBillsLinksSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    ListBillsLinksSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type ListBillsLinksSourceModifiedDatePurchaseOrderRefs struct {
	ID                  *string `json:"id,omitempty"`
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
}

type ListBillsLinksSourceModifiedDateStatusEnum string

const (
	ListBillsLinksSourceModifiedDateStatusEnumUnknown       ListBillsLinksSourceModifiedDateStatusEnum = "Unknown"
	ListBillsLinksSourceModifiedDateStatusEnumOpen          ListBillsLinksSourceModifiedDateStatusEnum = "Open"
	ListBillsLinksSourceModifiedDateStatusEnumPartiallyPaid ListBillsLinksSourceModifiedDateStatusEnum = "PartiallyPaid"
	ListBillsLinksSourceModifiedDateStatusEnumPaid          ListBillsLinksSourceModifiedDateStatusEnum = "Paid"
	ListBillsLinksSourceModifiedDateStatusEnumVoid          ListBillsLinksSourceModifiedDateStatusEnum = "Void"
	ListBillsLinksSourceModifiedDateStatusEnumDraft         ListBillsLinksSourceModifiedDateStatusEnum = "Draft"
)

type ListBillsLinksSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// ListBillsLinksSourceModifiedDateSupplierRef
// Reference to the supplier the bill was received from.
type ListBillsLinksSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type ListBillsLinksSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// ListBillsLinksSourceModifiedDate
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
type ListBillsLinksSourceModifiedDate struct {
	AmountDue          *float64                                             `json:"amountDue,omitempty"`
	Currency           *string                                              `json:"currency,omitempty"`
	CurrencyRate       *float64                                             `json:"currencyRate,omitempty"`
	DueDate            *time.Time                                           `json:"dueDate,omitempty"`
	ID                 *string                                              `json:"id,omitempty"`
	IssueDate          time.Time                                            `json:"issueDate"`
	LineItems          []ListBillsLinksSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata           *ListBillsLinksSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                           `json:"modifiedDate,omitempty"`
	Note               *string                                              `json:"note,omitempty"`
	PaymentAllocations []ListBillsLinksSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	PurchaseOrderRefs  []ListBillsLinksSourceModifiedDatePurchaseOrderRefs  `json:"purchaseOrderRefs,omitempty"`
	Reference          *string                                              `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                           `json:"sourceModifiedDate,omitempty"`
	Status             ListBillsLinksSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal           float64                                              `json:"subTotal"`
	SupplementalData   *ListBillsLinksSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef        *ListBillsLinksSourceModifiedDateSupplierRef         `json:"supplierRef,omitempty"`
	TaxAmount          float64                                              `json:"taxAmount"`
	TotalAmount        float64                                              `json:"totalAmount"`
	WithholdingTax     []ListBillsLinksSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

// ListBillsLinks
// Codat's Paging Model
type ListBillsLinks struct {
	Links        ListBillsLinksLinks                `json:"_links"`
	PageNumber   int64                              `json:"pageNumber"`
	PageSize     int64                              `json:"pageSize"`
	Results      []ListBillsLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                              `json:"totalResults"`
}

type ListBillsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListBillsLinks
}
