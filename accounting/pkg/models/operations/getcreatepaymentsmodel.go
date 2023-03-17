package operations

import (
	"net/http"
)

type GetCreatePaymentsModelRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                           `json:"description,omitempty"`
	DisplayName *string                                                                           `json:"displayName,omitempty"`
	Required    *bool                                                                             `json:"required,omitempty"`
	Type        *GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                           `json:"value,omitempty"`
}

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                             `json:"description,omitempty"`
	DisplayName *string                                                                                             `json:"displayName,omitempty"`
	Required    *bool                                                                                               `json:"required,omitempty"`
	Type        *GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                             `json:"value,omitempty"`
}

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                               `json:"description,omitempty"`
	DisplayName *string                                                                                                               `json:"displayName,omitempty"`
	Required    *bool                                                                                                                 `json:"required,omitempty"`
	Type        *GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                               `json:"value,omitempty"`
}

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                                 `json:"description,omitempty"`
	DisplayName *string                                                                                                                                 `json:"displayName,omitempty"`
	Required    *bool                                                                                                                                   `json:"required,omitempty"`
	Type        *GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                                 `json:"value,omitempty"`
}

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                      `json:"description"`
	DisplayName string                                                                                                                      `json:"displayName"`
	Options     []GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice  `json:"options,omitempty"`
	Required    bool                                                                                                                        `json:"required"`
	Type        GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum      `json:"type"`
	Validation  *GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo `json:"validation,omitempty"`
}

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                              `json:"description"`
	DisplayName string                                                                                                              `json:"displayName"`
	Options     []GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                                `json:"required"`
	Type        GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                            `json:"description"`
	DisplayName string                                                                                            `json:"displayName"`
	Options     []GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                              `json:"required"`
	Type        GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnum string

const (
	GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnumArray     GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnum = "Array"
	GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnumObject    GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnum = "Object"
	GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnumString    GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnum = "String"
	GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnumNumber    GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnum = "Number"
	GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnumBoolean   GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnumDateTime  GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnumFile      GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnum = "File"
	GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnumMultiPart GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreatePaymentsModelPushOptionPushOptionPropertyPushValidationInfo struct {
	Information []GetCreatePaymentsModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreatePaymentsModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreatePaymentsModelPushOptionPushOptionProperty struct {
	Description string                                                                          `json:"description"`
	DisplayName string                                                                          `json:"displayName"`
	Options     []GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreatePaymentsModelPushOptionPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                            `json:"required"`
	Type        GetCreatePaymentsModelPushOptionPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreatePaymentsModelPushOptionPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreatePaymentsModelPushOptionOptionTypeEnum string

const (
	GetCreatePaymentsModelPushOptionOptionTypeEnumArray     GetCreatePaymentsModelPushOptionOptionTypeEnum = "Array"
	GetCreatePaymentsModelPushOptionOptionTypeEnumObject    GetCreatePaymentsModelPushOptionOptionTypeEnum = "Object"
	GetCreatePaymentsModelPushOptionOptionTypeEnumString    GetCreatePaymentsModelPushOptionOptionTypeEnum = "String"
	GetCreatePaymentsModelPushOptionOptionTypeEnumNumber    GetCreatePaymentsModelPushOptionOptionTypeEnum = "Number"
	GetCreatePaymentsModelPushOptionOptionTypeEnumBoolean   GetCreatePaymentsModelPushOptionOptionTypeEnum = "Boolean"
	GetCreatePaymentsModelPushOptionOptionTypeEnumDateTime  GetCreatePaymentsModelPushOptionOptionTypeEnum = "DateTime"
	GetCreatePaymentsModelPushOptionOptionTypeEnumFile      GetCreatePaymentsModelPushOptionOptionTypeEnum = "File"
	GetCreatePaymentsModelPushOptionOptionTypeEnumMultiPart GetCreatePaymentsModelPushOptionOptionTypeEnum = "MultiPart"
)

type GetCreatePaymentsModelPushOption struct {
	Description *string                                                       `json:"description,omitempty"`
	DisplayName string                                                        `json:"displayName"`
	Properties  map[string]GetCreatePaymentsModelPushOptionPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                          `json:"required"`
	Type        GetCreatePaymentsModelPushOptionOptionTypeEnum                `json:"type"`
}

type GetCreatePaymentsModelResponse struct {
	ContentType string
	PushOption  *GetCreatePaymentsModelPushOption
	StatusCode  int
	RawResponse *http.Response
}
