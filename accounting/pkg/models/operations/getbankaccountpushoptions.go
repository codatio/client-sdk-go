package operations

import (
	"net/http"
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

type GetBankAccountPushOptionsRequest struct {
	PathParams  GetBankAccountPushOptionsPathParams
	QueryParams GetBankAccountPushOptionsQueryParams
}

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                              `json:"description,omitempty"`
	DisplayName *string                                                                              `json:"displayName,omitempty"`
	Required    *bool                                                                                `json:"required,omitempty"`
	Type        *GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                              `json:"value,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                `json:"description,omitempty"`
	DisplayName *string                                                                                                `json:"displayName,omitempty"`
	Required    *bool                                                                                                  `json:"required,omitempty"`
	Type        *GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                `json:"value,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                  `json:"description,omitempty"`
	DisplayName *string                                                                                                                  `json:"displayName,omitempty"`
	Required    *bool                                                                                                                    `json:"required,omitempty"`
	Type        *GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                  `json:"value,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                                    `json:"description,omitempty"`
	DisplayName *string                                                                                                                                    `json:"displayName,omitempty"`
	Required    *bool                                                                                                                                      `json:"required,omitempty"`
	Type        *GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                                    `json:"value,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                         `json:"description"`
	DisplayName string                                                                                                                         `json:"displayName"`
	Options     []GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice  `json:"options,omitempty"`
	Required    bool                                                                                                                           `json:"required"`
	Type        GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum      `json:"type"`
	Validation  *GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo `json:"validation,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                 `json:"description"`
	DisplayName string                                                                                                                 `json:"displayName"`
	Options     []GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                                   `json:"required"`
	Type        GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                               `json:"description"`
	DisplayName string                                                                                               `json:"displayName"`
	Options     []GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                 `json:"required"`
	Type        GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnum string

const (
	GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnumArray     GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnum = "Array"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnumObject    GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnum = "Object"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnumString    GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnum = "String"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnumNumber    GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnum = "Number"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnumBoolean   GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnum = "Boolean"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnumDateTime  GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnum = "DateTime"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnumFile      GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnum = "File"
	GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnumMultiPart GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionPropertyPushValidationInfo struct {
	Information []GetBankAccountPushOptionsPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetBankAccountPushOptionsPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetBankAccountPushOptionsPushOptionPushOptionProperty struct {
	Description string                                                                             `json:"description"`
	DisplayName string                                                                             `json:"displayName"`
	Options     []GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetBankAccountPushOptionsPushOptionPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                               `json:"required"`
	Type        GetBankAccountPushOptionsPushOptionPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetBankAccountPushOptionsPushOptionPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
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

type GetBankAccountPushOptionsPushOption struct {
	Description *string                                                          `json:"description,omitempty"`
	DisplayName string                                                           `json:"displayName"`
	Properties  map[string]GetBankAccountPushOptionsPushOptionPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                             `json:"required"`
	Type        GetBankAccountPushOptionsPushOptionOptionTypeEnum                `json:"type"`
}

type GetBankAccountPushOptionsResponse struct {
	ContentType string
	PushOption  *GetBankAccountPushOptionsPushOption
	StatusCode  int
	RawResponse *http.Response
}
