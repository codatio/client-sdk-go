// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type AccountingCreateInvoiceResponseAllocation struct {
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
	// Where the currency rate is provided by the underlying accounting software, it will be available from Codat with the same precision (up to a maximum of 9 decimal places).
	//
	// For accounting software which do not provide an explicit currency rate, it is calculated as `baseCurrency / foreignCurrency` and will be returned to 9 decimal places.
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
	// | QuickBooks Online | Transaction currency differs from base currency | If currency rate value is left `null`, a rate of 1 will be used by QBO by default. To override this, specify a currencyRate in the request body.  |
	CurrencyRate *decimal.Big `decimal:"number" json:"currencyRate,omitempty"`
	// The total amount that has been allocated.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount,omitempty"`
}

func (a AccountingCreateInvoiceResponseAllocation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingCreateInvoiceResponseAllocation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingCreateInvoiceResponseAllocation) GetAllocatedOnDate() *string {
	if o == nil {
		return nil
	}
	return o.AllocatedOnDate
}

func (o *AccountingCreateInvoiceResponseAllocation) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountingCreateInvoiceResponseAllocation) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *AccountingCreateInvoiceResponseAllocation) GetTotalAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

type AccountingCreateInvoiceResponseAccountingPaymentAllocation struct {
	Allocation AccountingCreateInvoiceResponseAllocation `json:"allocation"`
	Payment    PaymentAllocationPayment                  `json:"payment"`
}

func (o *AccountingCreateInvoiceResponseAccountingPaymentAllocation) GetAllocation() AccountingCreateInvoiceResponseAllocation {
	if o == nil {
		return AccountingCreateInvoiceResponseAllocation{}
	}
	return o.Allocation
}

func (o *AccountingCreateInvoiceResponseAccountingPaymentAllocation) GetPayment() PaymentAllocationPayment {
	if o == nil {
		return PaymentAllocationPayment{}
	}
	return o.Payment
}

// AccountingCreateInvoiceResponseDataType - The underlying data type associated to the reference `id`.
type AccountingCreateInvoiceResponseDataType string

const (
	AccountingCreateInvoiceResponseDataTypeSalesOrders AccountingCreateInvoiceResponseDataType = "salesOrders"
)

func (e AccountingCreateInvoiceResponseDataType) ToPointer() *AccountingCreateInvoiceResponseDataType {
	return &e
}
func (e *AccountingCreateInvoiceResponseDataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "salesOrders":
		*e = AccountingCreateInvoiceResponseDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountingCreateInvoiceResponseDataType: %v", v)
	}
}

type AccountingCreateInvoiceResponseSalesOrderReference struct {
	// The underlying data type associated to the reference `id`.
	DataType *AccountingCreateInvoiceResponseDataType `json:"dataType,omitempty"`
	// Unique identifier to a record in `dataType`.
	ID *string `json:"id,omitempty"`
}

func (o *AccountingCreateInvoiceResponseSalesOrderReference) GetDataType() *AccountingCreateInvoiceResponseDataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *AccountingCreateInvoiceResponseSalesOrderReference) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type AccountingCreateInvoiceResponseWithholdingTax struct {
	// Amount of tax withheld.
	Amount *decimal.Big `decimal:"number" json:"amount"`
	// Name assigned to withheld tax.
	Name string `json:"name"`
}

func (a AccountingCreateInvoiceResponseWithholdingTax) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingCreateInvoiceResponseWithholdingTax) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingCreateInvoiceResponseWithholdingTax) GetAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.Amount
}

