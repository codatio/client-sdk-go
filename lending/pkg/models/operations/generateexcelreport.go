// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v7/pkg/models/shared"
	"net/http"
)

type GenerateExcelReportRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// The type of report you want to generate and download.
	ReportType shared.ExcelReportTypes `queryParam:"style=form,explode=true,name=reportType"`
}

func (o *GenerateExcelReportRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GenerateExcelReportRequest) GetReportType() shared.ExcelReportTypes {
	if o == nil {
		return shared.ExcelReportTypes("")
	}
	return o.ReportType
}

type GenerateExcelReportResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	ExcelStatus *shared.ExcelStatus
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GenerateExcelReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GenerateExcelReportResponse) GetExcelStatus() *shared.ExcelStatus {
	if o == nil {
		return nil
	}
	return o.ExcelStatus
}

func (o *GenerateExcelReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GenerateExcelReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
