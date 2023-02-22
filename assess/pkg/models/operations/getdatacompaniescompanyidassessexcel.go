package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type GetDataCompaniesCompanyIDAssessExcelPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetDataCompaniesCompanyIDAssessExcelQueryParams struct {
	ReportType shared.ExcelReportTypeEnum `queryParam:"style=form,explode=true,name=reportType"`
}

type GetDataCompaniesCompanyIDAssessExcelRequest struct {
	PathParams  GetDataCompaniesCompanyIDAssessExcelPathParams
	QueryParams GetDataCompaniesCompanyIDAssessExcelQueryParams
}

type GetDataCompaniesCompanyIDAssessExcelResponse struct {
	ContentType string
	ExcelStatus *shared.ExcelStatus
	StatusCode  int
}
