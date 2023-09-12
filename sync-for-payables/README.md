# Sync for Payables
    


## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/sync-for-payables
```## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/sync-for-payables
```<!-- Start SDK Installation -->

<!-- End SDK Installation -->

## Example Usage


```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payables"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/operations"
)

func main() {
    s := codatsyncpayables.New(
        codatsyncpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.Create(ctx, operations.CreateAccountRequest{
        Account: &shared.Account{
            Currency: codatsyncpayables.String("USD"),
            CurrentBalance: types.MustNewDecimalFromString("0"),
            Description: codatsyncpayables.String("Invoices the business has issued but has not yet collected payment on."),
            FullyQualifiedCategory: codatsyncpayables.String("Asset.Current"),
            FullyQualifiedName: codatsyncpayables.String("Fixed Asset"),
            ID: codatsyncpayables.String("1b6266d1-1e44-46c5-8eb5-a8f98e03124e"),
            IsBankAccount: codatsyncpayables.Bool(false),
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncpayables.Bool(false),
            },
            ModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Name: codatsyncpayables.String("Accounts Receivable"),
            NominalCode: codatsyncpayables.String("610"),
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.AccountStatusActive.ToPointer(),
            Type: shared.AccountTypeAsset.ToPointer(),
            ValidDatatypeLinks: []shared.AccountValidDataTypeLinks{
                shared.AccountValidDataTypeLinks{
                    Links: []string{
                        "unde",
                    },
                    Property: codatsyncpayables.String("nulla"),
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncpayables.Int(544883),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateAccountResponse != nil {
        // handle response
    }
}
```

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-payables"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/operations"
)

