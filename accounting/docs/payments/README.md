# Payments

## Overview

Payments

### Available Operations

* [Create](#create) - Create payment
* [Get](#get) - Get payment
* [GetCreateModel](#getcreatemodel) - Get create payment model
* [List](#list) - List payments

## Create

Posts a new payment to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create payment model](https://docs.codat.io/accounting-api#/operations/get-create-payments-model).

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=payments) for integrations that support creating payments.

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
    res, err := s.Payments.Create(ctx, operations.CreatePaymentRequest{
        Payment: &shared.Payment{
            AccountRef: &shared.AccountRef{
                ID: codataccounting.String("def9c765-dfd7-4354-a5cb-94977017a262"),
                Name: codataccounting.String("Monica Reinger"),
            },
            Currency: codataccounting.String("eum"),
            CurrencyRate: codataccounting.Float64(7979.17),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("laborum"),
                ID: "4e999828-79de-4fc0-b239-606cf90ad989",
            },
            Date: "recusandae",
            ID: codataccounting.String("1a34715a-cda0-444f-aaed-6e13a620e2e9"),
            Lines: []shared.PaymentLine{
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("totam"),
                    Amount: 7539.55,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(3390.23),
                            CurrencyRate: codataccounting.Float64(7169.63),
                            ID: codataccounting.String("0486cf39-8a0b-4374-a17d-d95ce3044be4"),
                            Type: shared.PaymentLinkTypeDiscount,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(7139.18),
                            CurrencyRate: codataccounting.Float64(2076.75),
                            ID: codataccounting.String("b31cb503-c314-40d8-b72c-535e1dd6bf64"),
                            Type: shared.PaymentLinkTypePaymentOnAccount,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("non"),
            Note: codataccounting.String("quis"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: codataccounting.String("4e9831e7-95f0-4e57-b54e-bf2d2b46097e"),
                Name: codataccounting.String("Carlton Grady"),
            },
            Reference: codataccounting.String("voluptatum"),
            SourceModifiedDate: codataccounting.String("quibusdam"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "earum": map[string]interface{}{
                        "sit": "cumque",
                        "quibusdam": "quibusdam",
                    },
                    "inventore": map[string]interface{}{
                        "enim": "perferendis",
                        "soluta": "tenetur",
                        "ipsam": "dolorum",
                    },
                    "ipsa": map[string]interface{}{
                        "soluta": "impedit",
                        "quas": "facilis",
                        "quam": "blanditiis",
                        "commodi": "eaque",
                    },
                    "similique": map[string]interface{}{
                        "voluptates": "similique",
                        "autem": "nobis",
                        "laboriosam": "non",
                        "corporis": "ab",
                    },
                },
            },
            TotalAmount: codataccounting.Float64(1991.9),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(179221),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatePaymentResponse != nil {
        // handle response
    }
}
```

## Get

Get payment

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
    res, err := s.Payments.Get(ctx, operations.GetPaymentRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PaymentID: "repellendus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Payment != nil {
        // handle response
    }
}
```

## GetCreateModel

Get create payment model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=payments) for integrations that support creating payments.

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
    res, err := s.Payments.GetCreateModel(ctx, operations.GetCreatePaymentsModelRequest{
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

Gets the latest payments for a company, with pagination

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
    res, err := s.Payments.List(ctx, operations.ListPaymentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("ipsam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Payments != nil {
        // handle response
    }
}
```
