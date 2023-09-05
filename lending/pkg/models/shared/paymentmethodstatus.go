// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PaymentMethodStatus - Status of the Payment Method
type PaymentMethodStatus string

const (
	PaymentMethodStatusActive   PaymentMethodStatus = "Active"
	PaymentMethodStatusArchived PaymentMethodStatus = "Archived"
	PaymentMethodStatusUnknown  PaymentMethodStatus = "Unknown"
)

func (e PaymentMethodStatus) ToPointer() *PaymentMethodStatus {
	return &e
}

func (e *PaymentMethodStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Active":
		fallthrough
	case "Archived":
		fallthrough
	case "Unknown":
		*e = PaymentMethodStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentMethodStatus: %v", v)
	}
}