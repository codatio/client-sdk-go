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
                    City: codataccounting.String("Blickland"),
                    Country: codataccounting.String("New Caledonia"),
                    Line1: codataccounting.String("nihil"),
                    Line2: codataccounting.String("commodi"),
                    PostalCode: codataccounting.String("48210-0859"),
                    Region: codataccounting.String("sed"),
                    Type: shared.AddressTypeDelivery,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Nataliatown"),
                    Country: codataccounting.String("Spain"),
                    Line1: codataccounting.String("dignissimos"),
                    Line2: codataccounting.String("quaerat"),
                    PostalCode: codataccounting.String("81647"),
                    Region: codataccounting.String("molestiae"),
                    Type: shared.AddressTypeDelivery,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Fort Mariannatown"),
                    Country: codataccounting.String("Turkey"),
                    Line1: codataccounting.String("quam"),
                    Line2: codataccounting.String("iste"),
                    PostalCode: codataccounting.String("65779-3904"),
                    Region: codataccounting.String("sint"),
                    Type: shared.AddressTypeDelivery,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Flossietown"),
                    Country: codataccounting.String("Reunion"),
                    Line1: codataccounting.String("quasi"),
                    Line2: codataccounting.String("earum"),
                    PostalCode: codataccounting.String("04948"),
                    Region: codataccounting.String("dolore"),
                    Type: shared.AddressTypeUnknown,
                },
            },
            ContactName: codataccounting.String("molestias"),
            DefaultCurrency: codataccounting.String("ea"),
            EmailAddress: codataccounting.String("odio"),
            ID: codataccounting.String("13bacce0-72ab-4d61-918d-279c10c18516"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Phone: codataccounting.String("01224 658 999"),
            RegistrationNumber: codataccounting.String("in"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.SupplierStatusArchived,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "fugit": map[string]interface{}{
                        "fugit": "ab",
                        "quia": "esse",
                    },
                    "sed": map[string]interface{}{
                        "odit": "quos",
                        "doloribus": "officia",
                    },
                    "ullam": map[string]interface{}{
                        "ratione": "natus",
                    },
                    "vel": map[string]interface{}{
                        "blanditiis": "eum",
                    },
                },
            },
            SupplierName: codataccounting.String("esse"),
            TaxNumber: codataccounting.String("debitis"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(456179),
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
        SupplierID: "aspernatur",
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
        SupplierID: "nam",
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
        SupplierID: "neque",
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
        Query: codataccounting.String("laborum"),
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
        SupplierID: "autem",
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
                    City: codataccounting.String("East Edisoncester"),
                    Country: codataccounting.String("Barbados"),
                    Line1: codataccounting.String("minima"),
                    Line2: codataccounting.String("molestiae"),
                    PostalCode: codataccounting.String("57739-9416"),
                    Region: codataccounting.String("ad"),
                    Type: shared.AddressTypeUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Jarretland"),
                    Country: codataccounting.String("Taiwan"),
                    Line1: codataccounting.String("unde"),
                    Line2: codataccounting.String("provident"),
                    PostalCode: codataccounting.String("34064-8990"),
                    Region: codataccounting.String("commodi"),
                    Type: shared.AddressTypeBilling,
                },
            },
            ContactName: codataccounting.String("nam"),
            DefaultCurrency: codataccounting.String("vel"),
            EmailAddress: codataccounting.String("impedit"),
            ID: codataccounting.String("cb2822b4-a985-40ed-af4a-1e9c4ae55140"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Phone: codataccounting.String("(877) 492-8687"),
            RegistrationNumber: codataccounting.String("corporis"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.SupplierStatusUnknown,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "earum": map[string]interface{}{
                        "consequatur": "nesciunt",
                    },
                    "porro": map[string]interface{}{
                        "asperiores": "aut",
                    },
                },
            },
            SupplierName: codataccounting.String("consequuntur"),
            TaxNumber: codataccounting.String("natus"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        SupplierID: "aliquam",
        TimeoutInMinutes: codataccounting.Int(98946),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateSupplierResponse != nil {
        // handle response
    }
}
```
