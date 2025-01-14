// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type BillPaymentPrototype struct {
	// Amount of the payment in the bill currency.
	Amount *decimal.Big `decimal:"number" json:"amount"`
	Date   string       `json:"date"`
	// Additional information associated with the payment.
	Reference *string `json:"reference,omitempty"`
	// Reference to the bank account / credit card which you are using to pay the bill.
	AccountRef BillPaymentAccountRef `json:"accountRef"`
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
}

func (b BillPaymentPrototype) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BillPaymentPrototype) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BillPaymentPrototype) GetAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.Amount
}

func (o *BillPaymentPrototype) GetDate() string {
	if o == nil {
		return ""
	}
	return o.Date
}

func (o *BillPaymentPrototype) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *BillPaymentPrototype) GetAccountRef() BillPaymentAccountRef {
	if o == nil {
		return BillPaymentAccountRef{}
	}
	return o.AccountRef
}

func (o *BillPaymentPrototype) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}
