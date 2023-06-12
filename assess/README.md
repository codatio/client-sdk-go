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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
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


### [Categories](docs/sdks/categories/README.md)

* [~~GetAccountCategory~~](docs/sdks/categories/README.md#getaccountcategory) - Get suggested and/or confirmed category for a specific account :warning: **Deprecated**
* [~~ListAccountsCategories~~](docs/sdks/categories/README.md#listaccountscategories) - List suggested and confirmed account categories :warning: **Deprecated**
* [~~ListAvailableAccountCategories~~](docs/sdks/categories/README.md#listavailableaccountcategories) - List account categories :warning: **Deprecated**
* [~~UpdateAccountCategory~~](docs/sdks/categories/README.md#updateaccountcategory) - Update account categories :warning: **Deprecated**
* [~~UpdateAccountsCategories~~](docs/sdks/categories/README.md#updateaccountscategories) - Confirm categories for accounts :warning: **Deprecated**

### [DataIntegrity](docs/sdks/dataintegrity/README.md)

* [GetDataIntegrityStatus](docs/sdks/dataintegrity/README.md#getdataintegritystatus) - Get data integrity status
* [GetDataIntegritySummaries](docs/sdks/dataintegrity/README.md#getdataintegritysummaries) - Get data integrity summary
* [ListDataTypeDataIntegrityDetails](docs/sdks/dataintegrity/README.md#listdatatypedataintegritydetails) - List data type data integrity

### [ExcelReports](docs/sdks/excelreports/README.md)

* [GenerateExcelReport](docs/sdks/excelreports/README.md#generateexcelreport) - Generate Excel report
* [GetAccountingMarketingMetrics](docs/sdks/excelreports/README.md#getaccountingmarketingmetrics) - Get marketing metrics report
* [GetExcelReport](docs/sdks/excelreports/README.md#getexcelreport) - Download Excel report
* [GetExcelReportGenerationStatus](docs/sdks/excelreports/README.md#getexcelreportgenerationstatus) - Get Excel report status

### [Reports](docs/sdks/reports/README.md)

* [GetAccountsForEnhancedBalanceSheet](docs/sdks/reports/README.md#getaccountsforenhancedbalancesheet) - Get enhanced balance sheet accounts
* [GetAccountsForEnhancedProfitAndLoss](docs/sdks/reports/README.md#getaccountsforenhancedprofitandloss) - Get enhanced profit and loss accounts
* [GetCommerceCustomerRetentionMetrics](docs/sdks/reports/README.md#getcommercecustomerretentionmetrics) - Get customer retention metrics
* [GetCommerceLifetimeValueMetrics](docs/sdks/reports/README.md#getcommercelifetimevaluemetrics) - Get lifetime value metric
* [GetCommerceOrdersMetrics](docs/sdks/reports/README.md#getcommerceordersmetrics) - Get orders report
* [GetCommerceRefundsMetrics](docs/sdks/reports/README.md#getcommercerefundsmetrics) - Get refunds report
* [GetCommerceRevenueMetrics](docs/sdks/reports/README.md#getcommercerevenuemetrics) - Get commerce revenue metrics
* [~~GetEnhancedBalanceSheet~~](docs/sdks/reports/README.md#getenhancedbalancesheet) - Get enhanced balance sheet report :warning: **Deprecated**
* [GetEnhancedCashFlowTransactions](docs/sdks/reports/README.md#getenhancedcashflowtransactions) - Get enhanced cash flow report
* [~~GetEnhancedFinancialMetrics~~](docs/sdks/reports/README.md#getenhancedfinancialmetrics) - List financial metrics :warning: **Deprecated**
* [GetEnhancedInvoicesReport](docs/sdks/reports/README.md#getenhancedinvoicesreport) - Get enhanced invoices report
* [~~GetEnhancedProfitAndLoss~~](docs/sdks/reports/README.md#getenhancedprofitandloss) - Get enhanced profit and loss report :warning: **Deprecated**
* [GetRecurringRevenueMetrics](docs/sdks/reports/README.md#getrecurringrevenuemetrics) - Get key subscription revenue metrics
* [RequestRecurringRevenueMetrics](docs/sdks/reports/README.md#requestrecurringrevenuemetrics) - Generate key subscription revenue metrics
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
