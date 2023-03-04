package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/types"
	"net/http"
)

type GetEnhancedFinancialMetricsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetEnhancedFinancialMetricsQueryParams struct {
	NumberOfPeriods  int64  `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	PeriodLength     int64  `queryParam:"style=form,explode=true,name=periodLength"`
	ReportDate       string `queryParam:"style=form,explode=true,name=reportDate"`
	ShowMetricInputs *bool  `queryParam:"style=form,explode=true,name=showMetricInputs"`
}

type GetEnhancedFinancialMetricsRequest struct {
	PathParams  GetEnhancedFinancialMetricsPathParams
	QueryParams GetEnhancedFinancialMetricsQueryParams
}

type GetEnhancedFinancialMetrics200ApplicationJSONErrorsTypeEnum string

const (
	GetEnhancedFinancialMetrics200ApplicationJSONErrorsTypeEnumDataNotSynced      GetEnhancedFinancialMetrics200ApplicationJSONErrorsTypeEnum = "DataNotSynced"
	GetEnhancedFinancialMetrics200ApplicationJSONErrorsTypeEnumDataNotSupported   GetEnhancedFinancialMetrics200ApplicationJSONErrorsTypeEnum = "DataNotSupported"
	GetEnhancedFinancialMetrics200ApplicationJSONErrorsTypeEnumDataSyncFailed     GetEnhancedFinancialMetrics200ApplicationJSONErrorsTypeEnum = "DataSyncFailed"
	GetEnhancedFinancialMetrics200ApplicationJSONErrorsTypeEnumDataTypeNotEnabled GetEnhancedFinancialMetrics200ApplicationJSONErrorsTypeEnum = "DataTypeNotEnabled"
)

type GetEnhancedFinancialMetrics200ApplicationJSONErrors struct {
	Message *string                                                      `json:"message,omitempty"`
	Type    *GetEnhancedFinancialMetrics200ApplicationJSONErrorsTypeEnum `json:"type,omitempty"`
}

// GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricErrorsAssessErrorDetails
// Dictionary list outlining the missing properties or allowed values.
type GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricErrorsAssessErrorDetails struct {
	PropertyDetail1 []string `json:"propertyDetail1,omitempty"`
	PropertyDetail2 []string `json:"propertyDetail2,omitempty"`
	PropertyDetailN []string `json:"propertyDetailN,omitempty"`
}

type GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricErrorsTypeEnum string

const (
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricErrorsTypeEnumUncategorizedAccounts GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricErrorsTypeEnum = "UncategorizedAccounts"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricErrorsTypeEnumMissingInputData      GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricErrorsTypeEnum = "MissingInputData"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricErrorsTypeEnumInputDataError        GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricErrorsTypeEnum = "InputDataError"
)

type GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricErrors struct {
	Details *GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricErrorsAssessErrorDetails `json:"details,omitempty"`
	Message *string                                                                               `json:"message,omitempty"`
	Type    *GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricErrorsTypeEnum           `json:"type,omitempty"`
}

type GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum string

const (
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumUnknown                     GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "Unknown"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumEbitda                      GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "EBITDA"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumDebtServiceCoverageRatio    GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "DebtServiceCoverageRatio"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumCurrentRatioQuickRatio      GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "CurrentRatio QuickRatio"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumGrossProfitMargin           GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "GrossProfitMargin"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumFixedChargeCoverageRatio    GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "FixedChargeCoverageRatio"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumWorkingCapital              GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "WorkingCapital"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumFreeCashFlow                GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "FreeCashFlow"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumNetProfitMargin             GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "NetProfitMargin"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumReturnOnAssetsRatio         GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "ReturnOnAssetsRatio"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumReturnOnEquityRatio         GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "ReturnOnEquityRatio"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumOperatingProfitMargin       GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "OperatingProfitMargin"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumDeptToEquity                GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "DeptToEquity"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumDebtToAssets                GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "DebtToAssets"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumInterestCoverageRatio       GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "InterestCoverageRatio"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumCashRatio                   GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "CashRatio"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumInventoryTurnoverRatio      GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "InventoryTurnoverRatio"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumAssetTurnoverRatio          GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "AssetTurnoverRatio"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumWorkingCapitalTurnoverRatio GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "WorkingCapitalTurnoverRatio"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumDaysSalesOutstanding        GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "DaysSalesOutstanding"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnumDaysPayablesOutstanding     GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum = "DaysPayablesOutstanding"
)

type GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricMetricUnitEnum string

const (
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricMetricUnitEnumRatio GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricMetricUnitEnum = "Ratio"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricMetricUnitEnumMoney GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricMetricUnitEnum = "Money"
)

// GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsAssessErrorDetails
// Dictionary list outlining the missing properties or allowed values.
type GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsAssessErrorDetails struct {
	PropertyDetail1 []string `json:"propertyDetail1,omitempty"`
	PropertyDetail2 []string `json:"propertyDetail2,omitempty"`
	PropertyDetailN []string `json:"propertyDetailN,omitempty"`
}

type GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsTypeEnum string

const (
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsTypeEnumMissingAccountData GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsTypeEnum = "MissingAccountData"
	GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsTypeEnumDatesOutOfRange    GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsTypeEnum = "DatesOutOfRange"
)

type GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrors struct {
	Details *GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsAssessErrorDetails `json:"details,omitempty"`
	Massage *string                                                                                      `json:"massage,omitempty"`
	Type    *GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrorsTypeEnum           `json:"type,omitempty"`
}

type GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricPeriodsInputs struct {
	Name  *string  `json:"name,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

type GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricPeriods struct {
	Errors   []GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricPeriodsErrors `json:"errors,omitempty"`
	FromDate *types.Date                                                                 `json:"fromDate,omitempty"`
	Inputs   []GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricPeriodsInputs `json:"inputs,omitempty"`
	ToDate   *types.Date                                                                 `json:"toDate,omitempty"`
	Value    *float64                                                                    `json:"value,omitempty"`
}

type GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetric struct {
	Errors     []GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricErrors        `json:"errors,omitempty"`
	Key        *GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricKeyEnum        `json:"key,omitempty"`
	MetricUnit *GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricMetricUnitEnum `json:"metricUnit,omitempty"`
	Name       *string                                                                     `json:"name,omitempty"`
	Periods    []GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetricPeriods       `json:"periods,omitempty"`
}

type GetEnhancedFinancialMetrics200ApplicationJSONPeriodUnitEnum string

const (
	GetEnhancedFinancialMetrics200ApplicationJSONPeriodUnitEnumMonth GetEnhancedFinancialMetrics200ApplicationJSONPeriodUnitEnum = "Month"
	GetEnhancedFinancialMetrics200ApplicationJSONPeriodUnitEnumWeek  GetEnhancedFinancialMetrics200ApplicationJSONPeriodUnitEnum = "Week"
	GetEnhancedFinancialMetrics200ApplicationJSONPeriodUnitEnumDay   GetEnhancedFinancialMetrics200ApplicationJSONPeriodUnitEnum = "Day"
)

type GetEnhancedFinancialMetrics200ApplicationJSON struct {
	Currency   *string                                                        `json:"currency,omitempty"`
	Errors     []GetEnhancedFinancialMetrics200ApplicationJSONErrors          `json:"errors,omitempty"`
	Metrics    []GetEnhancedFinancialMetrics200ApplicationJSONFinancialMetric `json:"metrics,omitempty"`
	PeriodUnit *GetEnhancedFinancialMetrics200ApplicationJSONPeriodUnitEnum   `json:"periodUnit,omitempty"`
}

type GetEnhancedFinancialMetricsResponse struct {
	ContentType                                         string
	StatusCode                                          int
	RawResponse                                         *http.Response
	GetEnhancedFinancialMetrics200ApplicationJSONObject *GetEnhancedFinancialMetrics200ApplicationJSON
}
