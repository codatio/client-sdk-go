// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"time"
)

type GetBankingAccountRequest struct {
	// Unique identifier for an account
	AccountID    string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

// GetBankingAccountSourceModifiedDateAccountBalanceAmounts - An object containing bank balance data.
type GetBankingAccountSourceModifiedDateAccountBalanceAmounts struct {
	// The balance available in the account, including any pending transactions. This doesn't include additional funds available from any overdrafts.
	Available *float64 `json:"available,omitempty"`
	// The balance of the account only including cleared transactions.
	Current *float64 `json:"current,omitempty"`
	// The minimum allowed balance for the account. For example, a $100.00 overdraft would show as a limit of -100.00
	Limit *float64 `json:"limit,omitempty"`
}

// GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnum - Type of account
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

// GetBankingAccountSourceModifiedDateAccountIdentifiers - An object containing bank account identification information.
type GetBankingAccountSourceModifiedDateAccountIdentifiers struct {
	// The local (usually national) routing number for the account.
	//
	// This is known by different names in different countries:
	// * BSB code (Australia)
	// * routing number (Canada, USA)
	// * sort code (UK)
	BankCode *string `json:"bankCode,omitempty"`
	// The ISO 9362 code (commonly called SWIFT code, SWIFT-BIC or BIC) for the account.
	Bic *string `json:"bic,omitempty"`
	// The international bank account number (IBAN) for the account, if known.
	Iban *string `json:"iban,omitempty"`
	// A portion of the actual account `number` to help account identification where number is tokenised (Plaid only)
	MaskedAccountNumber *string `json:"maskedAccountNumber,omitempty"`
	// The account number for the account. When combined with the`bankCode`, this is usually enough to uniquely identify an account within a jurisdiction.
	Number *string `json:"number,omitempty"`
	// Detailed account category
	Subtype *string `json:"subtype,omitempty"`
	// Type of account
	Type GetBankingAccountSourceModifiedDateAccountIdentifiersTypeEnum `json:"type"`
}

// GetBankingAccountSourceModifiedDateAccountInstitution - The bank or other financial institution providing the account.
type GetBankingAccountSourceModifiedDateAccountInstitution struct {
	// The institution's ID, according to the provider.
	ID *string `json:"id,omitempty"`
	// The institution's name, according to the underlying provider.
	Name *string `json:"name,omitempty"`
}

// GetBankingAccountSourceModifiedDateTypeEnum - The type of transactions and balances on the account.
// For Credit accounts, positive balances are liabilities and positive transactions reduce liabilities.
// For Debit accounts, positive balances are assets and positive transactions increase assets.
type GetBankingAccountSourceModifiedDateTypeEnum string

const (
	GetBankingAccountSourceModifiedDateTypeEnumUnknown GetBankingAccountSourceModifiedDateTypeEnum = "Unknown"
	GetBankingAccountSourceModifiedDateTypeEnumCredit  GetBankingAccountSourceModifiedDateTypeEnum = "Credit"
	GetBankingAccountSourceModifiedDateTypeEnumDebit   GetBankingAccountSourceModifiedDateTypeEnum = "Debit"
)

// GetBankingAccountSourceModifiedDate - An account where payments are made or received, and bank transactions are recorded.
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-accounts).
type GetBankingAccountSourceModifiedDate struct {
	// An object containing bank balance data.
	Balance GetBankingAccountSourceModifiedDateAccountBalanceAmounts `json:"balance"`
	// The currency code for the account.
	Currency string `json:"currency"`
	// The name of the person or company who holds the account.
	Holder *string `json:"holder,omitempty"`
	// The ID of the account from the provider.
	ID string `json:"id"`
	// An object containing bank account identification information.
	Identifiers GetBankingAccountSourceModifiedDateAccountIdentifiers `json:"identifiers"`
	// The friendly name of the account, chosen by the holder. This may not have been set by the account holder and therefore is not always available.
	InformalName *string `json:"informalName,omitempty"`
	// The bank or other financial institution providing the account.
	Institution GetBankingAccountSourceModifiedDateAccountInstitution `json:"institution"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	// The name of the account according to the provider.
	Name string `json:"name"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
	// The type of transactions and balances on the account.
	// For Credit accounts, positive balances are liabilities and positive transactions reduce liabilities.
	// For Debit accounts, positive balances are assets and positive transactions increase assets.
	Type GetBankingAccountSourceModifiedDateTypeEnum `json:"type"`
}

type GetBankingAccountResponse struct {
	ContentType string
	// Success
	SourceModifiedDate *GetBankingAccountSourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
