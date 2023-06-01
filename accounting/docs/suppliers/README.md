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
                    City: codataccounting.String("Stephaniastead"),
                    Country: codataccounting.String("Swaziland"),
                    Line1: codataccounting.String("laudantium"),
                    Line2: codataccounting.String("eos"),
                    PostalCode: codataccounting.String("38203-9334"),
                    Region: codataccounting.String("aliquid"),
                    Type: shared.AddressTypeDelivery,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Port Rosalinda"),
                    Country: codataccounting.String("Turkmenistan"),
                    Line1: codataccounting.String("cum"),
                    Line2: codataccounting.String("consequatur"),
                    PostalCode: codataccounting.String("54807"),
                    Region: codataccounting.String("maxime"),
                    Type: shared.AddressTypeUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Lake Arnoldo"),
                    Country: codataccounting.String("Martinique"),
                    Line1: codataccounting.String("id"),
                    Line2: codataccounting.String("expedita"),
                    PostalCode: codataccounting.String("45919"),
                    Region: codataccounting.String("quisquam"),
                    Type: shared.AddressTypeDelivery,
                },
            },
            ContactName: codataccounting.String("maiores"),
            DefaultCurrency: codataccounting.String("laudantium"),
            EmailAddress: codataccounting.String("beatae"),
            ID: codataccounting.String("ddf7e088-f74e-4f54-8921-6e8926313bb6"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("hic"),
            Phone: codataccounting.String("375-814-0006 x46446"),
            RegistrationNumber: codataccounting.String("fugiat"),
            SourceModifiedDate: codataccounting.String("laboriosam"),
            Status: shared.SupplierStatusArchived,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "recusandae": map[string]interface{}{
                        "pariatur": "excepturi",
                    },
                },
            },
            SupplierName: codataccounting.String("fugiat"),
            TaxNumber: codataccounting.String("ipsum"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(730689),
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
        SupplierID: "voluptas",
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
        SupplierID: "aliquid",
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
        SupplierID: "perferendis",
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
        Query: codataccounting.String("nesciunt"),
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
        SupplierID: "non",
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
                    City: codataccounting.String("Carterland"),
                    Country: codataccounting.String("Northern Mariana Islands"),
                    Line1: codataccounting.String("id"),
                    Line2: codataccounting.String("ab"),
                    PostalCode: codataccounting.String("38112-4885"),
                    Region: codataccounting.String("libero"),
                    Type: shared.AddressTypeUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Enterprise"),
                    Country: codataccounting.String("Honduras"),
                    Line1: codataccounting.String("quasi"),
                    Line2: codataccounting.String("ducimus"),
                    PostalCode: codataccounting.String("84456"),
                    Region: codataccounting.String("perspiciatis"),
                    Type: shared.AddressTypeBilling,
                },
            },
            ContactName: codataccounting.String("nobis"),
            DefaultCurrency: codataccounting.String("facilis"),
            EmailAddress: codataccounting.String("sequi"),
            ID: codataccounting.String("98788398-eba1-4bbf-b143-356f6349a164"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("explicabo"),
            Phone: codataccounting.String("671-217-9347 x630"),
            RegistrationNumber: codataccounting.String("vel"),
            SourceModifiedDate: codataccounting.String("enim"),
            Status: shared.SupplierStatusUnknown,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "beatae": map[string]interface{}{
                        "atque": "optio",
                        "culpa": "excepturi",
                    },
                    "et": map[string]interface{}{
                        "odit": "reiciendis",
                        "voluptatem": "veniam",
                    },
                    "consequuntur": map[string]interface{}{
                        "dolor": "quia",
                        "harum": "sequi",
                    },
                },
            },
            SupplierName: codataccounting.String("quae"),
            TaxNumber: codataccounting.String("porro"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        SupplierID: "harum",
        TimeoutInMinutes: codataccounting.Int(868742),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateSupplierResponse != nil {
        // handle response
    }
}
```
