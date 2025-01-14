// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// TransactionStatus - Status of the transaction.
type TransactionStatus string

const (
	TransactionStatusUnknown         TransactionStatus = "Unknown"
	TransactionStatusPending         TransactionStatus = "Pending"
	TransactionStatusValidationError TransactionStatus = "ValidationError"
	TransactionStatusCompleted       TransactionStatus = "Completed"
	TransactionStatusPushError       TransactionStatus = "PushError"
)

func (e TransactionStatus) ToPointer() *TransactionStatus {
	return &e
}
func (e *TransactionStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Pending":
		fallthrough
	case "ValidationError":
		fallthrough
	case "Completed":
		fallthrough
	case "PushError":
		*e = TransactionStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TransactionStatus: %v", v)
	}
}
