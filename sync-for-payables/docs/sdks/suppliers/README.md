# Suppliers
(*Suppliers*)

## Overview

Get, create, and update Suppliers.

### Available Operations

* [List](#list) - List suppliers
* [Create](#create) - Create supplier

## List

The *List suppliers* endpoint returns a list of [suppliers](https://docs.codat.io/sync-for-payables-api#/schemas/Supplier) for a given company's connection.

[Suppliers](https://docs.codat.io/sync-for-payables-api#/schemas/Supplier) are people or organizations that provide something, such as a product or service.

By default, this endpoint returns a list of active and archived suppliers. You can use [querying](https://docs.codat.io/using-the-api/querying) to change that. 

For example, to retrieve only active suppliers (i.e. `status=Active`) or suppliers created within the specified number of days (e.g. `sourceModifiedDate>2023-12-15T00:00:00.000Z`), query the endpoint as follows: `/payables/suppliers?query=sourceModifiedDate>2023-12-15T00:00:00.000Z&&status=Active`.For example, to retrieve active suppliers modified after a particular date use `query=sourceModifiedDate>2023-12-15T00:00:00.000Z&&status=Active`.

### Example Usage

```go
package main

import(
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v5"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Suppliers.List(ctx, operations.ListSuppliersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ContinuationToken: syncforpayables.String("continuationToken=eyJwYWdlIjoyLCJwYWdlU2l6ZSI6MTAwLCJwYWdlQ291bnQiOjExfQ=="),
        Query: syncforpayables.String("sourceModifiedDate>2023-12-15T00:00:00.000Z"),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListSuppliersRequest](../../pkg/models/operations/listsuppliersrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListSuppliersResponse](../../pkg/models/operations/listsuppliersresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.ErrorMessage                      | 400, 401, 402, 403, 404, 409, 429, 500, 503 | application/json                            |
| sdkerrors.SDKError                          | 4XX, 5XX                                    | \*/\*                                       |

## Create

The *Create supplier* endpoint creates a new [supplier](https://docs.codat.io/sync-for-payables-api#/schemas/Supplier) for a given company's connection.

[Suppliers](https://docs.codat.io/sync-for-payables-api#/schemas/Supplier) are people or organizations that provide something, such as a product or service.


### Example Usage

```go
package main

import(
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v5"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := syncforpayables.New(
        syncforpayables.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Suppliers.Create(ctx, operations.CreateSupplierRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        SupplierPrototype: &shared.SupplierPrototype{
            SupplierName: "Greggs",
            ContactName: syncforpayables.String("Greg Greggs"),
            EmailAddress: syncforpayables.String("greg@greggs.com"),
            Phone: syncforpayables.String("+44 (0)1223 322410"),
            Addresses: []shared.Address{
                shared.Address{
                    Type: shared.AddressTypeBilling.ToPointer(),
                    Line1: syncforpayables.String("Flat 1"),
                    Line2: syncforpayables.String("2 Dennis Avenue"),
                    City: syncforpayables.String("London"),
                    Region: syncforpayables.String("Camden"),
                    Country: syncforpayables.String("GBR"),
                    PostalCode: syncforpayables.String("EC1N 7TE"),
                },
            },
            Status: shared.SupplierStatusActive,
            DefaultCurrency: syncforpayables.String("GBP"),
        },
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateSupplierRequest](../../pkg/models/operations/createsupplierrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateSupplierResponse](../../pkg/models/operations/createsupplierresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |