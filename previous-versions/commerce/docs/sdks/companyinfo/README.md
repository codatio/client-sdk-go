# CompanyInfo
(*CompanyInfo*)

## Overview

Retrieve standardized data from linked commerce software.

### Available Operations

* [Get](#get) - Get company info

## Get

Retrieve information about the company, as seen in the commerce software.

This may include information like addresses, tax registration details and social media or website information.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/commerce"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/operations"
	"log"
)

func main() {
    s := commerce.New(
        commerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.CompanyInfo.Get(ctx, operations.GetCompanyInfoRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CompanyInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetCompanyInfoRequest](../../pkg/models/operations/getcompanyinforequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetCompanyInfoResponse](../../pkg/models/operations/getcompanyinforesponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 401,402,403,404,409,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
