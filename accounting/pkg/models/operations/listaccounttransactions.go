// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"time"
)

type ListAccountTransactionsRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Field to order results by. [Read more](https://docs.codat.io/using-the-api/ordering-results).
	OrderBy *string `queryParam:"style=form,explode=true,name=orderBy"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page int `queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `queryParam:"style=form,explode=true,name=pageSize"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

type ListAccountTransactionsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListAccountTransactionsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListAccountTransactionsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListAccountTransactionsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListAccountTransactionsLinksLinks struct {
	Current  ListAccountTransactionsLinksLinksCurrent   `json:"current"`
	Next     *ListAccountTransactionsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListAccountTransactionsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListAccountTransactionsLinksLinksSelf      `json:"self"`
}

// ListAccountTransactionsLinksSourceModifiedDateBankAccountRef - Reference to the bank account the account transaction is recorded against.
type ListAccountTransactionsLinksSourceModifiedDateBankAccountRef struct {
	// Bank account 'id' for the account transaction.
	ID *string `json:"id,omitempty"`
	// bank account 'name' for the account transaction.
	Name *string `json:"name,omitempty"`
}

// ListAccountTransactionsLinksSourceModifiedDateLinesRecordRef - Links an account transaction line to the underlying record that created it.
type ListAccountTransactionsLinksSourceModifiedDateLinesRecordRef struct {
	// Name of the 'dataType'.
	DataType *string `json:"dataType,omitempty"`
	// 'id' of the underlying record or data type.
	ID *string `json:"id,omitempty"`
}

type ListAccountTransactionsLinksSourceModifiedDateLines struct {
	// Amount in the bill payment currency.
	Amount *float64 `json:"amount,omitempty"`
	// Description of the account transaction.
	Description *string `json:"description,omitempty"`
	// Links an account transaction line to the underlying record that created it.
	RecordRef *ListAccountTransactionsLinksSourceModifiedDateLinesRecordRef `json:"recordRef,omitempty"`
}

type ListAccountTransactionsLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// ListAccountTransactionsLinksSourceModifiedDateStatusEnum - The status of the account transaction.
type ListAccountTransactionsLinksSourceModifiedDateStatusEnum string

const (
	ListAccountTransactionsLinksSourceModifiedDateStatusEnumUnknown      ListAccountTransactionsLinksSourceModifiedDateStatusEnum = "Unknown"
	ListAccountTransactionsLinksSourceModifiedDateStatusEnumUnreconciled ListAccountTransactionsLinksSourceModifiedDateStatusEnum = "Unreconciled"
	ListAccountTransactionsLinksSourceModifiedDateStatusEnumReconciled   ListAccountTransactionsLinksSourceModifiedDateStatusEnum = "Reconciled"
	ListAccountTransactionsLinksSourceModifiedDateStatusEnumVoid         ListAccountTransactionsLinksSourceModifiedDateStatusEnum = "Void"
)

// ListAccountTransactionsLinksSourceModifiedDate - > **Language tip:** In Codat, account transactions represent all transactions posted to a bank account within an accounting platform. For bank transactions posted within a banking platform, refer to [Banking transactions](https://docs.codat.io/banking-api#/operations/list-all-banking-transactions).
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
// Account transactions is the parent data type of [payments](https://docs.codat.io/accounting-api#/schemas/Payment), [bill payments](https://docs.codat.io/accounting-api#/schemas/BillPayment), [direct costs](https://docs.codat.io/accounting-api#/schemas/DirectCost), [direct incomes](https://docs.codat.io/accounting-api#/schemas/DirectIncome), and [transfers](https://docs.codat.io/accounting-api#/schemas/Transfer).
type ListAccountTransactionsLinksSourceModifiedDate struct {
	// Reference to the bank account the account transaction is recorded against.
	BankAccountRef *ListAccountTransactionsLinksSourceModifiedDateBankAccountRef `json:"bankAccountRef,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code. e.g. _GBP_.
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
	CurrencyRate *float64 `json:"currencyRate,omitempty"`
	// The date the account transaction was recorded in the platform.
	Date *time.Time `json:"date,omitempty"`
	// Identifier of the direct cost (unique to the company).
	ID *string `json:"id,omitempty"`
	// Array of account transaction lines.
	Lines    []ListAccountTransactionsLinksSourceModifiedDateLines   `json:"lines,omitempty"`
	Metadata *ListAccountTransactionsLinksSourceModifiedDateMetadata `json:"metadata,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	// Additional information about the account transaction, if available.
	Note *string `json:"note,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
	// The status of the account transaction.
	Status *ListAccountTransactionsLinksSourceModifiedDateStatusEnum `json:"status,omitempty"`
	// Total amount of the account transactions, inclusive of tax.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	// Identifier of the transaction (unique to the company).
	TransactionID *string `json:"transactionId,omitempty"`
}

// ListAccountTransactionsLinks - Codat's Paging Model
type ListAccountTransactionsLinks struct {
	Links        ListAccountTransactionsLinksLinks                `json:"_links"`
	PageNumber   int64                                            `json:"pageNumber"`
	PageSize     int64                                            `json:"pageSize"`
	Results      []ListAccountTransactionsLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                            `json:"totalResults"`
}

type ListAccountTransactionsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	Links *ListAccountTransactionsLinks
}
