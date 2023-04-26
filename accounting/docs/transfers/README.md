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

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers) for integrations that support creating transfers.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CreateTransferRequest{
        Transfer: &shared.Transfer{
            ContactRef: &shared.TransferContactRef{
                DataType: codataccounting.String("laborum"),
                ID: "9ffc5619-29cc-4a95-a0a1-395918da1d48",
            },
            Date: codataccounting.String("recusandae"),
            DepositedRecordRefs: []string{
                "quas",
                "officiis",
            },
            Description: codataccounting.String("ipsum"),
            From: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("cf8e1143-da93-408b-a7a0-8af22184439b"),
                    Name: codataccounting.String("Desiree Walsh"),
                },
                Amount: codataccounting.Float64(3395.66),
                Currency: codataccounting.String("eum"),
            },
            ID: codataccounting.String("ccce470c-d214-47b6-a615-2cf01d0d8c3a"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("magnam"),
            SourceModifiedDate: codataccounting.String("facilis"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "laborum": map[string]interface{}{
                        "quidem": "repellat",
                        "molestias": "amet",
                    },
                    "veniam": map[string]interface{}{
                        "voluptatibus": "vero",
                        "provident": "iure",
                        "incidunt": "repellat",
                        "similique": "ut",
                    },
                    "tempore": map[string]interface{}{
                        "voluptates": "excepturi",
                    },
                },
            },
            To: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("c097eda6-2344-42e1-a923-7e9984c80b47"),
                    Name: codataccounting.String("Bert Lind V"),
                },
                Amount: codataccounting.Float64(1752.16),
                Currency: codataccounting.String("dolorem"),
            },
            TrackingCategoryRefs: []shared.TrackingCategoryRef{
                shared.TrackingCategoryRef{
                    ID: "18ca8d69-c568-4921-8fa2-0207e4fae038",
                    Name: codataccounting.String("Carroll Klocko DDS"),
                },
                shared.TrackingCategoryRef{
                    ID: "c2cabaf7-fc2c-4cba-8bef-0df68eaedb2e",
                    Name: codataccounting.String("Darryl Altenwerth"),
                },
                shared.TrackingCategoryRef{
                    ID: "069fb36a-dd70-4408-8e0a-3fc73a5a034b",
                    Name: codataccounting.String("Rebecca Graham"),
                },
                shared.TrackingCategoryRef{
                    ID: "243afa69-87a4-472b-b09a-153e22301068",
                    Name: codataccounting.String("Tracy Monahan"),
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Transfers.Create(ctx, req)
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetTransferRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TransferID: "ipsa",
    }

    res, err := s.Transfers.Get(ctx, req)
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
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers) for integrations that support creating transfers.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetCreateTransfersModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Transfers.GetCreateModel(ctx, req)
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListTransfersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("perspiciatis"),
    }

    res, err := s.Transfers.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Transfers != nil {
        // handle response
    }
}
```
