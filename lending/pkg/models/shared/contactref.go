// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ContactRefDataType - Allowed name of the 'dataType'.
type ContactRefDataType string

const (
	ContactRefDataTypeCustomers ContactRefDataType = "customers"
	ContactRefDataTypeSuppliers ContactRefDataType = "suppliers"
)

func (e ContactRefDataType) ToPointer() *ContactRefDataType {
	return &e
}
func (e *ContactRefDataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "customers":
		fallthrough
	case "suppliers":
		*e = ContactRefDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ContactRefDataType: %v", v)
	}
}

type ContactRef struct {
	// Allowed name of the 'dataType'.
	DataType *ContactRefDataType `json:"dataType,omitempty"`
	// Unique identifier for a customer or supplier.
	ID string `json:"id"`
}

func (o *ContactRef) GetDataType() *ContactRefDataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *ContactRef) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
