// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/utils"
)

type Configuration struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	Fees                 *Fees                  `json:"fees,omitempty"`
	NewPayments          *NewPayments           `json:"newPayments,omitempty"`
	Payments             *Payments              `json:"payments,omitempty"`
	Sales                *Sales                 `json:"sales,omitempty"`
}

func (c Configuration) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *Configuration) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Configuration) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *Configuration) GetFees() *Fees {
	if o == nil {
		return nil
	}
	return o.Fees
}

func (o *Configuration) GetNewPayments() *NewPayments {
	if o == nil {
		return nil
	}
	return o.NewPayments
}

func (o *Configuration) GetPayments() *Payments {
	if o == nil {
		return nil
	}
	return o.Payments
}

func (o *Configuration) GetSales() *Sales {
	if o == nil {
		return nil
	}
	return o.Sales
}
