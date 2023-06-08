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

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=payments) to see which integrations support this endpoint.

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
    res, err := s.Payments.Create(ctx, operations.CreatePaymentRequest{
        Payment: &shared.Payment{
            AccountRef: &shared.AccountRef{
                ID: codataccounting.String("6b786b3d-37bd-4204-a1f3-40bb36f677a4"),
                Name: codataccounting.String("Zachary Bode"),
            },
            Currency: codataccounting.String("GBP"),
            CurrencyRate: codataccounting.Float64(2424.79),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("quam"),
                ID: "49028488-26bb-403c-bfd2-25e47871a88e",
            },
            Date: "2022-10-23T00:00:00.000Z",
            ID: codataccounting.String("72a2d4af-4158-4ac2-90f0-f58c3b87b470"),
            Lines: []shared.PaymentLine{
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Amount: 8220.93,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(8291.15),
                            CurrencyRate: codataccounting.Float64(5697.11),
                            ID: codataccounting.String("8e9d82c5-e306-4f55-b6f5-cdeb0286d0bc"),
                            Type: shared.PaymentLinkTypeCreditNote,
                        },
                    },
                },
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Amount: 6889.3,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(5572.63),
                            CurrencyRate: codataccounting.Float64(6607.69),
                            ID: codataccounting.String("b378f2fc-ff81-4ddf-be08-8f74ef54c921"),
                            Type: shared.PaymentLinkTypeCreditNote,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("atque"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: codataccounting.String("926313bb-6fc2-4c8d-a701-096b66ad6e3e"),
                Name: codataccounting.String("Freda Marquardt"),
            },
            Reference: codataccounting.String("libero"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "perferendis": map[string]interface{}{
                        "non": "quaerat",
                    },
                    "mollitia": map[string]interface{}{
                        "sunt": "id",
                    },
                },
            },
            TotalAmount: codataccounting.Float64(6613.56),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(67786),
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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Payments.Get(ctx, operations.GetPaymentRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PaymentID: "nulla",
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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
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
