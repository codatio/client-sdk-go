package operations

import (
	"time"
)

type PostDataCompaniesCompanyIDAssessExcelPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type PostDataCompaniesCompanyIDAssessExcelReportTypeEnum string

const (
	PostDataCompaniesCompanyIDAssessExcelReportTypeEnumAudit PostDataCompaniesCompanyIDAssessExcelReportTypeEnum = "audit"
)

type PostDataCompaniesCompanyIDAssessExcelQueryParams struct {
	ReportType PostDataCompaniesCompanyIDAssessExcelReportTypeEnum `queryParam:"style=form,explode=true,name=reportType"`
}

type PostDataCompaniesCompanyIDAssessExcelRequest struct {
	PathParams  PostDataCompaniesCompanyIDAssessExcelPathParams
	QueryParams PostDataCompaniesCompanyIDAssessExcelQueryParams
}

type PostDataCompaniesCompanyIDAssessExcel200ApplicationJSON struct {
	ErrorMessage     *string    `json:"errorMessage,omitempty"`
	FileSize         *int64     `json:"fileSize,omitempty"`
	InProgress       *bool      `json:"inProgress,omitempty"`
	LastGenerated    *time.Time `json:"lastGenerated,omitempty"`
	LastInvocationID *string    `json:"lastInvocationId,omitempty"`
	Queued           *string    `json:"queued,omitempty"`
	ReportType       *string    `json:"reportType,omitempty"`
	Success          *bool      `json:"success,omitempty"`
}

type PostDataCompaniesCompanyIDAssessExcelResponse struct {
	ContentType                                                   string
	StatusCode                                                    int
	PostDataCompaniesCompanyIDAssessExcel200ApplicationJSONObject *PostDataCompaniesCompanyIDAssessExcel200ApplicationJSON
}
