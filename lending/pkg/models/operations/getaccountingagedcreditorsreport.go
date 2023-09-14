// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v3/pkg/types"
	"net/http"
)

type GetAccountingAgedCreditorsReportRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Number of periods to include in the report.
	NumberOfPeriods *int `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	// The length of period in days.
	PeriodLengthDays *int `queryParam:"style=form,explode=true,name=periodLengthDays"`
	// Date the report is generated up to.
	ReportDate *types.Date `queryParam:"style=form,explode=true,name=reportDate"`
}

func (o *GetAccountingAgedCreditorsReportRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingAgedCreditorsReportRequest) GetNumberOfPeriods() *int {
	if o == nil {
		return nil
	}
	return o.NumberOfPeriods
}

func (o *GetAccountingAgedCreditorsReportRequest) GetPeriodLengthDays() *int {
	if o == nil {
		return nil
	}
	return o.PeriodLengthDays
}

func (o *GetAccountingAgedCreditorsReportRequest) GetReportDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.ReportDate
}

type GetAccountingAgedCreditorsReportResponse struct {
	// OK
	AccountingAgedCreditorReport *shared.AccountingAgedCreditorReport
	ContentType                  string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAccountingAgedCreditorsReportResponse) GetAccountingAgedCreditorReport() *shared.AccountingAgedCreditorReport {
	if o == nil {
		return nil
	}
	return o.AccountingAgedCreditorReport
}

func (o *GetAccountingAgedCreditorsReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingAgedCreditorsReportResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAccountingAgedCreditorsReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingAgedCreditorsReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
