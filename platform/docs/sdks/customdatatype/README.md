# CustomDataType
(*CustomDataType*)

## Overview

View and configure custom data types for supported integrations.

### Available Operations

* [Configure](#configure) - Configure custom data type
* [GetConfiguration](#getconfiguration) - Get custom data configuration
* [List](#list) - List custom data type records
* [Refresh](#refresh) - Refresh custom data type

## Configure

The *Configure custom data type* endpoint allows you to maintain or change the configuration required to return a custom data type for a specific integration. 

A [custom data type](https://docs.codat.io/using-the-api/custom-data) is an additional data type you can create that is not included in Codat's standardized data model.

### Tips and traps

- You can only configure a single custom data type for a single platform at a time. Use the endpoint multiple times if you need to configure it for multiple platforms. 

- You can only indicate a single data source for each customer data type. 

- Make your custom configuration as similar as possible to our standard data types so you can interact with them in exactly the same way.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"context"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/operations"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "<YOUR_API_KEY_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.CustomDataType.Configure(ctx, operations.ConfigureCustomDataTypeRequest{
        CustomDataTypeConfiguration: &shared.CustomDataTypeConfiguration{
            DataSource: platform.String("api/purchaseOrders?$filter=currencyCode eq 'NOK'"),
            KeyBy: []string{
                "$[*].id",
            },
            RequiredData: map[string]string{
                "currencyCode": "$[*].currencyCode",
                "id": "$[*].id",
                "number": "$[*].number",
                "orderDate": "$[*].orderDate",
                "totalAmountExcludingTax": "$[*].totalAmountExcludingTax",
                "totalTaxAmount": "$[*].totalTaxAmount",
                "vendorName": "$[*].number",
            },
            SourceModifiedDate: []string{
                "$[*].lastModifiedDateTime",
            },
        },
        CustomDataIdentifier: "DynamicsPurchaseOrders",
        PlatformKey: "gbol",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CustomDataTypeConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ConfigureCustomDataTypeRequest](../../pkg/models/operations/configurecustomdatatyperequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |


### Response

**[*operations.ConfigureCustomDataTypeResponse](../../pkg/models/operations/configurecustomdatatyperesponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## GetConfiguration

The *Get custom data configuration* endpoint returns existing configuration details for the specified custom data type and integration pair you previously configured.

A [custom data type](https://docs.codat.io/using-the-api/custom-data) is an additional data type you can create that is not included in Codat's standardized data model.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"context"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/operations"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "<YOUR_API_KEY_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.CustomDataType.GetConfiguration(ctx, operations.GetCustomDataTypeConfigurationRequest{
        CustomDataIdentifier: "DynamicsPurchaseOrders",
        PlatformKey: "gbol",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CustomDataTypeRecords != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.GetCustomDataTypeConfigurationRequest](../../pkg/models/operations/getcustomdatatypeconfigurationrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |


### Response

**[*operations.GetCustomDataTypeConfigurationResponse](../../pkg/models/operations/getcustomdatatypeconfigurationresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## List

The *List custom data type records* endpoint returns a paginated list of records pulled for the specified custom data type you previously configured.

A [custom data type](https://docs.codat.io/using-the-api/custom-data) is an additional data type you can create that is not included in Codat's standardized data model.s endpoint returns a paginated list of records whose schema is defined [Configure custom data type](https://docs.codat.io/platform-api#/operations/configure-custom-data-type)

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"context"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/operations"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "<YOUR_API_KEY_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.CustomDataType.List(ctx, operations.ListCustomDataTypeRecordsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomDataIdentifier: "DynamicsPurchaseOrders",
        Page: platform.Int(1),
        PageSize: platform.Int(100),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CustomDataTypeRecords != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListCustomDataTypeRecordsRequest](../../pkg/models/operations/listcustomdatatyperecordsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |


### Response

**[*operations.ListCustomDataTypeRecordsResponse](../../pkg/models/operations/listcustomdatatyperecordsresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.ErrorMessage              | 400,401,402,403,404,429,451,500,503 | application/json                    |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |

## Refresh

The *Refresh custom data type* endpoint refreshes the specified custom data type for a given company. This is an asynchronous operation that will sync updated data from the linked integration into Codat for you to view.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"context"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/operations"
	"log"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "<YOUR_API_KEY_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.CustomDataType.Refresh(ctx, operations.RefreshCustomDataTypeRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomDataIdentifier: "DynamicsPurchaseOrders",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PullOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RefreshCustomDataTypeRequest](../../pkg/models/operations/refreshcustomdatatyperequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.RefreshCustomDataTypeResponse](../../pkg/models/operations/refreshcustomdatatyperesponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 401,402,403,404,429,451,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
