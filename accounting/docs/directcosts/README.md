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
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

func main() {
    s := codataccounting.New(
        codataccounting.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.DirectCosts.Create(ctx, operations.CreateDirectCostRequest{
        DirectCost: &shared.DirectCost{
            ContactRef: &shared.ContactRef{
                DataType: codataccounting.String("quis"),
                ID: "93260525-1e66-4bb4-a689-7d99a2d33567",
            },
            Currency: "ipsa",
            CurrencyRate: codataccounting.Float64(8842.06),
            ID: codataccounting.String("93ee6cf5-9f35-48aa-aaca-e323a31bf7ba"),
            IssueDate: "architecto",
            LineItems: []shared.DirectCostLineItem{
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("c97716c8-02cc-49e0-87d9-d323f1aa63ed"),
                        Name: codataccounting.String("Miss Earnest Zemlak"),
                    },
                    Description: codataccounting.String("nemo"),
                    DiscountAmount: codataccounting.Float64(4125.09),
                    DiscountPercentage: codataccounting.Float64(7425.62),
                    ItemRef: &shared.ItemRef{
                        ID: "cba51ef2-454a-447f-acf1-16cdd5444a75",
                        Name: codataccounting.String("Andrea Lang"),
                    },
                    Quantity: 8116.96,
                    SubTotal: codataccounting.Float64(4938),
                    TaxAmount: codataccounting.Float64(8445.24),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(8127.53),
                        ID: codataccounting.String("9efaf43d-c623-4620-b313-8f30df3db022"),
                        Name: codataccounting.String("Donnie Ondricka"),
                    },
                    TotalAmount: codataccounting.Float64(3413.59),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingRecordReference{
                            DataType: codataccounting.String("delectus"),
                            ID: codataccounting.String("b8f652eb-b9d3-4838-b879-0243b293dab3"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("vero"),
                                ID: codataccounting.String("917f50fd-a04c-48b1-bb55-a292b0bc3bb7"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "4664eb1d-0338-48b0-91bb-17afee74b6fe",
                            Name: codataccounting.String("Marion Greenfelder"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "c7edaf39-d16f-4bf7-afd1-62b303e3023b",
                            Name: codataccounting.String("Manuel Tillman"),
                        },
                    },
                    UnitAmount: 1973.88,
                },
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("16cf55b4-3135-453c-8f1c-204c4adcc990"),
                        Name: codataccounting.String("Ms. Jody Hamill"),
                    },
                    Description: codataccounting.String("cum"),
                    DiscountAmount: codataccounting.Float64(5405.93),
                    DiscountPercentage: codataccounting.Float64(3840.98),
                    ItemRef: &shared.ItemRef{
                        ID: "48cefa78-f1e2-4d3b-901e-0952bbb4cbb1",
                        Name: codataccounting.String("Mrs. Courtney Kuhic"),
                    },
                    Quantity: 5685.48,
                    SubTotal: codataccounting.Float64(3158.92),
                    TaxAmount: codataccounting.Float64(6366.25),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2930.95),
                        ID: codataccounting.String("169c1387-271e-418e-a9e4-5118c2cc57fb"),
                        Name: codataccounting.String("Tyler Abernathy MD"),
                    },
                    TotalAmount: codataccounting.Float64(4645.41),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingRecordReference{
                            DataType: codataccounting.String("molestias"),
                            ID: codataccounting.String("ed29a9d4-eea8-4565-8c2d-4f4c88be4f27"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("tenetur"),
                                ID: codataccounting.String("d9667e46-c51d-42ff-aa58-dcef234c955b"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("natus"),
                                ID: codataccounting.String("bdf2190a-bd9b-4bcc-a725-ec2659ce0280"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("corrupti"),
                                ID: codataccounting.String("40c69ef6-8e45-4c8a-9dfa-c754500430c6"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "32b4391f-df01-4c3e-91e8-f7bc69d460a7",
                            Name: codataccounting.String("Alyssa Satterfield"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "26d10f1e-f263-41c7-80f0-f873f9d5c25f",
                            Name: codataccounting.String("Miss Alfred VonRueden"),
                        },
                    },
                    UnitAmount: 6464.87,
                },
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("4a4253c3-0257-411f-82c7-e7dc548be09e"),
                        Name: codataccounting.String("Diane Paucek"),
                    },
                    Description: codataccounting.String("dolores"),
                    DiscountAmount: codataccounting.Float64(816.73),
                    DiscountPercentage: codataccounting.Float64(3686.17),
                    ItemRef: &shared.ItemRef{
                        ID: "ca12a4ba-9d59-4988-992c-fd0c77c53e7e",
                        Name: codataccounting.String("Chelsea Gibson"),
                    },
                    Quantity: 3952.6,
                    SubTotal: codataccounting.Float64(9251.57),
                    TaxAmount: codataccounting.Float64(5344.54),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7493.37),
                        ID: codataccounting.String("90bac384-e239-4670-bfec-31c50824d189"),
                        Name: codataccounting.String("Curtis Keebler"),
                    },
                    TotalAmount: codataccounting.Float64(6908.65),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingRecordReference{
                            DataType: codataccounting.String("sunt"),
                            ID: codataccounting.String("d27eb707-aa60-4c8f-a46e-6177db9db3b7"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("voluptatibus"),
                                ID: codataccounting.String("fbb6970e-e770-4e36-897e-f7c206e61b0d"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "08714c20-a3d9-4863-bca8-5c3fe65574db",
                            Name: codataccounting.String("Dominick Mraz"),
                        },
                    },
                    UnitAmount: 4622.78,
                },
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("c98f13af-28db-42cf-abf4-f3ded356d7e1"),
                        Name: codataccounting.String("Miss Whitney Dach"),
                    },
                    Description: codataccounting.String("occaecati"),
                    DiscountAmount: codataccounting.Float64(5317.63),
                    DiscountPercentage: codataccounting.Float64(1078.63),
                    ItemRef: &shared.ItemRef{
                        ID: "96d55af6-9a1c-44b7-9ae3-3681c23c39a7",
                        Name: codataccounting.String("Ms. Michael Torphy"),
                    },
                    Quantity: 7162.96,
                    SubTotal: codataccounting.Float64(938.38),
                    TaxAmount: codataccounting.Float64(1709.49),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7710.42),
                        ID: codataccounting.String("5ba825fe-22cd-45cb-a6fb-fec932af6813"),
                        Name: codataccounting.String("Angel Harvey"),
                    },
                    TotalAmount: codataccounting.Float64(9353.01),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingRecordReference{
                            DataType: codataccounting.String("placeat"),
                            ID: codataccounting.String("ec2dd691-6f7f-4c7d-9a70-ec60e6075894"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("autem"),
                                ID: codataccounting.String("1c14cd90-227e-437c-8d97-7f1a5491abe9"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("dignissimos"),
                                ID: codataccounting.String("51b106d2-3e03-4e69-815a-ae99fcde9e72"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("iste"),
                                ID: codataccounting.String("c9d4f2d8-a446-440c-a60d-b73a2f93f467"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("at"),
                                ID: codataccounting.String("c0d8da56-1220-426a-b8f2-77485c1976af"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "80da7a08-9fc4-44db-a745-30e5cc7c6d0c",
                            Name: codataccounting.String("Spencer Wintheiser"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "d334b6f6-23bc-4eca-b50a-ee5e0da8b9af",
                            Name: codataccounting.String("Mrs. Alberta Stoltenberg"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "86e7b413-cbe2-4d17-adc1-c43d40f61d17",
                            Name: codataccounting.String("Christine Hane"),
                        },
                    },
                    UnitAmount: 7322.16,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("debitis"),
            Note: codataccounting.String("enim"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("saepe"),
                        Currency: codataccounting.String("non"),
                        CurrencyRate: codataccounting.Float64(9407.97),
                        TotalAmount: codataccounting.Float64(4601.8),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("21184077-2f32-4e3b-89db-e0f23b7b6d99"),
                            Name: codataccounting.String("Dianne Simonis"),
                        },
                        Currency: codataccounting.String("facere"),
                        CurrencyRate: codataccounting.Float64(8999.7),
                        ID: codataccounting.String("d477680f-c7a1-47a8-ae5e-82fd28d1040a"),
                        Note: codataccounting.String("iusto"),
                        PaidOnDate: codataccounting.String("debitis"),
                        Reference: codataccounting.String("sint"),
                        TotalAmount: codataccounting.Float64(1083.49),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("ratione"),
                        Currency: codataccounting.String("omnis"),
                        CurrencyRate: codataccounting.Float64(1850.41),
                        TotalAmount: codataccounting.Float64(6781.29),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("b44cb183-5008-4f46-9ce5-3e914498a9ba"),
                            Name: codataccounting.String("Alma Beahan"),
                        },
                        Currency: codataccounting.String("pariatur"),
                        CurrencyRate: codataccounting.Float64(9692.94),
                        ID: codataccounting.String("de410c37-daa9-4182-a49d-9625d3caffc1"),
                        Note: codataccounting.String("cupiditate"),
                        PaidOnDate: codataccounting.String("blanditiis"),
                        Reference: codataccounting.String("voluptates"),
                        TotalAmount: codataccounting.Float64(9025.46),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("animi"),
                        Currency: codataccounting.String("modi"),
                        CurrencyRate: codataccounting.Float64(3101.3),
                        TotalAmount: codataccounting.Float64(3390.36),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("2792bcd4-40ea-498b-acce-0486de0d56d7"),
                            Name: codataccounting.String("Mrs. Verna Anderson"),
                        },
                        Currency: codataccounting.String("accusantium"),
                        CurrencyRate: codataccounting.Float64(2062.3),
                        ID: codataccounting.String("e8dc626f-f77c-4656-b5f5-b70e3e4cfcc6"),
                        Note: codataccounting.String("dolorum"),
                        PaidOnDate: codataccounting.String("cupiditate"),
                        Reference: codataccounting.String("ab"),
                        TotalAmount: codataccounting.Float64(8969.21),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("maxime"),
                        Currency: codataccounting.String("veniam"),
                        CurrencyRate: codataccounting.Float64(1816.73),
                        TotalAmount: codataccounting.Float64(3974.18),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("24d00014-ef45-4cea-91ac-53ebb6587f34"),
                            Name: codataccounting.String("Megan Bergnaum"),
                        },
                        Currency: codataccounting.String("ipsam"),
                        CurrencyRate: codataccounting.Float64(7395.69),
                        ID: codataccounting.String("9acee400-ae9f-492c-af1b-025f1d14718c"),
                        Note: codataccounting.String("nisi"),
                        PaidOnDate: codataccounting.String("voluptatibus"),
                        Reference: codataccounting.String("est"),
                        TotalAmount: codataccounting.Float64(1274.87),
                    },
                },
            },
            Reference: codataccounting.String("doloribus"),
            SourceModifiedDate: codataccounting.String("mollitia"),
            SubTotal: 8418.47,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "cumque": map[string]interface{}{
                        "commodi": "porro",
                    },
                },
            },
            TaxAmount: 3298.49,
            TotalAmount: 8155.84,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(571303),
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.DirectCosts.DownloadAttachment(ctx, operations.DownloadDirectCostAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.DirectCosts.Get(ctx, operations.GetDirectCostRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.DirectCosts.GetAttachment(ctx, operations.GetDirectCostAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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
            AuthHeader: "YOUR_API_KEY_HERE",
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.DirectCosts.List(ctx, operations.ListDirectCostsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("exercitationem"),
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.DirectCosts.ListAttachments(ctx, operations.ListDirectCostAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.DirectCosts.UploadAttachment(ctx, operations.UploadDirectCostAttachmentRequest{
        RequestBody: &operations.UploadDirectCostAttachmentRequestBody{
            Content: []byte("quaerat"),
            RequestBody: "in",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
