package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
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

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONPeriodUnitEnum string

const (
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONPeriodUnitEnumMonth GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONPeriodUnitEnum = "Month"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONPeriodUnitEnumWeek  GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONPeriodUnitEnum = "Week"
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONPeriodUnitEnumDay   GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONPeriodUnitEnum = "Day"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSON struct {
	Currency   *string                                                                                                 `json:"currency,omitempty"`
	Errors     []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONErrors        `json:"errors,omitempty"`
	Metrics    []shared.FinancialMetric                                                                                `json:"metrics,omitempty"`
	PeriodUnit *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONPeriodUnitEnum `json:"periodUnit,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetricsResponse struct {
	ContentType                                                                                    string
	StatusCode                                                                                     int
	GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSONObject *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessFinancialMetrics200ApplicationJSON
}
