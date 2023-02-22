package shared

import (
	"time"
)

// CashFlowStatement
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
type CashFlowStatement struct {
	CashPayments *Assets    `json:"cashPayments,omitempty"`
	CashReceipts *Assets    `json:"cashReceipts,omitempty"`
	FromDate     *time.Time `json:"fromDate,omitempty"`
	ToDate       *time.Time `json:"toDate,omitempty"`
}
