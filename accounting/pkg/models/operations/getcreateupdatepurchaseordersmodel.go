package operations

import (
	"net/http"
)

type GetCreateUpdatePurchaseOrdersModelRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                       `json:"description,omitempty"`
	DisplayName *string                                                                                       `json:"displayName,omitempty"`
	Required    *bool                                                                                         `json:"required,omitempty"`
	Type        *GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                       `json:"value,omitempty"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                         `json:"description,omitempty"`
	DisplayName *string                                                                                                         `json:"displayName,omitempty"`
	Required    *bool                                                                                                           `json:"required,omitempty"`
	Type        *GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                         `json:"value,omitempty"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                           `json:"description,omitempty"`
	DisplayName *string                                                                                                                           `json:"displayName,omitempty"`
	Required    *bool                                                                                                                             `json:"required,omitempty"`
	Type        *GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                           `json:"value,omitempty"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                                             `json:"description,omitempty"`
	DisplayName *string                                                                                                                                             `json:"displayName,omitempty"`
	Required    *bool                                                                                                                                               `json:"required,omitempty"`
	Type        *GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                                             `json:"value,omitempty"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                                  `json:"description"`
	DisplayName string                                                                                                                                  `json:"displayName"`
	Options     []GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice  `json:"options,omitempty"`
	Required    bool                                                                                                                                    `json:"required"`
	Type        GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum      `json:"type"`
	Validation  *GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo `json:"validation,omitempty"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                          `json:"description"`
	DisplayName string                                                                                                                          `json:"displayName"`
	Options     []GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                                            `json:"required"`
	Type        GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                        `json:"description"`
	DisplayName string                                                                                                        `json:"displayName"`
	Options     []GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                          `json:"required"`
	Type        GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnum string

const (
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnumArray     GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnum = "Array"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnumObject    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnum = "Object"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnumString    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnum = "String"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnumNumber    GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnum = "Number"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnumBoolean   GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnumDateTime  GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnumFile      GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnum = "File"
	GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnumMultiPart GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushValidationInfo struct {
	Information []GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionProperty struct {
	Description string                                                                                      `json:"description"`
	DisplayName string                                                                                      `json:"displayName"`
	Options     []GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                        `json:"required"`
	Type        GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnum string

const (
	GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnumArray     GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnum = "Array"
	GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnumObject    GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnum = "Object"
	GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnumString    GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnum = "String"
	GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnumNumber    GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnum = "Number"
	GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnumBoolean   GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnum = "Boolean"
	GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnumDateTime  GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnum = "DateTime"
	GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnumFile      GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnum = "File"
	GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnumMultiPart GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnum = "MultiPart"
)

type GetCreateUpdatePurchaseOrdersModelPushOption struct {
	Description *string                                                                   `json:"description,omitempty"`
	DisplayName string                                                                    `json:"displayName"`
	Properties  map[string]GetCreateUpdatePurchaseOrdersModelPushOptionPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                      `json:"required"`
	Type        GetCreateUpdatePurchaseOrdersModelPushOptionOptionTypeEnum                `json:"type"`
}

type GetCreateUpdatePurchaseOrdersModelResponse struct {
	ContentType string
	PushOption  *GetCreateUpdatePurchaseOrdersModelPushOption
	StatusCode  int
	RawResponse *http.Response
}
