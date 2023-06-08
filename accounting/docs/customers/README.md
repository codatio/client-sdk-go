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
                    City: codataccounting.String("Diamond Bar"),
                    Country: codataccounting.String("Switzerland"),
                    Line1: codataccounting.String("vero"),
                    Line2: codataccounting.String("natus"),
                    PostalCode: codataccounting.String("86228-8842"),
                    Region: codataccounting.String("incidunt"),
                    Type: shared.AddressTypeBilling,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Fort Dena"),
                    Country: codataccounting.String("Palau"),
                    Line1: codataccounting.String("esse"),
                    Line2: codataccounting.String("id"),
                    PostalCode: codataccounting.String("59289-2483"),
                    Region: codataccounting.String("nostrum"),
                    Type: shared.AddressTypeDelivery,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Lenoreborough"),
                    Country: codataccounting.String("Grenada"),
                    Line1: codataccounting.String("facere"),
                    Line2: codataccounting.String("aliquam"),
                    PostalCode: codataccounting.String("29026"),
                    Region: codataccounting.String("numquam"),
                    Type: shared.AddressTypeBilling,
                },
            },
            ContactName: codataccounting.String("dolores"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("North Kristatown"),
                        Country: codataccounting.String("Japan"),
                        Line1: codataccounting.String("cum"),
                        Line2: codataccounting.String("facere"),
                        PostalCode: codataccounting.String("32705"),
                        Region: codataccounting.String("sed"),
                        Type: shared.AddressTypeDelivery,
                    },
                    Email: codataccounting.String("Janie95@hotmail.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Name: codataccounting.String("Marjorie Fritsch"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeUnknown,
                        },
                        shared.PhoneNumbersitems{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypeFax,
                        },
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
            },
            CustomerName: codataccounting.String("asperiores"),
            DefaultCurrency: codataccounting.String("EUR"),
            EmailAddress: codataccounting.String("nostrum"),
            ID: codataccounting.String("d911cbfe-749c-4af4-9a27-f69e2c9e6d10"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Phone: codataccounting.String("(862) 782-7470 x2005"),
            RegistrationNumber: codataccounting.String("facere"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusArchived,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "sequi": map[string]interface{}{
                        "modi": "nihil",
                        "velit": "aut",
                    },
                },
            },
            TaxNumber: codataccounting.String("rem"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(142520),
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
        CustomerID: "incidunt",
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
        Query: codataccounting.String("asperiores"),
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
        CustomerID: "eos",
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
                    City: codataccounting.String("Cartershire"),
                    Country: codataccounting.String("Somalia"),
                    Line1: codataccounting.String("nemo"),
                    Line2: codataccounting.String("eum"),
                    PostalCode: codataccounting.String("09572"),
                    Region: codataccounting.String("dolores"),
                    Type: shared.AddressTypeBilling,
                },
                shared.Addressesitems{
                    City: codataccounting.String("West Amira"),
                    Country: codataccounting.String("Nicaragua"),
                    Line1: codataccounting.String("tempora"),
                    Line2: codataccounting.String("eum"),
                    PostalCode: codataccounting.String("13245"),
                    Region: codataccounting.String("natus"),
                    Type: shared.AddressTypeDelivery,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Allen"),
                    Country: codataccounting.String("Tuvalu"),
                    Line1: codataccounting.String("iste"),
                    Line2: codataccounting.String("error"),
                    PostalCode: codataccounting.String("36286"),
                    Region: codataccounting.String("amet"),
                    Type: shared.AddressTypeBilling,
                },
            },
            ContactName: codataccounting.String("ducimus"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Beerberg"),
                        Country: codataccounting.String("Guatemala"),
                        Line1: codataccounting.String("sed"),
                        Line2: codataccounting.String("a"),
                        PostalCode: codataccounting.String("26372-8757"),
                        Region: codataccounting.String("odio"),
                        Type: shared.AddressTypeBilling,
                    },
                    Email: codataccounting.String("Taryn.Raynor88@yahoo.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Name: codataccounting.String("Brandon Rutherford"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypeFax,
                        },
                    },
                    Status: shared.CustomerStatusArchived,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Plymouth"),
                        Country: codataccounting.String("Brunei Darussalam"),
                        Line1: codataccounting.String("eligendi"),
                        Line2: codataccounting.String("nam"),
                        PostalCode: codataccounting.String("33238-3943"),
                        Region: codataccounting.String("at"),
                        Type: shared.AddressTypeBilling,
                    },
                    Email: codataccounting.String("Charity.Lubowitz88@gmail.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Name: codataccounting.String("Dr. Marie O'Hara"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeUnknown,
                        },
                    },
                    Status: shared.CustomerStatusActive,
                },
            },
            CustomerName: codataccounting.String("sint"),
            DefaultCurrency: codataccounting.String("USD"),
            EmailAddress: codataccounting.String("commodi"),
            ID: codataccounting.String("489d7b78-673e-413a-92a6-b99249459448"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Phone: codataccounting.String("475.425.1475 x36278"),
            RegistrationNumber: codataccounting.String("a"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusUnknown,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "minima": map[string]interface{}{
                        "sit": "incidunt",
                        "magnam": "molestias",
                        "hic": "error",
                    },
                },
            },
            TaxNumber: codataccounting.String("repellendus"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "delectus",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(118986),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCustomerResponse != nil {
        // handle response
    }
}
```
