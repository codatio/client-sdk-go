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
    req := operations.GetAccountTransactionRequest{
        AccountTransactionID: "corrupti",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.AccountTransactions.GetAccountTransaction(ctx, req)
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

* [GetAccountTransaction](docs/accounttransactions/README.md#getaccounttransaction) - Get account transaction
* [ListAccountTransactions](docs/accounttransactions/README.md#listaccounttransactions) - List account transactions

### [Accounts](docs/accounts/README.md)

* [CreateAccount](docs/accounts/README.md#createaccount) - Create account
* [GetAccount](docs/accounts/README.md#getaccount) - Get account
* [GetCreateChartOfAccountsModel](docs/accounts/README.md#getcreatechartofaccountsmodel) - Get create account model
* [ListAccounts](docs/accounts/README.md#listaccounts) - List accounts

### [BankAccountTransactions](docs/bankaccounttransactions/README.md)

* [CreateBankTransactions](docs/bankaccounttransactions/README.md#createbanktransactions) - Create bank transactions
* [GetCreateBankAccountModel](docs/bankaccounttransactions/README.md#getcreatebankaccountmodel) - List push options for bank account bank transactions
* [ListBankAccountTransactions](docs/bankaccounttransactions/README.md#listbankaccounttransactions) - List bank transactions for bank account
* [ListBankTransactions](docs/bankaccounttransactions/README.md#listbanktransactions) - List all bank transactions

### [BankAccounts](docs/bankaccounts/README.md)

* [CreateBankAccount](docs/bankaccounts/README.md#createbankaccount) - Create bank account
* [GetAllBankAccount](docs/bankaccounts/README.md#getallbankaccount) - Get bank account
* [GetBankAccount](docs/bankaccounts/README.md#getbankaccount) - Get bank account
* [GetCreateUpdateBankAccountsModel](docs/bankaccounts/README.md#getcreateupdatebankaccountsmodel) - Get create/update bank account model
* [ListBankAccounts](docs/bankaccounts/README.md#listbankaccounts) - List bank accounts
* [UpdateBankAccount](docs/bankaccounts/README.md#updatebankaccount) - Update bank account

### [BillCreditNotes](docs/billcreditnotes/README.md)

* [CreateBillCreditNote](docs/billcreditnotes/README.md#createbillcreditnote) - Create bill credit note
* [GetBillCreditNote](docs/billcreditnotes/README.md#getbillcreditnote) - Get bill credit note
* [GetCreateUpdateBillCreditNotesModel](docs/billcreditnotes/README.md#getcreateupdatebillcreditnotesmodel) - Get create/update bill credit note model
* [ListBillCreditNotes](docs/billcreditnotes/README.md#listbillcreditnotes) - List bill credit notes
* [UpdateBillCreditNote](docs/billcreditnotes/README.md#updatebillcreditnote) - Update bill credit note

### [BillPayments](docs/billpayments/README.md)

* [CreateBillPayment](docs/billpayments/README.md#createbillpayment) - Create bill payments
* [DeleteBillPayment](docs/billpayments/README.md#deletebillpayment) - Delete bill payment
* [GetBillPayments](docs/billpayments/README.md#getbillpayments) - Get bill payment
* [GetCreateBillPaymentsModel](docs/billpayments/README.md#getcreatebillpaymentsmodel) - Get create bill payment model
* [ListBillPayments](docs/billpayments/README.md#listbillpayments) - List bill payments

### [Bills](docs/bills/README.md)

* [CreateBill](docs/bills/README.md#createbill) - Create bill
* [DeleteBill](docs/bills/README.md#deletebill) - Delete bill
* [DownloadBillAttachment](docs/bills/README.md#downloadbillattachment) - Download bill attachment
* [GetBill](docs/bills/README.md#getbill) - Get bill
* [GetBillAttachment](docs/bills/README.md#getbillattachment) - Get bill attachment
* [GetBillAttachments](docs/bills/README.md#getbillattachments) - List bill attachments
* [GetCreateUpdateBillsModel](docs/bills/README.md#getcreateupdatebillsmodel) - Get create/update bill model
* [ListBills](docs/bills/README.md#listbills) - List bills
* [UpdateBill](docs/bills/README.md#updatebill) - Update bill
* [UploadBillAttachments](docs/bills/README.md#uploadbillattachments) - Upload bill attachments

### [CompanyInfo](docs/companyinfo/README.md)

* [GetCompanyInfo](docs/companyinfo/README.md#getcompanyinfo) - Get company info
* [PostSyncInfo](docs/companyinfo/README.md#postsyncinfo) - Refresh company info

### [CreditNotes](docs/creditnotes/README.md)

* [CreateCreditNote](docs/creditnotes/README.md#createcreditnote) - Create credit note
* [GetCreateUpdateCreditNotesModel](docs/creditnotes/README.md#getcreateupdatecreditnotesmodel) - Get create/update credit note model
* [GetCreditNote](docs/creditnotes/README.md#getcreditnote) - Get credit note
* [ListCreditNotes](docs/creditnotes/README.md#listcreditnotes) - List credit notes
* [UpdateCreditNote](docs/creditnotes/README.md#updatecreditnote) - Update creditNote

### [Customers](docs/customers/README.md)

* [CreateCustomer](docs/customers/README.md#createcustomer) - Create customer
* [DownloadCustomerAttachment](docs/customers/README.md#downloadcustomerattachment) - Download customer attachment
* [GetCreateUpdateCustomersModel](docs/customers/README.md#getcreateupdatecustomersmodel) - Get create/update customer model
* [GetCustomer](docs/customers/README.md#getcustomer) - Get customer
* [GetCustomerAttachment](docs/customers/README.md#getcustomerattachment) - Get customer attachment
* [GetCustomerAttachments](docs/customers/README.md#getcustomerattachments) - List customer attachments
* [GetCustomers](docs/customers/README.md#getcustomers) - List customers
* [UpdateCustomer](docs/customers/README.md#updatecustomer) - Update customer

### [DirectCosts](docs/directcosts/README.md)

* [CreateDirectCost](docs/directcosts/README.md#createdirectcost) - Create direct cost
* [DownloadDirectCostAttachment](docs/directcosts/README.md#downloaddirectcostattachment) - Download direct cost attachment
* [GetCreateDirectCostsModel](docs/directcosts/README.md#getcreatedirectcostsmodel) - Get create direct cost model
* [GetDirectCost](docs/directcosts/README.md#getdirectcost) - Get direct cost
* [GetDirectCostAttachment](docs/directcosts/README.md#getdirectcostattachment) - Get direct cost attachment
* [GetDirectCosts](docs/directcosts/README.md#getdirectcosts) - List direct costs
* [ListDirectCostAttachments](docs/directcosts/README.md#listdirectcostattachments) - List direct cost attachments
* [UploadDirectCostAttachment](docs/directcosts/README.md#uploaddirectcostattachment) - Upload direct cost attachment

### [DirectIncomes](docs/directincomes/README.md)

* [CreateDirectIncome](docs/directincomes/README.md#createdirectincome) - Create direct income
* [DownloadDirectIncomeAttachment](docs/directincomes/README.md#downloaddirectincomeattachment) - Download direct income attachment
* [GetCreateDirectIncomesModel](docs/directincomes/README.md#getcreatedirectincomesmodel) - Get create direct income model
* [GetDirectIncome](docs/directincomes/README.md#getdirectincome) - Get direct income
* [GetDirectIncomeAttachment](docs/directincomes/README.md#getdirectincomeattachment) - Get direct income attachment
* [GetDirectIncomes](docs/directincomes/README.md#getdirectincomes) - Get direct incomes
* [ListDirectIncomeAttachments](docs/directincomes/README.md#listdirectincomeattachments) - List direct income attachments
* [UploadDirectIncomeAttachment](docs/directincomes/README.md#uploaddirectincomeattachment) - Create direct income attachment

### [Financials](docs/financials/README.md)

* [GetBalanceSheet](docs/financials/README.md#getbalancesheet) - Get balance sheet
* [GetCashFlowStatement](docs/financials/README.md#getcashflowstatement) - Get cash flow statement
* [GetProfitAndLoss](docs/financials/README.md#getprofitandloss) - Get profit and loss

### [Invoices](docs/invoices/README.md)

* [DownloadInvoicePdf](docs/invoices/README.md#downloadinvoicepdf) - Get invoice as PDF
* [CreateInvoice](docs/invoices/README.md#createinvoice) - Create invoice
* [DeleteInvoice](docs/invoices/README.md#deleteinvoice) - Delete invoice
* [DownloadInvoiceAttachment](docs/invoices/README.md#downloadinvoiceattachment) - Download invoice attachment
* [GetCreateUpdateInvoicesModel](docs/invoices/README.md#getcreateupdateinvoicesmodel) - Get create/update invoice model
* [GetInvoice](docs/invoices/README.md#getinvoice) - Get invoice
* [GetInvoiceAttachment](docs/invoices/README.md#getinvoiceattachment) - Get invoice attachment
* [GetInvoiceAttachments](docs/invoices/README.md#getinvoiceattachments) - Get invoice attachments
* [ListInvoices](docs/invoices/README.md#listinvoices) - List invoices
* [UpdateInvoice](docs/invoices/README.md#updateinvoice) - Update invoice
* [UploadInvoiceAttachment](docs/invoices/README.md#uploadinvoiceattachment) - Push invoice attachment

### [Items](docs/items/README.md)

* [CreateItem](docs/items/README.md#createitem) - Create item
* [GetCreateItemsModel](docs/items/README.md#getcreateitemsmodel) - Get create item model
* [GetItem](docs/items/README.md#getitem) - Get item
* [ListItems](docs/items/README.md#listitems) - List items

### [JournalEntries](docs/journalentries/README.md)

* [CreateJournalEntry](docs/journalentries/README.md#createjournalentry) - Create journal entry
* [DeleteJournalEntry](docs/journalentries/README.md#deletejournalentry) - Delete journal entry
* [GetCreateJournalEntriesModel](docs/journalentries/README.md#getcreatejournalentriesmodel) - Get create journal entry model
* [GetJournalEntry](docs/journalentries/README.md#getjournalentry) - Get journal entry
* [ListJournalEntries](docs/journalentries/README.md#listjournalentries) - List journal entries

### [Journals](docs/journals/README.md)

* [GetCreateJournalsModel](docs/journals/README.md#getcreatejournalsmodel) - Get create journal model
* [GetJournal](docs/journals/README.md#getjournal) - Get journal
* [ListJournals](docs/journals/README.md#listjournals) - List journals
* [PushJournal](docs/journals/README.md#pushjournal) - Create journal

### [PaymentMethods](docs/paymentmethods/README.md)

* [GetPaymentMethod](docs/paymentmethods/README.md#getpaymentmethod) - Get payment method
* [ListPaymentMethods](docs/paymentmethods/README.md#listpaymentmethods) - List all payment methods

### [Payments](docs/payments/README.md)

* [CreatePayment](docs/payments/README.md#createpayment) - Create payment
* [GetCreatePaymentsModel](docs/payments/README.md#getcreatepaymentsmodel) - Get create payment model
* [GetPayment](docs/payments/README.md#getpayment) - Get payment
* [ListPayments](docs/payments/README.md#listpayments) - List payments

### [PurchaseOrders](docs/purchaseorders/README.md)

* [CreatePurchaseOrder](docs/purchaseorders/README.md#createpurchaseorder) - Create purchase order
* [GetCreateUpdatePurchaseOrdersModel](docs/purchaseorders/README.md#getcreateupdatepurchaseordersmodel) - Get create/update purchase order model
* [GetPurchaseOrder](docs/purchaseorders/README.md#getpurchaseorder) - Get purchase order
* [ListPurchaseOrders](docs/purchaseorders/README.md#listpurchaseorders) - List purchase orders
* [UpdatePurchaseOrder](docs/purchaseorders/README.md#updatepurchaseorder) - Update purchase order

### [Reports](docs/reports/README.md)

* [GetAgedCreditorsReport](docs/reports/README.md#getagedcreditorsreport) - Aged creditors report
* [GetAgedDebtorsReport](docs/reports/README.md#getageddebtorsreport) - Aged debtors report
* [IsAgedCreditorsReportAvailable](docs/reports/README.md#isagedcreditorsreportavailable) - Aged creditors report available
* [IsAgedDebtorReportAvailable](docs/reports/README.md#isageddebtorreportavailable) - Aged debtors report available

### [SalesOrders](docs/salesorders/README.md)

* [GetSalesOrder](docs/salesorders/README.md#getsalesorder) - Get sales order
* [ListSalesOrders](docs/salesorders/README.md#listsalesorders) - List sales orders

### [Suppliers](docs/suppliers/README.md)

* [CreateSupplier](docs/suppliers/README.md#createsupplier) - Create suppliers
* [DownloadSupplierAttachment](docs/suppliers/README.md#downloadsupplierattachment) - Download supplier attachment
* [GetCreateUpdateSuppliersModel](docs/suppliers/README.md#getcreateupdatesuppliersmodel) - Get create/update supplier model
* [GetSupplier](docs/suppliers/README.md#getsupplier) - Get supplier
* [GetSupplierAttachment](docs/suppliers/README.md#getsupplierattachment) - Get supplier attachment
* [ListSupplierAttachments](docs/suppliers/README.md#listsupplierattachments) - List supplier attachments
* [ListSuppliers](docs/suppliers/README.md#listsuppliers) - List suppliers
* [PutSupplier](docs/suppliers/README.md#putsupplier) - Update supplier

### [TaxRates](docs/taxrates/README.md)

* [GetTaxRate](docs/taxrates/README.md#gettaxrate) - Get tax rate
* [ListTaxRates](docs/taxrates/README.md#listtaxrates) - List all tax rates

### [TrackingCategories](docs/trackingcategories/README.md)

* [GetTrackingCategory](docs/trackingcategories/README.md#gettrackingcategory) - Get tracking categories
* [ListTrackingCategories](docs/trackingcategories/README.md#listtrackingcategories) - List tracking categories

### [Transfers](docs/transfers/README.md)

* [CreateTransfer](docs/transfers/README.md#createtransfer) - Create transfer
* [GetCreateTransfersModel](docs/transfers/README.md#getcreatetransfersmodel) - Get create transfer model
* [GetTransfer](docs/transfers/README.md#gettransfer) - Get transfer
* [ListTransfers](docs/transfers/README.md#listtransfers) - List transfers
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
