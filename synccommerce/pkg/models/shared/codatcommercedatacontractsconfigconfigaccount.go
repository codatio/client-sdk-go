package shared

type CodatCommerceDataContractsConfigConfigAccount struct {
	AccountOptions    []CodatCommerceDataContractsConfigAccountOption `json:"accountOptions,omitempty"`
	DescriptionText   *string                                         `json:"descriptionText,omitempty"`
	LabelText         *string                                         `json:"labelText,omitempty"`
	Required          *bool                                           `json:"required,omitempty"`
	SelectedAccountID *string                                         `json:"selectedAccountId,omitempty"`
}
