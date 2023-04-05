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

func (e *SalesOrderStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Draft":
		fallthrough
	case "Open":
		fallthrough
	case "Closed":
		fallthrough
	case "Void":
		*e = SalesOrderStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SalesOrderStatusEnum: %s", s)
	}
}