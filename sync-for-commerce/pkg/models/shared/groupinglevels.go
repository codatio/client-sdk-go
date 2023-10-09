// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/utils"
)

type GroupingLevels struct {
	AdditionalProperties map[string]interface{}     `additionalProperties:"true" json:"-"`
	InvoiceLevel         *InvoiceLevelSelection     `json:"invoiceLevel,omitempty"`
	InvoiceLineLevel     *InvoiceLineLevelSelection `json:"invoiceLineLevel,omitempty"`
}

func (g GroupingLevels) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GroupingLevels) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GroupingLevels) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *GroupingLevels) GetInvoiceLevel() *InvoiceLevelSelection {
	if o == nil {
		return nil
	}
	return o.InvoiceLevel
}

func (o *GroupingLevels) GetInvoiceLineLevel() *InvoiceLineLevelSelection {
	if o == nil {
		return nil
	}
	return o.InvoiceLineLevel
}
