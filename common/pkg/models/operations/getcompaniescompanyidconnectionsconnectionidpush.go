package operations

type GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumInvoices             GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "invoices"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumAccounts             GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "accounts"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCommercePayments     GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "commerce-payments"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumBankingAccounts      GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "banking-accounts"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCompany              GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "company"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumProfitAndLoss        GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "profitAndLoss"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCommerceTransactions GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "commerce-transactions"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumBills                GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "bills"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCustomers            GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "customers"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushPathParams struct {
	CompanyID    string                                                       `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string                                                       `pathParam:"style=simple,explode=false,name=connectionId"`
	DataType     GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum `pathParam:"style=simple,explode=false,name=dataType"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushRequest struct {
	PathParams GetCompaniesCompanyIDConnectionsConnectionIDPushPathParams
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumArray     GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "Array"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumObject    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "Object"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumString    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "String"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumNumber    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "Number"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumBoolean   GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "Boolean"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumDateTime  GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "DateTime"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumFile      GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "File"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumMultiPart GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   string  `json:"field"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfo struct {
	Information []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoice struct {
	Description string                                                                                                                                        `json:"description"`
	DisplayName string                                                                                                                                        `json:"displayName"`
	Rel         *string                                                                                                                                       `json:"rel,omitempty"`
	Required    bool                                                                                                                                          `json:"required"`
	Type        GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum      `json:"type"`
	Validation  *GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfo `json:"validation,omitempty"`
	Value       string                                                                                                                                        `json:"value"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumArray     GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "Array"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumObject    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "Object"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumString    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "String"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumNumber    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "Number"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumBoolean   GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "Boolean"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumDateTime  GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "DateTime"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumFile      GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "File"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnumMultiPart GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   string  `json:"field"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfo struct {
	Information []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoice struct {
	Description string                                                                                                                        `json:"description"`
	DisplayName string                                                                                                                        `json:"displayName"`
	Options     []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushOptionChoice  `json:"options,omitempty"`
	Rel         *string                                                                                                                       `json:"rel,omitempty"`
	Required    bool                                                                                                                          `json:"required"`
	Type        GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoiceOptionTypeEnum      `json:"type"`
	Validation  *GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoicePushValidationInfo `json:"validation,omitempty"`
	Value       string                                                                                                                        `json:"value"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnumArray     GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum = "Array"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnumObject    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum = "Object"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnumString    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum = "String"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnumNumber    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum = "Number"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnumBoolean   GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum = "Boolean"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnumDateTime  GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum = "DateTime"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnumFile      GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum = "File"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnumMultiPart GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   string  `json:"field"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushValidationInfo struct {
	Information []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoice struct {
	Description string                                                                                                        `json:"description"`
	DisplayName string                                                                                                        `json:"displayName"`
	Options     []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushOptionChoice  `json:"options,omitempty"`
	Rel         *string                                                                                                       `json:"rel,omitempty"`
	Required    bool                                                                                                          `json:"required"`
	Type        GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoiceOptionTypeEnum      `json:"type"`
	Validation  *GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoicePushValidationInfo `json:"validation,omitempty"`
	Value       string                                                                                                        `json:"value"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnumArray     GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnum = "Array"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnumObject    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnum = "Object"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnumString    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnum = "String"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnumNumber    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnum = "Number"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnumBoolean   GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnumDateTime  GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnumFile      GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnum = "File"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnumMultiPart GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   string  `json:"field"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushValidationInfo struct {
	Information []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoice struct {
	Description string                                                                                        `json:"description"`
	DisplayName string                                                                                        `json:"displayName"`
	Options     []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushOptionChoice  `json:"options,omitempty"`
	Rel         *string                                                                                       `json:"rel,omitempty"`
	Required    bool                                                                                          `json:"required"`
	Type        GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoiceOptionTypeEnum      `json:"type"`
	Validation  *GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionChoicePushValidationInfo `json:"validation,omitempty"`
	Value       string                                                                                        `json:"value"`
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
	Information []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
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
