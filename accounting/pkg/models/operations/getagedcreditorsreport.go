// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/types"
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

type GetAgedCreditorsReportResponse struct {
	// OK
	AgedCreditorReport *shared.AgedCreditorReport
	ContentType        string
	StatusCode         int
	RawResponse        *http.Response
	// Your API request was not properly authorized.
	Schema *shared.Schema
}
