# github.com/codatio/client-sdk-go/accounting

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

import (
    "log"
    "github.com/codatio/client-sdk-go/accounting"
    "github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
    "github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    opts := []codatio.SDKOption{
        codatio.WithSecurity(
            shared.Security{
                APIKey: shared.SchemeAPIKey{
                    APIKey: "YOUR_API_KEY_HERE",
                },
            }
        ),
    }

    s := codatio.New(opts...)
    
    req := operations.GetAccountTransactionRequest{
        Security: operations.GetAccountTransactionSecurity{
            APIKey: shared.SchemeAPIKey{
                APIKey: "YOUR_API_KEY_HERE",
            },
        },
        PathParams: operations.GetAccountTransactionPathParams{
            AccountTransactionID: "unde",
            CompanyID: "deserunt",
            ConnectionID: "porro",
        },
    }
    
    res, err := s.AccountTransactions.GetAccountTransaction(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountTransaction != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations


### AccountTransactions

* `GetAccountTransaction` - Get account transaction
* `ListAccountTransactions` - List account transactions

### Accounts

* `GetAccount` - Get account
* `GetAccounts` - List accounts
* `PostAccount` - Create account

### BankAccountTransactions

* `GetBankAccountPushOptions` - List push options for bank account bank transactions
* `ListAllBankTransactionscount` - List bank transactions for bank account
* `ListBankTransactions` - List all bank transactions
* `PostBankTransactions` - Create bank transactions

### BankAccounts

* `GetAllBankAccount` - Get bank account
* `GetBankAccount` - Get bank account
* `ListBankAccounts` - List bank accounts
* `PostBankAccount` - Create bank account
* `PutBankAccount` - Update bank account

### BillCreditNotes

* `GetBillCreditNote` - Get bill credit note
* `ListBillCreditNotes` - List bill credit notes
* `PostBillCreditNote` - Create bill credit note
* `UpdateBillCreditNote` - Update bill credit note

### BillPayments

* `GetBillPayments` - Get bill payment
* `ListBillPayments` - List bill payments
* `PostBillPayment` - Create bill payment

### Bills

* `DownloadBillAttachment` - Download bill attachment
* `GetBill` - Get bill
* `GetBillAttachment` - Get bill attachment
* `GetBillAttachments` - List bill attachments
* `ListBills` - List bills
* `PostBill` - Create bill
* `PostBillAttachments` - Create bill attachments
* `UpdateBill` - Update bill

### CreditNotes

* `GetCreditNote` - Get credit note
* `ListCreditNotes` - List credit notes
* `PostCreditNote` - Update creditNote
* `PushCreditNote` - Create credit note

### Customers

* `DownloadCustomerAttachment` - Download customer attachment
* `GetCustomer` - Get customer
* `GetCustomerAttachment` - Get customer attachment
* `GetCustomerAttachments` - List customer attachments
* `GetCustomers` - List customers
* `PostCustomer` - Create customer
* `UpdateCustomer` - Update customer

### DirectCosts

* `DownloadDirectCostAttachment` - Download direct cost attachment
* `GetDirectCost` - Get directCost
* `GetDirectCostAttachment` - Get directCost attachment
* `GetDirectCosts` - List directCosts
* `ListDirectCostAttachments` - List direct cost attachments
* `PostDirectCost` - Create direct cost
* `PostDirectCostAttachment` - Create direct cost attachment

### DirectIncomes

* `DownloadDirectIncomeAttachment` - Download directIncome attachment
* `GetDirectIncome` - Get direct income
* `GetDirectIncomeAttachment` - Get direct income attachment
* `GetDirectIncomes` - Get direct incomes
* `ListDirectIncomeAttachments` - List direct income attachments
* `PostDirectIncome` - Create direct income
* `PostDirectIncomeAttachment` - Create direct income attachment

### Financials

* `GetBalanceSheet` - Get balance sheet
* `GetCashFlowStatement` - Get cash flow statement
* `GetProfitAndLoss` - Get profit and loss

### Info

* `GetCompanyInfo` - Get company info
* `PostSyncInfo` - Refresh company info

### Invoices

* `DonwloadInvoiceAttachment` - Download invoice attachment
* `GetInvoice` - Get invoice
* `GetInvoiceAttachment` - Get invoice attachment
* `GetInvoiceAttachments` - Get invoice attachments
* `GetInvoicePdf` - Get invoice as PDF
* `ListInvoices` - List invoices
* `PostInvoice` - Create invoice
* `PushInvoiceAttachment` - Push invoice attachment
* `UpdateInvoice` - Update invoice

### Items

* `GetItem` - Get item
* `ListItems` - List items
* `PostItem` - Create item

### JournalEntries

* `GetJournalEntry` - Get journal entry
* `ListJournalEntries` - List journal entries
* `PostJournalEntry` - Create journal entry

### Journals

* `GetJournal` - Get journal
* `ListJournals` - List journals
* `PushJournal` - Create journal

### PaymentMethods

* `GetPaymentMethod` - Get payment method
* `ListPaymentMethods` - List all payment methods

### Payments

* `GetPayment` - Get payment
* `ListPayments` - List payments
* `PostPayment` - Create payment

### PurchaseOrders

* `GetPurchaseOrder` - Get purchase order
* `ListPurchaseOrders` - List purchase orders
* `PostPurchaseOrder` - Create purchase order
* `UpdatePurchaseOrder` - Update purchase order

### SalesOrders

* `GetSalesOrder` - Get sales order
* `ListSalesOrders` - List sales orders

### Suppliers

* `DownloadSupplierAttachment` - Download supplier attachment
* `GetSupplier` - Get supplier
* `GetSupplierAttachment` - Get supplier attachment
* `ListSupplierAttachments` - List supplier attachments
* `ListSuppliers` - List suppliers
* `PostSuppliers` - Create suppliers
* `PutSupplier` - Update supplier

### TaxRates

* `GetTaxRate` - Get tax rate
* `ListTaxRates` - List all tax rates

### TrackingCategories

* `GetTrackingCategory` - Get tracking categories
* `ListTrackingCategories` - List tracking categories

### Transfers

* `GetTransfer` - Get transfer
* `ListTransfers` - List transfers
* `PostTransfer` - Create transfer
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
