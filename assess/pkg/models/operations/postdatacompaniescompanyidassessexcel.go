package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type PostDataCompaniesCompanyIDAssessExcelPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type PostDataCompaniesCompanyIDAssessExcelQueryParams struct {
	ReportType shared.ExcelReportTypeEnum `queryParam:"style=form,explode=true,name=reportType"`
}

type PostDataCompaniesCompanyIDAssessExcelRequest struct {
	PathParams  PostDataCompaniesCompanyIDAssessExcelPathParams
	QueryParams PostDataCompaniesCompanyIDAssessExcelQueryParams
}

type PostDataCompaniesCompanyIDAssessExcelResponse struct {
	ContentType string
	ExcelStatus *shared.ExcelStatus
	StatusCode  int
}
