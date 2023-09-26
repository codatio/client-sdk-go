// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type UpdateExpenseRequest struct {
	ContactRef *ContactRef `json:"contactRef,omitempty"`
	// Currency the transaction was recorded in.
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
	// Date the transaction was recorded.
	IssueDate string `json:"issueDate"`
	// Array of transaction lines.
	Lines []ExpenseTransactionLine `json:"lines,omitempty"`
	// Name of the merchant where the purchase took place
	MerchantName *string `json:"merchantName,omitempty"`
	// Any private, company notes about the transaction.
	Notes *string     `json:"notes,omitempty"`
	Type  interface{} `json:"type"`
}

func (u UpdateExpenseRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateExpenseRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateExpenseRequest) GetContactRef() *ContactRef {
	if o == nil {
		return nil
	}
	return o.ContactRef
}

func (o *UpdateExpenseRequest) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *UpdateExpenseRequest) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *UpdateExpenseRequest) GetIssueDate() string {
	if o == nil {
		return ""
	}
	return o.IssueDate
}

func (o *UpdateExpenseRequest) GetLines() []ExpenseTransactionLine {
	if o == nil {
		return nil
	}
	return o.Lines
}

func (o *UpdateExpenseRequest) GetMerchantName() *string {
	if o == nil {
		return nil
	}
	return o.MerchantName
}

func (o *UpdateExpenseRequest) GetNotes() *string {
	if o == nil {
		return nil
	}
	return o.Notes
}

func (o *UpdateExpenseRequest) GetType() interface{} {
	if o == nil {
		return nil
	}
	return o.Type
}
