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
                DataType: codataccounting.String("doloremque"),
                ID: "6d212494-5081-49d7-83b1-b41844060e00",
            },
            Currency: "nesciunt",
            CurrencyRate: codataccounting.Float64(827.23),
            ID: codataccounting.String("0d023dc9-01f5-4afd-aa6c-44846ae9d892"),
            IssueDate: "veniam",
            LineItems: []shared.DirectCostLineItem{
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("c8962f48-96bf-451e-8652-d3c343d61778"),
                        Name: codataccounting.String("Amos Hahn Sr."),
                    },
                    Description: codataccounting.String("numquam"),
                    DiscountAmount: codataccounting.Float64(4553.55),
                    DiscountPercentage: codataccounting.Float64(4837.72),
                    ItemRef: &shared.ItemRef{
                        ID: "25e62190-9e91-4044-a5de-59ac7706670c",
                        Name: codataccounting.String("Henry Ruecker"),
                    },
                    Quantity: 6035.57,
                    SubTotal: codataccounting.Float64(2347.4),
                    TaxAmount: codataccounting.Float64(1862.22),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(4003.27),
                        ID: codataccounting.String("05251e66-bb42-4689-bd99-a2d335670e93"),
                        Name: codataccounting.String("Clay Hyatt"),
                    },
                    TotalAmount: codataccounting.Float64(3195.13),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingRecordReference{
                            DataType: codataccounting.String("perspiciatis"),
                            ID: codataccounting.String("f358aaea-cae3-423a-b1bf-7ba1cc97716c"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("accusantium"),
                                ID: codataccounting.String("2cc9e0c7-d9d3-423f-9aa6-3ed9cf1c856b"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("quo"),
                                ID: codataccounting.String("ba51ef24-54a4-47fa-8f11-6cdd5444a756"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("consequuntur"),
                                ID: codataccounting.String("873c7dd9-efaf-443d-8623-620f3138f30d"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "3db022fa-a565-4fb8-b652-ebb9d3838387",
                            Name: codataccounting.String("Jason Considine"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b293dab3-0e91-47f5-8fda-04c8b1bb55a2",
                            Name: codataccounting.String("Miss Russell Ritchie"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "3bb74466-4eb1-4d03-b88b-0d1bb17afee7",
                            Name: codataccounting.String("Latoya Hodkiewicz"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b9457c7e-daf3-49d1-afbf-76fd162b303e",
                            Name: codataccounting.String("Sarah Cremin"),
                        },
                    },
                    UnitAmount: 5677.5,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("ipsum"),
            Note: codataccounting.String("accusamus"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("tempora"),
                        Currency: codataccounting.String("sequi"),
                        CurrencyRate: codataccounting.Float64(893.2),
                        TotalAmount: codataccounting.Float64(3986.87),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("cf55b431-3553-4ccf-9c20-4c4adcc9904c"),
                            Name: codataccounting.String("Debra Medhurst"),
                        },
                        Currency: codataccounting.String("atque"),
                        CurrencyRate: codataccounting.Float64(3840.98),
                        ID: codataccounting.String("48cefa78-f1e2-4d3b-901e-0952bbb4cbb1"),
                        Note: codataccounting.String("iste"),
                        PaidOnDate: codataccounting.String("voluptatibus"),
                        Reference: codataccounting.String("odio"),
                        TotalAmount: codataccounting.Float64(665.27),
                    },
                },
            },
            Reference: codataccounting.String("neque"),
            SourceModifiedDate: codataccounting.String("pariatur"),
            SubTotal: 5685.48,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "culpa": map[string]interface{}{
                        "sunt": "nisi",
                        "molestias": "impedit",
                    },
                    "quasi": map[string]interface{}{
                        "corrupti": "in",
                    },
                },
            },
            TaxAmount: 1589.19,
            TotalAmount: 4888.02,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(119013),
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
        DirectCostID: "eveniet",
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
        DirectCostID: "vitae",
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
        DirectCostID: "quos",
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
        Query: codataccounting.String("eveniet"),
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
        DirectCostID: "officia",
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
            Content: []byte("perspiciatis"),
            RequestBody: "debitis",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "non",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
