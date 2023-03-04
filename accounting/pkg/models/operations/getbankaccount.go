package operations

import (
	"net/http"
	"time"
)

type GetBankAccountPathParams struct {
	AccountID    string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetBankAccountRequest struct {
	PathParams GetBankAccountPathParams
}

type GetBankAccountSourceModifiedDateAccountTypeEnum string

const (
	GetBankAccountSourceModifiedDateAccountTypeEnumUnknown GetBankAccountSourceModifiedDateAccountTypeEnum = "Unknown"
	GetBankAccountSourceModifiedDateAccountTypeEnumCredit  GetBankAccountSourceModifiedDateAccountTypeEnum = "Credit"
	GetBankAccountSourceModifiedDateAccountTypeEnumDebit   GetBankAccountSourceModifiedDateAccountTypeEnum = "Debit"
)

type GetBankAccountSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// GetBankAccountSourceModifiedDate
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
type GetBankAccountSourceModifiedDate struct {
	AccountName        *string                                          `json:"accountName,omitempty"`
	AccountNumber      *string                                          `json:"accountNumber,omitempty"`
	AccountType        *GetBankAccountSourceModifiedDateAccountTypeEnum `json:"accountType,omitempty"`
	AvailableBalance   *float64                                         `json:"availableBalance,omitempty"`
	Balance            *float64                                         `json:"balance,omitempty"`
	Currency           *string                                          `json:"currency,omitempty"`
	IBan               *string                                          `json:"iBan,omitempty"`
	ID                 *string                                          `json:"id,omitempty"`
	Institution        *string                                          `json:"institution,omitempty"`
	Metadata           *GetBankAccountSourceModifiedDateMetadata        `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                       `json:"modifiedDate,omitempty"`
	NominalCode        *string                                          `json:"nominalCode,omitempty"`
	OverdraftLimit     *float64                                         `json:"overdraftLimit,omitempty"`
	SortCode           *string                                          `json:"sortCode,omitempty"`
	SourceModifiedDate *time.Time                                       `json:"sourceModifiedDate,omitempty"`
}

type GetBankAccountResponse struct {
	ContentType        string
	SourceModifiedDate *GetBankAccountSourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
