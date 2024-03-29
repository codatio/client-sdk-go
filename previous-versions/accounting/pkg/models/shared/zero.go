// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ZeroDataType - Allowed name of the 'dataType'.
type ZeroDataType string

const (
	ZeroDataTypePurchaseOrders ZeroDataType = "purchaseOrders"
)

func (e ZeroDataType) ToPointer() *ZeroDataType {
	return &e
}

func (e *ZeroDataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "purchaseOrders":
		*e = ZeroDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ZeroDataType: %v", v)
	}
}

// Zero - Reference to the purchase order line this line was generated from.
type Zero struct {
	// Allowed name of the 'dataType'.
	DataType *ZeroDataType `json:"dataType,omitempty"`
	// 'id' of the underlying record.
	ID *string `json:"id,omitempty"`
	// Line number of the underlying record.
	LineNumber *string `json:"lineNumber,omitempty"`
}

func (o *Zero) GetDataType() *ZeroDataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *Zero) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Zero) GetLineNumber() *string {
	if o == nil {
		return nil
	}
	return o.LineNumber
}
