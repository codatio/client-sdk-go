package operations

import (
	"time"
)

type MakeRequestToDownloadExcelReportPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type MakeRequestToDownloadExcelReportReportTypeEnum string

const (
	MakeRequestToDownloadExcelReportReportTypeEnumAudit MakeRequestToDownloadExcelReportReportTypeEnum = "audit"
)

type MakeRequestToDownloadExcelReportQueryParams struct {
	ReportType MakeRequestToDownloadExcelReportReportTypeEnum `queryParam:"style=form,explode=true,name=reportType"`
}

type MakeRequestToDownloadExcelReportRequest struct {
	PathParams  MakeRequestToDownloadExcelReportPathParams
	QueryParams MakeRequestToDownloadExcelReportQueryParams
}

type MakeRequestToDownloadExcelReport200ApplicationJSON struct {
	ErrorMessage     *string    `json:"errorMessage,omitempty"`
	FileSize         *int64     `json:"fileSize,omitempty"`
	InProgress       *bool      `json:"inProgress,omitempty"`
	LastGenerated    *time.Time `json:"lastGenerated,omitempty"`
	LastInvocationID *string    `json:"lastInvocationId,omitempty"`
	Queued           *string    `json:"queued,omitempty"`
	ReportType       *string    `json:"reportType,omitempty"`
	Success          *bool      `json:"success,omitempty"`
}

type MakeRequestToDownloadExcelReportResponse struct {
	ContentType                                              string
	StatusCode                                               int
	MakeRequestToDownloadExcelReport200ApplicationJSONObject *MakeRequestToDownloadExcelReport200ApplicationJSON
}
