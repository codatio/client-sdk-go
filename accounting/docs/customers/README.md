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

﻿Posts an individual customer for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update customer model](https://docs.codat.io/accounting-api#/operations/get-create-update-customers-model).

> **Supported Integrations**
> 
> Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support creating customers.

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
                    City: codataccounting.String("Port Clevelandtown"),
                    Country: codataccounting.String("Turkmenistan"),
                    Line1: codataccounting.String("culpa"),
                    Line2: codataccounting.String("at"),
                    PostalCode: codataccounting.String("10296"),
                    Region: codataccounting.String("dolore"),
                    Type: shared.AddressTypeUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Abbyport"),
                    Country: codataccounting.String("Philippines"),
                    Line1: codataccounting.String("voluptatibus"),
                    Line2: codataccounting.String("sequi"),
                    PostalCode: codataccounting.String("37221"),
                    Region: codataccounting.String("quas"),
                    Type: shared.AddressTypeUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Burdetteland"),
                    Country: codataccounting.String("Vietnam"),
                    Line1: codataccounting.String("nulla"),
                    Line2: codataccounting.String("libero"),
                    PostalCode: codataccounting.String("94593"),
                    Region: codataccounting.String("unde"),
                    Type: shared.AddressTypeUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("South Paul"),
                    Country: codataccounting.String("Lesotho"),
                    Line1: codataccounting.String("ab"),
                    Line2: codataccounting.String("quo"),
                    PostalCode: codataccounting.String("58178-2135"),
                    Region: codataccounting.String("temporibus"),
                    Type: shared.AddressTypeUnknown,
                },
            },
            ContactName: codataccounting.String("amet"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Margaretestad"),
                        Country: codataccounting.String("Burkina Faso"),
                        Line1: codataccounting.String("placeat"),
                        Line2: codataccounting.String("rem"),
                        PostalCode: codataccounting.String("59914"),
                        Region: codataccounting.String("ullam"),
                        Type: shared.AddressTypeUnknown,
                    },
                    Email: codataccounting.String("Buster49@hotmail.com"),
                    ModifiedDate: codataccounting.String("quod"),
                    Name: codataccounting.String("Marjorie Funk"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "quaerat",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                    },
                    Status: shared.CustomerStatusArchived,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Howellland"),
                        Country: codataccounting.String("Mayotte"),
                        Line1: codataccounting.String("distinctio"),
                        Line2: codataccounting.String("cum"),
                        PostalCode: codataccounting.String("82083"),
                        Region: codataccounting.String("laborum"),
                        Type: shared.AddressTypeUnknown,
                    },
                    Email: codataccounting.String("Warren19@gmail.com"),
                    ModifiedDate: codataccounting.String("autem"),
                    Name: codataccounting.String("Edgar Cremin"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "velit",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                    },
                    Status: shared.CustomerStatusArchived,
                },
            },
            CustomerName: codataccounting.String("natus"),
            DefaultCurrency: codataccounting.String("minima"),
            EmailAddress: codataccounting.String("minus"),
            ID: codataccounting.String("9d237397-c785-4b5d-b4f5-00183febdf67"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("aliquid"),
            Phone: codataccounting.String("510.586.7430 x0316"),
            RegistrationNumber: codataccounting.String("ad"),
            SourceModifiedDate: codataccounting.String("nisi"),
            Status: shared.CustomerStatusUnknown,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "officiis": map[string]interface{}{
                        "minus": "tempora",
                        "sequi": "natus",
                        "saepe": "quibusdam",
                        "corrupti": "maxime",
                    },
                    "aliquam": map[string]interface{}{
                        "explicabo": "eaque",
                    },
                },
            },
            TaxNumber: codataccounting.String("hic"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(276828),
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.DownloadAttachment(ctx, operations.DownloadCustomerAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "quae",
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.Get(ctx, operations.GetCustomerRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CustomerID: "eos",
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.GetAttachment(ctx, operations.GetCustomerAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "eius",
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.List(ctx, operations.ListCustomersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: codataccounting.Int(1),
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("voluptatem"),
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.ListAttachments(ctx, operations.ListCustomerAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "temporibus",
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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.Update(ctx, operations.UpdateCustomerRequest{
        Customer: &shared.Customer{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("New Jayde"),
                    Country: codataccounting.String("Nigeria"),
                    Line1: codataccounting.String("porro"),
                    Line2: codataccounting.String("voluptas"),
                    PostalCode: codataccounting.String("27537-1668"),
                    Region: codataccounting.String("magni"),
                    Type: shared.AddressTypeUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Joanhaven"),
                    Country: codataccounting.String("Lebanon"),
                    Line1: codataccounting.String("unde"),
                    Line2: codataccounting.String("ad"),
                    PostalCode: codataccounting.String("62198-2033"),
                    Region: codataccounting.String("laboriosam"),
                    Type: shared.AddressTypeBilling,
                },
            },
            ContactName: codataccounting.String("tenetur"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("O'Konton"),
                        Country: codataccounting.String("Argentina"),
                        Line1: codataccounting.String("voluptas"),
                        Line2: codataccounting.String("pariatur"),
                        PostalCode: codataccounting.String("01252"),
                        Region: codataccounting.String("ullam"),
                        Type: shared.AddressTypeUnknown,
                    },
                    Email: codataccounting.String("Aurelia.Mraz77@yahoo.com"),
                    ModifiedDate: codataccounting.String("ipsum"),
                    Name: codataccounting.String("Carl Pollich IV"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "incidunt",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                        shared.PhoneNumbersitems{
                            Number: "nisi",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                    },
                    Status: shared.CustomerStatusArchived,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("North Cloydton"),
                        Country: codataccounting.String("Austria"),
                        Line1: codataccounting.String("possimus"),
                        Line2: codataccounting.String("perferendis"),
                        PostalCode: codataccounting.String("28760"),
                        Region: codataccounting.String("beatae"),
                        Type: shared.AddressTypeDelivery,
                    },
                    Email: codataccounting.String("Mae.Willms66@yahoo.com"),
                    ModifiedDate: codataccounting.String("commodi"),
                    Name: codataccounting.String("Tom Gutkowski"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "id",
                            Type: shared.PhoneNumberTypeUnknown,
                        },
                        shared.PhoneNumbersitems{
                            Number: "molestias",
                            Type: shared.PhoneNumberTypeUnknown,
                        },
                    },
                    Status: shared.CustomerStatusActive,
                },
            },
            CustomerName: codataccounting.String("occaecati"),
            DefaultCurrency: codataccounting.String("eos"),
            EmailAddress: codataccounting.String("veniam"),
            ID: codataccounting.String("3c8962f4-896b-4f51-a465-2d3c343d6177"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("atque"),
            Phone: codataccounting.String("1-935-212-4413 x9410"),
            RegistrationNumber: codataccounting.String("iste"),
            SourceModifiedDate: codataccounting.String("aut"),
            Status: shared.CustomerStatusActive,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "error": map[string]interface{}{
                        "ipsa": "dolore",
                    },
                    "labore": map[string]interface{}{
                        "ullam": "quibusdam",
                        "recusandae": "ad",
                        "omnis": "mollitia",
                    },
                    "placeat": map[string]interface{}{
                        "ducimus": "eaque",
                        "aliquid": "ea",
                    },
                    "odio": map[string]interface{}{
                        "quisquam": "delectus",
                    },
                },
            },
            TaxNumber: codataccounting.String("et"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "optio",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(953679),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCustomerResponse != nil {
        // handle response
    }
}
```
