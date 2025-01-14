# Lending

<!-- Start Codat Library Description -->
Lending helps you make smarter credit decisions on small businesses by enabling you to pull your customers' latest data from the operating systems they are already using. You can use that data for automating decisioning and surfacing new insights on the customer, all via one API.
<!-- End Codat Library Description -->

<!-- Start Summary [summary] -->
## Summary

Lending API: Our Lending API helps you make smarter credit decisions on small businesses by enabling you to pull your customers' latest data from accounting, banking, and commerce software they are already using. It also includes features to help providers verify the accuracy of data and process it more efficiently.

The Lending API is built on top of the latest accounting, commerce, and banking data, providing you with the most important data points you need to get a full picture of SMB creditworthiness and make a comprehensive assessment of your customers.

[Explore product](https://docs.codat.io/lending/overview) | [See OpenAPI spec](https://github.com/codatio/oas)

<!-- Start Codat Tags Table -->
## Endpoints

| Endpoints | Description |
| :- |:- |
| Companies | Create and manage your SMB users' companies. |
| Connections | Create new and manage existing data connections for a company. |
| Bank statements | Retrieve banking data from linked bank accounts. |
| Sales | Retrieve standardized sales data from a linked commerce software. |
| Financial statements | Financial data and reports from a linked accounting software. |
| Liabilities | Debt and other liabilities. |
| Accounts payable | Data from a linked accounting software representing money the business owes money to its suppliers. |
| Accounts receivable | Data from a linked accounting software representing money owed to the business for sold goods or services. |
| Transactions | Data from a linked accounting software representing transactions. |
| Company info | View company information fetched from the source platform. |
| Data integrity | Match mutable accounting data with immutable banking data to increase confidence in financial data. |
| Excel reports | Download reports in Excel format. |
| Manage data | Control how data is retrieved from an integration. |
| File upload | Endpoints to manage uploaded files. |
| Loan writeback | Implement the [loan writeback](https://docs.codat.io/lending/guides/loan-writeback/introduction) procedure in your lending process to maintain an accurate position of a loan during the entire lending cycle. |
<!-- End Codat Tags Table -->
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [Lending](#lending)
  * [Endpoints](#endpoints)
  * [SDK Installation](#sdk-installation)
  * [Example Usage](#example-usage)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Special Types](#special-types)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Authentication](#authentication)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/codatio/client-sdk-go/lending
```
<!-- End SDK Installation [installation] -->

## Example Usage
<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	lending "github.com/codatio/client-sdk-go/lending/v8"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := lending.New()

	res, err := s.AccountCategoriesUpdated(ctx, &shared.AccountCategoriesUpdatedWebhook{
		AlertID:    lending.String("a9367074-b5c3-42c4-9be4-be129f43577e"),
		ClientID:   lending.String("bae71d36-ff47-420a-b4a6-f8c9ddf41140"),
		ClientName: lending.String("Bank of Dave"),
		CompanyID:  lending.String("8a210b68-6988-11ed-a1eb-0242ac120002"),
		Data: &shared.AccountCategoriesUpdatedWebhookData{
			ModifiedDate: lending.String("2022-10-23"),
		},
		DataConnectionID: lending.String("2e9d2c44-f675-40ba-8049-353bfcb5e171"),
		Message:          lending.String("Account categories updated for company f1c35bdc-1546-41b9-baf4-3f31135af968."),
		RuleID:           lending.String("70af3071-65d9-4ec3-b3cb-5283e8d55dac"),
		RuleType:         lending.String("Account Categories Updated"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AccountingBankData](docs/sdks/codatlendingaccountingbankdata/README.md)

* [ListTransactions](docs/sdks/codatlendingaccountingbankdata/README.md#listtransactions) - List bank account transactions

#### [AccountingBankData.Accounts](docs/sdks/accounts/README.md)

* [Get](docs/sdks/accounts/README.md#get) - Get bank account
* [List](docs/sdks/accounts/README.md#list) - List bank accounts

### [AccountsPayable](docs/sdks/accountspayable/README.md)


#### [AccountsPayable.BillCreditNotes](docs/sdks/billcreditnotes/README.md)

* [Get](docs/sdks/billcreditnotes/README.md#get) - Get bill credit note
* [List](docs/sdks/billcreditnotes/README.md#list) - List bill credit notes

#### [AccountsPayable.BillPayments](docs/sdks/billpayments/README.md)

* [Get](docs/sdks/billpayments/README.md#get) - Get bill payment
* [List](docs/sdks/billpayments/README.md#list) - List bill payments

#### [AccountsPayable.Bills](docs/sdks/bills/README.md)

* [DownloadAttachment](docs/sdks/bills/README.md#downloadattachment) - Download bill attachment
* [Get](docs/sdks/bills/README.md#get) - Get bill
* [GetAttachment](docs/sdks/bills/README.md#getattachment) - Get bill attachment
* [List](docs/sdks/bills/README.md#list) - List bills
* [ListAttachments](docs/sdks/bills/README.md#listattachments) - List bill attachments

#### [AccountsPayable.Suppliers](docs/sdks/suppliers/README.md)

* [DownloadAttachment](docs/sdks/suppliers/README.md#downloadattachment) - Download supplier attachment
* [Get](docs/sdks/suppliers/README.md#get) - Get supplier
* [GetAttachment](docs/sdks/suppliers/README.md#getattachment) - Get supplier attachment
* [List](docs/sdks/suppliers/README.md#list) - List suppliers
* [ListAttachments](docs/sdks/suppliers/README.md#listattachments) - List supplier attachments

### [AccountsReceivable](docs/sdks/accountsreceivable/README.md)


#### [AccountsReceivable.CreditNotes](docs/sdks/creditnotes/README.md)

* [Get](docs/sdks/creditnotes/README.md#get) - Get credit note
* [List](docs/sdks/creditnotes/README.md#list) - List credit notes

#### [AccountsReceivable.Customers](docs/sdks/customers/README.md)

* [DownloadAttachment](docs/sdks/customers/README.md#downloadattachment) - Download customer attachment
* [Get](docs/sdks/customers/README.md#get) - Get customer
* [GetAttachment](docs/sdks/customers/README.md#getattachment) - Get customer attachment
* [List](docs/sdks/customers/README.md#list) - List customers
* [ListAttachments](docs/sdks/customers/README.md#listattachments) - List customer attachments

#### [AccountsReceivable.DirectIncomes](docs/sdks/directincomes/README.md)

* [DownloadAttachment](docs/sdks/directincomes/README.md#downloadattachment) - Download direct income attachment
* [Get](docs/sdks/directincomes/README.md#get) - Get direct income
* [GetAttachment](docs/sdks/directincomes/README.md#getattachment) - Get direct income attachment
* [List](docs/sdks/directincomes/README.md#list) - List direct incomes
* [ListAttachments](docs/sdks/directincomes/README.md#listattachments) - List direct income attachments

#### [AccountsReceivable.Invoices](docs/sdks/invoices/README.md)

* [DownloadAttachment](docs/sdks/invoices/README.md#downloadattachment) - Download invoice attachment
* [DownloadPdf](docs/sdks/invoices/README.md#downloadpdf) - Get invoice as PDF
* [Get](docs/sdks/invoices/README.md#get) - Get invoice
* [GetAttachment](docs/sdks/invoices/README.md#getattachment) - Get invoice attachment
* [List](docs/sdks/invoices/README.md#list) - List invoices
* [ListAttachments](docs/sdks/invoices/README.md#listattachments) - List invoice attachments
* [ListReconciled](docs/sdks/invoices/README.md#listreconciled) - List reconciled invoices

#### [AccountsReceivable.Payments](docs/sdks/payments/README.md)

* [Get](docs/sdks/payments/README.md#get) - Get payment
* [List](docs/sdks/payments/README.md#list) - List payments

#### [AccountsReceivable.Reports](docs/sdks/reports/README.md)

* [GetAgedCreditors](docs/sdks/reports/README.md#getagedcreditors) - Aged creditors report
* [GetAgedDebtors](docs/sdks/reports/README.md#getageddebtors) - Aged debtors report
* [IsAgedCreditorsAvailable](docs/sdks/reports/README.md#isagedcreditorsavailable) - Aged creditors report available
* [IsAgedDebtorsAvailable](docs/sdks/reports/README.md#isageddebtorsavailable) - Aged debtors report available

### [Banking](docs/sdks/banking/README.md)


#### [Banking.AccountBalances](docs/sdks/accountbalances/README.md)

* [List](docs/sdks/accountbalances/README.md#list) - List account balances

#### [Banking.Accounts](docs/sdks/codatlendingaccounts/README.md)

* [Get](docs/sdks/codatlendingaccounts/README.md#get) - Get account
* [List](docs/sdks/codatlendingaccounts/README.md#list) - List accounts

#### [Banking.CategorizedStatement](docs/sdks/categorizedstatement/README.md)

* [Get](docs/sdks/categorizedstatement/README.md#get) - Get categorized bank statement

#### [Banking.TransactionCategories](docs/sdks/transactioncategories/README.md)

* [Get](docs/sdks/transactioncategories/README.md#get) - Get transaction category
* [List](docs/sdks/transactioncategories/README.md#list) - List transaction categories

#### [Banking.Transactions](docs/sdks/codatlendingbankingtransactions/README.md)

* [Get](docs/sdks/codatlendingbankingtransactions/README.md#get) - Get bank transaction
* [List](docs/sdks/codatlendingbankingtransactions/README.md#list) - List transactions

### [BankStatements](docs/sdks/bankstatements/README.md)

* [EndUploadSession](docs/sdks/bankstatements/README.md#enduploadsession) - End upload session
* [GetUploadConfiguration](docs/sdks/bankstatements/README.md#getuploadconfiguration) - Get upload configuration
* [SetUploadConfiguration](docs/sdks/bankstatements/README.md#setuploadconfiguration) - Set upload configuration
* [StartUploadSession](docs/sdks/bankstatements/README.md#startuploadsession) - Start upload session
* [UploadBankStatementData](docs/sdks/bankstatements/README.md#uploadbankstatementdata) - Upload data


### [Companies](docs/sdks/companies/README.md)

* [Create](docs/sdks/companies/README.md#create) - Create company
* [Delete](docs/sdks/companies/README.md#delete) - Delete a company
* [Get](docs/sdks/companies/README.md#get) - Get company
* [List](docs/sdks/companies/README.md#list) - List companies
* [Update](docs/sdks/companies/README.md#update) - Update company

### [CompanyInfo](docs/sdks/companyinfo/README.md)

* [GetAccountingProfile](docs/sdks/companyinfo/README.md#getaccountingprofile) - Get company accounting profile
* [GetCommerceProfile](docs/sdks/companyinfo/README.md#getcommerceprofile) - Get company commerce profile

### [Connections](docs/sdks/connections/README.md)

* [Create](docs/sdks/connections/README.md#create) - Create connection
* [Delete](docs/sdks/connections/README.md#delete) - Delete connection
* [Get](docs/sdks/connections/README.md#get) - Get connection
* [List](docs/sdks/connections/README.md#list) - List connections
* [Unlink](docs/sdks/connections/README.md#unlink) - Unlink connection

### [DataIntegrity](docs/sdks/dataintegrity/README.md)

* [Details](docs/sdks/dataintegrity/README.md#details) - List data integrity details
* [Status](docs/sdks/dataintegrity/README.md#status) - Get data integrity status
* [Summaries](docs/sdks/dataintegrity/README.md#summaries) - Get data integrity summaries

### [ExcelReports](docs/sdks/excelreports/README.md)

* [Download](docs/sdks/excelreports/README.md#download) - Download Excel report
* [Generate](docs/sdks/excelreports/README.md#generate) - Generate Excel report
* [GetStatus](docs/sdks/excelreports/README.md#getstatus) - Get Excel report status

### [FileUpload](docs/sdks/fileupload/README.md)

* [Download](docs/sdks/fileupload/README.md#download) - Download all files for a company
* [ListUploaded](docs/sdks/fileupload/README.md#listuploaded) - List all files uploaded by a company
* [Upload](docs/sdks/fileupload/README.md#upload) - Upload files for a company

### [FinancialStatements](docs/sdks/financialstatements/README.md)


#### [FinancialStatements.Accounts](docs/sdks/codatlendingfinancialstatementsaccounts/README.md)

* [Get](docs/sdks/codatlendingfinancialstatementsaccounts/README.md#get) - Get account
* [List](docs/sdks/codatlendingfinancialstatementsaccounts/README.md#list) - List accounts

#### [FinancialStatements.BalanceSheet](docs/sdks/balancesheet/README.md)

* [Get](docs/sdks/balancesheet/README.md#get) - Get balance sheet
* [GetCategorizedAccounts](docs/sdks/balancesheet/README.md#getcategorizedaccounts) - Get categorized balance sheet statement

#### [FinancialStatements.CashFlow](docs/sdks/cashflow/README.md)

* [Get](docs/sdks/cashflow/README.md#get) - Get cash flow statement

#### [FinancialStatements.ProfitAndLoss](docs/sdks/profitandloss/README.md)

* [Get](docs/sdks/profitandloss/README.md#get) - Get profit and loss
* [GetCategorizedAccounts](docs/sdks/profitandloss/README.md#getcategorizedaccounts) - Get categorized profit and loss statement

### [Liabilities](docs/sdks/liabilities/README.md)

* [GenerateLoanSummary](docs/sdks/liabilities/README.md#generateloansummary) - Generate loan summaries report
* [GenerateLoanTransactions](docs/sdks/liabilities/README.md#generateloantransactions) - Generate loan transactions report
* [GetLoanSummary](docs/sdks/liabilities/README.md#getloansummary) - Get loan summaries
* [ListLoanTransactions](docs/sdks/liabilities/README.md#listloantransactions) - List loan transactions

### [LoanWriteback](docs/sdks/loanwriteback/README.md)


#### [LoanWriteback.Accounts](docs/sdks/codatlendingloanwritebackaccounts/README.md)

* [Create](docs/sdks/codatlendingloanwritebackaccounts/README.md#create) - Create account
* [GetCreateModel](docs/sdks/codatlendingloanwritebackaccounts/README.md#getcreatemodel) - Get create account model

#### [LoanWriteback.BankAccounts](docs/sdks/bankaccounts/README.md)

* [Create](docs/sdks/bankaccounts/README.md#create) - Create bank account
* [GetCreateUpdateModel](docs/sdks/bankaccounts/README.md#getcreateupdatemodel) - Get create/update bank account model

#### [LoanWriteback.BankTransactions](docs/sdks/banktransactions/README.md)

* [Create](docs/sdks/banktransactions/README.md#create) - Create bank account transactions
* [GetCreateModel](docs/sdks/banktransactions/README.md#getcreatemodel) - Get create bank account transactions model

#### [LoanWriteback.CreateOperations](docs/sdks/createoperations/README.md)

* [Get](docs/sdks/createoperations/README.md#get) - Get create operation
* [List](docs/sdks/createoperations/README.md#list) - List create operations

#### [LoanWriteback.DirectCosts](docs/sdks/directcosts/README.md)

* [Create](docs/sdks/directcosts/README.md#create) - Create direct cost
* [GetCreateModel](docs/sdks/directcosts/README.md#getcreatemodel) - Get create direct cost model

#### [LoanWriteback.Payments](docs/sdks/codatlendingpayments/README.md)

* [Create](docs/sdks/codatlendingpayments/README.md#create) - Create payment
* [GetCreateModel](docs/sdks/codatlendingpayments/README.md#getcreatemodel) - Get create payment model

#### [LoanWriteback.SourceAccounts](docs/sdks/sourceaccounts/README.md)

* [Create](docs/sdks/sourceaccounts/README.md#create) - Create source account
* [CreateMapping](docs/sdks/sourceaccounts/README.md#createmapping) - Create bank feed account mapping
* [ListMappings](docs/sdks/sourceaccounts/README.md#listmappings) - List bank feed account mappings

#### [LoanWriteback.Suppliers](docs/sdks/codatlendingsuppliers/README.md)

* [Create](docs/sdks/codatlendingsuppliers/README.md#create) - Create supplier
* [GetCreateUpdateModel](docs/sdks/codatlendingsuppliers/README.md#getcreateupdatemodel) - Get create/update supplier model

#### [LoanWriteback.Transfers](docs/sdks/transfers/README.md)

* [Create](docs/sdks/transfers/README.md#create) - Create transfer
* [GetCreateModel](docs/sdks/transfers/README.md#getcreatemodel) - Get create transfer model

### [ManageData](docs/sdks/managedata/README.md)

* [GetStatus](docs/sdks/managedata/README.md#getstatus) - Get data status

#### [ManageData.PullOperations](docs/sdks/pulloperations/README.md)

* [Get](docs/sdks/pulloperations/README.md#get) - Get pull operation
* [List](docs/sdks/pulloperations/README.md#list) - List pull operations

#### [ManageData.Refresh](docs/sdks/refresh/README.md)

* [AllDataTypes](docs/sdks/refresh/README.md#alldatatypes) - Refresh all data
* [DataType](docs/sdks/refresh/README.md#datatype) - Refresh data type

### [ManageReports](docs/sdks/managereports/README.md)

* [GenerateReport](docs/sdks/managereports/README.md#generatereport) - Generate report
* [ListReports](docs/sdks/managereports/README.md#listreports) - List reports

### [Sales](docs/sdks/sales/README.md)


#### [Sales.Customers](docs/sdks/codatlendingcustomers/README.md)

* [Get](docs/sdks/codatlendingcustomers/README.md#get) - Get customer
* [List](docs/sdks/codatlendingcustomers/README.md#list) - List customers

#### [Sales.Disputes](docs/sdks/disputes/README.md)

* [Get](docs/sdks/disputes/README.md#get) - Get dispute
* [List](docs/sdks/disputes/README.md#list) - List disputes

#### [Sales.Locations](docs/sdks/locations/README.md)

* [Get](docs/sdks/locations/README.md#get) - Get location
* [List](docs/sdks/locations/README.md#list) - List locations

#### [Sales.Metrics](docs/sdks/metrics/README.md)

* [GetCustomerRetention](docs/sdks/metrics/README.md#getcustomerretention) - Get customer retention metrics
* [GetLifetimeValue](docs/sdks/metrics/README.md#getlifetimevalue) - Get lifetime value metrics
* [GetRevenue](docs/sdks/metrics/README.md#getrevenue) - Get commerce revenue metrics

#### [Sales.Orders](docs/sdks/orders/README.md)

* [Get](docs/sdks/orders/README.md#get) - Get order
* [List](docs/sdks/orders/README.md#list) - List orders

#### [Sales.PaymentMethods](docs/sdks/paymentmethods/README.md)

* [Get](docs/sdks/paymentmethods/README.md#get) - Get payment method
* [List](docs/sdks/paymentmethods/README.md#list) - List payment methods

#### [Sales.Payments](docs/sdks/codatlendingsalespayments/README.md)

* [Get](docs/sdks/codatlendingsalespayments/README.md#get) - Get payment
* [List](docs/sdks/codatlendingsalespayments/README.md#list) - List payments

#### [Sales.ProductCategories](docs/sdks/productcategories/README.md)

* [Get](docs/sdks/productcategories/README.md#get) - Get product category
* [List](docs/sdks/productcategories/README.md#list) - List product categories

#### [Sales.Products](docs/sdks/products/README.md)

* [Get](docs/sdks/products/README.md#get) - Get product
* [List](docs/sdks/products/README.md#list) - List products

#### [Sales.Reports](docs/sdks/codatlendingreports/README.md)

* [GetOrders](docs/sdks/codatlendingreports/README.md#getorders) - Get orders report
* [GetRefunds](docs/sdks/codatlendingreports/README.md#getrefunds) - Get refunds report

#### [Sales.Transactions](docs/sdks/codatlendingtransactions/README.md)

* [Get](docs/sdks/codatlendingtransactions/README.md#get) - Get transaction
* [List](docs/sdks/codatlendingtransactions/README.md#list) - List transactions

### [Transactions](docs/sdks/transactions/README.md)


#### [Transactions.AccountTransactions](docs/sdks/accounttransactions/README.md)

* [Get](docs/sdks/accounttransactions/README.md#get) - Get account transaction
* [List](docs/sdks/accounttransactions/README.md#list) - List account transactions

#### [Transactions.DirectCosts](docs/sdks/codatlendingdirectcosts/README.md)

* [DownloadAttachment](docs/sdks/codatlendingdirectcosts/README.md#downloadattachment) - Download direct cost attachment
* [Get](docs/sdks/codatlendingdirectcosts/README.md#get) - Get direct cost
* [GetAttachment](docs/sdks/codatlendingdirectcosts/README.md#getattachment) - Get direct cost attachment
* [List](docs/sdks/codatlendingdirectcosts/README.md#list) - List direct costs
* [ListAttachments](docs/sdks/codatlendingdirectcosts/README.md#listattachments) - List direct cost attachments

#### [Transactions.JournalEntries](docs/sdks/journalentries/README.md)

* [Get](docs/sdks/journalentries/README.md#get) - Get journal entry
* [List](docs/sdks/journalentries/README.md#list) - List journal entries

#### [Transactions.Journals](docs/sdks/journals/README.md)

* [Get](docs/sdks/journals/README.md#get) - Get journal
* [List](docs/sdks/journals/README.md#list) - List journals

#### [Transactions.Transfers](docs/sdks/codatlendingtransfers/README.md)

* [Get](docs/sdks/codatlendingtransfers/README.md#get) - Get transfer
* [List](docs/sdks/codatlendingtransfers/README.md#list) - List transfers

</details>
<!-- End Available Resources and Operations [operations] -->





<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	lending "github.com/codatio/client-sdk-go/lending/v8"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/retry"
	"log"
	"pkg/models/operations"
)

func main() {
	ctx := context.Background()

	s := lending.New(
		lending.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: lending.String("Requested early access to the new financing scheme."),
		Name:        "Technicalium",
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Company != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	lending "github.com/codatio/client-sdk-go/lending/v8"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := lending.New(
		lending.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		lending.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: lending.String("Requested early access to the new financing scheme."),
		Name:        "Technicalium",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Company != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Create` function may return the following errors:

| Error Type             | Status Code                       | Content Type     |
| ---------------------- | --------------------------------- | ---------------- |
| sdkerrors.ErrorMessage | 400, 401, 402, 403, 429, 500, 503 | application/json |
| sdkerrors.SDKError     | 4XX, 5XX                          | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	lending "github.com/codatio/client-sdk-go/lending/v8"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/sdkerrors"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := lending.New(
		lending.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: lending.String("Requested early access to the new financing scheme."),
		Name:        "Technicalium",
	})
	if err != nil {

		var e *sdkerrors.ErrorMessage
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	lending "github.com/codatio/client-sdk-go/lending/v8"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := lending.New(
		lending.WithServerURL("https://api.codat.io"),
		lending.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: lending.String("Requested early access to the new financing scheme."),
		Name:        "Technicalium",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Company != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type   | Scheme  |
| ------------ | ------ | ------- |
| `AuthHeader` | apiKey | API key |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	lending "github.com/codatio/client-sdk-go/lending/v8"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := lending.New(
		lending.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: lending.String("Requested early access to the new financing scheme."),
		Name:        "Technicalium",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Company != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

<!-- Start Codat Support Notes -->
### Support

If you encounter any challenges while utilizing our SDKs, please don't hesitate to reach out for assistance. 
You can raise any issues by contacting your dedicated Codat representative or reaching out to our [support team](mailto:support@codat.io).
We're here to help ensure a smooth experience for you.
<!-- End Codat Support Notes -->

<!-- Start Codat Generated By -->
### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
<!-- End Codat Generated By -->