// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type Payment struct {
	// Payment amount.
	Amount                 *decimal.Big            `decimal:"number" json:"amount,omitempty"`
	BankingTransactionRefs []BankingTransactionRef `json:"bankingTransactionRefs,omitempty"`
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
	Date *string `json:"date,omitempty"`
	// ID of the invoice, which may be a GUID but it may be something else depending on the accounting platform.
	ID *string `json:"id,omitempty"`
	// The type of payment.
	PaymentType *string `json:"paymentType,omitempty"`
}

func (p Payment) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Payment) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Payment) GetAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *Payment) GetBankingTransactionRefs() []BankingTransactionRef {
	if o == nil {
		return nil
	}
	return o.BankingTransactionRefs
}

func (o *Payment) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *Payment) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *Payment) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *Payment) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Payment) GetPaymentType() *string {
	if o == nil {
		return nil
	}
	return o.PaymentType
}
