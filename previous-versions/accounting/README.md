# Accounting

<!-- Start Codat Library Description -->
﻿Codat's Accounting API is a flexible API for pulling and pushing up-to-date accounting data to your customer's accounting software.
It gives you a simple way to view, create, update adn delete data without having to worry about each platform's specific complexities.

<!-- End Codat Library Description -->

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/previous-versions/accounting
```
<!-- End SDK Installation -->

## Example Usage
<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
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


### [AccountTransactions](docs/sdks/accounttransactions/README.md)

* [Get](docs/sdks/accounttransactions/README.md#get) - Get account transaction
* [List](docs/sdks/accounttransactions/README.md#list) - List account transactions

### [Accounts](docs/sdks/accounts/README.md)

* [Create](docs/sdks/accounts/README.md#create) - Create account
* [Get](docs/sdks/accounts/README.md#get) - Get account
* [GetCreateModel](docs/sdks/accounts/README.md#getcreatemodel) - Get create account model
* [List](docs/sdks/accounts/README.md#list) - List accounts

### [BankAccountTransactions](docs/sdks/bankaccounttransactions/README.md)

* [Create](docs/sdks/bankaccounttransactions/README.md#create) - Create bank account transactions
* [GetCreateModel](docs/sdks/bankaccounttransactions/README.md#getcreatemodel) - Get create bank account transactions model
* [List](docs/sdks/bankaccounttransactions/README.md#list) - List bank account transactions

### [BankAccounts](docs/sdks/bankaccounts/README.md)

* [Create](docs/sdks/bankaccounts/README.md#create) - Create bank account
* [Get](docs/sdks/bankaccounts/README.md#get) - Get bank account
* [GetCreateUpdateModel](docs/sdks/bankaccounts/README.md#getcreateupdatemodel) - Get create/update bank account model
* [List](docs/sdks/bankaccounts/README.md#list) - List bank accounts
* [Update](docs/sdks/bankaccounts/README.md#update) - Update bank account

### [BillCreditNotes](docs/sdks/billcreditnotes/README.md)

* [Create](docs/sdks/billcreditnotes/README.md#create) - Create bill credit note
* [Get](docs/sdks/billcreditnotes/README.md#get) - Get bill credit note
* [GetCreateUpdateModel](docs/sdks/billcreditnotes/README.md#getcreateupdatemodel) - Get create/update bill credit note model
* [List](docs/sdks/billcreditnotes/README.md#list) - List bill credit notes
* [Update](docs/sdks/billcreditnotes/README.md#update) - Update bill credit note
* [UploadAttachment](docs/sdks/billcreditnotes/README.md#uploadattachment) - Upload bill credit note attachment

### [BillPayments](docs/sdks/billpayments/README.md)

* [Create](docs/sdks/billpayments/README.md#create) - Create bill payments
* [Delete](docs/sdks/billpayments/README.md#delete) - Delete bill payment
* [Get](docs/sdks/billpayments/README.md#get) - Get bill payment
* [GetCreateModel](docs/sdks/billpayments/README.md#getcreatemodel) - Get create bill payment model
* [List](docs/sdks/billpayments/README.md#list) - List bill payments

### [Bills](docs/sdks/bills/README.md)

* [Create](docs/sdks/bills/README.md#create) - Create bill
* [Delete](docs/sdks/bills/README.md#delete) - Delete bill
* [DownloadAttachment](docs/sdks/bills/README.md#downloadattachment) - Download bill attachment
* [Get](docs/sdks/bills/README.md#get) - Get bill
* [GetAttachment](docs/sdks/bills/README.md#getattachment) - Get bill attachment
* [GetCreateUpdateModel](docs/sdks/bills/README.md#getcreateupdatemodel) - Get create/update bill model
* [List](docs/sdks/bills/README.md#list) - List bills
* [ListAttachments](docs/sdks/bills/README.md#listattachments) - List bill attachments
* [Update](docs/sdks/bills/README.md#update) - Update bill
* [UploadAttachment](docs/sdks/bills/README.md#uploadattachment) - Upload bill attachment

### [CompanyInfo](docs/sdks/companyinfo/README.md)

* [Get](docs/sdks/companyinfo/README.md#get) - Get company info
* [Refresh](docs/sdks/companyinfo/README.md#refresh) - Refresh company info

### [CreditNotes](docs/sdks/creditnotes/README.md)

* [Create](docs/sdks/creditnotes/README.md#create) - Create credit note
* [Get](docs/sdks/creditnotes/README.md#get) - Get credit note
* [GetCreateUpdateModel](docs/sdks/creditnotes/README.md#getcreateupdatemodel) - Get create/update credit note model
* [List](docs/sdks/creditnotes/README.md#list) - List credit notes
* [Update](docs/sdks/creditnotes/README.md#update) - Update creditNote

### [Customers](docs/sdks/customers/README.md)

* [Create](docs/sdks/customers/README.md#create) - Create customer
* [DownloadAttachment](docs/sdks/customers/README.md#downloadattachment) - Download customer attachment
* [Get](docs/sdks/customers/README.md#get) - Get customer
* [GetAttachment](docs/sdks/customers/README.md#getattachment) - Get customer attachment
* [GetCreateUpdateModel](docs/sdks/customers/README.md#getcreateupdatemodel) - Get create/update customer model
* [List](docs/sdks/customers/README.md#list) - List customers
* [ListAttachments](docs/sdks/customers/README.md#listattachments) - List customer attachments
* [Update](docs/sdks/customers/README.md#update) - Update customer

### [DirectCosts](docs/sdks/directcosts/README.md)

* [Create](docs/sdks/directcosts/README.md#create) - Create direct cost
* [DownloadAttachment](docs/sdks/directcosts/README.md#downloadattachment) - Download direct cost attachment
* [Get](docs/sdks/directcosts/README.md#get) - Get direct cost
* [GetAttachment](docs/sdks/directcosts/README.md#getattachment) - Get direct cost attachment
* [GetCreateModel](docs/sdks/directcosts/README.md#getcreatemodel) - Get create direct cost model
* [List](docs/sdks/directcosts/README.md#list) - List direct costs
* [ListAttachments](docs/sdks/directcosts/README.md#listattachments) - List direct cost attachments
* [UploadAttachment](docs/sdks/directcosts/README.md#uploadattachment) - Upload direct cost attachment

### [DirectIncomes](docs/sdks/directincomes/README.md)

* [Create](docs/sdks/directincomes/README.md#create) - Create direct income
* [DownloadAttachment](docs/sdks/directincomes/README.md#downloadattachment) - Download direct income attachment
* [Get](docs/sdks/directincomes/README.md#get) - Get direct income
* [GetAttachment](docs/sdks/directincomes/README.md#getattachment) - Get direct income attachment
* [GetCreateModel](docs/sdks/directincomes/README.md#getcreatemodel) - Get create direct income model
* [List](docs/sdks/directincomes/README.md#list) - List direct incomes
* [ListAttachments](docs/sdks/directincomes/README.md#listattachments) - List direct income attachments
* [UploadAttachment](docs/sdks/directincomes/README.md#uploadattachment) - Create direct income attachment

### [Invoices](docs/sdks/invoices/README.md)

* [Create](docs/sdks/invoices/README.md#create) - Create invoice
* [Delete](docs/sdks/invoices/README.md#delete) - Delete invoice
* [DownloadAttachment](docs/sdks/invoices/README.md#downloadattachment) - Download invoice attachment
* [DownloadPdf](docs/sdks/invoices/README.md#downloadpdf) - Get invoice as PDF
* [Get](docs/sdks/invoices/README.md#get) - Get invoice
* [GetAttachment](docs/sdks/invoices/README.md#getattachment) - Get invoice attachment
* [GetCreateUpdateModel](docs/sdks/invoices/README.md#getcreateupdatemodel) - Get create/update invoice model
* [List](docs/sdks/invoices/README.md#list) - List invoices
* [ListAttachments](docs/sdks/invoices/README.md#listattachments) - List invoice attachments
* [Update](docs/sdks/invoices/README.md#update) - Update invoice
* [UploadAttachment](docs/sdks/invoices/README.md#uploadattachment) - Push invoice attachment

### [Items](docs/sdks/items/README.md)

* [Create](docs/sdks/items/README.md#create) - Create item
* [Get](docs/sdks/items/README.md#get) - Get item
* [GetCreateModel](docs/sdks/items/README.md#getcreatemodel) - Get create item model
* [List](docs/sdks/items/README.md#list) - List items

### [JournalEntries](docs/sdks/journalentries/README.md)

* [Create](docs/sdks/journalentries/README.md#create) - Create journal entry
* [Delete](docs/sdks/journalentries/README.md#delete) - Delete journal entry
* [Get](docs/sdks/journalentries/README.md#get) - Get journal entry
* [GetCreateModel](docs/sdks/journalentries/README.md#getcreatemodel) - Get create journal entry model
* [List](docs/sdks/journalentries/README.md#list) - List journal entries

### [Journals](docs/sdks/journals/README.md)

* [Create](docs/sdks/journals/README.md#create) - Create journal
* [Get](docs/sdks/journals/README.md#get) - Get journal
* [GetCreateModel](docs/sdks/journals/README.md#getcreatemodel) - Get create journal model
* [List](docs/sdks/journals/README.md#list) - List journals

### [PaymentMethods](docs/sdks/paymentmethods/README.md)

* [Get](docs/sdks/paymentmethods/README.md#get) - Get payment method
* [List](docs/sdks/paymentmethods/README.md#list) - List payment methods

### [Payments](docs/sdks/payments/README.md)

* [Create](docs/sdks/payments/README.md#create) - Create payment
* [Get](docs/sdks/payments/README.md#get) - Get payment
* [GetCreateModel](docs/sdks/payments/README.md#getcreatemodel) - Get create payment model
* [List](docs/sdks/payments/README.md#list) - List payments

### [PurchaseOrders](docs/sdks/purchaseorders/README.md)

* [Create](docs/sdks/purchaseorders/README.md#create) - Create purchase order
* [Get](docs/sdks/purchaseorders/README.md#get) - Get purchase order
* [GetCreateUpdateModel](docs/sdks/purchaseorders/README.md#getcreateupdatemodel) - Get create/update purchase order model
* [List](docs/sdks/purchaseorders/README.md#list) - List purchase orders
* [Update](docs/sdks/purchaseorders/README.md#update) - Update purchase order

### [Reports](docs/sdks/reports/README.md)

* [GetAgedCreditorsReport](docs/sdks/reports/README.md#getagedcreditorsreport) - Aged creditors report
* [GetAgedDebtorsReport](docs/sdks/reports/README.md#getageddebtorsreport) - Aged debtors report
* [GetBalanceSheet](docs/sdks/reports/README.md#getbalancesheet) - Get balance sheet
* [GetCashFlowStatement](docs/sdks/reports/README.md#getcashflowstatement) - Get cash flow statement
* [GetProfitAndLoss](docs/sdks/reports/README.md#getprofitandloss) - Get profit and loss
* [IsAgedCreditorsReportAvailable](docs/sdks/reports/README.md#isagedcreditorsreportavailable) - Aged creditors report available
* [IsAgedDebtorReportAvailable](docs/sdks/reports/README.md#isageddebtorreportavailable) - Aged debtors report available

### [SalesOrders](docs/sdks/salesorders/README.md)

* [Get](docs/sdks/salesorders/README.md#get) - Get sales order
* [List](docs/sdks/salesorders/README.md#list) - List sales orders

### [Suppliers](docs/sdks/suppliers/README.md)

* [Create](docs/sdks/suppliers/README.md#create) - Create supplier
* [DownloadAttachment](docs/sdks/suppliers/README.md#downloadattachment) - Download supplier attachment
* [Get](docs/sdks/suppliers/README.md#get) - Get supplier
* [GetAttachment](docs/sdks/suppliers/README.md#getattachment) - Get supplier attachment
* [GetCreateUpdateModel](docs/sdks/suppliers/README.md#getcreateupdatemodel) - Get create/update supplier model
* [List](docs/sdks/suppliers/README.md#list) - List suppliers
* [ListAttachments](docs/sdks/suppliers/README.md#listattachments) - List supplier attachments
* [Update](docs/sdks/suppliers/README.md#update) - Update supplier

### [TaxRates](docs/sdks/taxrates/README.md)

* [Get](docs/sdks/taxrates/README.md#get) - Get tax rate
* [List](docs/sdks/taxrates/README.md#list) - List all tax rates

### [TrackingCategories](docs/sdks/trackingcategories/README.md)

* [Get](docs/sdks/trackingcategories/README.md#get) - Get tracking categories
* [List](docs/sdks/trackingcategories/README.md#list) - List tracking categories

### [Transfers](docs/sdks/transfers/README.md)

* [Create](docs/sdks/transfers/README.md#create) - Create transfer
* [Get](docs/sdks/transfers/README.md#get) - Get transfer
* [GetCreateModel](docs/sdks/transfers/README.md#getcreatemodel) - Get create transfer model
* [List](docs/sdks/transfers/README.md#list) - List transfers
* [UploadAttachment](docs/sdks/transfers/README.md#uploadattachment) - Push invoice attachment
<!-- End SDK Available Operations -->