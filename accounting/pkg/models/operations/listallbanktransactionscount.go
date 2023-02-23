package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type ListAllBankTransactionscountPathParams struct {
	AccountID    string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListAllBankTransactionscountQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListAllBankTransactionscountSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListAllBankTransactionscountRequest struct {
	PathParams  ListAllBankTransactionscountPathParams
	QueryParams ListAllBankTransactionscountQueryParams
	Security    ListAllBankTransactionscountSecurity
}

type ListAllBankTransactionscountLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListAllBankTransactionscountLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListAllBankTransactionscountLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListAllBankTransactionscountLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListAllBankTransactionscountLinksLinks struct {
	Current  ListAllBankTransactionscountLinksLinksCurrent   `json:"current"`
	Next     *ListAllBankTransactionscountLinksLinksNext     `json:"next,omitempty"`
	Previous *ListAllBankTransactionscountLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListAllBankTransactionscountLinksLinksSelf      `json:"self"`
}

type ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum string

const (
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumUnknown     ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "Unknown"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumCredit      ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "Credit"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumDebit       ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "Debit"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumInt         ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "Int"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumDiv         ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "Div"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumFee         ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "Fee"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumSerChg      ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "SerChg"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumDep         ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "Dep"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumAtm         ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "Atm"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumPos         ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "Pos"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumXfer        ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "Xfer"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumCheck       ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "Check"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumPayment     ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "Payment"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumCash        ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "Cash"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumDirectDep   ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "DirectDep"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumDirectDebit ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "DirectDebit"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumRepeatPmt   ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "RepeatPmt"
	ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnumOther       ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum = "Other"
)

type ListAllBankTransactionscountLinksResultsTransactions struct {
	Amount             float64                                                                 `json:"amount"`
	Balance            float64                                                                 `json:"balance"`
	Counterparty       *string                                                                 `json:"counterparty,omitempty"`
	Date               time.Time                                                               `json:"date"`
	Description        *string                                                                 `json:"description,omitempty"`
	ID                 *string                                                                 `json:"id,omitempty"`
	ModifiedDate       *time.Time                                                              `json:"modifiedDate,omitempty"`
	Reconciled         bool                                                                    `json:"reconciled"`
	Reference          *string                                                                 `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                                              `json:"sourceModifiedDate,omitempty"`
	TransactionType    ListAllBankTransactionscountLinksResultsTransactionsTransactionTypeEnum `json:"transactionType"`
}

// ListAllBankTransactionscountLinksResults
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
type ListAllBankTransactionscountLinksResults struct {
	AccountID       *string                                                `json:"accountId,omitempty"`
	ContractVersion *string                                                `json:"contractVersion,omitempty"`
	Transactions    []ListAllBankTransactionscountLinksResultsTransactions `json:"transactions,omitempty"`
}

// ListAllBankTransactionscountLinks
// Codat's Paging Model
type ListAllBankTransactionscountLinks struct {
	Links        ListAllBankTransactionscountLinksLinks     `json:"_links"`
	PageNumber   int64                                      `json:"pageNumber"`
	PageSize     int64                                      `json:"pageSize"`
	Results      []ListAllBankTransactionscountLinksResults `json:"results,omitempty"`
	TotalResults int64                                      `json:"totalResults"`
}

type ListAllBankTransactionscountResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListAllBankTransactionscountLinks
}
