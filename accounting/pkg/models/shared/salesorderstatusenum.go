// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// SalesOrderStatusEnum - Current state of the sales order.
type SalesOrderStatusEnum string

const (
	SalesOrderStatusEnumUnknown SalesOrderStatusEnum = "Unknown"
	SalesOrderStatusEnumDraft   SalesOrderStatusEnum = "Draft"
	SalesOrderStatusEnumOpen    SalesOrderStatusEnum = "Open"
	SalesOrderStatusEnumClosed  SalesOrderStatusEnum = "Closed"
	SalesOrderStatusEnumVoid    SalesOrderStatusEnum = "Void"
)

func (e SalesOrderStatusEnum) ToPointer() *SalesOrderStatusEnum {
	return &e
}

func (e *SalesOrderStatusEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Draft":
		fallthrough
	case "Open":
		fallthrough
	case "Closed":
		fallthrough
	case "Void":
		*e = SalesOrderStatusEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SalesOrderStatusEnum: %v", v)
	}
}
