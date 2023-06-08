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
                    City: codataccounting.String("Port Camylleburgh"),
                    Country: codataccounting.String("American Samoa"),
                    Line1: codataccounting.String("voluptates"),
                    Line2: codataccounting.String("nulla"),
                    PostalCode: codataccounting.String("42511-6113"),
                    Region: codataccounting.String("non"),
                    Type: shared.AddressTypeBilling,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Lake Ryannmouth"),
                    Country: codataccounting.String("Tunisia"),
                    Line1: codataccounting.String("repellendus"),
                    Line2: codataccounting.String("enim"),
                    PostalCode: codataccounting.String("30441"),
                    Region: codataccounting.String("veritatis"),
                    Type: shared.AddressTypeDelivery,
                },
                shared.Addressesitems{
                    City: codataccounting.String("New Winnifred"),
                    Country: codataccounting.String("French Guiana"),
                    Line1: codataccounting.String("nulla"),
                    Line2: codataccounting.String("iusto"),
                    PostalCode: codataccounting.String("53342"),
                    Region: codataccounting.String("cumque"),
                    Type: shared.AddressTypeUnknown,
                },
            },
            ContactName: codataccounting.String("accusantium"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("South Beverlyland"),
                        Country: codataccounting.String("Netherlands"),
                        Line1: codataccounting.String("omnis"),
                        Line2: codataccounting.String("commodi"),
                        PostalCode: codataccounting.String("81264"),
                        Region: codataccounting.String("nulla"),
                        Type: shared.AddressTypeDelivery,
                    },
                    Email: codataccounting.String("Karli24@hotmail.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Name: codataccounting.String("Dave Metz"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeMobile,
                        },
                        shared.PhoneNumbersitems{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeFax,
                        },
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Leschport"),
                        Country: codataccounting.String("Madagascar"),
                        Line1: codataccounting.String("aspernatur"),
                        Line2: codataccounting.String("eius"),
                        PostalCode: codataccounting.String("79325"),
                        Region: codataccounting.String("a"),
                        Type: shared.AddressTypeBilling,
                    },
                    Email: codataccounting.String("Tyrel.Kuvalis68@yahoo.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Name: codataccounting.String("Hannah Schmidt"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Ryleighbury"),
                        Country: codataccounting.String("Syrian Arab Republic"),
                        Line1: codataccounting.String("quas"),
                        Line2: codataccounting.String("sequi"),
                        PostalCode: codataccounting.String("80050"),
                        Region: codataccounting.String("voluptatem"),
                        Type: shared.AddressTypeBilling,
                    },
                    Email: codataccounting.String("Virgie.Hodkiewicz@gmail.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Name: codataccounting.String("Dora Flatley"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeMobile,
                        },
                        shared.PhoneNumbersitems{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypeFax,
                        },
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeMobile,
                        },
                    },
                    Status: shared.CustomerStatusActive,
                },
            },
            CustomerName: codataccounting.String("blanditiis"),
            DefaultCurrency: codataccounting.String("EUR"),
            EmailAddress: codataccounting.String("impedit"),
            ID: codataccounting.String("e742409a-215e-4086-8148-9a5f63e3af3d"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Phone: codataccounting.String("(886) 328-8842 x2529"),
            RegistrationNumber: codataccounting.String("non"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusActive,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "natus": map[string]interface{}{
                        "saepe": "modi",
                        "assumenda": "maiores",
                        "neque": "in",
                    },
                    "debitis": map[string]interface{}{
                        "nostrum": "libero",
                        "totam": "omnis",
                    },
                    "veniam": map[string]interface{}{
                        "facere": "aliquam",
                        "vitae": "ipsum",
                    },
                },
            },
            TaxNumber: codataccounting.String("recusandae"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(78486),
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
        CustomerID: "ipsum",
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
        CustomerID: "error",
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
        CustomerID: "numquam",
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
        Query: codataccounting.String("praesentium"),
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
        CustomerID: "dolores",
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
                    City: codataccounting.String("North Kristatown"),
                    Country: codataccounting.String("Japan"),
                    Line1: codataccounting.String("cum"),
                    Line2: codataccounting.String("facere"),
                    PostalCode: codataccounting.String("32705"),
                    Region: codataccounting.String("sed"),
                    Type: shared.AddressTypeDelivery,
                },
            },
            ContactName: codataccounting.String("vero"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("West Velma"),
                        Country: codataccounting.String("Albania"),
                        Line1: codataccounting.String("consequuntur"),
                        Line2: codataccounting.String("numquam"),
                        PostalCode: codataccounting.String("58549"),
                        Region: codataccounting.String("modi"),
                        Type: shared.AddressTypeDelivery,
                    },
                    Email: codataccounting.String("Carmine_Bechtelar34@hotmail.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Name: codataccounting.String("Peter Casper"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeMobile,
                        },
                        shared.PhoneNumbersitems{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypeMobile,
                        },
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeFax,
                        },
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Corwinport"),
                        Country: codataccounting.String("Venezuela"),
                        Line1: codataccounting.String("ex"),
                        Line2: codataccounting.String("error"),
                        PostalCode: codataccounting.String("17594-8008"),
                        Region: codataccounting.String("omnis"),
                        Type: shared.AddressTypeDelivery,
                    },
                    Email: codataccounting.String("Consuelo.OReilly78@hotmail.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Name: codataccounting.String("Mr. Brian Fisher"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeFax,
                        },
                        shared.PhoneNumbersitems{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
            },
            CustomerName: codataccounting.String("rem"),
            DefaultCurrency: codataccounting.String("GBP"),
            EmailAddress: codataccounting.String("libero"),
            ID: codataccounting.String("94f2ab1f-d567-41e9-8326-350a46714378"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Phone: codataccounting.String("909-661-3628 x62642"),
            RegistrationNumber: codataccounting.String("impedit"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusUnknown,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "sed": map[string]interface{}{
                        "saepe": "dolor",
                        "quidem": "quaerat",
                        "cum": "dolore",
                        "quibusdam": "rerum",
                    },
                    "atque": map[string]interface{}{
                        "odio": "reprehenderit",
                        "quas": "voluptates",
                        "distinctio": "nam",
                    },
                },
            },
            TaxNumber: codataccounting.String("nisi"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "officiis",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(95168),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCustomerResponse != nil {
        // handle response
    }
}
```
