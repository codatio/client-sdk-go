// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// IntegrationTypeEnum - Type of transaction that has been processed e.g. Expense or Bank Feed.
type IntegrationTypeEnum string

const (
	IntegrationTypeEnumExpenses  IntegrationTypeEnum = "expenses"
	IntegrationTypeEnumBankfeeds IntegrationTypeEnum = "bankfeeds"
)

func (e IntegrationTypeEnum) ToPointer() *IntegrationTypeEnum {
	return &e
}

func (e *IntegrationTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "expenses":
		fallthrough
	case "bankfeeds":
		*e = IntegrationTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IntegrationTypeEnum: %v", v)
	}
}
