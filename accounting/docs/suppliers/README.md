# Suppliers

## Overview

Suppliers

### Available Operations

* [Create](#create) - Create supplier
* [DownloadAttachment](#downloadattachment) - Download supplier attachment
* [Get](#get) - Get supplier
* [GetAttachment](#getattachment) - Get supplier attachment
* [GetCreateUpdateModel](#getcreateupdatemodel) - Get create/update supplier model
* [List](#list) - List suppliers
* [ListAttachments](#listattachments) - List supplier attachments
* [Update](#update) - Update supplier

## Create

Push suppliers

Required data may vary by integration. To see what data to post, first call [Get create/update supplier model](https://docs.codat.io/accounting-api#/operations/get-create-update-suppliers-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers) for integrations that support creating suppliers.

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
    res, err := s.Suppliers.Create(ctx, operations.CreateSupplierRequest{
        Supplier: &shared.Supplier{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("Thousand Oaks"),
                    Country: codataccounting.String("Papua New Guinea"),
                    Line1: codataccounting.String("amet"),
                    Line2: codataccounting.String("tempore"),
                    PostalCode: codataccounting.String("81317"),
                    Region: codataccounting.String("adipisci"),
                    Type: shared.AddressTypeEnumUnknown,
                },
            },
            ContactName: codataccounting.String("alias"),
            DefaultCurrency: codataccounting.String("occaecati"),
            EmailAddress: codataccounting.String("perspiciatis"),
            ID: codataccounting.String("83663c66-dcbb-47df-acb0-9c8b408e0713"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("molestiae"),
            Phone: codataccounting.String("489.499.8000 x854"),
            RegistrationNumber: codataccounting.String("praesentium"),
            SourceModifiedDate: codataccounting.String("aperiam"),
            Status: shared.SupplierStatusEnumArchived,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "doloremque": map[string]interface{}{
                        "eius": "odio",
                        "rerum": "provident",
                        "nostrum": "perferendis",
                        "aliquam": "accusantium",
                    },
                },
            },
            SupplierName: codataccounting.String("possimus"),
            TaxNumber: codataccounting.String("vel"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(796063),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateSupplierResponse != nil {
        // handle response
    }
}
```

## DownloadAttachment

Download supplier attachment

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
    res, err := s.Suppliers.DownloadAttachment(ctx, operations.DownloadSupplierAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        SupplierID: "8a210b68-6988-11ed-a1eb-0242ac120002",
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

Gets a single supplier corresponding to the given ID.

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
    res, err := s.Suppliers.Get(ctx, operations.GetSupplierRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SupplierID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Supplier != nil {
        // handle response
    }
}
```

## GetAttachment

Get supplier attachment

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
    res, err := s.Suppliers.GetAttachment(ctx, operations.GetSupplierAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        SupplierID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Attachment != nil {
        // handle response
    }
}
```

## GetCreateUpdateModel

Get create/update supplier model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers) for integrations that support creating and updating suppliers.

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
    res, err := s.Suppliers.GetCreateUpdateModel(ctx, operations.GetCreateUpdateSuppliersModelRequest{
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

Gets the latest suppliers for a company, with pagination

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
    res, err := s.Suppliers.List(ctx, operations.ListSuppliersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("blanditiis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Suppliers != nil {
        // handle response
    }
}
```

## ListAttachments

Get supplier attachments

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
    res, err := s.Suppliers.ListAttachments(ctx, operations.ListSupplierAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        SupplierID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachmentsDataset != nil {
        // handle response
    }
}
```

## Update

Update supplier

Required data may vary by integration. To see what data to post, first call [Get create/update supplier model](https://docs.codat.io/accounting-api#/operations/get-create-update-suppliers-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers) for integrations that support updating suppliers.

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
    res, err := s.Suppliers.Update(ctx, operations.UpdateSupplierRequest{
        Supplier: &shared.Supplier{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("Lake Gabriellashire"),
                    Country: codataccounting.String("Afghanistan"),
                    Line1: codataccounting.String("perferendis"),
                    Line2: codataccounting.String("aspernatur"),
                    PostalCode: codataccounting.String("04820"),
                    Region: codataccounting.String("dolore"),
                    Type: shared.AddressTypeEnumBilling,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Mountain View"),
                    Country: codataccounting.String("Armenia"),
                    Line1: codataccounting.String("alias"),
                    Line2: codataccounting.String("sit"),
                    PostalCode: codataccounting.String("98150-1458"),
                    Region: codataccounting.String("quidem"),
                    Type: shared.AddressTypeEnumUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Watersfurt"),
                    Country: codataccounting.String("Syrian Arab Republic"),
                    Line1: codataccounting.String("suscipit"),
                    Line2: codataccounting.String("ut"),
                    PostalCode: codataccounting.String("40961"),
                    Region: codataccounting.String("corporis"),
                    Type: shared.AddressTypeEnumUnknown,
                },
            },
            ContactName: codataccounting.String("alias"),
            DefaultCurrency: codataccounting.String("ratione"),
            EmailAddress: codataccounting.String("sed"),
            ID: codataccounting.String("3b2c09b9-2477-41f5-a69e-5b7ec7626649"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("possimus"),
            Phone: codataccounting.String("487-692-7981 x1448"),
            RegistrationNumber: codataccounting.String("sit"),
            SourceModifiedDate: codataccounting.String("expedita"),
            Status: shared.SupplierStatusEnumActive,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "repellat": map[string]interface{}{
                        "atque": "iure",
                        "nulla": "aliquid",
                        "asperiores": "similique",
                    },
                    "veniam": map[string]interface{}{
                        "vel": "earum",
                        "corrupti": "temporibus",
                        "libero": "sapiente",
                    },
                    "praesentium": map[string]interface{}{
                        "qui": "asperiores",
                    },
                },
            },
            SupplierName: codataccounting.String("blanditiis"),
            TaxNumber: codataccounting.String("nesciunt"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        SupplierID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TimeoutInMinutes: codataccounting.Int(721212),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateSupplierResponse != nil {
        // handle response
    }
}
```
