// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Status of the Payment Method.
type Status string

const (
	StatusUnknown  Status = "Unknown"
	StatusActive   Status = "Active"
	StatusArchived Status = "Archived"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Active":
		fallthrough
	case "Archived":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

// CommercePaymentMethod - A Payment Method represents the payment method(s) used to make payments.
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=commerce-paymentMethods) for this data type.
type CommercePaymentMethod struct {
	// A unique, persistent identifier for this record
	ID           string  `json:"id"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// The name of the PaymentMethod
	Name               *string `json:"name,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Status of the Payment Method.
	Status *Status `json:"status,omitempty"`
}

func (o *CommercePaymentMethod) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CommercePaymentMethod) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *CommercePaymentMethod) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CommercePaymentMethod) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *CommercePaymentMethod) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}
