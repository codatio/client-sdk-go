// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"net/http"
)

type GetExcelReportGenerationStatusRequest struct {
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
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// OK
	ExcelStatus *shared.ExcelStatus
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetExcelReportGenerationStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetExcelReportGenerationStatusResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
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