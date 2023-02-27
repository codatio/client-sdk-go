package operations

import (
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
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

type GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum string

const (
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumArray     GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "Array"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumObject    GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "Object"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumString    GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "String"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumNumber    GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "Number"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumBoolean   GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "Boolean"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumDateTime  GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "DateTime"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumFile      GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "File"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumMultiPart GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   string  `json:"field"`
	Ref     *string `json:"ref,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfo struct {
	Information []GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoice struct {
	Description string                                                                                                                 `json:"description"`
	DisplayName string                                                                                                                 `json:"displayName"`
	Rel         *string                                                                                                                `json:"rel,omitempty"`
	Required    bool                                                                                                                   `json:"required"`
	Type        GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum      `json:"type"`
	Validation  *GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfo `json:"validation,omitempty"`
	Value       string                                                                                                                 `json:"value"`
}

type GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum string

const (
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumArray     GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "Array"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumObject    GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "Object"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumString    GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "String"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumNumber    GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "Number"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumBoolean   GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "Boolean"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumDateTime  GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "DateTime"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumFile      GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "File"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumMultiPart GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   string  `json:"field"`
	Ref     *string `json:"ref,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfo struct {
	Information []GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoice struct {
	Description string                                                                                                 `json:"description"`
	DisplayName string                                                                                                 `json:"displayName"`
	Options     []GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoice  `json:"options,omitempty"`
	Rel         *string                                                                                                `json:"rel,omitempty"`
	Required    bool                                                                                                   `json:"required"`
	Type        GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum      `json:"type"`
	Validation  *GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfo `json:"validation,omitempty"`
	Value       string                                                                                                 `json:"value"`
}

type GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum string

const (
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnumArray     GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum = "Array"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnumObject    GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum = "Object"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnumString    GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum = "String"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnumNumber    GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum = "Number"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnumBoolean   GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum = "Boolean"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnumDateTime  GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum = "DateTime"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnumFile      GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum = "File"
	GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnumMultiPart GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   string  `json:"field"`
	Ref     *string `json:"ref,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushValidationInfo struct {
	Information []GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoice struct {
	Description string                                                                                 `json:"description"`
	DisplayName string                                                                                 `json:"displayName"`
	Options     []GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushOptionChoice  `json:"options,omitempty"`
	Rel         *string                                                                                `json:"rel,omitempty"`
	Required    bool                                                                                   `json:"required"`
	Type        GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum      `json:"type"`
	Validation  *GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoicePushValidationInfo `json:"validation,omitempty"`
	Value       string                                                                                 `json:"value"`
}

type GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnum string

const (
	GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnumArray     GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnum = "Array"
	GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnumObject    GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnum = "Object"
	GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnumString    GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnum = "String"
	GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnumNumber    GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnum = "Number"
	GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnumBoolean   GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnum = "Boolean"
	GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnumDateTime  GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnum = "DateTime"
	GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnumFile      GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnum = "File"
	GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnumMultiPart GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetBankAccountPushOptionsPushOptionPushOptionChoicePushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   string  `json:"field"`
	Ref     *string `json:"ref,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionChoicePushValidationInfo struct {
	Information []GetBankAccountPushOptionsPushOptionPushOptionChoicePushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetBankAccountPushOptionsPushOptionPushOptionChoicePushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionChoice struct {
	Description string                                                                 `json:"description"`
	DisplayName string                                                                 `json:"displayName"`
	Options     []GetBankAccountPushOptionsPushOptionPushOptionChoicePushOptionChoice  `json:"options,omitempty"`
	Rel         *string                                                                `json:"rel,omitempty"`
	Required    bool                                                                   `json:"required"`
	Type        GetBankAccountPushOptionsPushOptionPushOptionChoiceOptionTypeEnum      `json:"type"`
	Validation  *GetBankAccountPushOptionsPushOptionPushOptionChoicePushValidationInfo `json:"validation,omitempty"`
	Value       string                                                                 `json:"value"`
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
	Information []GetBankAccountPushOptionsPushOptionPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetBankAccountPushOptionsPushOptionPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
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
