package shared

import (
	"github.com/codatio/client-sdk-go/assess/pkg/types"
)

// FinancialMetricErrorsAssessErrorDetails
// Dictionary list outlining the missing properties or allowed values.
type FinancialMetricErrorsAssessErrorDetails struct {
	PropertyDetail1 []string `json:"propertyDetail1,omitempty"`
	PropertyDetail2 []string `json:"propertyDetail2,omitempty"`
	PropertyDetailN []string `json:"propertyDetailN,omitempty"`
}

type FinancialMetricErrorsTypeEnum string

const (
	FinancialMetricErrorsTypeEnumUncategorizedAccounts FinancialMetricErrorsTypeEnum = "UncategorizedAccounts"
	FinancialMetricErrorsTypeEnumMissingInputData      FinancialMetricErrorsTypeEnum = "MissingInputData"
	FinancialMetricErrorsTypeEnumInputDataError        FinancialMetricErrorsTypeEnum = "InputDataError"
)

type FinancialMetricErrors struct {
	Details *FinancialMetricErrorsAssessErrorDetails `json:"details,omitempty"`
	Message *string                                  `json:"message,omitempty"`
	Type    *FinancialMetricErrorsTypeEnum           `json:"type,omitempty"`
}

type FinancialMetricKeyEnum string

const (
	FinancialMetricKeyEnumUnknown                     FinancialMetricKeyEnum = "Unknown"
	FinancialMetricKeyEnumEbitda                      FinancialMetricKeyEnum = "EBITDA"
	FinancialMetricKeyEnumDebtServiceCoverageRatio    FinancialMetricKeyEnum = "DebtServiceCoverageRatio"
	FinancialMetricKeyEnumCurrentRatioQuickRatio      FinancialMetricKeyEnum = "CurrentRatio QuickRatio"
	FinancialMetricKeyEnumGrossProfitMargin           FinancialMetricKeyEnum = "GrossProfitMargin"
	FinancialMetricKeyEnumFixedChargeCoverageRatio    FinancialMetricKeyEnum = "FixedChargeCoverageRatio"
	FinancialMetricKeyEnumWorkingCapital              FinancialMetricKeyEnum = "WorkingCapital"
	FinancialMetricKeyEnumFreeCashFlow                FinancialMetricKeyEnum = "FreeCashFlow"
	FinancialMetricKeyEnumNetProfitMargin             FinancialMetricKeyEnum = "NetProfitMargin"
	FinancialMetricKeyEnumReturnOnAssetsRatio         FinancialMetricKeyEnum = "ReturnOnAssetsRatio"
	FinancialMetricKeyEnumReturnOnEquityRatio         FinancialMetricKeyEnum = "ReturnOnEquityRatio"
	FinancialMetricKeyEnumOperatingProfitMargin       FinancialMetricKeyEnum = "OperatingProfitMargin"
	FinancialMetricKeyEnumDeptToEquity                FinancialMetricKeyEnum = "DeptToEquity"
	FinancialMetricKeyEnumDebtToAssets                FinancialMetricKeyEnum = "DebtToAssets"
	FinancialMetricKeyEnumInterestCoverageRatio       FinancialMetricKeyEnum = "InterestCoverageRatio"
	FinancialMetricKeyEnumCashRatio                   FinancialMetricKeyEnum = "CashRatio"
	FinancialMetricKeyEnumInventoryTurnoverRatio      FinancialMetricKeyEnum = "InventoryTurnoverRatio"
	FinancialMetricKeyEnumAssetTurnoverRatio          FinancialMetricKeyEnum = "AssetTurnoverRatio"
	FinancialMetricKeyEnumWorkingCapitalTurnoverRatio FinancialMetricKeyEnum = "WorkingCapitalTurnoverRatio"
	FinancialMetricKeyEnumDaysSalesOutstanding        FinancialMetricKeyEnum = "DaysSalesOutstanding"
	FinancialMetricKeyEnumDaysPayablesOutstanding     FinancialMetricKeyEnum = "DaysPayablesOutstanding"
)

type FinancialMetricMetricUnitEnum string

const (
	FinancialMetricMetricUnitEnumRatio FinancialMetricMetricUnitEnum = "Ratio"
	FinancialMetricMetricUnitEnumMoney FinancialMetricMetricUnitEnum = "Money"
)

type FinancialMetricPeriodsErrorsTypeEnum string

const (
	FinancialMetricPeriodsErrorsTypeEnumMissingAccountData FinancialMetricPeriodsErrorsTypeEnum = "MissingAccountData"
	FinancialMetricPeriodsErrorsTypeEnumDatesOutOfRange    FinancialMetricPeriodsErrorsTypeEnum = "DatesOutOfRange"
)

type FinancialMetricPeriodsErrors struct {
	Details *Details                              `json:"details,omitempty"`
	Massage *string                               `json:"massage,omitempty"`
	Type    *FinancialMetricPeriodsErrorsTypeEnum `json:"type,omitempty"`
}

type FinancialMetricPeriodsInputs struct {
	Name  *string  `json:"name,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

type FinancialMetricPeriods struct {
	Errors   []FinancialMetricPeriodsErrors `json:"errors,omitempty"`
	FromDate *types.Date                    `json:"fromDate,omitempty"`
	Inputs   []FinancialMetricPeriodsInputs `json:"inputs,omitempty"`
	ToDate   *types.Date                    `json:"toDate,omitempty"`
	Value    *float64                       `json:"value,omitempty"`
}

type FinancialMetric struct {
	Errors     []FinancialMetricErrors        `json:"errors,omitempty"`
	Key        *FinancialMetricKeyEnum        `json:"key,omitempty"`
	MetricUnit *FinancialMetricMetricUnitEnum `json:"metricUnit,omitempty"`
	Name       *string                        `json:"name,omitempty"`
	Periods    []FinancialMetricPeriods       `json:"periods,omitempty"`
}
