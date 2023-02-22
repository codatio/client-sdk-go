package shared

import (
	"time"
)

type ProfitAndLossResponseReportBasisEnum string

const (
	ProfitAndLossResponseReportBasisEnumUnknown ProfitAndLossResponseReportBasisEnum = "Unknown"
	ProfitAndLossResponseReportBasisEnumAccrual ProfitAndLossResponseReportBasisEnum = "Accrual"
	ProfitAndLossResponseReportBasisEnumCash    ProfitAndLossResponseReportBasisEnum = "Cash"
)

type ProfitAndLossResponse struct {
	Currency                 string                               `json:"currency"`
	EarliestAvailableMonth   *time.Time                           `json:"earliestAvailableMonth,omitempty"`
	MostRecentAvailableMonth *time.Time                           `json:"mostRecentAvailableMonth,omitempty"`
	ReportBasis              ProfitAndLossResponseReportBasisEnum `json:"reportBasis"`
	Reports                  []ProfitAndLossReport                `json:"reports"`
}
