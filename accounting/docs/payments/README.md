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
                ID: codataccounting.String("61e91500-323b-42c0-9b92-4771f5669e5b"),
                Name: codataccounting.String("Tricia Sawayn"),
            },
            Currency: codataccounting.String("magni"),
            CurrencyRate: codataccounting.Float64(4374.89),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("ea"),
                ID: "49d84eb9-e4cf-4d22-b6e0-b88fb87d6fa5",
            },
            Date: "cum",
            ID: codataccounting.String("6e8dbf81-2f83-4b1c-a6a9-ffc561929cca"),
            Lines: []shared.PaymentLine{
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("nemo"),
                    Amount: 3864.41,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(6814.58),
                            CurrencyRate: codataccounting.Float64(977.35),
                            ID: codataccounting.String("395918da-1d48-4e78-a3cf-8e1143da9308"),
                            Type: shared.PaymentLinkTypePaymentOnAccount,
                        },
                    },
                },
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("magni"),
                    Amount: 4682.52,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(137),
                            CurrencyRate: codataccounting.Float64(5312.36),
                            ID: codataccounting.String("af221844-39b3-4de8-b56c-cce470cd2147"),
                            Type: shared.PaymentLinkTypePayment,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(4308.75),
                            CurrencyRate: codataccounting.Float64(9267.48),
                            ID: codataccounting.String("6152cf01-d0d8-4c3a-8b9a-5bf935dfe974"),
                            Type: shared.PaymentLinkTypeDiscount,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(6309.83),
                            CurrencyRate: codataccounting.Float64(2828),
                            ID: codataccounting.String("b1e9c097-eda6-4234-82e1-a9237e9984c8"),
                            Type: shared.PaymentLinkTypeUnknown,
                        },
                    },
                },
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("facilis"),
                    Amount: 2810.64,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(5745.91),
                            CurrencyRate: codataccounting.Float64(9051.54),
                            ID: codataccounting.String("891923c1-8ca8-4d69-8568-9214fa20207e"),
                            Type: shared.PaymentLinkTypeInvoice,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(9643.29),
                            CurrencyRate: codataccounting.Float64(6295.82),
                            ID: codataccounting.String("e038cd7f-1bc2-4cab-af7f-c2ccba4bef0d"),
                            Type: shared.PaymentLinkTypeDiscount,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("nisi"),
            Note: codataccounting.String("voluptatum"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: codataccounting.String("eaedb2ee-70be-4069-bb36-add704080e0a"),
                Name: codataccounting.String("Lorene Schneider"),
            },
            Reference: codataccounting.String("animi"),
            SourceModifiedDate: codataccounting.String("ullam"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "aperiam": map[string]interface{}{
                        "aliquam": "soluta",
                    },
                    "inventore": map[string]interface{}{
                        "ut": "sint",
                    },
                    "sint": map[string]interface{}{
                        "eius": "ratione",
                    },
                },
            },
            TotalAmount: codataccounting.Float64(6253.92),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(979832),
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
        PaymentID: "mollitia",
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
        Query: codataccounting.String("suscipit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Payments != nil {
        // handle response
    }
}
```
