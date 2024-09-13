// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v5/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type TransactionCategory struct {
	// Returns the aggregate confidence of the suggested category for the transaction. The value is between 0 and 100.
	Confidence *decimal.Big `decimal:"number" json:"confidence,omitempty"`
	// An ordered array of category level confidences where each element is the confidence of the corresponding item in the `levels` array.
	Confidences []*decimal.Big `decimal:"number" json:"confidences,omitempty"`
	// The suggested category is an ordered array of category levels where each element (or level) is a subcategory of the previous element (or level).
	Levels []string `json:"levels,omitempty"`
}

func (t TransactionCategory) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TransactionCategory) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TransactionCategory) GetConfidence() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Confidence
}

func (o *TransactionCategory) GetConfidences() []*decimal.Big {
	if o == nil {
		return nil
	}
	return o.Confidences
}

func (o *TransactionCategory) GetLevels() []string {
	if o == nil {
		return nil
	}
	return o.Levels
}
