// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// SalesOrderContact - Details of the named contact at the delivery address.
type SalesOrderContact struct {
	// Email address of the contact at the delivery address.
	Email *string `json:"email,omitempty"`
	// Name of the contact at the delivery address.
	Name *string `json:"name,omitempty"`
	// Phone number of the contact at the delivery address.
	Phone *string `json:"phone,omitempty"`
}

func (o *SalesOrderContact) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *SalesOrderContact) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SalesOrderContact) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

// SalesOrderShipTo - Delivery details for any goods that have been ordered.
type SalesOrderShipTo struct {
	Address *Items `json:"address,omitempty"`
	// Details of the named contact at the delivery address.
	Contact *SalesOrderContact `json:"contact,omitempty"`
}

func (o *SalesOrderShipTo) GetAddress() *Items {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *SalesOrderShipTo) GetContact() *SalesOrderContact {
	if o == nil {
		return nil
	}
	return o.Contact
}

// SalesOrder - > View the coverage for sales orders in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=salesOrders" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A sales order represents a customer's intention to purchase goods or services from a seller and usually includes information such as the expected delivery date and shipping details. This information can be used to provide visibility on a business's expected receivables and to track sales through the full procurement process.
//
// A sales order is typically converted to an [invoice](https://docs.codat.io/accounting-api#/schemas/Invoice) after approval.
type SalesOrder struct {
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
	// A customer-supplied identifier for the purchase order in the customer's system.
	CustomerPurchaseOrderNumber *string                `json:"customerPurchaseOrderNumber,omitempty"`
	CustomerRef                 *AccountingCustomerRef `json:"customerRef,omitempty"`
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
	ExpectedDeliveryDate *string `json:"expectedDeliveryDate,omitempty"`
	// Identifier for the sales order, unique for the company in the accounting software.
	ID *string `json:"id,omitempty"`
	// If the sales order is converted to an invoice, or will be in future, the invoicingStatus field indicates the current stage of the invoicing process.
	InvoicingStatus *SalesOrderInvoiceStatus `json:"invoicingStatus,omitempty"`
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
	// An array of line items.
	LineItems    []SalesOrderLineItem `json:"lineItems,omitempty"`
	Metadata     *Metadata            `json:"metadata,omitempty"`
	ModifiedDate *string              `json:"modifiedDate,omitempty"`
	// Any additional information associated with the sales order.
	Note *string `json:"note,omitempty"`
	// Friendly reference for the sales order, commonly generated by the accounting software.
	SalesOrderNumber *string `json:"salesOrderNumber,omitempty"`
	// Delivery details for any goods that have been ordered.
	ShipTo             *SalesOrderShipTo `json:"shipTo,omitempty"`
	SourceModifiedDate *string           `json:"sourceModifiedDate,omitempty"`
	// Current state of the sales order.
	Status *SalesOrderStatus `json:"status,omitempty"`
	// Total amount of the sales order, including discounts but excluding tax.
	SubTotal *decimal.Big `decimal:"number" json:"subTotal,omitempty"`
	// Total amount of the sales order, including discounts and tax.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount,omitempty"`
	// Total value of any discounts applied to the sales order.
	TotalDiscount *decimal.Big `decimal:"number" json:"totalDiscount,omitempty"`
	// Total amount of tax included in the sales order.
	TotalTaxAmount *decimal.Big `decimal:"number" json:"totalTaxAmount,omitempty"`
}

func (s SalesOrder) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SalesOrder) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SalesOrder) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *SalesOrder) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *SalesOrder) GetCustomerPurchaseOrderNumber() *string {
	if o == nil {
		return nil
	}
	return o.CustomerPurchaseOrderNumber
}

func (o *SalesOrder) GetCustomerRef() *AccountingCustomerRef {
	if o == nil {
		return nil
	}
	return o.CustomerRef
}

func (o *SalesOrder) GetExpectedDeliveryDate() *string {
	if o == nil {
		return nil
	}
	return o.ExpectedDeliveryDate
}

func (o *SalesOrder) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SalesOrder) GetInvoicingStatus() *SalesOrderInvoiceStatus {
	if o == nil {
		return nil
	}
	return o.InvoicingStatus
}

func (o *SalesOrder) GetIssueDate() *string {
	if o == nil {
		return nil
	}
	return o.IssueDate
}

func (o *SalesOrder) GetLineItems() []SalesOrderLineItem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *SalesOrder) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *SalesOrder) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *SalesOrder) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *SalesOrder) GetSalesOrderNumber() *string {
	if o == nil {
		return nil
	}
	return o.SalesOrderNumber
}

func (o *SalesOrder) GetShipTo() *SalesOrderShipTo {
	if o == nil {
		return nil
	}
	return o.ShipTo
}

func (o *SalesOrder) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *SalesOrder) GetStatus() *SalesOrderStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *SalesOrder) GetSubTotal() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.SubTotal
}

func (o *SalesOrder) GetTotalAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *SalesOrder) GetTotalDiscount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalDiscount
}

func (o *SalesOrder) GetTotalTaxAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalTaxAmount
}
