// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AccountStatusEnum - Status of the account
type AccountStatusEnum string

const (
	AccountStatusEnumUnknown  AccountStatusEnum = "Unknown"
	AccountStatusEnumActive   AccountStatusEnum = "Active"
	AccountStatusEnumArchived AccountStatusEnum = "Archived"
	AccountStatusEnumPending  AccountStatusEnum = "Pending"
)

func (e *AccountStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Active":
		fallthrough
	case "Archived":
		fallthrough
	case "Pending":
		*e = AccountStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountStatusEnum: %s", s)
	}
}