# Assess

Assess helps you make smarter credit decisions on small businesses by enabling you to pull your customers' latest data from the operating systems they are already using.
You can use that data for automating decisioning and surfacing new insights on the customer, all via one API.

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/assess
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Categories.GetAccountCategory(ctx, operations.GetAccountCategoryRequest{
        AccountID: "corrupti",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CategorisedAccount != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Categories](docs/categories/README.md)

* [~~GetAccountCategory~~](docs/categories/README.md#getaccountcategory) - Get suggested and/or confirmed category for a specific account :warning: **Deprecated**
* [~~ListAccountsCategories~~](docs/categories/README.md#listaccountscategories) - List suggested and confirmed account categories :warning: **Deprecated**
* [~~ListAvailableAccountCategories~~](docs/categories/README.md#listavailableaccountcategories) - List account categories :warning: **Deprecated**
* [~~UpdateAccountCategory~~](docs/categories/README.md#updateaccountcategory) - Patch account categories :warning: **Deprecated**
* [~~UpdateAccountsCategories~~](docs/categories/README.md#updateaccountscategories) - Confirm categories for accounts :warning: **Deprecated**

### [DataIntegrity](docs/dataintegrity/README.md)

* [GetDataIntegrityDetails](docs/dataintegrity/README.md#getdataintegritydetails) - Lists data integrity details for date type
* [GetDataIntegrityStatus](docs/dataintegrity/README.md#getdataintegritystatus) - Get data integrity status
* [GetDataIntegritySummaries](docs/dataintegrity/README.md#getdataintegritysummaries) - Get data integrity summary

### [ExcelReports](docs/excelreports/README.md)

* [~~DownloadExcelReport~~](docs/excelreports/README.md#downloadexcelreport) - Download generated excel report :warning: **Deprecated**
* [GenerateExcelReport](docs/excelreports/README.md#generateexcelreport) - Generate an Excel report
* [GetAccountingMarketingMetrics](docs/excelreports/README.md#getaccountingmarketingmetrics) - Get the marketing metrics from an accounting source for a given company.
* [GetExcelReport](docs/excelreports/README.md#getexcelreport) - Download generated excel report
* [GetExcelReportGenerationStatus](docs/excelreports/README.md#getexcelreportgenerationstatus) - Get status of Excel report

### [Reports](docs/reports/README.md)

* [GetAccountsForEnhancedBalanceSheet](docs/reports/README.md#getaccountsforenhancedbalancesheet) - Enhanced Balance Sheet Accounts
* [GetAccountsForEnhancedProfitAndLoss](docs/reports/README.md#getaccountsforenhancedprofitandloss) - Enhanced Profit and Loss Accounts
* [GetCommerceCustomerRetentionMetrics](docs/reports/README.md#getcommercecustomerretentionmetrics) - Get the customer retention metrics for a specific company.
* [GetCommerceLifetimeValueMetrics](docs/reports/README.md#getcommercelifetimevaluemetrics) - Get the lifetime value metric for a specific company.
* [GetCommerceOrdersMetrics](docs/reports/README.md#getcommerceordersmetrics) - Get order information for a specific company
* [GetCommerceRefundsMetrics](docs/reports/README.md#getcommercerefundsmetrics) - Get the refunds information for a specific company
* [GetCommerceRevenueMetrics](docs/reports/README.md#getcommercerevenuemetrics) - Commerce Revenue Metrics
* [~~GetEnhancedBalanceSheet~~](docs/reports/README.md#getenhancedbalancesheet) - Enhanced Balance Sheet :warning: **Deprecated**
* [GetEnhancedCashFlowTransactions](docs/reports/README.md#getenhancedcashflowtransactions) - Get enhanced cash flow report
* [~~GetEnhancedFinancialMetrics~~](docs/reports/README.md#getenhancedfinancialmetrics) - List financial metrics :warning: **Deprecated**
* [GetEnhancedInvoicesReport](docs/reports/README.md#getenhancedinvoicesreport) - Enhanced Invoices Report
* [~~GetEnhancedProfitAndLoss~~](docs/reports/README.md#getenhancedprofitandloss) - Enhanced Profit and Loss :warning: **Deprecated**
* [GetRecurringRevenueMetrics](docs/reports/README.md#getrecurringrevenuemetrics) - Get key metrics for subscription revenue
* [RequestRecurringRevenueMetrics](docs/reports/README.md#requestrecurringrevenuemetrics) - Request production of key subscription revenue metrics
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
