// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"time"
)

type ListBankAccountsRequest struct {
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

// ListBankAccountsLinksSourceModifiedDateAccountTypeEnum - The type of the account.
type ListBankAccountsLinksSourceModifiedDateAccountTypeEnum string

const (
	ListBankAccountsLinksSourceModifiedDateAccountTypeEnumUnknown ListBankAccountsLinksSourceModifiedDateAccountTypeEnum = "Unknown"
	ListBankAccountsLinksSourceModifiedDateAccountTypeEnumCredit  ListBankAccountsLinksSourceModifiedDateAccountTypeEnum = "Credit"
	ListBankAccountsLinksSourceModifiedDateAccountTypeEnumDebit   ListBankAccountsLinksSourceModifiedDateAccountTypeEnum = "Debit"
)

type ListBankAccountsLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// ListBankAccountsLinksSourceModifiedDate - > **Accessing Bank Accounts through Banking API**
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
	// Name of the bank account in the accounting platform.
	AccountName *string `json:"accountName,omitempty"`
	// Account number for the bank account.
	//
	// Xero integrations
	// Only a UK account number shows for bank accounts with GBP currency and a combined total of sort code and account number that equals 14 digits, For non-GBP accounts, the full bank account number is populated.
	//
	// FreeAgent integrations
	// For Credit accounts, only the last four digits are required. For other types, the field is optional.
	AccountNumber *string `json:"accountNumber,omitempty"`
	// The type of the account.
	AccountType *ListBankAccountsLinksSourceModifiedDateAccountTypeEnum `json:"accountType,omitempty"`
	// Total available balance of the bank account as reported by the underlying data source. This may take into account overdrafts or pending transactions for example.
	AvailableBalance *float64 `json:"availableBalance,omitempty"`
	// Balance of the bank account.
	Balance *float64 `json:"balance,omitempty"`
	// Base currency of the bank account.
	Currency *string `json:"currency,omitempty"`
	// International bank account number of the account. Often used when making or receiving international payments.
	IBan *string `json:"iBan,omitempty"`
	// Identifier for the account, unique for the company in the accounting platform.
	ID *string `json:"id,omitempty"`
	// The institution of the bank account.
	Institution *string                                          `json:"institution,omitempty"`
	Metadata    *ListBankAccountsLinksSourceModifiedDateMetadata `json:"metadata,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	// Code used to identify each nominal account for a business.
	NominalCode *string `json:"nominalCode,omitempty"`
	// Pre-arranged overdraft limit of the account.
	//
	// The value is always positive. For example, an overdraftLimit of `1000` means that the balance of the account can go down to `-1000`.
	OverdraftLimit *float64 `json:"overdraftLimit,omitempty"`
	// Sort code for the bank account.
	//
	// Xero integrations
	// The sort code is only displayed when the currency = GBP and the sort code and account number sum to 14 digits. For non-GBP accounts, this field is not populated.
	SortCode *string `json:"sortCode,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
}

// ListBankAccountsLinks - Codat's Paging Model
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
	// Success
	Links *ListBankAccountsLinks
}
