package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushPathParams struct {
	CompanyID    string              `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string              `pathParam:"style=simple,explode=false,name=connectionId"`
	DataType     shared.DataTypeEnum `pathParam:"style=simple,explode=false,name=dataType"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushRequest struct {
	PathParams GetCompaniesCompanyIDConnectionsConnectionIDPushPathParams
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoice struct {
	Description string                                                                                                                                                                                            `json:"description"`
	DisplayName string                                                                                                                                                                                            `json:"displayName"`
	Options     []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1Percent7BdataTypePercent7DGetResponses200ContentApplication1jsonSchemaPropertiesOptionsItems `json:"options,omitempty"`
	Rel         *string                                                                                                                                                                                           `json:"rel,omitempty"`
	Required    bool                                                                                                                                                                                              `json:"required"`
	Type        shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1Percent7BdataTypePercent7DGetResponses200ContentApplication1jsonSchemaPropertiesTypeEnum       `json:"type"`
	Validation  *shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1Percent7BdataTypePercent7DGetResponses200ContentApplication1jsonSchemaPropertiesValidation    `json:"validation,omitempty"`
	Value       string                                                                                                                                                                                            `json:"value"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnumArray     GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnum = "Array"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnumObject    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnum = "Object"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnumString    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnum = "String"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnumNumber    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnum = "Number"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnumBoolean   GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnum = "Boolean"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnumDateTime  GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnum = "DateTime"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnumFile      GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnum = "File"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnumMultiPart GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnum = "MultiPart"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   string  `json:"field"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushValidationInfo struct {
	Information []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1Percent7BdataTypePercent7DGetResponses200ContentApplication1jsonSchemaPropertiesValidationPropertiesWarningsItems `json:"information,omitempty"`
	Warnings    []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushValidationInfoPushFieldValidation                                                                                                                      `json:"warnings,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOption struct {
	Description string                                                                        `json:"description"`
	DisplayName string                                                                        `json:"displayName"`
	Options     []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoice  `json:"options,omitempty"`
	Rel         *string                                                                       `json:"rel,omitempty"`
	Required    bool                                                                          `json:"required"`
	Type        GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnum      `json:"type"`
	Validation  *GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushValidationInfo `json:"validation,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushResponse struct {
	ContentType string
	PushOption  *GetCompaniesCompanyIDConnectionsConnectionIDPushPushOption
	StatusCode  int
}
