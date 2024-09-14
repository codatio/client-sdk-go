// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// BankAccountStatus - The current status of the bank account.
type BankAccountStatus string

const (
	BankAccountStatusActive   BankAccountStatus = "Active"
	BankAccountStatusArchived BankAccountStatus = "Archived"
)

func (e BankAccountStatus) ToPointer() *BankAccountStatus {
	return &e
}
func (e *BankAccountStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Active":
		fallthrough
	case "Archived":
		*e = BankAccountStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BankAccountStatus: %v", v)
	}
}
