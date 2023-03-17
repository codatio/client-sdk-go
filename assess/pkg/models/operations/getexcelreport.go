package operations

import (
	"net/http"
)

type GetExcelReportReportTypeEnum string

const (
	GetExcelReportReportTypeEnumAudit GetExcelReportReportTypeEnum = "audit"
)

type GetExcelReportRequest struct {
	CompanyID  string                       `pathParam:"style=simple,explode=false,name=companyId"`
	ReportType GetExcelReportReportTypeEnum `queryParam:"style=form,explode=true,name=reportType"`
}

type GetExcelReportResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
