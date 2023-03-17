package operations

import (
	"net/http"
	"time"
)

type GetBankingAccountRequest struct {
	AccountID    string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

// GetBankingAccountSourceModifiedDateAccountBalanceAmounts
// An object containing bank balance data.
type GetBankingAccountSourceModifiedDateAccountBalanceAmounts struct {
	Available *float64 `json:"available,omitempty"`
	Current   *float64 `json:"current,omitempty"`
	Limit     *float64 `json:"limit,omitempty"`
}

type GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnum string

const (
	GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnumAccount    GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnum = "Account"
	GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnumCard       GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnum = "Card"
	GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnumCredit     GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnum = "Credit"
	GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnumDepository GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnum = "Depository"
	GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnumInvestment GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnum = "Investment"
	GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnumLoan       GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnum = "Loan"
	GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnumOther      GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnum = "Other"
)

// GetBankingAccountSourceModifiedDateAccountIdentifiers
// An object containing bank account identification information.
type GetBankingAccountSourceModifiedDateAccountIdentifiers struct {
	BankCode            *string                                                       `json:"bankCode,omitempty"`
	Bic                 *string                                                       `json:"bic,omitempty"`
	Iban                *string                                                       `json:"iban,omitempty"`
	MaskedAccountNumber *string                                                       `json:"maskedAccountNumber,omitempty"`
	Number              *string                                                       `json:"number,omitempty"`
	Subtype             *string                                                       `json:"subtype,omitempty"`
	Type                GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnum `json:"type"`
}

// GetBankingAccountSourceModifiedDateAccountInstitution
// The bank or other financial institution providing the account.
type GetBankingAccountSourceModifiedDateAccountInstitution struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetBankingAccountSourceModifiedDateTypeEnum string

const (
	GetBankingAccountSourceModifiedDateTypeEnumUnknown GetBankingAccountSourceModifiedDateTypeEnum = "Unknown"
	GetBankingAccountSourceModifiedDateTypeEnumCredit  GetBankingAccountSourceModifiedDateTypeEnum = "Credit"
	GetBankingAccountSourceModifiedDateTypeEnumDebit   GetBankingAccountSourceModifiedDateTypeEnum = "Debit"
)

// GetBankingAccountSourceModifiedDate
// An account where payments are made or received, and bank transactions are recorded.
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-accounts).
type GetBankingAccountSourceModifiedDate struct {
	Balance            GetBankingAccountSourceModifiedDateAccountBalanceAmounts `json:"balance"`
	Currency           string                                                   `json:"currency"`
	Holder             *string                                                  `json:"holder,omitempty"`
	ID                 string                                                   `json:"id"`
	Identifiers        GetBankingAccountSourceModifiedDateAccountIdentifiers    `json:"identifiers"`
	InformalName       *string                                                  `json:"informalName,omitempty"`
	Institution        GetBankingAccountSourceModifiedDateAccountInstitution    `json:"institution"`
	ModifiedDate       *time.Time                                               `json:"modifiedDate,omitempty"`
	Name               string                                                   `json:"name"`
	SourceModifiedDate *time.Time                                               `json:"sourceModifiedDate,omitempty"`
	Type               GetBankingAccountSourceModifiedDateTypeEnum              `json:"type"`
}

type GetBankingAccountResponse struct {
	ContentType        string
	SourceModifiedDate *GetBankingAccountSourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
