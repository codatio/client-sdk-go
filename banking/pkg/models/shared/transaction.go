package shared

import (
	"time"
)

type TransactionCodeEnum string

const (
	TransactionCodeEnumUnknown       TransactionCodeEnum = "Unknown"
	TransactionCodeEnumFee           TransactionCodeEnum = "Fee"
	TransactionCodeEnumPayment       TransactionCodeEnum = "Payment"
	TransactionCodeEnumCash          TransactionCodeEnum = "Cash"
	TransactionCodeEnumTransfer      TransactionCodeEnum = "Transfer"
	TransactionCodeEnumInterest      TransactionCodeEnum = "Interest"
	TransactionCodeEnumCashback      TransactionCodeEnum = "Cashback"
	TransactionCodeEnumCheque        TransactionCodeEnum = "Cheque"
	TransactionCodeEnumDirectDebit   TransactionCodeEnum = "DirectDebit"
	TransactionCodeEnumPurchase      TransactionCodeEnum = "Purchase"
	TransactionCodeEnumStandingOrder TransactionCodeEnum = "StandingOrder"
	TransactionCodeEnumAdjustment    TransactionCodeEnum = "Adjustment"
	TransactionCodeEnumCredit        TransactionCodeEnum = "Credit"
	TransactionCodeEnumOther         TransactionCodeEnum = "Other"
	TransactionCodeEnumNotSupported  TransactionCodeEnum = "NotSupported"
)

// TransactionTransactionCategoryRef
// An object of bank transaction category reference data.
type TransactionTransactionCategoryRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// Transaction
// The Banking Transactions data type provides an immutable source of up-to-date information on income and expenditure.
//
// View the coverage for banking transactions in the [Data Coverage Explorer](https://knowledge.codat.io/supported-features/banking?view=tab-by-data-type&dataType=banking-transactions).
type Transaction struct {
	AccountID              string                             `json:"accountId"`
	Amount                 *float64                           `json:"amount,omitempty"`
	AuthorizedDate         *time.Time                         `json:"authorizedDate,omitempty"`
	Code                   *TransactionCodeEnum               `json:"code,omitempty"`
	Currency               string                             `json:"currency"`
	Description            *string                            `json:"description,omitempty"`
	ID                     string                             `json:"id"`
	MerchantName           *string                            `json:"merchantName,omitempty"`
	ModifiedDate           *time.Time                         `json:"modifiedDate,omitempty"`
	PostedDate             *time.Time                         `json:"postedDate,omitempty"`
	SourceModifiedDate     *time.Time                         `json:"sourceModifiedDate,omitempty"`
	TransactionCategoryRef *TransactionTransactionCategoryRef `json:"transactionCategoryRef,omitempty"`
}
