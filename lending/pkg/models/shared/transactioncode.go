// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// TransactionCode - Code to identify the underlying transaction.
type TransactionCode string

const (
	TransactionCodeUnknown       TransactionCode = "Unknown"
	TransactionCodeFee           TransactionCode = "Fee"
	TransactionCodePayment       TransactionCode = "Payment"
	TransactionCodeCash          TransactionCode = "Cash"
	TransactionCodeTransfer      TransactionCode = "Transfer"
	TransactionCodeInterest      TransactionCode = "Interest"
	TransactionCodeCashback      TransactionCode = "Cashback"
	TransactionCodeCheque        TransactionCode = "Cheque"
	TransactionCodeDirectDebit   TransactionCode = "DirectDebit"
	TransactionCodePurchase      TransactionCode = "Purchase"
	TransactionCodeStandingOrder TransactionCode = "StandingOrder"
	TransactionCodeAdjustment    TransactionCode = "Adjustment"
	TransactionCodeCredit        TransactionCode = "Credit"
	TransactionCodeOther         TransactionCode = "Other"
	TransactionCodeNotSupported  TransactionCode = "NotSupported"
)

func (e TransactionCode) ToPointer() *TransactionCode {
	return &e
}
func (e *TransactionCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Fee":
		fallthrough
	case "Payment":
		fallthrough
	case "Cash":
		fallthrough
	case "Transfer":
		fallthrough
	case "Interest":
		fallthrough
	case "Cashback":
		fallthrough
	case "Cheque":
		fallthrough
	case "DirectDebit":
		fallthrough
	case "Purchase":
		fallthrough
	case "StandingOrder":
		fallthrough
	case "Adjustment":
		fallthrough
	case "Credit":
		fallthrough
	case "Other":
		fallthrough
	case "NotSupported":
		*e = TransactionCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TransactionCode: %v", v)
	}
}
