# Transfers

## Overview

Transfers

### Available Operations

* [Create](#create) - Create transfer
* [Get](#get) - Get transfer
* [GetCreateModel](#getcreatemodel) - Get create transfer model
* [List](#list) - List transfers

## Create

Posts a new transfer to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create transfer model](https://docs.codat.io/accounting-api#/operations/get-create-transfers-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers) to see which integrations support this endpoint.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transfers.Create(ctx, operations.CreateTransferRequest{
        Transfer: &shared.Transfer{
            ContactRef: &shared.TransferContactRef{
                DataType: codataccounting.String("atque"),
                ID: "cee41c99-9f46-49f6-b1cf-1a3f023c669e",
            },
            Date: codataccounting.String("2022-10-23T00:00:00.000Z"),
            DepositedRecordRefs: []shared.InvoiceTo{
                shared.InvoiceTo{
                    DataType: codataccounting.String("commodi"),
                    ID: codataccounting.String("26012eba-0579-488c-a720-c3103f1a40c0"),
                },
                shared.InvoiceTo{
                    DataType: codataccounting.String("tenetur"),
                    ID: codataccounting.String("3ec8688f-d8ec-46fc-8312-8f0aaaeee004"),
                },
                shared.InvoiceTo{
                    DataType: codataccounting.String("accusamus"),
                    ID: codataccounting.String("ba7bf873-2be5-409c-9087-131f06f0bce5"),
                },
            },
            Description: codataccounting.String("enim"),
            From: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("a8687143-c979-405f-b797-a5da664b7e77"),
                    Name: codataccounting.String("Shaun Koss"),
                },
                Amount: codataccounting.Float64(6254.03),
                Currency: codataccounting.String("USD"),
            },
            ID: codataccounting.String("a2832bb6-5862-4d2a-b1f9-b14aa6bdec7f"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "dolores": map[string]interface{}{
                        "qui": "voluptates",
                    },
                    "unde": map[string]interface{}{
                        "veniam": "at",
                        "eveniet": "debitis",
                        "inventore": "laborum",
                    },
                },
            },
            To: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("cd72a899-81b5-48fe-a82e-1c2dbe23d58e"),
                    Name: codataccounting.String("Louis Gutkowski"),
                },
                Amount: codataccounting.Float64(1047.48),
                Currency: codataccounting.String("GBP"),
            },
            TrackingCategoryRefs: []shared.TrackingCategoryRef{
                shared.TrackingCategoryRef{
                    ID: "c9f67678-fa27-4958-b673-63da079096fa",
                    Name: codataccounting.String("Alton Larson"),
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateTransferResponse != nil {
        // handle response
    }
}
```

## Get

Gets the specified transfer for a given company.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transfers.Get(ctx, operations.GetTransferRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TransferID: "atque",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transfer != nil {
        // handle response
    }
}
```

## GetCreateModel

Get create transfer model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers) for integrations that support creating transfers.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transfers.GetCreateModel(ctx, operations.GetCreateTransfersModelRequest{
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

## List

Gets the transfers for a given company.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Transfers.List(ctx, operations.ListTransfersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("consequatur"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transfers != nil {
        // handle response
    }
}
```
