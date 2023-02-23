package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type GetCashFlowStatementPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCashFlowStatementQueryParams struct {
	PeriodLength     int        `queryParam:"style=form,explode=true,name=periodLength"`
	PeriodsToCompare int        `queryParam:"style=form,explode=true,name=periodsToCompare"`
	StartMonth       *time.Time `queryParam:"style=form,explode=true,name=startMonth"`
}

type GetCashFlowStatementSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetCashFlowStatementRequest struct {
	PathParams  GetCashFlowStatementPathParams
	QueryParams GetCashFlowStatementQueryParams
	Security    GetCashFlowStatementSecurity
}

type GetCashFlowStatement200ApplicationJSONReportBasisEnum string

const (
	GetCashFlowStatement200ApplicationJSONReportBasisEnumUnknown GetCashFlowStatement200ApplicationJSONReportBasisEnum = "Unknown"
	GetCashFlowStatement200ApplicationJSONReportBasisEnumAccrual GetCashFlowStatement200ApplicationJSONReportBasisEnum = "Accrual"
	GetCashFlowStatement200ApplicationJSONReportBasisEnumCash    GetCashFlowStatement200ApplicationJSONReportBasisEnum = "Cash"
)

type GetCashFlowStatement200ApplicationJSONReportInputEnum string

const (
	GetCashFlowStatement200ApplicationJSONReportInputEnumUnknown  GetCashFlowStatement200ApplicationJSONReportInputEnum = "Unknown"
	GetCashFlowStatement200ApplicationJSONReportInputEnumIndirect GetCashFlowStatement200ApplicationJSONReportInputEnum = "Indirect"
	GetCashFlowStatement200ApplicationJSONReportInputEnumDirect   GetCashFlowStatement200ApplicationJSONReportInputEnum = "Direct"
)

type GetCashFlowStatement200ApplicationJSONCashFlowStatementReportLineReportLineReportLineReportLine struct {
	AccountID *string `json:"accountId,omitempty"`
	Name      *string `json:"name,omitempty"`
	Value     float64 `json:"value"`
}

type GetCashFlowStatement200ApplicationJSONCashFlowStatementReportLineReportLineReportLine struct {
	AccountID *string                                                                                           `json:"accountId,omitempty"`
	Items     []GetCashFlowStatement200ApplicationJSONCashFlowStatementReportLineReportLineReportLineReportLine `json:"items,omitempty"`
	Name      *string                                                                                           `json:"name,omitempty"`
	Value     float64                                                                                           `json:"value"`
}

type GetCashFlowStatement200ApplicationJSONCashFlowStatementReportLineReportLine struct {
	AccountID *string                                                                                 `json:"accountId,omitempty"`
	Items     []GetCashFlowStatement200ApplicationJSONCashFlowStatementReportLineReportLineReportLine `json:"items,omitempty"`
	Name      *string                                                                                 `json:"name,omitempty"`
	Value     float64                                                                                 `json:"value"`
}

// GetCashFlowStatement200ApplicationJSONCashFlowStatementReportLine
// ReportLines for cash payments to suppliers for the purchase of goods or services.
type GetCashFlowStatement200ApplicationJSONCashFlowStatementReportLine struct {
	AccountID *string                                                                       `json:"accountId,omitempty"`
	Items     []GetCashFlowStatement200ApplicationJSONCashFlowStatementReportLineReportLine `json:"items,omitempty"`
	Name      *string                                                                       `json:"name,omitempty"`
	Value     float64                                                                       `json:"value"`
}

// GetCashFlowStatement200ApplicationJSONCashFlowStatement
// > View the coverage for cash flow statement in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=cashFlowStatement" target="_blank">Data coverage explorer</a>.
//
// > ** Operating activities only**
// >
// > Currently, the cash flow statement shows cash that flows into and out of the company from operating activities *only*. Operating activities generate cash from the sale of goods or services.
//
// ## Overview
//
// A cash flow statement is a financial report that records all cash that is received or spent by a company during a given period. It gives you a clearer picture of the company’s performance, and their ability to pay creditors and finance growth.
//
// > **Cash flow statement or balance sheet?**
// >
// > Look at the cash flow statement to understand a company's ability to pay its bills. Although the balance sheet may show healthy earnings at a specific point in time, the cash flow statement allows you to see whether the company is meeting its financial commitments, such as paying creditors or its employees.
type GetCashFlowStatement200ApplicationJSONCashFlowStatement struct {
	CashPayments *GetCashFlowStatement200ApplicationJSONCashFlowStatementReportLine `json:"cashPayments,omitempty"`
	CashReceipts *GetCashFlowStatement200ApplicationJSONCashFlowStatementReportLine `json:"cashReceipts,omitempty"`
	FromDate     *time.Time                                                         `json:"fromDate,omitempty"`
	ToDate       *time.Time                                                         `json:"toDate,omitempty"`
}

type GetCashFlowStatement200ApplicationJSON struct {
	Currency                 string                                                    `json:"currency"`
	EarliestAvailableMonth   *time.Time                                                `json:"earliestAvailableMonth,omitempty"`
	MostRecentAvailableMonth *time.Time                                                `json:"mostRecentAvailableMonth,omitempty"`
	ReportBasis              GetCashFlowStatement200ApplicationJSONReportBasisEnum     `json:"reportBasis"`
	ReportInput              GetCashFlowStatement200ApplicationJSONReportInputEnum     `json:"reportInput"`
	Reports                  []GetCashFlowStatement200ApplicationJSONCashFlowStatement `json:"reports"`
}

type GetCashFlowStatementResponse struct {
	ContentType                                  string
	StatusCode                                   int
	GetCashFlowStatement200ApplicationJSONObject *GetCashFlowStatement200ApplicationJSON
}
