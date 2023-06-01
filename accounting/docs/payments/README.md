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
                ID: codataccounting.String("aa558a65-e208-4301-aca3-4bb87d4f6212"),
                Name: codataccounting.String("Ms. Kristi Jast"),
            },
            Currency: codataccounting.String("architecto"),
            CurrencyRate: codataccounting.Float64(3766.39),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("voluptatem"),
                ID: "6294514c-3db9-4ca9-b38b-d2be87870349",
            },
            Date: "adipisci",
            ID: codataccounting.String("f49aa846-5a32-4832-b9b7-19d1cea673d8"),
            Lines: []shared.PaymentLine{
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("itaque"),
                    Amount: 1984.56,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(1976.2),
                            CurrencyRate: codataccounting.Float64(3400.87),
                            ID: codataccounting.String("e49a3135-778c-4e54-8acb-0e3ea975045b"),
                            Type: shared.PaymentLinkTypePayment,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(7977.88),
                            CurrencyRate: codataccounting.Float64(9916.87),
                            ID: codataccounting.String("63b21518-6ab5-4e3a-8226-14315d156829"),
                            Type: shared.PaymentLinkTypeRefund,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(9207.5),
                            CurrencyRate: codataccounting.Float64(3778.77),
                            ID: codataccounting.String("1afc7186-ff20-4b7a-b3df-40ca0d7657c1"),
                            Type: shared.PaymentLinkTypeOther,
                        },
                    },
                },
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("eius"),
                    Amount: 698.78,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(7083.87),
                            CurrencyRate: codataccounting.Float64(9823),
                            ID: codataccounting.String("055271b2-511d-4d60-add1-b28272bc9c32"),
                            Type: shared.PaymentLinkTypeUnlinked,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(1125.17),
                            CurrencyRate: codataccounting.Float64(4058.4),
                            ID: codataccounting.String("97b1880f-cbb2-4b93-815f-670bd1784831"),
                            Type: shared.PaymentLinkTypeOther,
                        },
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(3516.07),
                            CurrencyRate: codataccounting.Float64(2234.48),
                            ID: codataccounting.String("eeb3b6e2-41c3-4109-9836-63c66dcbb7df"),
                            Type: shared.PaymentLinkTypeCreditNote,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("porro"),
            Note: codataccounting.String("soluta"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: codataccounting.String("09c8b408-e071-4377-8de4-fee101d9780a"),
                Name: codataccounting.String("Laura Schimmel"),
            },
            Reference: codataccounting.String("rerum"),
            SourceModifiedDate: codataccounting.String("provident"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "perferendis": map[string]interface{}{
                        "accusantium": "possimus",
                        "vel": "minus",
                    },
                    "blanditiis": map[string]interface{}{
                        "quia": "similique",
                        "ipsam": "a",
                        "alias": "perferendis",
                    },
                },
            },
            TotalAmount: codataccounting.Float64(1333.46),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(152643),
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
        PaymentID: "sit",
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
        Query: codataccounting.String("esse"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Payments != nil {
        // handle response
    }
}
```
