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
    "context"
    "log"
    "github.com/codatio/client-sdk-go/accounting"
    "github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
    "github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codatio.New(
        WithSecurity(        shared.Security{
            AuthHeader: shared.SchemeAuthHeader{
                APIKey: "YOUR_API_KEY_HERE",
            },
        }),
    )

    req := operations.GetCreateUpdateAccountTransactionsModelRequest{
        PathParams: operations.GetCreateUpdateAccountTransactionsModelPathParams{
            AccountTransactionID: "unde",
            CompanyID: "deserunt",
            ConnectionID: "porro",
        },
    }

    ctx := context.Background()
    res, err := s.AccountTransactions.GetCreateUpdateAccountTransactionsModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SourceModifiedDate != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations


### AccountTransactions

* `GetCreateUpdateAccountTransactionsModel` - Get account transaction
* `ListAccountTransactions` - List account transactions

### Accounts

* `CreateAccount` - Create account
* `GetAccount` - Get account
* `GetAccounts` - List accounts
* `GetCreateChartOfAccountsModel` - Get create account model

### BankAccountTransactions

* `GetBankAccountPushOptions` - List push options for bank account bank transactions
* `ListBankAccountTransactions` - List bank transactions for bank account
* `ListBankTransactions` - List all bank transactions
* `PostBankTransactions` - Create bank transactions

### BankAccounts

* `CreateBankAccount` - Create bank account
* `GetAllBankAccount` - Get bank account
* `GetBankAccount` - Get bank account
* `GetCreateUpdateBankAccountsModel` - Get create/update bank account model
* `ListBankAccounts` - List bank accounts
* `PutBankAccount` - Update bank account

### BillCreditNotes

* `CreateBillCreditNote` - Create bill credit note
* `GetBillCreditNote` - Get bill credit note
* `GetCreateUpdateBillCreditNotesModel` - Get create/update bill credit note model
* `ListBillCreditNotes` - List bill credit notes
* `UpdateBillCreditNote` - Update bill credit note

### BillPayments

* `CreateBillPayment` - Create bill payments
* `DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillPaymentsBillPaymentID` - Delete bill payment
* `GetBillPayments` - Get bill payment
* `GetCreateBillPaymentsModel` - Get create bill payment model
* `ListBillPayments` - List bill payments

### Bills

* `CreateBill` - Create bill
* `CreateBillAttachments` - Create bill attachments
* `DeleteCompaniesCompanyIDConnectionsConnectionIDPushBillsBillID` - Delete bill
* `DownloadBillAttachment` - Download bill attachment
* `GetBill` - Get bill
* `GetBillAttachment` - Get bill attachment
* `GetBillAttachments` - List bill attachments
* `GetCreateUpdateBillsModel` - Get create/update bill model
* `ListBills` - List bills
* `UpdateBill` - Update bill

### CompanyInfo

* `GetCompanyInfo` - Get company info
* `PostSyncInfo` - Refresh company info

### CreditNotes

* `CreateCreditNote` - Update creditNote
* `GetCreateUpdateCreditNotesModel` - Get create/update credit note model
* `GetCreditNote` - Get credit note
* `ListCreditNotes` - List credit notes
* `PushCreditNote` - Create credit note

### Customers

* `CreateCustomer` - Create customer
* `DownloadCustomerAttachment` - Download customer attachment
* `GetCreateUpdateCustomersModel` - Get create/update customer model
* `GetCustomer` - Get customer
* `GetCustomerAttachment` - Get customer attachment
* `GetCustomerAttachments` - List customer attachments
* `GetCustomers` - List customers
* `UpdateCustomer` - Update customer

### DirectCosts

* `CreateDirectCost` - Create direct cost
* `DownloadDirectCostAttachment` - Download direct cost attachment
* `GetCreateDirectCostsModel` - Get create direct cost model
* `GetDirectCost` - Get direct cost
* `GetDirectCostAttachment` - Get direct cost attachment
* `GetDirectCosts` - List direct costs
* `ListDirectCostAttachments` - List direct cost attachments
* `PostDirectCostAttachment` - Create direct cost attachment

