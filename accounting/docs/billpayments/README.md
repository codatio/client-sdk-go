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
    res, err := s.BillPayments.Create(ctx, operations.CreateBillPaymentRequest{
        BillPayment: &shared.BillPayment{
            AccountRef: &shared.AccountRef{
                ID: codataccounting.String("57eb809e-2810-4331-b398-1d4c700b607f"),
                Name: codataccounting.String("Francis McKenzie"),
            },
            Currency: codataccounting.String("dignissimos"),
            CurrencyRate: codataccounting.Float64(2358.34),
            Date: "soluta",
            ID: codataccounting.String("3d5a8e00-d108-4045-8823-7f342676cffa"),
            Lines: []shared.BillPaymentLine{
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("temporibus"),
                    Amount: 6396.22,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(9485.41),
                            CurrencyRate: codataccounting.Float64(1339.49),
                            ID: codataccounting.String("ceda7e23-f225-4741-9faf-4b7544e472e8"),
                            Type: shared.BillPaymentLineLinkTypeEnumUnknown,
                        },
                    },
                },
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("odit"),
                    Amount: 5358.02,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(4527.3),
                            CurrencyRate: codataccounting.Float64(6267.07),
                            ID: codataccounting.String("5b40463a-7d57-45f1-800e-764ad7334ec1"),
                            Type: shared.BillPaymentLineLinkTypeEnumRefund,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(4813.75),
                            CurrencyRate: codataccounting.Float64(5580.51),
                            ID: codataccounting.String("1b36a080-88d1-400e-bada-200ef0422eb2"),
                            Type: shared.BillPaymentLineLinkTypeEnumUnlinked,
                        },
                    },
                },
                shared.BillPaymentLine{
                    AllocatedOnDate: codataccounting.String("aliquid"),
                    Amount: 2646.49,
                    Links: []shared.BillPaymentLineLink{
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(9750.95),
                            CurrencyRate: codataccounting.Float64(5629.48),
                            ID: codataccounting.String("ab8366c7-23ff-4da9-a06b-ee4825c1fc0e"),
                            Type: shared.BillPaymentLineLinkTypeEnumUnlinked,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(1002.51),
                            CurrencyRate: codataccounting.Float64(3178.98),
                            ID: codataccounting.String("c80bff91-8544-4ec4-adef-cce8f1977773"),
                            Type: shared.BillPaymentLineLinkTypeEnumManualJournal,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(4235.88),
                            CurrencyRate: codataccounting.Float64(2086.83),
                            ID: codataccounting.String("562a7b40-8f05-4e3d-88fd-af313a1f5fd9"),
                            Type: shared.BillPaymentLineLinkTypeEnumBill,
                        },
                        shared.BillPaymentLineLink{
                            Amount: codataccounting.Float64(1280.21),
                            CurrencyRate: codataccounting.Float64(3684.91),
                            ID: codataccounting.String("9c0b36f2-5ea9-444f-bb75-6c11f6c37a51"),
                            Type: shared.BillPaymentLineLinkTypeEnumUnlinked,
                        },
                    },
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("aliquid"),
            Note: codataccounting.String("Bill Payment against bill c13e37b6-dfaa-4894-b3be-9fe97bda9f44"),
            PaymentMethodRef: &shared.PaymentMethodRef{
                ID: codataccounting.String("243835bb-c05a-423a-85ce-fc5fde10a0ce"),
                Name: codataccounting.String("Mildred Kautzer"),
            },
            Reference: codataccounting.String("ullam"),
            SourceModifiedDate: codataccounting.String("architecto"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "perferendis": map[string]interface{}{
                        "provident": "cumque",
                    },
                },
            },
            SupplierRef: &shared.SupplierRef{
                ID: "6dc5e347-6279-49bf-bbe6-949fb2bb4eca",
                SupplierName: codataccounting.String("saepe"),
            },
            TotalAmount: codataccounting.Float64(1329.54),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(423054),
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
    res, err := s.BillPayments.Delete(ctx, operations.DeleteBillPaymentRequest{
        BillPaymentID: "quo",
        CompanyID: "nesciunt",
        ConnectionID: "illum",
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
    res, err := s.BillPayments.Get(ctx, operations.GetBillPaymentsRequest{
        BillPaymentID: "nemo",
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
    res, err := s.BillPayments.List(ctx, operations.ListBillPaymentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("illum"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillPayments != nil {
        // handle response
    }
}
```
