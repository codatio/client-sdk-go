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
                DataType: codataccounting.String("maiores"),
                ID: "c8745005-e9d3-4d93-8e03-6f5c388664f6",
            },
            Date: codataccounting.String("error"),
            DepositedRecordRefs: []shared.InvoiceTo{
                shared.InvoiceTo{
                    DataType: codataccounting.String("corporis"),
                    ID: codataccounting.String("530a2e2a-ed6a-4af8-a3c2-8d040c69a3d9"),
                },
                shared.InvoiceTo{
                    DataType: codataccounting.String("voluptatem"),
                    ID: codataccounting.String("6f6ebd5a-d7ec-4739-8f25-f634b3730714"),
                },
                shared.InvoiceTo{
                    DataType: codataccounting.String("itaque"),
                    ID: codataccounting.String("6be8c3e0-9c64-4d34-aac2-99a6e5e7aef1"),
                },
            },
            Description: codataccounting.String("amet"),
            From: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("402e945f-5374-43ef-9e11-98221f9b1f7d"),
                    Name: codataccounting.String("Gerard Weimann"),
                },
                Amount: codataccounting.Float64(4322.15),
                Currency: codataccounting.String("omnis"),
            },
            ID: codataccounting.String("682aceef-b04f-48c5-92ca-abea708ed579"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("rem"),
            SourceModifiedDate: codataccounting.String("facere"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "quas": map[string]interface{}{
                        "illum": "labore",
                        "ea": "aperiam",
                    },
                },
            },
            To: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("599d5c33-4957-46d5-9209-e9a2253b6d76"),
                    Name: codataccounting.String("Irma Larkin"),
                },
                Amount: codataccounting.Float64(9345.12),
                Currency: codataccounting.String("similique"),
            },
            TrackingCategoryRefs: []shared.TrackingCategoryRef{
                shared.TrackingCategoryRef{
                    ID: "5fd4b39f-8a14-4906-b8f1-3c686d839fc9",
                    Name: codataccounting.String("Joshua Koelpin"),
                },
                shared.TrackingCategoryRef{
                    ID: "fa906ae5-59b7-42eb-a746-030fe18376c2",
                    Name: codataccounting.String("Merle Strosin"),
                },
                shared.TrackingCategoryRef{
                    ID: "76790ed0-c16a-47ba-8784-04489f6770ef",
                    Name: codataccounting.String("Ms. Suzanne Lang MD"),
                },
                shared.TrackingCategoryRef{
                    ID: "2ba25ee6-c75a-4f8a-a0a7-ae346e0979e5",
                    Name: codataccounting.String("Timmy Trantow V"),
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
        TransferID: "impedit",
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
        Query: codataccounting.String("culpa"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transfers != nil {
        // handle response
    }
}
```
