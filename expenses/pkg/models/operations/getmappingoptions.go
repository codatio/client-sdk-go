package operations

import (
	"time"
)

type GetMappingOptionsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetMappingOptionsRequest struct {
	PathParams GetMappingOptionsPathParams
}

type GetMappingOptions200TextJSONAccountsAccountTypeEnum string

const (
	GetMappingOptions200TextJSONAccountsAccountTypeEnumAsset     GetMappingOptions200TextJSONAccountsAccountTypeEnum = "Asset"
	GetMappingOptions200TextJSONAccountsAccountTypeEnumLiability GetMappingOptions200TextJSONAccountsAccountTypeEnum = "Liability"
	GetMappingOptions200TextJSONAccountsAccountTypeEnumIncome    GetMappingOptions200TextJSONAccountsAccountTypeEnum = "Income"
	GetMappingOptions200TextJSONAccountsAccountTypeEnumExpense   GetMappingOptions200TextJSONAccountsAccountTypeEnum = "Expense"
	GetMappingOptions200TextJSONAccountsAccountTypeEnumEquity    GetMappingOptions200TextJSONAccountsAccountTypeEnum = "Equity"
)

type GetMappingOptions200TextJSONAccountsValidTransactionTypesEnum string

const (
	GetMappingOptions200TextJSONAccountsValidTransactionTypesEnumPayment       GetMappingOptions200TextJSONAccountsValidTransactionTypesEnum = "Payment"
	GetMappingOptions200TextJSONAccountsValidTransactionTypesEnumRefund        GetMappingOptions200TextJSONAccountsValidTransactionTypesEnum = "Refund"
	GetMappingOptions200TextJSONAccountsValidTransactionTypesEnumReward        GetMappingOptions200TextJSONAccountsValidTransactionTypesEnum = "Reward"
	GetMappingOptions200TextJSONAccountsValidTransactionTypesEnumChargeback    GetMappingOptions200TextJSONAccountsValidTransactionTypesEnum = "Chargeback"
	GetMappingOptions200TextJSONAccountsValidTransactionTypesEnumTransferIn    GetMappingOptions200TextJSONAccountsValidTransactionTypesEnum = "TransferIn"
	GetMappingOptions200TextJSONAccountsValidTransactionTypesEnumTransferOut   GetMappingOptions200TextJSONAccountsValidTransactionTypesEnum = "TransferOut"
	GetMappingOptions200TextJSONAccountsValidTransactionTypesEnumAdjustmentIn  GetMappingOptions200TextJSONAccountsValidTransactionTypesEnum = "AdjustmentIn"
	GetMappingOptions200TextJSONAccountsValidTransactionTypesEnumAdjustmentOut GetMappingOptions200TextJSONAccountsValidTransactionTypesEnum = "AdjustmentOut"
)

type GetMappingOptions200TextJSONAccounts struct {
	AccountType           *GetMappingOptions200TextJSONAccountsAccountTypeEnum            `json:"accountType,omitempty"`
	Currency              *string                                                         `json:"currency,omitempty"`
	ID                    *string                                                         `json:"id,omitempty"`
	Name                  *string                                                         `json:"name,omitempty"`
	ValidTransactionTypes []GetMappingOptions200TextJSONAccountsValidTransactionTypesEnum `json:"validTransactionTypes,omitempty"`
}

type GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnum string

const (
	GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnumPayment       GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnum = "Payment"
	GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnumRefund        GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnum = "Refund"
	GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnumReward        GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnum = "Reward"
	GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnumChargeback    GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnum = "Chargeback"
	GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnumTransferIn    GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnum = "TransferIn"
	GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnumTransferOut   GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnum = "TransferOut"
	GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnumAdjustmentIn  GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnum = "AdjustmentIn"
	GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnumAdjustmentOut GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnum = "AdjustmentOut"
)

type GetMappingOptions200TextJSONTaxRates struct {
	Code                  *string                                                         `json:"code,omitempty"`
	EffectiveTaxRate      *float64                                                        `json:"effectiveTaxRate,omitempty"`
	ID                    *string                                                         `json:"id,omitempty"`
	Name                  *string                                                         `json:"name,omitempty"`
	TotalTaxRate          *float64                                                        `json:"totalTaxRate,omitempty"`
	ValidTransactionTypes []GetMappingOptions200TextJSONTaxRatesValidTransactionTypesEnum `json:"validTransactionTypes,omitempty"`
}

type GetMappingOptions200TextJSONTrackingCategories struct {
	HasChildren  *bool      `json:"hasChildren,omitempty"`
	ID           *string    `json:"id,omitempty"`
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	Name         *string    `json:"name,omitempty"`
	ParentID     *string    `json:"parentId,omitempty"`
}

