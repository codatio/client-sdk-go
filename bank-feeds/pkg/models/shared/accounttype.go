// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AccountType - The type of bank account e.g. checking, savings, loan, creditCard, prepaidCard.
type AccountType string

const (
	AccountTypeChecking    AccountType = "checking"
	AccountTypeSavings     AccountType = "savings"
	AccountTypeLoan        AccountType = "loan"
	AccountTypeCreditCard  AccountType = "creditCard"
	AccountTypePrepaidCard AccountType = "prepaidCard"
)

func (e AccountType) ToPointer() *AccountType {
	return &e
}
func (e *AccountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "checking":
		fallthrough
	case "savings":
		fallthrough
	case "loan":
		fallthrough
	case "creditCard":
		fallthrough
	case "prepaidCard":
		*e = AccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountType: %v", v)
	}
}
