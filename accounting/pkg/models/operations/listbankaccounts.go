package operations

import (
	"net/http"
	"time"
)

type ListBankAccountsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListBankAccountsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListBankAccountsRequest struct {
	PathParams  ListBankAccountsPathParams
	QueryParams ListBankAccountsQueryParams
}

type ListBankAccountsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListBankAccountsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListBankAccountsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListBankAccountsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListBankAccountsLinksLinks struct {
	Current  ListBankAccountsLinksLinksCurrent   `json:"current"`
	Next     *ListBankAccountsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListBankAccountsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListBankAccountsLinksLinksSelf      `json:"self"`
}

type ListBankAccountsLinksSourceModifiedDateAccountTypeEnum string

const (
	ListBankAccountsLinksSourceModifiedDateAccountTypeEnumUnknown ListBankAccountsLinksSourceModifiedDateAccountTypeEnum = "Unknown"
	ListBankAccountsLinksSourceModifiedDateAccountTypeEnumCredit  ListBankAccountsLinksSourceModifiedDateAccountTypeEnum = "Credit"
	ListBankAccountsLinksSourceModifiedDateAccountTypeEnumDebit   ListBankAccountsLinksSourceModifiedDateAccountTypeEnum = "Debit"
)

type ListBankAccountsLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// ListBankAccountsLinksSourceModifiedDate
// > **Accessing Bank Accounts through Banking API**
// >
// > This datatype was originally used for accessing bank account data both in accounting integrations and open banking aggregators.
// >
// > To view bank account data through the Banking API, please refer to the new datatype [here](https://docs.codat.io/banking-api#/schemas/Account)
//
// > View the coverage for bank accounts in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A list of bank accounts associated with a company and a specific [data connection](https://api.codat.io/swagger/index.html#/Connection/get_companies__companyId__connections__connectionId_).
//
// Bank accounts data includes:
// * The name and ID of the account in the accounting platform.
// * The currency and balance of the account.
// * The sort code and account number.
type ListBankAccountsLinksSourceModifiedDate struct {
	AccountName        *string                                                 `json:"accountName,omitempty"`
	AccountNumber      *string                                                 `json:"accountNumber,omitempty"`
	AccountType        *ListBankAccountsLinksSourceModifiedDateAccountTypeEnum `json:"accountType,omitempty"`
	AvailableBalance   *float64                                                `json:"availableBalance,omitempty"`
	Balance            *float64                                                `json:"balance,omitempty"`
	Currency           *string                                                 `json:"currency,omitempty"`
	IBan               *string                                                 `json:"iBan,omitempty"`
	ID                 *string                                                 `json:"id,omitempty"`
	Institution        *string                                                 `json:"institution,omitempty"`
	Metadata           *ListBankAccountsLinksSourceModifiedDateMetadata        `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                              `json:"modifiedDate,omitempty"`
	NominalCode        *string                                                 `json:"nominalCode,omitempty"`
	OverdraftLimit     *float64                                                `json:"overdraftLimit,omitempty"`
	SortCode           *string                                                 `json:"sortCode,omitempty"`
	SourceModifiedDate *time.Time                                              `json:"sourceModifiedDate,omitempty"`
}

// ListBankAccountsLinks
// Codat's Paging Model
type ListBankAccountsLinks struct {
	Links        ListBankAccountsLinksLinks                `json:"_links"`
	PageNumber   int64                                     `json:"pageNumber"`
	PageSize     int64                                     `json:"pageSize"`
	Results      []ListBankAccountsLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                     `json:"totalResults"`
}

type ListBankAccountsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListBankAccountsLinks
}