func (o *AccountingCreateInvoiceResponseWithholdingTax) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// AccountingCreateInvoiceResponseAccountingInvoice - > **Invoices or bills?**
// >
// > We distinguish between invoices where the company *owes money* vs. *is owed money*. If the company issued an invoice, and is owed money (accounts receivable) we call this an Invoice.
// >
// > See [Bills](https://docs.codat.io/accounting-api#/schemas/Bill) for the accounts payable equivalent of bills.
//
// View the coverage for invoices in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=invoices" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// An invoice is an itemized record of goods sold or services provided to a [customer](https://docs.codat.io/accounting-api#/schemas/Customer).
//
// In Codat, an invoice contains details of:
//
// - The timeline of the invoice—when it was raised, marked as paid, last edited, and so on.
// - How much the invoice is for, what portion of the invoice is tax or discounts, and what currency the amounts are represented in.
// - Who the invoice has been raised to; the _customer_.
// - The breakdown of what the invoice is for; the _line items_.
// - Any [payments](https://docs.codat.io/accounting-api#/schemas/Payment) assigned to the invoice; the _payment allocations_.
//
// > **Invoice PDF downloads**
// >
// > You can <a className="external" href="https://docs.codat.io/accounting-api#/operations/get-invoice-pdf" target="_blank">download a PDF version</a> of an invoice for supported integrations.
// >
// > The filename will be invoice-{number}.pdf.
//
// > **Referencing an invoice in Sage 50 and ClearBooks**
// >
// > In Sage 50 and ClearBooks, you may prefer to use the **invoiceNumber** to identify an invoice rather than the invoice **id**. Each time a draft invoice is submitted or printed, the draft **id** becomes void and a submitted invoice with a new **id** exists in its place. In both platforms, the **invoiceNumber** should remain the same.
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type AccountingCreateInvoiceResponseAccountingInvoice struct {
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
	// Where the currency rate is provided by the underlying accounting software, it will be available from Codat with the same precision (up to a maximum of 9 decimal places).
	//
	// For accounting software which do not provide an explicit currency rate, it is calculated as `baseCurrency / foreignCurrency` and will be returned to 9 decimal places.
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
	// | QuickBooks Online | Transaction currency differs from base currency | If currency rate value is left `null`, a rate of 1 will be used by QBO by default. To override this, specify a currencyRate in the request body.  |
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
	// Identifier for the invoice, unique to the company in the accounting software.
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
	// Any additional information about the invoice. Where possible, Codat links to a data field in the accounting software that is publicly available. This means that the contents of the note field are included when an invoice is emailed from the accounting software to the customer.
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
	PaymentAllocations []AccountingCreateInvoiceResponseAccountingPaymentAllocation `json:"paymentAllocations,omitempty"`
	// List of references to related Sales orders.
	SalesOrderRefs     []AccountingCreateInvoiceResponseSalesOrderReference `json:"salesOrderRefs,omitempty"`
	SourceModifiedDate *string                                              `json:"sourceModifiedDate,omitempty"`
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
	// It is referenced as a configured dynamic key value pair that is unique to the accounting software. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// Amount of the invoice, inclusive of tax.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount"`
	// Numerical value of discounts applied to the invoice.
	TotalDiscount *decimal.Big `decimal:"number" json:"totalDiscount,omitempty"`
	// Amount of tax on the invoice.
	TotalTaxAmount *decimal.Big                                    `decimal:"number" json:"totalTaxAmount"`
	WithholdingTax []AccountingCreateInvoiceResponseWithholdingTax `json:"withholdingTax,omitempty"`
}

