// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v2/pkg/models/shared"
	"net/http"
)

type DownloadExcelReportRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// The type of report you want to generate and download.
	ReportType shared.ExcelReportTypes `queryParam:"style=form,explode=true,name=reportType"`
}

func (o *DownloadExcelReportRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *DownloadExcelReportRequest) GetReportType() shared.ExcelReportTypes {
	if o == nil {
		return shared.ExcelReportTypes("")
	}
	return o.ReportType
}

type DownloadExcelReportResponse struct {
	Body        []byte
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *DownloadExcelReportResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *DownloadExcelReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DownloadExcelReportResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *DownloadExcelReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DownloadExcelReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
