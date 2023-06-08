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
                ID: codataccounting.String("efbd02ba-e0be-42d7-8225-9e3ea4b5197f"),
                Name: codataccounting.String("Steve Fritsch"),
            },
            Currency: codataccounting.String("EUR"),
            CurrencyRate: codataccounting.Float64(6378.56),
            Date: "2022-10-23T00:00:00.000Z",
            ID: codataccounting.String("3d5a8e00-d108-4045-8823-7f342676cffa"),
            Lines: []shared.BillPaymentLine{
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Amount: 3591.11,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(7088.98),
                            CurrencyRate: codataccounting.Float64(5326.69),
                            ID: codataccounting.String("95c537c6-454e-4fb0-b348-96c3ca5acfbe"),
                            Type: shared.BillPaymentLineLinkTypeUnlinked,
                        },
                    },
                },
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Amount: 8659.46,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(4413.21),
                            CurrencyRate: codataccounting.Float64(452.34),
                            ID: codataccounting.String("75779291-77de-4ac6-86ec-b573409e3eb1"),
                            Type: shared.BillPaymentLineLinkTypeDiscount,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(3344.74),
                            CurrencyRate: codataccounting.Float64(6592.68),
                            ID: codataccounting.String("2b12eb07-f116-4db9-9545-fc95fa88970e"),
                            Type: shared.BillPaymentLineLinkTypeUnlinked,
                        },
                    },
                },
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Amount: 6144.38,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(7316.34),
                            CurrencyRate: codataccounting.Float64(7255.74),
                            ID: codataccounting.String("30fcb33e-a055-4b19-bcd4-4e2f52d82d35"),
                            Type: shared.BillPaymentLineLinkTypeUnknown,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(2426.37),
                            CurrencyRate: codataccounting.Float64(7057.1),
                            ID: codataccounting.String("b6f48b65-6bcd-4b35-bf2e-4b27537a8cd9"),
                            Type: shared.BillPaymentLineLinkTypeManualJournal,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(4956.17),
                            CurrencyRate: codataccounting.Float64(2201.04),
                            ID: codataccounting.String("19c177d5-25f7-47b1-94ee-b52ff785fc37"),
                            Type: shared.BillPaymentLineLinkTypeBillPayment,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(755.66),
                            CurrencyRate: codataccounting.Float64(2902.48),
                            ID: codataccounting.String("d4c98e0c-2bb8-49eb-b5da-d636c600503d"),
                            Type: shared.BillPaymentLineLinkTypeBillPayment,
                        },
                    },
                },
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Amount: 7368.53,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(976.76),
                            CurrencyRate: codataccounting.Float64(1181.26),
                            ID: codataccounting.String("80f739ae-9e05-47eb-809e-2810331f3981"),
                            Type: shared.BillPaymentLineLinkTypeManualJournal,
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
                ID: codataccounting.String("c700b607-f3c9-43c7-bb9d-a3f2ceda7e23"),
                Name: codataccounting.String("Roy Christiansen"),
            },
            Reference: codataccounting.String("non"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "hic": map[string]interface{}{
                        "delectus": "non",
                        "distinctio": "in",
                        "exercitationem": "labore",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "4e472e80-2857-4a5b-8046-3a7d575f1400",
                SupplierName: codataccounting.String("eveniet"),
            },
            TotalAmount: codataccounting.Float64(1329.54),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(483753),
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
        BillPaymentID: "commodi",
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
        BillPaymentID: "numquam",
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
        Query: codataccounting.String("dolorum"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillPayments != nil {
        // handle response
    }
}
```
