// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v4/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type AccountingBillPurchaseOrderReference struct {
	// Identifier for the purchase order, unique for the company in the accounting platform.
	ID *string `json:"id,omitempty"`
	// Friendly reference for the purchase order, commonly generated by the accounting platform.
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
}

func (o *AccountingBillPurchaseOrderReference) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingBillPurchaseOrderReference) GetPurchaseOrderNumber() *string {
	if o == nil {
		return nil
	}
	return o.PurchaseOrderNumber
}

type AccountingBillWithholdingTax struct {
	Amount *decimal.Big `decimal:"number" json:"amount"`
	Name   string       `json:"name"`
}

func (a AccountingBillWithholdingTax) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingBillWithholdingTax) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingBillWithholdingTax) GetAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.Amount
}

func (o *AccountingBillWithholdingTax) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// AccountingBill - > **Invoices or bills?**
// >
// > We distinguish between invoices where the company *owes money* vs. *is owed money*. If the company has received an invoice, and owes money to someone else (accounts payable) we call this a Bill.
// >
// > See [Invoices](https://docs.codat.io/accounting-api#/schemas/Invoice) for the accounts receivable equivalent of bills.
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
type AccountingBill struct {
	// Amount outstanding on the bill.
	AmountDue *decimal.Big `decimal:"number" json:"amountDue,omitempty"`
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
	DueDate      *string      `json:"dueDate,omitempty"`
	// Identifier for the bill, unique for the company in the accounting platform.
	ID        *string `json:"id,omitempty"`
	IssueDate string  `json:"issueDate"`
	// Array of Bill line items.
	LineItems    []BillLineItem `json:"lineItems,omitempty"`
	Metadata     *Metadata      `json:"metadata,omitempty"`
	ModifiedDate *string        `json:"modifiedDate,omitempty"`
	// Any private, company notes about the bill, such as payment information.
	Note *string `json:"note,omitempty"`
	// An array of payment allocations.
	PaymentAllocations []AccountingPaymentAllocation          `json:"paymentAllocations,omitempty"`
	PurchaseOrderRefs  []AccountingBillPurchaseOrderReference `json:"purchaseOrderRefs,omitempty"`
	// User-friendly reference for the bill.
	Reference          *string `json:"reference,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Current state of the bill.
	Status BillStatus `json:"status"`
	// Total amount of the bill, excluding any taxes.
	SubTotal *decimal.Big `decimal:"number" json:"subTotal"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// Reference to the supplier the record relates to.
	SupplierRef *SupplierRef `json:"supplierRef,omitempty"`
	// Amount of tax on the bill.
	TaxAmount *decimal.Big `decimal:"number" json:"taxAmount"`
	// Amount of the bill, including tax.
	TotalAmount    *decimal.Big                   `decimal:"number" json:"totalAmount"`
	WithholdingTax []AccountingBillWithholdingTax `json:"withholdingTax,omitempty"`
}

func (a AccountingBill) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingBill) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingBill) GetAmountDue() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.AmountDue
}

func (o *AccountingBill) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountingBill) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *AccountingBill) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *AccountingBill) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingBill) GetIssueDate() string {
	if o == nil {
		return ""
	}
	return o.IssueDate
}

func (o *AccountingBill) GetLineItems() []BillLineItem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *AccountingBill) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AccountingBill) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *AccountingBill) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *AccountingBill) GetPaymentAllocations() []AccountingPaymentAllocation {
	if o == nil {
		return nil
	}
	return o.PaymentAllocations
}

func (o *AccountingBill) GetPurchaseOrderRefs() []AccountingBillPurchaseOrderReference {
	if o == nil {
		return nil
	}
	return o.PurchaseOrderRefs
}

func (o *AccountingBill) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *AccountingBill) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *AccountingBill) GetStatus() BillStatus {
	if o == nil {
		return BillStatus("")
	}
	return o.Status
}

func (o *AccountingBill) GetSubTotal() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.SubTotal
}

func (o *AccountingBill) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *AccountingBill) GetSupplierRef() *SupplierRef {
	if o == nil {
		return nil
	}
	return o.SupplierRef
}

func (o *AccountingBill) GetTaxAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TaxAmount
}

func (o *AccountingBill) GetTotalAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TotalAmount
}

func (o *AccountingBill) GetWithholdingTax() []AccountingBillWithholdingTax {
	if o == nil {
		return nil
	}
	return o.WithholdingTax
}
