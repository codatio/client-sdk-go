package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsQueryParams struct {
	NumberOfPeriods int64  `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	ReportDate      string `queryParam:"style=form,explode=true,name=reportDate"`
}

type GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsRequest struct {
	PathParams  GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsPathParams
	QueryParams GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsQueryParams
}

type GetCompaniesCompanyIDReportsEnhancedProfitAndLossAccountsResponse struct {
	ContentType    string
	EnhancedReport *shared.EnhancedReport
	StatusCode     int
}
