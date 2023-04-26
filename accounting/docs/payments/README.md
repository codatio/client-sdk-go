# Payments

## Overview

Payments

### Available Operations

* [CreatePayment](#createpayment) - Create payment
* [GetCreatePaymentsModel](#getcreatepaymentsmodel) - Get create payment model
* [GetPayment](#getpayment) - Get payment
* [ListPayments](#listpayments) - List payments

## CreatePayment

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
                ID: codataccounting.String("cf63b215-186a-4b5e-ba02-2614315d1568"),
                Name: codataccounting.String("Violet McClure"),
            },
            Currency: codataccounting.String("architecto"),
            CurrencyRate: codataccounting.Float64(6396.52),
            CustomerRef: &shared.CustomerRef{
                CompanyName: codataccounting.String("reiciendis"),
                ID: "c7186ff2-0b7a-473d-b40c-a0d7657c1641",
            },
            Date: "distinctio",
            ID: codataccounting.String("bf055271-b251-41dd-a06d-d1b28272bc9c"),
            Lines: []shared.PaymentLine{
                shared.PaymentLine{
                    AllocatedOnDate: codataccounting.String("magni"),
                    Amount: 1257.69,
                    Links: []shared.PaymentLineLink{
                        shared.PaymentLineLink{
                            Amount: codataccounting.Float64(4058.4),
                            CurrencyRate: codataccounting.Float64(5724.81),
                            ID: codataccounting.String("7b1880fc-bb2b-493c-95f6-70bd17848316"),
                            Type: shared.PaymentLinkTypeEnumCreditNote,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("dolor"),
            Note: codataccounting.String("debitis"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: codataccounting.String("eb3b6e24-1c31-4099-8366-3c66dcbb7df6"),
                Name: codataccounting.String("Gerardo Baumbach"),
            },
            Reference: codataccounting.String("totam"),
            SourceModifiedDate: codataccounting.String("distinctio"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "aperiam": map[string]interface{}{
                        "recusandae": "eaque",
                        "nihil": "dicta",
                        "adipisci": "molestiae",
                    },
                    "in": map[string]interface{}{
                        "repellendus": "saepe",
                        "non": "a",
                    },
                },
            },
            TotalAmount: codataccounting.Float64(9140.93),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(878742),
    }

    res, err := s.Payments.CreatePayment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatePaymentResponse != nil {
        // handle response
    }
}
```

## GetCreatePaymentsModel

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

    res, err := s.Payments.GetCreatePaymentsModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## GetPayment

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
        PaymentID: "quae",
    }

    res, err := s.Payments.GetPayment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Payment != nil {
        // handle response
    }
}
```

## ListPayments

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
        Query: codataccounting.String("doloremque"),
    }

    res, err := s.Payments.ListPayments(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Payments != nil {
        // handle response
    }
}
```
