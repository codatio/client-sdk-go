// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type ProfitAndLossReport struct {
	CostOfSales *ReportLine `json:"costOfSales,omitempty"`
	Expenses    *ReportLine `json:"expenses,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	FromDate *string `json:"fromDate,omitempty"`
	// Gross profit of the company in the given date range.
	GrossProfit *decimal.Big `decimal:"number" json:"grossProfit"`
	Income      *ReportLine  `json:"income,omitempty"`
	// Net operating profit of the company in the given date range.
	NetOperatingProfit *decimal.Big `decimal:"number" json:"netOperatingProfit"`
	// Net other income of the company in the given date range.
	NetOtherIncome *decimal.Big `decimal:"number" json:"netOtherIncome"`
	// Net profit of the company in the given date range.
	NetProfit     *decimal.Big `decimal:"number" json:"netProfit"`
	OtherExpenses *ReportLine  `json:"otherExpenses,omitempty"`
	OtherIncome   *ReportLine  `json:"otherIncome,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	ToDate *string `json:"toDate,omitempty"`
}

func (p ProfitAndLossReport) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProfitAndLossReport) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ProfitAndLossReport) GetCostOfSales() *ReportLine {
	if o == nil {
		return nil
	}
	return o.CostOfSales
}

func (o *ProfitAndLossReport) GetExpenses() *ReportLine {
	if o == nil {
		return nil
	}
	return o.Expenses
}

func (o *ProfitAndLossReport) GetFromDate() *string {
	if o == nil {
		return nil
	}
	return o.FromDate
}

func (o *ProfitAndLossReport) GetGrossProfit() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.GrossProfit
}

func (o *ProfitAndLossReport) GetIncome() *ReportLine {
	if o == nil {
		return nil
	}
	return o.Income
}

func (o *ProfitAndLossReport) GetNetOperatingProfit() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.NetOperatingProfit
}

func (o *ProfitAndLossReport) GetNetOtherIncome() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.NetOtherIncome
}

func (o *ProfitAndLossReport) GetNetProfit() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.NetProfit
}

func (o *ProfitAndLossReport) GetOtherExpenses() *ReportLine {
	if o == nil {
		return nil
	}
	return o.OtherExpenses
}

func (o *ProfitAndLossReport) GetOtherIncome() *ReportLine {
	if o == nil {
		return nil
	}
	return o.OtherIncome
}

func (o *ProfitAndLossReport) GetToDate() *string {
	if o == nil {
		return nil
	}
	return o.ToDate
}
