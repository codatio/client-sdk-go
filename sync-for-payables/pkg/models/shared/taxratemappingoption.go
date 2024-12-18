// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// TaxRateStatus - Status of the tax rate in the accounting software.
// - `Active` - An active tax rate in use by a company.
// - `Archived` - A tax rate that has been archived or is inactive in the accounting software.
type TaxRateStatus string

const (
	TaxRateStatusActive   TaxRateStatus = "Active"
	TaxRateStatusArchived TaxRateStatus = "Archived"
)

func (e TaxRateStatus) ToPointer() *TaxRateStatus {
	return &e
}
func (e *TaxRateStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Active":
		fallthrough
	case "Archived":
		*e = TaxRateStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaxRateStatus: %v", v)
	}
}

type TaxRateMappingOption struct {
	// Identifier for the tax rate, unique for the company in the accounting software.
	ID *string `json:"id,omitempty"`
	// Codat-augmented name of the tax rate in the accounting software.
	Name *string `json:"name,omitempty"`
	// Code for the tax rate from the accounting software.
	Code *string `json:"code,omitempty"`
	// See Effective tax rates description.
	EffectiveTaxRate *decimal.Big `decimal:"number" json:"effectiveTaxRate,omitempty"`
	// Total (not compounded) sum of the components of a tax rate.
	TotalTaxRate *decimal.Big `decimal:"number" json:"totalTaxRate,omitempty"`
	// Status of the tax rate in the accounting software.
	// - `Active` - An active tax rate in use by a company.
	// - `Archived` - A tax rate that has been archived or is inactive in the accounting software.
	Status *TaxRateStatus `json:"status,omitempty"`
}

func (t TaxRateMappingOption) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaxRateMappingOption) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaxRateMappingOption) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TaxRateMappingOption) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TaxRateMappingOption) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *TaxRateMappingOption) GetEffectiveTaxRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.EffectiveTaxRate
}

func (o *TaxRateMappingOption) GetTotalTaxRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalTaxRate
}

func (o *TaxRateMappingOption) GetStatus() *TaxRateStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
