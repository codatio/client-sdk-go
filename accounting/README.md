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

* [GetAccountTransaction](docs/accounttransactions/getaccounttransaction.md) - Get account transaction
* [ListAccountTransactions](docs/accounttransactions/listaccounttransactions.md) - List account transactions

### [Accounts](docs/accounts/README.md)

* [CreateAccount](docs/accounts/createaccount.md) - Create account
* [GetAccount](docs/accounts/getaccount.md) - Get account
* [GetCreateChartOfAccountsModel](docs/accounts/getcreatechartofaccountsmodel.md) - Get create account model
* [ListAccounts](docs/accounts/listaccounts.md) - List accounts

### [BankAccountTransactions](docs/bankaccounttransactions/README.md)

* [CreateBankTransactions](docs/bankaccounttransactions/createbanktransactions.md) - Create bank transactions
* [GetCreateBankAccountModel](docs/bankaccounttransactions/getcreatebankaccountmodel.md) - List push options for bank account bank transactions
* [ListBankAccountTransactions](docs/bankaccounttransactions/listbankaccounttransactions.md) - List bank transactions for bank account
* [ListBankTransactions](docs/bankaccounttransactions/listbanktransactions.md) - List all bank transactions

### [BankAccounts](docs/bankaccounts/README.md)

* [CreateBankAccount](docs/bankaccounts/createbankaccount.md) - Create bank account
* [GetAllBankAccount](docs/bankaccounts/getallbankaccount.md) - Get bank account
* [GetBankAccount](docs/bankaccounts/getbankaccount.md) - Get bank account
* [GetCreateUpdateBankAccountsModel](docs/bankaccounts/getcreateupdatebankaccountsmodel.md) - Get create/update bank account model
* [ListBankAccounts](docs/bankaccounts/listbankaccounts.md) - List bank accounts
* [UpdateBankAccount](docs/bankaccounts/updatebankaccount.md) - Update bank account

### [BillCreditNotes](docs/billcreditnotes/README.md)

* [CreateBillCreditNote](docs/billcreditnotes/createbillcreditnote.md) - Create bill credit note
* [GetBillCreditNote](docs/billcreditnotes/getbillcreditnote.md) - Get bill credit note
* [GetCreateUpdateBillCreditNotesModel](docs/billcreditnotes/getcreateupdatebillcreditnotesmodel.md) - Get create/update bill credit note model
* [ListBillCreditNotes](docs/billcreditnotes/listbillcreditnotes.md) - List bill credit notes
* [UpdateBillCreditNote](docs/billcreditnotes/updatebillcreditnote.md) - Update bill credit note

### [BillPayments](docs/billpayments/README.md)

* [CreateBillPayment](docs/billpayments/createbillpayment.md) - Create bill payments
* [DeleteBillPayment](docs/billpayments/deletebillpayment.md) - Delete bill payment
* [GetBillPayments](docs/billpayments/getbillpayments.md) - Get bill payment
* [GetCreateBillPaymentsModel](docs/billpayments/getcreatebillpaymentsmodel.md) - Get create bill payment model
* [ListBillPayments](docs/billpayments/listbillpayments.md) - List bill payments

### [Bills](docs/bills/README.md)

* [CreateBill](docs/bills/createbill.md) - Create bill
* [DeleteBill](docs/bills/deletebill.md) - Delete bill
* [DownloadBillAttachment](docs/bills/downloadbillattachment.md) - Download bill attachment
* [GetBill](docs/bills/getbill.md) - Get bill
* [GetBillAttachment](docs/bills/getbillattachment.md) - Get bill attachment
* [GetBillAttachments](docs/bills/getbillattachments.md) - List bill attachments
* [GetCreateUpdateBillsModel](docs/bills/getcreateupdatebillsmodel.md) - Get create/update bill model
* [ListBills](docs/bills/listbills.md) - List bills
* [UpdateBill](docs/bills/updatebill.md) - Update bill
* [UploadBillAttachments](docs/bills/uploadbillattachments.md) - Upload bill attachments

### [CompanyInfo](docs/companyinfo/README.md)

* [GetCompanyInfo](docs/companyinfo/getcompanyinfo.md) - Get company info
* [PostSyncInfo](docs/companyinfo/postsyncinfo.md) - Refresh company info

### [CreditNotes](docs/creditnotes/README.md)

* [CreateCreditNote](docs/creditnotes/createcreditnote.md) - Create credit note
* [GetCreateUpdateCreditNotesModel](docs/creditnotes/getcreateupdatecreditnotesmodel.md) - Get create/update credit note model
* [GetCreditNote](docs/creditnotes/getcreditnote.md) - Get credit note
* [ListCreditNotes](docs/creditnotes/listcreditnotes.md) - List credit notes
* [UpdateCreditNote](docs/creditnotes/updatecreditnote.md) - Update creditNote

