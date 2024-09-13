// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"net/http"
)

type GetExcelReportGenerationStatusRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// The type of report you want to generate and download.
	ReportType shared.ExcelReportType `queryParam:"style=form,explode=true,name=reportType"`
}

func (o *GetExcelReportGenerationStatusRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetExcelReportGenerationStatusRequest) GetReportType() shared.ExcelReportType {
	if o == nil {
		return shared.ExcelReportType("")
	}
	return o.ReportType
}

type GetExcelReportGenerationStatusResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	ExcelStatus *shared.ExcelStatus
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetExcelReportGenerationStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetExcelReportGenerationStatusResponse) GetExcelStatus() *shared.ExcelStatus {
	if o == nil {
		return nil
	}
	return o.ExcelStatus
}

func (o *GetExcelReportGenerationStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetExcelReportGenerationStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
