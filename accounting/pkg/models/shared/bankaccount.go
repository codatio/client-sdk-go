// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// BankAccountBankAccountTypeEnum - The type of transactions and balances on the account.
// For Credit accounts, positive balances are liabilities, and positive transactions **reduce** liabilities.
// For Debit accounts, positive balances are assets, and positive transactions **increase** assets.
type BankAccountBankAccountTypeEnum string

const (
	BankAccountBankAccountTypeEnumUnknown BankAccountBankAccountTypeEnum = "Unknown"
	BankAccountBankAccountTypeEnumCredit  BankAccountBankAccountTypeEnum = "Credit"
	BankAccountBankAccountTypeEnumDebit   BankAccountBankAccountTypeEnum = "Debit"
)

func (e *BankAccountBankAccountTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Credit":
		fallthrough
	case "Debit":
		*e = BankAccountBankAccountTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for BankAccountBankAccountTypeEnum: %s", s)
	}
}

// BankAccount - > **Accessing Bank Accounts through Banking API**
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
type BankAccount struct {
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
	AccountType *BankAccountBankAccountTypeEnum `json:"accountType,omitempty"`
	// Total available balance of the bank account as reported by the underlying data source. This may take into account overdrafts or pending transactions for example.
	AvailableBalance *float64 `json:"availableBalance,omitempty"`
	// Balance of the bank account.
	Balance *float64 `json:"balance,omitempty"`
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
	Institution *string   `json:"institution,omitempty"`
	Metadata    *Metadata `json:"metadata,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
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
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
}