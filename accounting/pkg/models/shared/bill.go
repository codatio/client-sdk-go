package shared

import (
	"time"
)

// BillLineItemsItemRef
// Reference to the product, service type, or inventory item to which the line item is linked.
type BillLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type BillLineItemsTrackingIsBilledToEnum string

const (
	BillLineItemsTrackingIsBilledToEnumUnknown       BillLineItemsTrackingIsBilledToEnum = "Unknown"
	BillLineItemsTrackingIsBilledToEnumNotApplicable BillLineItemsTrackingIsBilledToEnum = "NotApplicable"
	BillLineItemsTrackingIsBilledToEnumCustomer      BillLineItemsTrackingIsBilledToEnum = "Customer"
	BillLineItemsTrackingIsBilledToEnumProject       BillLineItemsTrackingIsBilledToEnum = "Project"
)

type BillLineItemsTracking struct {
	CategoryRefs []Items                             `json:"categoryRefs"`
	CustomerRef  *CustomerRef                        `json:"customerRef,omitempty"`
	IsBilledTo   BillLineItemsTrackingIsBilledToEnum `json:"isBilledTo"`
	IsRebilledTo IsBilledToEnum                      `json:"isRebilledTo"`
	ProjectRef   *ProjectRef                         `json:"projectRef,omitempty"`
}

type BillLineItems struct {
	AccountRef           *AccountRef            `json:"accountRef,omitempty"`
	Description          *string                `json:"description,omitempty"`
	DiscountAmount       *float64               `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64               `json:"discountPercentage,omitempty"`
	IsDirectCost         *bool                  `json:"isDirectCost,omitempty"`
	ItemRef              *BillLineItemsItemRef  `json:"itemRef,omitempty"`
	Quantity             float64                `json:"quantity"`
	SubTotal             *float64               `json:"subTotal,omitempty"`
	TaxAmount            *float64               `json:"taxAmount,omitempty"`
	TaxRateRef           *TaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64               `json:"totalAmount,omitempty"`
	Tracking             *BillLineItemsTracking `json:"tracking,omitempty"`
	TrackingCategoryRefs []Items                `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                `json:"unitAmount"`
}

type BillMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type BillPaymentAllocationsAllocation struct {
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrencyRate    *float64   `json:"currencyRate,omitempty"`
	TotalAmount     *float64   `json:"totalAmount,omitempty"`
}

type BillPaymentAllocationsPayment struct {
	AccountRef   *AccountRef `json:"accountRef,omitempty"`
	Currency     *string     `json:"currency,omitempty"`
	CurrencyRate *float64    `json:"currencyRate,omitempty"`
	ID           *string     `json:"id,omitempty"`
	Note         *string     `json:"note,omitempty"`
	PaidOnDate   *time.Time  `json:"paidOnDate,omitempty"`
	Reference    *string     `json:"reference,omitempty"`
	TotalAmount  *float64    `json:"totalAmount,omitempty"`
}

type BillPaymentAllocations struct {
	Allocation BillPaymentAllocationsAllocation `json:"allocation"`
	Payment    BillPaymentAllocationsPayment    `json:"payment"`
}

type BillPurchaseOrderRefs struct {
	ID                  *string `json:"id,omitempty"`
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
}

type BillStatusEnum string

const (
	BillStatusEnumUnknown       BillStatusEnum = "Unknown"
	BillStatusEnumOpen          BillStatusEnum = "Open"
	BillStatusEnumPartiallyPaid BillStatusEnum = "PartiallyPaid"
	BillStatusEnumPaid          BillStatusEnum = "Paid"
	BillStatusEnumVoid          BillStatusEnum = "Void"
	BillStatusEnumDraft         BillStatusEnum = "Draft"
)

type BillSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// BillSupplierRef
// Reference to the supplier the bill was received from.
type BillSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

type BillWithholdingTax struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}

// Bill
// > **Invoices or bills?**
// >
// > In Codat, bills are for accounts payable only. For the accounts receivable equivalent of bills, see [Invoices](https://docs.codat.io/docs/datamodel-accounting-invoices).
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
// You can find these types of transactions in our [Direct costs](https://docs.codat.io/docs/datamodel-accounting-directcosts) data model.
type Bill struct {
	AmountDue          *float64                 `json:"amountDue,omitempty"`
	Currency           *string                  `json:"currency,omitempty"`
	CurrencyRate       *float64                 `json:"currencyRate,omitempty"`
	DueDate            *time.Time               `json:"dueDate,omitempty"`
	ID                 *string                  `json:"id,omitempty"`
	IssueDate          time.Time                `json:"issueDate"`
	LineItems          []BillLineItems          `json:"lineItems,omitempty"`
	Metadata           *BillMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time               `json:"modifiedDate,omitempty"`
	Note               *string                  `json:"note,omitempty"`
	PaymentAllocations []BillPaymentAllocations `json:"paymentAllocations,omitempty"`
	PurchaseOrderRefs  []BillPurchaseOrderRefs  `json:"purchaseOrderRefs,omitempty"`
	Reference          *string                  `json:"reference,omitempty"`
	SourceModifiedDate *time.Time               `json:"sourceModifiedDate,omitempty"`
	Status             BillStatusEnum           `json:"status"`
	SubTotal           float64                  `json:"subTotal"`
	SupplementalData   *BillSupplementalData    `json:"supplementalData,omitempty"`
	SupplierRef        *BillSupplierRef         `json:"supplierRef,omitempty"`
	TaxAmount          float64                  `json:"taxAmount"`
	TotalAmount        float64                  `json:"totalAmount"`
	WithholdingTax     []BillWithholdingTax     `json:"withholdingTax,omitempty"`
}
