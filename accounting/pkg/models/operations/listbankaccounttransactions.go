package operations

import (
	"net/http"
	"time"
)

type ListBankAccountTransactionsRequest struct {
	AccountID    string  `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string  `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string  `pathParam:"style=simple,explode=false,name=connectionId"`
	OrderBy      *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page         int     `queryParam:"style=form,explode=true,name=page"`
	PageSize     *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query        *string `queryParam:"style=form,explode=true,name=query"`
}

type ListBankAccountTransactionsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListBankAccountTransactionsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListBankAccountTransactionsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListBankAccountTransactionsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListBankAccountTransactionsLinksLinks struct {
	Current  ListBankAccountTransactionsLinksLinksCurrent   `json:"current"`
	Next     *ListBankAccountTransactionsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListBankAccountTransactionsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListBankAccountTransactionsLinksLinksSelf      `json:"self"`
}

type ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum string

const (
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumUnknown     ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "Unknown"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumCredit      ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "Credit"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumDebit       ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "Debit"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumInt         ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "Int"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumDiv         ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "Div"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumFee         ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "Fee"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumSerChg      ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "SerChg"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumDep         ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "Dep"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumAtm         ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "Atm"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumPos         ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "Pos"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumXfer        ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "Xfer"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumCheck       ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "Check"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumPayment     ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "Payment"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumCash        ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "Cash"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumDirectDep   ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "DirectDep"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumDirectDebit ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "DirectDebit"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumRepeatPmt   ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "RepeatPmt"
	ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnumOther       ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum = "Other"
)

type ListBankAccountTransactionsLinksResultsTransactions struct {
	Amount             float64                                                                `json:"amount"`
	Balance            float64                                                                `json:"balance"`
	Counterparty       *string                                                                `json:"counterparty,omitempty"`
	Date               time.Time                                                              `json:"date"`
	Description        *string                                                                `json:"description,omitempty"`
	ID                 *string                                                                `json:"id,omitempty"`
	ModifiedDate       *time.Time                                                             `json:"modifiedDate,omitempty"`
	Reconciled         bool                                                                   `json:"reconciled"`
	Reference          *string                                                                `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                                             `json:"sourceModifiedDate,omitempty"`
	TransactionType    ListBankAccountTransactionsLinksResultsTransactionsTransactionTypeEnum `json:"transactionType"`
}

// ListBankAccountTransactionsLinksResults
// > **Accessing Bank Accounts through Banking API**
// >
// > This datatype was originally used for accessing bank account data both in accounting integrations and open banking aggregators.
// >
// > To view bank account data through the Banking API, please refer to the new datatype [here](https://docs.codat.io/banking-api#/operations/list-all-banking-transactions)
//
// > View the coverage for bank transactions in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankTransactions" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Transactional banking data for a specific company and account.
//
// Bank transactions include the:
// * Amount of the transaction.
// * Current account balance.
// * Transaction type, for example, credit, debit, or transfer.
type ListBankAccountTransactionsLinksResults struct {
	AccountID       *string                                               `json:"accountId,omitempty"`
	ContractVersion *string                                               `json:"contractVersion,omitempty"`
	Transactions    []ListBankAccountTransactionsLinksResultsTransactions `json:"transactions,omitempty"`
}

// ListBankAccountTransactionsLinks
// Codat's Paging Model
type ListBankAccountTransactionsLinks struct {
	Links        ListBankAccountTransactionsLinksLinks     `json:"_links"`
	PageNumber   int64                                     `json:"pageNumber"`
	PageSize     int64                                     `json:"pageSize"`
	Results      []ListBankAccountTransactionsLinksResults `json:"results,omitempty"`
	TotalResults int64                                     `json:"totalResults"`
}

type ListBankAccountTransactionsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListBankAccountTransactionsLinks
}
