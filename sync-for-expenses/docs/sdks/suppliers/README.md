# Suppliers

## Overview

Suppliers

### Available Operations

* [Create](#create) - Create supplier
* [Get](#get) - Get supplier
* [List](#list) - List suppliers
* [Update](#update) - Update supplier

## Create

The *Create supplier* endpoint creates a new [supplier](https://docs.codat.io/accounting-api#/schemas/Supplier) for a given company's connection.

[Suppliers](https://docs.codat.io/accounting-api#/schemas/Supplier) are people or organizations that provide something, such as a product or service.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update supplier model](https://docs.codat.io/accounting-api#/operations/get-create-update-suppliers-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers) for integrations that support creating an account.


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
    res, err := s.Suppliers.Create(ctx, operations.CreateSupplierRequest{
        Supplier: &shared.Supplier{
            ID: codatsyncexpenses.String("124"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncexpenses.Int(673660),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateSupplierResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateSupplierRequest](../../models/operations/createsupplierrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.CreateSupplierResponse](../../models/operations/createsupplierresponse.md), error**


## Get

The *Get supplier* endpoint returns a single supplier for a given supplierId.

[Suppliers](https://docs.codat.io/accounting-api#/schemas/Supplier) are people or organizations that provide something, such as a product or service.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers) for integrations that support getting a specific supplier.

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
    res, err := s.Suppliers.Get(ctx, operations.GetSupplierRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SupplierID: "13d946f0-c5d5-42bc-b092-97ece17923ab",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Supplier != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetSupplierRequest](../../models/operations/getsupplierrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |


### Response

**[*operations.GetSupplierResponse](../../models/operations/getsupplierresponse.md), error**


## List

The *List suppliers* endpoint returns a list of [suppliers](https://docs.codat.io/accounting-api#/schemas/Supplier) for a given company's connection.

[Suppliers](https://docs.codat.io/accounting-api#/schemas/Supplier) are people or organizations that provide something, such as a product or service.

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
    res, err := s.Suppliers.List(ctx, operations.ListSuppliersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codatsyncexpenses.String("-modifiedDate"),
        Page: codatsyncexpenses.Int(1),
        PageSize: codatsyncexpenses.Int(100),
        Query: codatsyncexpenses.String("reiciendis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Suppliers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListSuppliersRequest](../../models/operations/listsuppliersrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.ListSuppliersResponse](../../models/operations/listsuppliersresponse.md), error**


## Update

The *Update supplier* endpoint updates an existing [supplier](https://docs.codat.io/accounting-api#/schemas/Supplier) for a given company's connection.

[Suppliers](https://docs.codat.io/accounting-api#/schemas/Supplier) are people or organizations that provide something, such as a product or service.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update supplier model](https://docs.codat.io/accounting-api#/operations/get-create-update-suppliers-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers) for integrations that support creating an account.


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
    res, err := s.Suppliers.Update(ctx, operations.UpdateSupplierRequest{
        Supplier: &shared.Supplier{
            ID: codatsyncexpenses.String("124"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codatsyncexpenses.Bool(false),
        SupplierID: "EILBDVJVNUAGVKRQ",
        TimeoutInMinutes: codatsyncexpenses.Int(878194),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateSupplierResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateSupplierRequest](../../models/operations/updatesupplierrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.UpdateSupplierResponse](../../models/operations/updatesupplierresponse.md), error**
