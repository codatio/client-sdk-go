// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// DirectCostContactRef - A customer or supplier associated with the direct cost.
type DirectCostContactRef struct {
	// Available Data types
	DataType *DataType `json:"dataType,omitempty"`
	// Unique identifier for a customer or supplier.
	ID string `json:"id"`
}

func (o *DirectCostContactRef) GetDataType() *DataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *DirectCostContactRef) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DirectCostAllocation struct {
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
	CurrencyRate *decimal.Big `decimal:"number" json:"currencyRate,omitempty"`
	// The total amount that has been allocated.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount,omitempty"`
}

func (d DirectCostAllocation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DirectCostAllocation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DirectCostAllocation) GetAllocatedOnDate() *string {
	if o == nil {
		return nil
	}
	return o.AllocatedOnDate
}

func (o *DirectCostAllocation) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *DirectCostAllocation) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *DirectCostAllocation) GetTotalAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

type DirectCostAccountingPaymentAllocation struct {
	Allocation DirectCostAllocation     `json:"allocation"`
	Payment    PaymentAllocationPayment `json:"payment"`
}

func (o *DirectCostAccountingPaymentAllocation) GetAllocation() DirectCostAllocation {
	if o == nil {
		return DirectCostAllocation{}
	}
	return o.Allocation
}

func (o *DirectCostAccountingPaymentAllocation) GetPayment() PaymentAllocationPayment {
	if o == nil {
		return PaymentAllocationPayment{}
	}
	return o.Payment
}

// DirectCost - > **Language tip:** Direct costs may also be referred to as **Spend transactions**, **Spend money transactions**, or **Payments** in various accounting platforms.
//
// > View the coverage for direct costs in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Direct costs are the expenses associated with a business' operations. For example, purchases of raw materials and service fees are considered direct costs.
//
// Direct costs include:
//   - Purchasing an item and paying it off at the point of the purchase
//   - Receiving cash from a refunded item if the refund is made by the supplier
//   - Withdrawing money from a bank account
//   - Writing a cheque
//
// Direct costs is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type DirectCost struct {
	// A customer or supplier associated with the direct cost.
	ContactRef *DirectCostContactRef `json:"contactRef,omitempty"`
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
	//
	//
	// ### Integration-specific details
	//
	// | Integration       | Scenario                                        | System behavior                                                                                                                                                      |
	// |-------------------|-------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	// | QuickBooks Online | Transaction currency differs from base currency | If currency rate value is left `null`, a rate of 1 will be used by QBO by default. To override this, include the required currency rate in the expense transaction.  |
	CurrencyRate *decimal.Big `decimal:"number" json:"currencyRate,omitempty"`
	// Identifier of the direct cost, unique for the company.
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
	LineItems    []DirectCostLineItem `json:"lineItems"`
	Metadata     *Metadata            `json:"metadata,omitempty"`
	ModifiedDate *string              `json:"modifiedDate,omitempty"`
	// A note attached to the direct cost.
	Note *string `json:"note,omitempty"`
	// An array of payment allocations.
	PaymentAllocations []DirectCostAccountingPaymentAllocation `json:"paymentAllocations"`
	// User-friendly reference for the direct cost.
	Reference          *string `json:"reference,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// The total amount of the direct costs, excluding any taxes.
	SubTotal *decimal.Big `decimal:"number" json:"subTotal"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// The total amount of tax on the direct costs.
	TaxAmount *decimal.Big `decimal:"number" json:"taxAmount"`
	// The amount of the direct costs, inclusive of tax.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount"`
}

func (d DirectCost) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DirectCost) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DirectCost) GetContactRef() *DirectCostContactRef {
	if o == nil {
		return nil
	}
	return o.ContactRef
}

func (o *DirectCost) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *DirectCost) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *DirectCost) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DirectCost) GetIssueDate() string {
	if o == nil {
		return ""
	}
	return o.IssueDate
}

func (o *DirectCost) GetLineItems() []DirectCostLineItem {
	if o == nil {
		return []DirectCostLineItem{}
	}
	return o.LineItems
}

func (o *DirectCost) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *DirectCost) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *DirectCost) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *DirectCost) GetPaymentAllocations() []DirectCostAccountingPaymentAllocation {
	if o == nil {
		return []DirectCostAccountingPaymentAllocation{}
	}
	return o.PaymentAllocations
}

func (o *DirectCost) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *DirectCost) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *DirectCost) GetSubTotal() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.SubTotal
}

func (o *DirectCost) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *DirectCost) GetTaxAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TaxAmount
}

func (o *DirectCost) GetTotalAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TotalAmount
}
