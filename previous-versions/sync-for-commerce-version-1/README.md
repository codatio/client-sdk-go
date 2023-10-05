# Sync for Commerce version 1

<!-- Start Codat Library Description -->
Embedded accounting integrations for POS and eCommerce platforms.
<!-- End Codat Library Description -->

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1
```
<!-- End SDK Installation -->

## Example Usage
<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountingAccounts.CreateAccountingAccount(ctx, operations.CreateAccountingAccountRequest{
        AccountingAccount: &shared.AccountingAccount{
            Currency: codatsynccommerce.String("USD"),
            CurrentBalance: types.MustNewDecimalFromString("0"),
            Description: codatsynccommerce.String("Invoices the business has issued but has not yet collected payment on."),
            FullyQualifiedCategory: codatsynccommerce.String("Asset.Current"),
            FullyQualifiedName: codatsynccommerce.String("Fixed Asset"),
            ID: codatsynccommerce.String("1b6266d1-1e44-46c5-8eb5-a8f98e03124e"),
            IsBankAccount: codatsynccommerce.Bool(false),
            Metadata: &shared.AccountingAccountMetadata{
                IsDeleted: codatsynccommerce.Bool(false),
            },
            ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Name: codatsynccommerce.String("Accounts Receivable"),
            NominalCode: codatsynccommerce.String("610"),
            SourceModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Status: shared.AccountStatusActive.ToPointer(),
            Type: shared.AccountTypeAsset.ToPointer(),
            ValidDatatypeLinks: []shared.AccountingAccountValidDataTypeLinks{
                shared.AccountingAccountValidDataTypeLinks{
                    Links: []string{
                        "unde",
                    },
                    Property: codatsynccommerce.String("nulla"),
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsynccommerce.Int(544883),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountingCreateAccountResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [AccountingAccounts](docs/sdks/accountingaccounts/README.md)

* [CreateAccountingAccount](docs/sdks/accountingaccounts/README.md#createaccountingaccount) - Create account
* [GetAccountingAccount](docs/sdks/accountingaccounts/README.md#getaccountingaccount) - Get account
* [ListAccountingAccounts](docs/sdks/accountingaccounts/README.md#listaccountingaccounts) - List accounts

### [AccountingBankAccounts](docs/sdks/accountingbankaccounts/README.md)

* [GetAccountingBankAccount](docs/sdks/accountingbankaccounts/README.md#getaccountingbankaccount) - Get bank account
* [ListAccountingBankAccounts](docs/sdks/accountingbankaccounts/README.md#listaccountingbankaccounts) - List bank accounts

### [AccountingCompanyInfo](docs/sdks/accountingcompanyinfo/README.md)

* [GetAccountingCompanyInfo](docs/sdks/accountingcompanyinfo/README.md#getaccountingcompanyinfo) - Get company info
* [Refresh](docs/sdks/accountingcompanyinfo/README.md#refresh) - Refresh company info

### [AccountingCreditNotes](docs/sdks/accountingcreditnotes/README.md)

* [CreateAccountingCreditNote](docs/sdks/accountingcreditnotes/README.md#createaccountingcreditnote) - Create credit note

### [AccountingCustomers](docs/sdks/accountingcustomers/README.md)

* [CreateAccountingCustomer](docs/sdks/accountingcustomers/README.md#createaccountingcustomer) - Create customer

### [AccountingDirectIncomes](docs/sdks/accountingdirectincomes/README.md)

* [CreateAccountingDirectIncome](docs/sdks/accountingdirectincomes/README.md#createaccountingdirectincome) - Create direct income

### [AccountingInvoices](docs/sdks/accountinginvoices/README.md)

* [CreateAccountingInvoice](docs/sdks/accountinginvoices/README.md#createaccountinginvoice) - Create invoice

### [AccountingJournalEntries](docs/sdks/accountingjournalentries/README.md)

* [CreateAccountingJournalEntry](docs/sdks/accountingjournalentries/README.md#createaccountingjournalentry) - Create journal entry

### [AccountingPayments](docs/sdks/accountingpayments/README.md)

* [CreateAccountingPayment](docs/sdks/accountingpayments/README.md#createaccountingpayment) - Create payment

### [CommerceCompanyInfo](docs/sdks/commercecompanyinfo/README.md)

* [GetCommerceCompanyInfo](docs/sdks/commercecompanyinfo/README.md#getcommercecompanyinfo) - Get company info

### [CommerceCustomers](docs/sdks/commercecustomers/README.md)

* [GetCommerceCustomer](docs/sdks/commercecustomers/README.md#getcommercecustomer) - Get customer
* [ListCommerceCustomers](docs/sdks/commercecustomers/README.md#listcommercecustomers) - List customers

### [CommerceLocations](docs/sdks/commercelocations/README.md)

* [GetCommerceLocation](docs/sdks/commercelocations/README.md#getcommercelocation) - Get location
* [ListCommerceLocations](docs/sdks/commercelocations/README.md#listcommercelocations) - List locations

### [CommerceOrders](docs/sdks/commerceorders/README.md)

* [GetCommerceOrder](docs/sdks/commerceorders/README.md#getcommerceorder) - Get order
* [ListCommerceOrders](docs/sdks/commerceorders/README.md#listcommerceorders) - List orders

### [CommercePayments](docs/sdks/commercepayments/README.md)

* [GetCommercePayment](docs/sdks/commercepayments/README.md#getcommercepayment) - Get payment
* [GetMethod](docs/sdks/commercepayments/README.md#getmethod) - Get payment method
* [ListCommercePayments](docs/sdks/commercepayments/README.md#listcommercepayments) - List payments
* [ListMethods](docs/sdks/commercepayments/README.md#listmethods) - List payment methods

### [CommerceProducts](docs/sdks/commerceproducts/README.md)

* [GetCommerceProduct](docs/sdks/commerceproducts/README.md#getcommerceproduct) - Get product
* [ListCommerceProducts](docs/sdks/commerceproducts/README.md#listcommerceproducts) - List products

### [CommerceTransactions](docs/sdks/commercetransactions/README.md)

* [GetCommerceTransaction](docs/sdks/commercetransactions/README.md#getcommercetransaction) - Get transaction
* [ListCommerceTransactions](docs/sdks/commercetransactions/README.md#listcommercetransactions) - List transactions

### [Companies](docs/sdks/companies/README.md)

* [DeleteCompany](docs/sdks/companies/README.md#deletecompany) - Delete a company
* [GetCompany](docs/sdks/companies/README.md#getcompany) - Get company
* [UpdateCompany](docs/sdks/companies/README.md#updatecompany) - Update company

### [CompanyManagement](docs/sdks/companymanagement/README.md)

* [CreateCompany](docs/sdks/companymanagement/README.md#createcompany) - Create Sync for Commerce company
* [CreateConnection](docs/sdks/companymanagement/README.md#createconnection) - Create connection
* [ListCompanies](docs/sdks/companymanagement/README.md#listcompanies) - List companies
* [ListConnections](docs/sdks/companymanagement/README.md#listconnections) - List data connections
* [UpdateConnection](docs/sdks/companymanagement/README.md#updateconnection) - Update data connection

### [Configuration](docs/sdks/configuration/README.md)

* [GetConfiguration](docs/sdks/configuration/README.md#getconfiguration) - Retrieve config preferences set for a company.
* [SetConfiguration](docs/sdks/configuration/README.md#setconfiguration) - Create or update configuration.

### [Connections](docs/sdks/connections/README.md)

* [DeleteConnection](docs/sdks/connections/README.md#deleteconnection) - Delete connection
* [GetConnection](docs/sdks/connections/README.md#getconnection) - Get connection
* [Unlink](docs/sdks/connections/README.md#unlink) - Unlink connection

### [Integrations](docs/sdks/integrations/README.md)

* [GetIntegrationBranding](docs/sdks/integrations/README.md#getintegrationbranding) - Get branding for an integration
* [ListIntegrations](docs/sdks/integrations/README.md#listintegrations) - List information on Codat's supported integrations

### [PushData](docs/sdks/pushdata/README.md)

* [GetOperation](docs/sdks/pushdata/README.md#getoperation) - Get push operation
* [ListOperations](docs/sdks/pushdata/README.md#listoperations) - List push operations

### [RefreshData](docs/sdks/refreshdata/README.md)

* [All](docs/sdks/refreshdata/README.md#all) - Refresh all data
* [ByDataType](docs/sdks/refreshdata/README.md#bydatatype) - Refresh data type
* [GetCompanyDataStatus](docs/sdks/refreshdata/README.md#getcompanydatastatus) - Get data status
* [GetPullOperation](docs/sdks/refreshdata/README.md#getpulloperation) - Get pull operation
* [ListPullOperations](docs/sdks/refreshdata/README.md#listpulloperations) - List pull operations

### [Sync](docs/sdks/sync/README.md)

* [GetSyncStatus](docs/sdks/sync/README.md#getsyncstatus) - Get status for a company's syncs
* [RequestSync](docs/sdks/sync/README.md#requestsync) - Sync new
* [RequestSyncForDateRange](docs/sdks/sync/README.md#requestsyncfordaterange) - Sync range

### [SyncFlowPreferences](docs/sdks/syncflowpreferences/README.md)

* [GetConfigTextSyncFlow](docs/sdks/syncflowpreferences/README.md#getconfigtextsyncflow) - Retrieve preferences for text fields on Sync Flow
* [GetSyncFlowURL](docs/sdks/syncflowpreferences/README.md#getsyncflowurl) - Retrieve sync flow url
* [GetVisibleAccounts](docs/sdks/syncflowpreferences/README.md#getvisibleaccounts) - List visible accounts
* [UpdateConfigTextSyncFlow](docs/sdks/syncflowpreferences/README.md#updateconfigtextsyncflow) - Update preferences for text fields on sync flow
* [UpdateVisibleAccountsSyncFlow](docs/sdks/syncflowpreferences/README.md#updatevisibleaccountssyncflow) - Update the visible accounts on Sync Flow
<!-- End SDK Available Operations -->
### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
