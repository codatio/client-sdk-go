// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"net/http"
)

type GetCategorizedProfitAndLossStatementRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// The number of periods to return. If not provided, 12 periods will be used as the default value.
	NumberOfPeriods *int64 `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	// The date in which the report is created up to. Users must specify a specific date, however the response will be provided for the full month.
	ReportDate *string `queryParam:"style=form,explode=true,name=reportDate"`
}

func (o *GetCategorizedProfitAndLossStatementRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetCategorizedProfitAndLossStatementRequest) GetNumberOfPeriods() *int64 {
	if o == nil {
		return nil
	}
	return o.NumberOfPeriods
}

func (o *GetCategorizedProfitAndLossStatementRequest) GetReportDate() *string {
	if o == nil {
		return nil
	}
	return o.ReportDate
}

type GetCategorizedProfitAndLossStatementResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	EnhancedFinancialReport *shared.EnhancedFinancialReport
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetCategorizedProfitAndLossStatementResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCategorizedProfitAndLossStatementResponse) GetEnhancedFinancialReport() *shared.EnhancedFinancialReport {
	if o == nil {
		return nil
	}
	return o.EnhancedFinancialReport
}

func (o *GetCategorizedProfitAndLossStatementResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCategorizedProfitAndLossStatementResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
