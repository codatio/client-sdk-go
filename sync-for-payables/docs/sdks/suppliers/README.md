# Suppliers

## Overview

Suppliers

### Available Operations

* [Create](#create) - Create supplier
* [Get](#get) - Get supplier
* [GetCreateUpdateModel](#getcreateupdatemodel) - Get create/update supplier model
* [List](#list) - List suppliers
* [Update](#update) - Update supplier

## Create

The *Create supplier* endpoint creates a new [supplier](https://docs.codat.io/sync-for-payables-api#/schemas/Supplier) for a given company's connection.

[Suppliers](https://docs.codat.io/sync-for-payables-api#/schemas/Supplier) are people or organizations that provide something, such as a product or service.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update supplier model](https://docs.codat.io/sync-for-payables-api#/operations/get-create-update-suppliers-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers) for integrations that support creating a supplier.


### Example Usage

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
    res, err := s.Suppliers.Create(ctx, operations.CreateSupplierRequest{
        Supplier: &shared.Supplier{
            Addresses: []shared.SupplierAccountingAddress{
                shared.SupplierAccountingAddress{
                    City: codatsyncpayables.String("West Jacey"),
                    Country: codatsyncpayables.String("Central African Republic"),
                    Line1: codatsyncpayables.String("excepturi"),
                    Line2: codatsyncpayables.String("esse"),
                    PostalCode: codatsyncpayables.String("82708-8277"),
                    Region: codatsyncpayables.String("impedit"),
                    Type: shared.AccountingAddressTypeDelivery,
                },
                shared.SupplierAccountingAddress{
                    City: codatsyncpayables.String("West Jaquanchester"),
                    Country: codatsyncpayables.String("Kenya"),
                    Line1: codatsyncpayables.String("aliquid"),
                    Line2: codatsyncpayables.String("blanditiis"),
                    PostalCode: codatsyncpayables.String("99930"),
                    Region: codatsyncpayables.String("beatae"),
                    Type: shared.AccountingAddressTypeUnknown,
                },
            },
            ContactName: codatsyncpayables.String("laboriosam"),
            DefaultCurrency: codatsyncpayables.String("temporibus"),
            EmailAddress: codatsyncpayables.String("in"),
            ID: codatsyncpayables.String("1cffbd0e-b74b-4842-9953-b44bd3c43159"),
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncpayables.Bool(false),
            },
            ModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Phone: codatsyncpayables.String("+44 25691 154789"),
            RegistrationNumber: codatsyncpayables.String("velit"),
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.SupplierStatusUnknown,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "nostrum": map[string]interface{}{
                        "quod": "consequatur",
                    },
                    "accusantium": map[string]interface{}{
                        "illo": "amet",
                    },
                    "occaecati": map[string]interface{}{
                        "aliquid": "sequi",
                        "culpa": "fuga",
                        "modi": "et",
                    },
                },
            },
            SupplierName: codatsyncpayables.String("eveniet"),
            TaxNumber: codatsyncpayables.String("aliquid"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsyncpayables.Int(760942),
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

The *Get supplier* endpoint returns a single supplier for a given `supplierId`.

[Suppliers](https://docs.codat.io/sync-for-payables-api#/schemas/Supplier) are people or organizations that provide something, such as a product or service.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers) for integrations that support getting a specific supplier.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-payables-api#/operations/refresh-company-data).


### Example Usage

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
    res, err := s.Suppliers.Get(ctx, operations.GetSupplierRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SupplierID: "adipisci",
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


## GetCreateUpdateModel

The *Get create/update supplier model* endpoint returns the expected data for the request payload when creating and updating a [supplier](https://docs.codat.io/sync-for-payables-api#/schemas/Supplier) for a given company and integration.

[Suppliers](https://docs.codat.io/sync-for-payables-api#/schemas/Supplier) are people or organizations that provide something, such as a product or service.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers) for integrations that support creating and updating a supplier.


### Example Usage

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
    res, err := s.Suppliers.GetCreateUpdateModel(ctx, operations.GetCreateUpdateSupplierModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetCreateUpdateSupplierModelRequest](../../models/operations/getcreateupdatesuppliermodelrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |


### Response

**[*operations.GetCreateUpdateSupplierModelResponse](../../models/operations/getcreateupdatesuppliermodelresponse.md), error**


## List

The *List suppliers* endpoint returns a list of [suppliers](https://docs.codat.io/sync-for-payables-api#/schemas/Supplier) for a given company's connection.

[Suppliers](https://docs.codat.io/sync-for-payables-api#/schemas/Supplier) are people or organizations that provide something, such as a product or service.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/sync-for-payables-api#/operations/refresh-company-data).
    

### Example Usage

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
    res, err := s.Suppliers.List(ctx, operations.ListSuppliersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codatsyncpayables.String("-modifiedDate"),
        Page: codatsyncpayables.Int(1),
        PageSize: codatsyncpayables.Int(100),
        Query: codatsyncpayables.String("ab"),
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

The *Update supplier* endpoint updates an existing [supplier](https://docs.codat.io/sync-for-payables-api#/schemas/Supplier) for a given company's connection.

[Suppliers](https://docs.codat.io/sync-for-payables-api#/schemas/Supplier) are people or organizations that provide something, such as a product or service.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update supplier model](https://docs.codat.io/sync-for-payables-api#/operations/get-create-update-suppliers-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers) for integrations that support creating a supplier.


### Example Usage

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
    res, err := s.Suppliers.Update(ctx, operations.UpdateSupplierRequest{
        Supplier: &shared.Supplier{
            Addresses: []shared.SupplierAccountingAddress{
                shared.SupplierAccountingAddress{
                    City: codatsyncpayables.String("Champlinworth"),
                    Country: codatsyncpayables.String("Brazil"),
                    Line1: codatsyncpayables.String("hic"),
                    Line2: codatsyncpayables.String("porro"),
                    PostalCode: codatsyncpayables.String("31756-2099"),
                    Region: codatsyncpayables.String("distinctio"),
                    Type: shared.AccountingAddressTypeDelivery,
                },
                shared.SupplierAccountingAddress{
                    City: codatsyncpayables.String("Schmittchester"),
                    Country: codatsyncpayables.String("Sri Lanka"),
                    Line1: codatsyncpayables.String("iure"),
                    Line2: codatsyncpayables.String("natus"),
                    PostalCode: codatsyncpayables.String("99438"),
                    Region: codatsyncpayables.String("doloremque"),
                    Type: shared.AccountingAddressTypeBilling,
                },
                shared.SupplierAccountingAddress{
                    City: codatsyncpayables.String("Oswaldofield"),
                    Country: codatsyncpayables.String("Kyrgyz Republic"),
                    Line1: codatsyncpayables.String("mollitia"),
                    Line2: codatsyncpayables.String("quidem"),
                    PostalCode: codatsyncpayables.String("41396-3740"),
                    Region: codatsyncpayables.String("vel"),
                    Type: shared.AccountingAddressTypeUnknown,
                },
                shared.SupplierAccountingAddress{
                    City: codatsyncpayables.String("New Dellboro"),
                    Country: codatsyncpayables.String("Saint Barthelemy"),
                    Line1: codatsyncpayables.String("occaecati"),
                    Line2: codatsyncpayables.String("sit"),
                    PostalCode: codatsyncpayables.String("18058-0681"),
                    Region: codatsyncpayables.String("a"),
                    Type: shared.AccountingAddressTypeDelivery,
                },
            },
            ContactName: codatsyncpayables.String("explicabo"),
            DefaultCurrency: codatsyncpayables.String("delectus"),
            EmailAddress: codatsyncpayables.String("natus"),
            ID: codatsyncpayables.String("e2e10594-4b93-45d2-b7a7-2f90849d6aed"),
            Metadata: &shared.Metadata{
                IsDeleted: codatsyncpayables.Bool(false),
            },
            ModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Phone: codatsyncpayables.String("(877) 492-8687"),
            RegistrationNumber: codatsyncpayables.String("eveniet"),
            SourceModifiedDate: codatsyncpayables.String("2022-10-23T00:00:00.000Z"),
            Status: shared.SupplierStatusArchived,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "nemo": map[string]interface{}{
                        "esse": "placeat",
                    },
                    "at": map[string]interface{}{
                        "eos": "odit",
                        "quia": "maxime",
                        "excepturi": "sapiente",
                    },
                },
            },
            SupplierName: codatsyncpayables.String("maiores"),
            TaxNumber: codatsyncpayables.String("exercitationem"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codatsyncpayables.Bool(false),
        SupplierID: "ducimus",
        TimeoutInMinutes: codatsyncpayables.Int(293512),
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

