// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PushOperationStatusEnum - The status of the push operation.
type PushOperationStatusEnum string

const (
	PushOperationStatusEnumPending  PushOperationStatusEnum = "Pending"
	PushOperationStatusEnumFailed   PushOperationStatusEnum = "Failed"
	PushOperationStatusEnumSuccess  PushOperationStatusEnum = "Success"
	PushOperationStatusEnumTimedOut PushOperationStatusEnum = "TimedOut"
)

func (e PushOperationStatusEnum) ToPointer() *PushOperationStatusEnum {
	return &e
}

func (e *PushOperationStatusEnum) UnmarshalJSON(data []byte) error {
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
		*e = PushOperationStatusEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PushOperationStatusEnum: %v", v)
	}
}
