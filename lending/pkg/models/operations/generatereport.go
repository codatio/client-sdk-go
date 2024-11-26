// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"net/http"
)

type GenerateReportRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// The type of the report
	ReportType shared.ReportType `pathParam:"style=simple,explode=false,name=reportType"`
}

func (o *GenerateReportRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GenerateReportRequest) GetReportType() shared.ReportType {
	if o == nil {
		return shared.ReportType("")
	}
	return o.ReportType
}

type GenerateReportResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	ReportOperation *shared.ReportOperation
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GenerateReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GenerateReportResponse) GetReportOperation() *shared.ReportOperation {
	if o == nil {
		return nil
	}
	return o.ReportOperation
}

func (o *GenerateReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GenerateReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
