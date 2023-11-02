# Assess

<!-- Start Codat Library Description -->
﻿Assess helps you make smarter credit decisions on small businesses by enabling you to pull your customers' latest data from the operating systems they are already using.
You can use that data for automating decisioning and surfacing new insights on the customer, all via one API.
<!-- End Codat Library Description -->

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/previous-versions/assess
```
<!-- End SDK Installation -->

## Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"log"
)

func main() {
	s := assess.New(
		assess.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.DataIntegrity.Details(ctx, operations.ListDataTypeDataIntegrityDetailsRequest{
		CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
		DataType:  shared.DataIntegrityDataTypeBankingAccounts,
		OrderBy:   assess.String("-modifiedDate"),
		Page:      assess.Int(1),
		PageSize:  assess.Int(100),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Details != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [DataIntegrity](docs/sdks/dataintegrity/README.md)

* [Details](docs/sdks/dataintegrity/README.md#details) - List data type data integrity
* [Status](docs/sdks/dataintegrity/README.md#status) - Get data integrity status
* [Summary](docs/sdks/dataintegrity/README.md#summary) - Get data integrity summary

### [ExcelReports](docs/sdks/excelreports/README.md)

* [GenerateExcelReport](docs/sdks/excelreports/README.md#generateexcelreport) - Generate Excel report
* [GetAccountingMarketingMetrics](docs/sdks/excelreports/README.md#getaccountingmarketingmetrics) - Get marketing metrics report
* [GetExcelReport](docs/sdks/excelreports/README.md#getexcelreport) - Download Excel report
* [GetExcelReportGenerationStatus](docs/sdks/excelreports/README.md#getexcelreportgenerationstatus) - Get Excel report status

### [Reports](docs/sdks/reports/README.md)

* [GenerateLoanSummary](docs/sdks/reports/README.md#generateloansummary) - Generate loan summaries report
* [GenerateLoanTransactions](docs/sdks/reports/README.md#generateloantransactions) - Generate loan transactions report
* [GetAccountsForEnhancedBalanceSheet](docs/sdks/reports/README.md#getaccountsforenhancedbalancesheet) - Get enhanced balance sheet accounts
* [GetAccountsForEnhancedProfitAndLoss](docs/sdks/reports/README.md#getaccountsforenhancedprofitandloss) - Get enhanced profit and loss accounts
* [GetCommerceCustomerRetentionMetrics](docs/sdks/reports/README.md#getcommercecustomerretentionmetrics) - Get customer retention metrics
* [GetCommerceLifetimeValueMetrics](docs/sdks/reports/README.md#getcommercelifetimevaluemetrics) - Get lifetime value metric
* [GetCommerceOrdersMetrics](docs/sdks/reports/README.md#getcommerceordersmetrics) - Get orders report
* [GetCommerceRefundsMetrics](docs/sdks/reports/README.md#getcommercerefundsmetrics) - Get refunds report
* [GetCommerceRevenueMetrics](docs/sdks/reports/README.md#getcommercerevenuemetrics) - Get commerce revenue metrics
* [GetEnhancedCashFlowTransactions](docs/sdks/reports/README.md#getenhancedcashflowtransactions) - Get enhanced cash flow report
* [GetEnhancedInvoicesReport](docs/sdks/reports/README.md#getenhancedinvoicesreport) - Get enhanced invoices report
* [GetLoanSummary](docs/sdks/reports/README.md#getloansummary) - Get loan summaries
* [GetRecurringRevenueMetrics](docs/sdks/reports/README.md#getrecurringrevenuemetrics) - Get key subscription revenue metrics
* [ListLoanTransactions](docs/sdks/reports/README.md#listloantransactions) - List loan transactions
* [RequestRecurringRevenueMetrics](docs/sdks/reports/README.md#requestrecurringrevenuemetrics) - Generate key subscription revenue metrics
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->



<!-- End Dev Containers -->

