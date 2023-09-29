# Customers
(*Customers*)

## Overview

Customers

### Available Operations

* [Create](#create) - Create customer
* [Get](#get) - Get customer
* [List](#list) - List customers
* [Update](#update) - Update customer

## Create

The *Create customer* endpoint creates a new [customer](https://docs.codat.io/sync-for-expenses-api#/schemas/Customer) for a given company's connection.

[Customers](https://docs.codat.io/sync-for-expenses-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update customer model](https://docs.codat.io/sync-for-expenses-api#/operations/get-create-update-customers-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v2"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/operations"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.Create(ctx, operations.CreateCustomerRequest{
        Customer: &shared.Customer{
            Addresses: []shared.Items{
                shared.Items{
                    City: syncforexpenses.String("Jenafurt"),
                    Country: syncforexpenses.String("Sweden"),
                    Line1: syncforexpenses.String("innovative blue"),
                    Line2: syncforexpenses.String("grey technology East"),
                    PostalCode: syncforexpenses.String("30778"),
                    Region: syncforexpenses.String("quantify Polestar mobile"),
                    Type: shared.AccountingAddressTypeBilling,
                },
            },
            ContactName: syncforexpenses.String("Durham after"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Items{
                        City: syncforexpenses.String("Darenberg"),
                        Country: syncforexpenses.String("Cote d'Ivoire"),
                        Line1: syncforexpenses.String("Buckinghamshire functionalities Grocery"),
                        Line2: syncforexpenses.String("Metal"),
                        PostalCode: syncforexpenses.String("61380"),
                        Region: syncforexpenses.String("Interactions Senior Mouse"),
                        Type: shared.AccountingAddressTypeUnknown,
                    },
                    Email: syncforexpenses.String("Judd27@hotmail.com"),
                    ModifiedDate: syncforexpenses.String("2022-10-23T00:00:00.000Z"),
                    Name: syncforexpenses.String("transmit likewise"),
                    Phone: []shared.ContactPhone{
                        shared.ContactPhone{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeFax,
                        },
                    },
                    Status: shared.CustomerStatusActive,
                },
            },
            CustomerName: syncforexpenses.String("Rubber silver Indiana"),
            DefaultCurrency: syncforexpenses.String("EUR"),
            EmailAddress: syncforexpenses.String("Thea_Ritchie76@hotmail.com"),
            ID: syncforexpenses.String("<ID>"),
            Metadata: &shared.Metadata{
                IsDeleted: syncforexpenses.Bool(false),
            },
            ModifiedDate: syncforexpenses.String("2022-10-23T00:00:00.000Z"),
            Phone: syncforexpenses.String("948.595.2034"),
            RegistrationNumber: syncforexpenses.String("digital"),
            SourceModifiedDate: syncforexpenses.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusUnknown,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "recusandae": map[string]interface{}{
                        "maiores": "Mongolia",
                    },
                },
            },
            TaxNumber: syncforexpenses.String("discrete"),
        },
        AllowSyncOnPushComplete: syncforexpenses.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: syncforexpenses.Int(522311),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateCustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateCustomerRequest](../../models/operations/createcustomerrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.CreateCustomerResponse](../../models/operations/createcustomerresponse.md), error**


## Get

The *Get customer* endpoint returns a single customer for a given customerId.

[Customers](https://docs.codat.io/sync-for-expenses-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support getting a specific customer.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-expenses-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v2"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/operations"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.Get(ctx, operations.GetCustomerRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CustomerID: "Northeast Hatchback Kia",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Customer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetCustomerRequest](../../models/operations/getcustomerrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |


### Response

**[*operations.GetCustomerResponse](../../models/operations/getcustomerresponse.md), error**


## List

The *List customers* endpoint returns a list of [customers](https://docs.codat.io/sync-for-expenses-api#/schemas/Customer) for a given company's connection.

[Customers](https://docs.codat.io/sync-for-expenses-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-expenses-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v2"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/operations"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.List(ctx, operations.ListCustomersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: syncforexpenses.String("-modifiedDate"),
        Page: syncforexpenses.Int(1),
        PageSize: syncforexpenses.Int(100),
        Query: syncforexpenses.String("Northeast Metal Canada"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Customers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListCustomersRequest](../../models/operations/listcustomersrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.ListCustomersResponse](../../models/operations/listcustomersresponse.md), error**


## Update

The *Update customer* endpoint updates an existing [customer](https://docs.codat.io/sync-for-expenses-api#/schemas/Customer) for a given company's connection.

[Customers](https://docs.codat.io/sync-for-expenses-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update customer model](https://docs.codat.io/sync-for-expenses-api#/operations/get-create-update-customers-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"context"
	"log"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v2"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/operations"
)

func main() {
    s := syncforexpenses.New(
        syncforexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.Update(ctx, operations.UpdateCustomerRequest{
        Customer: &shared.Customer{
            Addresses: []shared.Items{
                shared.Items{
                    City: syncforexpenses.String("Ann Arbor"),
                    Country: syncforexpenses.String("Montserrat"),
                    Line1: syncforexpenses.String("Reactive"),
                    Line2: syncforexpenses.String("Metal cheater Islands"),
                    PostalCode: syncforexpenses.String("43372"),
                    Region: syncforexpenses.String("Carolina syndicate"),
                    Type: shared.AccountingAddressTypeBilling,
                },
            },
            ContactName: syncforexpenses.String("East"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Items{
                        City: syncforexpenses.String("Olenfurt"),
                        Country: syncforexpenses.String("Paraguay"),
                        Line1: syncforexpenses.String("Home users Sharable"),
                        Line2: syncforexpenses.String("Lev Wooden"),
                        PostalCode: syncforexpenses.String("36848"),
                        Region: syncforexpenses.String("brightly"),
                        Type: shared.AccountingAddressTypeDelivery,
                    },
                    Email: syncforexpenses.String("Josie49@yahoo.com"),
                    ModifiedDate: syncforexpenses.String("2022-10-23T00:00:00.000Z"),
                    Name: syncforexpenses.String("possimus navigating Diesel"),
                    Phone: []shared.ContactPhone{
                        shared.ContactPhone{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
            },
            CustomerName: syncforexpenses.String("Reactive Global Northeast"),
            DefaultCurrency: syncforexpenses.String("USD"),
            EmailAddress: syncforexpenses.String("Abe.Bogan@hotmail.com"),
            ID: syncforexpenses.String("<ID>"),
            Metadata: &shared.Metadata{
                IsDeleted: syncforexpenses.Bool(false),
            },
            ModifiedDate: syncforexpenses.String("2022-10-23T00:00:00.000Z"),
            Phone: syncforexpenses.String("(883) 732-4217 x6499"),
            RegistrationNumber: syncforexpenses.String("redundant ew"),
            SourceModifiedDate: syncforexpenses.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusActive,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "asperiores": map[string]interface{}{
                        "quibusdam": "Omnigender",
                    },
                },
            },
            TaxNumber: syncforexpenses.String("Volkswagen Specialist Bacon"),
        },
        AllowSyncOnPushComplete: syncforexpenses.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "Copper port East",
        ForceUpdate: syncforexpenses.Bool(false),
        TimeoutInMinutes: syncforexpenses.Int(373959),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateCustomerRequest](../../models/operations/updatecustomerrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.UpdateCustomerResponse](../../models/operations/updatecustomerresponse.md), error**

