package operations

import (
	"time"
)

type RequestExcelReportForDownloadPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type RequestExcelReportForDownloadReportTypeEnum string

const (
	RequestExcelReportForDownloadReportTypeEnumAudit RequestExcelReportForDownloadReportTypeEnum = "audit"
)

type RequestExcelReportForDownloadQueryParams struct {
	ReportType RequestExcelReportForDownloadReportTypeEnum `queryParam:"style=form,explode=true,name=reportType"`
}

type RequestExcelReportForDownloadRequest struct {
	PathParams  RequestExcelReportForDownloadPathParams
	QueryParams RequestExcelReportForDownloadQueryParams
}

type RequestExcelReportForDownload200ApplicationJSON struct {
	ErrorMessage     *string    `json:"errorMessage,omitempty"`
	FileSize         *int64     `json:"fileSize,omitempty"`
	InProgress       *bool      `json:"inProgress,omitempty"`
	LastGenerated    *time.Time `json:"lastGenerated,omitempty"`
	LastInvocationID *string    `json:"lastInvocationId,omitempty"`
	Queued           *string    `json:"queued,omitempty"`
	ReportType       *string    `json:"reportType,omitempty"`
	Success          *bool      `json:"success,omitempty"`
}

type RequestExcelReportForDownloadResponse struct {
	ContentType                                           string
	StatusCode                                            int
	RequestExcelReportForDownload200ApplicationJSONObject *RequestExcelReportForDownload200ApplicationJSON
}
