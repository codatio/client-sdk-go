// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// BankTransactionType - Type of transaction for the bank statement line.
type BankTransactionType string

const (
	BankTransactionTypeUnknown     BankTransactionType = "Unknown"
	BankTransactionTypeCredit      BankTransactionType = "Credit"
	BankTransactionTypeDebit       BankTransactionType = "Debit"
	BankTransactionTypeInt         BankTransactionType = "Int"
	BankTransactionTypeDiv         BankTransactionType = "Div"
	BankTransactionTypeFee         BankTransactionType = "Fee"
	BankTransactionTypeSerChg      BankTransactionType = "SerChg"
	BankTransactionTypeDep         BankTransactionType = "Dep"
	BankTransactionTypeAtm         BankTransactionType = "Atm"
	BankTransactionTypePos         BankTransactionType = "Pos"
	BankTransactionTypeXfer        BankTransactionType = "Xfer"
	BankTransactionTypeCheck       BankTransactionType = "Check"
	BankTransactionTypePayment     BankTransactionType = "Payment"
	BankTransactionTypeCash        BankTransactionType = "Cash"
	BankTransactionTypeDirectDep   BankTransactionType = "DirectDep"
	BankTransactionTypeDirectDebit BankTransactionType = "DirectDebit"
	BankTransactionTypeRepeatPmt   BankTransactionType = "RepeatPmt"
	BankTransactionTypeOther       BankTransactionType = "Other"
)

func (e BankTransactionType) ToPointer() *BankTransactionType {
	return &e
}
func (e *BankTransactionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Credit":
		fallthrough
	case "Debit":
		fallthrough
	case "Int":
		fallthrough
	case "Div":
		fallthrough
	case "Fee":
		fallthrough
	case "SerChg":
		fallthrough
	case "Dep":
		fallthrough
	case "Atm":
		fallthrough
	case "Pos":
		fallthrough
	case "Xfer":
		fallthrough
	case "Check":
		fallthrough
	case "Payment":
		fallthrough
	case "Cash":
		fallthrough
	case "DirectDep":
		fallthrough
	case "DirectDebit":
		fallthrough
	case "RepeatPmt":
		fallthrough
	case "Other":
		*e = BankTransactionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BankTransactionType: %v", v)
	}
}
