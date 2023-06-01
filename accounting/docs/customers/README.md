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
                    City: codataccounting.String("South Carterboro"),
                    Country: codataccounting.String("Ireland"),
                    Line1: codataccounting.String("reprehenderit"),
                    Line2: codataccounting.String("quidem"),
                    PostalCode: codataccounting.String("62077-0396"),
                    Region: codataccounting.String("libero"),
                    Type: shared.AddressTypeUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Hicksville"),
                    Country: codataccounting.String("Greenland"),
                    Line1: codataccounting.String("perferendis"),
                    Line2: codataccounting.String("itaque"),
                    PostalCode: codataccounting.String("91165-2810"),
                    Region: codataccounting.String("itaque"),
                    Type: shared.AddressTypeDelivery,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Alfonsoside"),
                    Country: codataccounting.String("Niue"),
                    Line1: codataccounting.String("aliquam"),
                    Line2: codataccounting.String("quasi"),
                    PostalCode: codataccounting.String("19233-9530"),
                    Region: codataccounting.String("minima"),
                    Type: shared.AddressTypeBilling,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Wesley Chapel"),
                    Country: codataccounting.String("Uruguay"),
                    Line1: codataccounting.String("molestiae"),
                    Line2: codataccounting.String("amet"),
                    PostalCode: codataccounting.String("89329-8839"),
                    Region: codataccounting.String("deserunt"),
                    Type: shared.AddressTypeBilling,
                },
            },
            ContactName: codataccounting.String("minima"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("East Korbinstad"),
                        Country: codataccounting.String("Svalbard & Jan Mayen Islands"),
                        Line1: codataccounting.String("laborum"),
                        Line2: codataccounting.String("sapiente"),
                        PostalCode: codataccounting.String("21658-4225"),
                        Region: codataccounting.String("quos"),
                        Type: shared.AddressTypeDelivery,
                    },
                    Email: codataccounting.String("Samanta20@gmail.com"),
                    ModifiedDate: codataccounting.String("error"),
                    Name: codataccounting.String("Trevor Hauck"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "sunt",
                            Type: shared.PhoneNumberTypeMobile,
                        },
                        shared.PhoneNumbersitems{
                            Number: "similique",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                        shared.PhoneNumbersitems{
                            Number: "tempora",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                    },
                    Status: shared.CustomerStatusArchived,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Hicksville"),
                        Country: codataccounting.String("Bahrain"),
                        Line1: codataccounting.String("omnis"),
                        Line2: codataccounting.String("aut"),
                        PostalCode: codataccounting.String("68688"),
                        Region: codataccounting.String("velit"),
                        Type: shared.AddressTypeBilling,
                    },
                    Email: codataccounting.String("Roosevelt_Batz38@hotmail.com"),
                    ModifiedDate: codataccounting.String("esse"),
                    Name: codataccounting.String("Dianna Ruecker"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "vero",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                    },
                    Status: shared.CustomerStatusActive,
                },
            },
            CustomerName: codataccounting.String("asperiores"),
            DefaultCurrency: codataccounting.String("quasi"),
            EmailAddress: codataccounting.String("veniam"),
            ID: codataccounting.String("920c90d1-b490-41f2-bd89-c8a32639da5b"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("in"),
            Phone: codataccounting.String("560.375.5065 x2942"),
            RegistrationNumber: codataccounting.String("sequi"),
            SourceModifiedDate: codataccounting.String("nisi"),
            Status: shared.CustomerStatusActive,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "mollitia": map[string]interface{}{
                        "hic": "doloremque",
                        "id": "asperiores",
                        "rem": "quod",
                    },
                    "commodi": map[string]interface{}{
                        "beatae": "placeat",
                        "molestiae": "dolor",
                        "quia": "nulla",
                    },
                },
            },
            TaxNumber: codataccounting.String("occaecati"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(983596),
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
        CustomerID: "libero",
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
        CustomerID: "culpa",
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
        CustomerID: "tenetur",
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
        Query: codataccounting.String("molestias"),
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
        CustomerID: "magnam",
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
                    City: codataccounting.String("Lake Christ"),
                    Country: codataccounting.String("Norfolk Island"),
                    Line1: codataccounting.String("vero"),
                    Line2: codataccounting.String("quas"),
                    PostalCode: codataccounting.String("77308-5623"),
                    Region: codataccounting.String("dicta"),
                    Type: shared.AddressTypeUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Kshlerinberg"),
                    Country: codataccounting.String("Kazakhstan"),
                    Line1: codataccounting.String("blanditiis"),
                    Line2: codataccounting.String("dolore"),
                    PostalCode: codataccounting.String("52043-0600"),
                    Region: codataccounting.String("accusamus"),
                    Type: shared.AddressTypeBilling,
                },
            },
            ContactName: codataccounting.String("commodi"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Rockwall"),
                        Country: codataccounting.String("Liberia"),
                        Line1: codataccounting.String("dolorem"),
                        Line2: codataccounting.String("eum"),
                        PostalCode: codataccounting.String("22063-1657"),
                        Region: codataccounting.String("molestiae"),
                        Type: shared.AddressTypeBilling,
                    },
                    Email: codataccounting.String("Nils22@gmail.com"),
                    ModifiedDate: codataccounting.String("cupiditate"),
                    Name: codataccounting.String("Sheldon Stiedemann"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "deserunt",
                            Type: shared.PhoneNumberTypeFax,
                        },
                    },
                    Status: shared.CustomerStatusActive,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Pasco"),
                        Country: codataccounting.String("Ecuador"),
                        Line1: codataccounting.String("placeat"),
                        Line2: codataccounting.String("veniam"),
                        PostalCode: codataccounting.String("65336-6407"),
                        Region: codataccounting.String("reiciendis"),
                        Type: shared.AddressTypeDelivery,
                    },
                    Email: codataccounting.String("Salvador_Johnson@gmail.com"),
                    ModifiedDate: codataccounting.String("iste"),
                    Name: codataccounting.String("Mrs. Harvey Crooks"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "quae",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                        shared.PhoneNumbersitems{
                            Number: "iure",
                            Type: shared.PhoneNumberTypeMobile,
                        },
                    },
                    Status: shared.CustomerStatusArchived,
                },
            },
            CustomerName: codataccounting.String("et"),
            DefaultCurrency: codataccounting.String("perspiciatis"),
            EmailAddress: codataccounting.String("accusamus"),
            ID: codataccounting.String("a83d492e-d14b-48a2-8195-4545e955dcc1"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("praesentium"),
            Phone: codataccounting.String("(963) 601-7473 x168"),
            RegistrationNumber: codataccounting.String("explicabo"),
            SourceModifiedDate: codataccounting.String("nulla"),
            Status: shared.CustomerStatusArchived,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "quam": map[string]interface{}{
                        "incidunt": "similique",
                        "nobis": "culpa",
                        "ratione": "illum",
                    },
                    "sed": map[string]interface{}{
                        "aut": "voluptates",
                    },
                    "nulla": map[string]interface{}{
                        "dignissimos": "dolor",
                        "totam": "beatae",
                        "vitae": "laborum",
                        "beatae": "vitae",
                    },
                },
            },
            TaxNumber: codataccounting.String("veniam"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "non",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(516231),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCustomerResponse != nil {
        // handle response
    }
}
```
