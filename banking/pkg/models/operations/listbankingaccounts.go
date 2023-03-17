package operations

import (
	"net/http"
	"time"
)

type ListBankingAccountsRequest struct {
	CompanyID    string  `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string  `pathParam:"style=simple,explode=false,name=connectionId"`
	OrderBy      *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page         int     `queryParam:"style=form,explode=true,name=page"`
	PageSize     *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query        *string `queryParam:"style=form,explode=true,name=query"`
}

type ListBankingAccountsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListBankingAccountsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListBankingAccountsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListBankingAccountsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListBankingAccountsLinksLinks struct {
	Current  ListBankingAccountsLinksLinksCurrent   `json:"current"`
	Next     *ListBankingAccountsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListBankingAccountsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListBankingAccountsLinksLinksSelf      `json:"self"`
}

// ListBankingAccountsLinksSourceModifiedDateAccountBalanceAmounts
// An object containing bank balance data.
type ListBankingAccountsLinksSourceModifiedDateAccountBalanceAmounts struct {
	Available *float64 `json:"available,omitempty"`
	Current   *float64 `json:"current,omitempty"`
	Limit     *float64 `json:"limit,omitempty"`
}

type ListBankingAccountsLinksSourceModifiedDateAccountIdentifiersTypeEnum string

const (
	ListBankingAccountsLinksSourceModifiedDateAccountIdentifiersTypeEnumAccount    ListBankingAccountsLinksSourceModifiedDateAccountIdentifiersTypeEnum = "Account"
	ListBankingAccountsLinksSourceModifiedDateAccountIdentifiersTypeEnumCard       ListBankingAccountsLinksSourceModifiedDateAccountIdentifiersTypeEnum = "Card"
	ListBankingAccountsLinksSourceModifiedDateAccountIdentifiersTypeEnumCredit     ListBankingAccountsLinksSourceModifiedDateAccountIdentifiersTypeEnum = "Credit"
	ListBankingAccountsLinksSourceModifiedDateAccountIdentifiersTypeEnumDepository ListBankingAccountsLinksSourceModifiedDateAccountIdentifiersTypeEnum = "Depository"
	ListBankingAccountsLinksSourceModifiedDateAccountIdentifiersTypeEnumInvestment ListBankingAccountsLinksSourceModifiedDateAccountIdentifiersTypeEnum = "Investment"
	ListBankingAccountsLinksSourceModifiedDateAccountIdentifiersTypeEnumLoan       ListBankingAccountsLinksSourceModifiedDateAccountIdentifiersTypeEnum = "Loan"
	ListBankingAccountsLinksSourceModifiedDateAccountIdentifiersTypeEnumOther      ListBankingAccountsLinksSourceModifiedDateAccountIdentifiersTypeEnum = "Other"
)

// ListBankingAccountsLinksSourceModifiedDateAccountIdentifiers
// An object containing bank account identification information.
type ListBankingAccountsLinksSourceModifiedDateAccountIdentifiers struct {
	BankCode            *string                                                              `json:"bankCode,omitempty"`
	Bic                 *string                                                              `json:"bic,omitempty"`
	Iban                *string                                                              `json:"iban,omitempty"`
	MaskedAccountNumber *string                                                              `json:"maskedAccountNumber,omitempty"`
	Number              *string                                                              `json:"number,omitempty"`
	Subtype             *string                                                              `json:"subtype,omitempty"`
	Type                ListBankingAccountsLinksSourceModifiedDateAccountIdentifiersTypeEnum `json:"type"`
}

// ListBankingAccountsLinksSourceModifiedDateAccountInstitution
// The bank or other financial institution providing the account.
type ListBankingAccountsLinksSourceModifiedDateAccountInstitution struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ListBankingAccountsLinksSourceModifiedDateTypeEnum string

const (
	ListBankingAccountsLinksSourceModifiedDateTypeEnumUnknown ListBankingAccountsLinksSourceModifiedDateTypeEnum = "Unknown"
	ListBankingAccountsLinksSourceModifiedDateTypeEnumCredit  ListBankingAccountsLinksSourceModifiedDateTypeEnum = "Credit"
	ListBankingAccountsLinksSourceModifiedDateTypeEnumDebit   ListBankingAccountsLinksSourceModifiedDateTypeEnum = "Debit"
)

// ListBankingAccountsLinksSourceModifiedDate
// An account where payments are made or received, and bank transactions are recorded.
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-accounts).
type ListBankingAccountsLinksSourceModifiedDate struct {
	Balance            ListBankingAccountsLinksSourceModifiedDateAccountBalanceAmounts `json:"balance"`
	Currency           string                                                          `json:"currency"`
	Holder             *string                                                         `json:"holder,omitempty"`
	ID                 string                                                          `json:"id"`
	Identifiers        ListBankingAccountsLinksSourceModifiedDateAccountIdentifiers    `json:"identifiers"`
	InformalName       *string                                                         `json:"informalName,omitempty"`
	Institution        ListBankingAccountsLinksSourceModifiedDateAccountInstitution    `json:"institution"`
	ModifiedDate       *time.Time                                                      `json:"modifiedDate,omitempty"`
	Name               string                                                          `json:"name"`
	SourceModifiedDate *time.Time                                                      `json:"sourceModifiedDate,omitempty"`
	Type               ListBankingAccountsLinksSourceModifiedDateTypeEnum              `json:"type"`
}

// ListBankingAccountsLinks
// Codat's Paging Model
type ListBankingAccountsLinks struct {
	Links        ListBankingAccountsLinksLinks               `json:"_links"`
	PageNumber   int64                                       `json:"pageNumber"`
	PageSize     int64                                       `json:"pageSize"`
	Results      *ListBankingAccountsLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                       `json:"totalResults"`
}

type ListBankingAccountsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListBankingAccountsLinks
}
