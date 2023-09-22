// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// AccountingDirectIncomeContactRef - A customer or supplier associated with the direct income.
type AccountingDirectIncomeContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

func (o *AccountingDirectIncomeContactRef) GetDataType() *string {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *AccountingDirectIncomeContactRef) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// AccountingDirectIncome - > **Language tip:**  Direct incomes may also be referred to as **Receive transactions**, **Receive money transactions**, **Sales receipts**, or **Cash sales** in various accounting platforms.
//
// > View the coverage for direct incomes in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Direct incomes are incomes received directly from the business' operations. For example, cash sales of items to a customer, referral commissions, and service fee refunds are considered direct incomes.
//
// Direct incomes include:
//
// - Selling an item directly to a contact, and receiving payment at the point of the sale.
// - Refunding an item in cash to a contact.
// - Depositing money into a bank account.
//
// Direct incomes is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type AccountingDirectIncome struct {
	// A customer or supplier associated with the direct income.
	ContactRef *AccountingDirectIncomeContactRef `json:"contactRef,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency string `json:"currency"`
	// Rate to convert the total amount of the payment into the base currency for the company at the time of the payment.
	//
	// Currency rates in Codat are implemented as the multiple of foreign currency units to each base currency unit.
	//
	// It is not possible to perform the currency conversion with two or more non-base currencies participating in the transaction. For example, if a company's base currency is USD, and it has a bill issued in EUR, then the bill payment must happen in USD or EUR.
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
	CurrencyRate *decimal.Big `decimal:"number" json:"currencyRate,omitempty"`
	// Identifier of the direct income, unique for the company.
	ID *string `json:"id,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	IssueDate string `json:"issueDate"`
	// An array of line items.
	LineItems          []DirectIncomeLineItem    `json:"lineItems"`
	Metadata           *Metadata                 `json:"metadata,omitempty"`
	ModifiedDate       *string                   `json:"modifiedDate,omitempty"`
	Note               *string                   `json:"note,omitempty"`
	PaymentAllocations []PaymentAllocationsitems `json:"paymentAllocations"`
	// User-friendly reference for the direct income.
	Reference          *string `json:"reference,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// The total amount of the direct incomes, excluding any taxes.
	SubTotal *decimal.Big `decimal:"number" json:"subTotal"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// The total amount of tax on the direct incomes.
	TaxAmount *decimal.Big `decimal:"number" json:"taxAmount"`
	// The amount of the direct incomes, inclusive of tax.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount"`
}

func (a AccountingDirectIncome) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingDirectIncome) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingDirectIncome) GetContactRef() *AccountingDirectIncomeContactRef {
	if o == nil {
		return nil
	}
	return o.ContactRef
}

func (o *AccountingDirectIncome) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *AccountingDirectIncome) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *AccountingDirectIncome) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingDirectIncome) GetIssueDate() string {
	if o == nil {
		return ""
	}
	return o.IssueDate
}

func (o *AccountingDirectIncome) GetLineItems() []DirectIncomeLineItem {
	if o == nil {
		return []DirectIncomeLineItem{}
	}
	return o.LineItems
}

func (o *AccountingDirectIncome) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AccountingDirectIncome) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *AccountingDirectIncome) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *AccountingDirectIncome) GetPaymentAllocations() []PaymentAllocationsitems {
	if o == nil {
		return []PaymentAllocationsitems{}
	}
	return o.PaymentAllocations
}

func (o *AccountingDirectIncome) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *AccountingDirectIncome) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *AccountingDirectIncome) GetSubTotal() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.SubTotal
}

func (o *AccountingDirectIncome) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *AccountingDirectIncome) GetTaxAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TaxAmount
}

func (o *AccountingDirectIncome) GetTotalAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TotalAmount
}
