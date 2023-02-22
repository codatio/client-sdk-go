package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsQueryParams struct {
	NumberOfPeriods int64  `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	ReportDate      string `queryParam:"style=form,explode=true,name=reportDate"`
}

type GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsRequest struct {
	PathParams  GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsPathParams
	QueryParams GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsQueryParams
}

type GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsResponse struct {
	ContentType    string
	EnhancedReport *shared.EnhancedReport
	StatusCode     int
}
