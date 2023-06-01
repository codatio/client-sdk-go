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
                DataType: codataccounting.String("consequuntur"),
                ID: "bd7ed565-0762-41c5-8f4d-7396564c20a0",
            },
            Currency: "reprehenderit",
            CurrencyRate: codataccounting.Float64(1171.17),
            ID: codataccounting.String("1a961d24-a7db-4b8f-932d-892cf7812cb5"),
            IssueDate: "architecto",
            LineItems: []shared.DirectCostLineItem{
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("c878240b-f548-4f88-b8f1-bf0bc8e1f206"),
                        Name: codataccounting.String("Corey Streich"),
                    },
                    Description: codataccounting.String("illo"),
                    DiscountAmount: codataccounting.Float64(8384.02),
                    DiscountPercentage: codataccounting.Float64(403.46),
                    ItemRef: &shared.ItemRef{
                        ID: "081090f6-7066-473f-ba68-1c5768dce742",
                        Name: codataccounting.String("Donna Mann"),
                    },
                    Quantity: 1058.74,
                    SubTotal: codataccounting.Float64(3497.12),
                    TaxAmount: codataccounting.Float64(8881.17),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(532.16),
                        ID: codataccounting.String("8601489a-5f63-4e3a-b3dd-9dda33dcd634"),
                        Name: codataccounting.String("Nathan Ward"),
                    },
                    TotalAmount: codataccounting.Float64(4568.09),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingRecordReference{
                            DataType: codataccounting.String("id"),
                            ID: codataccounting.String("98e4df37-e45b-4895-9d41-3e13a4823109"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("iure"),
                                ID: codataccounting.String("bd354c09-2bd7-434f-8244-9d86f4bb20fe"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "d911cbfe-749c-4af4-9a27-f69e2c9e6d10",
                            Name: codataccounting.String("Arturo Smith"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "ad4c6b03-108d-49c3-b747-3082b94f2ab1",
                            Name: codataccounting.String("Rufus Hickle"),
                        },
                    },
                    UnitAmount: 837,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("voluptates"),
            Note: codataccounting.String("unde"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("amet"),
                        Currency: codataccounting.String("dolores"),
                        CurrencyRate: codataccounting.Float64(4167.82),
                        TotalAmount: codataccounting.Float64(2051.5),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("50a46714-3789-4ce0-a991-594d93a74c02"),
                            Name: codataccounting.String("Jane Wilkinson"),
                        },
                        Currency: codataccounting.String("quidem"),
                        CurrencyRate: codataccounting.Float64(3102.12),
                        ID: codataccounting.String("b4db8b77-8ebb-46e1-92cf-502bafb2cbc4"),
                        Note: codataccounting.String("aliquid"),
                        PaidOnDate: codataccounting.String("adipisci"),
                        Reference: codataccounting.String("ipsam"),
                        TotalAmount: codataccounting.Float64(8526.23),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("enim"),
                        Currency: codataccounting.String("eveniet"),
                        CurrencyRate: codataccounting.Float64(4330.83),
                        TotalAmount: codataccounting.Float64(3470.5),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("da028c3e-951a-41e3-8fda-966489d7b786"),
                            Name: codataccounting.String("Mrs. Josephine Tromp"),
                        },
                        Currency: codataccounting.String("quasi"),
                        CurrencyRate: codataccounting.Float64(1772.5),
                        ID: codataccounting.String("a6b99249-4594-4487-b5c8-43836b86b3cd"),
                        Note: codataccounting.String("a"),
                        PaidOnDate: codataccounting.String("ex"),
                        Reference: codataccounting.String("dolore"),
                        TotalAmount: codataccounting.Float64(1158.7),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("minima"),
                        Currency: codataccounting.String("facilis"),
                        CurrencyRate: codataccounting.Float64(223.76),
                        TotalAmount: codataccounting.Float64(2799.72),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("49f9df13-f4ee-4dbe-b8bf-606825894ea7"),
                            Name: codataccounting.String("Grace Stehr"),
                        },
                        Currency: codataccounting.String("in"),
                        CurrencyRate: codataccounting.Float64(1479.33),
                        ID: codataccounting.String("795b7851-48d6-4d54-9e56-35b33bc0f970"),
                        Note: codataccounting.String("placeat"),
                        PaidOnDate: codataccounting.String("dolore"),
                        Reference: codataccounting.String("magni"),
                        TotalAmount: codataccounting.Float64(9730.03),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("quod"),
                        Currency: codataccounting.String("provident"),
                        CurrencyRate: codataccounting.Float64(9624.68),
                        TotalAmount: codataccounting.Float64(2930.13),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("844225e7-5b79-4606-9c0e-fa6f93b90a1b"),
                            Name: codataccounting.String("Colin Mills"),
                        },
                        Currency: codataccounting.String("accusamus"),
                        CurrencyRate: codataccounting.Float64(1147.52),
                        ID: codataccounting.String("254b739f-4fe7-4721-8d1f-6558c99c722d"),
                        Note: codataccounting.String("fugit"),
                        PaidOnDate: codataccounting.String("nam"),
                        Reference: codataccounting.String("optio"),
                        TotalAmount: codataccounting.Float64(349.2),
                    },
                },
            },
            Reference: codataccounting.String("earum"),
            SourceModifiedDate: codataccounting.String("excepturi"),
            SubTotal: 2568.9,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "voluptatum": map[string]interface{}{
                        "possimus": "unde",
                        "maxime": "culpa",
                    },
                },
            },
            TaxAmount: 6428.58,
            TotalAmount: 9268.79,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(42929),
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
        DirectCostID: "magnam",
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
        DirectCostID: "quia",
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
        DirectCostID: "quibusdam",
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
        Query: codataccounting.String("temporibus"),
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
        DirectCostID: "voluptate",
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
            Content: []byte("placeat"),
            RequestBody: "est",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "est",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
