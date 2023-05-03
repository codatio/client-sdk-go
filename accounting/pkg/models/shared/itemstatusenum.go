// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ItemStatusEnum - Current state of the item, either:
//
// - `Active`: Available for use
// - `Archived`: Unavailable
// - `Unknown`
//
// Due to a [limitation in Xero's API](https://docs.codat.io/integrations/accounting/xero/xero-faq#why-do-all-of-my-items-from-xero-have-their-status-as-unknown), all items from Xero are mapped as `Unknown`.
type ItemStatusEnum string

const (
	ItemStatusEnumUnknown  ItemStatusEnum = "Unknown"
	ItemStatusEnumActive   ItemStatusEnum = "Active"
	ItemStatusEnumArchived ItemStatusEnum = "Archived"
)

func (e ItemStatusEnum) ToPointer() *ItemStatusEnum {
	return &e
}

func (e *ItemStatusEnum) UnmarshalJSON(data []byte) error {
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
		*e = ItemStatusEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ItemStatusEnum: %v", v)
	}
}
