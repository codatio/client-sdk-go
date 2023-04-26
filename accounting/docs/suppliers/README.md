# Suppliers

## Overview

Suppliers

### Available Operations

* [CreateSupplier](#createsupplier) - Create suppliers
* [DownloadSupplierAttachment](#downloadsupplierattachment) - Download supplier attachment
* [GetCreateUpdateSuppliersModel](#getcreateupdatesuppliersmodel) - Get create/update supplier model
* [GetSupplier](#getsupplier) - Get supplier
* [GetSupplierAttachment](#getsupplierattachment) - Get supplier attachment
* [ListSupplierAttachments](#listsupplierattachments) - List supplier attachments
* [ListSuppliers](#listsuppliers) - List suppliers
* [PutSupplier](#putsupplier) - Update supplier

## CreateSupplier

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
    req := operations.CreateSupplierRequest{
        Supplier: &shared.Supplier{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("South Gavin"),
                    Country: codataccounting.String("Saint Lucia"),
                    Line1: codataccounting.String("explicabo"),
                    Line2: codataccounting.String("expedita"),
                    PostalCode: codataccounting.String("89844-4509"),
                    Region: codataccounting.String("temporibus"),
                    Type: shared.AddressTypeEnumUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Blockview"),
                    Country: codataccounting.String("Oman"),
                    Line1: codataccounting.String("molestiae"),
                    Line2: codataccounting.String("harum"),
                    PostalCode: codataccounting.String("24520-2256"),
                    Region: codataccounting.String("repellat"),
                    Type: shared.AddressTypeEnumBilling,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Jeniferton"),
                    Country: codataccounting.String("Ukraine"),
                    Line1: codataccounting.String("earum"),
                    Line2: codataccounting.String("ipsa"),
                    PostalCode: codataccounting.String("50506"),
                    Region: codataccounting.String("dolores"),
                    Type: shared.AddressTypeEnumDelivery,
                },
            },
            ContactName: codataccounting.String("culpa"),
            DefaultCurrency: codataccounting.String("fugit"),
            EmailAddress: codataccounting.String("nemo"),
            ID: codataccounting.String("ee6c75af-8a60-4a7a-a346-e0979e5afe60"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("culpa"),
            Phone: codataccounting.String("776-533-8955 x34331"),
            RegistrationNumber: codataccounting.String("deserunt"),
            SourceModifiedDate: codataccounting.String("iste"),
            Status: shared.SupplierStatusEnumArchived,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "eveniet": map[string]interface{}{
                        "quae": "voluptates",
                        "impedit": "sunt",
                    },
                    "optio": map[string]interface{}{
                        "occaecati": "officia",
                        "consectetur": "excepturi",
                    },
                    "fuga": map[string]interface{}{
                        "ipsam": "fuga",
                        "magnam": "assumenda",
                        "nemo": "id",
                        "laboriosam": "nostrum",
                    },
                    "expedita": map[string]interface{}{
                        "fugiat": "exercitationem",
                        "veniam": "ea",
                    },
                },
            },
            SupplierName: codataccounting.String("aspernatur"),
            TaxNumber: codataccounting.String("assumenda"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(587240),
    }

    res, err := s.Suppliers.CreateSupplier(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateSupplierResponse != nil {
        // handle response
    }
}
```

## DownloadSupplierAttachment

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
    req := operations.DownloadSupplierAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        SupplierID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Suppliers.DownloadSupplierAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

## GetCreateUpdateSuppliersModel

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
    req := operations.GetCreateUpdateSuppliersModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Suppliers.GetCreateUpdateSuppliersModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## GetSupplier

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
    req := operations.GetSupplierRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        SupplierID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Suppliers.GetSupplier(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Supplier != nil {
        // handle response
    }
}
```

## GetSupplierAttachment

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
    req := operations.GetSupplierAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        SupplierID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Suppliers.GetSupplierAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Attachment != nil {
        // handle response
    }
}
```

## ListSupplierAttachments

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
    req := operations.ListSupplierAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        SupplierID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Suppliers.ListSupplierAttachments(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachmentsDataset != nil {
        // handle response
    }
}
```

## ListSuppliers

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
    req := operations.ListSuppliersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("expedita"),
    }

    res, err := s.Suppliers.ListSuppliers(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Suppliers != nil {
        // handle response
    }
}
```

## PutSupplier

Push supplier

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
    req := operations.PutSupplierRequest{
        Supplier: &shared.Supplier{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("Missoula"),
                    Country: codataccounting.String("Tuvalu"),
                    Line1: codataccounting.String("eos"),
                    Line2: codataccounting.String("facere"),
                    PostalCode: codataccounting.String("97933"),
                    Region: codataccounting.String("esse"),
                    Type: shared.AddressTypeEnumBilling,
                },
                shared.Addressesitems{
                    City: codataccounting.String("New Sabrynachester"),
                    Country: codataccounting.String("Malta"),
                    Line1: codataccounting.String("quam"),
                    Line2: codataccounting.String("ad"),
                    PostalCode: codataccounting.String("16550-1516"),
                    Region: codataccounting.String("enim"),
                    Type: shared.AddressTypeEnumUnknown,
                },
            },
            ContactName: codataccounting.String("delectus"),
            DefaultCurrency: codataccounting.String("magnam"),
            EmailAddress: codataccounting.String("illo"),
            ID: codataccounting.String("cf6796ed-3d72-44c1-8f58-1e98cce3f716"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("aliquid"),
            Phone: codataccounting.String("286-282-6630"),
            RegistrationNumber: codataccounting.String("optio"),
            SourceModifiedDate: codataccounting.String("ex"),
            Status: shared.SupplierStatusEnumArchived,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "alias": map[string]interface{}{
                        "assumenda": "totam",
                        "minima": "explicabo",
                        "soluta": "ad",
                    },
                    "adipisci": map[string]interface{}{
                        "nesciunt": "eos",
                        "placeat": "blanditiis",
                        "cumque": "dignissimos",
                    },
                    "placeat": map[string]interface{}{
                        "eligendi": "esse",
                    },
                    "quasi": map[string]interface{}{
                        "accusamus": "inventore",
                    },
                },
            },
            SupplierName: codataccounting.String("voluptas"),
            TaxNumber: codataccounting.String("molestiae"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        SupplierID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TimeoutInMinutes: codataccounting.Int(219664),
    }

    res, err := s.Suppliers.PutSupplier(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PutSupplier200ApplicationJSONObject != nil {
        // handle response
    }
}
```
