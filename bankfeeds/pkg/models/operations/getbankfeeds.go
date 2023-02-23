package operations

import (
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
	"time"
)

type GetBankFeedsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetBankFeedsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetBankFeedsRequest struct {
	PathParams GetBankFeedsPathParams
	Security   GetBankFeedsSecurity
}

type GetBankFeedsBankFeedBankAccountAccountTypeEnum string

const (
	GetBankFeedsBankFeedBankAccountAccountTypeEnumUnknown GetBankFeedsBankFeedBankAccountAccountTypeEnum = "Unknown"
	GetBankFeedsBankFeedBankAccountAccountTypeEnumCredit  GetBankFeedsBankFeedBankAccountAccountTypeEnum = "Credit"
	GetBankFeedsBankFeedBankAccountAccountTypeEnumDebit   GetBankFeedsBankFeedBankAccountAccountTypeEnum = "Debit"
)

// GetBankFeedsBankFeedBankAccount
// The target bank account in a supported accounting package for ingestion into a bank feed.
type GetBankFeedsBankFeedBankAccount struct {
	AccountName   *string                                         `json:"accountName,omitempty"`
	AccountNumber *string                                         `json:"accountNumber,omitempty"`
	AccountType   *GetBankFeedsBankFeedBankAccountAccountTypeEnum `json:"accountType,omitempty"`
	Balance       *float64                                        `json:"balance,omitempty"`
	Currency      *string                                         `json:"currency,omitempty"`
	FeedStartDate *time.Time                                      `json:"feedStartDate,omitempty"`
	ID            string                                          `json:"id"`
	ModifiedDate  *time.Time                                      `json:"modifiedDate,omitempty"`
	SortCode      *string                                         `json:"sortCode,omitempty"`
	Status        *string                                         `json:"status,omitempty"`
}

type GetBankFeedsResponse struct {
	BankFeedBankAccounts []GetBankFeedsBankFeedBankAccount
	ContentType          string
	StatusCode           int
}
