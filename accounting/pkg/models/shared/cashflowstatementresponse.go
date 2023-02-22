package shared

import (
	"time"
)

type CashFlowStatementResponseReportBasisEnum string

const (
	CashFlowStatementResponseReportBasisEnumUnknown CashFlowStatementResponseReportBasisEnum = "Unknown"
	CashFlowStatementResponseReportBasisEnumAccrual CashFlowStatementResponseReportBasisEnum = "Accrual"
	CashFlowStatementResponseReportBasisEnumCash    CashFlowStatementResponseReportBasisEnum = "Cash"
)

type CashFlowStatementResponseReportInputEnum string

const (
	CashFlowStatementResponseReportInputEnumUnknown  CashFlowStatementResponseReportInputEnum = "Unknown"
	CashFlowStatementResponseReportInputEnumIndirect CashFlowStatementResponseReportInputEnum = "Indirect"
	CashFlowStatementResponseReportInputEnumDirect   CashFlowStatementResponseReportInputEnum = "Direct"
)

type CashFlowStatementResponse struct {
	Currency                 string                                   `json:"currency"`
	EarliestAvailableMonth   *time.Time                               `json:"earliestAvailableMonth,omitempty"`
	MostRecentAvailableMonth *time.Time                               `json:"mostRecentAvailableMonth,omitempty"`
	ReportBasis              CashFlowStatementResponseReportBasisEnum `json:"reportBasis"`
	ReportInput              CashFlowStatementResponseReportInputEnum `json:"reportInput"`
	Reports                  []CashFlowStatement                      `json:"reports"`
}
