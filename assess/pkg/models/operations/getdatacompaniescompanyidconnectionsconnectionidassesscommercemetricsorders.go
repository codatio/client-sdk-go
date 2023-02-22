package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsOrdersPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsOrdersQueryParams struct {
	IncludeDisplayNames *bool                 `queryParam:"style=form,explode=true,name=includeDisplayNames"`
	NumberOfPeriods     int64                 `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	PeriodLength        int64                 `queryParam:"style=form,explode=true,name=periodLength"`
	PeriodUnit          shared.PeriodUnitEnum `queryParam:"style=form,explode=true,name=periodUnit"`
	ReportDate          string                `queryParam:"style=form,explode=true,name=reportDate"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsOrdersRequest struct {
	PathParams  GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsOrdersPathParams
	QueryParams GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsOrdersQueryParams
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsOrdersResponse struct {
	ContentType string
	Report      *shared.Report
	StatusCode  int
}
