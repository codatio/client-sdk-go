// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AccountIdentifierType - Type of account
type AccountIdentifierType string

const (
	AccountIdentifierTypeAccount    AccountIdentifierType = "Account"
	AccountIdentifierTypeCard       AccountIdentifierType = "Card"
	AccountIdentifierTypeCredit     AccountIdentifierType = "Credit"
	AccountIdentifierTypeDepository AccountIdentifierType = "Depository"
	AccountIdentifierTypeInvestment AccountIdentifierType = "Investment"
	AccountIdentifierTypeLoan       AccountIdentifierType = "Loan"
	AccountIdentifierTypeOther      AccountIdentifierType = "Other"
)

func (e AccountIdentifierType) ToPointer() *AccountIdentifierType {
	return &e
}

func (e *AccountIdentifierType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Account":
		fallthrough
	case "Card":
		fallthrough
	case "Credit":
		fallthrough
	case "Depository":
		fallthrough
	case "Investment":
		fallthrough
	case "Loan":
		fallthrough
	case "Other":
		*e = AccountIdentifierType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountIdentifierType: %v", v)
	}
}