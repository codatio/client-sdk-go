package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type ListAccountTransactionsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListAccountTransactionsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListAccountTransactionsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListAccountTransactionsRequest struct {
	PathParams  ListAccountTransactionsPathParams
	QueryParams ListAccountTransactionsQueryParams
	Security    ListAccountTransactionsSecurity
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

// ListAccountTransactionsLinksSourceModifiedDateBankAccountRef
// Reference to the bank account the account transaction is recorded against.
type ListAccountTransactionsLinksSourceModifiedDateBankAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ListAccountTransactionsLinksSourceModifiedDateLinesRecordRef
// Links an account transaction line to the underlying record that created it.
type ListAccountTransactionsLinksSourceModifiedDateLinesRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type ListAccountTransactionsLinksSourceModifiedDateLines struct {
	Amount      *float64                                                      `json:"amount,omitempty"`
	Description *string                                                       `json:"description,omitempty"`
	RecordRef   *ListAccountTransactionsLinksSourceModifiedDateLinesRecordRef `json:"recordRef,omitempty"`
}

type ListAccountTransactionsLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type ListAccountTransactionsLinksSourceModifiedDateStatusEnum string

const (
	ListAccountTransactionsLinksSourceModifiedDateStatusEnumUnknown      ListAccountTransactionsLinksSourceModifiedDateStatusEnum = "Unknown"
	ListAccountTransactionsLinksSourceModifiedDateStatusEnumUnreconciled ListAccountTransactionsLinksSourceModifiedDateStatusEnum = "Unreconciled"
	ListAccountTransactionsLinksSourceModifiedDateStatusEnumReconciled   ListAccountTransactionsLinksSourceModifiedDateStatusEnum = "Reconciled"
	ListAccountTransactionsLinksSourceModifiedDateStatusEnumVoid         ListAccountTransactionsLinksSourceModifiedDateStatusEnum = "Void"
)

// ListAccountTransactionsLinksSourceModifiedDate
// > **Language tip:** In Codat, account transactions represent all transactions posted to a bank account within an accounting platform. For bank transactions posted within a banking platform, refer to [Banking transactions](https://docs.codat.io/banking-api#/operations/list-all-banking-transactions).
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
	BankAccountRef     *ListAccountTransactionsLinksSourceModifiedDateBankAccountRef `json:"bankAccountRef,omitempty"`
	Currency           *string                                                       `json:"currency,omitempty"`
	CurrencyRate       *float64                                                      `json:"currencyRate,omitempty"`
	Date               *time.Time                                                    `json:"date,omitempty"`
	ID                 *string                                                       `json:"id,omitempty"`
	Lines              []ListAccountTransactionsLinksSourceModifiedDateLines         `json:"lines,omitempty"`
	Metadata           *ListAccountTransactionsLinksSourceModifiedDateMetadata       `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                    `json:"modifiedDate,omitempty"`
	Note               *string                                                       `json:"note,omitempty"`
	SourceModifiedDate *time.Time                                                    `json:"sourceModifiedDate,omitempty"`
	Status             *ListAccountTransactionsLinksSourceModifiedDateStatusEnum     `json:"status,omitempty"`
	TotalAmount        *float64                                                      `json:"totalAmount,omitempty"`
	TransactionID      *string                                                       `json:"transactionId,omitempty"`
}

// ListAccountTransactionsLinks
// Codat's Paging Model
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
	Links       *ListAccountTransactionsLinks
}
