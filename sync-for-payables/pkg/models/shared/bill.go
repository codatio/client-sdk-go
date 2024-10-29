// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// Bill - Bills are invoices that represent the SMB's financial obligations to their supplier for a purchase of goods or services.
type Bill struct {
	// Identifier for the bill, unique for the company in the accounting software.
	ID *string `json:"id,omitempty"`
	// User-friendly reference for the bill.
	Reference *string `json:"reference,omitempty"`
	// Reference to the supplier the record relates to.
	SupplierRef SupplierRef `json:"supplierRef"`
	IssueDate   string      `json:"issueDate"`
	DueDate     string      `json:"dueDate"`
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
	// Array of Bill line items.
	LineItems []BillLineItem `json:"lineItems,omitempty"`
	// Current state of the bill. If creating a bill the status must be `Open`.
	Status BillStatus `json:"status"`
	// Amount of the bill, including tax.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount,omitempty"`
	// Amount outstanding on the bill.
	AmountDue          *decimal.Big `decimal:"number" json:"amountDue,omitempty"`
	SourceModifiedDate *string      `json:"sourceModifiedDate,omitempty"`
}

func (b Bill) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *Bill) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Bill) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Bill) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *Bill) GetSupplierRef() SupplierRef {
	if o == nil {
		return SupplierRef{}
	}
	return o.SupplierRef
}

func (o *Bill) GetIssueDate() string {
	if o == nil {
		return ""
	}
	return o.IssueDate
}

func (o *Bill) GetDueDate() string {
	if o == nil {
		return ""
	}
	return o.DueDate
}

func (o *Bill) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *Bill) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *Bill) GetLineItems() []BillLineItem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *Bill) GetStatus() BillStatus {
	if o == nil {
		return BillStatus("")
	}
	return o.Status
}

func (o *Bill) GetTotalAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *Bill) GetAmountDue() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.AmountDue
}

func (o *Bill) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}
