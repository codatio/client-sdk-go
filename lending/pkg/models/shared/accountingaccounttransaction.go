// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// AccountingAccountTransactionStatus - The status of the account transaction.
type AccountingAccountTransactionStatus string

const (
	AccountingAccountTransactionStatusUnknown      AccountingAccountTransactionStatus = "Unknown"
	AccountingAccountTransactionStatusUnreconciled AccountingAccountTransactionStatus = "Unreconciled"
	AccountingAccountTransactionStatusReconciled   AccountingAccountTransactionStatus = "Reconciled"
	AccountingAccountTransactionStatusVoid         AccountingAccountTransactionStatus = "Void"
)

func (e AccountingAccountTransactionStatus) ToPointer() *AccountingAccountTransactionStatus {
	return &e
}
func (e *AccountingAccountTransactionStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Unreconciled":
		fallthrough
	case "Reconciled":
		fallthrough
	case "Void":
		*e = AccountingAccountTransactionStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountingAccountTransactionStatus: %v", v)
	}
}

// AccountingAccountTransaction - > **Language tip:** In Codat, account transactions represent all transactions posted to a bank account within an accounting software. For bank transactions posted within a banking platform, refer to [Banking transactions](https://docs.codat.io/lending-api#/operations/list-all-banking-transactions).
//
// ## Overview
//
// In Codat’s data model, account transactions represent bank activity within an accounting software. All transactions that go through a bank account are recorded as account transactions.
//
// Account transactions are created as a result of different business activities, for example:
//
// * Payments: for example, receiving money for payment against an invoice.
// * Bill payments: for example, spending money for a payment against a bill.
// * Direct costs: for example, withdrawing money from a bank account, either for cash purposes or to make a payment.
// * Direct incomes: for example, selling an item directly to a contact and receiving payment at point of sale.
// * Transfers: for example, transferring money between two bank accounts.
//
// Account transactions is the parent data type of [payments](https://docs.codat.io/lending-api#/schemas/Payment), [bill payments](https://docs.codat.io/lending-api#/schemas/BillPayment), [direct costs](https://docs.codat.io/lending-api#/schemas/DirectCost), [direct incomes](https://docs.codat.io/lending-api#/schemas/DirectIncome), and [transfers](https://docs.codat.io/lending-api#/schemas/Transfer).
type AccountingAccountTransaction struct {
	// Links to the Account transactions data type.
	BankAccountRef *BankAccountRef `json:"bankAccountRef,omitempty"`
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
	// Identifier of the direct cost (unique to the company).
	ID *string `json:"id,omitempty"`
	// Array of account transaction lines.
	Lines        []AccountTransactionLine `json:"lines,omitempty"`
	Metadata     *Metadata                `json:"metadata,omitempty"`
	ModifiedDate *string                  `json:"modifiedDate,omitempty"`
	// Additional information about the account transaction, if available.
	Note               *string `json:"note,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// The status of the account transaction.
	Status *AccountingAccountTransactionStatus `json:"status,omitempty"`
	// Total amount of the account transactions, inclusive of tax.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount,omitempty"`
	// Identifier of the transaction (unique to the company).
	TransactionID *string `json:"transactionId,omitempty"`
}

func (a AccountingAccountTransaction) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingAccountTransaction) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingAccountTransaction) GetBankAccountRef() *BankAccountRef {
	if o == nil {
		return nil
	}
	return o.BankAccountRef
}

func (o *AccountingAccountTransaction) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountingAccountTransaction) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *AccountingAccountTransaction) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *AccountingAccountTransaction) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingAccountTransaction) GetLines() []AccountTransactionLine {
	if o == nil {
		return nil
	}
	return o.Lines
}

func (o *AccountingAccountTransaction) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AccountingAccountTransaction) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *AccountingAccountTransaction) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *AccountingAccountTransaction) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *AccountingAccountTransaction) GetStatus() *AccountingAccountTransactionStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *AccountingAccountTransaction) GetTotalAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *AccountingAccountTransaction) GetTransactionID() *string {
	if o == nil {
		return nil
	}
	return o.TransactionID
}
