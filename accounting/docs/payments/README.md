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
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=payments) for integrations that support creating payments.

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
    req := operations.CreatePaymentRequest{
        Payment: &shared.Payment{
            AccountRef: &shared.AccountRef{
                ID: codataccounting.String("2de7b356-2201-4a6a-ab4a-e7b1a5b908d4"),
                Name: codataccounting.String("Jeffery Aufderhar"),
            },
            Currency: codataccounting.String("quae"),
            CurrencyRate: codataccounting.Float64(6765.76),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("fuga"),
                ID: "35d4a839-f03b-4ab7-bb91-8f0313984507",
            },
            Date: "officiis",
            ID: codataccounting.String("0e39c7e2-3ecb-4060-8652-e23a3d6c657e"),
            Lines: []shared.PaymentLine{
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("quibusdam"),
                    Amount: 8936.05,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(9387.2),
                            CurrencyRate: codataccounting.Float64(4758.76),
                            ID: codataccounting.String("f002d198-6aa9-49d3-a1d3-2329e45837e8"),
                            Type: shared.PaymentLinkTypeEnumDiscount,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(1859.89),
                            CurrencyRate: codataccounting.Float64(6377.7),
                            ID: codataccounting.String("d6bb10e2-55fd-4c48-8d6e-3308675cbf18"),
                            Type: shared.PaymentLinkTypeEnumCreditNote,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(5604.72),
                            CurrencyRate: codataccounting.Float64(3424.33),
                            ID: codataccounting.String("6a7e82cd-f9d0-4fc2-82c6-66af3c3f5589"),
                            Type: shared.PaymentLinkTypeEnumPaymentOnAccount,
                        },
                    },
                },
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("accusamus"),
                    Amount: 6668.05,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(8213.45),
                            CurrencyRate: codataccounting.Float64(1736.08),
                            ID: codataccounting.String("64e41e2c-a848-422e-913f-6d9d2ad37c30"),
                            Type: shared.PaymentLinkTypeEnumPayment,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(5821.15),
                            CurrencyRate: codataccounting.Float64(328.36),
                            ID: codataccounting.String("77c10b68-7921-463e-a7d4-8860543c0a30"),
                            Type: shared.PaymentLinkTypeEnumInvoice,
                        },
                    },
                },
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("excepturi"),
                    Amount: 7879.41,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(8004.56),
                            CurrencyRate: codataccounting.Float64(9757.5),
                            ID: codataccounting.String("6c0276e7-b21b-4ad9-8d27-43fd6c2a10e6"),
                            Type: shared.PaymentLinkTypeEnumManualJournal,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("odit"),
            Note: codataccounting.String("natus"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: codataccounting.String("78ec256a-5b09-4227-bcc4-7996c977bbc5"),
                Name: codataccounting.String("Jeannie Dibbert"),
            },
            Reference: codataccounting.String("eos"),
            SourceModifiedDate: codataccounting.String("quos"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "blanditiis": map[string]interface{}{
                        "ipsa": "eaque",
                        "quo": "ad",
                    },
                    "atque": map[string]interface{}{
                        "eum": "iusto",
                        "facere": "ea",
                        "sequi": "voluptates",
                        "tempora": "similique",
                    },
                    "officia": map[string]interface{}{
                        "laboriosam": "quos",
                        "aliquam": "vel",
                    },
                },
            },
            TotalAmount: codataccounting.Float64(2546.16),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(321921),
    }

    res, err := s.Payments.Create(ctx, req)
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
    req := operations.GetPaymentRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PaymentID: "odio",
    }

    res, err := s.Payments.Get(ctx, req)
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
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=payments) for integrations that support creating payments.

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
    req := operations.GetCreatePaymentsModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Payments.GetCreateModel(ctx, req)
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
    req := operations.ListPaymentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("omnis"),
    }

    res, err := s.Payments.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Payments != nil {
        // handle response
    }
}
```
