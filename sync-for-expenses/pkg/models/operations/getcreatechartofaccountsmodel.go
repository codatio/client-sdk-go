// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/models/shared"
	"net/http"
)

type GetCreateChartOfAccountsModelRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *GetCreateChartOfAccountsModelRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetCreateChartOfAccountsModelRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

// GetCreateChartOfAccountsModelPushOption - OK
type GetCreateChartOfAccountsModelPushOption struct {
	// A description of the property.
	Description *string `json:"description,omitempty"`
	// The property's display name.
	DisplayName string                                                                                                                                                                                                 `json:"displayName"`
	Options     []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1chartOfAccountsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1chartOfAccountsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionProperty `json:"properties,omitempty"`
	// The property is required if `True`.
	Required bool `json:"required"`
	// The option type.
	Type       shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1chartOfAccountsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType      `json:"type"`
	Validation *shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1chartOfAccountsGetResponses200ContentApplication1jsonSchemaDefinitionsPushValidationInfo `json:"validation,omitempty"`
}

func (o *GetCreateChartOfAccountsModelPushOption) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *GetCreateChartOfAccountsModelPushOption) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}

func (o *GetCreateChartOfAccountsModelPushOption) GetOptions() []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1chartOfAccountsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionChoice {
	if o == nil {
		return nil
	}
	return o.Options
}

func (o *GetCreateChartOfAccountsModelPushOption) GetProperties() map[string]shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1chartOfAccountsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionProperty {
	if o == nil {
		return nil
	}
	return o.Properties
}

func (o *GetCreateChartOfAccountsModelPushOption) GetRequired() bool {
	if o == nil {
		return false
	}
	return o.Required
}

func (o *GetCreateChartOfAccountsModelPushOption) GetType() shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1chartOfAccountsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType {
	if o == nil {
		return shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1chartOfAccountsGetResponses200ContentApplication1jsonSchemaDefinitionsPushOptionType("")
	}
	return o.Type
}

func (o *GetCreateChartOfAccountsModelPushOption) GetValidation() *shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1chartOfAccountsGetResponses200ContentApplication1jsonSchemaDefinitionsPushValidationInfo {
	if o == nil {
		return nil
	}
	return o.Validation
}

type GetCreateChartOfAccountsModelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// OK
	PushOption *GetCreateChartOfAccountsModelPushOption
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetCreateChartOfAccountsModelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCreateChartOfAccountsModelResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetCreateChartOfAccountsModelResponse) GetPushOption() *GetCreateChartOfAccountsModelPushOption {
	if o == nil {
		return nil
	}
	return o.PushOption
}

func (o *GetCreateChartOfAccountsModelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCreateChartOfAccountsModelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
