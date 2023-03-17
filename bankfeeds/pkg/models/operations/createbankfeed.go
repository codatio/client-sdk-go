package operations

import (
	"net/http"
	"time"
)

type CreateBankFeedBankFeedBankAccountAccountTypeEnum string

const (
	CreateBankFeedBankFeedBankAccountAccountTypeEnumUnknown CreateBankFeedBankFeedBankAccountAccountTypeEnum = "Unknown"
	CreateBankFeedBankFeedBankAccountAccountTypeEnumCredit  CreateBankFeedBankFeedBankAccountAccountTypeEnum = "Credit"
	CreateBankFeedBankFeedBankAccountAccountTypeEnumDebit   CreateBankFeedBankFeedBankAccountAccountTypeEnum = "Debit"
)

// CreateBankFeedBankFeedBankAccount
// The target bank account in a supported accounting package for ingestion into a bank feed.
type CreateBankFeedBankFeedBankAccount struct {
	AccountName   *string                                           `json:"accountName,omitempty"`
	AccountNumber *string                                           `json:"accountNumber,omitempty"`
	AccountType   *CreateBankFeedBankFeedBankAccountAccountTypeEnum `json:"accountType,omitempty"`
	Balance       *float64                                          `json:"balance,omitempty"`
	Currency      *string                                           `json:"currency,omitempty"`
	FeedStartDate *time.Time                                        `json:"feedStartDate,omitempty"`
	ID            string                                            `json:"id"`
	ModifiedDate  *time.Time                                        `json:"modifiedDate,omitempty"`
	SortCode      *string                                           `json:"sortCode,omitempty"`
	Status        *string                                           `json:"status,omitempty"`
}

type CreateBankFeedRequest struct {
	RequestBody  []CreateBankFeedBankFeedBankAccount `request:"mediaType=application/json"`
	CompanyID    string                              `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string                              `pathParam:"style=simple,explode=false,name=connectionId"`
}

type CreateBankFeedResponse struct {
	BankFeedBankAccounts []CreateBankFeedBankFeedBankAccount
	ContentType          string
	StatusCode           int
	RawResponse          *http.Response
}
