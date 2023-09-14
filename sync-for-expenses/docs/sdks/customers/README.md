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
	"github.com/codatio/client-sdk-go/sync-for-expenses"
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/models/operations"
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
                    City: codatsyncexpenses.String("West Christa"),
                    Country: codatsyncexpenses.String("Iceland"),
                    Line1: codatsyncexpenses.String("cupiditate"),
                    Line2: codatsyncexpenses.String("quos"),
                    PostalCode: codatsyncexpenses.String("18301"),
                    Region: codatsyncexpenses.String("dolorum"),
                    Type: shared.AccountingAddressTypeBilling,
                },
            },
            ContactName: codatsyncexpenses.String("tempora"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Items{
                        City: codatsyncexpenses.String("Riceboro"),
                        Country: codatsyncexpenses.String("Vanuatu"),
                        Line1: codatsyncexpenses.String("eum"),
                        Line2: codatsyncexpenses.String("non"),
                        PostalCode: codatsyncexpenses.String("53585-6289"),
                        Region: codatsyncexpenses.String("dolorum"),
                        Type: shared.AccountingAddressTypeBilling,
                    },
                    Email: codatsyncexpenses.String("Rose.Wolff29@yahoo.com"),
                    ModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
                    Name: codatsyncexpenses.String("Nathaniel Hyatt"),
                    Phone: []shared.ContactPhone{
                        shared.ContactPhone{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypeMobile,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
            },
            CustomerName: codatsyncexpenses.String("accusamus"),
            DefaultCurrency: codatsyncexpenses.String("EUR"),
            EmailAddress: codatsyncexpenses.String("quidem"),
            ID: codatsyncexpenses.String("9ba88f3a-6699-4707-8ba4-469b6e214195"),
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncexpenses.Bool(false),
            },
            ModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
            Phone: codatsyncexpenses.String("606-963-4281 x3049"),
            RegistrationNumber: codatsyncexpenses.String("debitis"),
            SourceModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusArchived,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "deleniti": map[string]interface{}{
                        "facilis": "in",
                    },
                },
            },
            TaxNumber: codatsyncexpenses.String("architecto"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncexpenses.Int(99569),
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
	"github.com/codatio/client-sdk-go/sync-for-expenses"
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/models/operations"
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
        CustomerID: "repudiandae",
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
	"github.com/codatio/client-sdk-go/sync-for-expenses"
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/models/operations"
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
        Query: codatsyncexpenses.String("ullam"),
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
	"github.com/codatio/client-sdk-go/sync-for-expenses"
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/models/operations"
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
                    City: codatsyncexpenses.String("Kossworth"),
                    Country: codatsyncexpenses.String("Sudan"),
                    Line1: codatsyncexpenses.String("sed"),
                    Line2: codatsyncexpenses.String("saepe"),
                    PostalCode: codatsyncexpenses.String("01561-1788"),
                    Region: codatsyncexpenses.String("maxime"),
                    Type: shared.AccountingAddressTypeBilling,
                },
            },
            ContactName: codatsyncexpenses.String("excepturi"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Items{
                        City: codatsyncexpenses.String("South Alexanneton"),
                        Country: codatsyncexpenses.String("Wallis and Futuna"),
                        Line1: codatsyncexpenses.String("quidem"),
                        Line2: codatsyncexpenses.String("ipsam"),
                        PostalCode: codatsyncexpenses.String("47083"),
                        Region: codatsyncexpenses.String("voluptatibus"),
                        Type: shared.AccountingAddressTypeUnknown,
                    },
                    Email: codatsyncexpenses.String("Darian.Anderson94@hotmail.com"),
                    ModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
                    Name: codatsyncexpenses.String("Ernest Hayes"),
                    Phone: []shared.ContactPhone{
                        shared.ContactPhone{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
            },
            CustomerName: codatsyncexpenses.String("eos"),
            DefaultCurrency: codatsyncexpenses.String("GBP"),
            EmailAddress: codatsyncexpenses.String("dolores"),
            ID: codatsyncexpenses.String("c73d5fe9-b90c-4289-89b3-fe49a8d9cbf4"),
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncexpenses.Bool(false),
            },
            ModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
            Phone: codatsyncexpenses.String("1-322-329-5744 x926"),
            RegistrationNumber: codatsyncexpenses.String("numquam"),
            SourceModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusUnknown,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "ipsa": map[string]interface{}{
                        "iure": "odio",
                    },
                },
            },
            TaxNumber: codatsyncexpenses.String("quaerat"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "accusamus",
        ForceUpdate: codatsyncexpenses.Bool(false),
        TimeoutInMinutes: codatsyncexpenses.Int(696344),
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

