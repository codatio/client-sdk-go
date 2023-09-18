# Customers

## Overview

Customers

### Available Operations

* [Create](#create) - Create customer
* [Get](#get) - Get customer
* [List](#list) - List customers
* [Update](#update) - Update customer

## Create

The *Create customer* endpoint creates a new [customer](https://docs.codat.io/accounting-api#/schemas/Customer) for a given company's connection.

[Customers](https://docs.codat.io/accounting-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update customer model](https://docs.codat.io/accounting-api#/operations/get-create-update-customers-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support creating an account.


### Example Usage

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
    res, err := s.Customers.Create(ctx, operations.CreateCustomerRequest{
        Customer: &shared.Customer{
            Addresses: []shared.Items{
                shared.Items{
                    City: codatsyncexpenses.String("New Humberto"),
                    Country: codatsyncexpenses.String("Trinidad and Tobago"),
                    Line1: codatsyncexpenses.String("quidem"),
                    Line2: codatsyncexpenses.String("architecto"),
                    PostalCode: codatsyncexpenses.String("96661"),
                    Region: codatsyncexpenses.String("dolorem"),
                    Type: shared.AccountingAddressTypeBilling,
                },
            },
            ContactName: codatsyncexpenses.String("explicabo"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Items{
                        City: codatsyncexpenses.String("Halvorsonstead"),
                        Country: codatsyncexpenses.String("Guinea"),
                        Line1: codatsyncexpenses.String("minima"),
                        Line2: codatsyncexpenses.String("excepturi"),
                        PostalCode: codatsyncexpenses.String("46991"),
                        Region: codatsyncexpenses.String("mollitia"),
                        Type: shared.AccountingAddressTypeUnknown,
                    },
                    Email: codatsyncexpenses.String("Caroline_Ziemann@yahoo.com"),
                    ModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
                    Name: codatsyncexpenses.String("Claudia Krajcik"),
                    Phone: []shared.ContactPhone{
                        shared.ContactPhone{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
            },
            CustomerName: codatsyncexpenses.String("laborum"),
            DefaultCurrency: codatsyncexpenses.String("USD"),
            EmailAddress: codatsyncexpenses.String("enim"),
            ID: codatsyncexpenses.String("2c3f5ad0-19da-41ff-a78f-097b0074f154"),
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncexpenses.Bool(false),
            },
            ModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
            Phone: codatsyncexpenses.String("(738) 590-2655"),
            RegistrationNumber: codatsyncexpenses.String("pariatur"),
            SourceModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusActive,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "rem": map[string]interface{}{
                        "voluptates": "quasi",
                    },
                },
            },
            TaxNumber: codatsyncexpenses.String("repudiandae"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncexpenses.Int(575947),
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

[Customers](https://docs.codat.io/accounting-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support getting a specific customer.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-expenses-api#/operations/refresh-company-data).


### Example Usage

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
    res, err := s.Customers.Get(ctx, operations.GetCustomerRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CustomerID: "veritatis",
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

The *List customers* endpoint returns a list of [customers](https://docs.codat.io/accounting-api#/schemas/Customer) for a given company's connection.

[Customers](https://docs.codat.io/accounting-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-expenses-api#/operations/refresh-company-data).
    

### Example Usage

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
    res, err := s.Customers.List(ctx, operations.ListCustomersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codatsyncexpenses.String("-modifiedDate"),
        Page: codatsyncexpenses.Int(1),
        PageSize: codatsyncexpenses.Int(100),
        Query: codatsyncexpenses.String("itaque"),
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

The *Update customer* endpoint updates an existing [customer](https://docs.codat.io/accounting-api#/schemas/Customer) for a given company's connection.

[Customers](https://docs.codat.io/accounting-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update customer model](https://docs.codat.io/accounting-api#/operations/get-create-update-customers-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support creating an account.


### Example Usage

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
    res, err := s.Customers.Update(ctx, operations.UpdateCustomerRequest{
        Customer: &shared.Customer{
            Addresses: []shared.Items{
                shared.Items{
                    City: codatsyncexpenses.String("West Adele"),
                    Country: codatsyncexpenses.String("Norway"),
                    Line1: codatsyncexpenses.String("quibusdam"),
                    Line2: codatsyncexpenses.String("explicabo"),
                    PostalCode: codatsyncexpenses.String("78221-3550"),
                    Region: codatsyncexpenses.String("magni"),
                    Type: shared.AccountingAddressTypeDelivery,
                },
            },
            ContactName: codatsyncexpenses.String("ipsam"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Items{
                        City: codatsyncexpenses.String("East Marianostead"),
                        Country: codatsyncexpenses.String("Estonia"),
                        Line1: codatsyncexpenses.String("facilis"),
                        Line2: codatsyncexpenses.String("tempore"),
                        PostalCode: codatsyncexpenses.String("94275"),
                        Region: codatsyncexpenses.String("aliquid"),
                        Type: shared.AccountingAddressTypeBilling,
                    },
                    Email: codatsyncexpenses.String("Kianna89@hotmail.com"),
                    ModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
                    Name: codatsyncexpenses.String("Arnold Kirlin"),
                    Phone: []shared.ContactPhone{
                        shared.ContactPhone{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
            },
            CustomerName: codatsyncexpenses.String("cumque"),
            DefaultCurrency: codatsyncexpenses.String("EUR"),
            EmailAddress: codatsyncexpenses.String("ea"),
            ID: codatsyncexpenses.String("6ae395ef-b9ba-488f-ba66-997074ba4469"),
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncexpenses.Bool(false),
            },
            ModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
            Phone: codatsyncexpenses.String("1-911-405-3555 x069"),
            RegistrationNumber: codatsyncexpenses.String("mollitia"),
            SourceModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusActive,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "dolor": map[string]interface{}{
                        "necessitatibus": "odit",
                    },
                },
            },
            TaxNumber: codatsyncexpenses.String("nemo"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "quasi",
        ForceUpdate: codatsyncexpenses.Bool(false),
        TimeoutInMinutes: codatsyncexpenses.Int(435865),
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

