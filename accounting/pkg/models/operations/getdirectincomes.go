// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"time"
)

type GetDirectIncomesRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Field to order results by. [Read more](https://docs.codat.io/using-the-api/ordering-results).
	OrderBy *string `queryParam:"style=form,explode=true,name=orderBy"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page int `queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `queryParam:"style=form,explode=true,name=pageSize"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
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

// GetDirectIncomesLinksSourceModifiedDateContactRef - A customer or supplier associated with the direct income.
type GetDirectIncomesLinksSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// GetDirectIncomesLinksSourceModifiedDateLineItemsAccountRef - Reference to the account to which the line item is linked.
type GetDirectIncomesLinksSourceModifiedDateLineItemsAccountRef struct {
	// 'id' from the Accounts data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the Accounts data type.
	Name *string `json:"name,omitempty"`
}

// GetDirectIncomesLinksSourceModifiedDateLineItemsItemRef - Reference to the product, service type, or inventory item to which the direct cost is linked.
type GetDirectIncomesLinksSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// GetDirectIncomesLinksSourceModifiedDateLineItemsTaxRateRef - Reference to the tax rate to which the line item is linked.
type GetDirectIncomesLinksSourceModifiedDateLineItemsTaxRateRef struct {
	// Applicable tax rate.
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	// 'id' from the 'taxRates' data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the 'taxRates' data type.
	Name *string `json:"name,omitempty"`
}

// GetDirectIncomesLinksSourceModifiedDateLineItemsTrackingCategoryRefs - References a category against which the item is tracked.
type GetDirectIncomesLinksSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type GetDirectIncomesLinksSourceModifiedDateLineItems struct {
	// Reference to the account to which the line item is linked.
	AccountRef *GetDirectIncomesLinksSourceModifiedDateLineItemsAccountRef `json:"accountRef,omitempty"`
	// A user-friendly name of the goods or services.
	Description *string `json:"description,omitempty"`
	// Discount amount for the line before tax.
	DiscountAmount *float64 `json:"discountAmount,omitempty"`
	// Discount percentage for the line before tax.
	DiscountPercentage *float64 `json:"discountPercentage,omitempty"`
	// Reference to the product, service type, or inventory item to which the direct cost is linked.
	ItemRef *GetDirectIncomesLinksSourceModifiedDateLineItemsItemRef `json:"itemRef,omitempty"`
	// The number of units of goods or services received.
	//
	// Note: If the platform does not provide this information, the quantity will be mapped as 1.
	Quantity float64 `json:"quantity"`
	// The amount of the line, inclusive of discounts, but exclusive of tax.
	SubTotal *float64 `json:"subTotal,omitempty"`
	// The amount of tax for the line.
	// Note: If the platform does not provide this information, the quantity will be mapped as 0.00.
	TaxAmount *float64 `json:"taxAmount,omitempty"`
	// Reference to the tax rate to which the line item is linked.
	TaxRateRef *GetDirectIncomesLinksSourceModifiedDateLineItemsTaxRateRef `json:"taxRateRef,omitempty"`
	// The total amount of the line, including tax.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	// An array of categories against which this direct cost is tracked.
	TrackingCategoryRefs []GetDirectIncomesLinksSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	// The price of each unit of goods or services.
	// Note: If the platform does not provide this information, the unit amount will be mapped to the total amount.
	UnitAmount float64 `json:"unitAmount"`
}

type GetDirectIncomesLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetDirectIncomesLinksSourceModifiedDatePaymentAllocationsAllocation struct {
	// The date the payment was allocated.
	AllocatedOnDate *time.Time `json:"allocatedOnDate,omitempty"`
	// The currency of the transaction.
	Currency *string `json:"currency,omitempty"`
	// Rate to convert the total amount of the payment into the base currency for the company at the time of the payment.
	//
	// Currency rates in Codat are implemented as the multiple of foreign currency units to each base currency unit.
	//
	// Where the currency rate is provided by the underlying accounting platform, it will be available from Codat with the same precision (up to a maximum of 9 decimal places).
	//
	// For accounting platforms which do not provide an explicit currency rate, it is calculated as `baseCurrency / foreignCurrency` and will be returned to 9 decimal places.
	//
	// ## Examples with base currency of GBP
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (GBP) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **USD**          | $20            | 0.781         | £15.62                     |
	// | **EUR**          | €20            | 0.885         | £17.70                     |
	// | **RUB**          | ₽20            | 0.011         | £0.22                      |
	//
	// ## Examples with base currency of USD
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (USD) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **GBP**          | £20            | 1.277         | $25.54                     |
	// | **EUR**          | €20            | 1.134         | $22.68                     |
	// | **RUB**          | ₽20            | 0.015         | $0.30                      |
	CurrencyRate *float64 `json:"currencyRate,omitempty"`
	// The total amount that has been allocated.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
}

// GetDirectIncomesLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef - The account that the allocated payment is made from or to.
type GetDirectIncomesLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef struct {
	// 'id' from the Accounts data type.
	ID *string `json:"id,omitempty"`
	// 'name' from the Accounts data type.
	Name *string `json:"name,omitempty"`
}

type GetDirectIncomesLinksSourceModifiedDatePaymentAllocationsPayment struct {
	// The account that the allocated payment is made from or to.
	AccountRef *GetDirectIncomesLinksSourceModifiedDatePaymentAllocationsPaymentAccountRef `json:"accountRef,omitempty"`
	// Currency the payment has been made in.
	Currency *string `json:"currency,omitempty"`
	// Rate to convert the total amount of the payment into the base currency for the company at the time of the payment.
	//
	// Currency rates in Codat are implemented as the multiple of foreign currency units to each base currency unit.
	//
	// Where the currency rate is provided by the underlying accounting platform, it will be available from Codat with the same precision (up to a maximum of 9 decimal places).
	//
	// For accounting platforms which do not provide an explicit currency rate, it is calculated as `baseCurrency / foreignCurrency` and will be returned to 9 decimal places.
	//
	// ## Examples with base currency of GBP
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (GBP) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **USD**          | $20            | 0.781         | £15.62                     |
	// | **EUR**          | €20            | 0.885         | £17.70                     |
	// | **RUB**          | ₽20            | 0.011         | £0.22                      |
	//
	// ## Examples with base currency of USD
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (USD) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **GBP**          | £20            | 1.277         | $25.54                     |
	// | **EUR**          | €20            | 1.134         | $22.68                     |
	// | **RUB**          | ₽20            | 0.015         | $0.30                      |
	CurrencyRate *float64 `json:"currencyRate,omitempty"`
	// Identifier of the allocated payment.
	ID *string `json:"id,omitempty"`
	// Notes attached to the allocated payment.
	Note *string `json:"note,omitempty"`
	// The date the payment was paid.
	PaidOnDate *time.Time `json:"paidOnDate,omitempty"`
	// Reference to the allocated payment.
	Reference *string `json:"reference,omitempty"`
	// Total amount that was paid.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
}

type GetDirectIncomesLinksSourceModifiedDatePaymentAllocations struct {
	Allocation GetDirectIncomesLinksSourceModifiedDatePaymentAllocationsAllocation `json:"allocation"`
	Payment    GetDirectIncomesLinksSourceModifiedDatePaymentAllocationsPayment    `json:"payment"`
}

// GetDirectIncomesLinksSourceModifiedDateSupplementalData - Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type GetDirectIncomesLinksSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// GetDirectIncomesLinksSourceModifiedDate - > **Language tip:**  Direct incomes may also be referred to as **Receive transactions**, **Receive money transactions**, **Sales receipts**, or **Cash sales** in various accounting platforms.
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
	// A customer or supplier associated with the direct income.
	ContactRef *GetDirectIncomesLinksSourceModifiedDateContactRef `json:"contactRef,omitempty"`
	// The currency of the direct income.
	Currency string `json:"currency"`
	// Rate to convert the total amount of the payment into the base currency for the company at the time of the payment.
	//
	// Currency rates in Codat are implemented as the multiple of foreign currency units to each base currency unit.
	//
	// Where the currency rate is provided by the underlying accounting platform, it will be available from Codat with the same precision (up to a maximum of 9 decimal places).
	//
	// For accounting platforms which do not provide an explicit currency rate, it is calculated as `baseCurrency / foreignCurrency` and will be returned to 9 decimal places.
	//
	// ## Examples with base currency of GBP
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (GBP) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **USD**          | $20            | 0.781         | £15.62                     |
	// | **EUR**          | €20            | 0.885         | £17.70                     |
	// | **RUB**          | ₽20            | 0.011         | £0.22                      |
	//
	// ## Examples with base currency of USD
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (USD) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **GBP**          | £20            | 1.277         | $25.54                     |
	// | **EUR**          | €20            | 1.134         | $22.68                     |
	// | **RUB**          | ₽20            | 0.015         | $0.30                      |
	CurrencyRate *float64 `json:"currencyRate,omitempty"`
	// Identifier of the direct income, unique for the company.
	ID *string `json:"id,omitempty"`
	// The date of the direct income as recorded in the accounting platform.
	IssueDate time.Time `json:"issueDate"`
	// An array of line items.
	LineItems []GetDirectIncomesLinksSourceModifiedDateLineItems `json:"lineItems"`
	Metadata  *GetDirectIncomesLinksSourceModifiedDateMetadata   `json:"metadata,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate       *time.Time                                                  `json:"modifiedDate,omitempty"`
	Note               *string                                                     `json:"note,omitempty"`
	PaymentAllocations []GetDirectIncomesLinksSourceModifiedDatePaymentAllocations `json:"paymentAllocations"`
	// User-friendly reference for the direct income.
	Reference *string `json:"reference,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
	// The total amount of the direct incomes, excluding any taxes.
	SubTotal float64 `json:"subTotal"`
	// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
	SupplementalData *GetDirectIncomesLinksSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	// The total amount of tax on the direct incomes.
	TaxAmount float64 `json:"taxAmount"`
	// The amount of the direct incomes, inclusive of tax.
	TotalAmount float64 `json:"totalAmount"`
}

// GetDirectIncomesLinks - Codat's Paging Model
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
	// Success
	Links *GetDirectIncomesLinks
}
