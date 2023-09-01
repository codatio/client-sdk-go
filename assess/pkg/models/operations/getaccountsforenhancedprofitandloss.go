// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"net/http"
)

type GetAccountsForEnhancedProfitAndLossRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// The number of periods to return. If not provided, 12 periods will be used as the default value.
	NumberOfPeriods *int64 `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	// The date in which the report is created up to. Users must specify a specific date, however the response will be provided for the full month.
	ReportDate string `queryParam:"style=form,explode=true,name=reportDate"`
}

func (o *GetAccountsForEnhancedProfitAndLossRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountsForEnhancedProfitAndLossRequest) GetNumberOfPeriods() *int64 {
	if o == nil {
		return nil
	}
	return o.NumberOfPeriods
}

func (o *GetAccountsForEnhancedProfitAndLossRequest) GetReportDate() string {
	if o == nil {
		return ""
	}
	return o.ReportDate
}

type GetAccountsForEnhancedProfitAndLossResponse struct {
	ContentType string
	// OK
	EnhancedReport *shared.EnhancedReport
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAccountsForEnhancedProfitAndLossResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountsForEnhancedProfitAndLossResponse) GetEnhancedReport() *shared.EnhancedReport {
	if o == nil {
		return nil
	}
	return o.EnhancedReport
}

func (o *GetAccountsForEnhancedProfitAndLossResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAccountsForEnhancedProfitAndLossResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountsForEnhancedProfitAndLossResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
