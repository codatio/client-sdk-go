package operations

import (
	"time"
)

type GetDataCompaniesCompanyIDAssessExcelPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetDataCompaniesCompanyIDAssessExcelReportTypeEnum string

const (
	GetDataCompaniesCompanyIDAssessExcelReportTypeEnumAudit GetDataCompaniesCompanyIDAssessExcelReportTypeEnum = "audit"
)

type GetDataCompaniesCompanyIDAssessExcelQueryParams struct {
	ReportType GetDataCompaniesCompanyIDAssessExcelReportTypeEnum `queryParam:"style=form,explode=true,name=reportType"`
}

type GetDataCompaniesCompanyIDAssessExcelRequest struct {
	PathParams  GetDataCompaniesCompanyIDAssessExcelPathParams
	QueryParams GetDataCompaniesCompanyIDAssessExcelQueryParams
}

type GetDataCompaniesCompanyIDAssessExcel200ApplicationJSON struct {
	ErrorMessage     *string    `json:"errorMessage,omitempty"`
	FileSize         *int64     `json:"fileSize,omitempty"`
	InProgress       *bool      `json:"inProgress,omitempty"`
	LastGenerated    *time.Time `json:"lastGenerated,omitempty"`
	LastInvocationID *string    `json:"lastInvocationId,omitempty"`
	Queued           *string    `json:"queued,omitempty"`
	ReportType       *string    `json:"reportType,omitempty"`
	Success          *bool      `json:"success,omitempty"`
}

type GetDataCompaniesCompanyIDAssessExcelResponse struct {
	ContentType                                                  string
	StatusCode                                                   int
	GetDataCompaniesCompanyIDAssessExcel200ApplicationJSONObject *GetDataCompaniesCompanyIDAssessExcel200ApplicationJSON
}
