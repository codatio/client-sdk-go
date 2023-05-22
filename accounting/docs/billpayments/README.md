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

﻿Posts a new bill payment to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create bill payment model](https://docs.codat.io/accounting-api#/operations/get-create-billPayments-model).

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billPayments) for integrations that support creating bill payments.

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
    res, err := s.BillPayments.Create(ctx, operations.CreateBillPaymentRequest{
        BillPayment: &shared.BillPayment{
            AccountRef: &shared.AccountRef{
                ID: codataccounting.String("2af03102-d514-4f4c-86f1-8bf9621a6a4f"),
                Name: codataccounting.String("Tamara O'Kon"),
            },
            Currency: codataccounting.String("eveniet"),
            CurrencyRate: codataccounting.Float64(9351.61),
            Date: "velit",
            ID: codataccounting.String("3d5a8e00-d108-4045-8823-7f342676cffa"),
            Lines: []shared.BillPaymentLine{
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("eius"),
                    Amount: 7019.78,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(4896.85),
                            CurrencyRate: codataccounting.Float64(3734.49),
                            ID: codataccounting.String("2c65b344-18e3-4bb9-9c8d-975e0e8419d8"),
                            Type: shared.BillPaymentLineLinkTypeDiscount,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(5379.46),
                            CurrencyRate: codataccounting.Float64(2649.58),
                            ID: codataccounting.String("f144f3e0-7edc-4c4a-a5f3-cabd905a972e"),
                            Type: shared.BillPaymentLineLinkTypeUnknown,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(3214.73),
                            CurrencyRate: codataccounting.Float64(3927.52),
                            ID: codataccounting.String("728227b2-d309-4470-bf7a-4fa87cf535a6"),
                            Type: shared.BillPaymentLineLinkTypeDiscount,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(6578.62),
                            CurrencyRate: codataccounting.Float64(9259.94),
                            ID: codataccounting.String("54ebf60c-321f-4023-b75d-2367fe1a0cc8"),
                            Type: shared.BillPaymentLineLinkTypeManualJournal,
                        },
                    },
                },
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("maiores"),
                    Amount: 4857.95,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(9609.33),
                            CurrencyRate: codataccounting.Float64(455.1),
                            ID: codataccounting.String("a396d90c-364b-47c1-9dfb-ace188b1c4ee"),
                            Type: shared.BillPaymentLineLinkTypeUnlinked,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(7726.28),
                            CurrencyRate: codataccounting.Float64(5582.83),
                            ID: codataccounting.String("c6ce611f-eeb1-4c7c-bdb6-eec74378ba25"),
                            Type: shared.BillPaymentLineLinkTypeUnlinked,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(1068.06),
                            CurrencyRate: codataccounting.Float64(4810.42),
                            ID: codataccounting.String("747dc915-ad2c-4af5-9d67-23dc0f5ae2f3"),
                            Type: shared.BillPaymentLineLinkTypePaymentOnAccount,
                        },
                    },
                },
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("suscipit"),
                    Amount: 6886.49,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(424.54),
                            CurrencyRate: codataccounting.Float64(201.41),
                            ID: codataccounting.String("87875614-3f5a-46c9-8b55-554080d40bca"),
                            Type: shared.BillPaymentLineLinkTypeRefund,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(7697.89),
                            CurrencyRate: codataccounting.Float64(3947.24),
                            ID: codataccounting.String("cbd6b5f3-ec90-4930-8f92-6bad2553819b"),
                            Type: shared.BillPaymentLineLinkTypeBill,
                        },
                    },
                },
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("voluptate"),
                    Amount: 2611.7,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(463.84),
                            CurrencyRate: codataccounting.Float64(9154.08),
                            ID: codataccounting.String("d20e5624-8fff-4639-a910-abdcab626766"),
                            Type: shared.BillPaymentLineLinkTypeBillPayment,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(3857.6),
                            CurrencyRate: codataccounting.Float64(8815.68),
                            ID: codataccounting.String("1ec00221-b335-4d89-acb3-ecfda8d0c549"),
                            Type: shared.BillPaymentLineLinkTypeManualJournal,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(9380.94),
                            CurrencyRate: codataccounting.Float64(427.63),
                            ID: codataccounting.String("3004978a-61fa-41cf-a068-8f77c1ffc71d"),
                            Type: shared.BillPaymentLineLinkTypeManualJournal,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("fuga"),
            Note: codataccounting.String("Bill Payment against bill c13e37b6-dfaa-4894-b3be-9fe97bda9f44"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: codataccounting.String("163f2a3c-80a9-47ff-b34c-ddf857a9e618"),
                Name: codataccounting.String("Tonya Sauer"),
            },
            Reference: codataccounting.String("quidem"),
            SourceModifiedDate: codataccounting.String("explicabo"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "nulla": map[string]interface{}{
                        "natus": "illum",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "fc94d6fe-cd79-4939-8066-a6d2d0003553",
                SupplierName: codataccounting.String("ratione"),
            },
            TotalAmount: codataccounting.Float64(1329.54),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(555386),
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.BillPayments.Delete(ctx, operations.DeleteBillPaymentRequest{
        BillPaymentID: "maxime",
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.BillPayments.Get(ctx, operations.GetBillPaymentsRequest{
        BillPaymentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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
            AuthHeader: "YOUR_API_KEY_HERE",
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.BillPayments.List(ctx, operations.ListBillPaymentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("recusandae"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillPayments != nil {
        // handle response
    }
}
```
