// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/utils"
)

// ConfigAccount - G/L account object for configuration.
type ConfigAccount struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// Object containing account options.
	AccountOptions []AccountOption `json:"accountOptions,omitempty"`
	// Descriprtive text for sales configuration section.
	DescriptionText *string `json:"descriptionText,omitempty"`
	// Label text for sales configuration section.
	LabelText *string `json:"labelText,omitempty"`
	// Required section to be configured for sync.
	Required *bool `json:"required,omitempty"`
	// Selected account id from the list of available accounts.
	SelectedAccountID *string `json:"selectedAccountId,omitempty"`
}

func (c ConfigAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConfigAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConfigAccount) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *ConfigAccount) GetAccountOptions() []AccountOption {
	if o == nil {
		return nil
	}
	return o.AccountOptions
}

func (o *ConfigAccount) GetDescriptionText() *string {
	if o == nil {
		return nil
	}
	return o.DescriptionText
}

func (o *ConfigAccount) GetLabelText() *string {
	if o == nil {
		return nil
	}
	return o.LabelText
}

func (o *ConfigAccount) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *ConfigAccount) GetSelectedAccountID() *string {
	if o == nil {
		return nil
	}
	return o.SelectedAccountID
}
