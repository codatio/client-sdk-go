// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"net/http"
)

type GetAccountingMarketingMetricsRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Shows the dimensionDisplayName and itemDisplayName in measures to make the report data human-readable.
	IncludeDisplayNames *bool `queryParam:"style=form,explode=true,name=includeDisplayNames"`
	// The number of periods to return. There will be no pagination as a query parameter.
	NumberOfPeriods int64 `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	// The number of months per period. E.g. 2 = 2 months per period.
	PeriodLength int64 `queryParam:"style=form,explode=true,name=periodLength"`
	// The period unit of time returned.
	PeriodUnit shared.PeriodUnit `queryParam:"style=form,explode=true,name=periodUnit"`
	// The date in which the report is created up to. Users must specify a specific date, however the response will be provided for the full month.
	ReportDate string `queryParam:"style=form,explode=true,name=reportDate"`
	// If set to true, then the system includes the input values within the response.
	ShowInputValues *bool `queryParam:"style=form,explode=true,name=showInputValues"`
}

func (o *GetAccountingMarketingMetricsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingMarketingMetricsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetAccountingMarketingMetricsRequest) GetIncludeDisplayNames() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeDisplayNames
}

func (o *GetAccountingMarketingMetricsRequest) GetNumberOfPeriods() int64 {
	if o == nil {
		return 0
	}
	return o.NumberOfPeriods
}

func (o *GetAccountingMarketingMetricsRequest) GetPeriodLength() int64 {
	if o == nil {
		return 0
	}
	return o.PeriodLength
}

func (o *GetAccountingMarketingMetricsRequest) GetPeriodUnit() shared.PeriodUnit {
	if o == nil {
		return shared.PeriodUnit("")
	}
	return o.PeriodUnit
}

func (o *GetAccountingMarketingMetricsRequest) GetReportDate() string {
	if o == nil {
		return ""
	}
	return o.ReportDate
}

func (o *GetAccountingMarketingMetricsRequest) GetShowInputValues() *bool {
	if o == nil {
		return nil
	}
	return o.ShowInputValues
}

type GetAccountingMarketingMetricsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	Report *shared.Report
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAccountingMarketingMetricsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingMarketingMetricsResponse) GetReport() *shared.Report {
	if o == nil {
		return nil
	}
	return o.Report
}

func (o *GetAccountingMarketingMetricsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingMarketingMetricsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
