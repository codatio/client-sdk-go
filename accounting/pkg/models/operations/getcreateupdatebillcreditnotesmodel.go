package operations

import (
	"net/http"
)

type GetCreateUpdateBillCreditNotesModelRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                        `json:"description,omitempty"`
	DisplayName *string                                                                                        `json:"displayName,omitempty"`
	Required    *bool                                                                                          `json:"required,omitempty"`
	Type        *GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                        `json:"value,omitempty"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                          `json:"description,omitempty"`
	DisplayName *string                                                                                                          `json:"displayName,omitempty"`
	Required    *bool                                                                                                            `json:"required,omitempty"`
	Type        *GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                          `json:"value,omitempty"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                            `json:"description,omitempty"`
	DisplayName *string                                                                                                                            `json:"displayName,omitempty"`
	Required    *bool                                                                                                                              `json:"required,omitempty"`
	Type        *GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                            `json:"value,omitempty"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                                              `json:"description,omitempty"`
	DisplayName *string                                                                                                                                              `json:"displayName,omitempty"`
	Required    *bool                                                                                                                                                `json:"required,omitempty"`
	Type        *GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                                              `json:"value,omitempty"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                                   `json:"description"`
	DisplayName string                                                                                                                                   `json:"displayName"`
	Options     []GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice  `json:"options,omitempty"`
	Required    bool                                                                                                                                     `json:"required"`
	Type        GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum      `json:"type"`
	Validation  *GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo `json:"validation,omitempty"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                           `json:"description"`
	DisplayName string                                                                                                                           `json:"displayName"`
	Options     []GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                                             `json:"required"`
	Type        GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                         `json:"description"`
	DisplayName string                                                                                                         `json:"displayName"`
	Options     []GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                           `json:"required"`
	Type        GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnum string

const (
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnumArray     GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnumObject    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnumString    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnum = "String"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnumNumber    GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnumBoolean   GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnumDateTime  GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnumFile      GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnum = "File"
	GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnumMultiPart GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionPushOptionProperty struct {
	Description string                                                                                       `json:"description"`
	DisplayName string                                                                                       `json:"displayName"`
	Options     []GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                         `json:"required"`
	Type        GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateUpdateBillCreditNotesModelPushOptionPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnum string

const (
	GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnumArray     GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnum = "Array"
	GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnumObject    GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnum = "Object"
	GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnumString    GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnum = "String"
	GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnumNumber    GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnum = "Number"
	GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnumBoolean   GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnum = "Boolean"
	GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnumDateTime  GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnum = "DateTime"
	GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnumFile      GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnum = "File"
	GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnumMultiPart GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnum = "MultiPart"
)

type GetCreateUpdateBillCreditNotesModelPushOption struct {
	Description *string                                                                    `json:"description,omitempty"`
	DisplayName string                                                                     `json:"displayName"`
	Properties  map[string]GetCreateUpdateBillCreditNotesModelPushOptionPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                       `json:"required"`
	Type        GetCreateUpdateBillCreditNotesModelPushOptionOptionTypeEnum                `json:"type"`
}

type GetCreateUpdateBillCreditNotesModelResponse struct {
	ContentType string
	PushOption  *GetCreateUpdateBillCreditNotesModelPushOption
	StatusCode  int
	RawResponse *http.Response
}
