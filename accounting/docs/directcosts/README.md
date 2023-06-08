# DirectCosts

## Overview

Direct costs

### Available Operations

* [Create](#create) - Create direct cost
* [DownloadAttachment](#downloadattachment) - Download direct cost attachment
* [Get](#get) - Get direct cost
* [GetAttachment](#getattachment) - Get direct cost attachment
* [GetCreateModel](#getcreatemodel) - Get create direct cost model
* [List](#list) - List direct costs
* [ListAttachments](#listattachments) - List direct cost attachments
* [UploadAttachment](#uploadattachment) - Upload direct cost attachment

## Create

Posts a new direct cost to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create direct cost model](https://docs.codat.io/accounting-api#/operations/get-create-directCosts-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts) to see which integrations support this endpoint.

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
    res, err := s.DirectCosts.Create(ctx, operations.CreateDirectCostRequest{
        DirectCost: &shared.DirectCost{
            ContactRef: &shared.ContactRef{
                DataType: codataccounting.String("fugiat"),
                ID: "2cf502ba-fb2c-4bc4-a35d-5e65da028c3e",
            },
            Currency: "USD",
            CurrencyRate: codataccounting.Float64(3186.72),
            ID: codataccounting.String("1a1e30fd-a966-4489-97b7-8673e13a12a6"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.DirectCostLineItem{
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("92494594-487f-45c8-8383-6b86b3cdf641"),
                        Name: codataccounting.String("Kristy Bahringer"),
                    },
                    Description: codataccounting.String("molestias"),
                    DiscountAmount: codataccounting.Float64(9417.1),
                    DiscountPercentage: codataccounting.Float64(6236.93),
                    ItemRef: &shared.ItemRef{
                        ID: "df13f4ee-dbe7-48bf-a068-25894ea763d5",
                        Name: codataccounting.String("Roland Corkery"),
                    },
                    Quantity: 3378.33,
                    SubTotal: codataccounting.Float64(7137.41),
                    TaxAmount: codataccounting.Float64(4626.46),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5144.41),
                        ID: codataccounting.String("5148d6d5-49e5-4635-b33b-c0f970c42fc9"),
                        Name: codataccounting.String("Micheal Larson"),
                    },
                    TotalAmount: codataccounting.Float64(1253.94),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingRecordReference{
                            DataType: codataccounting.String("consequuntur"),
                            ID: codataccounting.String("5e75b796-065c-40ef-a6f9-3b90a1b8c95b"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("vitae"),
                                ID: codataccounting.String("254b739f-4fe7-4721-8d1f-6558c99c722d"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("fugit"),
                                ID: codataccounting.String("bc0f9408-7d9c-4aae-842d-d7caac9b4caa"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("beatae"),
                                ID: codataccounting.String("cfe9e15d-f903-4907-b378-31983d42e54a"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("voluptatum"),
                                ID: codataccounting.String("5466597c-5023-43c1-871d-51aaa6ddf5ab"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "6487c5fc-2b86-42a0-8bef-69e100157630",
                            Name: codataccounting.String("Taylor Paucek"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "fded84a3-5a41-4238-a1a7-35ac26ae33be",
                            Name: codataccounting.String("Miss Terrence Kulas"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "f46bca11-06fe-4965-b711-d08cf88ec9f7",
                            Name: codataccounting.String("Freddie Mayert"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "50a656ed-333b-4b0c-a8aa-65432a986eb7",
                            Name: codataccounting.String("Jonathan Grimes"),
                        },
                    },
                    UnitAmount: 3702.19,
                },
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("64089150-0970-419a-88f8-8ece7bf904e0"),
                        Name: codataccounting.String("Mildred Balistreri"),
                    },
                    Description: codataccounting.String("sequi"),
                    DiscountAmount: codataccounting.Float64(5122.2),
                    DiscountPercentage: codataccounting.Float64(5700.04),
                    ItemRef: &shared.ItemRef{
                        ID: "08162c6b-eb68-4a0f-a57b-7d03a1480f8d",
                        Name: codataccounting.String("Dale Barrows III"),
                    },
                    Quantity: 5877.48,
                    SubTotal: codataccounting.Float64(8550.63),
                    TaxAmount: codataccounting.Float64(5429.86),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(1068),
                        ID: codataccounting.String("0618d97e-1522-4975-90da-80312292cc61"),
                        Name: codataccounting.String("Louis Pacocha Sr."),
                    },
                    TotalAmount: codataccounting.Float64(7196.54),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingRecordReference{
                            DataType: codataccounting.String("soluta"),
                            ID: codataccounting.String("97ee102d-a2de-435f-8e01-bf33eaab4540"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("dolorum"),
                                ID: codataccounting.String("c1704bf1-cc9f-4c61-aae5-eb5f0c492b57"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "4d08a226-7aae-4e79-a3c7-1ad31becb83d",
                            Name: codataccounting.String("Cindy Koelpin"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e3bfc23d-9450-4a98-aa49-5bac707f06b2",
                            Name: codataccounting.String("Cornelius Schinner"),
                        },
                    },
                    UnitAmount: 4203.54,
                },
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("492386f6-2c96-49c4-8c6b-78890a3fd3c8"),
                        Name: codataccounting.String("Mr. Doreen Nicolas"),
                    },
                    Description: codataccounting.String("deleniti"),
                    DiscountAmount: codataccounting.Float64(7596.13),
                    DiscountPercentage: codataccounting.Float64(1614.81),
                    ItemRef: &shared.ItemRef{
                        ID: "3df931da-3edb-451f-ad94-acc943513772",
                        Name: codataccounting.String("Doreen Boehm"),
                    },
                    Quantity: 1584.51,
                    SubTotal: codataccounting.Float64(1199.27),
                    TaxAmount: codataccounting.Float64(7214.48),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(5545.08),
                        ID: codataccounting.String("32a56d69-180f-4f60-ab9a-6658e69a4b84"),
                        Name: codataccounting.String("Desiree Fisher"),
                    },
                    TotalAmount: codataccounting.Float64(8447.03),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingRecordReference{
                            DataType: codataccounting.String("harum"),
                            ID: codataccounting.String("ec75c68c-6065-4946-8ce3-04d8849bf821"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("eligendi"),
                                ID: codataccounting.String("337f96bb-0c69-4e37-adb1-344ba9f78a5c"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("accusantium"),
                                ID: codataccounting.String("ed7aab62-e972-461f-b0c5-8d27b51996b5"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "4b50eef7-12b7-4a7a-b034-4b1710688dee",
                            Name: codataccounting.String("Santiago Zboncak"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "7f3dd0cc-d33f-411b-be4e-080aa104186e",
                            Name: codataccounting.String("Darren Herman"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "02f3702c-5c8e-42d3-8ead-3104fa44707b",
                            Name: codataccounting.String("Johnny Kunze"),
                        },
                    },
                    UnitAmount: 2839.36,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("sed"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(1449.13),
                        TotalAmount: codataccounting.Float64(909.18),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("fdb2f69e-5926-47c7-9cc8-d3cd4258d035"),
                            Name: codataccounting.String("Cameron Macejkovic"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(161.32),
                        ID: codataccounting.String("8fe2751a-2047-4c04-89e1-43f9619bb7d4"),
                        Note: codataccounting.String("aperiam"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("corporis"),
                        TotalAmount: codataccounting.Float64(6731.58),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(9726.54),
                        TotalAmount: codataccounting.Float64(6701.98),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("436e6259-233f-495c-9d23-7397c785b5db"),
                            Name: codataccounting.String("Mr. Shawna Hayes V"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(9715.31),
                        ID: codataccounting.String("ebdf676b-7206-4dab-b500-52a5647edc43"),
                        Note: codataccounting.String("natus"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("quibusdam"),
                        TotalAmount: codataccounting.Float64(5481.43),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(2387.23),
                        TotalAmount: codataccounting.Float64(1301.47),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("0f41240d-4487-4ac6-93b9-4c3b9d2488d7"),
                            Name: codataccounting.String("Tommy Nienow"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(9901.25),
                        ID: codataccounting.String("c405669f-69a0-406d-a124-9450819d7c3b"),
                        Note: codataccounting.String("et"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("magnam"),
                        TotalAmount: codataccounting.Float64(745.3),
                    },
                },
            },
            Reference: codataccounting.String("atque"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SubTotal: 2793.53,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "nisi": map[string]interface{}{
                        "saepe": "accusantium",
                    },
                },
            },
            TaxAmount: 263.34,
            TotalAmount: 1989.34,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(82723),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDirectCostResponse != nil {
        // handle response
    }
}
```

## DownloadAttachment

Downloads an attachment for the specified direct cost for a given company.

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
    res, err := s.DirectCosts.DownloadAttachment(ctx, operations.DownloadDirectCostAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "ipsa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

## Get

Gets the specified direct cost for a given company.

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
    res, err := s.DirectCosts.Get(ctx, operations.GetDirectCostRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "possimus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectCost != nil {
        // handle response
    }
}
```

## GetAttachment

Gets the specified direct cost attachment for a given company.

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
    res, err := s.DirectCosts.GetAttachment(ctx, operations.GetDirectCostAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "perferendis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Attachment != nil {
        // handle response
    }
}
```

## GetCreateModel

Get create direct cost model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts) for integrations that support creating direct costs.

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
    res, err := s.DirectCosts.GetCreateModel(ctx, operations.GetCreateDirectCostsModelRequest{
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

Gets the direct costs for the company.

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
    res, err := s.DirectCosts.List(ctx, operations.ListDirectCostsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("qui"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectCosts != nil {
        // handle response
    }
}
```

## ListAttachments

Gets all attachments for the specified direct cost for a given company.

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
    res, err := s.DirectCosts.ListAttachments(ctx, operations.ListDirectCostAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "velit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachmentsDataset != nil {
        // handle response
    }
}
```

## UploadAttachment

Posts a new direct cost attachment for a given company.

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
    res, err := s.DirectCosts.UploadAttachment(ctx, operations.UploadDirectCostAttachmentRequest{
        RequestBody: &operations.UploadDirectCostAttachmentRequestBody{
            Content: []byte("repellendus"),
            RequestBody: "quod",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "error",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
