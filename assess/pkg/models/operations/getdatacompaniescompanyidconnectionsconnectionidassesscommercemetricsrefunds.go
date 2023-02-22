package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRefundsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRefundsQueryParams struct {
	IncludeDisplayNames *bool                 `queryParam:"style=form,explode=true,name=includeDisplayNames"`
	NumberOfPeriods     int64                 `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	PeriodLength        int64                 `queryParam:"style=form,explode=true,name=periodLength"`
	PeriodUnit          shared.PeriodUnitEnum `queryParam:"style=form,explode=true,name=periodUnit"`
	ReportDate          string                `queryParam:"style=form,explode=true,name=reportDate"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRefundsRequest struct {
	PathParams  GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRefundsPathParams
	QueryParams GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRefundsQueryParams
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessCommerceMetricsRefundsResponse struct {
	ContentType string
	Report      *shared.Report
	StatusCode  int
}
