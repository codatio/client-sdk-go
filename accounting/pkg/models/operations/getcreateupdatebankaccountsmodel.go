package operations

import (
	"net/http"
)

type GetCreateUpdateBankAccountsModelRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                     `json:"description,omitempty"`
	DisplayName *string                                                                                     `json:"displayName,omitempty"`
	Required    *bool                                                                                       `json:"required,omitempty"`
	Type        *GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                     `json:"value,omitempty"`
}

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                       `json:"description,omitempty"`
	DisplayName *string                                                                                                       `json:"displayName,omitempty"`
	Required    *bool                                                                                                         `json:"required,omitempty"`
	Type        *GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                       `json:"value,omitempty"`
}

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                         `json:"description,omitempty"`
	DisplayName *string                                                                                                                         `json:"displayName,omitempty"`
	Required    *bool                                                                                                                           `json:"required,omitempty"`
	Type        *GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                         `json:"value,omitempty"`
}

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                                           `json:"description,omitempty"`
	DisplayName *string                                                                                                                                           `json:"displayName,omitempty"`
	Required    *bool                                                                                                                                             `json:"required,omitempty"`
	Type        *GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                                           `json:"value,omitempty"`
}

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                                `json:"description"`
	DisplayName string                                                                                                                                `json:"displayName"`
	Options     []GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice  `json:"options,omitempty"`
	Required    bool                                                                                                                                  `json:"required"`
	Type        GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum      `json:"type"`
	Validation  *GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo `json:"validation,omitempty"`
}

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                        `json:"description"`
	DisplayName string                                                                                                                        `json:"displayName"`
	Options     []GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                                          `json:"required"`
	Type        GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                      `json:"description"`
	DisplayName string                                                                                                      `json:"displayName"`
	Options     []GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                        `json:"required"`
	Type        GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnum string

const (
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnumArray     GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnumObject    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnumString    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnum = "String"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnumNumber    GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnumBoolean   GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnumDateTime  GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnumFile      GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnum = "File"
	GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnumMultiPart GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateUpdateBankAccountsModelPushOptionPushOptionProperty struct {
	Description string                                                                                    `json:"description"`
	DisplayName string                                                                                    `json:"displayName"`
	Options     []GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                      `json:"required"`
	Type        GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateUpdateBankAccountsModelPushOptionPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnum string

const (
	GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnumArray     GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnum = "Array"
	GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnumObject    GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnum = "Object"
	GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnumString    GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnum = "String"
	GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnumNumber    GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnum = "Number"
	GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnumBoolean   GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnum = "Boolean"
	GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnumDateTime  GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnum = "DateTime"
	GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnumFile      GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnum = "File"
	GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnumMultiPart GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBankAccountsModelPushOption struct {
	Description *string                                                                 `json:"description,omitempty"`
	DisplayName string                                                                  `json:"displayName"`
	Properties  map[string]GetCreateUpdateBankAccountsModelPushOptionPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                    `json:"required"`
	Type        GetCreateUpdateBankAccountsModelPushOptionOptionTypeEnum                `json:"type"`
}

type GetCreateUpdateBankAccountsModelResponse struct {
	ContentType string
	PushOption  *GetCreateUpdateBankAccountsModelPushOption
	StatusCode  int
	RawResponse *http.Response
}