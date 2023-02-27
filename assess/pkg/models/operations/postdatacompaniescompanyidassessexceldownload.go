package operations

type PostDataCompaniesCompanyIDAssessExcelDownloadPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type PostDataCompaniesCompanyIDAssessExcelDownloadReportTypeEnum string

const (
	PostDataCompaniesCompanyIDAssessExcelDownloadReportTypeEnumAudit PostDataCompaniesCompanyIDAssessExcelDownloadReportTypeEnum = "audit"
)

type PostDataCompaniesCompanyIDAssessExcelDownloadQueryParams struct {
	ReportType PostDataCompaniesCompanyIDAssessExcelDownloadReportTypeEnum `queryParam:"style=form,explode=true,name=reportType"`
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
