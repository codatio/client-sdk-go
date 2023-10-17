// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// AccountingBankAccountBankAccountType - The type of transactions and balances on the account.
// For Credit accounts, positive balances are liabilities, and positive transactions **reduce** liabilities.
// For Debit accounts, positive balances are assets, and positive transactions **increase** assets.
type AccountingBankAccountBankAccountType string

const (
	AccountingBankAccountBankAccountTypeUnknown AccountingBankAccountBankAccountType = "Unknown"
	AccountingBankAccountBankAccountTypeCredit  AccountingBankAccountBankAccountType = "Credit"
	AccountingBankAccountBankAccountTypeDebit   AccountingBankAccountBankAccountType = "Debit"
)

func (e AccountingBankAccountBankAccountType) ToPointer() *AccountingBankAccountBankAccountType {
	return &e
}

func (e *AccountingBankAccountBankAccountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Credit":
		fallthrough
	case "Debit":
		*e = AccountingBankAccountBankAccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountingBankAccountBankAccountType: %v", v)
	}
}

// AccountingBankAccount - > **Accessing Bank Accounts through Banking API**
// >
// > This datatype was originally used for accessing bank account data both in accounting integrations and open banking aggregators.
// >
// > To view bank account data through the Banking API, please refer to the new datatype [here](https://docs.codat.io/banking-api#/schemas/Account)
//
// > View the coverage for bank accounts in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankAccounts" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A list of bank accounts associated with a company and a specific data connection.
//
// Bank accounts data includes:
// * The name and ID of the account in the accounting platform.
// * The currency and balance of the account.
// * The sort code and account number.
type AccountingBankAccount struct {
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
	// The type of transactions and balances on the account.
	// For Credit accounts, positive balances are liabilities, and positive transactions **reduce** liabilities.
	// For Debit accounts, positive balances are assets, and positive transactions **increase** assets.
	AccountType *AccountingBankAccountBankAccountType `json:"accountType,omitempty"`
	// Total available balance of the bank account as reported by the underlying data source. This may take into account overdrafts or pending transactions for example.
	AvailableBalance *decimal.Big `decimal:"number" json:"availableBalance,omitempty"`
	// Balance of the bank account.
	Balance *decimal.Big `decimal:"number" json:"balance,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// International bank account number of the account. Often used when making or receiving international payments.
	IBan *string `json:"iBan,omitempty"`
	// Identifier for the account, unique for the company in the accounting platform.
	ID *string `json:"id,omitempty"`
	// The institution of the bank account.
	Institution  *string   `json:"institution,omitempty"`
	Metadata     *Metadata `json:"metadata,omitempty"`
	ModifiedDate *string   `json:"modifiedDate,omitempty"`
	// Code used to identify each nominal account for a business.
	NominalCode *string `json:"nominalCode,omitempty"`
	// Pre-arranged overdraft limit of the account.
	//
	// The value is always positive. For example, an overdraftLimit of `1000` means that the balance of the account can go down to `-1000`.
	OverdraftLimit *decimal.Big `decimal:"number" json:"overdraftLimit,omitempty"`
	// Sort code for the bank account.
	//
	// Xero integrations
	// The sort code is only displayed when the currency = GBP and the sort code and account number sum to 14 digits. For non-GBP accounts, this field is not populated.
	SortCode           *string `json:"sortCode,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
}

func (a AccountingBankAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingBankAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingBankAccount) GetAccountName() *string {
	if o == nil {
		return nil
	}
	return o.AccountName
}

func (o *AccountingBankAccount) GetAccountNumber() *string {
	if o == nil {
		return nil
	}
	return o.AccountNumber
}

func (o *AccountingBankAccount) GetAccountType() *AccountingBankAccountBankAccountType {
	if o == nil {
		return nil
	}
	return o.AccountType
}

func (o *AccountingBankAccount) GetAvailableBalance() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.AvailableBalance
}

func (o *AccountingBankAccount) GetBalance() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *AccountingBankAccount) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountingBankAccount) GetIBan() *string {
	if o == nil {
		return nil
	}
	return o.IBan
}

func (o *AccountingBankAccount) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingBankAccount) GetInstitution() *string {
	if o == nil {
		return nil
	}
	return o.Institution
}

func (o *AccountingBankAccount) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AccountingBankAccount) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *AccountingBankAccount) GetNominalCode() *string {
	if o == nil {
		return nil
	}
	return o.NominalCode
}

func (o *AccountingBankAccount) GetOverdraftLimit() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.OverdraftLimit
}

func (o *AccountingBankAccount) GetSortCode() *string {
	if o == nil {
		return nil
	}
	return o.SortCode
}

func (o *AccountingBankAccount) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *AccountingBankAccount) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}
