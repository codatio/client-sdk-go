// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v6/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type Items struct {
	// Amount of tax withheld.
	Amount *decimal.Big `decimal:"number" json:"amount"`
	// Name assigned to withheld tax.
	Name string `json:"name"`
}

func (i Items) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *Items) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Items) GetAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.Amount
}

func (o *Items) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