### DirectIncomes

* `CreateDirectIncome` - Create direct income
* `DownloadDirectIncomeAttachment` - Download direct income attachment
* `GetCreateDirectIncomesModel` - Get create direct income model
* `GetDirectIncome` - Get direct income
* `GetDirectIncomeAttachment` - Get direct income attachment
* `GetDirectIncomes` - Get direct incomes
* `ListDirectIncomeAttachments` - List direct income attachments
* `PostDirectIncomeAttachment` - Create direct income attachment

### Financials

* `GetBalanceSheet` - Get balance sheet
* `GetCashFlowStatement` - Get cash flow statement
* `GetProfitAndLoss` - Get profit and loss

### Invoices

* `CreateInvoice` - Create invoice
* `DonwloadInvoiceAttachment` - Download invoice attachment
* `GetCreateUpdateInvoicesModel` - Get create/update invoice model
* `GetInvoice` - Get invoice
* `GetInvoiceAttachment` - Get invoice attachment
* `GetInvoiceAttachments` - Get invoice attachments
* `GetInvoicePdf` - Get invoice as PDF
* `ListInvoices` - List invoices
* `PushInvoiceAttachment` - Push invoice attachment
* `UpdateInvoice` - Update invoice

### Items

* `CreateItem` - Create item
* `GetCreateItemsModel` - Get create item model
* `GetItem` - Get item
* `ListItems` - List items

### JournalEntries

* `CreateJournalEntry` - Create journal entry
* `GetCreateJournalEntriesModel` - Get create journal entry model
* `GetJournalEntry` - Get journal entry
* `ListJournalEntries` - List journal entries

### Journals

* `GetCreateJournalsModel` - Get create journal model
* `GetJournal` - Get journal
* `ListJournals` - List journals
* `PushJournal` - Create journal

### PaymentMethods

* `GetPaymentMethod` - Get payment method
* `ListPaymentMethods` - List all payment methods

### Payments

* `CreatePayment` - Create payment
* `GetCreatePaymentsModel` - Get create payment model
* `GetPayment` - Get payment
* `ListPayments` - List payments

### PurchaseOrders

* `CreatePurchaseOrder` - Create purchase order
* `GetCreateUpdatePurchaseOrdersModel` - Get create/update purchase order model
* `GetPurchaseOrder` - Get purchase order
* `ListPurchaseOrders` - List purchase orders
* `UpdatePurchaseOrder` - Update purchase order

### Reports

* `GetAgedCreditorsReport` - Aged creditors report
* `GetAgedDebtorsReport` - Aged debtors report
* `IsAgedCreditorsReportAvailable` - Aged creditors report available
* `IsAgedDebtorReportAvailable` - Aged debtors report available

### SalesOrders

* `GetSalesOrder` - Get sales order
* `ListSalesOrders` - List sales orders

### Suppliers

* `CreateSuppliers` - Create suppliers
* `DownloadSupplierAttachment` - Download supplier attachment
* `GetCreateUpdateSuppliersModel` - Get create/update supplier model
* `GetSupplier` - Get supplier
* `GetSupplierAttachment` - Get supplier attachment
* `ListSupplierAttachments` - List supplier attachments
* `ListSuppliers` - List suppliers
* `PutSupplier` - Update supplier

### TaxRates

* `GetTaxRate` - Get tax rate
* `ListTaxRates` - List all tax rates

### TrackingCategories

* `GetTrackingCategory` - Get tracking categories
* `ListTrackingCategories` - List tracking categories

### Transfers

* `CreateTransfer` - Create transfer
* `GetCreateTransfersModel` - Get create transfer model
* `GetTransfer` - Get transfer
* `ListTransfers` - List transfers
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