### [Customers](docs/customers/README.md)

* [CreateCustomer](docs/customers/createcustomer.md) - Create customer
* [DownloadCustomerAttachment](docs/customers/downloadcustomerattachment.md) - Download customer attachment
* [GetCreateUpdateCustomersModel](docs/customers/getcreateupdatecustomersmodel.md) - Get create/update customer model
* [GetCustomer](docs/customers/getcustomer.md) - Get customer
* [GetCustomerAttachment](docs/customers/getcustomerattachment.md) - Get customer attachment
* [GetCustomerAttachments](docs/customers/getcustomerattachments.md) - List customer attachments
* [GetCustomers](docs/customers/getcustomers.md) - List customers
* [UpdateCustomer](docs/customers/updatecustomer.md) - Update customer

### [DirectCosts](docs/directcosts/README.md)

* [CreateDirectCost](docs/directcosts/createdirectcost.md) - Create direct cost
* [DownloadDirectCostAttachment](docs/directcosts/downloaddirectcostattachment.md) - Download direct cost attachment
* [GetCreateDirectCostsModel](docs/directcosts/getcreatedirectcostsmodel.md) - Get create direct cost model
* [GetDirectCost](docs/directcosts/getdirectcost.md) - Get direct cost
* [GetDirectCostAttachment](docs/directcosts/getdirectcostattachment.md) - Get direct cost attachment
* [GetDirectCosts](docs/directcosts/getdirectcosts.md) - List direct costs
* [ListDirectCostAttachments](docs/directcosts/listdirectcostattachments.md) - List direct cost attachments
* [UploadDirectCostAttachment](docs/directcosts/uploaddirectcostattachment.md) - Upload direct cost attachment

### [DirectIncomes](docs/directincomes/README.md)

* [CreateDirectIncome](docs/directincomes/createdirectincome.md) - Create direct income
* [DownloadDirectIncomeAttachment](docs/directincomes/downloaddirectincomeattachment.md) - Download direct income attachment
* [GetCreateDirectIncomesModel](docs/directincomes/getcreatedirectincomesmodel.md) - Get create direct income model
* [GetDirectIncome](docs/directincomes/getdirectincome.md) - Get direct income
* [GetDirectIncomeAttachment](docs/directincomes/getdirectincomeattachment.md) - Get direct income attachment
* [GetDirectIncomes](docs/directincomes/getdirectincomes.md) - Get direct incomes
* [ListDirectIncomeAttachments](docs/directincomes/listdirectincomeattachments.md) - List direct income attachments
* [UploadDirectIncomeAttachment](docs/directincomes/uploaddirectincomeattachment.md) - Create direct income attachment

### [Financials](docs/financials/README.md)

* [GetBalanceSheet](docs/financials/getbalancesheet.md) - Get balance sheet
* [GetCashFlowStatement](docs/financials/getcashflowstatement.md) - Get cash flow statement
* [GetProfitAndLoss](docs/financials/getprofitandloss.md) - Get profit and loss

### [Invoices](docs/invoices/README.md)

* [DownloadInvoicePdf](docs/invoices/downloadinvoicepdf.md) - Get invoice as PDF
* [CreateInvoice](docs/invoices/createinvoice.md) - Create invoice
* [DeleteInvoice](docs/invoices/deleteinvoice.md) - Delete invoice
* [DownloadInvoiceAttachment](docs/invoices/downloadinvoiceattachment.md) - Download invoice attachment
* [GetCreateUpdateInvoicesModel](docs/invoices/getcreateupdateinvoicesmodel.md) - Get create/update invoice model
* [GetInvoice](docs/invoices/getinvoice.md) - Get invoice
* [GetInvoiceAttachment](docs/invoices/getinvoiceattachment.md) - Get invoice attachment
* [GetInvoiceAttachments](docs/invoices/getinvoiceattachments.md) - Get invoice attachments
* [ListInvoices](docs/invoices/listinvoices.md) - List invoices
* [UpdateInvoice](docs/invoices/updateinvoice.md) - Update invoice
* [UploadInvoiceAttachment](docs/invoices/uploadinvoiceattachment.md) - Push invoice attachment

### [Items](docs/items/README.md)

* [CreateItem](docs/items/createitem.md) - Create item
* [GetCreateItemsModel](docs/items/getcreateitemsmodel.md) - Get create item model
* [GetItem](docs/items/getitem.md) - Get item
* [ListItems](docs/items/listitems.md) - List items

