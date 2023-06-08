# DirectIncomes

## Overview

Direct incomes

### Available Operations

* [Create](#create) - Create direct income
* [DownloadAttachment](#downloadattachment) - Download direct income attachment
* [Get](#get) - Get direct income
* [GetAttachment](#getattachment) - Get direct income attachment
* [GetCreateModel](#getcreatemodel) - Get create direct income model
* [List](#list) - List direct incomes
* [ListAttachments](#listattachments) - List direct income attachments
* [UploadAttachment](#uploadattachment) - Create direct income attachment

## Create

Posts a new direct income to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create direct income model](https://docs.codat.io/accounting-api#/operations/get-create-directIncomes-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) to see which integrations support this endpoint.

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
    res, err := s.DirectIncomes.Create(ctx, operations.CreateDirectIncomeRequest{
        DirectIncome: &shared.DirectIncome{
            ContactRef: &shared.ContactRef{
                DataType: codataccounting.String("ipsa"),
                ID: "1f5afd2a-6c44-4846-ae9d-89253c8962f4",
            },
            Currency: "USD",
            CurrencyRate: codataccounting.Float64(5683.23),
            ID: codataccounting.String("6bf51e46-52d3-4c34-bd61-778af4912477"),
            IssueDate: "2022-10-23T00:00:00.000Z",
            LineItems: []shared.DirectIncomeLineItem{
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("e621909e-9104-44a5-9e59-ac7706670cf1"),
                        Name: codataccounting.String("Randal Hayes"),
                    },
                    Description: codataccounting.String("qui"),
                    DiscountAmount: codataccounting.Float64(4003.27),
                    DiscountPercentage: codataccounting.Float64(437.63),
                    ItemRef: &shared.ItemRef{
                        ID: "5251e66b-b426-4897-999a-2d335670e93e",
                        Name: codataccounting.String("Gilbert Rutherford"),
                    },
                    Quantity: 5978.75,
                    SubTotal: codataccounting.Float64(9668.01),
                    TaxAmount: codataccounting.Float64(2112.46),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(3687.85),
                        ID: codataccounting.String("8aaeacae-323a-431b-b7ba-1cc97716c802"),
                        Name: codataccounting.String("Wilson Marvin DDS"),
                    },
                    TotalAmount: codataccounting.Float64(4453.35),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "9d323f1a-a63e-4d9c-b1c8-56bcba51ef24",
                            Name: codataccounting.String("Marjorie Nicolas"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "facf116c-dd54-444a-b562-873c7dd9efaf",
                            Name: codataccounting.String("Sylvia Swift"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "23620f31-38f3-40df-bdb0-22faa565fb8f",
                            Name: codataccounting.String("Sally Crooks"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "b9d38383-8790-4243-b293-dab30e917f50",
                            Name: codataccounting.String("Mrs. Taylor O'Kon"),
                        },
                    },
                    UnitAmount: 5328.66,
                },
                shared.DirectIncomeLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("b1bb55a2-92b0-4bc3-bb74-4664eb1d0338"),
                        Name: codataccounting.String("Cesar Adams MD"),
                    },
                    Description: codataccounting.String("rerum"),
                    DiscountAmount: codataccounting.Float64(991.83),
                    DiscountPercentage: codataccounting.Float64(4480.63),
                    ItemRef: &shared.ItemRef{
                        ID: "afee74b6-feb9-4457-87ed-af39d16fbf76",
                        Name: codataccounting.String("Sammy Bernhard"),
                    },
                    Quantity: 6954.63,
                    SubTotal: codataccounting.Float64(1945.68),
                    TaxAmount: codataccounting.Float64(418.71),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(2210),
                        ID: codataccounting.String("e3023b93-e343-416c-b55b-4313553ccf1c"),
                        Name: codataccounting.String("Barbara Gottlieb"),
                    },
                    TotalAmount: codataccounting.Float64(6364.29),
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "cc9904c5-195b-4864-8cef-a78f1e2d3b90",
                            Name: codataccounting.String("Gwen Barrows"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "2bbb4cbb-19f7-413d-95a4-169c1387271e",
                            Name: codataccounting.String("Billie Vandervort"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "e45118c2-cc57-4fbd-a0b1-a78ed29a9d4e",
                            Name: codataccounting.String("Matt Larson"),
                        },
                        shared.TrackingCategoryRef{
                            ID: "58c2d4f4-c88b-4e4f-a78f-d9667e46c51d",
                            Name: codataccounting.String("Camille Zulauf"),
                        },
                    },
                    UnitAmount: 3237.44,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Note: codataccounting.String("temporibus"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(1490.64),
                        TotalAmount: codataccounting.Float64(1908.74),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("4c955b9b-df21-490a-bd9b-bcc2725ec265"),
                            Name: codataccounting.String("Mr. Sylvester Trantow"),
                        },
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(5459.03),
                        ID: codataccounting.String("40c69ef6-8e45-4c8a-9dfa-c754500430c6"),
                        Note: codataccounting.String("commodi"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("quia"),
                        TotalAmount: codataccounting.Float64(7036.48),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(5888.01),
                        TotalAmount: codataccounting.Float64(814.31),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("fdf01c3e-91e8-4f7b-869d-460a77eceb26"),
                            Name: codataccounting.String("Joe Barton DVM"),
                        },
                        Currency: codataccounting.String("EUR"),
                        CurrencyRate: codataccounting.Float64(1818.77),
                        ID: codataccounting.String("631c7c0f-0f87-43f9-95c2-5fd3e0b4a4a4"),
                        Note: codataccounting.String("quia"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("dolor"),
                        TotalAmount: codataccounting.Float64(8051.65),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("GBP"),
                        CurrencyRate: codataccounting.Float64(1626.75),
                        TotalAmount: codataccounting.Float64(3545.33),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("711f42c7-e7dc-4548-be09-e41a7a215ca1"),
                            Name: codataccounting.String("Hattie Green"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(8445.45),
                        ID: codataccounting.String("59988192-cfd0-4c77-853e-7e7d4ee6e8b9"),
                        Note: codataccounting.String("consequatur"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("similique"),
                        TotalAmount: codataccounting.Float64(7821.55),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(2671.84),
                        TotalAmount: codataccounting.Float64(9027.31),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("2396703f-ec31-4c50-824d-189a36a6b2d2"),
                            Name: codataccounting.String("Elvira Price IV"),
                        },
                        Currency: codataccounting.String("USD"),
                        CurrencyRate: codataccounting.Float64(6663.31),
                        ID: codataccounting.String("60c8fe46-e617-47db-9db3-b70ffbb6970e"),
                        Note: codataccounting.String("repudiandae"),
                        PaidOnDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                        Reference: codataccounting.String("nihil"),
                        TotalAmount: codataccounting.Float64(580.86),
                    },
                },
            },
            Reference: codataccounting.String("eveniet"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            SubTotal: 4184.07,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "occaecati": map[string]interface{}{
                        "debitis": "voluptatibus",
                        "esse": "maxime",
                    },
                },
            },
            TaxAmount: 1795.31,
            TotalAmount: 412.89,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(425065),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDirectIncomeResponse != nil {
        // handle response
    }
}
```

## DownloadAttachment

Downloads an attachment for the specified direct income for a given company.

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
    res, err := s.DirectIncomes.DownloadAttachment(ctx, operations.DownloadDirectIncomeAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "itaque",
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

Gets the specified direct income for a given company and connection.

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
    res, err := s.DirectIncomes.Get(ctx, operations.GetDirectIncomeRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "suscipit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectIncome != nil {
        // handle response
    }
}
```

## GetAttachment

Gets the specified direct income attachment for a given company.

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
    res, err := s.DirectIncomes.GetAttachment(ctx, operations.GetDirectIncomeAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "vitae",
        TimeoutInMinutes: codataccounting.Int(730863),
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

Get create direct income model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=directIncomes) for integrations that support creating direct incomes.

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
    res, err := s.DirectIncomes.GetCreateModel(ctx, operations.GetCreateDirectIncomesModelRequest{
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

Lists the direct incomes for a given company.

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
    res, err := s.DirectIncomes.List(ctx, operations.ListDirectIncomesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("sit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectIncomes != nil {
        // handle response
    }
}
```

## ListAttachments

Gets all attachments for the specified direct income for a given company.

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
    res, err := s.DirectIncomes.ListAttachments(ctx, operations.ListDirectIncomeAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "quibusdam",
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

Posts a new direct income attachment for a given company.

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
    res, err := s.DirectIncomes.UploadAttachment(ctx, operations.UploadDirectIncomeAttachmentRequest{
        RequestBody: &operations.UploadDirectIncomeAttachmentRequestBody{
            Content: []byte("dolorem"),
            RequestBody: "aut",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectIncomeID: "blanditiis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
