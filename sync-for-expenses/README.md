# Sync for Expenses
    
﻿Embedded accounting integrations for corporate card providers.

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
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.Create(ctx, operations.CreateAccountRequest{
        Account: &shared.Account{
            Currency: codatsyncexpenses.String("USD"),
            CurrentBalance: types.MustNewDecimalFromString("0"),
            Description: codatsyncexpenses.String("Invoices the business has issued but has not yet collected payment on."),
            FullyQualifiedCategory: codatsyncexpenses.String("Asset.Current"),
            FullyQualifiedName: codatsyncexpenses.String("Fixed Asset"),
            ID: codatsyncexpenses.String("1b6266d1-1e44-46c5-8eb5-a8f98e03124e"),
            IsBankAccount: codatsyncexpenses.Bool(false),
            Metadata: &shared.AccountMetadata{
                IsDeleted: codatsyncexpenses.Bool(false),
            },
            ModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
            Name: codatsyncexpenses.String("Accounts Receivable"),
            NominalCode: codatsyncexpenses.String("610"),
            SourceModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
            Status: shared.AccountStatusActive.ToPointer(),
            Type: shared.AccountTypeAsset.ToPointer(),
            ValidDatatypeLinks: []shared.AccountValidDataTypeLinks{
                shared.AccountValidDataTypeLinks{
                    Links: []string{
                        "unde",
                    },
                    Property: codatsyncexpenses.String("nulla"),
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncexpenses.Int(544883),
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
* [CreatePartnerExpenseConnection](docs/sdks/connections/README.md#createpartnerexpenseconnection) - Create Partner Expense connection
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

* [Get](docs/sdks/sync/README.md#get) - Get Sync status
* [GetLastSuccessfulSync](docs/sdks/sync/README.md#getlastsuccessfulsync) - Last successful sync
* [GetLatestSync](docs/sdks/sync/README.md#getlatestsync) - Latest sync status
* [InitiateSync](docs/sdks/sync/README.md#initiatesync) - Initiate sync
* [List](docs/sdks/sync/README.md#list) - List sync statuses

### [TransactionStatus](docs/sdks/transactionstatus/README.md)

* [Get](docs/sdks/transactionstatus/README.md#get) - Get Sync Transaction
* [List](docs/sdks/transactionstatus/README.md#list) - List sync transactions
<!-- End SDK Available Operations -->
### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
