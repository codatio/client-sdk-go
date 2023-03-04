package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/types"
	"net/http"
	"time"
)

type GetAgedDebtorsReportPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetAgedDebtorsReportQueryParams struct {
	NumberOfPeriods  *int        `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	PeriodLengthDays *int        `queryParam:"style=form,explode=true,name=periodLengthDays"`
	ReportDate       *types.Date `queryParam:"style=form,explode=true,name=reportDate"`
}

type GetAgedDebtorsReportRequest struct {
	PathParams  GetAgedDebtorsReportPathParams
	QueryParams GetAgedDebtorsReportQueryParams
}

type GetAgedDebtorsReportAgedDebtorsReportAgedDebtorAgedCurrencyOutstandingAgedOutstandingAmountAmountsOutstandingByDataType struct {
	Amount *float64 `json:"amount,omitempty"`
	Name   *string  `json:"name,omitempty"`
}

type GetAgedDebtorsReportAgedDebtorsReportAgedDebtorAgedCurrencyOutstandingAgedOutstandingAmount struct {
	Amount   *float64                                                                                                                  `json:"amount,omitempty"`
	Details  []GetAgedDebtorsReportAgedDebtorsReportAgedDebtorAgedCurrencyOutstandingAgedOutstandingAmountAmountsOutstandingByDataType `json:"details,omitempty"`
	FromDate *time.Time                                                                                                                `json:"fromDate,omitempty"`
	ToDate   *time.Time                                                                                                                `json:"toDate,omitempty"`
}

type GetAgedDebtorsReportAgedDebtorsReportAgedDebtorAgedCurrencyOutstanding struct {
	AgedOutstandingAmounts []GetAgedDebtorsReportAgedDebtorsReportAgedDebtorAgedCurrencyOutstandingAgedOutstandingAmount `json:"agedOutstandingAmounts,omitempty"`
	Currency               *string                                                                                       `json:"currency,omitempty"`
}

type GetAgedDebtorsReportAgedDebtorsReportAgedDebtor struct {
	AgedCurrencyOutstanding []GetAgedDebtorsReportAgedDebtorsReportAgedDebtorAgedCurrencyOutstanding `json:"agedCurrencyOutstanding,omitempty"`
	CustomerID              *string                                                                  `json:"customerId,omitempty"`
	CustomerName            *string                                                                  `json:"customerName,omitempty"`
}

type GetAgedDebtorsReportAgedDebtorsReport struct {
	Data       []GetAgedDebtorsReportAgedDebtorsReportAgedDebtor `json:"data,omitempty"`
	Generated  *time.Time                                        `json:"generated,omitempty"`
	ReportDate *time.Time                                        `json:"reportDate,omitempty"`
}

type GetAgedDebtorsReportResponse struct {
	AgedDebtorsReport *GetAgedDebtorsReportAgedDebtorsReport
	ContentType       string
	StatusCode        int
	RawResponse       *http.Response
}
