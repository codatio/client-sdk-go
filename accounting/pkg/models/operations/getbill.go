package operations

import (
	"net/http"
	"time"
)

type GetBillPathParams struct {
	BillID    string `pathParam:"style=simple,explode=false,name=billId"`
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetBillRequest struct {
	PathParams GetBillPathParams
}

// GetBillSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type GetBillSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetBillSourceModifiedDateLineItemsItemRef
// Reference to the product, service type, or inventory item to which the line item is linked.
type GetBillSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// GetBillSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type GetBillSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type GetBillSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type GetBillSourceModifiedDateLineItemsTrackingCustomerRef struct {
	CompanyName *string `json:"companyName,omitempty"`
	ID          string  `json:"id"`
}

type GetBillSourceModifiedDateLineItemsTrackingIsBilledToEnum string

const (
	GetBillSourceModifiedDateLineItemsTrackingIsBilledToEnumUnknown       GetBillSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Unknown"
	GetBillSourceModifiedDateLineItemsTrackingIsBilledToEnumNotApplicable GetBillSourceModifiedDateLineItemsTrackingIsBilledToEnum = "NotApplicable"
	GetBillSourceModifiedDateLineItemsTrackingIsBilledToEnumCustomer      GetBillSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Customer"
	GetBillSourceModifiedDateLineItemsTrackingIsBilledToEnumProject       GetBillSourceModifiedDateLineItemsTrackingIsBilledToEnum = "Project"
)

type GetBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum string

const (
	GetBillSourceModifiedDateLineItemsTrackingIsRebilledToEnumUnknown       GetBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Unknown"
	GetBillSourceModifiedDateLineItemsTrackingIsRebilledToEnumNotApplicable GetBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "NotApplicable"
	GetBillSourceModifiedDateLineItemsTrackingIsRebilledToEnumCustomer      GetBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Customer"
	GetBillSourceModifiedDateLineItemsTrackingIsRebilledToEnumProject       GetBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum = "Project"
)

type GetBillSourceModifiedDateLineItemsTrackingProjectRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type GetBillSourceModifiedDateLineItemsTracking struct {
	CategoryRefs []GetBillSourceModifiedDateLineItemsTrackingCategoryRefs   `json:"categoryRefs"`
	CustomerRef  *GetBillSourceModifiedDateLineItemsTrackingCustomerRef     `json:"customerRef,omitempty"`
	IsBilledTo   GetBillSourceModifiedDateLineItemsTrackingIsBilledToEnum   `json:"isBilledTo"`
	IsRebilledTo GetBillSourceModifiedDateLineItemsTrackingIsRebilledToEnum `json:"isRebilledTo"`
	ProjectRef   *GetBillSourceModifiedDateLineItemsTrackingProjectRef      `json:"projectRef,omitempty"`
}

type GetBillSourceModifiedDateLineItems struct {
	AccountRef           *GetBillSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                  `json:"description,omitempty"`
	DiscountAmount       *float64                                                 `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                 `json:"discountPercentage,omitempty"`
	IsDirectCost         *bool                                                    `json:"isDirectCost,omitempty"`
	ItemRef              *GetBillSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             float64                                                  `json:"quantity"`
	SubTotal             *float64                                                 `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                 `json:"taxAmount,omitempty"`
	TaxRateRef           *GetBillSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                 `json:"totalAmount,omitempty"`
	Tracking             *GetBillSourceModifiedDateLineItemsTracking              `json:"tracking,omitempty"`
	TrackingCategoryRefs []GetBillSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                                                  `json:"unitAmount"`
}

type GetBillSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetBillSourceModifiedDatePaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

// GetBillSourceModifiedDatePaymentAllocationsPaymentAccountRef
// The account that the allocated payment is made from or to.
type GetBillSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetBillSourceModifiedDatePaymentAllocationsPayment struct {
	AccountRef   *GetBillSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	Currency     *string                                                       `json:"currency,omitempty"`
	CurrencyRate *float64                                                      `json:"currencyRate,omitempty"`
	ID           *string                                                       `json:"id,omitempty"`
	Note         *string                                                       `json:"note,omitempty"`
	PaidOnDate   *time.Time                                                    `json:"paidOnDate,omitempty"`
	Reference    *string                                                       `json:"reference,omitempty"`
	TotalAmount  *float64                                                      `json:"totalAmount,omitempty"`
}

type GetBillSourceModifiedDatePaymentAllocations struct {
	Allocation GetBillSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    GetBillSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

type GetBillSourceModifiedDatePurchaseOrderRefs struct {
	ID                  *string `json:"id,omitempty"`
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
}

type GetBillSourceModifiedDateStatusEnum string

const (
	GetBillSourceModifiedDateStatusEnumUnknown       GetBillSourceModifiedDateStatusEnum = "Unknown"
	GetBillSourceModifiedDateStatusEnumOpen          GetBillSourceModifiedDateStatusEnum = "Open"
	GetBillSourceModifiedDateStatusEnumPartiallyPaid GetBillSourceModifiedDateStatusEnum = "PartiallyPaid"
	GetBillSourceModifiedDateStatusEnumPaid          GetBillSourceModifiedDateStatusEnum = "Paid"
	GetBillSourceModifiedDateStatusEnumVoid          GetBillSourceModifiedDateStatusEnum = "Void"
	GetBillSourceModifiedDateStatusEnumDraft         GetBillSourceModifiedDateStatusEnum = "Draft"
)

type GetBillSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// GetBillSourceModifiedDateSupplierRef
// Reference to the supplier the bill was received from.
type GetBillSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type GetBillSourceModifiedDateWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// GetBillSourceModifiedDate
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
type GetBillSourceModifiedDate struct {
	AmountDue          *float64                                      `json:"amountDue,omitempty"`
	Currency           *string                                       `json:"currency,omitempty"`
	CurrencyRate       *float64                                      `json:"currencyRate,omitempty"`
	DueDate            *time.Time                                    `json:"dueDate,omitempty"`
	ID                 *string                                       `json:"id,omitempty"`
	IssueDate          time.Time                                     `json:"issueDate"`
	LineItems          []GetBillSourceModifiedDateLineItems          `json:"lineItems,omitempty"`
	Metadata           *GetBillSourceModifiedDateMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                    `json:"modifiedDate,omitempty"`
	Note               *string                                       `json:"note,omitempty"`
	PaymentAllocations []GetBillSourceModifiedDatePaymentAllocations `json:"paymentAllocations,omitempty"`
	PurchaseOrderRefs  []GetBillSourceModifiedDatePurchaseOrderRefs  `json:"purchaseOrderRefs,omitempty"`
	Reference          *string                                       `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                    `json:"sourceModifiedDate,omitempty"`
	Status             GetBillSourceModifiedDateStatusEnum           `json:"status"`
	SubTotal           float64                                       `json:"subTotal"`
	SupplementalData   *GetBillSourceModifiedDateSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef        *GetBillSourceModifiedDateSupplierRef         `json:"supplierRef,omitempty"`
	TaxAmount          float64                                       `json:"taxAmount"`
	TotalAmount        float64                                       `json:"totalAmount"`
	WithholdingTax     []GetBillSourceModifiedDateWithholdingTax     `json:"withholdingTax,omitempty"`
}

type GetBillResponse struct {
	ContentType        string
	SourceModifiedDate *GetBillSourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
