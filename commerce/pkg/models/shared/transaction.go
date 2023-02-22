package shared

import (
	"time"
)

type TransactionTypeEnum string

const (
	TransactionTypeEnumPayment          TransactionTypeEnum = "Payment"
	TransactionTypeEnumRefund           TransactionTypeEnum = "Refund"
	TransactionTypeEnumPayout           TransactionTypeEnum = "Payout"
	TransactionTypeEnumFailedPayout     TransactionTypeEnum = "FailedPayout"
	TransactionTypeEnumTransfer         TransactionTypeEnum = "Transfer"
	TransactionTypeEnumPaymentFee       TransactionTypeEnum = "PaymentFee"
	TransactionTypeEnumPaymentFeeRefund TransactionTypeEnum = "PaymentFeeRefund"
	TransactionTypeEnumUnknown          TransactionTypeEnum = "Unknown"
)

// Transaction
// A financial transaction recorded in the commerce or point of sale system
type Transaction struct {
	CreatedDate          *time.Time           `json:"createdDate,omitempty"`
	Currency             *string              `json:"currency,omitempty"`
	ID                   string               `json:"id"`
	ModifiedDate         *time.Time           `json:"modifiedDate,omitempty"`
	SourceCreatedDate    *time.Time           `json:"sourceCreatedDate,omitempty"`
	SourceModifiedDate   *time.Time           `json:"sourceModifiedDate,omitempty"`
	SubType              *string              `json:"subType,omitempty"`
	TotalAmount          *float64             `json:"totalAmount,omitempty"`
	TransactionSourceRef *One                 `json:"transactionSourceRef,omitempty"`
	Type                 *TransactionTypeEnum `json:"type,omitempty"`
}
