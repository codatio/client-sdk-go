# Lending
    
Lending helps you make smarter credit decisions on small businesses by enabling you to pull your customers' latest data from the operating systems they are already using. You can use that data for automating decisioning and surfacing new insights on the customer, all via one API.

## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/lending
```## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/lending
```## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/lending
```<!-- Start SDK Installation -->

<!-- End SDK Installation -->

## Example Usage


```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/lending"
	"github.com/codatio/client-sdk-go/lending/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/pkg/models/operations"
)

func main() {
    s := codatlending.New(
        codatlending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountingBankData.GetAccount(ctx, operations.GetAccountingBankAccountRequest{
        AccountID: "corrupti",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountingBankAccount != nil {
        // handle response
    }
}
```

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/lending"
	"github.com/codatio/client-sdk-go/lending/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/pkg/models/operations"
)

func main() {
    s := codatlending.New(
        codatlending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.GetAccountingProfile(ctx, operations.GetAccountingProfileRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountingCompanyInfo != nil {
        // handle response
    }
}
```

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/lending"
	"github.com/codatio/client-sdk-go/lending/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/pkg/models/operations"
)

func main() {
    s := codatlending.New(
        codatlending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountingBankData.ListTransactions(ctx, operations.ListAccountingBankAccountTransactionsRequest{
        AccountID: "corrupti",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatlending.String("-modifiedDate"),
        Page: codatlending.Int(1),
        PageSize: codatlending.Int(100),
        Query: codatlending.String("provident"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountingBankTransactions != nil {
        // handle response
    }
}
```<!-- Start SDK Example Usage -->

<!-- End SDK Example Usage -->

## Available Resources and Operations


### [AccountingBankData](docs/sdks/accountingbankdata/README.md)

* [GetAccount](docs/sdks/accountingbankdata/README.md#getaccount) - Get bank account
* [ListAccounts](docs/sdks/accountingbankdata/README.md#listaccounts) - List bank accounts
* [ListTransactions](docs/sdks/accountingbankdata/README.md#listtransactions) - List bank account transactions

### [AccountsPayable](docs/sdks/accountspayable/README.md)

* [DownloadBillAttachment](docs/sdks/accountspayable/README.md#downloadbillattachment) - Download bill attachment
* [DownloadSupplierAttachment](docs/sdks/accountspayable/README.md#downloadsupplierattachment) - Download supplier attachment
* [GetAgedCreditorsReport](docs/sdks/accountspayable/README.md#getagedcreditorsreport) - Aged creditors report
* [GetBill](docs/sdks/accountspayable/README.md#getbill) - Get bill
* [GetBillAttachment](docs/sdks/accountspayable/README.md#getbillattachment) - Get bill attachment
* [GetBillCreditNote](docs/sdks/accountspayable/README.md#getbillcreditnote) - Get bill credit note
* [GetBillPayment](docs/sdks/accountspayable/README.md#getbillpayment) - Get bill payment
* [GetSupplier](docs/sdks/accountspayable/README.md#getsupplier) - Get supplier
* [GetSupplierAttachment](docs/sdks/accountspayable/README.md#getsupplierattachment) - Get supplier attachment
* [IsAgedCreditorsReportAvailable](docs/sdks/accountspayable/README.md#isagedcreditorsreportavailable) - Aged creditors report available
* [ListBillAttachments](docs/sdks/accountspayable/README.md#listbillattachments) - List bill attachments
* [ListBillCreditNotes](docs/sdks/accountspayable/README.md#listbillcreditnotes) - List bill credit notes
* [ListBillPayments](docs/sdks/accountspayable/README.md#listbillpayments) - List bill payments
* [ListBills](docs/sdks/accountspayable/README.md#listbills) - List bills
* [ListSupplierAttachments](docs/sdks/accountspayable/README.md#listsupplierattachments) - List supplier attachments
* [ListSuppliers](docs/sdks/accountspayable/README.md#listsuppliers) - List suppliers

### [AccountsReceivable](docs/sdks/accountsreceivable/README.md)

* [DownloadCustomerAttachment](docs/sdks/accountsreceivable/README.md#downloadcustomerattachment) - Download customer attachment
* [DownloadDirectIncomeAttachment](docs/sdks/accountsreceivable/README.md#downloaddirectincomeattachment) - Download direct income attachment
* [DownloadInvoiceAttachment](docs/sdks/accountsreceivable/README.md#downloadinvoiceattachment) - Download invoice attachment
* [DownloadInvoicePdf](docs/sdks/accountsreceivable/README.md#downloadinvoicepdf) - Get invoice as PDF
* [GetAgedDebtorsReport](docs/sdks/accountsreceivable/README.md#getageddebtorsreport) - Aged debtors report
* [GetCreditNote](docs/sdks/accountsreceivable/README.md#getcreditnote) - Get credit note
* [GetCustomer](docs/sdks/accountsreceivable/README.md#getcustomer) - Get customer
* [GetCustomerAttachment](docs/sdks/accountsreceivable/README.md#getcustomerattachment) - Get customer attachment
* [GetDirectIncome](docs/sdks/accountsreceivable/README.md#getdirectincome) - Get direct income
* [GetDirectIncomeAttachment](docs/sdks/accountsreceivable/README.md#getdirectincomeattachment) - Get direct income attachment
* [GetInvoice](docs/sdks/accountsreceivable/README.md#getinvoice) - Get invoice
* [GetInvoiceAttachment](docs/sdks/accountsreceivable/README.md#getinvoiceattachment) - Get invoice attachment
* [GetPayment](docs/sdks/accountsreceivable/README.md#getpayment) - Get payment
* [GetReconciledInvoices](docs/sdks/accountsreceivable/README.md#getreconciledinvoices) - Get reconciled invoices
* [IsAgedDebtorReportAvailable](docs/sdks/accountsreceivable/README.md#isageddebtorreportavailable) - Aged debtors report available
* [ListCreditNotes](docs/sdks/accountsreceivable/README.md#listcreditnotes) - List credit notes
* [ListCustomerAttachments](docs/sdks/accountsreceivable/README.md#listcustomerattachments) - List customer attachments
* [ListCustomers](docs/sdks/accountsreceivable/README.md#listcustomers) - List customers
* [ListDirectIncomeAttachments](docs/sdks/accountsreceivable/README.md#listdirectincomeattachments) - List direct income attachments
* [ListDirectIncomes](docs/sdks/accountsreceivable/README.md#listdirectincomes) - List direct incomes
* [ListInvoiceAttachments](docs/sdks/accountsreceivable/README.md#listinvoiceattachments) - List invoice attachments
* [ListInvoices](docs/sdks/accountsreceivable/README.md#listinvoices) - List invoices
* [ListPayments](docs/sdks/accountsreceivable/README.md#listpayments) - List payments

### [Banking](docs/sdks/banking/README.md)

* [GetBankAccount](docs/sdks/banking/README.md#getbankaccount) - Get account
* [GetBankTransaction](docs/sdks/banking/README.md#getbanktransaction) - Get bank transaction
* [GetBankTransactionCategory](docs/sdks/banking/README.md#getbanktransactioncategory) - Get transaction category
* [GetCategorizedBankStatement](docs/sdks/banking/README.md#getcategorizedbankstatement) - Get categorized bank statement
* [ListBankAccountBalances](docs/sdks/banking/README.md#listbankaccountbalances) - List account balances
* [ListBankAccounts](docs/sdks/banking/README.md#listbankaccounts) - List accounts
* [ListBankTransactionCategories](docs/sdks/banking/README.md#listbanktransactioncategories) - List transaction categories
* [ListBankTransactions](docs/sdks/banking/README.md#listbanktransactions) - List transactions

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

* [GetAccount](docs/sdks/financialstatements/README.md#getaccount) - Get account
* [GetBalanceSheet](docs/sdks/financialstatements/README.md#getbalancesheet) - Get balance sheet
* [GetCashFlowStatement](docs/sdks/financialstatements/README.md#getcashflowstatement) - Get cash flow statement
* [GetCategorizedBalanceSheet](docs/sdks/financialstatements/README.md#getcategorizedbalancesheet) - Get categorized balance sheet statement
* [GetCategorizedProfitAndLoss](docs/sdks/financialstatements/README.md#getcategorizedprofitandloss) - Get categorized profit and loss statement
* [GetProfitAndLoss](docs/sdks/financialstatements/README.md#getprofitandloss) - Get profit and loss
* [ListAccounts](docs/sdks/financialstatements/README.md#listaccounts) - List accounts

### [Liabilities](docs/sdks/liabilities/README.md)

* [GetLoanSummary](docs/sdks/liabilities/README.md#getloansummary) - Get loan summaries
* [ListLoanTransactions](docs/sdks/liabilities/README.md#listloantransactions) - List loan transactions

### [ManageData](docs/sdks/managedata/README.md)

* [Get](docs/sdks/managedata/README.md#get) - Get data status
* [GetPullOperation](docs/sdks/managedata/README.md#getpulloperation) - Get pull operation
* [ListPullOperations](docs/sdks/managedata/README.md#listpulloperations) - List pull operations
* [RefreshAllDataTypes](docs/sdks/managedata/README.md#refreshalldatatypes) - Refresh all data
* [RefreshDataType](docs/sdks/managedata/README.md#refreshdatatype) - Refresh data type

### [Sales](docs/sdks/sales/README.md)

* [GetCustomer](docs/sdks/sales/README.md#getcustomer) - Get customer
* [GetCustomerRetentionMetrics](docs/sdks/sales/README.md#getcustomerretentionmetrics) - Get customer retention metrics
* [GetDispute](docs/sdks/sales/README.md#getdispute) - Get dispute
* [GetLifetimeValueMetrics](docs/sdks/sales/README.md#getlifetimevaluemetrics) - Get lifetime value metrics
* [GetLocation](docs/sdks/sales/README.md#getlocation) - Get location
* [GetOrder](docs/sdks/sales/README.md#getorder) - Get order
* [GetOrdersReport](docs/sdks/sales/README.md#getordersreport) - Get orders report
* [GetPayment](docs/sdks/sales/README.md#getpayment) - Get payment
* [GetPaymentMethod](docs/sdks/sales/README.md#getpaymentmethod) - Get payment method
* [GetProduct](docs/sdks/sales/README.md#getproduct) - Get product
* [GetProductCategory](docs/sdks/sales/README.md#getproductcategory) - Get product category
* [GetRefundsReport](docs/sdks/sales/README.md#getrefundsreport) - Get refunds report
* [GetRevenueMetrics](docs/sdks/sales/README.md#getrevenuemetrics) - Get commerce revenue metrics
* [GetTransaction](docs/sdks/sales/README.md#gettransaction) - Get transaction
* [ListCustomers](docs/sdks/sales/README.md#listcustomers) - List customers
* [ListDisputes](docs/sdks/sales/README.md#listdisputes) - List disputes
* [ListLocations](docs/sdks/sales/README.md#listlocations) - List locations
* [ListOrders](docs/sdks/sales/README.md#listorders) - List orders
* [ListPaymentMethods](docs/sdks/sales/README.md#listpaymentmethods) - List payment methods
* [ListPayments](docs/sdks/sales/README.md#listpayments) - List payments
* [ListProductCategories](docs/sdks/sales/README.md#listproductcategories) - List product categories
* [ListProducts](docs/sdks/sales/README.md#listproducts) - List products
* [ListTransactions](docs/sdks/sales/README.md#listtransactions) - List transactions

### [Transactions](docs/sdks/transactions/README.md)

* [DownloadDirectCostAttachment](docs/sdks/transactions/README.md#downloaddirectcostattachment) - Download direct cost attachment
* [GetAccountTransaction](docs/sdks/transactions/README.md#getaccounttransaction) - Get account transaction
* [GetDirectCost](docs/sdks/transactions/README.md#getdirectcost) - Get direct cost
* [GetDirectCostAttachment](docs/sdks/transactions/README.md#getdirectcostattachment) - Get direct cost attachment
* [GetJournal](docs/sdks/transactions/README.md#getjournal) - Get journal
* [GetJournalEntry](docs/sdks/transactions/README.md#getjournalentry) - Get journal entry
* [GetTransfer](docs/sdks/transactions/README.md#gettransfer) - Get transfer
* [ListAccountTransactions](docs/sdks/transactions/README.md#listaccounttransactions) - List account transactions
* [ListDirectCostAttachments](docs/sdks/transactions/README.md#listdirectcostattachments) - List direct cost attachments
* [ListDirectCosts](docs/sdks/transactions/README.md#listdirectcosts) - List direct costs
* [ListJournalEntries](docs/sdks/transactions/README.md#listjournalentries) - List journal entries
* [ListJournals](docs/sdks/transactions/README.md#listjournals) - List journals
* [ListTransfers](docs/sdks/transactions/README.md#listtransfers) - List transfers## Available Resources and Operations

### [CodatLending SDK](docs/sdks/codatlending/README.md)

* [GetAccountingProfile](docs/sdks/codatlending/README.md#getaccountingprofile) - Get company accounting profile

### [AccountingBankData](docs/sdks/accountingbankdata/README.md)

* [GetAccount](docs/sdks/accountingbankdata/README.md#getaccount) - Get bank account
* [ListAccounts](docs/sdks/accountingbankdata/README.md#listaccounts) - List bank accounts
* [ListTransactions](docs/sdks/accountingbankdata/README.md#listtransactions) - List bank account transactions

### [AccountsPayable](docs/sdks/accountspayable/README.md)

* [DownloadBillAttachment](docs/sdks/accountspayable/README.md#downloadbillattachment) - Download bill attachment
* [DownloadSupplierAttachment](docs/sdks/accountspayable/README.md#downloadsupplierattachment) - Download supplier attachment
* [GetAgedCreditorsReport](docs/sdks/accountspayable/README.md#getagedcreditorsreport) - Aged creditors report
* [GetBill](docs/sdks/accountspayable/README.md#getbill) - Get bill
* [GetBillAttachment](docs/sdks/accountspayable/README.md#getbillattachment) - Get bill attachment
* [GetBillCreditNote](docs/sdks/accountspayable/README.md#getbillcreditnote) - Get bill credit note
* [GetBillPayment](docs/sdks/accountspayable/README.md#getbillpayment) - Get bill payment
* [GetSupplier](docs/sdks/accountspayable/README.md#getsupplier) - Get supplier
* [GetSupplierAttachment](docs/sdks/accountspayable/README.md#getsupplierattachment) - Get supplier attachment
* [IsAgedCreditorsReportAvailable](docs/sdks/accountspayable/README.md#isagedcreditorsreportavailable) - Aged creditors report available
* [ListBillAttachments](docs/sdks/accountspayable/README.md#listbillattachments) - List bill attachments
* [ListBillCreditNotes](docs/sdks/accountspayable/README.md#listbillcreditnotes) - List bill credit notes
* [ListBillPayments](docs/sdks/accountspayable/README.md#listbillpayments) - List bill payments
* [ListBills](docs/sdks/accountspayable/README.md#listbills) - List bills
* [ListSupplierAttachments](docs/sdks/accountspayable/README.md#listsupplierattachments) - List supplier attachments
* [ListSuppliers](docs/sdks/accountspayable/README.md#listsuppliers) - List suppliers

### [AccountsReceivable](docs/sdks/accountsreceivable/README.md)

* [DownloadCustomerAttachment](docs/sdks/accountsreceivable/README.md#downloadcustomerattachment) - Download customer attachment
* [DownloadDirectIncomeAttachment](docs/sdks/accountsreceivable/README.md#downloaddirectincomeattachment) - Download direct income attachment
* [DownloadInvoiceAttachment](docs/sdks/accountsreceivable/README.md#downloadinvoiceattachment) - Download invoice attachment
* [DownloadInvoicePdf](docs/sdks/accountsreceivable/README.md#downloadinvoicepdf) - Get invoice as PDF
* [GetAgedDebtorsReport](docs/sdks/accountsreceivable/README.md#getageddebtorsreport) - Aged debtors report
* [GetCreditNote](docs/sdks/accountsreceivable/README.md#getcreditnote) - Get credit note
* [GetCustomer](docs/sdks/accountsreceivable/README.md#getcustomer) - Get customer
* [GetCustomerAttachment](docs/sdks/accountsreceivable/README.md#getcustomerattachment) - Get customer attachment
* [GetDirectIncome](docs/sdks/accountsreceivable/README.md#getdirectincome) - Get direct income
* [GetDirectIncomeAttachment](docs/sdks/accountsreceivable/README.md#getdirectincomeattachment) - Get direct income attachment
* [GetInvoice](docs/sdks/accountsreceivable/README.md#getinvoice) - Get invoice
* [GetInvoiceAttachment](docs/sdks/accountsreceivable/README.md#getinvoiceattachment) - Get invoice attachment
* [GetPayment](docs/sdks/accountsreceivable/README.md#getpayment) - Get payment
* [GetReconciledInvoices](docs/sdks/accountsreceivable/README.md#getreconciledinvoices) - Get reconciled invoices
* [IsAgedDebtorReportAvailable](docs/sdks/accountsreceivable/README.md#isageddebtorreportavailable) - Aged debtors report available
* [ListCreditNotes](docs/sdks/accountsreceivable/README.md#listcreditnotes) - List credit notes
* [ListCustomerAttachments](docs/sdks/accountsreceivable/README.md#listcustomerattachments) - List customer attachments
* [ListCustomers](docs/sdks/accountsreceivable/README.md#listcustomers) - List customers
* [ListDirectIncomeAttachments](docs/sdks/accountsreceivable/README.md#listdirectincomeattachments) - List direct income attachments
* [ListDirectIncomes](docs/sdks/accountsreceivable/README.md#listdirectincomes) - List direct incomes
* [ListInvoiceAttachments](docs/sdks/accountsreceivable/README.md#listinvoiceattachments) - List invoice attachments
* [ListInvoices](docs/sdks/accountsreceivable/README.md#listinvoices) - List invoices
* [ListPayments](docs/sdks/accountsreceivable/README.md#listpayments) - List payments

### [BankStatements](docs/sdks/bankstatements/README.md)

* [GetBankAccount](docs/sdks/bankstatements/README.md#getbankaccount) - Get account
* [GetBankTransaction](docs/sdks/bankstatements/README.md#getbanktransaction) - Get bank transaction
* [GetBankTransactionCategory](docs/sdks/bankstatements/README.md#getbanktransactioncategory) - Get transaction category
* [GetCategorizedBankStatement](docs/sdks/bankstatements/README.md#getcategorizedbankstatement) - Get categorized bank statement
* [ListBankAccountBalances](docs/sdks/bankstatements/README.md#listbankaccountbalances) - List account balances
* [ListBankAccounts](docs/sdks/bankstatements/README.md#listbankaccounts) - List accounts
* [ListBankTransactionCategories](docs/sdks/bankstatements/README.md#listbanktransactioncategories) - List transaction categories
* [ListBankTransactions](docs/sdks/bankstatements/README.md#listbanktransactions) - List transactions

### [Companies](docs/sdks/companies/README.md)

* [Create](docs/sdks/companies/README.md#create) - Create company
* [Delete](docs/sdks/companies/README.md#delete) - Delete a company
* [Get](docs/sdks/companies/README.md#get) - Get company
* [List](docs/sdks/companies/README.md#list) - List companies
* [Update](docs/sdks/companies/README.md#update) - Update company

### [CompanyInfo](docs/sdks/companyinfo/README.md)

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

* [GetAccount](docs/sdks/financialstatements/README.md#getaccount) - Get account
* [GetBalanceSheet](docs/sdks/financialstatements/README.md#getbalancesheet) - Get balance sheet
* [GetCashFlowStatement](docs/sdks/financialstatements/README.md#getcashflowstatement) - Get cash flow statement
* [GetCategorizedBalanceSheet](docs/sdks/financialstatements/README.md#getcategorizedbalancesheet) - Get categorized balance sheet statement
* [GetCategorizedProfitAndLoss](docs/sdks/financialstatements/README.md#getcategorizedprofitandloss) - Get categorized profit and loss statement
* [GetProfitAndLoss](docs/sdks/financialstatements/README.md#getprofitandloss) - Get profit and loss
* [ListAccounts](docs/sdks/financialstatements/README.md#listaccounts) - List accounts

### [Liabilities](docs/sdks/liabilities/README.md)

* [GetLoanSummary](docs/sdks/liabilities/README.md#getloansummary) - Get loan summaries
* [ListLoanTransactions](docs/sdks/liabilities/README.md#listloantransactions) - List loan transactions

### [ManageData](docs/sdks/managedata/README.md)

* [Get](docs/sdks/managedata/README.md#get) - Get data status
* [GetPullOperation](docs/sdks/managedata/README.md#getpulloperation) - Get pull operation
* [ListPullOperations](docs/sdks/managedata/README.md#listpulloperations) - List pull operations
* [RefreshAllDataTypes](docs/sdks/managedata/README.md#refreshalldatatypes) - Refresh all data
* [RefreshDataType](docs/sdks/managedata/README.md#refreshdatatype) - Refresh data type

### [Sales](docs/sdks/sales/README.md)

* [GetCustomer](docs/sdks/sales/README.md#getcustomer) - Get customer
* [GetCustomerRetentionMetrics](docs/sdks/sales/README.md#getcustomerretentionmetrics) - Get customer retention metrics
* [GetDispute](docs/sdks/sales/README.md#getdispute) - Get dispute
* [GetLifetimeValueMetrics](docs/sdks/sales/README.md#getlifetimevaluemetrics) - Get lifetime value metrics
* [GetLocation](docs/sdks/sales/README.md#getlocation) - Get location
* [GetOrder](docs/sdks/sales/README.md#getorder) - Get order
* [GetOrdersReport](docs/sdks/sales/README.md#getordersreport) - Get orders report
* [GetPayment](docs/sdks/sales/README.md#getpayment) - Get payment
* [GetPaymentMethod](docs/sdks/sales/README.md#getpaymentmethod) - Get payment method
* [GetProduct](docs/sdks/sales/README.md#getproduct) - Get product
* [GetProductCategory](docs/sdks/sales/README.md#getproductcategory) - Get product category
* [GetRefundsReport](docs/sdks/sales/README.md#getrefundsreport) - Get refunds report
* [GetRevenueMetrics](docs/sdks/sales/README.md#getrevenuemetrics) - Get commerce revenue metrics
* [GetTransaction](docs/sdks/sales/README.md#gettransaction) - Get transaction
* [ListCustomers](docs/sdks/sales/README.md#listcustomers) - List customers
* [ListDisputes](docs/sdks/sales/README.md#listdisputes) - List disputes
* [ListLocations](docs/sdks/sales/README.md#listlocations) - List locations
* [ListOrders](docs/sdks/sales/README.md#listorders) - List orders
* [ListPaymentMethods](docs/sdks/sales/README.md#listpaymentmethods) - List payment methods
* [ListPayments](docs/sdks/sales/README.md#listpayments) - List payments
* [ListProductCategories](docs/sdks/sales/README.md#listproductcategories) - List product categories
* [ListProducts](docs/sdks/sales/README.md#listproducts) - List products
* [ListTransactions](docs/sdks/sales/README.md#listtransactions) - List transactions

### [Transactions](docs/sdks/transactions/README.md)

* [DownloadDirectCostAttachment](docs/sdks/transactions/README.md#downloaddirectcostattachment) - Download direct cost attachment
* [GetAccountTransaction](docs/sdks/transactions/README.md#getaccounttransaction) - Get account transaction
* [GetDirectCost](docs/sdks/transactions/README.md#getdirectcost) - Get direct cost
* [GetDirectCostAttachment](docs/sdks/transactions/README.md#getdirectcostattachment) - Get direct cost attachment
* [GetJournal](docs/sdks/transactions/README.md#getjournal) - Get journal
* [GetJournalEntry](docs/sdks/transactions/README.md#getjournalentry) - Get journal entry
* [GetTransfer](docs/sdks/transactions/README.md#gettransfer) - Get transfer
* [ListAccountTransactions](docs/sdks/transactions/README.md#listaccounttransactions) - List account transactions
* [ListDirectCostAttachments](docs/sdks/transactions/README.md#listdirectcostattachments) - List direct cost attachments
* [ListDirectCosts](docs/sdks/transactions/README.md#listdirectcosts) - List direct costs
* [ListJournalEntries](docs/sdks/transactions/README.md#listjournalentries) - List journal entries
* [ListJournals](docs/sdks/transactions/README.md#listjournals) - List journals
* [ListTransfers](docs/sdks/transactions/README.md#listtransfers) - List transfers## Available Resources and Operations


### [AccountingBankData](docs/sdks/accountingbankdata/README.md)

* [ListTransactions](docs/sdks/accountingbankdata/README.md#listtransactions) - List bank account transactions

### [AccountsPayable](docs/sdks/accountspayable/README.md)

* [DownloadBillAttachment](docs/sdks/accountspayable/README.md#downloadbillattachment) - Download bill attachment
* [GetBillAttachment](docs/sdks/accountspayable/README.md#getbillattachment) - Get bill attachment

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

* [GetLoanSummary](docs/sdks/liabilities/README.md#getloansummary) - Get loan summaries
* [ListLoanTransactions](docs/sdks/liabilities/README.md#listloantransactions) - List loan transactions


### [AccountingBankDataAccounts](docs/sdks/accountingbankdataaccounts/README.md)

* [Get](docs/sdks/accountingbankdataaccounts/README.md#get) - Get bank account
* [List](docs/sdks/accountingbankdataaccounts/README.md#list) - List bank accounts


### [AccountsPayableBillCreditNotes](docs/sdks/accountspayablebillcreditnotes/README.md)

* [Get](docs/sdks/accountspayablebillcreditnotes/README.md#get) - Get bill credit note
* [List](docs/sdks/accountspayablebillcreditnotes/README.md#list) - List bill credit notes

### [AccountsPayableBillPayments](docs/sdks/accountspayablebillpayments/README.md)

* [Get](docs/sdks/accountspayablebillpayments/README.md#get) - Get bill payment
* [List](docs/sdks/accountspayablebillpayments/README.md#list) - List bill payments

### [AccountsPayableBills](docs/sdks/accountspayablebills/README.md)

* [Get](docs/sdks/accountspayablebills/README.md#get) - Get bill
* [List](docs/sdks/accountspayablebills/README.md#list) - List bills
* [ListAttachments](docs/sdks/accountspayablebills/README.md#listattachments) - List bill attachments

### [AccountsPayableSuppliers](docs/sdks/accountspayablesuppliers/README.md)

* [DownloadAttachment](docs/sdks/accountspayablesuppliers/README.md#downloadattachment) - Download supplier attachment
* [Get](docs/sdks/accountspayablesuppliers/README.md#get) - Get supplier
* [GetAttachment](docs/sdks/accountspayablesuppliers/README.md#getattachment) - Get supplier attachment
* [List](docs/sdks/accountspayablesuppliers/README.md#list) - List suppliers
* [ListAttachments](docs/sdks/accountspayablesuppliers/README.md#listattachments) - List supplier attachments


### [AccountsReceivableCreditNotes](docs/sdks/accountsreceivablecreditnotes/README.md)

* [Get](docs/sdks/accountsreceivablecreditnotes/README.md#get) - Get credit note
* [List](docs/sdks/accountsreceivablecreditnotes/README.md#list) - List credit notes

### [AccountsReceivableCustomers](docs/sdks/accountsreceivablecustomers/README.md)

* [DownloadAttachment](docs/sdks/accountsreceivablecustomers/README.md#downloadattachment) - Download customer attachment
* [Get](docs/sdks/accountsreceivablecustomers/README.md#get) - Get customer
* [GetAttachment](docs/sdks/accountsreceivablecustomers/README.md#getattachment) - Get customer attachment
* [List](docs/sdks/accountsreceivablecustomers/README.md#list) - List customers
* [ListAttachments](docs/sdks/accountsreceivablecustomers/README.md#listattachments) - List customer attachments

### [AccountsReceivableDirectIncomes](docs/sdks/accountsreceivabledirectincomes/README.md)

* [DownloadAttachment](docs/sdks/accountsreceivabledirectincomes/README.md#downloadattachment) - Download direct income attachment
* [Get](docs/sdks/accountsreceivabledirectincomes/README.md#get) - Get direct income
* [GetAttachment](docs/sdks/accountsreceivabledirectincomes/README.md#getattachment) - Get direct income attachment
* [List](docs/sdks/accountsreceivabledirectincomes/README.md#list) - List direct incomes
* [ListAttachments](docs/sdks/accountsreceivabledirectincomes/README.md#listattachments) - List direct income attachments

### [AccountsReceivableInvoices](docs/sdks/accountsreceivableinvoices/README.md)

* [DownloadAttachment](docs/sdks/accountsreceivableinvoices/README.md#downloadattachment) - Download invoice attachment
* [DownloadPdf](docs/sdks/accountsreceivableinvoices/README.md#downloadpdf) - Get invoice as PDF
* [Get](docs/sdks/accountsreceivableinvoices/README.md#get) - Get invoice
* [GetAttachment](docs/sdks/accountsreceivableinvoices/README.md#getattachment) - Get invoice attachment
* [List](docs/sdks/accountsreceivableinvoices/README.md#list) - List invoices
* [ListAttachments](docs/sdks/accountsreceivableinvoices/README.md#listattachments) - List invoice attachments
* [ListReconciled](docs/sdks/accountsreceivableinvoices/README.md#listreconciled) - List reconciled invoices

### [AccountsReceivablePayments](docs/sdks/accountsreceivablepayments/README.md)

* [Get](docs/sdks/accountsreceivablepayments/README.md#get) - Get payment
* [List](docs/sdks/accountsreceivablepayments/README.md#list) - List payments

### [AccountsReceivableReports](docs/sdks/accountsreceivablereports/README.md)

* [GetAgedCreditors](docs/sdks/accountsreceivablereports/README.md#getagedcreditors) - Aged creditors report
* [GetAgedDebtors](docs/sdks/accountsreceivablereports/README.md#getageddebtors) - Aged debtors report
* [IsAgedCreditorsAvailable](docs/sdks/accountsreceivablereports/README.md#isagedcreditorsavailable) - Aged creditors report available
* [IsAgedDebtorsAvailable](docs/sdks/accountsreceivablereports/README.md#isageddebtorsavailable) - Aged debtors report available


### [BankingAccountBalances](docs/sdks/bankingaccountbalances/README.md)

* [List](docs/sdks/bankingaccountbalances/README.md#list) - List account balances

### [Banking.Accounts](docs/sdks/bankingaccounts/README.md)

* [Get](docs/sdks/bankingaccounts/README.md#get) - Get account
* [List](docs/sdks/bankingaccounts/README.md#list) - List accounts

### [BankingCategorizedStatement](docs/sdks/bankingcategorizedstatement/README.md)

* [Get](docs/sdks/bankingcategorizedstatement/README.md#get) - Get categorized bank statement

### [BankingTransactionCategories](docs/sdks/bankingtransactioncategories/README.md)

* [Get](docs/sdks/bankingtransactioncategories/README.md#get) - Get transaction category
* [List](docs/sdks/bankingtransactioncategories/README.md#list) - List transaction categories

### [Banking.Transactions](docs/sdks/bankingtransactions/README.md)

* [Get](docs/sdks/bankingtransactions/README.md#get) - Get bank transaction
* [List](docs/sdks/bankingtransactions/README.md#list) - List transactions


### [FinancialStatementsAccounts](docs/sdks/financialstatementsaccounts/README.md)

* [Get](docs/sdks/financialstatementsaccounts/README.md#get) - Get account
* [List](docs/sdks/financialstatementsaccounts/README.md#list) - List accounts

### [FinancialStatementsBalanceSheet](docs/sdks/financialstatementsbalancesheet/README.md)

* [Get](docs/sdks/financialstatementsbalancesheet/README.md#get) - Get balance sheet
* [GetCategorizedAccounts](docs/sdks/financialstatementsbalancesheet/README.md#getcategorizedaccounts) - Get categorized balance sheet statement

### [FinancialStatementsCashFlow](docs/sdks/financialstatementscashflow/README.md)

* [Get](docs/sdks/financialstatementscashflow/README.md#get) - Get cash flow statement

### [FinancialStatementsProfitAndLoss](docs/sdks/financialstatementsprofitandloss/README.md)

* [Get](docs/sdks/financialstatementsprofitandloss/README.md#get) - Get profit and loss
* [GetCategorizedAccounts](docs/sdks/financialstatementsprofitandloss/README.md#getcategorizedaccounts) - Get categorized profit and loss statement

### [ManageData](docs/sdks/managedata/README.md)

* [GetStatus](docs/sdks/managedata/README.md#getstatus) - Get data status

### [ManageDataPullOperations](docs/sdks/managedatapulloperations/README.md)

* [Get](docs/sdks/managedatapulloperations/README.md#get) - Get pull operation
* [List](docs/sdks/managedatapulloperations/README.md#list) - List pull operations

### [ManageDataRefresh](docs/sdks/managedatarefresh/README.md)

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

### [SalesPaymentMethods](docs/sdks/salespaymentmethods/README.md)

* [Get](docs/sdks/salespaymentmethods/README.md#get) - Get payment method
* [List](docs/sdks/salespaymentmethods/README.md#list) - List payment methods

### [Sales.Payments](docs/sdks/salespayments/README.md)

* [Get](docs/sdks/salespayments/README.md#get) - Get payment
* [List](docs/sdks/salespayments/README.md#list) - List payments

### [SalesProductCategories](docs/sdks/salesproductcategories/README.md)

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


### [TransactionsAccountTransactions](docs/sdks/transactionsaccounttransactions/README.md)

* [Get](docs/sdks/transactionsaccounttransactions/README.md#get) - Get account transaction
* [List](docs/sdks/transactionsaccounttransactions/README.md#list) - List account transactions

### [TransactionsDirectCosts](docs/sdks/transactionsdirectcosts/README.md)

* [DownloadAttachment](docs/sdks/transactionsdirectcosts/README.md#downloadattachment) - Download direct cost attachment
* [Get](docs/sdks/transactionsdirectcosts/README.md#get) - Get direct cost
* [GetAttachment](docs/sdks/transactionsdirectcosts/README.md#getattachment) - Get direct cost attachment
* [List](docs/sdks/transactionsdirectcosts/README.md#list) - List direct costs
* [ListAttachments](docs/sdks/transactionsdirectcosts/README.md#listattachments) - List direct cost attachments

### [TransactionsJournalEntries](docs/sdks/transactionsjournalentries/README.md)

* [Get](docs/sdks/transactionsjournalentries/README.md#get) - Get journal entry
* [List](docs/sdks/transactionsjournalentries/README.md#list) - List journal entries

### [Transactions.Journals](docs/sdks/transactionsjournals/README.md)

* [Get](docs/sdks/transactionsjournals/README.md#get) - Get journal
* [List](docs/sdks/transactionsjournals/README.md#list) - List journals

### [Transactions.Transfers](docs/sdks/transactionstransfers/README.md)

* [Get](docs/sdks/transactionstransfers/README.md#get) - Get transfer
* [List](docs/sdks/transactionstransfers/README.md#list) - List transfers<!-- Start SDK Available Operations -->

<!-- End SDK Available Operations -->
### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
