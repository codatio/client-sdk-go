// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// AccountCategoryLevel - An object containing an ordered list of account category levels.
type AccountCategoryLevel struct {
	// Confidence level of the category. This will only be populated where `status` is `Suggested`.
	Confidence *decimal.Big `decimal:"number" json:"confidence,omitempty"`
	// Account category name.
	LevelName *string `json:"levelName,omitempty"`
}

func (a AccountCategoryLevel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountCategoryLevel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountCategoryLevel) GetConfidence() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Confidence
}

func (o *AccountCategoryLevel) GetLevelName() *string {
	if o == nil {
		return nil
	}
	return o.LevelName
}
