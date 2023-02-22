package shared

import (
	"time"
)

// AccountAccountBalanceAmounts
// An object containing bank balance data.
type AccountAccountBalanceAmounts struct {
	Available *float64 `json:"available,omitempty"`
	Current   *float64 `json:"current,omitempty"`
	Limit     *float64 `json:"limit,omitempty"`
}

type AccountAccountIdentifiersTypeEnum string

const (
	AccountAccountIdentifiersTypeEnumAccount    AccountAccountIdentifiersTypeEnum = "Account"
	AccountAccountIdentifiersTypeEnumCard       AccountAccountIdentifiersTypeEnum = "Card"
	AccountAccountIdentifiersTypeEnumCredit     AccountAccountIdentifiersTypeEnum = "Credit"
	AccountAccountIdentifiersTypeEnumDepository AccountAccountIdentifiersTypeEnum = "Depository"
	AccountAccountIdentifiersTypeEnumInvestment AccountAccountIdentifiersTypeEnum = "Investment"
	AccountAccountIdentifiersTypeEnumLoan       AccountAccountIdentifiersTypeEnum = "Loan"
	AccountAccountIdentifiersTypeEnumOther      AccountAccountIdentifiersTypeEnum = "Other"
)

// AccountAccountIdentifiers
// An object containing bank account identification information.
type AccountAccountIdentifiers struct {
	BankCode            *string                           `json:"bankCode,omitempty"`
	Bic                 *string                           `json:"bic,omitempty"`
	Iban                *string                           `json:"iban,omitempty"`
	MaskedAccountNumber *string                           `json:"maskedAccountNumber,omitempty"`
	Number              *string                           `json:"number,omitempty"`
	Subtype             *string                           `json:"subtype,omitempty"`
	Type                AccountAccountIdentifiersTypeEnum `json:"type"`
}

// AccountAccountInstitution
// The bank or other financial institution providing the account.
type AccountAccountInstitution struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type AccountTypeEnum string

const (
	AccountTypeEnumUnknown AccountTypeEnum = "Unknown"
	AccountTypeEnumCredit  AccountTypeEnum = "Credit"
	AccountTypeEnumDebit   AccountTypeEnum = "Debit"
)

// Account
// An account where payments are made or received, and bank transactions are recorded.
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-accounts).
type Account struct {
	Balance            *AccountAccountBalanceAmounts `json:"balance,omitempty"`
	Currency           *string                       `json:"currency,omitempty"`
	Holder             *string                       `json:"holder,omitempty"`
	ID                 *string                       `json:"id,omitempty"`
	Identifiers        *AccountAccountIdentifiers    `json:"identifiers,omitempty"`
	InformalName       *string                       `json:"informalName,omitempty"`
	Institution        *AccountAccountInstitution    `json:"institution,omitempty"`
	ModifiedDate       *time.Time                    `json:"modifiedDate,omitempty"`
	Name               *string                       `json:"name,omitempty"`
	SourceModifiedDate *time.Time                    `json:"sourceModifiedDate,omitempty"`
	Type               *AccountTypeEnum              `json:"type,omitempty"`
}
