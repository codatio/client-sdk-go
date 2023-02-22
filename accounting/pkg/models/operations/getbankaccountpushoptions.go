package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetBankAccountPushOptionsPathParams struct {
	AccountID    string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetBankAccountPushOptionsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type GetBankAccountPushOptionsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetBankAccountPushOptionsRequest struct {
	PathParams  GetBankAccountPushOptionsPathParams
	QueryParams GetBankAccountPushOptionsQueryParams
	Security    GetBankAccountPushOptionsSecurity
}

type GetBankAccountPushOptionsPushOptionPushOptionChoice struct {
	Description string                                                                                                                                                                                                                           `json:"description"`
	DisplayName string                                                                                                                                                                                                                           `json:"displayName"`
	Options     []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaPropertiesOptionsItems `json:"options,omitempty"`
	Rel         *string                                                                                                                                                                                                                          `json:"rel,omitempty"`
	Required    bool                                                                                                                                                                                                                             `json:"required"`
	Type        shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaPropertiesTypeEnum       `json:"type"`
	Validation  *shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaPropertiesValidation    `json:"validation,omitempty"`
	Value       string                                                                                                                                                                                                                           `json:"value"`
}

type GetBankAccountPushOptionsPushOptionOptionTypeEnum string

const (
	GetBankAccountPushOptionsPushOptionOptionTypeEnumArray     GetBankAccountPushOptionsPushOptionOptionTypeEnum = "Array"
	GetBankAccountPushOptionsPushOptionOptionTypeEnumObject    GetBankAccountPushOptionsPushOptionOptionTypeEnum = "Object"
	GetBankAccountPushOptionsPushOptionOptionTypeEnumString    GetBankAccountPushOptionsPushOptionOptionTypeEnum = "String"
	GetBankAccountPushOptionsPushOptionOptionTypeEnumNumber    GetBankAccountPushOptionsPushOptionOptionTypeEnum = "Number"
	GetBankAccountPushOptionsPushOptionOptionTypeEnumBoolean   GetBankAccountPushOptionsPushOptionOptionTypeEnum = "Boolean"
	GetBankAccountPushOptionsPushOptionOptionTypeEnumDateTime  GetBankAccountPushOptionsPushOptionOptionTypeEnum = "DateTime"
	GetBankAccountPushOptionsPushOptionOptionTypeEnumFile      GetBankAccountPushOptionsPushOptionOptionTypeEnum = "File"
	GetBankAccountPushOptionsPushOptionOptionTypeEnumMultiPart GetBankAccountPushOptionsPushOptionOptionTypeEnum = "MultiPart"
)

type GetBankAccountPushOptionsPushOptionPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   string  `json:"field"`
	Ref     *string `json:"ref,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushValidationInfo struct {
	Information []shared.Onecompanies1Percent7BcompanyIDPercent7D1connections1Percent7BconnectionIDPercent7D1options1bankAccounts1Percent7BaccountIDPercent7D1bankTransactionsGetResponses200ContentApplication1jsonSchemaPropertiesValidationPropertiesWarningsItems `json:"information,omitempty"`
	Warnings    []GetBankAccountPushOptionsPushOptionPushValidationInfoPushFieldValidation                                                                                                                                                                            `json:"warnings,omitempty"`
}

type GetBankAccountPushOptionsPushOption struct {
	Description string                                                 `json:"description"`
	DisplayName string                                                 `json:"displayName"`
	Options     []GetBankAccountPushOptionsPushOptionPushOptionChoice  `json:"options,omitempty"`
	Rel         *string                                                `json:"rel,omitempty"`
	Required    bool                                                   `json:"required"`
	Type        GetBankAccountPushOptionsPushOptionOptionTypeEnum      `json:"type"`
	Validation  *GetBankAccountPushOptionsPushOptionPushValidationInfo `json:"validation,omitempty"`
}

type GetBankAccountPushOptionsResponse struct {
	ContentType string
	PushOption  *GetBankAccountPushOptionsPushOption
	StatusCode  int
}
