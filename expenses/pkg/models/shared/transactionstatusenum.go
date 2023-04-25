// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// TransactionStatusEnum - Status of the transaction.
type TransactionStatusEnum string

const (
	TransactionStatusEnumUnknown         TransactionStatusEnum = "Unknown"
	TransactionStatusEnumPending         TransactionStatusEnum = "Pending"
	TransactionStatusEnumValidationError TransactionStatusEnum = "ValidationError"
	TransactionStatusEnumCompleted       TransactionStatusEnum = "Completed"
	TransactionStatusEnumPushError       TransactionStatusEnum = "PushError"
)

func (e TransactionStatusEnum) ToPointer() *TransactionStatusEnum {
	return &e
}

func (e *TransactionStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Pending":
		fallthrough
	case "ValidationError":
		fallthrough
	case "Completed":
		fallthrough
	case "PushError":
		*e = TransactionStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for TransactionStatusEnum: %s", s)
	}
}
