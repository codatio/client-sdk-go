// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// TaxRateStatusEnum - Status of the tax rate in the accounting platform.
// - `Active` - An active tax rate in use by a company.
// - `Archived` - A tax rate that has been archived or is inactive in the accounting platform.
// - `Unknown` - Where the status of the tax rate cannot be determined from the underlying platform.
type TaxRateStatusEnum string

const (
	TaxRateStatusEnumUnknown  TaxRateStatusEnum = "Unknown"
	TaxRateStatusEnumActive   TaxRateStatusEnum = "Active"
	TaxRateStatusEnumArchived TaxRateStatusEnum = "Archived"
)

func (e *TaxRateStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Active":
		fallthrough
	case "Archived":
		*e = TaxRateStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for TaxRateStatusEnum: %s", s)
	}
}