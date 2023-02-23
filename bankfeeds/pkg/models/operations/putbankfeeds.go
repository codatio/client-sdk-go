package operations

import (
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
	"time"
)

type PutBankFeedsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PutBankFeedsBankFeedBankAccountAccountTypeEnum string

const (
	PutBankFeedsBankFeedBankAccountAccountTypeEnumUnknown PutBankFeedsBankFeedBankAccountAccountTypeEnum = "Unknown"
	PutBankFeedsBankFeedBankAccountAccountTypeEnumCredit  PutBankFeedsBankFeedBankAccountAccountTypeEnum = "Credit"
	PutBankFeedsBankFeedBankAccountAccountTypeEnumDebit   PutBankFeedsBankFeedBankAccountAccountTypeEnum = "Debit"
)

// PutBankFeedsBankFeedBankAccount
// The target bank account in a supported accounting package for ingestion into a bank feed.
type PutBankFeedsBankFeedBankAccount struct {
	AccountName   *string                                         `json:"accountName,omitempty"`
	AccountNumber *string                                         `json:"accountNumber,omitempty"`
	AccountType   *PutBankFeedsBankFeedBankAccountAccountTypeEnum `json:"accountType,omitempty"`
	Balance       *float64                                        `json:"balance,omitempty"`
	Currency      *string                                         `json:"currency,omitempty"`
	FeedStartDate *time.Time                                      `json:"feedStartDate,omitempty"`
	ID            string                                          `json:"id"`
	ModifiedDate  *time.Time                                      `json:"modifiedDate,omitempty"`
	SortCode      *string                                         `json:"sortCode,omitempty"`
	Status        *string                                         `json:"status,omitempty"`
}

type PutBankFeedsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PutBankFeedsRequest struct {
	PathParams PutBankFeedsPathParams
	Request    []PutBankFeedsBankFeedBankAccount `request:"mediaType=application/json"`
	Security   PutBankFeedsSecurity
}

type PutBankFeedsResponse struct {
	BankFeedBankAccounts []PutBankFeedsBankFeedBankAccount
	ContentType          string
	StatusCode           int
}
