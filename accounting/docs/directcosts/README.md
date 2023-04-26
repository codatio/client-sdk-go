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
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts) for integrations that support creating direct costs.

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
    req := operations.CreateDirectCostRequest{
        DirectCost: &shared.DirectCost{
            ContactRef: &shared.ContactRef{
                DataType: codataccounting.String("tempore"),
                ID: "e3444eac-8b3a-4287-9c6c-1fe606d07d2a",
            },
            Currency: "error",
            CurrencyRate: codataccounting.Float64(7699.22),
            ID: codataccounting.String("87ae50c1-6661-4a1d-9136-a7e8d53213f3"),
            IssueDate: "asperiores",
            LineItems: []shared.DirectCostLineItem{
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("58752db7-64c5-49f0-a56c-ebcada29ca79"),
                        Name: codataccounting.String("Margie Bosco"),
                    },
                    Description: codataccounting.String("ipsam"),
                    DiscountAmount: codataccounting.Float64(3887.15),
                    DiscountPercentage: codataccounting.Float64(4752.38),
                    ItemRef: &shared.ItemRef{
                        ID: "1663c530-b566-4516-ba36-38512ab2521b",
                        Name: codataccounting.String("Emmett Daugherty IV"),
                    },
                    Quantity: 1632.92,
                    SubTotal: codataccounting.Float64(3028.14),
                    TaxAmount: codataccounting.Float64(4199.95),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4450.23),
                        ID: codataccounting.String("b8a40bc0-5fab-40d6-90ed-f22a94d20ec9"),
                        Name: codataccounting.String("Raquel O'Keefe PhD"),
                    },
                    TotalAmount: codataccounting.Float64(1004.36),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingInvoiceTo{
                            DataType: codataccounting.String("hic"),
                            ID: codataccounting.String("465e8515-6fff-473f-9f54-fdd5ea954339"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("illum"),
                                ID: codataccounting.String("afb42a8d-6338-48e4-9803-9ea5f9b18a24"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("modi"),
                                ID: codataccounting.String("fd619039-dacd-438e-90dc-671dc7f1e3af"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("quasi"),
                                ID: codataccounting.String("5920c90d-1b49-401f-abd8-9c8a32639da5"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "7b6902b8-81a9-44f6-8366-4a8f0af8c691",
                            Name: codataccounting.String("Clinton Ernser"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "9fbaf947-6a2a-4e8d-8c50-c8a3512c7378",
                            Name: codataccounting.String("Gwendolyn McLaughlin IV"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "50a00e96-6ec7-436d-8319-4398c783c923",
                            Name: codataccounting.String("Max Tillman"),
                        },
                    },
                    UnitAmount: 8639.57,
                },
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("3ab7ca3c-5ca8-4649-a70c-fd5d6989b720"),
                        Name: codataccounting.String("Mr. Danielle Hamill"),
                    },
                    Description: codataccounting.String("voluptate"),
                    DiscountAmount: codataccounting.Float64(8494.86),
                    DiscountPercentage: codataccounting.Float64(879.6),
                    ItemRef: &shared.ItemRef{
                        ID: "9ea83d49-2ed1-44b8-a2c1-954545e955dc",
                        Name: codataccounting.String("Henry Langosh"),
                    },
                    Quantity: 6691.93,
                    SubTotal: codataccounting.Float64(3007.59),
                    TaxAmount: codataccounting.Float64(6029.32),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(65.39),
                        ID: codataccounting.String("1c7c43ad-2daa-4784-aba3-d230edf73811"),
                        Name: codataccounting.String("Albert Bruen"),
                    },
                    TotalAmount: codataccounting.Float64(5162.31),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingInvoiceTo{
                            DataType: codataccounting.String("consequuntur"),
                            ID: codataccounting.String("bd7ed565-0762-41c5-8f4d-7396564c20a0"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("dicta"),
                                ID: codataccounting.String("1a961d24-a7db-4b8f-932d-892cf7812cb5"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("architecto"),
                                ID: codataccounting.String("2c878240-bf54-48f8-8f8f-1bf0bc8e1f20"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "d5d831d0-0810-490f-a706-673f3a681c57",
                            Name: codataccounting.String("Vickie Simonis"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "742409a2-15e0-4860-9489-a5f63e3af3dd",
                            Name: codataccounting.String("Marty Spencer"),
                        },
                    },
                    UnitAmount: 2439.65,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("vero"),
            Note: codataccounting.String("placeat"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("vel"),
                        Currency: codataccounting.String("non"),
                        CurrencyRate: codataccounting.Float64(2799.65),
                        TotalAmount: codataccounting.Float64(5083.73),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("3e4a7a98-e4df-437e-85b8-955d413e13a4"),
                            Name: codataccounting.String("Mr. Todd Feil"),
                        },
                        Currency: codataccounting.String("perferendis"),
                        CurrencyRate: codataccounting.Float64(4391.35),
                        ID: codataccounting.String("bd354c09-2bd7-434f-8244-9d86f4bb20fe"),
                        Note: codataccounting.String("nostrum"),
                        PaidOnDate: codataccounting.String("quibusdam"),
                        Reference: codataccounting.String("provident"),
                        TotalAmount: codataccounting.Float64(857.97),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("sunt"),
                        Currency: codataccounting.String("quod"),
                        CurrencyRate: codataccounting.Float64(7101.48),
                        TotalAmount: codataccounting.Float64(9611.71),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("e749caf4-5a27-4f69-a2c9-e6d10e9db3ad"),
                            Name: codataccounting.String("Rosalie Kautzer I"),
                        },
                        Currency: codataccounting.String("quasi"),
                        CurrencyRate: codataccounting.Float64(374.55),
                        ID: codataccounting.String("8d9c3374-7308-42b9-8f2a-b1fd5671e9c3"),
                        Note: codataccounting.String("dolores"),
                        PaidOnDate: codataccounting.String("commodi"),
                        Reference: codataccounting.String("neque"),
                        TotalAmount: codataccounting.Float64(3182.94),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("eaque"),
                        Currency: codataccounting.String("officia"),
                        CurrencyRate: codataccounting.Float64(2702.53),
                        TotalAmount: codataccounting.Float64(4310.35),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("7143789c-e0e9-4915-94d9-3a74c0252fe3"),
                            Name: codataccounting.String("Alex Rippin"),
                        },
                        Currency: codataccounting.String("rerum"),
                        CurrencyRate: codataccounting.Float64(5394.26),
                        ID: codataccounting.String("b778ebb6-e1d2-4cf5-82ba-fb2cbc4635d5"),
                        Note: codataccounting.String("eveniet"),
                        PaidOnDate: codataccounting.String("eum"),
                        Reference: codataccounting.String("exercitationem"),
                        TotalAmount: codataccounting.Float64(8718.88),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("culpa"),
                        Currency: codataccounting.String("alias"),
                        CurrencyRate: codataccounting.Float64(1759.37),
                        TotalAmount: codataccounting.Float64(5500.66),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("c3e951a1-e30f-4da9-a648-9d7b78673e13"),
                            Name: codataccounting.String("Arthur Dare"),
                        },
                        Currency: codataccounting.String("rerum"),
                        CurrencyRate: codataccounting.Float64(5872.48),
                        ID: codataccounting.String("92494594-487f-45c8-8383-6b86b3cdf641"),
                        Note: codataccounting.String("minima"),
                        PaidOnDate: codataccounting.String("facilis"),
                        Reference: codataccounting.String("sit"),
                        TotalAmount: codataccounting.Float64(2799.72),
                    },
                },
            },
            Reference: codataccounting.String("magnam"),
            SourceModifiedDate: codataccounting.String("molestias"),
            SubTotal: 9417.1,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "repellendus": map[string]interface{}{
                        "dicta": "ratione",
                        "delectus": "ut",
                        "officiis": "itaque",
                        "nulla": "distinctio",
                    },
                    "recusandae": map[string]interface{}{
                        "deleniti": "tempore",
                        "reiciendis": "commodi",
                    },
                    "sit": map[string]interface{}{
                        "molestias": "quia",
                        "ipsam": "rem",
                    },
                },
            },
            TaxAmount: 5640.67,
            TotalAmount: 2626.64,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(898366),
    }

    res, err := s.DirectCosts.Create(ctx, req)
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
    req := operations.DownloadDirectCostAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.DirectCosts.DownloadAttachment(ctx, req)
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
    req := operations.GetDirectCostRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.DirectCosts.Get(ctx, req)
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
    req := operations.GetDirectCostAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.DirectCosts.GetAttachment(ctx, req)
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
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directCosts) for integrations that support creating direct costs.

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
    req := operations.GetCreateDirectCostsModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.DirectCosts.GetCreateModel(ctx, req)
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
    req := operations.ListDirectCostsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("culpa"),
    }

    res, err := s.DirectCosts.List(ctx, req)
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
    req := operations.ListDirectCostAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.DirectCosts.ListAttachments(ctx, req)
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
    req := operations.UploadDirectCostAttachmentRequest{
        RequestBody: &operations.UploadDirectCostAttachmentRequestBody{
            Content: []byte("in"),
            RequestBody: "aliquid",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.DirectCosts.UploadAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
