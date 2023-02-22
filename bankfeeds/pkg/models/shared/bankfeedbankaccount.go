package shared

import (
	"time"
)

type BankFeedBankAccountAccountTypeEnum string

const (
	BankFeedBankAccountAccountTypeEnumUnknown BankFeedBankAccountAccountTypeEnum = "Unknown"
	BankFeedBankAccountAccountTypeEnumCredit  BankFeedBankAccountAccountTypeEnum = "Credit"
	BankFeedBankAccountAccountTypeEnumDebit   BankFeedBankAccountAccountTypeEnum = "Debit"
)

// BankFeedBankAccount
// The target bank account in a supported accounting package for ingestion into a bank feed.
type BankFeedBankAccount struct {
	AccountName   *string                             `json:"accountName,omitempty"`
	AccountNumber *string                             `json:"accountNumber,omitempty"`
	AccountType   *BankFeedBankAccountAccountTypeEnum `json:"accountType,omitempty"`
	Balance       *float64                            `json:"balance,omitempty"`
	Currency      *string                             `json:"currency,omitempty"`
	FeedStartDate *time.Time                          `json:"feedStartDate,omitempty"`
	ID            string                              `json:"id"`
	ModifiedDate  *time.Time                          `json:"modifiedDate,omitempty"`
	SortCode      *string                             `json:"sortCode,omitempty"`
	Status        *string                             `json:"status,omitempty"`
}
