package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type PostDataCompaniesCompanyIDAssessExcelDownloadPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type PostDataCompaniesCompanyIDAssessExcelDownloadQueryParams struct {
	ReportType shared.ExcelReportTypeEnum `queryParam:"style=form,explode=true,name=reportType"`
}

type PostDataCompaniesCompanyIDAssessExcelDownloadRequest struct {
	PathParams  PostDataCompaniesCompanyIDAssessExcelDownloadPathParams
	QueryParams PostDataCompaniesCompanyIDAssessExcelDownloadQueryParams
}

type PostDataCompaniesCompanyIDAssessExcelDownloadResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
}
