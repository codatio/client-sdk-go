# ConnectionManagement
(*ConnectionManagement*)

## Overview

Configure UI and retrieve access tokens for authentication used by **Connections SDK**.

### Available Operations

* [GetAccessToken](#getaccesstoken) - Get access token

## GetAccessToken

﻿Use the *Get access token* endpoint to retrieve a new access token for use with the [Connections SDK](https://docs.codat.io/auth-flow/optimize/connection-management). The token is only valid for one hour and applies to a single company.

The embeddable [Connections SDK](https://docs.codat.io/auth-flow/optimize/connection-management) lets your customers control access to their data by allowing them to manage their existing connections.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/operations"
	"log"
)

func main() {
    s := common.New(
        common.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.ConnectionManagement.GetAccessToken(ctx, operations.GetConnectionManagementAccessTokenRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectionManagementAccessToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.GetConnectionManagementAccessTokenRequest](../../pkg/models/operations/getconnectionmanagementaccesstokenrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.GetConnectionManagementAccessTokenResponse](../../pkg/models/operations/getconnectionmanagementaccesstokenresponse.md), error**

### Errors

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |
