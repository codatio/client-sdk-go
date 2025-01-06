# DataIntegrity
(*DataIntegrity*)

## Overview

Match mutable accounting data with immutable banking data to increase confidence in financial data.

### Available Operations

* [Details](#details) - List data integrity details
* [Status](#status) - Get data integrity status
* [Summaries](#summaries) - Get data integrity summaries

## Details

The *List data integrity details* endpoint returns the match result record by record for a given data type, filtered based on a query string in the same way as summary results.

The [details](https://docs.codat.io/lending-api#/schemas/DataIntegrityDetail) are paginated and support ordering, following the same conventions as our other data endpoints.

### Example Usage

```go
package main

import(
	"context"
	lending "github.com/codatio/client-sdk-go/lending/v8"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.DataIntegrity.Details(ctx, operations.ListDataIntegrityDetailsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        DataType: shared.DataIntegrityDataTypeBankingAccounts,
        OrderBy: lending.String("-modifiedDate"),
        Page: lending.Int(1),
        PageSize: lending.Int(100),
        Query: lending.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DataIntegrityDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListDataIntegrityDetailsRequest](../../pkg/models/operations/listdataintegritydetailsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListDataIntegrityDetailsResponse](../../pkg/models/operations/listdataintegritydetailsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

## Status

The *Get data integrity status* endpoint returns the [status](https://docs.codat.io/lending-api#/schemas/DataIntegrityStatus) for the company’s match results between the data type provided in the URL and other data types that Data Integrity uses to support matching.
This endpoint helps you understand whether match data is available and, if so, how to usefully query it.

The response tells you:

- Whether match results are available.
- When the results were generated, and their status.
- The connection IDs, amounts, and dates involved to support useful querying.

### Example Usage

```go
package main

import(
	"context"
	lending "github.com/codatio/client-sdk-go/lending/v8"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.DataIntegrity.Status(ctx, operations.GetDataIntegrityStatusRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        DataType: shared.DataIntegrityDataTypeBankingAccounts,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DataIntegrityStatuses != nil {
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

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.ErrorMessage            | 401, 402, 403, 404, 429, 500, 503 | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## Summaries

The *Get data integrity summary* endpoint returns a [summary](https://docs.codat.io/lending-api#/schemas/DataIntegritySummary) of match results for a given data type filtered by a query string in the [Codat query language](https://docs.codat.io/using-the-api/querying). 

For example, if you wanted to see summary match results only for transactions after 1 December 2020, you could include a query parameter of `query=date>2020-12-01`.

The endpoint response includes only the summary results, not transactions. To view match data for transactions, use the [List data integrity details](https://docs.codat.io/lending-api#/operations/get-data-integrity-summaries) endpoint.

### Example Usage

```go
package main

import(
	"context"
	lending "github.com/codatio/client-sdk-go/lending/v8"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := lending.New(
        lending.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    res, err := s.DataIntegrity.Summaries(ctx, operations.GetDataIntegritySummariesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        DataType: shared.DataIntegrityDataTypeBankingAccounts,
        Query: lending.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DataIntegritySummaries != nil {
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

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.ErrorMessage                 | 400, 401, 402, 403, 404, 429, 500, 503 | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |