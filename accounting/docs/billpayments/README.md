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
                ID: codataccounting.String("8502a55e-7f73-4bc8-85e3-20a319f4badf"),
                Name: codataccounting.String("Jesus Kreiger"),
            },
            Currency: codataccounting.String("animi"),
            CurrencyRate: codataccounting.Float64(5287.42),
            Date: "aliquid",
            ID: codataccounting.String("3d5a8e00-d108-4045-8823-7f342676cffa"),
            Lines: []shared.BillPaymentLine{
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("facilis"),
                    Amount: 7594.51,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(1782.01),
                            CurrencyRate: codataccounting.Float64(2993.79),
                            ID: codataccounting.String("26665816-ddca-48ef-91fc-b4c593ec12cd"),
                            Type: shared.BillPaymentLineLinkTypePaymentOnAccount,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(6615.78),
                            CurrencyRate: codataccounting.Float64(8409.92),
                            ID: codataccounting.String("0ec7afed-bd80-4df4-88a4-7f9390c58880"),
                            Type: shared.BillPaymentLineLinkTypeBillPayment,
                        },
                    },
                },
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("quas"),
                    Amount: 2184.13,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(6834.9),
                            CurrencyRate: codataccounting.Float64(7047.32),
                            ID: codataccounting.String("f9ef3ffd-d9f7-4f07-9af4-d35724cdb0f4"),
                            Type: shared.BillPaymentLineLinkTypeManualJournal,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(1776.51),
                            CurrencyRate: codataccounting.Float64(5556.63),
                            ID: codataccounting.String("1187d568-44ed-4ed8-9a90-65e628bdfc20"),
                            Type: shared.BillPaymentLineLinkTypeBill,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(1553.71),
                            CurrencyRate: codataccounting.Float64(6937.24),
                            ID: codataccounting.String("6c879923-b7e1-4358-8f7a-e12c6891f82c"),
                            Type: shared.BillPaymentLineLinkTypeDiscount,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(1171.42),
                            CurrencyRate: codataccounting.Float64(810.53),
                            ID: codataccounting.String("57172305-377d-4cfa-89df-975e35668609"),
                            Type: shared.BillPaymentLineLinkTypeUnlinked,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("itaque"),
            Note: codataccounting.String("Bill Payment against bill c13e37b6-dfaa-4894-b3be-9fe97bda9f44"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: codataccounting.String("9c3ddc5f-111d-4ea1-826d-541a4d190feb"),
                Name: codataccounting.String("Evelyn Kuhlman MD"),
            },
            Reference: codataccounting.String("placeat"),
            SourceModifiedDate: codataccounting.String("quod"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "sit": map[string]interface{}{
                        "distinctio": "distinctio",
                        "assumenda": "illum",
                        "soluta": "magnam",
                        "laudantium": "tempora",
                    },
                    "esse": map[string]interface{}{
                        "corrupti": "reiciendis",
                    },
                    "facilis": map[string]interface{}{
                        "repudiandae": "amet",
                        "natus": "ab",
                    },
                    "officiis": map[string]interface{}{
                        "rerum": "placeat",
                        "ab": "ad",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "8c4c4e54-599e-4a34-a260-e9b200ce78a1",
                SupplierName: codataccounting.String("expedita"),
            },
            TotalAmount: codataccounting.Float64(1329.54),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(842921),
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
        BillPaymentID: "quos",
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
        BillPaymentID: "maiores",
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
        Query: codataccounting.String("quidem"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillPayments != nil {
        // handle response
    }
}
```
