// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// TransactionTypeEnum - The type of the platform transaction:
// - `Unknown`
// - `FailedPayout` — Failed transfer of funds from the seller's merchant account to their bank account.
// - `Payment` — Credit and debit card payments.
// - `PaymentFee` — Payment provider's fee on each card payment.
// - `PaymentFeeRefund` — Payment provider's fee that has been refunded to the seller.
// - `Payout` — Transfer of funds from the seller's merchant account to their bank account.
// - `Refund` — Refunds to a customer's credit or debit card.
// - `Transfer` — Secure transfer of funds to the seller's bank account.
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

func (e TransactionTypeEnum) ToPointer() *TransactionTypeEnum {
	return &e
}

func (e *TransactionTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
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
		*e = TransactionTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for TransactionTypeEnum: %s", s)
	}
}
