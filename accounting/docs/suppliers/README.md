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

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers) to see which integrations support this endpoint.


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
    res, err := s.Suppliers.Create(ctx, operations.CreateSupplierRequest{
        Supplier: &shared.Supplier{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("Kirkland"),
                    Country: codataccounting.String("Netherlands"),
                    Line1: codataccounting.String("doloremque"),
                    Line2: codataccounting.String("autem"),
                    PostalCode: codataccounting.String("25088"),
                    Region: codataccounting.String("laudantium"),
                    Type: shared.AddressTypeDelivery,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Ankundingfort"),
                    Country: codataccounting.String("Kyrgyz Republic"),
                    Line1: codataccounting.String("repellendus"),
                    Line2: codataccounting.String("beatae"),
                    PostalCode: codataccounting.String("42514"),
                    Region: codataccounting.String("facilis"),
                    Type: shared.AddressTypeUnknown,
                },
            },
            ContactName: codataccounting.String("cumque"),
            DefaultCurrency: codataccounting.String("doloribus"),
            EmailAddress: codataccounting.String("minima"),
            ID: codataccounting.String("e6cb6eba-be5e-40b9-9f3b-1358d6a87bb7"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("est"),
            Phone: codataccounting.String("1-869-446-8407 x70355"),
            RegistrationNumber: codataccounting.String("sit"),
            SourceModifiedDate: codataccounting.String("dignissimos"),
            Status: shared.SupplierStatusArchived,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "praesentium": map[string]interface{}{
                        "incidunt": "incidunt",
                        "vitae": "incidunt",
                        "nostrum": "explicabo",
                    },
                    "culpa": map[string]interface{}{
                        "voluptatibus": "ipsa",
                        "quasi": "sapiente",
                        "dolorem": "quaerat",
                    },
                    "incidunt": map[string]interface{}{
                        "cumque": "vel",
                    },
                },
            },
            SupplierName: codataccounting.String("inventore"),
            TaxNumber: codataccounting.String("quidem"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(894398),
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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Suppliers.DownloadAttachment(ctx, operations.DownloadSupplierAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        SupplierID: "quae",
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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Suppliers.Get(ctx, operations.GetSupplierRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SupplierID: "ipsum",
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

﻿Get supplier attachment.

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
    res, err := s.Suppliers.GetAttachment(ctx, operations.GetSupplierAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        SupplierID: "nesciunt",
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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Suppliers.List(ctx, operations.ListSuppliersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("distinctio"),
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
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Suppliers.ListAttachments(ctx, operations.ListSupplierAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        SupplierID: "dolorum",
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
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers) for integrations that support updating suppliers.

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
    res, err := s.Suppliers.Update(ctx, operations.UpdateSupplierRequest{
        Supplier: &shared.Supplier{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("The Hammocks"),
                    Country: codataccounting.String("Guinea"),
                    Line1: codataccounting.String("neque"),
                    Line2: codataccounting.String("eos"),
                    PostalCode: codataccounting.String("43139-5315"),
                    Region: codataccounting.String("corporis"),
                    Type: shared.AddressTypeUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Towson"),
                    Country: codataccounting.String("Christmas Island"),
                    Line1: codataccounting.String("laudantium"),
                    Line2: codataccounting.String("enim"),
                    PostalCode: codataccounting.String("88211-1109"),
                    Region: codataccounting.String("earum"),
                    Type: shared.AddressTypeBilling,
                },
                shared.Addressesitems{
                    City: codataccounting.String("South Electa"),
                    Country: codataccounting.String("Saint Lucia"),
                    Line1: codataccounting.String("numquam"),
                    Line2: codataccounting.String("quae"),
                    PostalCode: codataccounting.String("19763-8760"),
                    Region: codataccounting.String("eum"),
                    Type: shared.AddressTypeBilling,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Littlehaven"),
                    Country: codataccounting.String("Cambodia"),
                    Line1: codataccounting.String("cupiditate"),
                    Line2: codataccounting.String("dicta"),
                    PostalCode: codataccounting.String("97500-6166"),
                    Region: codataccounting.String("atque"),
                    Type: shared.AddressTypeBilling,
                },
            },
            ContactName: codataccounting.String("aliquam"),
            DefaultCurrency: codataccounting.String("perspiciatis"),
            EmailAddress: codataccounting.String("labore"),
            ID: codataccounting.String("79edd4fc-f7b5-40cf-87f0-8f39271076a2"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("dolore"),
            Phone: codataccounting.String("407.690.5799 x0051"),
            RegistrationNumber: codataccounting.String("saepe"),
            SourceModifiedDate: codataccounting.String("laudantium"),
            Status: shared.SupplierStatusActive,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "blanditiis": map[string]interface{}{
                        "occaecati": "natus",
                        "voluptas": "optio",
                    },
                    "totam": map[string]interface{}{
                        "odit": "eos",
                        "libero": "eveniet",
                        "aut": "similique",
                        "ipsum": "maxime",
                    },
                    "tenetur": map[string]interface{}{
                        "voluptate": "blanditiis",
                        "sint": "dolorem",
                    },
                    "soluta": map[string]interface{}{
                        "fugit": "neque",
                        "asperiores": "corrupti",
                        "autem": "autem",
                        "alias": "eaque",
                    },
                },
            },
            SupplierName: codataccounting.String("minus"),
            TaxNumber: codataccounting.String("commodi"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        SupplierID: "inventore",
        TimeoutInMinutes: codataccounting.Int(784398),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateSupplierResponse != nil {
        // handle response
    }
}
```
