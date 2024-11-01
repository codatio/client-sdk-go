// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PushOperationStatus - The current status of the push operation.
type PushOperationStatus string

const (
	PushOperationStatusPending  PushOperationStatus = "Pending"
	PushOperationStatusFailed   PushOperationStatus = "Failed"
	PushOperationStatusSuccess  PushOperationStatus = "Success"
	PushOperationStatusTimedOut PushOperationStatus = "TimedOut"
)

func (e PushOperationStatus) ToPointer() *PushOperationStatus {
	return &e
}
func (e *PushOperationStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Pending":
		fallthrough
	case "Failed":
		fallthrough
	case "Success":
		fallthrough
	case "TimedOut":
		*e = PushOperationStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PushOperationStatus: %v", v)
	}
}
