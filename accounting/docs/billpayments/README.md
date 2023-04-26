# BillPayments

## Overview

Bill payments

### Available Operations

* [CreateBillPayment](#createbillpayment) - Create bill payments
* [DeleteBillPayment](#deletebillpayment) - Delete bill payment
* [GetBillPayments](#getbillpayments) - Get bill payment
* [GetCreateBillPaymentsModel](#getcreatebillpaymentsmodel) - Get create bill payment model
* [ListBillPayments](#listbillpayments) - List bill payments

## CreateBillPayment

Posts a new bill payment to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create bill payment model](https://docs.codat.io/accounting-api#/operations/get-create-billPayments-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billPayments) for integrations that support creating bill payments.

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
    req := operations.CreateBillPaymentRequest{
        BillPayment: &shared.BillPayment{
            AccountRef: &shared.AccountRef{
                ID: codataccounting.String("b7008787-5614-43f5-a6c9-8b55554080d4"),
                Name: codataccounting.String("Celia Rolfson"),
            },
            Currency: codataccounting.String("impedit"),
            CurrencyRate: codataccounting.Float64(3947.24),
            Date: "cumque",
            ID: codataccounting.String("bd6b5f3e-c909-4304-b926-bad2553819b4"),
            Lines: []shared.BillPaymentLine{
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("eius"),
                    Amount: 7137.18,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(9154.08),
                            CurrencyRate: codataccounting.Float64(8227.11),
                            ID: codataccounting.String("20e56248-fff6-439a-910a-bdcab6267669"),
                            Type: shared.BillPaymentLineLinkTypeEnumOther,
                        },
                    },
                },
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("accusamus"),
                    Amount: 682.92,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(7903.05),
                            CurrencyRate: codataccounting.Float64(535.29),
                            ID: codataccounting.String("0221b335-d89a-4cb3-acfd-a8d0c549ef03"),
                            Type: shared.BillPaymentLineLinkTypeEnumUnknown,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(437.15),
                            CurrencyRate: codataccounting.Float64(2737.32),
                            ID: codataccounting.String("978a61fa-1cf2-4068-8f77-c1ffc71dca16"),
                            Type: shared.BillPaymentLineLinkTypeEnumBill,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(9824.09),
                            CurrencyRate: codataccounting.Float64(1530.97),
                            ID: codataccounting.String("a3c80a97-ff33-44cd-9f85-7a9e61876c6a"),
                            Type: shared.BillPaymentLineLinkTypeEnumPaymentOnAccount,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(1265.12),
                            CurrencyRate: codataccounting.Float64(928.51),
                            ID: codataccounting.String("d29dfc94-d6fe-4cd7-9939-0066a6d2d000"),
                            Type: shared.BillPaymentLineLinkTypeEnumBill,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("ullam"),
            Note: codataccounting.String("quis"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: codataccounting.String("338cec08-6fa2-41e9-952c-b3119167b8e3"),
                Name: codataccounting.String("Julian Stanton I"),
            },
            Reference: codataccounting.String("quaerat"),
            SourceModifiedDate: codataccounting.String("consequatur"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "repellendus": map[string]interface{}{
                        "quibusdam": "consectetur",
                        "voluptas": "quaerat",
                    },
                    "earum": map[string]interface{}{
                        "assumenda": "dolore",
                        "enim": "ullam",
                        "perspiciatis": "alias",
                        "ex": "quibusdam",
                    },
                    "dicta": map[string]interface{}{
                        "commodi": "neque",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "d48e935c-2c9e-481f-b0be-3e43202d7216",
                SupplierName: codataccounting.String("ad"),
            },
            TotalAmount: codataccounting.Float64(4531.98),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(426594),
    }

    res, err := s.BillPayments.CreateBillPayment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBillPaymentResponse != nil {
        // handle response
    }
}
```

## DeleteBillPayment

Deletes a bill payment from the accounting package for a given company.

> **Supported Integrations**
> 
> This functionality is currently only supported for our Oracle NetSuite integration. Check out our [public roadmap](https://portal.productboard.com/codat/7-public-product-roadmap/tabs/46-accounting-api) to see what we're building next, and to submit ideas for new features.

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
    req := operations.DeleteBillPaymentRequest{
        BillPaymentID: "minima",
        CompanyID: "sit",
        ConnectionID: "vel",
    }

    res, err := s.BillPayments.DeleteBillPayment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOperationSummary != nil {
        // handle response
    }
}
```

## GetBillPayments

Get a bill payment

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
    req := operations.GetBillPaymentsRequest{
        BillPaymentID: "laboriosam",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.BillPayments.GetBillPayments(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BillPayment != nil {
        // handle response
    }
}
```

## GetCreateBillPaymentsModel

Get create bill payment model.

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billPayments) for integrations that support creating and deleting bill payments.

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
    req := operations.GetCreateBillPaymentsModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.BillPayments.GetCreateBillPaymentsModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## ListBillPayments

Gets the latest billPayments for a company, with pagination

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
    req := operations.ListBillPaymentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("quaerat"),
    }

    res, err := s.BillPayments.ListBillPayments(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BillPayments != nil {
        // handle response
    }
}
```
