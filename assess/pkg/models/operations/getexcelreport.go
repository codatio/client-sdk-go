// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"net/http"
)

type GetExcelReportRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// The type of report you want to generate and download.
	ReportType shared.ExcelReportType `queryParam:"style=form,explode=true,name=reportType"`
}

func (o *GetExcelReportRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetExcelReportRequest) GetReportType() shared.ExcelReportType {
	if o == nil {
		return shared.ExcelReportType("")
	}
	return o.ReportType
}

// GetExcelReport200ApplicationOctetStream - OK
type GetExcelReport200ApplicationOctetStream struct {
}

type GetExcelReportResponse struct {
	Body        []byte
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetExcelReportResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *GetExcelReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetExcelReportResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetExcelReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetExcelReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
