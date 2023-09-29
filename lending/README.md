# Lending
    
Lending helps you make smarter credit decisions on small businesses by enabling you to pull your customers' latest data from the operating systems they are already using. You can use that data for automating decisioning and surfacing new insights on the customer, all via one API.

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/lending
```
<!-- End SDK Installation -->

## Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	lending "github.com/codatio/client-sdk-go/lending/v4"
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/operations"
)

func main() {
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountingBankData.ListTransactions(ctx, operations.ListAccountingBankAccountTransactionsRequest{
        AccountID: "Anchorage Product",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: lending.String("-modifiedDate"),
        Page: lending.Int(1),
        PageSize: lending.Int(100),
        Query: lending.String("Future"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountingBankTransactions != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [AccountingBankData](docs/sdks/accountingbankdata/README.md)

* [ListTransactions](docs/sdks/accountingbankdata/README.md#listtransactions) - List bank account transactions

### [AccountingBankData.Accounts](docs/sdks/accountingbankdataaccounts/README.md)

* [Get](docs/sdks/accountingbankdataaccounts/README.md#get) - Get bank account
* [List](docs/sdks/accountingbankdataaccounts/README.md#list) - List bank accounts

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

### [Liabilities](docs/sdks/liabilities/README.md)

* [GenerateLoanSummary](docs/sdks/liabilities/README.md#generateloansummary) - Generate loan summaries report
* [GenerateLoanTransactions](docs/sdks/liabilities/README.md#generateloantransactions) - Generate loan transactions report
* [GetLoanSummary](docs/sdks/liabilities/README.md#getloansummary) - Get loan summaries
* [ListLoanTransactions](docs/sdks/liabilities/README.md#listloantransactions) - List loan transactions


### [AccountsPayable.BillCreditNotes](docs/sdks/accountspayablebillcreditnotes/README.md)

* [Get](docs/sdks/accountspayablebillcreditnotes/README.md#get) - Get bill credit note
* [List](docs/sdks/accountspayablebillcreditnotes/README.md#list) - List bill credit notes

### [AccountsPayable.BillPayments](docs/sdks/accountspayablebillpayments/README.md)

* [Get](docs/sdks/accountspayablebillpayments/README.md#get) - Get bill payment
* [List](docs/sdks/accountspayablebillpayments/README.md#list) - List bill payments

### [AccountsPayable.Bills](docs/sdks/accountspayablebills/README.md)

* [DownloadAttachment](docs/sdks/accountspayablebills/README.md#downloadattachment) - Download bill attachment
* [Get](docs/sdks/accountspayablebills/README.md#get) - Get bill
* [GetAttachment](docs/sdks/accountspayablebills/README.md#getattachment) - Get bill attachment
* [List](docs/sdks/accountspayablebills/README.md#list) - List bills
* [ListAttachments](docs/sdks/accountspayablebills/README.md#listattachments) - List bill attachments

### [AccountsPayable.Suppliers](docs/sdks/accountspayablesuppliers/README.md)

* [DownloadAttachment](docs/sdks/accountspayablesuppliers/README.md#downloadattachment) - Download supplier attachment
* [Get](docs/sdks/accountspayablesuppliers/README.md#get) - Get supplier
* [GetAttachment](docs/sdks/accountspayablesuppliers/README.md#getattachment) - Get supplier attachment
* [List](docs/sdks/accountspayablesuppliers/README.md#list) - List suppliers
* [ListAttachments](docs/sdks/accountspayablesuppliers/README.md#listattachments) - List supplier attachments


### [AccountsReceivable.CreditNotes](docs/sdks/accountsreceivablecreditnotes/README.md)

* [Get](docs/sdks/accountsreceivablecreditnotes/README.md#get) - Get credit note
* [List](docs/sdks/accountsreceivablecreditnotes/README.md#list) - List credit notes

### [AccountsReceivable.Customers](docs/sdks/accountsreceivablecustomers/README.md)

* [DownloadAttachment](docs/sdks/accountsreceivablecustomers/README.md#downloadattachment) - Download customer attachment
* [Get](docs/sdks/accountsreceivablecustomers/README.md#get) - Get customer
* [GetAttachment](docs/sdks/accountsreceivablecustomers/README.md#getattachment) - Get customer attachment
* [List](docs/sdks/accountsreceivablecustomers/README.md#list) - List customers
* [ListAttachments](docs/sdks/accountsreceivablecustomers/README.md#listattachments) - List customer attachments

### [AccountsReceivable.DirectIncomes](docs/sdks/accountsreceivabledirectincomes/README.md)

* [DownloadAttachment](docs/sdks/accountsreceivabledirectincomes/README.md#downloadattachment) - Download direct income attachment
* [Get](docs/sdks/accountsreceivabledirectincomes/README.md#get) - Get direct income
* [GetAttachment](docs/sdks/accountsreceivabledirectincomes/README.md#getattachment) - Get direct income attachment
* [List](docs/sdks/accountsreceivabledirectincomes/README.md#list) - List direct incomes
* [ListAttachments](docs/sdks/accountsreceivabledirectincomes/README.md#listattachments) - List direct income attachments

### [AccountsReceivable.Invoices](docs/sdks/accountsreceivableinvoices/README.md)

* [DownloadAttachment](docs/sdks/accountsreceivableinvoices/README.md#downloadattachment) - Download invoice attachment
* [DownloadPdf](docs/sdks/accountsreceivableinvoices/README.md#downloadpdf) - Get invoice as PDF
* [Get](docs/sdks/accountsreceivableinvoices/README.md#get) - Get invoice
* [GetAttachment](docs/sdks/accountsreceivableinvoices/README.md#getattachment) - Get invoice attachment
* [List](docs/sdks/accountsreceivableinvoices/README.md#list) - List invoices
* [ListAttachments](docs/sdks/accountsreceivableinvoices/README.md#listattachments) - List invoice attachments
* [ListReconciled](docs/sdks/accountsreceivableinvoices/README.md#listreconciled) - List reconciled invoices

### [AccountsReceivable.Payments](docs/sdks/accountsreceivablepayments/README.md)

* [Get](docs/sdks/accountsreceivablepayments/README.md#get) - Get payment
* [List](docs/sdks/accountsreceivablepayments/README.md#list) - List payments

### [AccountsReceivable.Reports](docs/sdks/accountsreceivablereports/README.md)

* [GetAgedCreditors](docs/sdks/accountsreceivablereports/README.md#getagedcreditors) - Aged creditors report
* [GetAgedDebtors](docs/sdks/accountsreceivablereports/README.md#getageddebtors) - Aged debtors report
* [IsAgedCreditorsAvailable](docs/sdks/accountsreceivablereports/README.md#isagedcreditorsavailable) - Aged creditors report available
* [IsAgedDebtorsAvailable](docs/sdks/accountsreceivablereports/README.md#isageddebtorsavailable) - Aged debtors report available


### [Banking.AccountBalances](docs/sdks/bankingaccountbalances/README.md)

* [List](docs/sdks/bankingaccountbalances/README.md#list) - List account balances

### [Banking.Accounts](docs/sdks/bankingaccounts/README.md)

* [Get](docs/sdks/bankingaccounts/README.md#get) - Get account
* [List](docs/sdks/bankingaccounts/README.md#list) - List accounts

### [Banking.CategorizedStatement](docs/sdks/bankingcategorizedstatement/README.md)

* [Get](docs/sdks/bankingcategorizedstatement/README.md#get) - Get categorized bank statement

### [Banking.TransactionCategories](docs/sdks/bankingtransactioncategories/README.md)

* [Get](docs/sdks/bankingtransactioncategories/README.md#get) - Get transaction category
* [List](docs/sdks/bankingtransactioncategories/README.md#list) - List transaction categories

### [Banking.Transactions](docs/sdks/bankingtransactions/README.md)

* [Get](docs/sdks/bankingtransactions/README.md#get) - Get bank transaction
* [List](docs/sdks/bankingtransactions/README.md#list) - List transactions


### [FinancialStatements.Accounts](docs/sdks/financialstatementsaccounts/README.md)

* [Get](docs/sdks/financialstatementsaccounts/README.md#get) - Get account
* [List](docs/sdks/financialstatementsaccounts/README.md#list) - List accounts

### [FinancialStatements.BalanceSheet](docs/sdks/financialstatementsbalancesheet/README.md)

* [Get](docs/sdks/financialstatementsbalancesheet/README.md#get) - Get balance sheet
* [GetCategorizedAccounts](docs/sdks/financialstatementsbalancesheet/README.md#getcategorizedaccounts) - Get categorized balance sheet statement

### [FinancialStatements.CashFlow](docs/sdks/financialstatementscashflow/README.md)

* [Get](docs/sdks/financialstatementscashflow/README.md#get) - Get cash flow statement

### [FinancialStatements.ProfitAndLoss](docs/sdks/financialstatementsprofitandloss/README.md)

* [Get](docs/sdks/financialstatementsprofitandloss/README.md#get) - Get profit and loss
* [GetCategorizedAccounts](docs/sdks/financialstatementsprofitandloss/README.md#getcategorizedaccounts) - Get categorized profit and loss statement


### [LoanWriteback.Accounts](docs/sdks/loanwritebackaccounts/README.md)

* [Create](docs/sdks/loanwritebackaccounts/README.md#create) - Create account
* [GetCreateModel](docs/sdks/loanwritebackaccounts/README.md#getcreatemodel) - Get create account model

### [LoanWriteback.BankAccounts](docs/sdks/loanwritebackbankaccounts/README.md)

* [Create](docs/sdks/loanwritebackbankaccounts/README.md#create) - Create bank account
* [GetCreateUpdateModel](docs/sdks/loanwritebackbankaccounts/README.md#getcreateupdatemodel) - Get create/update bank account model

### [LoanWriteback.BankTransactions](docs/sdks/loanwritebackbanktransactions/README.md)

* [Create](docs/sdks/loanwritebackbanktransactions/README.md#create) - Create bank account transactions
* [GetCreateModel](docs/sdks/loanwritebackbanktransactions/README.md#getcreatemodel) - Get create bank account transactions model

### [LoanWriteback.CreateOperations](docs/sdks/loanwritebackcreateoperations/README.md)

* [Get](docs/sdks/loanwritebackcreateoperations/README.md#get) - Get create operation
* [List](docs/sdks/loanwritebackcreateoperations/README.md#list) - List create operations

### [LoanWriteback.DirectCosts](docs/sdks/loanwritebackdirectcosts/README.md)

* [Create](docs/sdks/loanwritebackdirectcosts/README.md#create) - Create direct cost
* [GetCreateModel](docs/sdks/loanwritebackdirectcosts/README.md#getcreatemodel) - Get create direct cost model

### [LoanWriteback.Suppliers](docs/sdks/loanwritebacksuppliers/README.md)

* [Create](docs/sdks/loanwritebacksuppliers/README.md#create) - Create supplier
* [GetCreateUpdateModel](docs/sdks/loanwritebacksuppliers/README.md#getcreateupdatemodel) - Get create/update supplier model

### [LoanWriteback.Transfers](docs/sdks/loanwritebacktransfers/README.md)

* [Create](docs/sdks/loanwritebacktransfers/README.md#create) - Create transfer
* [GetCreateModel](docs/sdks/loanwritebacktransfers/README.md#getcreatemodel) - Get create transfer model

### [ManageData](docs/sdks/managedata/README.md)

* [GetStatus](docs/sdks/managedata/README.md#getstatus) - Get data status

### [ManageData.PullOperations](docs/sdks/managedatapulloperations/README.md)

* [Get](docs/sdks/managedatapulloperations/README.md#get) - Get pull operation
* [List](docs/sdks/managedatapulloperations/README.md#list) - List pull operations

### [ManageData.Refresh](docs/sdks/managedatarefresh/README.md)

* [AllDataTypes](docs/sdks/managedatarefresh/README.md#alldatatypes) - Refresh all data
* [DataType](docs/sdks/managedatarefresh/README.md#datatype) - Refresh data type


### [Sales.Customers](docs/sdks/salescustomers/README.md)

* [Get](docs/sdks/salescustomers/README.md#get) - Get customer
* [List](docs/sdks/salescustomers/README.md#list) - List customers

### [Sales.Disputes](docs/sdks/salesdisputes/README.md)

* [Get](docs/sdks/salesdisputes/README.md#get) - Get dispute
* [List](docs/sdks/salesdisputes/README.md#list) - List disputes

### [Sales.Locations](docs/sdks/saleslocations/README.md)

* [Get](docs/sdks/saleslocations/README.md#get) - Get location
* [List](docs/sdks/saleslocations/README.md#list) - List locations

### [Sales.Metrics](docs/sdks/salesmetrics/README.md)

* [GetCustomerRetention](docs/sdks/salesmetrics/README.md#getcustomerretention) - Get customer retention metrics
* [GetLifetimeValue](docs/sdks/salesmetrics/README.md#getlifetimevalue) - Get lifetime value metrics
* [GetRevenue](docs/sdks/salesmetrics/README.md#getrevenue) - Get commerce revenue metrics

### [Sales.Orders](docs/sdks/salesorders/README.md)

* [Get](docs/sdks/salesorders/README.md#get) - Get order
* [List](docs/sdks/salesorders/README.md#list) - List orders

### [Sales.PaymentMethods](docs/sdks/salespaymentmethods/README.md)

* [Get](docs/sdks/salespaymentmethods/README.md#get) - Get payment method
* [List](docs/sdks/salespaymentmethods/README.md#list) - List payment methods

### [Sales.Payments](docs/sdks/salespayments/README.md)

* [Get](docs/sdks/salespayments/README.md#get) - Get payment
* [List](docs/sdks/salespayments/README.md#list) - List payments

### [Sales.ProductCategories](docs/sdks/salesproductcategories/README.md)

* [Get](docs/sdks/salesproductcategories/README.md#get) - Get product category
* [List](docs/sdks/salesproductcategories/README.md#list) - List product categories

### [Sales.Products](docs/sdks/salesproducts/README.md)

* [Get](docs/sdks/salesproducts/README.md#get) - Get product
* [List](docs/sdks/salesproducts/README.md#list) - List products

### [Sales.Reports](docs/sdks/salesreports/README.md)

* [GetOrders](docs/sdks/salesreports/README.md#getorders) - Get orders report
* [GetRefunds](docs/sdks/salesreports/README.md#getrefunds) - Get refunds report

### [Sales.Transactions](docs/sdks/salestransactions/README.md)

* [Get](docs/sdks/salestransactions/README.md#get) - Get transaction
* [List](docs/sdks/salestransactions/README.md#list) - List transactions


### [Transactions.AccountTransactions](docs/sdks/transactionsaccounttransactions/README.md)

* [Get](docs/sdks/transactionsaccounttransactions/README.md#get) - Get account transaction
* [List](docs/sdks/transactionsaccounttransactions/README.md#list) - List account transactions

### [Transactions.DirectCosts](docs/sdks/transactionsdirectcosts/README.md)

* [DownloadAttachment](docs/sdks/transactionsdirectcosts/README.md#downloadattachment) - Download direct cost attachment
* [Get](docs/sdks/transactionsdirectcosts/README.md#get) - Get direct cost
* [GetAttachment](docs/sdks/transactionsdirectcosts/README.md#getattachment) - Get direct cost attachment
* [List](docs/sdks/transactionsdirectcosts/README.md#list) - List direct costs
* [ListAttachments](docs/sdks/transactionsdirectcosts/README.md#listattachments) - List direct cost attachments

### [Transactions.JournalEntries](docs/sdks/transactionsjournalentries/README.md)

* [Get](docs/sdks/transactionsjournalentries/README.md#get) - Get journal entry
* [List](docs/sdks/transactionsjournalentries/README.md#list) - List journal entries

### [Transactions.Journals](docs/sdks/transactionsjournals/README.md)

* [Get](docs/sdks/transactionsjournals/README.md#get) - Get journal
* [List](docs/sdks/transactionsjournals/README.md#list) - List journals

### [Transactions.Transfers](docs/sdks/transactionstransfers/README.md)

* [Get](docs/sdks/transactionstransfers/README.md#get) - Get transfer
* [List](docs/sdks/transactionstransfers/README.md#list) - List transfers
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->
# Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

## Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->


### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
