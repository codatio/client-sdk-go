package shared

import (
	"time"
)

type DirectIncomeLineItems struct {
	AccountRef           *AccountRef `json:"accountRef,omitempty"`
	Description          *string     `json:"description,omitempty"`
	DiscountAmount       *float64    `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64    `json:"discountPercentage,omitempty"`
	ItemRef              *ItemRef    `json:"itemRef,omitempty"`
	Quantity             float64     `json:"quantity"`
	SubTotal             *float64    `json:"subTotal,omitempty"`
	TaxAmount            *float64    `json:"taxAmount,omitempty"`
	TaxRateRef           *TaxRateRef `json:"taxRateRef,omitempty"`
	TotalAmount          *float64    `json:"totalAmount,omitempty"`
	TrackingCategoryRefs []Items     `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           float64     `json:"unitAmount"`
}

// DirectIncome
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
// Direct incomes is a child data type of [account transactions](https://docs.codat.io/docs/datamodel-accounting-account-transactions).
type DirectIncome struct {
	ContactRef         *ContactRef               `json:"contactRef,omitempty"`
	Currency           string                    `json:"currency"`
	CurrencyRate       *float64                  `json:"currencyRate,omitempty"`
	ID                 *string                   `json:"id,omitempty"`
	IssueDate          time.Time                 `json:"issueDate"`
	LineItems          []DirectIncomeLineItems   `json:"lineItems"`
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
