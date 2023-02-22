package shared

import (
	"time"
)

type DirectCostLineItemsTracking struct {
	InvoiceTo  *string  `json:"invoiceTo,omitempty"`
	RecordRefs []string `json:"recordRefs"`
}

type DirectCostLineItems struct {
	AccountRef           *AccountRef                  `json:"accountRef,omitempty"`
	Description          *string                      `json:"description,omitempty"`
	DiscountAmount       *float64                     `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                     `json:"discountPercentage,omitempty"`
	ItemRef              *ItemRef                     `json:"itemRef,omitempty"`
	Quantity             float64                      `json:"quantity"`
	SubTotal             *float64                     `json:"subTotal,omitempty"`
	TaxAmount            *float64                     `json:"taxAmount,omitempty"`
	TaxRateRef           *TaxRateRef                  `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                     `json:"totalAmount,omitempty"`
	Tracking             *DirectCostLineItemsTracking `json:"tracking,omitempty"`
	TrackingCategoryRefs []Items                      `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64                      `json:"unitAmount"`
}

// DirectCost
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
// Direct costs is a child data type of [account transactions](https://docs.codat.io/docs/datamodel-accounting-account-transactions).
type DirectCost struct {
	ContactRef         *ContactRef               `json:"contactRef,omitempty"`
	Currency           string                    `json:"currency"`
	CurrencyRate       *float64                  `json:"currencyRate,omitempty"`
	ID                 *string                   `json:"id,omitempty"`
	IssueDate          time.Time                 `json:"issueDate"`
	LineItems          []DirectCostLineItems     `json:"lineItems"`
	Metadata           *Metadata                 `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                `json:"modifiedDate,omitempty"`
	Note               *string                   `json:"note,omitempty"`
	PaymentAllocations []PaymentAllocationsitems `json:"paymentAllocations"`
	Reference          *string                   `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                `json:"sourceModifiedDate,omitempty"`
	SubTotal           float64                   `json:"subTotal"`
	SupplementalData   *SupplementalData         `json:"supplementalData,omitempty"`
	TaxAmount          float64                   `json:"taxAmount"`
	TotalAmount        float64                   `json:"totalAmount"`
}
