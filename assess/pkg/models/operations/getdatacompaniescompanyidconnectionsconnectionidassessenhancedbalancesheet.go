package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedBalanceSheetPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedBalanceSheetQueryParams struct {
	IncludeDisplayNames *bool  `queryParam:"style=form,explode=true,name=includeDisplayNames"`
	NumberOfPeriods     int64  `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	PeriodLength        int64  `queryParam:"style=form,explode=true,name=periodLength"`
	ReportDate          string `queryParam:"style=form,explode=true,name=reportDate"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedBalanceSheetRequest struct {
	PathParams  GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedBalanceSheetPathParams
	QueryParams GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedBalanceSheetQueryParams
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessEnhancedBalanceSheetResponse struct {
	ContentType string
	Report      *shared.Report
	StatusCode  int
}
