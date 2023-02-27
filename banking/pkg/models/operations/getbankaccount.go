package operations

import (
	"github.com/codatio/client-sdk-go/banking/pkg/models/shared"
	"time"
)

type GetBankAccountPathParams struct {
	AccountID    string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetBankAccountSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetBankAccountRequest struct {
	PathParams GetBankAccountPathParams
	Security   GetBankAccountSecurity
}

// GetBankAccountSourceModifiedDateAccountBalanceAmounts
// An object containing bank balance data.
type GetBankAccountSourceModifiedDateAccountBalanceAmounts struct {
	Available *float64 `json:"available,omitempty"`
	Current   *float64 `json:"current,omitempty"`
	Limit     *float64 `json:"limit,omitempty"`
}

type GetBankAccountSourceModifiedDateAccountIdentifiersTypeEnum string

const (
	GetBankAccountSourceModifiedDateAccountIdentifiersTypeEnumAccount    GetBankAccountSourceModifiedDateAccountIdentifiersTypeEnum = "Account"
	GetBankAccountSourceModifiedDateAccountIdentifiersTypeEnumCard       GetBankAccountSourceModifiedDateAccountIdentifiersTypeEnum = "Card"
	GetBankAccountSourceModifiedDateAccountIdentifiersTypeEnumCredit     GetBankAccountSourceModifiedDateAccountIdentifiersTypeEnum = "Credit"
	GetBankAccountSourceModifiedDateAccountIdentifiersTypeEnumDepository GetBankAccountSourceModifiedDateAccountIdentifiersTypeEnum = "Depository"
	GetBankAccountSourceModifiedDateAccountIdentifiersTypeEnumInvestment GetBankAccountSourceModifiedDateAccountIdentifiersTypeEnum = "Investment"
	GetBankAccountSourceModifiedDateAccountIdentifiersTypeEnumLoan       GetBankAccountSourceModifiedDateAccountIdentifiersTypeEnum = "Loan"
	GetBankAccountSourceModifiedDateAccountIdentifiersTypeEnumOther      GetBankAccountSourceModifiedDateAccountIdentifiersTypeEnum = "Other"
)

// GetBankAccountSourceModifiedDateAccountIdentifiers
// An object containing bank account identification information.
type GetBankAccountSourceModifiedDateAccountIdentifiers struct {
	BankCode            *string                                                    `json:"bankCode,omitempty"`
	Bic                 *string                                                    `json:"bic,omitempty"`
	Iban                *string                                                    `json:"iban,omitempty"`
	MaskedAccountNumber *string                                                    `json:"maskedAccountNumber,omitempty"`
	Number              *string                                                    `json:"number,omitempty"`
	Subtype             *string                                                    `json:"subtype,omitempty"`
	Type                GetBankAccountSourceModifiedDateAccountIdentifiersTypeEnum `json:"type"`
}

// GetBankAccountSourceModifiedDateAccountInstitution
// The bank or other financial institution providing the account.
type GetBankAccountSourceModifiedDateAccountInstitution struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetBankAccountSourceModifiedDateTypeEnum string

const (
	GetBankAccountSourceModifiedDateTypeEnumUnknown GetBankAccountSourceModifiedDateTypeEnum = "Unknown"
	GetBankAccountSourceModifiedDateTypeEnumCredit  GetBankAccountSourceModifiedDateTypeEnum = "Credit"
	GetBankAccountSourceModifiedDateTypeEnumDebit   GetBankAccountSourceModifiedDateTypeEnum = "Debit"
)

// GetBankAccountSourceModifiedDate
// An account where payments are made or received, and bank transactions are recorded.
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-accounts).
type GetBankAccountSourceModifiedDate struct {
	Balance            GetBankAccountSourceModifiedDateAccountBalanceAmounts `json:"balance"`
	Currency           string                                                `json:"currency"`
	Holder             *string                                               `json:"holder,omitempty"`
	ID                 string                                                `json:"id"`
	Identifiers        GetBankAccountSourceModifiedDateAccountIdentifiers    `json:"identifiers"`
	InformalName       *string                                               `json:"informalName,omitempty"`
	Institution        GetBankAccountSourceModifiedDateAccountInstitution    `json:"institution"`
	ModifiedDate       *time.Time                                            `json:"modifiedDate,omitempty"`
	Name               string                                                `json:"name"`
	SourceModifiedDate *time.Time                                            `json:"sourceModifiedDate,omitempty"`
	Type               GetBankAccountSourceModifiedDateTypeEnum              `json:"type"`
}

type GetBankAccountResponse struct {
	ContentType        string
	SourceModifiedDate *GetBankAccountSourceModifiedDate
	StatusCode         int
}