type GetMappingOptions200TextJSON struct {
	Accounts           []GetMappingOptions200TextJSONAccounts           `json:"accounts,omitempty"`
	ExpenseProvider    *string                                          `json:"expenseProvider,omitempty"`
	TaxRates           []GetMappingOptions200TextJSONTaxRates           `json:"taxRates,omitempty"`
	TrackingCategories []GetMappingOptions200TextJSONTrackingCategories `json:"trackingCategories,omitempty"`
}

type GetMappingOptions200ApplicationJSONAccountsAccountTypeEnum string

const (
	GetMappingOptions200ApplicationJSONAccountsAccountTypeEnumAsset     GetMappingOptions200ApplicationJSONAccountsAccountTypeEnum = "Asset"
	GetMappingOptions200ApplicationJSONAccountsAccountTypeEnumLiability GetMappingOptions200ApplicationJSONAccountsAccountTypeEnum = "Liability"
	GetMappingOptions200ApplicationJSONAccountsAccountTypeEnumIncome    GetMappingOptions200ApplicationJSONAccountsAccountTypeEnum = "Income"
	GetMappingOptions200ApplicationJSONAccountsAccountTypeEnumExpense   GetMappingOptions200ApplicationJSONAccountsAccountTypeEnum = "Expense"
	GetMappingOptions200ApplicationJSONAccountsAccountTypeEnumEquity    GetMappingOptions200ApplicationJSONAccountsAccountTypeEnum = "Equity"
)

type GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnum string

const (
	GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnumPayment       GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnum = "Payment"
	GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnumRefund        GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnum = "Refund"
	GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnumReward        GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnum = "Reward"
	GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnumChargeback    GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnum = "Chargeback"
	GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnumTransferIn    GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnum = "TransferIn"
	GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnumTransferOut   GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnum = "TransferOut"
	GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnumAdjustmentIn  GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnum = "AdjustmentIn"
	GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnumAdjustmentOut GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnum = "AdjustmentOut"
)

type GetMappingOptions200ApplicationJSONAccounts struct {
	AccountType           *GetMappingOptions200ApplicationJSONAccountsAccountTypeEnum            `json:"accountType,omitempty"`
	Currency              *string                                                                `json:"currency,omitempty"`
	ID                    *string                                                                `json:"id,omitempty"`
	Name                  *string                                                                `json:"name,omitempty"`
	ValidTransactionTypes []GetMappingOptions200ApplicationJSONAccountsValidTransactionTypesEnum `json:"validTransactionTypes,omitempty"`
}

type GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnum string

const (
	GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnumPayment       GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnum = "Payment"
	GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnumRefund        GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnum = "Refund"
	GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnumReward        GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnum = "Reward"
	GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnumChargeback    GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnum = "Chargeback"
	GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnumTransferIn    GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnum = "TransferIn"
	GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnumTransferOut   GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnum = "TransferOut"
	GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnumAdjustmentIn  GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnum = "AdjustmentIn"
	GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnumAdjustmentOut GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnum = "AdjustmentOut"
)

type GetMappingOptions200ApplicationJSONTaxRates struct {
	Code                  *string                                                                `json:"code,omitempty"`
	EffectiveTaxRate      *float64                                                               `json:"effectiveTaxRate,omitempty"`
	ID                    *string                                                                `json:"id,omitempty"`
	Name                  *string                                                                `json:"name,omitempty"`
	TotalTaxRate          *float64                                                               `json:"totalTaxRate,omitempty"`
	ValidTransactionTypes []GetMappingOptions200ApplicationJSONTaxRatesValidTransactionTypesEnum `json:"validTransactionTypes,omitempty"`
}

type GetMappingOptions200ApplicationJSONTrackingCategories struct {
	HasChildren  *bool      `json:"hasChildren,omitempty"`
	ID           *string    `json:"id,omitempty"`
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	Name         *string    `json:"name,omitempty"`
	ParentID     *string    `json:"parentId,omitempty"`
}

type GetMappingOptions200ApplicationJSON struct {
	Accounts           []GetMappingOptions200ApplicationJSONAccounts           `json:"accounts,omitempty"`
	ExpenseProvider    *string                                                 `json:"expenseProvider,omitempty"`
	TaxRates           []GetMappingOptions200ApplicationJSONTaxRates           `json:"taxRates,omitempty"`
	TrackingCategories []GetMappingOptions200ApplicationJSONTrackingCategories `json:"trackingCategories,omitempty"`
}

type GetMappingOptionsResponse struct {
	ContentType                               string
	StatusCode                                int
	GetMappingOptions200ApplicationJSONObject *GetMappingOptions200ApplicationJSON
	GetMappingOptions200TextJSONObject        *GetMappingOptions200TextJSON
	GetMappingOptions200TextPlainObject       *string
}
