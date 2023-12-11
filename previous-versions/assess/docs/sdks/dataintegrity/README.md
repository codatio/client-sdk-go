# DataIntegrity
(*DataIntegrity*)

## Overview

Match mutable accounting data with immutable banking data to increase confidence in financial data

### Available Operations

* [Details](#details) - List data type data integrity
* [Status](#status) - Get data integrity status
* [Summary](#summary) - Get data integrity summary

## Details

Gets record-by-record match results for a given company and datatype, optionally restricted by a Codat query string.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
	"log"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.DataIntegrity.Details(ctx, operations.ListDataTypeDataIntegrityDetailsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        DataType: shared.DataIntegrityDataTypeBankingAccounts,
        OrderBy: assess.String("-modifiedDate"),
        Page: assess.Int(1),
        PageSize: assess.Int(100),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Details != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.ListDataTypeDataIntegrityDetailsRequest](../../pkg/models/operations/listdatatypedataintegritydetailsrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |


### Response

**[*operations.ListDataTypeDataIntegrityDetailsResponse](../../pkg/models/operations/listdatatypedataintegritydetailsresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |

## Status

Gets match status for a given company and datatype.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
	"log"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.DataIntegrity.Status(ctx, operations.GetDataIntegrityStatusRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        DataType: shared.DataIntegrityDataTypeBankingAccounts,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Status != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetDataIntegrityStatusRequest](../../pkg/models/operations/getdataintegritystatusrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.GetDataIntegrityStatusResponse](../../pkg/models/operations/getdataintegritystatusresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 401,402,403,404,429,500,503 | application/json            |
| sdkerrors.SDKError          | 400-600                     | */*                         |

## Summary

Gets match summary for a given company and datatype, optionally restricted by a Codat query string.

### Example Usage

```go
package main

import(
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
	"log"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.DataIntegrity.Summary(ctx, operations.GetDataIntegritySummariesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        DataType: shared.DataIntegrityDataTypeBankingAccounts,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Summaries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetDataIntegritySummariesRequest](../../pkg/models/operations/getdataintegritysummariesrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |


### Response

**[*operations.GetDataIntegritySummariesResponse](../../pkg/models/operations/getdataintegritysummariesresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,404,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |
