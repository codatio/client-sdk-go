# Sync for Expenses

<!-- Start Codat Library Description -->
Embedded accounting integrations for corporate card providers.
<!-- End Codat Library Description  -->

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/sync-for-expenses
```
<!-- End SDK Installation -->

## Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v3"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/types"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.Create(ctx, operations.CreateAccountRequest{
        Account: &shared.Account{
            Currency: syncforexpenses.String("USD"),
            CurrentBalance: types.MustNewDecimalFromString("0"),
            Description: syncforexpenses.String("Invoices the business has issued but has not yet collected payment on."),
            FullyQualifiedCategory: syncforexpenses.String("Asset.Current"),
            FullyQualifiedName: syncforexpenses.String("Cash On Hand"),
            ID: syncforexpenses.String("1b6266d1-1e44-46c5-8eb5-a8f98e03124e"),
            IsBankAccount: syncforexpenses.Bool(false),
            Metadata: &shared.AccountMetadata{
                IsDeleted: syncforexpenses.Bool(false),
            },
            ModifiedDate: syncforexpenses.String("2022-10-23T00:00:00.000Z"),
            Name: syncforexpenses.String("Accounts Receivable"),
            NominalCode: syncforexpenses.String("610"),
            SourceModifiedDate: syncforexpenses.String("2022-10-23T00:00:00.000Z"),
            Status: shared.AccountStatusActive.ToPointer(),
            Type: shared.AccountTypeAsset.ToPointer(),
            ValidDatatypeLinks: []shared.AccountValidDataTypeLinks{
                shared.AccountValidDataTypeLinks{
                    Links: []string{
                        "Money",
                    },
                    Property: syncforexpenses.String("Cambridgeshire grey technology"),
                },
            },
        },
        AllowSyncOnPushComplete: syncforexpenses.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: syncforexpenses.Int(86),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateAccountResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Accounts](docs/sdks/accounts/README.md)

* [Create](docs/sdks/accounts/README.md#create) - Create account
* [GetCreateModel](docs/sdks/accounts/README.md#getcreatemodel) - Get create account model

### [Companies](docs/sdks/companies/README.md)

* [Create](docs/sdks/companies/README.md#create) - Create company
* [Delete](docs/sdks/companies/README.md#delete) - Delete a company
* [Get](docs/sdks/companies/README.md#get) - Get company
* [List](docs/sdks/companies/README.md#list) - List companies
* [Update](docs/sdks/companies/README.md#update) - Update company

### [Configuration](docs/sdks/configuration/README.md)

* [Get](docs/sdks/configuration/README.md#get) - Get company configuration
* [GetMappingOptions](docs/sdks/configuration/README.md#getmappingoptions) - Mapping options
* [Set](docs/sdks/configuration/README.md#set) - Set company configuration

### [Connections](docs/sdks/connections/README.md)

* [Create](docs/sdks/connections/README.md#create) - Create connection
* [CreatePartnerExpenseConnection](docs/sdks/connections/README.md#createpartnerexpenseconnection) - Create partner expense connection
* [Delete](docs/sdks/connections/README.md#delete) - Delete connection
* [Get](docs/sdks/connections/README.md#get) - Get connection
* [List](docs/sdks/connections/README.md#list) - List connections
* [Unlink](docs/sdks/connections/README.md#unlink) - Unlink connection

### [Customers](docs/sdks/customers/README.md)

* [Create](docs/sdks/customers/README.md#create) - Create customer
* [Get](docs/sdks/customers/README.md#get) - Get customer
* [List](docs/sdks/customers/README.md#list) - List customers
* [Update](docs/sdks/customers/README.md#update) - Update customer

### [Expenses](docs/sdks/expenses/README.md)

* [Create](docs/sdks/expenses/README.md#create) - Create expense transaction
* [Update](docs/sdks/expenses/README.md#update) - Update expense-transactions
* [UploadAttachment](docs/sdks/expenses/README.md#uploadattachment) - Upload attachment

### [ManageData](docs/sdks/managedata/README.md)

* [Get](docs/sdks/managedata/README.md#get) - Get data status
* [GetPullOperation](docs/sdks/managedata/README.md#getpulloperation) - Get pull operation
* [ListPullOperations](docs/sdks/managedata/README.md#listpulloperations) - List pull operations
* [RefreshAllDataTypes](docs/sdks/managedata/README.md#refreshalldatatypes) - Refresh all data
* [RefreshDataType](docs/sdks/managedata/README.md#refreshdatatype) - Refresh data type

### [PushOperations](docs/sdks/pushoperations/README.md)

* [Get](docs/sdks/pushoperations/README.md#get) - Get push operation
* [List](docs/sdks/pushoperations/README.md#list) - List push operations

### [Suppliers](docs/sdks/suppliers/README.md)

* [Create](docs/sdks/suppliers/README.md#create) - Create supplier
* [Get](docs/sdks/suppliers/README.md#get) - Get supplier
* [List](docs/sdks/suppliers/README.md#list) - List suppliers
* [Update](docs/sdks/suppliers/README.md#update) - Update supplier

### [Sync](docs/sdks/sync/README.md)

* [Get](docs/sdks/sync/README.md#get) - Get sync status
* [GetLastSuccessfulSync](docs/sdks/sync/README.md#getlastsuccessfulsync) - Last successful sync
* [GetLatestSync](docs/sdks/sync/README.md#getlatestsync) - Latest sync status
* [InitiateSync](docs/sdks/sync/README.md#initiatesync) - Initiate sync
* [List](docs/sdks/sync/README.md#list) - List sync statuses

### [TransactionStatus](docs/sdks/transactionstatus/README.md)

* [Get](docs/sdks/transactionstatus/README.md#get) - Get sync transaction
* [List](docs/sdks/transactionstatus/README.md#list) - List sync transactions
<!-- End SDK Available Operations -->

<!-- Start Dev Containers -->



<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:


<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->


### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
