# SupplementalData
(*SupplementalData*)

## Overview

View and configure supplemental data for supported data types.

### Available Operations

* [Configure](#configure) - Configure
* [GetConfiguration](#getconfiguration) - Get configuration

## Configure

The *Configure* endpoint allows you to maintain or change configuration required to return supplemental data for each integration and data type combination.

[Supplemental data](https://docs.codat.io/using-the-api/additional-data) is additional data you can include in Codat's standard data types.

**Integration-specific behaviour**
See the *examples* for integration-specific frequently requested properties.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"context"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := platform.New(
        platform.WithSecurity(shared.Security{
            AuthHeader: "<YOUR_API_KEY_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.SupplementalData.Configure(ctx, operations.ConfigureSupplementalDataRequest{
        SupplementalDataConfiguration: &shared.SupplementalDataConfiguration{
            SupplementalDataConfig: map[string]shared.SupplementalDataSourceConfiguration{
                "key": shared.SupplementalDataSourceConfiguration{
                    PullData: map[string]string{
                        "key": "string",
                    },
                    PushData: map[string]string{
                        "key": "string",
                    },
                },
            },
        },
        DataType: operations.DataTypeInvoices,
        PlatformKey: "gbol",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ConfigureSupplementalDataRequest](../../pkg/models/operations/configuresupplementaldatarequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |


### Response

**[*operations.ConfigureSupplementalDataResponse](../../pkg/models/operations/configuresupplementaldataresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## GetConfiguration

The *Get configuration* endpoint returns supplemental data configuration previously created for each integration and data type combination.

[Supplemental data](https://docs.codat.io/using-the-api/additional-data) is additional data you can include in Codat's standard data types.

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
    res, err := s.SupplementalData.GetConfiguration(ctx, operations.GetSupplementalDataConfigurationRequest{
        DataType: operations.PathParamDataTypeInvoices,
        PlatformKey: "gbol",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SupplementalDataConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetSupplementalDataConfigurationRequest](../../pkg/models/operations/getsupplementaldataconfigurationrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |


### Response

**[*operations.GetSupplementalDataConfigurationResponse](../../pkg/models/operations/getsupplementaldataconfigurationresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |
