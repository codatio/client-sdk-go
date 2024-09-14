// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AccountStatus - The current status of the account.
type AccountStatus string

const (
	AccountStatusActive   AccountStatus = "Active"
	AccountStatusArchived AccountStatus = "Archived"
)

func (e AccountStatus) ToPointer() *AccountStatus {
	return &e
}
func (e *AccountStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Active":
		fallthrough
	case "Archived":
		*e = AccountStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountStatus: %v", v)
	}
}