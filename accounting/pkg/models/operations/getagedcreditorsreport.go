package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/types"
	"net/http"
	"time"
)

type GetAgedCreditorsReportPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetAgedCreditorsReportQueryParams struct {
	NumberOfPeriods  *int        `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	PeriodLengthDays *int        `queryParam:"style=form,explode=true,name=periodLengthDays"`
	ReportDate       *types.Date `queryParam:"style=form,explode=true,name=reportDate"`
}

type GetAgedCreditorsReportRequest struct {
	PathParams  GetAgedCreditorsReportPathParams
	QueryParams GetAgedCreditorsReportQueryParams
}

type GetAgedCreditorsReportAgedCreditorsReportAgedCreditorAgedCurrencyOutstandingAgedOutstandingAmountAmountsOutstandingByDataType struct {
	Amount *float64 `json:"amount,omitempty"`
	Name   *string  `json:"name,omitempty"`
}

type GetAgedCreditorsReportAgedCreditorsReportAgedCreditorAgedCurrencyOutstandingAgedOutstandingAmount struct {
	Amount   *float64                                                                                                                        `json:"amount,omitempty"`
	Details  []GetAgedCreditorsReportAgedCreditorsReportAgedCreditorAgedCurrencyOutstandingAgedOutstandingAmountAmountsOutstandingByDataType `json:"details,omitempty"`
	FromDate *time.Time                                                                                                                      `json:"fromDate,omitempty"`
	ToDate   *time.Time                                                                                                                      `json:"toDate,omitempty"`
}

type GetAgedCreditorsReportAgedCreditorsReportAgedCreditorAgedCurrencyOutstanding struct {
	AgedOutstandingAmounts []GetAgedCreditorsReportAgedCreditorsReportAgedCreditorAgedCurrencyOutstandingAgedOutstandingAmount `json:"agedOutstandingAmounts,omitempty"`
	Currency               *string                                                                                             `json:"currency,omitempty"`
}

type GetAgedCreditorsReportAgedCreditorsReportAgedCreditor struct {
	AgedCurrencyOutstanding []GetAgedCreditorsReportAgedCreditorsReportAgedCreditorAgedCurrencyOutstanding `json:"agedCurrencyOutstanding,omitempty"`
	SupplierID              *string                                                                        `json:"supplierId,omitempty"`
	SupplierName            *string                                                                        `json:"supplierName,omitempty"`
}

type GetAgedCreditorsReportAgedCreditorsReport struct {
	Data       []GetAgedCreditorsReportAgedCreditorsReportAgedCreditor `json:"data,omitempty"`
	Generated  *time.Time                                              `json:"generated,omitempty"`
	ReportDate *time.Time                                              `json:"reportDate,omitempty"`
}

type GetAgedCreditorsReportResponse struct {
	AgedCreditorsReport *GetAgedCreditorsReportAgedCreditorsReport
	ContentType         string
	StatusCode          int
	RawResponse         *http.Response
}
