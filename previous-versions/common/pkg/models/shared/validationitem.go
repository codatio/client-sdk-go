// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/utils"
)

type ValidationItem struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// Unique identifier for a validation item.
	ItemID *string `json:"itemId,omitempty"`
	// A message outlining validation item's issue.
	Message *string `json:"message,omitempty"`
	// Name of validator.
	ValidatorName *string `json:"validatorName,omitempty"`
}

func (v ValidationItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *ValidationItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ValidationItem) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *ValidationItem) GetItemID() *string {
	if o == nil {
		return nil
	}
	return o.ItemID
}

func (o *ValidationItem) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ValidationItem) GetValidatorName() *string {
	if o == nil {
		return nil
	}
	return o.ValidatorName
}
