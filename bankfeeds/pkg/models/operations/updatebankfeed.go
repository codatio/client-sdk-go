package operations

import (
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
	"time"
)

type UpdateBankFeedPathParams struct {
	BankAccountID string `pathParam:"style=simple,explode=false,name=bankAccountId"`
	CompanyID     string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID  string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type UpdateBankFeedBankFeedBankAccountAccountTypeEnum string

const (
	UpdateBankFeedBankFeedBankAccountAccountTypeEnumUnknown UpdateBankFeedBankFeedBankAccountAccountTypeEnum = "Unknown"
	UpdateBankFeedBankFeedBankAccountAccountTypeEnumCredit  UpdateBankFeedBankFeedBankAccountAccountTypeEnum = "Credit"
	UpdateBankFeedBankFeedBankAccountAccountTypeEnumDebit   UpdateBankFeedBankFeedBankAccountAccountTypeEnum = "Debit"
)

// UpdateBankFeedBankFeedBankAccount
// The target bank account in a supported accounting package for ingestion into a bank feed.
type UpdateBankFeedBankFeedBankAccount struct {
	AccountName   *string                                           `json:"accountName,omitempty"`
	AccountNumber *string                                           `json:"accountNumber,omitempty"`
	AccountType   *UpdateBankFeedBankFeedBankAccountAccountTypeEnum `json:"accountType,omitempty"`
	Balance       *float64                                          `json:"balance,omitempty"`
	Currency      *string                                           `json:"currency,omitempty"`
	FeedStartDate *time.Time                                        `json:"feedStartDate,omitempty"`
	ID            string                                            `json:"id"`
	ModifiedDate  *time.Time                                        `json:"modifiedDate,omitempty"`
	SortCode      *string                                           `json:"sortCode,omitempty"`
	Status        *string                                           `json:"status,omitempty"`
}

type UpdateBankFeedSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateBankFeedRequest struct {
	PathParams UpdateBankFeedPathParams
	Request    *UpdateBankFeedBankFeedBankAccount `request:"mediaType=application/json"`
	Security   UpdateBankFeedSecurity
}

type UpdateBankFeedResponse struct {
	BankFeedBankAccount *UpdateBankFeedBankFeedBankAccount
	ContentType         string
	StatusCode          int
}
