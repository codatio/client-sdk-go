// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PurchaseOrderStatus - Current state of the purchase order
type PurchaseOrderStatus string

const (
	PurchaseOrderStatusUnknown PurchaseOrderStatus = "Unknown"
	PurchaseOrderStatusDraft   PurchaseOrderStatus = "Draft"
	PurchaseOrderStatusOpen    PurchaseOrderStatus = "Open"
	PurchaseOrderStatusClosed  PurchaseOrderStatus = "Closed"
	PurchaseOrderStatusVoid    PurchaseOrderStatus = "Void"
)

func (e PurchaseOrderStatus) ToPointer() *PurchaseOrderStatus {
	return &e
}
func (e *PurchaseOrderStatus) UnmarshalJSON(data []byte) error {
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
		*e = PurchaseOrderStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PurchaseOrderStatus: %v", v)
	}
}
