package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/types"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetricsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetricsQueryParams struct {
	NumberOfPeriods  int64  `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	PeriodLength     int64  `queryParam:"style=form,explode=true,name=periodLength"`
	ReportDate       string `queryParam:"style=form,explode=true,name=reportDate"`
	ShowMetricInputs *bool  `queryParam:"style=form,explode=true,name=showMetricInputs"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetricsRequest struct {
	PathParams  GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetricsPathParams
	QueryParams GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetricsQueryParams
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONErrorsTypeEnum string

const (
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONErrorsTypeEnumDataNotSynced      GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONErrorsTypeEnum = "DataNotSynced"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONErrorsTypeEnumDataNotSupported   GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONErrorsTypeEnum = "DataNotSupported"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONErrorsTypeEnumDataSyncFailed     GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONErrorsTypeEnum = "DataSyncFailed"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONErrorsTypeEnumDataTypeNotEnabled GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONErrorsTypeEnum = "DataTypeNotEnabled"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONErrors struct {
	Message *string                                                                                                 `json:"message,omitempty"`
	Type    *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONErrorsTypeEnum `json:"type,omitempty"`
}

// GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricErrorsAssessErrorDetails
// Dictionary list outlining the missing properties or allowed values.
type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricErrorsAssessErrorDetails struct {
	PropertyDetail1 []string `json:"propertyDetail1,omitempty"`
	PropertyDetail2 []string `json:"propertyDetail2,omitempty"`
	PropertyDetailN []string `json:"propertyDetailN,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricErrorsTypeEnum string

const (
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricErrorsTypeEnumUncategorizedAccounts GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricErrorsTypeEnum = "UncategorizedAccounts"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricErrorsTypeEnumMissingInputData      GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricErrorsTypeEnum = "MissingInputData"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricErrorsTypeEnumInputDataError        GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricErrorsTypeEnum = "InputDataError"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricErrors struct {
	Details *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricErrorsAssessErrorDetails `json:"details,omitempty"`
	Message *string                                                                                                                          `json:"message,omitempty"`
	Type    *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricErrorsTypeEnum           `json:"type,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum string

const (
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumUnknown                     GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "Unknown"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumEbitda                      GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "EBITDA"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumDebtServiceCoverageRatio    GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "DebtServiceCoverageRatio"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumCurrentRatioQuickRatio      GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "CurrentRatio QuickRatio"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumGrossProfitMargin           GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "GrossProfitMargin"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumFixedChargeCoverageRatio    GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "FixedChargeCoverageRatio"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumWorkingCapital              GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "WorkingCapital"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumFreeCashFlow                GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "FreeCashFlow"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumNetProfitMargin             GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "NetProfitMargin"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumReturnOnAssetsRatio         GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "ReturnOnAssetsRatio"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumReturnOnEquityRatio         GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "ReturnOnEquityRatio"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumOperatingProfitMargin       GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "OperatingProfitMargin"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumDeptToEquity                GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "DeptToEquity"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumDebtToAssets                GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "DebtToAssets"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumInterestCoverageRatio       GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "InterestCoverageRatio"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumCashRatio                   GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "CashRatio"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumInventoryTurnoverRatio      GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "InventoryTurnoverRatio"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumAssetTurnoverRatio          GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "AssetTurnoverRatio"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumWorkingCapitalTurnoverRatio GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "WorkingCapitalTurnoverRatio"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumDaysSalesOutstanding        GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "DaysSalesOutstanding"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumDaysPayablesOutstanding     GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "DaysPayablesOutstanding"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricMetricUnitEnum string

const (
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricMetricUnitEnumRatio GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricMetricUnitEnum = "Ratio"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricMetricUnitEnumMoney GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricMetricUnitEnum = "Money"
)

// GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsAssessErrorDetails
// Dictionary list outlining the missing properties or allowed values.
type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsAssessErrorDetails struct {
	PropertyDetail1 []string `json:"propertyDetail1,omitempty"`
	PropertyDetail2 []string `json:"propertyDetail2,omitempty"`
	PropertyDetailN []string `json:"propertyDetailN,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsTypeEnum string

const (
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsTypeEnumMissingAccountData GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsTypeEnum = "MissingAccountData"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsTypeEnumDatesOutOfRange    GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsTypeEnum = "DatesOutOfRange"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrors struct {
	Details *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsAssessErrorDetails `json:"details,omitempty"`
	Massage *string                                                                                                                                 `json:"massage,omitempty"`
	Type    *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsTypeEnum           `json:"type,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricPeriodsInputs struct {
	Name  *string  `json:"name,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricPeriods struct {
	Errors   []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrors `json:"errors,omitempty"`
	FromDate *types.Date                                                                                                            `json:"fromDate,omitempty"`
	Inputs   []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricPeriodsInputs `json:"inputs,omitempty"`
	ToDate   *types.Date                                                                                                            `json:"toDate,omitempty"`
	Value    *float64                                                                                                               `json:"value,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetric struct {
	Errors     []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricErrors        `json:"errors,omitempty"`
	Key        *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum        `json:"key,omitempty"`
	MetricUnit *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricMetricUnitEnum `json:"metricUnit,omitempty"`
	Name       *string                                                                                                                `json:"name,omitempty"`
	Periods    []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetricPeriods       `json:"periods,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONPeriodUnitEnum string

const (
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONPeriodUnitEnumMonth GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONPeriodUnitEnum = "Month"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONPeriodUnitEnumWeek  GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONPeriodUnitEnum = "Week"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONPeriodUnitEnumDay   GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONPeriodUnitEnum = "Day"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSON struct {
	Currency   *string                                                                                                   `json:"currency,omitempty"`
	Errors     []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONErrors          `json:"errors,omitempty"`
	Metrics    []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONFinancialMetric `json:"metrics,omitempty"`
	PeriodUnit *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONPeriodUnitEnum   `json:"periodUnit,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetricsResponse struct {
	ContentType                                                                                    string
	StatusCode                                                                                     int
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONObject *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSON
}
