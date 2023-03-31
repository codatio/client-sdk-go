// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"net/http"
)

type GetEnhancedFinancialMetricsRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// The number of periods to return.  There will be no pagination as a query parameter, however Codat will limit the number of periods to request to 12 periods.
	NumberOfPeriods int64 `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	// The number of months per period. E.g. 2 = 2 months per period.
	PeriodLength int64 `queryParam:"style=form,explode=true,name=periodLength"`
	// The date in which the report is created up to. Users must specify a specific date, however the response will be provided for the full month.
	ReportDate string `queryParam:"style=form,explode=true,name=reportDate"`
	// If set to true, then the system includes the input values within the response.
	ShowMetricInputs *bool `queryParam:"style=form,explode=true,name=showMetricInputs"`
}

type GetEnhancedFinancialMetricsResponse struct {
	ContentType string
	// OK
	FinancialMetrics *shared.FinancialMetrics
	StatusCode       int
	RawResponse      *http.Response
}
