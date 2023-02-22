package shared

import (
	"time"
)

// AccountTransactionBankAccountRef
// Reference to the bank account the account transaction is recorded against.
type AccountTransactionBankAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type AccountTransactionLines struct {
	Amount      *float64   `json:"amount,omitempty"`
	Description *string    `json:"description,omitempty"`
	RecordRef   *RecordRef `json:"recordRef,omitempty"`
}

type AccountTransactionStatusEnum string

const (
	AccountTransactionStatusEnumUnknown      AccountTransactionStatusEnum = "Unknown"
	AccountTransactionStatusEnumUnreconciled AccountTransactionStatusEnum = "Unreconciled"
	AccountTransactionStatusEnumReconciled   AccountTransactionStatusEnum = "Reconciled"
	AccountTransactionStatusEnumVoid         AccountTransactionStatusEnum = "Void"
)

// AccountTransaction
// > **Language tip:** In Codat, account transactions represent all transactions posted to a bank account within an accounting platform. For bank transactions posted within a banking platform, refer to [Banking transactions](https://docs.codat.io/docs/datamodel-banking-banking-transactions).
//
// > View the coverage for account transactions in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=accountTransactions" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// In Codat’s data model, account transactions represent bank activity within an accounting platform. All transactions that go through a bank account are recorded as account transactions.
//
// Account transactions are created as a result of different business activities, for example:
//
// * Payments: for example, receiving money for payment against an invoice.
// * Bill payments: for example, spending money for a payment against a bill.
// * Direct costs: for example, withdrawing money from a bank account, either for cash purposes or to make a payment.
// * Direct incomes: for example, selling an item directly to a contact and receiving payment at point of sale.
// * Transfers: for example, transferring money between two bank accounts.
//
// Account transactions is the parent data type of [payments](https://docs.codat.io/docs/datamodel-accounting-payments), [bill payments](https://docs.codat.io/docs/datamodel-accounting-billpayments), [direct costs](https://docs.codat.io/docs/datamodel-accounting-directcosts), [direct incomes](https://docs.codat.io/docs/datamodel-accounting-directincomes), and [transfers](https://docs.codat.io/docs/datamodel-accounting-transfers).
//
// Perform the following tasks using the <a className="external" href="https://api.codat.io/swagger/index.html#/AccountTransactions" target="_blank">Account transactions</a> endpoints:
//
// * <a className="external" href="https://api.codat.io/swagger/index.html#/AccountTransactions/get_companies__companyId__connections__connectionId__data_accountTransactions" target="_blank">Get a list of account transactions</a>
// * <a className="external" href="https://api.codat.io/swagger/index.html#/AccountTransactions/get_companies__companyId__connections__connectionId__data_accountTransactions__accountTransactionId_" target="_blank">Get a single account transaction</a>
type AccountTransaction struct {
	BankAccountRef     *AccountTransactionBankAccountRef `json:"bankAccountRef,omitempty"`
	Currency           *string                           `json:"currency,omitempty"`
	CurrencyRate       *float64                          `json:"currencyRate,omitempty"`
	Date               *time.Time                        `json:"date,omitempty"`
	ID                 *string                           `json:"id,omitempty"`
	Lines              []AccountTransactionLines         `json:"lines,omitempty"`
	Metadata           *Metadata                         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                        `json:"modifiedDate,omitempty"`
	Note               *string                           `json:"note,omitempty"`
	SourceModifiedDate *time.Time                        `json:"sourceModifiedDate,omitempty"`
	Status             *AccountTransactionStatusEnum     `json:"status,omitempty"`
	TotalAmount        *float64                          `json:"totalAmount,omitempty"`
	TransactionID      *string                           `json:"transactionId,omitempty"`
}
