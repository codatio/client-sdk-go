# DirectCosts

## Overview

Direct costs

### Available Operations

* [CreateDirectCost](#createdirectcost) - Create direct cost
* [DownloadDirectCostAttachment](#downloaddirectcostattachment) - Download direct cost attachment
* [GetCreateDirectCostsModel](#getcreatedirectcostsmodel) - Get create direct cost model
* [GetDirectCost](#getdirectcost) - Get direct cost
* [GetDirectCostAttachment](#getdirectcostattachment) - Get direct cost attachment
* [GetDirectCosts](#getdirectcosts) - List direct costs
* [ListDirectCostAttachments](#listdirectcostattachments) - List direct cost attachments
* [UploadDirectCostAttachment](#uploaddirectcostattachment) - Upload direct cost attachment

## CreateDirectCost

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
                DataType: codataccounting.String("sed"),
                ID: "de35f8e0-1bf3-43ea-ab45-402ac1704bf1",
            },
            Currency: "maxime",
            CurrencyRate: codataccounting.Float64(8111.67),
            ID: codataccounting.String("9fc61aae-5eb5-4f0c-892b-5744d08a2267"),
            IssueDate: "est",
            LineItems: []shared.DirectCostLineItem{
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("ee79e3c7-1ad3-41be-8b83-d2378ae3bfc2"),
                        Name: codataccounting.String("Christie Marquardt"),
                    },
                    Description: codataccounting.String("aperiam"),
                    DiscountAmount: codataccounting.Float64(6589.17),
                    DiscountPercentage: codataccounting.Float64(6038.63),
                    ItemRef: &shared.ItemRef{
                        ID: "86a495ba-c707-4f06-b28e-cc86492386f6",
                        Name: codataccounting.String("Lynne Miller"),
                    },
                    Quantity: 7548.72,
                    SubTotal: codataccounting.Float64(2765.07),
                    TaxAmount: codataccounting.Float64(7900.8),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(7977.92),
                        ID: codataccounting.String("6b78890a-3fd3-4c81-9a10-f8c23df931da"),
                        Name: codataccounting.String("Henrietta Schuppe"),
                    },
                    TotalAmount: codataccounting.Float64(1152.02),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingInvoiceTo{
                            DataType: codataccounting.String("repellat"),
                            ID: codataccounting.String("ad94acc9-4351-4377-a6d1-5321b832a56d"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("sint"),
                                ID: codataccounting.String("180ff60e-b9a6-4658-a69a-4b843d382dbe"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("porro"),
                                ID: codataccounting.String("75c68c60-6594-468c-a304-d8849bf8214c"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "37f96bb0-c69e-4372-9b13-44ba9f78a5c0",
                            Name: codataccounting.String("Bryant Kiehn"),
                        },
                    },
                    UnitAmount: 7429.37,
                },
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("62e97261-fb0c-458d-a7b5-1996b5b4b50e"),
                        Name: codataccounting.String("Mr. Darrin Kub"),
                    },
                    Description: codataccounting.String("ducimus"),
                    DiscountAmount: codataccounting.Float64(6577.63),
                    DiscountPercentage: codataccounting.Float64(4967.86),
                    ItemRef: &shared.ItemRef{
                        ID: "ab0344b1-7106-488d-aebe-f897f3dd0ccd",
                        Name: codataccounting.String("Mr. Cindy White"),
                    },
                    Quantity: 2054.73,
                    SubTotal: codataccounting.Float64(9131.52),
                    TaxAmount: codataccounting.Float64(3033.65),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(9165.77),
                        ID: codataccounting.String("080aa104-186e-4c75-9e02-f3702c5c8e2d"),
                        Name: codataccounting.String("Barbara Von"),
                    },
                    TotalAmount: codataccounting.Float64(2133.99),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingInvoiceTo{
                            DataType: codataccounting.String("sunt"),
                            ID: codataccounting.String("04fa4470-7bf3-475b-8428-2821fdb2f69e"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("unde"),
                                ID: codataccounting.String("267c71cc-8d3c-4d42-98d0-358a82c808fe"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("odit"),
                                ID: codataccounting.String("751a2047-c044-49e1-83f9-619bb7d40d5a"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "1fa436e6-2592-433f-95c9-d237397c785b",
                            Name: codataccounting.String("Betsy Reynolds"),
                        },
                    },
                    UnitAmount: 3401.67,
                },
                shared.DirectCostLineItem{
                    AccountRef: &shared.AccountRef{
                        ID: codataccounting.String("00183feb-df67-46b7-a06d-ab750052a564"),
                        Name: codataccounting.String("Gretchen Steuber"),
                    },
                    Description: codataccounting.String("sequi"),
                    DiscountAmount: codataccounting.Float64(6159.32),
                    DiscountPercentage: codataccounting.Float64(9053.57),
                    ItemRef: &shared.ItemRef{
                        ID: "d8c4320f-4124-40d4-887a-c693b94c3b9d",
                        Name: codataccounting.String("Esther Lind"),
                    },
                    Quantity: 4884.37,
                    SubTotal: codataccounting.Float64(6008.1),
                    TaxAmount: codataccounting.Float64(3215.82),
                    TaxRateRef: &shared.TaxRateRef{
                        EffectiveTaxRate: codataccounting.Float64(6383.23),
                        ID: codataccounting.String("a42fc405-669f-469a-806d-21249450819d"),
                        Name: codataccounting.String("Krista Emard MD"),
                    },
                    TotalAmount: codataccounting.Float64(3003.75),
                    Tracking: &shared.Tracking{
                        InvoiceTo: &shared.TrackingInvoiceTo{
                            DataType: codataccounting.String("illo"),
                            ID: codataccounting.String("844060e0-0310-4d02-bdc9-01f5afd2a6c4"),
                        },
                        RecordRefs: []shared.InvoiceTo{
                            shared.InvoiceTo{
                                DataType: codataccounting.String("praesentium"),
                                ID: codataccounting.String("46ae9d89-253c-4896-af48-96bf51e4652d"),
                            },
                            shared.InvoiceTo{
                                DataType: codataccounting.String("dolorem"),
                                ID: codataccounting.String("c343d617-78af-4491-a477-25e621909e91"),
                            },
                        },
                    },
                    TrackingCategoryRefs: []shared.TrackingCategoryRef{
                        shared.TrackingCategoryRef{
                            ID: "44a5de59-ac77-4066-b0cf-1cf593260525",
                            Name: codataccounting.String("Elvira Jacobson"),
                        },
                    },
                    UnitAmount: 7414,
                },
            },
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("non"),
            Note: codataccounting.String("quia"),
            PaymentAllocations: []shared.Items{
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("deleniti"),
                        Currency: codataccounting.String("molestias"),
                        CurrencyRate: codataccounting.Float64(4937.34),
                        TotalAmount: codataccounting.Float64(8134.63),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("99a2d335-670e-493e-a6cf-59f358aaeaca"),
                            Name: codataccounting.String("Philip Crooks"),
                        },
                        Currency: codataccounting.String("adipisci"),
                        CurrencyRate: codataccounting.Float64(799.07),
                        ID: codataccounting.String("bf7ba1cc-9771-46c8-82cc-9e0c7d9d323f"),
                        Note: codataccounting.String("quae"),
                        PaidOnDate: codataccounting.String("animi"),
                        Reference: codataccounting.String("est"),
                        TotalAmount: codataccounting.Float64(4209.27),
                    },
                },
                shared.Items{
                    Allocation: shared.ItemsAllocation{
                        AllocatedOnDate: codataccounting.String("sequi"),
                        Currency: codataccounting.String("officiis"),
                        CurrencyRate: codataccounting.Float64(8610.9),
                        TotalAmount: codataccounting.Float64(5823.51),
                    },
                    Payment: shared.PaymentAllocationPayment{
                        AccountRef: &shared.AccountRef{
                            ID: codataccounting.String("cf1c856b-cba5-41ef-a454-a47facf116cd"),
                            Name: codataccounting.String("Jorge Grady"),
                        },
                        Currency: codataccounting.String("officia"),
                        CurrencyRate: codataccounting.Float64(4413.58),
                        ID: codataccounting.String("562873c7-dd9e-4faf-83dc-623620f3138f"),
                        Note: codataccounting.String("nesciunt"),
                        PaidOnDate: codataccounting.String("doloremque"),
                        Reference: codataccounting.String("at"),
                        TotalAmount: codataccounting.Float64(9458.52),
                    },
                },
            },
            Reference: codataccounting.String("sequi"),
            SourceModifiedDate: codataccounting.String("temporibus"),
            SubTotal: 7364.8,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "magni": map[string]interface{}{
                        "earum": "similique",
                    },
                },
            },
            TaxAmount: 6633.25,
            TotalAmount: 3507.98,
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(415121),
    }

    res, err := s.DirectCosts.CreateDirectCost(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDirectCostResponse != nil {
        // handle response
    }
}
```

## DownloadDirectCostAttachment

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

    res, err := s.DirectCosts.DownloadDirectCostAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

## GetCreateDirectCostsModel

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

    res, err := s.DirectCosts.GetCreateDirectCostsModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## GetDirectCost

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

    res, err := s.DirectCosts.GetDirectCost(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectCost != nil {
        // handle response
    }
}
```

## GetDirectCostAttachment

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

    res, err := s.DirectCosts.GetDirectCostAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Attachment != nil {
        // handle response
    }
}
```

## GetDirectCosts

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
    req := operations.GetDirectCostsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("nostrum"),
    }

    res, err := s.DirectCosts.GetDirectCosts(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectCosts != nil {
        // handle response
    }
}
```

## ListDirectCostAttachments

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

    res, err := s.DirectCosts.ListDirectCostAttachments(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachmentsDataset != nil {
        // handle response
    }
}
```

## UploadDirectCostAttachment

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
            Content: []byte("delectus"),
            RequestBody: "quidem",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        DirectCostID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.DirectCosts.UploadDirectCostAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
