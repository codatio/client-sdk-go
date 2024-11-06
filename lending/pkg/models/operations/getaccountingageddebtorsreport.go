// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v6/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v6/pkg/types"
	"github.com/codatio/client-sdk-go/lending/v6/pkg/utils"
	"net/http"
)

type GetAccountingAgedDebtorsReportRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Number of periods to include in the report.
	NumberOfPeriods *int `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	// The length of period in days.
	PeriodLengthDays *int `queryParam:"style=form,explode=true,name=periodLengthDays"`
	// Date the report is generated up to.
	ReportDate *types.Date `queryParam:"style=form,explode=true,name=reportDate"`
}

func (g GetAccountingAgedDebtorsReportRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetAccountingAgedDebtorsReportRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetAccountingAgedDebtorsReportRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingAgedDebtorsReportRequest) GetNumberOfPeriods() *int {
	if o == nil {
		return nil
	}
	return o.NumberOfPeriods
}

func (o *GetAccountingAgedDebtorsReportRequest) GetPeriodLengthDays() *int {
	if o == nil {
		return nil
	}
	return o.PeriodLengthDays
}

func (o *GetAccountingAgedDebtorsReportRequest) GetReportDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.ReportDate
}

type GetAccountingAgedDebtorsReportResponse struct {
	// OK
	AccountingAgedDebtorReport *shared.AccountingAgedDebtorReport
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAccountingAgedDebtorsReportResponse) GetAccountingAgedDebtorReport() *shared.AccountingAgedDebtorReport {
	if o == nil {
		return nil
	}
	return o.AccountingAgedDebtorReport
}

func (o *GetAccountingAgedDebtorsReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingAgedDebtorsReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingAgedDebtorsReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
