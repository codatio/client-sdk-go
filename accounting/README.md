# Accounting API

Codat's Accounting API is a flexible API for pulling and pushing up-to-date accounting data to your customer's accounting software.
It gives you a simple way to view, create, update adn delete data without having to worry about each platform's specific complexities.

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/accounting
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountTransactions.Get(ctx, operations.GetAccountTransactionRequest{
        AccountTransactionID: "corrupti",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountTransaction != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [AccountTransactions](docs/accounttransactions/README.md)

* [Get](docs/accounttransactions/README.md#get) - Get account transaction
* [List](docs/accounttransactions/README.md#list) - List account transactions

### [Accounts](docs/accounts/README.md)

* [Create](docs/accounts/README.md#create) - Create account
* [Get](docs/accounts/README.md#get) - Get account
* [GetCreateModel](docs/accounts/README.md#getcreatemodel) - Get create account model
* [List](docs/accounts/README.md#list) - List accounts

### [BankAccountTransactions](docs/bankaccounttransactions/README.md)

* [Create](docs/bankaccounttransactions/README.md#create) - Create bank transactions
* [GetCreateModel](docs/bankaccounttransactions/README.md#getcreatemodel) - List push options for bank account bank transactions
* [List](docs/bankaccounttransactions/README.md#list) - List bank transactions for bank account
* [ListTransactions](docs/bankaccounttransactions/README.md#listtransactions) - List all bank transactions

### [BankAccounts](docs/bankaccounts/README.md)

* [Create](docs/bankaccounts/README.md#create) - Create bank account
* [Get](docs/bankaccounts/README.md#get) - Get bank account
* [GetCreateUpdateModel](docs/bankaccounts/README.md#getcreateupdatemodel) - Get create/update bank account model
* [List](docs/bankaccounts/README.md#list) - List bank accounts
* [Update](docs/bankaccounts/README.md#update) - Update bank account

### [BillCreditNotes](docs/billcreditnotes/README.md)

* [Create](docs/billcreditnotes/README.md#create) - Create bill credit note
* [Get](docs/billcreditnotes/README.md#get) - Get bill credit note
* [GetCreateUpdateModel](docs/billcreditnotes/README.md#getcreateupdatemodel) - Get create/update bill credit note model
* [List](docs/billcreditnotes/README.md#list) - List bill credit notes
* [Update](docs/billcreditnotes/README.md#update) - Update bill credit note

### [BillPayments](docs/billpayments/README.md)

* [Create](docs/billpayments/README.md#create) - Create bill payments
* [Delete](docs/billpayments/README.md#delete) - Delete bill payment
* [Get](docs/billpayments/README.md#get) - Get bill payment
* [GetCreateModel](docs/billpayments/README.md#getcreatemodel) - Get create bill payment model
* [List](docs/billpayments/README.md#list) - List bill payments

### [Bills](docs/bills/README.md)

* [Create](docs/bills/README.md#create) - Create bill
* [Delete](docs/bills/README.md#delete) - Delete bill
* [DownloadAttachment](docs/bills/README.md#downloadattachment) - Download bill attachment
* [Get](docs/bills/README.md#get) - Get bill
* [GetAttachment](docs/bills/README.md#getattachment) - Get bill attachment
* [GetCreateUpdateModel](docs/bills/README.md#getcreateupdatemodel) - Get create/update bill model
* [List](docs/bills/README.md#list) - List bills
* [ListAttachments](docs/bills/README.md#listattachments) - List bill attachments
* [Update](docs/bills/README.md#update) - Update bill
* [UploadAttachment](docs/bills/README.md#uploadattachment) - Upload bill attachment

### [CompanyInfo](docs/companyinfo/README.md)

* [Get](docs/companyinfo/README.md#get) - Get company info
* [Refresh](docs/companyinfo/README.md#refresh) - Refresh company info

### [CreditNotes](docs/creditnotes/README.md)

* [Create](docs/creditnotes/README.md#create) - Create credit note
* [Get](docs/creditnotes/README.md#get) - Get credit note
* [GetCreateUpdateModel](docs/creditnotes/README.md#getcreateupdatemodel) - Get create/update credit note model
* [List](docs/creditnotes/README.md#list) - List credit notes
* [Update](docs/creditnotes/README.md#update) - Update creditNote

### [Customers](docs/customers/README.md)

* [Create](docs/customers/README.md#create) - Create customer
* [DownloadAttachment](docs/customers/README.md#downloadattachment) - Download customer attachment
* [Get](docs/customers/README.md#get) - Get customer
* [GetAttachment](docs/customers/README.md#getattachment) - Get customer attachment
* [GetCreateUpdateModel](docs/customers/README.md#getcreateupdatemodel) - Get create/update customer model
* [List](docs/customers/README.md#list) - List customers
* [ListAttachments](docs/customers/README.md#listattachments) - List customer attachments
* [Update](docs/customers/README.md#update) - Update customer

### [DirectCosts](docs/directcosts/README.md)

* [Create](docs/directcosts/README.md#create) - Create direct cost
* [DownloadAttachment](docs/directcosts/README.md#downloadattachment) - Download direct cost attachment
* [Get](docs/directcosts/README.md#get) - Get direct cost
* [GetAttachment](docs/directcosts/README.md#getattachment) - Get direct cost attachment
* [GetCreateModel](docs/directcosts/README.md#getcreatemodel) - Get create direct cost model
* [List](docs/directcosts/README.md#list) - List direct costs
* [ListAttachments](docs/directcosts/README.md#listattachments) - List direct cost attachments
* [UploadAttachment](docs/directcosts/README.md#uploadattachment) - Upload direct cost attachment

### [DirectIncomes](docs/directincomes/README.md)

* [Create](docs/directincomes/README.md#create) - Create direct income
* [DownloadAttachment](docs/directincomes/README.md#downloadattachment) - Download direct income attachment
* [Get](docs/directincomes/README.md#get) - Get direct income
* [GetAttachment](docs/directincomes/README.md#getattachment) - Get direct income attachment
* [GetCreateModel](docs/directincomes/README.md#getcreatemodel) - Get create direct income model
* [List](docs/directincomes/README.md#list) - List direct incomes
* [ListAttachments](docs/directincomes/README.md#listattachments) - List direct income attachments
* [UploadAttachment](docs/directincomes/README.md#uploadattachment) - Create direct income attachment

### [Financials](docs/financials/README.md)

* [GetBalanceSheet](docs/financials/README.md#getbalancesheet) - Get balance sheet
* [GetCashFlowStatement](docs/financials/README.md#getcashflowstatement) - Get cash flow statement
* [GetProfitAndLoss](docs/financials/README.md#getprofitandloss) - Get profit and loss

### [Invoices](docs/invoices/README.md)

* [Create](docs/invoices/README.md#create) - Create invoice
* [Delete](docs/invoices/README.md#delete) - Delete invoice
* [DownloadAttachment](docs/invoices/README.md#downloadattachment) - Download invoice attachment
* [DownloadPdf](docs/invoices/README.md#downloadpdf) - Get invoice as PDF
* [Get](docs/invoices/README.md#get) - Get invoice
* [GetAttachment](docs/invoices/README.md#getattachment) - Get invoice attachment
* [GetCreateUpdateModel](docs/invoices/README.md#getcreateupdatemodel) - Get create/update invoice model
* [List](docs/invoices/README.md#list) - List invoices
* [ListAttachments](docs/invoices/README.md#listattachments) - List invoice attachments
* [Update](docs/invoices/README.md#update) - Update invoice
* [UploadAttachment](docs/invoices/README.md#uploadattachment) - Push invoice attachment

### [Items](docs/items/README.md)

* [Create](docs/items/README.md#create) - Create item
* [Get](docs/items/README.md#get) - Get item
* [GetCreateModel](docs/items/README.md#getcreatemodel) - Get create item model
* [List](docs/items/README.md#list) - List items

### [JournalEntries](docs/journalentries/README.md)

* [Create](docs/journalentries/README.md#create) - Create journal entry
* [Delete](docs/journalentries/README.md#delete) - Delete journal entry
* [Get](docs/journalentries/README.md#get) - Get journal entry
* [GetCreateModel](docs/journalentries/README.md#getcreatemodel) - Get create journal entry model
* [List](docs/journalentries/README.md#list) - List journal entries

### [Journals](docs/journals/README.md)

* [Create](docs/journals/README.md#create) - Create journal
* [Get](docs/journals/README.md#get) - Get journal
* [GetCreateModel](docs/journals/README.md#getcreatemodel) - Get create journal model
* [List](docs/journals/README.md#list) - List journals

### [PaymentMethods](docs/paymentmethods/README.md)

* [Get](docs/paymentmethods/README.md#get) - Get payment method
* [List](docs/paymentmethods/README.md#list) - List all payment methods

### [Payments](docs/payments/README.md)

* [Create](docs/payments/README.md#create) - Create payment
* [Get](docs/payments/README.md#get) - Get payment
* [GetCreateModel](docs/payments/README.md#getcreatemodel) - Get create payment model
* [List](docs/payments/README.md#list) - List payments

### [PurchaseOrders](docs/purchaseorders/README.md)

* [Create](docs/purchaseorders/README.md#create) - Create purchase order
* [Get](docs/purchaseorders/README.md#get) - Get purchase order
* [GetCreateUpdateModel](docs/purchaseorders/README.md#getcreateupdatemodel) - Get create/update purchase order model
* [List](docs/purchaseorders/README.md#list) - List purchase orders
* [Update](docs/purchaseorders/README.md#update) - Update purchase order

### [Reports](docs/reports/README.md)

* [GetAgedCreditorsReport](docs/reports/README.md#getagedcreditorsreport) - Aged creditors report
* [GetAgedDebtorsReport](docs/reports/README.md#getageddebtorsreport) - Aged debtors report
* [IsAgedCreditorsReportAvailable](docs/reports/README.md#isagedcreditorsreportavailable) - Aged creditors report available
* [IsAgedDebtorReportAvailable](docs/reports/README.md#isageddebtorreportavailable) - Aged debtors report available

### [SalesOrders](docs/salesorders/README.md)

* [Get](docs/salesorders/README.md#get) - Get sales order
* [List](docs/salesorders/README.md#list) - List sales orders

### [Suppliers](docs/suppliers/README.md)

* [Create](docs/suppliers/README.md#create) - Create supplier
* [DownloadAttachment](docs/suppliers/README.md#downloadattachment) - Download supplier attachment
* [Get](docs/suppliers/README.md#get) - Get supplier
* [GetAttachment](docs/suppliers/README.md#getattachment) - Get supplier attachment
* [GetCreateUpdateModel](docs/suppliers/README.md#getcreateupdatemodel) - Get create/update supplier model
* [List](docs/suppliers/README.md#list) - List suppliers
* [ListAttachments](docs/suppliers/README.md#listattachments) - List supplier attachments
* [Update](docs/suppliers/README.md#update) - Update supplier

### [TaxRates](docs/taxrates/README.md)

* [Get](docs/taxrates/README.md#get) - Get tax rate
* [List](docs/taxrates/README.md#list) - List all tax rates

### [TrackingCategories](docs/trackingcategories/README.md)

* [Get](docs/trackingcategories/README.md#get) - Get tracking categories
* [List](docs/trackingcategories/README.md#list) - List tracking categories

### [Transfers](docs/transfers/README.md)

* [Create](docs/transfers/README.md#create) - Create transfer
* [Get](docs/transfers/README.md#get) - Get transfer
* [GetCreateModel](docs/transfers/README.md#getcreatemodel) - Get create transfer model
* [List](docs/transfers/README.md#list) - List transfers
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
