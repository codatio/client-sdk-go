// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/types"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"net/http"
)

type GetAgedCreditorsReportRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Number of periods to include in the report.
	NumberOfPeriods *int `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	// The length of period in days.
	PeriodLengthDays *int `queryParam:"style=form,explode=true,name=periodLengthDays"`
	// Date the report is generated up to.
	ReportDate *types.Date `queryParam:"style=form,explode=true,name=reportDate"`
}

func (g GetAgedCreditorsReportRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetAgedCreditorsReportRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetAgedCreditorsReportRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAgedCreditorsReportRequest) GetNumberOfPeriods() *int {
	if o == nil {
		return nil
	}
	return o.NumberOfPeriods
}

func (o *GetAgedCreditorsReportRequest) GetPeriodLengthDays() *int {
	if o == nil {
		return nil
	}
	return o.PeriodLengthDays
}

func (o *GetAgedCreditorsReportRequest) GetReportDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.ReportDate
}

type GetAgedCreditorsReportResponse struct {
	// OK
	AgedCreditorReport *shared.AgedCreditorReport
	ContentType        string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAgedCreditorsReportResponse) GetAgedCreditorReport() *shared.AgedCreditorReport {
	if o == nil {
		return nil
	}
	return o.AgedCreditorReport
}

func (o *GetAgedCreditorsReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAgedCreditorsReportResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAgedCreditorsReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAgedCreditorsReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
