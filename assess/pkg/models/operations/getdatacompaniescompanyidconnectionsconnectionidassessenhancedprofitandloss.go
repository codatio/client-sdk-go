package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedProfitAndLossPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedProfitAndLossQueryParams struct {
	IncludeDisplayNames *bool  `queryParam:"style=form,explode=true,name=includeDisplayNames"`
	NumberOfPeriods     int64  `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	PeriodLength        int64  `queryParam:"style=form,explode=true,name=periodLength"`
	ReportDate          string `queryParam:"style=form,explode=true,name=reportDate"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedProfitAndLossRequest struct {
	PathParams  GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedProfitAndLossPathParams
	QueryParams GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedProfitAndLossQueryParams
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedProfitAndLossResponse struct {
	ContentType string
	Report      *shared.Report
	StatusCode  int
}
