package operations

import (
	"net/http"
	"time"
)

type ListBankTransactionsRequest struct {
	AccountID string  `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID string  `pathParam:"style=simple,explode=false,name=companyId"`
	OrderBy   *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page      int     `queryParam:"style=form,explode=true,name=page"`
	PageSize  *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query     *string `queryParam:"style=form,explode=true,name=query"`
}

type ListBankTransactionsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListBankTransactionsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListBankTransactionsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListBankTransactionsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListBankTransactionsLinksLinks struct {
	Current  ListBankTransactionsLinksLinksCurrent   `json:"current"`
	Next     *ListBankTransactionsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListBankTransactionsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListBankTransactionsLinksLinksSelf      `json:"self"`
}

type ListBankTransactionsLinksResultsTransactionTypeEnum string

const (
	ListBankTransactionsLinksResultsTransactionTypeEnumUnknown     ListBankTransactionsLinksResultsTransactionTypeEnum = "Unknown"
	ListBankTransactionsLinksResultsTransactionTypeEnumCredit      ListBankTransactionsLinksResultsTransactionTypeEnum = "Credit"
	ListBankTransactionsLinksResultsTransactionTypeEnumDebit       ListBankTransactionsLinksResultsTransactionTypeEnum = "Debit"
	ListBankTransactionsLinksResultsTransactionTypeEnumInt         ListBankTransactionsLinksResultsTransactionTypeEnum = "Int"
	ListBankTransactionsLinksResultsTransactionTypeEnumDiv         ListBankTransactionsLinksResultsTransactionTypeEnum = "Div"
	ListBankTransactionsLinksResultsTransactionTypeEnumFee         ListBankTransactionsLinksResultsTransactionTypeEnum = "Fee"
	ListBankTransactionsLinksResultsTransactionTypeEnumSerChg      ListBankTransactionsLinksResultsTransactionTypeEnum = "SerChg"
	ListBankTransactionsLinksResultsTransactionTypeEnumDep         ListBankTransactionsLinksResultsTransactionTypeEnum = "Dep"
	ListBankTransactionsLinksResultsTransactionTypeEnumAtm         ListBankTransactionsLinksResultsTransactionTypeEnum = "Atm"
	ListBankTransactionsLinksResultsTransactionTypeEnumPos         ListBankTransactionsLinksResultsTransactionTypeEnum = "Pos"
	ListBankTransactionsLinksResultsTransactionTypeEnumXfer        ListBankTransactionsLinksResultsTransactionTypeEnum = "Xfer"
	ListBankTransactionsLinksResultsTransactionTypeEnumCheck       ListBankTransactionsLinksResultsTransactionTypeEnum = "Check"
	ListBankTransactionsLinksResultsTransactionTypeEnumPayment     ListBankTransactionsLinksResultsTransactionTypeEnum = "Payment"
	ListBankTransactionsLinksResultsTransactionTypeEnumCash        ListBankTransactionsLinksResultsTransactionTypeEnum = "Cash"
	ListBankTransactionsLinksResultsTransactionTypeEnumDirectDep   ListBankTransactionsLinksResultsTransactionTypeEnum = "DirectDep"
	ListBankTransactionsLinksResultsTransactionTypeEnumDirectDebit ListBankTransactionsLinksResultsTransactionTypeEnum = "DirectDebit"
	ListBankTransactionsLinksResultsTransactionTypeEnumRepeatPmt   ListBankTransactionsLinksResultsTransactionTypeEnum = "RepeatPmt"
	ListBankTransactionsLinksResultsTransactionTypeEnumOther       ListBankTransactionsLinksResultsTransactionTypeEnum = "Other"
)

type ListBankTransactionsLinksResults struct {
	Amount             float64                                             `json:"amount"`
	Balance            float64                                             `json:"balance"`
	Counterparty       *string                                             `json:"counterparty,omitempty"`
	Date               time.Time                                           `json:"date"`
	Description        *string                                             `json:"description,omitempty"`
	ID                 *string                                             `json:"id,omitempty"`
	ModifiedDate       *time.Time                                          `json:"modifiedDate,omitempty"`
	Reconciled         bool                                                `json:"reconciled"`
	Reference          *string                                             `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                          `json:"sourceModifiedDate,omitempty"`
	TransactionType    ListBankTransactionsLinksResultsTransactionTypeEnum `json:"transactionType"`
}

// ListBankTransactionsLinks
// Codat's Paging Model
type ListBankTransactionsLinks struct {
	Links        ListBankTransactionsLinksLinks     `json:"_links"`
	PageNumber   int64                              `json:"pageNumber"`
	PageSize     int64                              `json:"pageSize"`
	Results      []ListBankTransactionsLinksResults `json:"results,omitempty"`
	TotalResults int64                              `json:"totalResults"`
}

type ListBankTransactionsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListBankTransactionsLinks
}
