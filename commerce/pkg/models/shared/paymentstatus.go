// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PaymentStatus - Status of the payment
type PaymentStatus string

const (
	PaymentStatusPending    PaymentStatus = "Pending"
	PaymentStatusAuthorized PaymentStatus = "Authorized"
	PaymentStatusPaid       PaymentStatus = "Paid"
	PaymentStatusFailed     PaymentStatus = "Failed"
	PaymentStatusCancelled  PaymentStatus = "Cancelled"
	PaymentStatusUnknown    PaymentStatus = "Unknown"
)

func (e PaymentStatus) ToPointer() *PaymentStatus {
	return &e
}

func (e *PaymentStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Pending":
		fallthrough
	case "Authorized":
		fallthrough
	case "Paid":
		fallthrough
	case "Failed":
		fallthrough
	case "Cancelled":
		fallthrough
	case "Unknown":
		*e = PaymentStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentStatus: %v", v)
	}
}
