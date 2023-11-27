// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// BillStatus - Current state of the bill.
type BillStatus string

const (
	BillStatusUnknown       BillStatus = "Unknown"
	BillStatusOpen          BillStatus = "Open"
	BillStatusPartiallyPaid BillStatus = "PartiallyPaid"
	BillStatusPaid          BillStatus = "Paid"
	BillStatusVoid          BillStatus = "Void"
	BillStatusDraft         BillStatus = "Draft"
)

func (e BillStatus) ToPointer() *BillStatus {
	return &e
}

func (e *BillStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Open":
		fallthrough
	case "PartiallyPaid":
		fallthrough
	case "Paid":
		fallthrough
	case "Void":
		fallthrough
	case "Draft":
		*e = BillStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BillStatus: %v", v)
	}
}
