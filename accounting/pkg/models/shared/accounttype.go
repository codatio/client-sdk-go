// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AccountType - Type of account
type AccountType string

const (
	AccountTypeUnknown   AccountType = "Unknown"
	AccountTypeAsset     AccountType = "Asset"
	AccountTypeExpense   AccountType = "Expense"
	AccountTypeIncome    AccountType = "Income"
	AccountTypeLiability AccountType = "Liability"
	AccountTypeEquity    AccountType = "Equity"
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
	case "Unknown":
		fallthrough
	case "Asset":
		fallthrough
	case "Expense":
		fallthrough
	case "Income":
		fallthrough
	case "Liability":
		fallthrough
	case "Equity":
		*e = AccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountType: %v", v)
	}
}