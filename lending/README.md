# Lending
    
Lending helps you make smarter credit decisions on small businesses by enabling you to pull your customers' latest data from the operating systems they are already using. You can use that data for automating decisioning and surfacing new insights on the customer, all via one API.

## SDK Installation

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
* [ListTransfers](docs/sdks/transactions/README.md#listtransfers) - List transfers<!-- Start SDK Available Operations -->

<!-- End SDK Available Operations -->
### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
