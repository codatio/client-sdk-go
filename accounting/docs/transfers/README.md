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
                DataType: codataccounting.String("magni"),
                ID: "73caa911-8b38-4f1b-a1a3-31a54dc10294",
            },
            Date: codataccounting.String("reiciendis"),
            DepositedRecordRefs: []shared.InvoiceTo{
                shared.InvoiceTo{
                    DataType: codataccounting.String("eos"),
                    ID: codataccounting.String("fed939ba-8f71-4e29-92c2-0ee1228ac3ad"),
                },
                shared.InvoiceTo{
                    DataType: codataccounting.String("optio"),
                    ID: codataccounting.String("647d240b-c11e-4a48-a824-ccc6a2f0f5b9"),
                },
                shared.InvoiceTo{
                    DataType: codataccounting.String("quibusdam"),
                    ID: codataccounting.String("3cb11a76-87d3-4100-a8e2-b9b0d746d2a7"),
                },
            },
            Description: codataccounting.String("quisquam"),
            From: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("7d1ea0e7-9fa9-4bbe-9f17-9f650b1e707e"),
                    Name: codataccounting.String("Rochelle Grimes"),
                },
                Amount: codataccounting.Float64(4097.14),
                Currency: codataccounting.String("odio"),
            },
            ID: codataccounting.String("13bacce0-72ab-4d61-918d-279c10c18516"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("doloribus"),
            SourceModifiedDate: codataccounting.String("repellendus"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "praesentium": map[string]interface{}{
                        "repudiandae": "fugit",
                        "vel": "fugit",
                        "ab": "quia",
                    },
                    "esse": map[string]interface{}{
                        "ea": "odit",
                    },
                },
            },
            To: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("8fa50396-2867-4e72-b3a6-5024b157f9bb"),
                    Name: codataccounting.String("Delia Zulauf"),
                },
                Amount: codataccounting.Float64(6278.1),
                Currency: codataccounting.String("ad"),
            },
            TrackingCategoryRefs: []shared.TrackingCategoryRef{
                shared.TrackingCategoryRef{
                    ID: "871d99b6-61a7-4def-968b-6ccb2822b4a9",
                    Name: codataccounting.String("Roberto Abshire"),
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
        TransferID: "magni",
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
        Query: codataccounting.String("doloribus"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transfers != nil {
        // handle response
    }
}
```
