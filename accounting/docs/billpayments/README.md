# BillPayments

## Overview

Bill payments

### Available Operations

* [Create](#create) - Create bill payments
* [Delete](#delete) - Delete bill payment
* [Get](#get) - Get bill payment
* [GetCreateModel](#getcreatemodel) - Get create bill payment model
* [List](#list) - List bill payments

## Create

Posts a new bill payment to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create bill payment model](https://docs.codat.io/accounting-api#/operations/get-create-billPayments-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billPayments) to see which integrations support this endpoint.


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
    res, err := s.BillPayments.Create(ctx, operations.CreateBillPaymentRequest{
        BillPayment: &shared.BillPayment{
            AccountRef: &shared.AccountRef{
                ID: codataccounting.String("ebf60c32-1f02-43b7-9d23-67fe1a0cc8df"),
                Name: codataccounting.String("Miss Daisy Willms"),
            },
            Currency: codataccounting.String("provident"),
            CurrencyRate: codataccounting.Float64(4047.74),
            Date: "repellendus",
            ID: codataccounting.String("3d5a8e00-d108-4045-8823-7f342676cffa"),
            Lines: []shared.BillPaymentLine{
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("alias"),
                    Amount: 7719.31,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(4130.86),
                            CurrencyRate: codataccounting.Float64(2871.41),
                            ID: codataccounting.String("b7c15dfb-ace1-488b-9c4e-e2c8c6ce611f"),
                            Type: shared.BillPaymentLineLinkTypeDiscount,
                        },
                    },
                },
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("vero"),
                    Amount: 6943.94,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(7782.42),
                            CurrencyRate: codataccounting.Float64(4909.66),
                            ID: codataccounting.String("cbdb6eec-7437-48ba-a531-7747dc915ad2"),
                            Type: shared.BillPaymentLineLinkTypeManualJournal,
                        },
                    },
                },
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("dolorum"),
                    Amount: 9983.55,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(8473.08),
                            CurrencyRate: codataccounting.Float64(8450.86),
                            ID: codataccounting.String("6723dc0f-5ae2-4f3a-ab70-0878756143f5"),
                            Type: shared.BillPaymentLineLinkTypePaymentOnAccount,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(4351.42),
                            CurrencyRate: codataccounting.Float64(7876.29),
                            ID: codataccounting.String("98b55554-080d-440b-8acc-6cbd6b5f3ec9"),
                            Type: shared.BillPaymentLineLinkTypeUnknown,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("provident"),
            Note: codataccounting.String("Bill Payment against bill c13e37b6-dfaa-4894-b3be-9fe97bda9f44"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: codataccounting.String("304f926b-ad25-4538-99b4-74b0ed20e562"),
                Name: codataccounting.String("Vickie Welch"),
            },
            Reference: codataccounting.String("autem"),
            SourceModifiedDate: codataccounting.String("nesciunt"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "animi": map[string]interface{}{
                        "beatae": "ipsa",
                        "mollitia": "nam",
                        "assumenda": "quo",
                    },
                    "fuga": map[string]interface{}{
                        "commodi": "fugit",
                        "suscipit": "voluptate",
                        "nisi": "aliquid",
                    },
                    "provident": map[string]interface{}{
                        "accusamus": "ab",
                        "itaque": "quisquam",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "00221b33-5d89-4acb-becf-da8d0c549ef0",
                SupplierName: codataccounting.String("ipsum"),
            },
            TotalAmount: codataccounting.Float64(1329.54),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(367),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBillPaymentResponse != nil {
        // handle response
    }
}
```

## Delete

﻿The _Delete Bill Payments_ endpoint allows you to delete a specified Bill Payment from an accounting platform.

### Process
1. Pass the `{billPaymentId}` to the _Delete Bill Payments_ endpoint and store the `pushOperationKey` returned.
2. Check the status of the delete operation by checking the status of push operation either via
    1. [Push operation webhook](/introduction/webhooks/core-rules-types#push-operation-status-has-changed) (advised),
    2. [Push operation status endpoint](https://docs.codat.io/codat-api#/operations/get-push-operation).

   A `Success` status indicates that the Bill Payment object was deleted from the accounting platform.
3. (Optional) Check that the Bill Payment was deleted from the accounting platform.

### Effect on related objects
Be aware that deleting a Bill Payment from an accounting platform might cause related objects to be modified.

## Integration specifics
Integrations that support soft delete do not permanently delete the object in the accounting platform.

| Integration | Soft Delete | Details                                                                                             |  
|-------------|-------------|-----------------------------------------------------------------------------------------------------|
| Oracle NetSuite   | No          | See [here](/integrations/accounting/netsuite/how-deleting-bill-payments-works) to learn more. |

> **Supported Integrations**
>
> This functionality is currently only supported for our QuickBooks Online abd Oracle NetSuite integrations. Check out our [public roadmap](https://portal.productboard.com/codat/7-public-product-roadmap/tabs/46-accounting-api) to see what we're building next, and to submit ideas for new features.

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
    res, err := s.BillPayments.Delete(ctx, operations.DeleteBillPaymentRequest{
        BillPaymentID: "doloremque",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOperationSummary != nil {
        // handle response
    }
}
```

## Get

﻿Get a bill payment.

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
    res, err := s.BillPayments.Get(ctx, operations.GetBillPaymentsRequest{
        BillPaymentID: "tempora",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillPayment != nil {
        // handle response
    }
}
```

## GetCreateModel

﻿Get create bill payment model.

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billPayments) for integrations that support creating and deleting bill payments.

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
    res, err := s.BillPayments.GetCreateModel(ctx, operations.GetCreateBillPaymentsModelRequest{
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

﻿Gets the latest billPayments for a company, with pagination.

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
    res, err := s.BillPayments.List(ctx, operations.ListBillPaymentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("perspiciatis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillPayments != nil {
        // handle response
    }
}
```
