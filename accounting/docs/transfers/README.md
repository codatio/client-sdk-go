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
                ID: "a2795836-7363-4da0-b909-6faeb8648073",
            },
            Date: codataccounting.String("2022-10-23T00:00:00.000Z"),
            DepositedRecordRefs: []shared.InvoiceTo{
                shared.InvoiceTo{
                    DataType: codataccounting.String("laudantium"),
                    ID: codataccounting.String("f8b89d9c-a607-4565-afc0-ebe67155e2d0"),
                },
                shared.InvoiceTo{
                    DataType: codataccounting.String("autem"),
                    ID: codataccounting.String("a3070d6e-297f-4581-baba-aa7d80108807"),
                },
                shared.InvoiceTo{
                    DataType: codataccounting.String("laboriosam"),
                    ID: codataccounting.String("ff5f6ed2-9814-4088-a69b-6a70b0dd82f9"),
                },
                shared.InvoiceTo{
                    DataType: codataccounting.String("numquam"),
                    ID: codataccounting.String("fffbd1e1-e21d-4dc6-9038-b1d187b51eb5"),
                },
            },
            Description: codataccounting.String("doloribus"),
            From: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("d30bfe03-490c-4f20-a54a-959043cb462d"),
                    Name: codataccounting.String("Karla Schimmel"),
                },
                Amount: codataccounting.Float64(944.75),
                Currency: codataccounting.String("USD"),
            },
            ID: codataccounting.String("f98e4792-b979-4a41-bd6a-8c91683bd861"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "natus": map[string]interface{}{
                        "quod": "quo",
                        "repellat": "voluptatum",
                        "excepturi": "illum",
                    },
                    "amet": map[string]interface{}{
                        "ex": "quae",
                        "beatae": "praesentium",
                        "commodi": "vero",
                    },
                    "temporibus": map[string]interface{}{
                        "nisi": "minus",
                        "eaque": "consequatur",
                    },
                    "magni": map[string]interface{}{
                        "est": "cumque",
                        "harum": "dicta",
                        "nesciunt": "dolorum",
                        "placeat": "sed",
                    },
                },
            },
            To: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("4c8143b8-66c5-475a-9e26-68730be37b0e"),
                    Name: codataccounting.String("Terrell Reichert"),
                },
                Amount: codataccounting.Float64(5391.45),
                Currency: codataccounting.String("EUR"),
            },
            TrackingCategoryRefs: []shared.TrackingCategoryRef{
                shared.TrackingCategoryRef{
                    ID: "c7e69b53-5105-4050-94dc-a105882484c3",
                    Name: codataccounting.String("Raquel Metz"),
                },
                shared.TrackingCategoryRef{
                    ID: "892782d3-4e0b-48fc-8d59-f57b9f9820be",
                    Name: codataccounting.String("Ms. Heidi Lind"),
                },
                shared.TrackingCategoryRef{
                    ID: "36c9e2f7-0344-4e00-b478-eb539483f748",
                    Name: codataccounting.String("Santiago Windler"),
                },
                shared.TrackingCategoryRef{
                    ID: "b69d541b-4b39-43f3-9666-25bea32201de",
                    Name: codataccounting.String("Earl Kiehn"),
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
        TransferID: "voluptas",
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
        Query: codataccounting.String("error"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transfers != nil {
        // handle response
    }
}
```
