// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v5/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type SalesOrderReference struct {
	// Available Data types
	DataType *DataType `json:"dataType,omitempty"`
	// Unique identifier to a record in `dataType`.
	ID *string `json:"id,omitempty"`
}

func (o *SalesOrderReference) GetDataType() *DataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *SalesOrderReference) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AccountingInvoice - > **Invoices or bills?**
// >
// > We distinguish between invoices where the company *owes money* vs. *is owed money*. If the company issued an invoice, and is owed money (accounts receivable) we call this an Invoice.
// >
// > See [Bills](https://docs.codat.io/lending-api#/schemas/Bill) for the accounts payable equivalent of bills.
//
// View the coverage for invoices in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// An invoice is an itemized record of goods sold or services provided to a [customer](https://docs.codat.io/lending-api#/schemas/Customer).
//
// In Codat, an invoice contains details of:
//
// - The timeline of the invoice—when it was raised, marked as paid, last edited, and so on.
// - How much the invoice is for, what portion of the invoice is tax or discounts, and what currency the amounts are represented in.
// - Who the invoice has been raised to; the _customer_.
// - The breakdown of what the invoice is for; the _line items_.
// - Any [payments](https://docs.codat.io/lending-api#/schemas/Payment) assigned to the invoice; the _payment allocations_.
//
// > **Invoice PDF downloads**
// >
// > You can <a className="external" href="https://docs.codat.io/lending-api#/operations/get-invoice-pdf" target="_blank">download a PDF version</a> of an invoice for supported integrations.
// >
// > The filename will be invoice-{number}.pdf.
//
// > **Referencing an invoice in Sage 50 and ClearBooks**
// >
// > In Sage 50 and ClearBooks, you may prefer to use the **invoiceNumber** to identify an invoice rather than the invoice **id**. Each time a draft invoice is submitted or printed, the draft **id** becomes void and a submitted invoice with a new **id** exists in its place. In both platforms, the **invoiceNumber** should remain the same.
type AccountingInvoice struct {
	// Additional tax amount applied to invoice.
	AdditionalTaxAmount *decimal.Big `decimal:"number" json:"additionalTaxAmount,omitempty"`
	// Percentage rate of any additional tax applied to the invoice.
	AdditionalTaxPercentage *decimal.Big `decimal:"number" json:"additionalTaxPercentage,omitempty"`
	// Amount outstanding on the invoice.
	AmountDue *decimal.Big `decimal:"number" json:"amountDue"`
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
	//
	//
	// ### Integration-specific details
	//
	// | Integration       | Scenario                                        | System behavior                                                                                                                                                      |
	// |-------------------|-------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	// | QuickBooks Online | Transaction currency differs from base currency | If currency rate value is left `null`, a rate of 1 will be used by QBO by default. To override this, include the required currency rate in the expense transaction.  |
	CurrencyRate *decimal.Big           `decimal:"number" json:"currencyRate,omitempty"`
	CustomerRef  *AccountingCustomerRef `json:"customerRef,omitempty"`
	// Percentage rate (from 0 to 100) of discounts applied to the invoice. For example: A 5% discount will return a value of `5`, not `0.05`.
	DiscountPercentage *decimal.Big `decimal:"number" json:"discountPercentage,omitempty"`
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
	DueDate *string `json:"dueDate,omitempty"`
	// Identifier for the invoice, unique to the company in the accounting platform.
	ID *string `json:"id,omitempty"`
	// Friendly reference for the invoice. If available, this appears in the file name of invoice attachments.
	InvoiceNumber *string `json:"invoiceNumber,omitempty"`
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
	LineItems    []InvoiceLineItem `json:"lineItems,omitempty"`
	Metadata     *Metadata         `json:"metadata,omitempty"`
	ModifiedDate *string           `json:"modifiedDate,omitempty"`
	// Any additional information about the invoice. Where possible, Codat links to a data field in the accounting platform that is publicly available. This means that the contents of the note field are included when an invoice is emailed from the accounting platform to the customer.
	Note *string `json:"note,omitempty"`
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
	PaidOnDate *string `json:"paidOnDate,omitempty"`
	// An array of payment allocations.
	PaymentAllocations []AccountingPaymentAllocation `json:"paymentAllocations,omitempty"`
	// List of references to related Sales orders.
	SalesOrderRefs     []SalesOrderReference `json:"salesOrderRefs,omitempty"`
	SourceModifiedDate *string               `json:"sourceModifiedDate,omitempty"`
	// Current state of the invoice:
	//
	// - `Draft` - Invoice hasn't been submitted to the supplier. It may be in a pending state or is scheduled for future submission, for example by email.
	// - `Submitted` - Invoice is no longer a draft. It has been processed and, or, sent to the customer. In this state, it will impact the ledger. It also has no payments made against it (amountDue == totalAmount).
	// - `PartiallyPaid` - The balance paid against the invoice is positive, but less than the total invoice amount (0 < amountDue < totalAmount).
	// - `Paid` - Invoice is paid in full. This includes if the invoice has been credited or overpaid (amountDue == 0).
	// - `Void` - An invoice can become Void when it's deleted, refunded, written off, or cancelled. A voided invoice may still be PartiallyPaid, and so all outstanding amounts on voided invoices are removed from the accounts receivable account.
	Status InvoiceStatus `json:"status"`
	// Total amount of the invoice excluding any taxes.
	SubTotal *decimal.Big `decimal:"number" json:"subTotal,omitempty"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// Amount of the invoice, inclusive of tax.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount"`
	// Numerical value of discounts applied to the invoice.
	TotalDiscount *decimal.Big `decimal:"number" json:"totalDiscount,omitempty"`
	// Amount of tax on the invoice.
	TotalTaxAmount *decimal.Big `decimal:"number" json:"totalTaxAmount"`
	WithholdingTax []Items      `json:"withholdingTax,omitempty"`
}

