// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type AgedOutstandingAmountDetail struct {
	// The amount outstanding.
	Amount *decimal.Big `decimal:"number" json:"amount,omitempty"`
	// Name of data type with outstanding amount for given period.
	Name *string `json:"name,omitempty"`
}

func (a AgedOutstandingAmountDetail) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AgedOutstandingAmountDetail) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AgedOutstandingAmountDetail) GetAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *AgedOutstandingAmountDetail) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
