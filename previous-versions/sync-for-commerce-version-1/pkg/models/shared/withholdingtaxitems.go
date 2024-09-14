// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type WithholdingTaxItems struct {
	// Amount of tax withheld.
	Amount *decimal.Big `decimal:"number" json:"amount"`
	// Name assigned to withheld tax.
	Name string `json:"name"`
}

func (w WithholdingTaxItems) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WithholdingTaxItems) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WithholdingTaxItems) GetAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.Amount
}

func (o *WithholdingTaxItems) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
