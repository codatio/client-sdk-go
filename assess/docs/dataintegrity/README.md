# DataIntegrity

## Overview

Data integrity is important

### Available Operations

* [GetDataIntegrityDetails](#getdataintegritydetails) - Lists data integrity details for date type
* [GetDataIntegrityStatus](#getdataintegritystatus) - Get data integrity status
* [GetDataIntegritySummaries](#getdataintegritysummaries) - Get data integrity summary

## GetDataIntegrityDetails

Gets record-by-record match results for a given company and datatype, optionally restricted by a Codat query string.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.DataIntegrity.GetDataIntegrityDetails(ctx, operations.GetDataIntegrityDetailsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        DataType: shared.DataIntegrityDataTypeBankingAccounts,
        OrderBy: codatassess.String("-modifiedDate"),
        Page: 1,
        PageSize: codatassess.Int(100),
        Query: codatassess.String("voluptatibus"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Details != nil {
        // handle response
    }
}
```

## GetDataIntegrityStatus

Gets match status for a given company and datatype.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.DataIntegrity.GetDataIntegrityStatus(ctx, operations.GetDataIntegrityStatusRequest{
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

## GetDataIntegritySummaries

Gets match summary for a given company and datatype, optionally restricted by a Codat query string.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/assess"
	"github.com/codatio/client-sdk-go/assess/pkg/models/operations"
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

func main() {
    s := codatassess.New(
        codatassess.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.DataIntegrity.GetDataIntegritySummaries(ctx, operations.GetDataIntegritySummariesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        DataType: shared.DataIntegrityDataTypeBankingAccounts,
        Query: codatassess.String("ipsa"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Summaries != nil {
        // handle response
    }
}
```
