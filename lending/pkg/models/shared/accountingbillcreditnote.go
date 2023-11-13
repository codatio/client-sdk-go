// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v5/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// AccountingBillCreditNote - > **Bill credit notes or credit notes?**
// >
// > In Codat, bill credit notes represent accounts payable only. For accounts receivable, see [Credit notes](https://docs.codat.io/lending-api#/schemas/CreditNote).
//
// View the coverage for bill credit notes in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billCreditNotes" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A bill credit note is issued by a supplier for the purpose of recording credit. For example, if a supplier was unable to fulfil an order that was placed by a business, or delivered damaged goods, they would issue a bill credit note. A bill credit note reduces the amount a business owes to the supplier. It can be refunded to the business or used to pay off future bills.
//
// In the Codat API, a bill credit note is an accounts payable record issued by a [supplier](https://docs.codat.io/lending-api#/schemas/Supplier).
//
// A bill credit note includes details of:
// * The original and remaining credit.
// * Any allocations of the credit against other records, such as [bills](https://docs.codat.io/lending-api#/schemas/Bill).
// * The supplier that issued the bill credit note.
type AccountingBillCreditNote struct {
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
	AllocatedOnDate *string `json:"allocatedOnDate,omitempty"`
	// Friendly reference for the bill credit note.
	BillCreditNoteNumber *string `json:"billCreditNoteNumber,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
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
	// Percentage rate of any discount applied to the bill credit note.
	DiscountPercentage *decimal.Big `decimal:"number" json:"discountPercentage"`
	// Identifier for the bill credit note that is unique to a company in the accounting platform.
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
	IssueDate *string `json:"issueDate,omitempty"`
	// An array of line
	LineItems    []BillCreditNoteLineItem `json:"lineItems,omitempty"`
	Metadata     *Metadata                `json:"metadata,omitempty"`
	ModifiedDate *string                  `json:"modifiedDate,omitempty"`
	// Any additional information about the bill credit note.
	Note *string `json:"note,omitempty"`
	// An array of payment allocations.
	PaymentAllocations []AccountingPaymentAllocation `json:"paymentAllocations,omitempty"`
	// Amount of the bill credit note that is still outstanding.
	RemainingCredit    *decimal.Big `decimal:"number" json:"remainingCredit,omitempty"`
	SourceModifiedDate *string      `json:"sourceModifiedDate,omitempty"`
	// Current state of the bill credit note
	Status BillCreditNoteStatus `json:"status"`
	// Total amount of the bill credit note, including discounts but excluding tax.
	SubTotal *decimal.Big `decimal:"number" json:"subTotal"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// Reference to the supplier the record relates to.
	SupplierRef *SupplierRef `json:"supplierRef,omitempty"`
	// Total amount of credit that has been applied to the business' account with the supplier, including discounts and tax.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount"`
	// Total value of any discounts applied.
	TotalDiscount *decimal.Big `decimal:"number" json:"totalDiscount"`
	// Amount of tax included in the bill credit note.
	TotalTaxAmount *decimal.Big `decimal:"number" json:"totalTaxAmount"`
	WithholdingTax []Items      `json:"withholdingTax,omitempty"`
}

func (a AccountingBillCreditNote) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingBillCreditNote) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingBillCreditNote) GetAllocatedOnDate() *string {
	if o == nil {
		return nil
	}
	return o.AllocatedOnDate
}

func (o *AccountingBillCreditNote) GetBillCreditNoteNumber() *string {
	if o == nil {
		return nil
	}
	return o.BillCreditNoteNumber
}

func (o *AccountingBillCreditNote) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountingBillCreditNote) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *AccountingBillCreditNote) GetDiscountPercentage() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.DiscountPercentage
}

func (o *AccountingBillCreditNote) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingBillCreditNote) GetIssueDate() *string {
	if o == nil {
		return nil
	}
	return o.IssueDate
}

func (o *AccountingBillCreditNote) GetLineItems() []BillCreditNoteLineItem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *AccountingBillCreditNote) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AccountingBillCreditNote) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *AccountingBillCreditNote) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *AccountingBillCreditNote) GetPaymentAllocations() []AccountingPaymentAllocation {
	if o == nil {
		return nil
	}
	return o.PaymentAllocations
}

func (o *AccountingBillCreditNote) GetRemainingCredit() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.RemainingCredit
}

func (o *AccountingBillCreditNote) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *AccountingBillCreditNote) GetStatus() BillCreditNoteStatus {
	if o == nil {
		return BillCreditNoteStatus("")
	}
	return o.Status
}

func (o *AccountingBillCreditNote) GetSubTotal() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.SubTotal
}

func (o *AccountingBillCreditNote) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *AccountingBillCreditNote) GetSupplierRef() *SupplierRef {
	if o == nil {
		return nil
	}
	return o.SupplierRef
}

func (o *AccountingBillCreditNote) GetTotalAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TotalAmount
}

func (o *AccountingBillCreditNote) GetTotalDiscount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TotalDiscount
}

func (o *AccountingBillCreditNote) GetTotalTaxAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TotalTaxAmount
}

func (o *AccountingBillCreditNote) GetWithholdingTax() []Items {
	if o == nil {
		return nil
	}
	return o.WithholdingTax
}
