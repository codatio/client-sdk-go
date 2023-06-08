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
                ID: codataccounting.String("99171b51-c1bd-4b1c-b4b8-88ebdfc4ccca"),
                Name: codataccounting.String("Tracy Reinger"),
            },
            Currency: codataccounting.String("EUR"),
            CurrencyRate: codataccounting.Float64(8022.16),
            Date: "2022-10-23T00:00:00.000Z",
            ID: codataccounting.String("3d5a8e00-d108-4045-8823-7f342676cffa"),
            Lines: []shared.BillPaymentLine{
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Amount: 8662.92,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(9274.9),
                            CurrencyRate: codataccounting.Float64(837.91),
                            ID: codataccounting.String("0873e42b-006d-4678-878b-a8581a58208c"),
                            Type: shared.BillPaymentLineLinkTypeOther,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(2864.53),
                            CurrencyRate: codataccounting.Float64(9580.68),
                            ID: codataccounting.String("efa9c95f-2eac-4556-9d30-7cfee81206e2"),
                            Type: shared.BillPaymentLineLinkTypeBillPayment,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(1227.44),
                            CurrencyRate: codataccounting.Float64(2001.9),
                            ID: codataccounting.String("fa4a41c4-80d3-4f21-b2af-03102d514f4c"),
                            Type: shared.BillPaymentLineLinkTypeRefund,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(4044.22),
                            CurrencyRate: codataccounting.Float64(9980.23),
                            ID: codataccounting.String("18bf9621-a6a4-4f77-a87e-e3e4be752c65"),
                            Type: shared.BillPaymentLineLinkTypeRefund,
                        },
                    },
                },
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Amount: 2848.85,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(1039.88),
                            CurrencyRate: codataccounting.Float64(5069.66),
                            ID: codataccounting.String("e3bb91c8-d975-4e0e-8419-d8f84f144f3e"),
                            Type: shared.BillPaymentLineLinkTypeUnknown,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(4587.23),
                            CurrencyRate: codataccounting.Float64(8913.02),
                            ID: codataccounting.String("dcc4aa5f-3cab-4d90-9a97-2e056728227b"),
                            Type: shared.BillPaymentLineLinkTypeUnlinked,
                        },
                    },
                },
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Amount: 1998.79,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(5751.39),
                            CurrencyRate: codataccounting.Float64(2913.89),
                            ID: codataccounting.String("70bf7a4f-a87c-4f53-9a6f-ae54ebf60c32"),
                            Type: shared.BillPaymentLineLinkTypeUnlinked,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("Bill Payment against bill c13e37b6-dfaa-4894-b3be-9fe97bda9f44"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: codataccounting.String("023b75d2-367f-4e1a-8cc8-df79f0a396d9"),
                Name: codataccounting.String("Leticia Donnelly"),
            },
            Reference: codataccounting.String("expedita"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "sunt": map[string]interface{}{
                        "nulla": "maiores",
                        "distinctio": "mollitia",
                    },
                    "impedit": map[string]interface{}{
                        "et": "quas",
                        "blanditiis": "cum",
                        "dicta": "impedit",
                        "tempora": "eveniet",
                    },
                    "repudiandae": map[string]interface{}{
                        "impedit": "quas",
                    },
                    "impedit": map[string]interface{}{
                        "eligendi": "recusandae",
                        "ex": "beatae",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "1feeb1c7-cbdb-46ee-8743-78ba25317747",
                SupplierName: codataccounting.String("nulla"),
            },
            TotalAmount: codataccounting.Float64(1329.54),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(771406),
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
        BillPaymentID: "cupiditate",
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
        BillPaymentID: "illo",
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
        Query: codataccounting.String("exercitationem"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillPayments != nil {
        // handle response
    }
}
```
