package operations

import (
	"net/http"
	"time"
)

type RequestExcelReportForDownloadReportTypeEnum string

const (
	RequestExcelReportForDownloadReportTypeEnumAudit RequestExcelReportForDownloadReportTypeEnum = "audit"
)

type RequestExcelReportForDownloadRequest struct {
	CompanyID  string                                      `pathParam:"style=simple,explode=false,name=companyId"`
	ReportType RequestExcelReportForDownloadReportTypeEnum `queryParam:"style=form,explode=true,name=reportType"`
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
	RawResponse                                           *http.Response
	RequestExcelReportForDownload200ApplicationJSONObject *RequestExcelReportForDownload200ApplicationJSON
}
