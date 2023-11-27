// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AccountType - The type of transactions and balances on the account.
// For Credit accounts, positive balances are liabilities, and positive transactions **reduce** liabilities.
// For Debit accounts, positive balances are assets, and positive transactions **increase** assets.
type AccountType string

const (
	AccountTypeUnknown AccountType = "Unknown"
	AccountTypeCredit  AccountType = "Credit"
	AccountTypeDebit   AccountType = "Debit"
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
	case "Credit":
		fallthrough
	case "Debit":
		*e = AccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountType: %v", v)
	}
}
