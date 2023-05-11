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

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support creating customers.

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
    res, err := s.Customers.Create(ctx, operations.CreateCustomerRequest{
        Customer: &shared.Customer{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("Cartwrightstad"),
                    Country: codataccounting.String("Ukraine"),
                    Line1: codataccounting.String("assumenda"),
                    Line2: codataccounting.String("optio"),
                    PostalCode: codataccounting.String("94549"),
                    Region: codataccounting.String("adipisci"),
                    Type: shared.AddressTypeEnumUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Apopka"),
                    Country: codataccounting.String("Equatorial Guinea"),
                    Line1: codataccounting.String("rerum"),
                    Line2: codataccounting.String("nesciunt"),
                    PostalCode: codataccounting.String("28807-3440"),
                    Region: codataccounting.String("recusandae"),
                    Type: shared.AddressTypeEnumDelivery,
                },
                shared.Addressesitems{
                    City: codataccounting.String("East Martina"),
                    Country: codataccounting.String("Svalbard & Jan Mayen Islands"),
                    Line1: codataccounting.String("dolor"),
                    Line2: codataccounting.String("porro"),
                    PostalCode: codataccounting.String("91773-0293"),
                    Region: codataccounting.String("quod"),
                    Type: shared.AddressTypeEnumUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Reichelfort"),
                    Country: codataccounting.String("South Africa"),
                    Line1: codataccounting.String("alias"),
                    Line2: codataccounting.String("deserunt"),
                    PostalCode: codataccounting.String("15095"),
                    Region: codataccounting.String("nemo"),
                    Type: shared.AddressTypeEnumBilling,
                },
            },
            ContactName: codataccounting.String("reiciendis"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Freedaton"),
                        Country: codataccounting.String("Mali"),
                        Line1: codataccounting.String("natus"),
                        Line2: codataccounting.String("culpa"),
                        PostalCode: codataccounting.String("48950-4669"),
                        Region: codataccounting.String("quae"),
                        Type: shared.AddressTypeEnumUnknown,
                    },
                    Email: codataccounting.String("Cristal_Fisher48@hotmail.com"),
                    ModifiedDate: codataccounting.String("unde"),
                    Name: codataccounting.String("Alton McKenzie"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "aut",
                            Type: shared.PhoneNumberTypeEnumLandline,
                        },
                        shared.PhoneNumbersitems{
                            Number: "quia",
                            Type: shared.PhoneNumberTypeEnumFax,
                        },
                    },
                    Status: shared.CustomerStatusEnumArchived,
                },
            },
            CustomerName: codataccounting.String("qui"),
            DefaultCurrency: codataccounting.String("commodi"),
            EmailAddress: codataccounting.String("a"),
            ID: codataccounting.String("d368ba92-16bc-4b41-9835-c73641723133"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("recusandae"),
            Phone: codataccounting.String("(802) 568-3142 x77862"),
            RegistrationNumber: codataccounting.String("molestias"),
            SourceModifiedDate: codataccounting.String("dolores"),
            Status: shared.CustomerStatusEnumUnknown,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "minus": map[string]interface{}{
                        "odit": "maxime",
                        "aspernatur": "magni",
                    },
                    "minus": map[string]interface{}{
                        "ipsam": "sequi",
                        "quaerat": "accusantium",
                    },
                },
            },
            TaxNumber: codataccounting.String("incidunt"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(583959),
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

Download customer attachment

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
    res, err := s.Customers.DownloadAttachment(ctx, operations.DownloadCustomerAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "minima",
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

Gets a single customer corresponding to the given ID.

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
    res, err := s.Customers.Get(ctx, operations.GetCustomerRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CustomerID: "quo",
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

Get  customer attachment

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
    res, err := s.Customers.GetAttachment(ctx, operations.GetCustomerAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "quis",
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

Get create/update customer model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support creating and updating customers.

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

Gets the latest customers for a company, with pagination

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
    res, err := s.Customers.List(ctx, operations.ListCustomersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("facere"),
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

List customer attachments

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
    res, err := s.Customers.ListAttachments(ctx, operations.ListCustomerAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "quidem",
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

Posts an updated customer for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update customer model](https://docs.codat.io/accounting-api#/operations/get-create-update-customers-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support updating customers.

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
    res, err := s.Customers.Update(ctx, operations.UpdateCustomerRequest{
        Customer: &shared.Customer{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("Port Eriberto"),
                    Country: codataccounting.String("Jersey"),
                    Line1: codataccounting.String("quo"),
                    Line2: codataccounting.String("vitae"),
                    PostalCode: codataccounting.String("26519-5661"),
                    Region: codataccounting.String("enim"),
                    Type: shared.AddressTypeEnumBilling,
                },
                shared.Addressesitems{
                    City: codataccounting.String("South San Francisco"),
                    Country: codataccounting.String("Sao Tome and Principe"),
                    Line1: codataccounting.String("quasi"),
                    Line2: codataccounting.String("sint"),
                    PostalCode: codataccounting.String("19688"),
                    Region: codataccounting.String("eum"),
                    Type: shared.AddressTypeEnumUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Williamsonfield"),
                    Country: codataccounting.String("Samoa"),
                    Line1: codataccounting.String("veniam"),
                    Line2: codataccounting.String("magnam"),
                    PostalCode: codataccounting.String("68300"),
                    Region: codataccounting.String("quis"),
                    Type: shared.AddressTypeEnumDelivery,
                },
            },
            ContactName: codataccounting.String("reiciendis"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Lesterview"),
                        Country: codataccounting.String("Burundi"),
                        Line1: codataccounting.String("voluptatem"),
                        Line2: codataccounting.String("voluptas"),
                        PostalCode: codataccounting.String("89160-6192"),
                        Region: codataccounting.String("quia"),
                        Type: shared.AddressTypeEnumDelivery,
                    },
                    Email: codataccounting.String("Lydia7@gmail.com"),
                    ModifiedDate: codataccounting.String("perferendis"),
                    Name: codataccounting.String("Andy Paucek"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "necessitatibus",
                            Type: shared.PhoneNumberTypeEnumFax,
                        },
                    },
                    Status: shared.CustomerStatusEnumUnknown,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("West Ethelyn"),
                        Country: codataccounting.String("Martinique"),
                        Line1: codataccounting.String("ea"),
                        Line2: codataccounting.String("fugiat"),
                        PostalCode: codataccounting.String("59595"),
                        Region: codataccounting.String("reprehenderit"),
                        Type: shared.AddressTypeEnumBilling,
                    },
                    Email: codataccounting.String("Mittie_Williamson13@gmail.com"),
                    ModifiedDate: codataccounting.String("nam"),
                    Name: codataccounting.String("Erik Stehr"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "deserunt",
                            Type: shared.PhoneNumberTypeEnumMobile,
                        },
                        shared.PhoneNumbersitems{
                            Number: "modi",
                            Type: shared.PhoneNumberTypeEnumFax,
                        },
                        shared.PhoneNumbersitems{
                            Number: "repellendus",
                            Type: shared.PhoneNumberTypeEnumMobile,
                        },
                        shared.PhoneNumbersitems{
                            Number: "unde",
                            Type: shared.PhoneNumberTypeEnumFax,
                        },
                    },
                    Status: shared.CustomerStatusEnumUnknown,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Port Gavinland"),
                        Country: codataccounting.String("Central African Republic"),
                        Line1: codataccounting.String("numquam"),
                        Line2: codataccounting.String("velit"),
                        PostalCode: codataccounting.String("21540"),
                        Region: codataccounting.String("cumque"),
                        Type: shared.AddressTypeEnumDelivery,
                    },
                    Email: codataccounting.String("Josiah19@yahoo.com"),
                    ModifiedDate: codataccounting.String("fuga"),
                    Name: codataccounting.String("Miss Don Dach"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "nesciunt",
                            Type: shared.PhoneNumberTypeEnumUnknown,
                        },
                        shared.PhoneNumbersitems{
                            Number: "beatae",
                            Type: shared.PhoneNumberTypeEnumUnknown,
                        },
                    },
                    Status: shared.CustomerStatusEnumUnknown,
                },
            },
            CustomerName: codataccounting.String("quo"),
            DefaultCurrency: codataccounting.String("libero"),
            EmailAddress: codataccounting.String("eaque"),
            ID: codataccounting.String("a0003eb2-2d9b-43a7-8d94-faa741c57d1f"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("vero"),
            Phone: codataccounting.String("(810) 408-1587 x17905"),
            RegistrationNumber: codataccounting.String("nostrum"),
            SourceModifiedDate: codataccounting.String("labore"),
            Status: shared.CustomerStatusEnumActive,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "tenetur": map[string]interface{}{
                        "necessitatibus": "necessitatibus",
                        "autem": "natus",
                        "quasi": "iure",
                    },
                },
            },
            TaxNumber: codataccounting.String("ex"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "error",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(535903),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCustomerResponse != nil {
        // handle response
    }
}
```