func main() {
    s := codatsyncpayables.New(
        codatsyncpayables.WithSecurity(shared.Security{
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

    if res.CompanyInformation != nil {
        // handle response
    }
}
```<!-- Start SDK Example Usage -->

<!-- End SDK Example Usage -->

## Available Resources and Operations


### [Accounts](docs/sdks/accounts/README.md)

* [Create](docs/sdks/accounts/README.md#create) - Create account
* [Get](docs/sdks/accounts/README.md#get) - Get account
* [GetCreateModel](docs/sdks/accounts/README.md#getcreatemodel) - Get create account model
* [List](docs/sdks/accounts/README.md#list) - List accounts

### [BillCreditNotes](docs/sdks/billcreditnotes/README.md)

* [Create](docs/sdks/billcreditnotes/README.md#create) - Create bill credit note
* [Get](docs/sdks/billcreditnotes/README.md#get) - Get bill credit note
* [GetCreateUpdateModel](docs/sdks/billcreditnotes/README.md#getcreateupdatemodel) - Get create/update bill credit note model
* [List](docs/sdks/billcreditnotes/README.md#list) - List bill credit notes
* [Update](docs/sdks/billcreditnotes/README.md#update) - Update bill credit note

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

### [Companies](docs/sdks/companies/README.md)

* [Create](docs/sdks/companies/README.md#create) - Create company
* [Delete](docs/sdks/companies/README.md#delete) - Delete a company
* [Get](docs/sdks/companies/README.md#get) - Get company
* [List](docs/sdks/companies/README.md#list) - List companies
* [Update](docs/sdks/companies/README.md#update) - Update company

### [Connections](docs/sdks/connections/README.md)

* [Create](docs/sdks/connections/README.md#create) - Create connection
* [Delete](docs/sdks/connections/README.md#delete) - Delete connection
* [Get](docs/sdks/connections/README.md#get) - Get connection
* [List](docs/sdks/connections/README.md#list) - List connections
* [Unlink](docs/sdks/connections/README.md#unlink) - Unlink connection

### [JournalEntries](docs/sdks/journalentries/README.md)

* [Create](docs/sdks/journalentries/README.md#create) - Create journal entry
* [GetCreateModel](docs/sdks/journalentries/README.md#getcreatemodel) - Get create journal entry model

### [Journals](docs/sdks/journals/README.md)

* [Create](docs/sdks/journals/README.md#create) - Create journal
* [Get](docs/sdks/journals/README.md#get) - Get journal
* [GetCreateModel](docs/sdks/journals/README.md#getcreatemodel) - Get create journal model
* [List](docs/sdks/journals/README.md#list) - List journals

### [ManageData](docs/sdks/managedata/README.md)

* [Get](docs/sdks/managedata/README.md#get) - Get data status
* [GetPullOperation](docs/sdks/managedata/README.md#getpulloperation) - Get pull operation
* [ListPullOperations](docs/sdks/managedata/README.md#listpulloperations) - List pull operations
* [RefreshAllDataTypes](docs/sdks/managedata/README.md#refreshalldatatypes) - Refresh all data
* [RefreshDataType](docs/sdks/managedata/README.md#refreshdatatype) - Refresh data type

### [PaymentMethods](docs/sdks/paymentmethods/README.md)

* [Get](docs/sdks/paymentmethods/README.md#get) - Get payment method
* [List](docs/sdks/paymentmethods/README.md#list) - List payment methods

### [PushOperations](docs/sdks/pushoperations/README.md)

* [Get](docs/sdks/pushoperations/README.md#get) - Get push operation
* [List](docs/sdks/pushoperations/README.md#list) - List push operations

### [Suppliers](docs/sdks/suppliers/README.md)

* [Create](docs/sdks/suppliers/README.md#create) - Create supplier
* [Get](docs/sdks/suppliers/README.md#get) - Get supplier
* [GetCreateUpdateModel](docs/sdks/suppliers/README.md#getcreateupdatemodel) - Get create/update supplier model
* [List](docs/sdks/suppliers/README.md#list) - List suppliers
* [Update](docs/sdks/suppliers/README.md#update) - Update supplier

### [TaxRates](docs/sdks/taxrates/README.md)

* [Get](docs/sdks/taxrates/README.md#get) - Get tax rate
* [List](docs/sdks/taxrates/README.md#list) - List all tax rates

### [TrackingCategories](docs/sdks/trackingcategories/README.md)

* [Get](docs/sdks/trackingcategories/README.md#get) - Get tracking categories
* [List](docs/sdks/trackingcategories/README.md#list) - List tracking categories## Available Resources and Operations

### [CodatSyncPayables SDK](docs/sdks/codatsyncpayables/README.md)

* [GetAccountingProfile](docs/sdks/codatsyncpayables/README.md#getaccountingprofile) - Get company accounting profile

### [Accounts](docs/sdks/accounts/README.md)

* [Create](docs/sdks/accounts/README.md#create) - Create account
* [Get](docs/sdks/accounts/README.md#get) - Get account
* [GetCreateModel](docs/sdks/accounts/README.md#getcreatemodel) - Get create account model
* [List](docs/sdks/accounts/README.md#list) - List accounts

### [BillCreditNotes](docs/sdks/billcreditnotes/README.md)

* [Create](docs/sdks/billcreditnotes/README.md#create) - Create bill credit note
* [Get](docs/sdks/billcreditnotes/README.md#get) - Get bill credit note
* [GetCreateUpdateModel](docs/sdks/billcreditnotes/README.md#getcreateupdatemodel) - Get create/update bill credit note model
* [List](docs/sdks/billcreditnotes/README.md#list) - List bill credit notes
* [Update](docs/sdks/billcreditnotes/README.md#update) - Update bill credit note

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

### [Companies](docs/sdks/companies/README.md)

* [Create](docs/sdks/companies/README.md#create) - Create company
* [Delete](docs/sdks/companies/README.md#delete) - Delete a company
* [Get](docs/sdks/companies/README.md#get) - Get company
* [List](docs/sdks/companies/README.md#list) - List companies
* [Update](docs/sdks/companies/README.md#update) - Update company

### [Connections](docs/sdks/connections/README.md)

* [Create](docs/sdks/connections/README.md#create) - Create connection
* [Delete](docs/sdks/connections/README.md#delete) - Delete connection
* [Get](docs/sdks/connections/README.md#get) - Get connection
* [List](docs/sdks/connections/README.md#list) - List connections
* [Unlink](docs/sdks/connections/README.md#unlink) - Unlink connection

### [JournalEntries](docs/sdks/journalentries/README.md)

* [Create](docs/sdks/journalentries/README.md#create) - Create journal entry
* [GetCreateModel](docs/sdks/journalentries/README.md#getcreatemodel) - Get create journal entry model

### [Journals](docs/sdks/journals/README.md)

* [Create](docs/sdks/journals/README.md#create) - Create journal
* [Get](docs/sdks/journals/README.md#get) - Get journal
* [GetCreateModel](docs/sdks/journals/README.md#getcreatemodel) - Get create journal model
* [List](docs/sdks/journals/README.md#list) - List journals

### [ManageData](docs/sdks/managedata/README.md)

* [Get](docs/sdks/managedata/README.md#get) - Get data status
* [GetPullOperation](docs/sdks/managedata/README.md#getpulloperation) - Get pull operation
* [ListPullOperations](docs/sdks/managedata/README.md#listpulloperations) - List pull operations
* [RefreshAllDataTypes](docs/sdks/managedata/README.md#refreshalldatatypes) - Refresh all data
* [RefreshDataType](docs/sdks/managedata/README.md#refreshdatatype) - Refresh data type

### [PaymentMethods](docs/sdks/paymentmethods/README.md)

* [Get](docs/sdks/paymentmethods/README.md#get) - Get payment method
* [List](docs/sdks/paymentmethods/README.md#list) - List payment methods

### [PushOperations](docs/sdks/pushoperations/README.md)

* [Get](docs/sdks/pushoperations/README.md#get) - Get push operation
* [List](docs/sdks/pushoperations/README.md#list) - List push operations

### [Suppliers](docs/sdks/suppliers/README.md)

* [Create](docs/sdks/suppliers/README.md#create) - Create supplier
* [Get](docs/sdks/suppliers/README.md#get) - Get supplier
* [GetCreateUpdateModel](docs/sdks/suppliers/README.md#getcreateupdatemodel) - Get create/update supplier model
* [List](docs/sdks/suppliers/README.md#list) - List suppliers
* [Update](docs/sdks/suppliers/README.md#update) - Update supplier

### [TaxRates](docs/sdks/taxrates/README.md)

* [Get](docs/sdks/taxrates/README.md#get) - Get tax rate
* [List](docs/sdks/taxrates/README.md#list) - List all tax rates

### [TrackingCategories](docs/sdks/trackingcategories/README.md)

* [Get](docs/sdks/trackingcategories/README.md#get) - Get tracking categories
* [List](docs/sdks/trackingcategories/README.md#list) - List tracking categories<!-- Start SDK Available Operations -->

<!-- End SDK Available Operations -->
### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
