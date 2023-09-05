// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// TransactionType - The type of the platform transaction:
// - `Unknown`
// - `FailedPayout` — Failed transfer of funds from the seller's merchant account to their bank account.
// - `Payment` — Credit and debit card payments.
// - `PaymentFee` — Payment provider's fee on each card payment.
// - `PaymentFeeRefund` — Payment provider's fee that has been refunded to the seller.
// - `Payout` — Transfer of funds from the seller's merchant account to their bank account.
// - `Refund` — Refunds to a customer's credit or debit card.
// - `Transfer` — Secure transfer of funds to the seller's bank account.
type TransactionType string

const (
	TransactionTypePayment          TransactionType = "Payment"
	TransactionTypeRefund           TransactionType = "Refund"
	TransactionTypePayout           TransactionType = "Payout"
	TransactionTypeFailedPayout     TransactionType = "FailedPayout"
	TransactionTypeTransfer         TransactionType = "Transfer"
	TransactionTypePaymentFee       TransactionType = "PaymentFee"
	TransactionTypePaymentFeeRefund TransactionType = "PaymentFeeRefund"
	TransactionTypeUnknown          TransactionType = "Unknown"
)

func (e TransactionType) ToPointer() *TransactionType {
	return &e
}

func (e *TransactionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Payment":
		fallthrough
	case "Refund":
		fallthrough
	case "Payout":
		fallthrough
	case "FailedPayout":
		fallthrough
	case "Transfer":
		fallthrough
	case "PaymentFee":
		fallthrough
	case "PaymentFeeRefund":
		fallthrough
	case "Unknown":
		*e = TransactionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TransactionType: %v", v)
	}
}