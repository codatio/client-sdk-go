package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountingMetricsMarketingPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountingMetricsMarketingQueryParams struct {
	IncludeDisplayNames *bool                 `queryParam:"style=form,explode=true,name=includeDisplayNames"`
	NumberOfPeriods     int64                 `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	PeriodLength        int64                 `queryParam:"style=form,explode=true,name=periodLength"`
	PeriodUnit          shared.PeriodUnitEnum `queryParam:"style=form,explode=true,name=periodUnit"`
	ReportDate          string                `queryParam:"style=form,explode=true,name=reportDate"`
	ShowInputValues     *bool                 `queryParam:"style=form,explode=true,name=showInputValues"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountingMetricsMarketingRequest struct {
	PathParams  GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountingMetricsMarketingPathParams
	QueryParams GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountingMetricsMarketingQueryParams
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountingMetricsMarketingResponse struct {
	ContentType string
	Report      *shared.Report
	StatusCode  int
}
