package operations

import (
	"net/http"
	"time"
)

type GetBankingTransactionPathParams struct {
	CompanyID     string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID  string `pathParam:"style=simple,explode=false,name=connectionId"`
	TransactionID string `pathParam:"style=simple,explode=false,name=transactionId"`
}

type GetBankingTransactionRequest struct {
	PathParams GetBankingTransactionPathParams
}

type GetBankingTransactionSourceModifiedDateCodeEnum string

const (
	GetBankingTransactionSourceModifiedDateCodeEnumUnknown       GetBankingTransactionSourceModifiedDateCodeEnum = "Unknown"
	GetBankingTransactionSourceModifiedDateCodeEnumFee           GetBankingTransactionSourceModifiedDateCodeEnum = "Fee"
	GetBankingTransactionSourceModifiedDateCodeEnumPayment       GetBankingTransactionSourceModifiedDateCodeEnum = "Payment"
	GetBankingTransactionSourceModifiedDateCodeEnumCash          GetBankingTransactionSourceModifiedDateCodeEnum = "Cash"
	GetBankingTransactionSourceModifiedDateCodeEnumTransfer      GetBankingTransactionSourceModifiedDateCodeEnum = "Transfer"
	GetBankingTransactionSourceModifiedDateCodeEnumInterest      GetBankingTransactionSourceModifiedDateCodeEnum = "Interest"
	GetBankingTransactionSourceModifiedDateCodeEnumCashback      GetBankingTransactionSourceModifiedDateCodeEnum = "Cashback"
	GetBankingTransactionSourceModifiedDateCodeEnumCheque        GetBankingTransactionSourceModifiedDateCodeEnum = "Cheque"
	GetBankingTransactionSourceModifiedDateCodeEnumDirectDebit   GetBankingTransactionSourceModifiedDateCodeEnum = "DirectDebit"
	GetBankingTransactionSourceModifiedDateCodeEnumPurchase      GetBankingTransactionSourceModifiedDateCodeEnum = "Purchase"
	GetBankingTransactionSourceModifiedDateCodeEnumStandingOrder GetBankingTransactionSourceModifiedDateCodeEnum = "StandingOrder"
	GetBankingTransactionSourceModifiedDateCodeEnumAdjustment    GetBankingTransactionSourceModifiedDateCodeEnum = "Adjustment"
	GetBankingTransactionSourceModifiedDateCodeEnumCredit        GetBankingTransactionSourceModifiedDateCodeEnum = "Credit"
	GetBankingTransactionSourceModifiedDateCodeEnumOther         GetBankingTransactionSourceModifiedDateCodeEnum = "Other"
	GetBankingTransactionSourceModifiedDateCodeEnumNotSupported  GetBankingTransactionSourceModifiedDateCodeEnum = "NotSupported"
)

// GetBankingTransactionSourceModifiedDateTransactionCategoryRef
// An object of bank transaction category reference data.
type GetBankingTransactionSourceModifiedDateTransactionCategoryRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// GetBankingTransactionSourceModifiedDate
// The Banking Transactions data type provides an immutable source of up-to-date information on income and expenditure.
//
// View the coverage for banking transactions in the [Data Coverage Explorer](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-transactions).
type GetBankingTransactionSourceModifiedDate struct {
	AccountID              string                                                         `json:"accountId"`
	Amount                 *float64                                                       `json:"amount,omitempty"`
	AuthorizedDate         *time.Time                                                     `json:"authorizedDate,omitempty"`
	Code                   *GetBankingTransactionSourceModifiedDateCodeEnum               `json:"code,omitempty"`
	Currency               string                                                         `json:"currency"`
	Description            *string                                                        `json:"description,omitempty"`
	ID                     string                                                         `json:"id"`
	MerchantName           *string                                                        `json:"merchantName,omitempty"`
	ModifiedDate           *time.Time                                                     `json:"modifiedDate,omitempty"`
	PostedDate             *time.Time                                                     `json:"postedDate,omitempty"`
	SourceModifiedDate     *time.Time                                                     `json:"sourceModifiedDate,omitempty"`
	TransactionCategoryRef *GetBankingTransactionSourceModifiedDateTransactionCategoryRef `json:"transactionCategoryRef,omitempty"`
}

type GetBankingTransactionResponse struct {
	ContentType        string
	SourceModifiedDate *GetBankingTransactionSourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
