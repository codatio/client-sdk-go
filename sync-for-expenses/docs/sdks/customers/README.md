# Customers

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
                    City: syncforexpenses.String("Fort Manuelachester"),
                    Country: syncforexpenses.String("Oman"),
                    Line1: syncforexpenses.String("dolores"),
                    Line2: syncforexpenses.String("dolorem"),
                    PostalCode: syncforexpenses.String("17363"),
                    Region: syncforexpenses.String("minima"),
                    Type: shared.AccountingAddressTypeBilling,
                },
            },
            ContactName: syncforexpenses.String("accusantium"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Items{
                        City: syncforexpenses.String("Lorenzaworth"),
                        Country: syncforexpenses.String("Uzbekistan"),
                        Line1: syncforexpenses.String("architecto"),
                        Line2: syncforexpenses.String("mollitia"),
                        PostalCode: syncforexpenses.String("61965"),
                        Region: syncforexpenses.String("numquam"),
                        Type: shared.AccountingAddressTypeBilling,
                    },
                    Email: syncforexpenses.String("Jarred.Frami@yahoo.com"),
                    ModifiedDate: syncforexpenses.String("2022-10-23T00:00:00.000Z"),
                    Name: syncforexpenses.String("Kayla O'Kon"),
                    Phone: []shared.ContactPhone{
                        shared.ContactPhone{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                    },
                    Status: shared.CustomerStatusArchived,
                },
            },
            CustomerName: syncforexpenses.String("ipsam"),
            DefaultCurrency: syncforexpenses.String("USD"),
            EmailAddress: syncforexpenses.String("possimus"),
            ID: syncforexpenses.String("019da1ff-e78f-4097-b007-4f15471b5e6e"),
            Metadata: &shared.Metadata{
                IsDeleted: syncforexpenses.Bool(false),
            },
            ModifiedDate: syncforexpenses.String("2022-10-23T00:00:00.000Z"),
            Phone: syncforexpenses.String("755.825.5909"),
            RegistrationNumber: syncforexpenses.String("sint"),
            SourceModifiedDate: syncforexpenses.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusArchived,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "incidunt": map[string]interface{}{
                        "enim": "consequatur",
                    },
                },
            },
            TaxNumber: syncforexpenses.String("est"),
        },
        AllowSyncOnPushComplete: syncforexpenses.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: syncforexpenses.Int(842342),
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
        CustomerID: "explicabo",
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
        Query: syncforexpenses.String("deserunt"),
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
                    City: syncforexpenses.String("Spencerboro"),
                    Country: syncforexpenses.String("Eritrea"),
                    Line1: syncforexpenses.String("qui"),
                    Line2: syncforexpenses.String("aliquid"),
                    PostalCode: syncforexpenses.String("50183-0165"),
                    Region: syncforexpenses.String("tempora"),
                    Type: shared.AccountingAddressTypeDelivery,
                },
            },
            ContactName: syncforexpenses.String("tempore"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Items{
                        City: syncforexpenses.String("Fort Horacio"),
                        Country: syncforexpenses.String("Ecuador"),
                        Line1: syncforexpenses.String("eligendi"),
                        Line2: syncforexpenses.String("sint"),
                        PostalCode: syncforexpenses.String("58562"),
                        Region: syncforexpenses.String("debitis"),
                        Type: shared.AccountingAddressTypeDelivery,
                    },
                    Email: syncforexpenses.String("Isadore_Kirlin69@hotmail.com"),
                    ModifiedDate: syncforexpenses.String("2022-10-23T00:00:00.000Z"),
                    Name: syncforexpenses.String("Blanca Schulist"),
                    Phone: []shared.ContactPhone{
                        shared.ContactPhone{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeUnknown,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
            },
            CustomerName: syncforexpenses.String("occaecati"),
            DefaultCurrency: syncforexpenses.String("GBP"),
            EmailAddress: syncforexpenses.String("accusamus"),
            ID: syncforexpenses.String("fb9ba88f-3a66-4997-874b-a4469b6e2141"),
            Metadata: &shared.Metadata{
                IsDeleted: syncforexpenses.Bool(false),
            },
            ModifiedDate: syncforexpenses.String("2022-10-23T00:00:00.000Z"),
            Phone: syncforexpenses.String("(655) 269-6342 x813"),
            RegistrationNumber: syncforexpenses.String("quasi"),
            SourceModifiedDate: syncforexpenses.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusArchived,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "debitis": map[string]interface{}{
                        "eius": "maxime",
                    },
                },
            },
            TaxNumber: syncforexpenses.String("deleniti"),
        },
        AllowSyncOnPushComplete: syncforexpenses.Bool(false),
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "facilis",
        ForceUpdate: syncforexpenses.Bool(false),
        TimeoutInMinutes: syncforexpenses.Int(447926),
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