### [JournalEntries](docs/journalentries/README.md)

* [CreateJournalEntry](docs/journalentries/createjournalentry.md) - Create journal entry
* [DeleteJournalEntry](docs/journalentries/deletejournalentry.md) - Delete journal entry
* [GetCreateJournalEntriesModel](docs/journalentries/getcreatejournalentriesmodel.md) - Get create journal entry model
* [GetJournalEntry](docs/journalentries/getjournalentry.md) - Get journal entry
* [ListJournalEntries](docs/journalentries/listjournalentries.md) - List journal entries

### [Journals](docs/journals/README.md)

* [GetCreateJournalsModel](docs/journals/getcreatejournalsmodel.md) - Get create journal model
* [GetJournal](docs/journals/getjournal.md) - Get journal
* [ListJournals](docs/journals/listjournals.md) - List journals
* [PushJournal](docs/journals/pushjournal.md) - Create journal

### [PaymentMethods](docs/paymentmethods/README.md)

* [GetPaymentMethod](docs/paymentmethods/getpaymentmethod.md) - Get payment method
* [ListPaymentMethods](docs/paymentmethods/listpaymentmethods.md) - List all payment methods

### [Payments](docs/payments/README.md)

* [CreatePayment](docs/payments/createpayment.md) - Create payment
* [GetCreatePaymentsModel](docs/payments/getcreatepaymentsmodel.md) - Get create payment model
* [GetPayment](docs/payments/getpayment.md) - Get payment
* [ListPayments](docs/payments/listpayments.md) - List payments

### [PurchaseOrders](docs/purchaseorders/README.md)

* [CreatePurchaseOrder](docs/purchaseorders/createpurchaseorder.md) - Create purchase order
* [GetCreateUpdatePurchaseOrdersModel](docs/purchaseorders/getcreateupdatepurchaseordersmodel.md) - Get create/update purchase order model
* [GetPurchaseOrder](docs/purchaseorders/getpurchaseorder.md) - Get purchase order
* [ListPurchaseOrders](docs/purchaseorders/listpurchaseorders.md) - List purchase orders
* [UpdatePurchaseOrder](docs/purchaseorders/updatepurchaseorder.md) - Update purchase order

### [Reports](docs/reports/README.md)

* [GetAgedCreditorsReport](docs/reports/getagedcreditorsreport.md) - Aged creditors report
* [GetAgedDebtorsReport](docs/reports/getageddebtorsreport.md) - Aged debtors report
* [IsAgedCreditorsReportAvailable](docs/reports/isagedcreditorsreportavailable.md) - Aged creditors report available
* [IsAgedDebtorReportAvailable](docs/reports/isageddebtorreportavailable.md) - Aged debtors report available

### [SalesOrders](docs/salesorders/README.md)

* [GetSalesOrder](docs/salesorders/getsalesorder.md) - Get sales order
* [ListSalesOrders](docs/salesorders/listsalesorders.md) - List sales orders

### [Suppliers](docs/suppliers/README.md)

* [CreateSupplier](docs/suppliers/createsupplier.md) - Create suppliers
* [DownloadSupplierAttachment](docs/suppliers/downloadsupplierattachment.md) - Download supplier attachment
* [GetCreateUpdateSuppliersModel](docs/suppliers/getcreateupdatesuppliersmodel.md) - Get create/update supplier model
* [GetSupplier](docs/suppliers/getsupplier.md) - Get supplier
* [GetSupplierAttachment](docs/suppliers/getsupplierattachment.md) - Get supplier attachment
* [ListSupplierAttachments](docs/suppliers/listsupplierattachments.md) - List supplier attachments
* [ListSuppliers](docs/suppliers/listsuppliers.md) - List suppliers
* [PutSupplier](docs/suppliers/putsupplier.md) - Update supplier

### [TaxRates](docs/taxrates/README.md)

* [GetTaxRate](docs/taxrates/gettaxrate.md) - Get tax rate
* [ListTaxRates](docs/taxrates/listtaxrates.md) - List all tax rates

### [TrackingCategories](docs/trackingcategories/README.md)

* [GetTrackingCategory](docs/trackingcategories/gettrackingcategory.md) - Get tracking categories
* [ListTrackingCategories](docs/trackingcategories/listtrackingcategories.md) - List tracking categories

### [Transfers](docs/transfers/README.md)

* [CreateTransfer](docs/transfers/createtransfer.md) - Create transfer
* [GetCreateTransfersModel](docs/transfers/getcreatetransfersmodel.md) - Get create transfer model
* [GetTransfer](docs/transfers/gettransfer.md) - Get transfer
* [ListTransfers](docs/transfers/listtransfers.md) - List transfers
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
