package operations

import (
	"net/http"
)

type GetCreateDirectIncomesModelPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCreateDirectIncomesModelRequest struct {
	PathParams GetCreateDirectIncomesModelPathParams
}

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                `json:"description,omitempty"`
	DisplayName *string                                                                                `json:"displayName,omitempty"`
	Required    *bool                                                                                  `json:"required,omitempty"`
	Type        *GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                `json:"value,omitempty"`
}

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                  `json:"description,omitempty"`
	DisplayName *string                                                                                                  `json:"displayName,omitempty"`
	Required    *bool                                                                                                    `json:"required,omitempty"`
	Type        *GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                  `json:"value,omitempty"`
}

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                    `json:"description,omitempty"`
	DisplayName *string                                                                                                                    `json:"displayName,omitempty"`
	Required    *bool                                                                                                                      `json:"required,omitempty"`
	Type        *GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                    `json:"value,omitempty"`
}

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                                      `json:"description,omitempty"`
	DisplayName *string                                                                                                                                      `json:"displayName,omitempty"`
	Required    *bool                                                                                                                                        `json:"required,omitempty"`
	Type        *GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                                      `json:"value,omitempty"`
}

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                           `json:"description"`
	DisplayName string                                                                                                                           `json:"displayName"`
	Options     []GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice  `json:"options,omitempty"`
	Required    bool                                                                                                                             `json:"required"`
	Type        GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum      `json:"type"`
	Validation  *GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo `json:"validation,omitempty"`
}

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                   `json:"description"`
	DisplayName string                                                                                                                   `json:"displayName"`
	Options     []GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                                     `json:"required"`
	Type        GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                 `json:"description"`
	DisplayName string                                                                                                 `json:"displayName"`
	Options     []GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                   `json:"required"`
	Type        GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnum string

const (
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnumArray     GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnumObject    GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnumString    GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnum = "String"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnumNumber    GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnumBoolean   GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnumDateTime  GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnumFile      GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnum = "File"
	GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnumMultiPart GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateDirectIncomesModelPushOptionPushOptionProperty struct {
	Description string                                                                               `json:"description"`
	DisplayName string                                                                               `json:"displayName"`
	Options     []GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                 `json:"required"`
	Type        GetCreateDirectIncomesModelPushOptionPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateDirectIncomesModelPushOptionPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateDirectIncomesModelPushOptionOptionTypeEnum string

const (
	GetCreateDirectIncomesModelPushOptionOptionTypeEnumArray     GetCreateDirectIncomesModelPushOptionOptionTypeEnum = "Array"
	GetCreateDirectIncomesModelPushOptionOptionTypeEnumObject    GetCreateDirectIncomesModelPushOptionOptionTypeEnum = "Object"
	GetCreateDirectIncomesModelPushOptionOptionTypeEnumString    GetCreateDirectIncomesModelPushOptionOptionTypeEnum = "String"
	GetCreateDirectIncomesModelPushOptionOptionTypeEnumNumber    GetCreateDirectIncomesModelPushOptionOptionTypeEnum = "Number"
	GetCreateDirectIncomesModelPushOptionOptionTypeEnumBoolean   GetCreateDirectIncomesModelPushOptionOptionTypeEnum = "Boolean"
	GetCreateDirectIncomesModelPushOptionOptionTypeEnumDateTime  GetCreateDirectIncomesModelPushOptionOptionTypeEnum = "DateTime"
	GetCreateDirectIncomesModelPushOptionOptionTypeEnumFile      GetCreateDirectIncomesModelPushOptionOptionTypeEnum = "File"
	GetCreateDirectIncomesModelPushOptionOptionTypeEnumMultiPart GetCreateDirectIncomesModelPushOptionOptionTypeEnum = "MultiPart"
)

type GetCreateDirectIncomesModelPushOption struct {
	Description *string                                                            `json:"description,omitempty"`
	DisplayName string                                                             `json:"displayName"`
	Properties  map[string]GetCreateDirectIncomesModelPushOptionPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                               `json:"required"`
	Type        GetCreateDirectIncomesModelPushOptionOptionTypeEnum                `json:"type"`
}

type GetCreateDirectIncomesModelResponse struct {
	ContentType string
	PushOption  *GetCreateDirectIncomesModelPushOption
	StatusCode  int
	RawResponse *http.Response
}
