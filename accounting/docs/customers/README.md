# Customers

## Overview

Customers

### Available Operations

* [Create](#create) - Create customer
* [DownloadAttachment](#downloadattachment) - Download customer attachment
* [Get](#get) - Get customer
* [GetAttachment](#getattachment) - Get customer attachment
* [GetCreateUpdateModel](#getcreateupdatemodel) - Get create/update customer model
* [List](#list) - List customers
* [ListAttachments](#listattachments) - List customer attachments
* [Update](#update) - Update customer

## Create

Posts an individual customer for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update customer model](https://docs.codat.io/accounting-api#/operations/get-create-update-customers-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) to see which integrations support this endpoint.


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
    res, err := s.Customers.Create(ctx, operations.CreateCustomerRequest{
        Customer: &shared.Customer{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("Kassandrafield"),
                    Country: codataccounting.String("Ireland"),
                    Line1: codataccounting.String("aut"),
                    Line2: codataccounting.String("nisi"),
                    PostalCode: codataccounting.String("53357"),
                    Region: codataccounting.String("debitis"),
                    Type: shared.AddressTypeUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("West Rorystad"),
                    Country: codataccounting.String("Martinique"),
                    Line1: codataccounting.String("magnam"),
                    Line2: codataccounting.String("cupiditate"),
                    PostalCode: codataccounting.String("95102-7224"),
                    Region: codataccounting.String("hic"),
                    Type: shared.AddressTypeBilling,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Monroefort"),
                    Country: codataccounting.String("Australia"),
                    Line1: codataccounting.String("porro"),
                    Line2: codataccounting.String("vel"),
                    PostalCode: codataccounting.String("92418-7122"),
                    Region: codataccounting.String("incidunt"),
                    Type: shared.AddressTypeDelivery,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Mrazworth"),
                    Country: codataccounting.String("Jersey"),
                    Line1: codataccounting.String("rem"),
                    Line2: codataccounting.String("est"),
                    PostalCode: codataccounting.String("70884"),
                    Region: codataccounting.String("laborum"),
                    Type: shared.AddressTypeBilling,
                },
            },
            ContactName: codataccounting.String("soluta"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Fort Leifmouth"),
                        Country: codataccounting.String("Bulgaria"),
                        Line1: codataccounting.String("ea"),
                        Line2: codataccounting.String("architecto"),
                        PostalCode: codataccounting.String("70835-8147"),
                        Region: codataccounting.String("exercitationem"),
                        Type: shared.AddressTypeUnknown,
                    },
                    Email: codataccounting.String("Krystel.Jakubowski71@yahoo.com"),
                    ModifiedDate: codataccounting.String("modi"),
                    Name: codataccounting.String("Jon Bashirian"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "iusto",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                        shared.PhoneNumbersitems{
                            Number: "odit",
                            Type: shared.PhoneNumberTypeFax,
                        },
                        shared.PhoneNumbersitems{
                            Number: "ducimus",
                            Type: shared.PhoneNumberTypeFax,
                        },
                        shared.PhoneNumbersitems{
                            Number: "ducimus",
                            Type: shared.PhoneNumberTypeFax,
                        },
                    },
                    Status: shared.CustomerStatusArchived,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("East Dorcasboro"),
                        Country: codataccounting.String("Saint Helena"),
                        Line1: codataccounting.String("inventore"),
                        Line2: codataccounting.String("ducimus"),
                        PostalCode: codataccounting.String("04558"),
                        Region: codataccounting.String("necessitatibus"),
                        Type: shared.AddressTypeDelivery,
                    },
                    Email: codataccounting.String("Stanley_Zboncak@hotmail.com"),
                    ModifiedDate: codataccounting.String("quam"),
                    Name: codataccounting.String("Lee Steuber DDS"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "pariatur",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                        shared.PhoneNumbersitems{
                            Number: "amet",
                            Type: shared.PhoneNumberTypeUnknown,
                        },
                        shared.PhoneNumbersitems{
                            Number: "quasi",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                        shared.PhoneNumbersitems{
                            Number: "rerum",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                    },
                    Status: shared.CustomerStatusArchived,
                },
            },
            CustomerName: codataccounting.String("aliquam"),
            DefaultCurrency: codataccounting.String("voluptates"),
            EmailAddress: codataccounting.String("alias"),
            ID: codataccounting.String("80aa1041-86ec-4759-a02f-3702c5c8e2d3"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("consequatur"),
            Phone: codataccounting.String("1-782-202-9622 x40479"),
            RegistrationNumber: codataccounting.String("sequi"),
            SourceModifiedDate: codataccounting.String("ducimus"),
            Status: shared.CustomerStatusUnknown,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "ut": map[string]interface{}{
                        "sed": "quas",
                        "aspernatur": "laudantium",
                    },
                    "fugit": map[string]interface{}{
                        "reiciendis": "nulla",
                    },
                    "libero": map[string]interface{}{
                        "hic": "eum",
                    },
                },
            },
            TaxNumber: codataccounting.String("sint"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(909351),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateCustomerResponse != nil {
        // handle response
    }
}
```

## DownloadAttachment

﻿Download customer attachment.

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
    res, err := s.Customers.DownloadAttachment(ctx, operations.DownloadCustomerAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "veniam",
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

﻿Gets a single customer corresponding to the given ID.

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
    res, err := s.Customers.Get(ctx, operations.GetCustomerRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CustomerID: "unde",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Customer != nil {
        // handle response
    }
}
```

## GetAttachment

﻿Get  customer attachment.

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
    res, err := s.Customers.GetAttachment(ctx, operations.GetCustomerAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "consequuntur",
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

﻿Get create/update customer model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support creating and updating customers.

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
    res, err := s.Customers.GetCreateUpdateModel(ctx, operations.GetCreateUpdateCustomersModelRequest{
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

﻿Gets the latest customers for a company, with pagination.

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
    res, err := s.Customers.List(ctx, operations.ListCustomersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("laboriosam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Customers != nil {
        // handle response
    }
}
```

## ListAttachments

﻿List customer attachments

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
    res, err := s.Customers.ListAttachments(ctx, operations.ListCustomerAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "iusto",
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

﻿Posts an updated customer for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update customer model](https://docs.codat.io/accounting-api#/operations/get-create-update-customers-model).

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support updating customers.

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
    res, err := s.Customers.Update(ctx, operations.UpdateCustomerRequest{
        Customer: &shared.Customer{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("Annabelfort"),
                    Country: codataccounting.String("Saint Martin"),
                    Line1: codataccounting.String("voluptatum"),
                    Line2: codataccounting.String("pariatur"),
                    PostalCode: codataccounting.String("78213"),
                    Region: codataccounting.String("voluptatum"),
                    Type: shared.AddressTypeDelivery,
                },
                shared.Addressesitems{
                    City: codataccounting.String("East Fridastad"),
                    Country: codataccounting.String("Oman"),
                    Line1: codataccounting.String("quas"),
                    Line2: codataccounting.String("odit"),
                    PostalCode: codataccounting.String("50599-1430"),
                    Region: codataccounting.String("laborum"),
                    Type: shared.AddressTypeUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("West Jeromefield"),
                    Country: codataccounting.String("American Samoa"),
                    Line1: codataccounting.String("numquam"),
                    Line2: codataccounting.String("numquam"),
                    PostalCode: codataccounting.String("90329-6305"),
                    Region: codataccounting.String("distinctio"),
                    Type: shared.AddressTypeDelivery,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Robbboro"),
                    Country: codataccounting.String("Aruba"),
                    Line1: codataccounting.String("vero"),
                    Line2: codataccounting.String("corporis"),
                    PostalCode: codataccounting.String("00962-1494"),
                    Region: codataccounting.String("sed"),
                    Type: shared.AddressTypeBilling,
                },
            },
            ContactName: codataccounting.String("natus"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("East Trystan"),
                        Country: codataccounting.String("Nauru"),
                        Line1: codataccounting.String("minima"),
                        Line2: codataccounting.String("minus"),
                        PostalCode: codataccounting.String("81241-6474"),
                        Region: codataccounting.String("quas"),
                        Type: shared.AddressTypeBilling,
                    },
                    Email: codataccounting.String("Fay.Stokes97@hotmail.com"),
                    ModifiedDate: codataccounting.String("quis"),
                    Name: codataccounting.String("Helen Brown"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "accusamus",
                            Type: shared.PhoneNumberTypeFax,
                        },
                        shared.PhoneNumbersitems{
                            Number: "vero",
                            Type: shared.PhoneNumberTypeUnknown,
                        },
                        shared.PhoneNumbersitems{
                            Number: "ea",
                            Type: shared.PhoneNumberTypeMobile,
                        },
                        shared.PhoneNumbersitems{
                            Number: "aliquid",
                            Type: shared.PhoneNumberTypeFax,
                        },
                    },
                    Status: shared.CustomerStatusActive,
                },
            },
            CustomerName: codataccounting.String("consequuntur"),
            DefaultCurrency: codataccounting.String("accusantium"),
            EmailAddress: codataccounting.String("autem"),
            ID: codataccounting.String("dab75005-2a56-447e-9c43-9ed8c4320f41"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("eos"),
            Phone: codataccounting.String("283-454-6735 x275"),
            RegistrationNumber: codataccounting.String("quaerat"),
            SourceModifiedDate: codataccounting.String("nobis"),
            Status: shared.CustomerStatusUnknown,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "unde": map[string]interface{}{
                        "magni": "modi",
                        "atque": "blanditiis",
                        "quibusdam": "odio",
                        "unde": "ad",
                    },
                    "officia": map[string]interface{}{
                        "incidunt": "aspernatur",
                        "asperiores": "maxime",
                        "dolore": "accusantium",
                    },
                    "corporis": map[string]interface{}{
                        "laboriosam": "omnis",
                        "tenetur": "vel",
                    },
                },
            },
            TaxNumber: codataccounting.String("iste"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "animi",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(60491),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCustomerResponse != nil {
        // handle response
    }
}
```
