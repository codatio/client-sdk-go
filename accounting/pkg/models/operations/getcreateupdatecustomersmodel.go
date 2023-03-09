package operations

import (
	"net/http"
)

type GetCreateUpdateCustomersModelPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCreateUpdateCustomersModelRequest struct {
	PathParams GetCreateUpdateCustomersModelPathParams
}

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                  `json:"description,omitempty"`
	DisplayName *string                                                                                  `json:"displayName,omitempty"`
	Required    *bool                                                                                    `json:"required,omitempty"`
	Type        *GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                  `json:"value,omitempty"`
}

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                    `json:"description,omitempty"`
	DisplayName *string                                                                                                    `json:"displayName,omitempty"`
	Required    *bool                                                                                                      `json:"required,omitempty"`
	Type        *GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                    `json:"value,omitempty"`
}

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                      `json:"description,omitempty"`
	DisplayName *string                                                                                                                      `json:"displayName,omitempty"`
	Required    *bool                                                                                                                        `json:"required,omitempty"`
	Type        *GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                      `json:"value,omitempty"`
}

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                                        `json:"description,omitempty"`
	DisplayName *string                                                                                                                                        `json:"displayName,omitempty"`
	Required    *bool                                                                                                                                          `json:"required,omitempty"`
	Type        *GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                                        `json:"value,omitempty"`
}

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                             `json:"description"`
	DisplayName string                                                                                                                             `json:"displayName"`
	Options     []GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice  `json:"options,omitempty"`
	Required    bool                                                                                                                               `json:"required"`
	Type        GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum      `json:"type"`
	Validation  *GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo `json:"validation,omitempty"`
}

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                     `json:"description"`
	DisplayName string                                                                                                                     `json:"displayName"`
	Options     []GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                                       `json:"required"`
	Type        GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                   `json:"description"`
	DisplayName string                                                                                                   `json:"displayName"`
	Options     []GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                     `json:"required"`
	Type        GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnum string

const (
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnumArray     GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnumObject    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnumString    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnum = "String"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnumNumber    GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnumBoolean   GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnumDateTime  GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnumFile      GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnum = "File"
	GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnumMultiPart GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateUpdateCustomersModelPushOptionPushOptionProperty struct {
	Description string                                                                                 `json:"description"`
	DisplayName string                                                                                 `json:"displayName"`
	Options     []GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                   `json:"required"`
	Type        GetCreateUpdateCustomersModelPushOptionPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateUpdateCustomersModelPushOptionPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateUpdateCustomersModelPushOptionOptionTypeEnum string

const (
	GetCreateUpdateCustomersModelPushOptionOptionTypeEnumArray     GetCreateUpdateCustomersModelPushOptionOptionTypeEnum = "Array"
	GetCreateUpdateCustomersModelPushOptionOptionTypeEnumObject    GetCreateUpdateCustomersModelPushOptionOptionTypeEnum = "Object"
	GetCreateUpdateCustomersModelPushOptionOptionTypeEnumString    GetCreateUpdateCustomersModelPushOptionOptionTypeEnum = "String"
	GetCreateUpdateCustomersModelPushOptionOptionTypeEnumNumber    GetCreateUpdateCustomersModelPushOptionOptionTypeEnum = "Number"
	GetCreateUpdateCustomersModelPushOptionOptionTypeEnumBoolean   GetCreateUpdateCustomersModelPushOptionOptionTypeEnum = "Boolean"
	GetCreateUpdateCustomersModelPushOptionOptionTypeEnumDateTime  GetCreateUpdateCustomersModelPushOptionOptionTypeEnum = "DateTime"
	GetCreateUpdateCustomersModelPushOptionOptionTypeEnumFile      GetCreateUpdateCustomersModelPushOptionOptionTypeEnum = "File"
	GetCreateUpdateCustomersModelPushOptionOptionTypeEnumMultiPart GetCreateUpdateCustomersModelPushOptionOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateCustomersModelPushOption struct {
	Description *string                                                              `json:"description,omitempty"`
	DisplayName string                                                               `json:"displayName"`
	Properties  map[string]GetCreateUpdateCustomersModelPushOptionPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                 `json:"required"`
	Type        GetCreateUpdateCustomersModelPushOptionOptionTypeEnum                `json:"type"`
}

type GetCreateUpdateCustomersModelResponse struct {
	ContentType string
	PushOption  *GetCreateUpdateCustomersModelPushOption
	StatusCode  int
	RawResponse *http.Response
}
