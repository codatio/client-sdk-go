package operations

import (
	"net/http"
)

type GetExcelReportPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetExcelReportReportTypeEnum string

const (
	GetExcelReportReportTypeEnumAudit GetExcelReportReportTypeEnum = "audit"
)

type GetExcelReportQueryParams struct {
	ReportType GetExcelReportReportTypeEnum `queryParam:"style=form,explode=true,name=reportType"`
}

type GetExcelReportRequest struct {
	PathParams  GetExcelReportPathParams
	QueryParams GetExcelReportQueryParams
}

type GetExcelReportResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
