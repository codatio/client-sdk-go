package operations

import (
	"net/http"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumAccountTransactions          GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "accountTransactions"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumBalanceSheet                 GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "balanceSheet"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumBankAccounts                 GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "bankAccounts"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumBankTransactions             GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "bankTransactions"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumBillCreditNotes              GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "billCreditNotes"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumBillPayments                 GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "billPayments"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumBills                        GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "bills"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCashFlowStatement            GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "cashFlowStatement"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumChartOfAccounts              GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "chartOfAccounts"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCompany                      GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "company"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCreditNotes                  GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "creditNotes"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCustomers                    GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "customers"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumDirectCosts                  GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "directCosts"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumDirectIncomes                GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "directIncomes"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumInvoices                     GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "invoices"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumItems                        GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "items"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumJournalEntries               GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "journalEntries"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumJournals                     GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "journals"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumPaymentMethods               GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "paymentMethods"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumPayments                     GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "payments"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumProfitAndLoss                GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "profitAndLoss"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumProjects                     GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "projects"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumPurchaseOrders               GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "purchaseOrders"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumSalesOrders                  GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "salesOrders"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumSuppliers                    GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "suppliers"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumTaxRates                     GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "taxRates"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumTrackingCategories           GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "trackingCategories"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumTransfers                    GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "transfers"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumBankingAccountBalances       GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "banking-accountBalances"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumBankingAccounts              GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "banking-accounts"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumBankingTransactionCategories GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "banking-transactionCategories"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumBankingTransactions          GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "banking-transactions"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCommerceCompanyInfo          GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "commerce-companyInfo"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCommerceCustomers            GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "commerce-customers"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCommerceDisputes             GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "commerce-disputes"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCommerceLocations            GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "commerce-locations"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCommerceOrders               GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "commerce-orders"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCommercePaymentMethods       GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "commerce-paymentMethods"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCommercePayments             GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "commerce-payments"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCommerceProductCategories    GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "commerce-productCategories"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCommerceProducts             GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "commerce-products"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCommerceTaxComponents        GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "commerce-taxComponents"
	GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnumCommerceTransactions         GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum = "commerce-transactions"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushRequest struct {
	CompanyID    string                                                       `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string                                                       `pathParam:"style=simple,explode=false,name=connectionId"`
	DataType     GetCompaniesCompanyIDConnectionsConnectionIDPushDataTypeEnum `pathParam:"style=simple,explode=false,name=dataType"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                     `json:"description,omitempty"`
	DisplayName *string                                                                                                     `json:"displayName,omitempty"`
	Required    *bool                                                                                                       `json:"required,omitempty"`
	Type        *GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                     `json:"value,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                       `json:"description,omitempty"`
	DisplayName *string                                                                                                                       `json:"displayName,omitempty"`
	Required    *bool                                                                                                                         `json:"required,omitempty"`
	Type        *GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                       `json:"value,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                                         `json:"description,omitempty"`
	DisplayName *string                                                                                                                                         `json:"displayName,omitempty"`
	Required    *bool                                                                                                                                           `json:"required,omitempty"`
	Type        *GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                                         `json:"value,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumArray     GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Array"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumObject    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Object"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumString    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "String"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumNumber    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Number"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumBoolean   GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "Boolean"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumDateTime  GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "DateTime"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumFile      GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "File"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnumMultiPart GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum = "MultiPart"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice struct {
	Description *string                                                                                                                                                           `json:"description,omitempty"`
	DisplayName *string                                                                                                                                                           `json:"displayName,omitempty"`
	Required    *bool                                                                                                                                                             `json:"required,omitempty"`
	Type        *GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoiceOptionTypeEnum `json:"type,omitempty"`
	Value       *string                                                                                                                                                           `json:"value,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                                                `json:"description"`
	DisplayName string                                                                                                                                                `json:"displayName"`
	Options     []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice  `json:"options,omitempty"`
	Required    bool                                                                                                                                                  `json:"required"`
	Type        GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum      `json:"type"`
	Validation  *GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo `json:"validation,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                                        `json:"description"`
	DisplayName string                                                                                                                                        `json:"displayName"`
	Options     []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                                                          `json:"required"`
	Type        GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumArray     GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Array"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumObject    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Object"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumString    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "String"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumNumber    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Number"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumBoolean   GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumDateTime  GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumFile      GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "File"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnumMultiPart GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo struct {
	Information []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionProperty struct {
	Description string                                                                                                                      `json:"description"`
	DisplayName string                                                                                                                      `json:"displayName"`
	Options     []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                                        `json:"required"`
	Type        GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnumArray     GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnum = "Array"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnumObject    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnum = "Object"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnumString    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnum = "String"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnumNumber    GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnum = "Number"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnumBoolean   GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnum = "Boolean"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnumDateTime  GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnum = "DateTime"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnumFile      GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnum = "File"
	GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnumMultiPart GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnum = "MultiPart"
)

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation struct {
	Details string  `json:"details"`
	Field   *string `json:"field,omitempty"`
	Ref     *string `json:"ref,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushValidationInfo struct {
	Information []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"information,omitempty"`
	Warnings    []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushValidationInfoPushFieldValidation `json:"warnings,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionProperty struct {
	Description string                                                                                                    `json:"description"`
	DisplayName string                                                                                                    `json:"displayName"`
	Options     []GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionChoice            `json:"options,omitempty"`
	Properties  map[string]GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                                      `json:"required"`
	Type        GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyOptionTypeEnum                `json:"type"`
	Validation  *GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionPropertyPushValidationInfo           `json:"validation,omitempty"`
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

type GetCompaniesCompanyIDConnectionsConnectionIDPushPushOption struct {
	Description *string                                                                                 `json:"description,omitempty"`
	DisplayName string                                                                                  `json:"displayName"`
	Properties  map[string]GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionPushOptionProperty `json:"properties,omitempty"`
	Required    bool                                                                                    `json:"required"`
	Type        GetCompaniesCompanyIDConnectionsConnectionIDPushPushOptionOptionTypeEnum                `json:"type"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDPushResponse struct {
	ContentType string
	PushOption  *GetCompaniesCompanyIDConnectionsConnectionIDPushPushOption
	StatusCode  int
	RawResponse *http.Response
}