func (a AccountingCreateInvoiceResponseAccountingInvoice) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingCreateInvoiceResponseAccountingInvoice) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetAdditionalTaxAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.AdditionalTaxAmount
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetAdditionalTaxPercentage() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.AdditionalTaxPercentage
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetAmountDue() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.AmountDue
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetCustomerRef() *AccountingCustomerRef {
	if o == nil {
		return nil
	}
	return o.CustomerRef
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetDiscountPercentage() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.DiscountPercentage
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetInvoiceNumber() *string {
	if o == nil {
		return nil
	}
	return o.InvoiceNumber
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetIssueDate() string {
	if o == nil {
		return ""
	}
	return o.IssueDate
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetLineItems() []InvoiceLineItem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetPaidOnDate() *string {
	if o == nil {
		return nil
	}
	return o.PaidOnDate
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetPaymentAllocations() []AccountingCreateInvoiceResponseAccountingPaymentAllocation {
	if o == nil {
		return nil
	}
	return o.PaymentAllocations
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetSalesOrderRefs() []AccountingCreateInvoiceResponseSalesOrderReference {
	if o == nil {
		return nil
	}
	return o.SalesOrderRefs
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetStatus() InvoiceStatus {
	if o == nil {
		return InvoiceStatus("")
	}
	return o.Status
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetSubTotal() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetTotalAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TotalAmount
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetTotalDiscount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalDiscount
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetTotalTaxAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TotalTaxAmount
}

func (o *AccountingCreateInvoiceResponseAccountingInvoice) GetWithholdingTax() []AccountingCreateInvoiceResponseWithholdingTax {
	if o == nil {
		return nil
	}
	return o.WithholdingTax
}

type AccountingCreateInvoiceResponse struct {
	// Contains a single entry that communicates which record has changed and the manner in which it changed.
	Changes []PushOperationChange `json:"changes,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID string `json:"companyId"`
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
	CompletedOnUtc *string                                           `json:"completedOnUtc,omitempty"`
	Data           *AccountingCreateInvoiceResponseAccountingInvoice `json:"data,omitempty"`
	// Unique identifier for a company's data connection.
	DataConnectionKey string `json:"dataConnectionKey"`
	// Available data types
	DataType *DataType `json:"dataType,omitempty"`
	// A message about the error.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// A unique identifier generated by Codat to represent this single push operation. This identifier can be used to track the status of the push, and should be persisted.
	PushOperationKey string `json:"pushOperationKey"`
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
	RequestedOnUtc string `json:"requestedOnUtc"`
	// The current status of the push operation.
	Status PushOperationStatus `json:"status"`
	// Push status code.
	StatusCode int64 `json:"statusCode"`
	// Number of minutes the push operation must complete within before it times out.
	TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
	// Number of seconds the push operation must complete within before it times out.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	TimeoutInSeconds *int `json:"timeoutInSeconds,omitempty"`
	// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
	Validation *Validation `json:"validation,omitempty"`
}

func (o *AccountingCreateInvoiceResponse) GetChanges() []PushOperationChange {
	if o == nil {
		return nil
	}
	return o.Changes
}

func (o *AccountingCreateInvoiceResponse) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *AccountingCreateInvoiceResponse) GetCompletedOnUtc() *string {
	if o == nil {
		return nil
	}
	return o.CompletedOnUtc
}

func (o *AccountingCreateInvoiceResponse) GetData() *AccountingCreateInvoiceResponseAccountingInvoice {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *AccountingCreateInvoiceResponse) GetDataConnectionKey() string {
	if o == nil {
		return ""
	}
	return o.DataConnectionKey
}

func (o *AccountingCreateInvoiceResponse) GetDataType() *DataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *AccountingCreateInvoiceResponse) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *AccountingCreateInvoiceResponse) GetPushOperationKey() string {
	if o == nil {
		return ""
	}
	return o.PushOperationKey
}

func (o *AccountingCreateInvoiceResponse) GetRequestedOnUtc() string {
	if o == nil {
		return ""
	}
	return o.RequestedOnUtc
}

func (o *AccountingCreateInvoiceResponse) GetStatus() PushOperationStatus {
	if o == nil {
		return PushOperationStatus("")
	}
	return o.Status
}

func (o *AccountingCreateInvoiceResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccountingCreateInvoiceResponse) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

func (o *AccountingCreateInvoiceResponse) GetTimeoutInSeconds() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInSeconds
}

func (o *AccountingCreateInvoiceResponse) GetValidation() *Validation {
	if o == nil {
		return nil
	}
	return o.Validation
}
