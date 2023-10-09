// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/utils"
)

type TaxRateMapping struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// Selected tax rate id from the list of tax rates on the accounting software.
	SelectedAccountingTaxRateID *string `json:"selectedAccountingTaxRateId,omitempty"`
	// Selected tax component id from the list of tax components on the commerce software.
	SelectedCommerceTaxRateIds []string `json:"selectedCommerceTaxRateIds,omitempty"`
}

func (t TaxRateMapping) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaxRateMapping) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaxRateMapping) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *TaxRateMapping) GetSelectedAccountingTaxRateID() *string {
	if o == nil {
		return nil
	}
	return o.SelectedAccountingTaxRateID
}

func (o *TaxRateMapping) GetSelectedCommerceTaxRateIds() []string {
	if o == nil {
		return nil
	}
	return o.SelectedCommerceTaxRateIds
}