func (a AccountingInvoice) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingInvoice) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingInvoice) GetAdditionalTaxAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.AdditionalTaxAmount
}

func (o *AccountingInvoice) GetAdditionalTaxPercentage() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.AdditionalTaxPercentage
}

func (o *AccountingInvoice) GetAmountDue() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.AmountDue
}

func (o *AccountingInvoice) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountingInvoice) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *AccountingInvoice) GetCustomerRef() *AccountingCustomerRef {
	if o == nil {
		return nil
	}
	return o.CustomerRef
}

func (o *AccountingInvoice) GetDiscountPercentage() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.DiscountPercentage
}

func (o *AccountingInvoice) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *AccountingInvoice) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingInvoice) GetInvoiceNumber() *string {
	if o == nil {
		return nil
	}
	return o.InvoiceNumber
}

func (o *AccountingInvoice) GetIssueDate() string {
	if o == nil {
		return ""
	}
	return o.IssueDate
}

func (o *AccountingInvoice) GetLineItems() []InvoiceLineItem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *AccountingInvoice) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AccountingInvoice) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *AccountingInvoice) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *AccountingInvoice) GetPaidOnDate() *string {
	if o == nil {
		return nil
	}
	return o.PaidOnDate
}

func (o *AccountingInvoice) GetPaymentAllocations() []AccountingPaymentAllocation {
	if o == nil {
		return nil
	}
	return o.PaymentAllocations
}

func (o *AccountingInvoice) GetSalesOrderRefs() []SalesOrderReference {
	if o == nil {
		return nil
	}
	return o.SalesOrderRefs
}

func (o *AccountingInvoice) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *AccountingInvoice) GetStatus() InvoiceStatus {
	if o == nil {
		return InvoiceStatus("")
	}
	return o.Status
}

func (o *AccountingInvoice) GetSubTotal() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *AccountingInvoice) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *AccountingInvoice) GetTotalAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TotalAmount
}

func (o *AccountingInvoice) GetTotalDiscount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalDiscount
}

func (o *AccountingInvoice) GetTotalTaxAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TotalTaxAmount
}

func (o *AccountingInvoice) GetWithholdingTax() []Items {
	if o == nil {
		return nil
	}
	return o.WithholdingTax
}
