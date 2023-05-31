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
                DataType: codataccounting.String("explicabo"),
                ID: "1f06d4d1-7852-4d28-be1d-b01d6919f831",
            },
            Date: codataccounting.String("nemo"),
            DepositedRecordRefs: []shared.InvoiceTo{
                shared.InvoiceTo{
                    DataType: codataccounting.String("dolorem"),
                    ID: codataccounting.String("a84ea7db-15c4-4c15-be6c-d097a675597e"),
                },
                shared.InvoiceTo{
                    DataType: codataccounting.String("cumque"),
                    ID: codataccounting.String("beb7982b-af9a-47da-ac29-b938e51a7e6e"),
                },
                shared.InvoiceTo{
                    DataType: codataccounting.String("nulla"),
                    ID: codataccounting.String("6f7ff04f-da04-4669-aae8-182403655aa9"),
                },
            },
            Description: codataccounting.String("consequuntur"),
            From: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("3c49919e-bd1c-4f77-9538-cbbfcdf4ece9"),
                    Name: codataccounting.String("Jodi Schamberger"),
                },
                Amount: codataccounting.Float64(7140.11),
                Currency: codataccounting.String("modi"),
            },
            ID: codataccounting.String("2c330496-17cb-471d-9c25-0b60c751d2ae"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("expedita"),
            SourceModifiedDate: codataccounting.String("necessitatibus"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "tempore": map[string]interface{}{
                        "rem": "consequuntur",
                        "molestias": "officiis",
                        "qui": "vel",
                    },
                    "aliquam": map[string]interface{}{
                        "ab": "dolorum",
                        "veniam": "officiis",
                    },
                    "minus": map[string]interface{}{
                        "corrupti": "reprehenderit",
                        "a": "quam",
                        "cupiditate": "incidunt",
                    },
                },
            },
            To: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("f04f4144-6f79-43d3-b100-20147cd1b831"),
                    Name: codataccounting.String("Amy Price"),
                },
                Amount: codataccounting.Float64(2218.09),
                Currency: codataccounting.String("voluptates"),
            },
            TrackingCategoryRefs: []shared.TrackingCategoryRef{
                shared.TrackingCategoryRef{
                    ID: "8960a0aa-fc7a-4867-8ba5-00a8f4cb72ed",
                    Name: codataccounting.String("Kara Wilderman"),
                },
                shared.TrackingCategoryRef{
                    ID: "25d55615-8803-4212-b7b5-9b7154642b9e",
                    Name: codataccounting.String("Stella Schroeder"),
                },
                shared.TrackingCategoryRef{
                    ID: "c3d3ca49-1837-4978-88d1-56f01ae36bb8",
                    Name: codataccounting.String("Jose Adams"),
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
        TransferID: "eveniet",
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
        Query: codataccounting.String("ratione"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transfers != nil {
        // handle response
    }
}
```
