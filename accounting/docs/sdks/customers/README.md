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

The *Create customer* endpoint creates a new [customer](https://docs.codat.io/accounting-api#/schemas/Customer) for a given company's connection.

[Customers](https://docs.codat.io/accounting-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update customer model](https://docs.codat.io/accounting-api#/operations/get-create-update-customers-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
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
                    City: codataccounting.String("San Antonio"),
                    Country: codataccounting.String("New Zealand"),
                    Line1: codataccounting.String("architecto"),
                    Line2: codataccounting.String("iure"),
                    PostalCode: codataccounting.String("95623-0684"),
                    Region: codataccounting.String("nisi"),
                    Type: shared.AddressTypeDelivery,
                },
            },
            ContactName: codataccounting.String("officiis"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Hilllton"),
                        Country: codataccounting.String("Mauritius"),
                        Line1: codataccounting.String("impedit"),
                        Line2: codataccounting.String("modi"),
                        PostalCode: codataccounting.String("61968-2330"),
                        Region: codataccounting.String("odit"),
                        Type: shared.AddressTypeDelivery,
                    },
                    Email: codataccounting.String("Gina93@gmail.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Name: codataccounting.String("Lance Cruickshank DVM"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypeMobile,
                        },
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeMobile,
                        },
                    },
                    Status: shared.CustomerStatusActive,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Adamsboro"),
                        Country: codataccounting.String("Chile"),
                        Line1: codataccounting.String("labore"),
                        Line2: codataccounting.String("earum"),
                        PostalCode: codataccounting.String("06184"),
                        Region: codataccounting.String("voluptates"),
                        Type: shared.AddressTypeDelivery,
                    },
                    Email: codataccounting.String("Domenick41@yahoo.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Name: codataccounting.String("Cristina Beer V"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeFax,
                        },
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeFax,
                        },
                    },
                    Status: shared.CustomerStatusActive,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Runolfssonside"),
                        Country: codataccounting.String("Uruguay"),
                        Line1: codataccounting.String("tenetur"),
                        Line2: codataccounting.String("ipsam"),
                        PostalCode: codataccounting.String("70098-3083"),
                        Region: codataccounting.String("aspernatur"),
                        Type: shared.AddressTypeBilling,
                    },
                    Email: codataccounting.String("Diana_Hammes@yahoo.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Name: codataccounting.String("Verna Lang"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeUnknown,
                        },
                        shared.PhoneNumbersitems{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypeUnknown,
                        },
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeFax,
                        },
                    },
                    Status: shared.CustomerStatusArchived,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Reychester"),
                        Country: codataccounting.String("French Guiana"),
                        Line1: codataccounting.String("veritatis"),
                        Line2: codataccounting.String("fugit"),
                        PostalCode: codataccounting.String("88364"),
                        Region: codataccounting.String("hic"),
                        Type: shared.AddressTypeDelivery,
                    },
                    Email: codataccounting.String("Jaden62@gmail.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Name: codataccounting.String("Verna Sauer"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeUnknown,
                        },
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeFax,
                        },
                    },
                    Status: shared.CustomerStatusActive,
                },
            },
            CustomerName: codataccounting.String("incidunt"),
            DefaultCurrency: codataccounting.String("EUR"),
            EmailAddress: codataccounting.String("nihil"),
            ID: codataccounting.String("7b4848bd-6a6f-4044-9d2c-3b808094373e"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Phone: codataccounting.String("1-223-679-7768 x019"),
            RegistrationNumber: codataccounting.String("sunt"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusActive,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "quidem": map[string]interface{}{
                        "a": "et",
                        "ipsam": "eos",
                        "exercitationem": "minima",
                        "laudantium": "quibusdam",
                    },
                    "fuga": map[string]interface{}{
                        "excepturi": "corporis",
                        "nam": "itaque",
                        "suscipit": "porro",
                    },
                },
            },
            TaxNumber: codataccounting.String("nulla"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(10447),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateCustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateCustomerRequest](../../models/operations/createcustomerrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.CreateCustomerResponse](../../models/operations/createcustomerresponse.md), error**


## DownloadAttachment

The *Download customer attachment* endpoint downloads a specific attachment for a given `customerId` and `attachmentId`.

[Customers](https://docs.codat.io/accounting-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support downloading a customer attachment.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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
        CustomerID: "qui",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.DownloadCustomerAttachmentRequest](../../models/operations/downloadcustomerattachmentrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |


### Response

**[*operations.DownloadCustomerAttachmentResponse](../../models/operations/downloadcustomerattachmentresponse.md), error**


## Get

The *Get customer* endpoint returns a single customer for a given customerId.

[Customers](https://docs.codat.io/accounting-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support getting a specific customer.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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
        CustomerID: "in",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Customer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetCustomerRequest](../../models/operations/getcustomerrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |


### Response

**[*operations.GetCustomerResponse](../../models/operations/getcustomerresponse.md), error**


## GetAttachment

The *Get customer attachment* endpoint returns a specific attachment for a given `customerId` and `attachmentId`.

[Customers](https://docs.codat.io/accounting-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support getting a customer attachment.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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
        CustomerID: "enim",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Attachment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetCustomerAttachmentRequest](../../models/operations/getcustomerattachmentrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |


### Response

**[*operations.GetCustomerAttachmentResponse](../../models/operations/getcustomerattachmentresponse.md), error**


## GetCreateUpdateModel

﻿The *Get create/update customer model* endpoint returns the expected data for the request payload when creating and updating a [customer](https://docs.codat.io/accounting-api#/schemas/Customer) for a given company and integration.

[Customers](https://docs.codat.io/accounting-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

**Integration-specific behaviour**

See the *response examples* for integration-specific indicative models.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support creating and updating a customer.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetCreateUpdateCustomersModelRequest](../../models/operations/getcreateupdatecustomersmodelrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |


### Response

**[*operations.GetCreateUpdateCustomersModelResponse](../../models/operations/getcreateupdatecustomersmodelresponse.md), error**


## List

The *List customers* endpoint returns a list of [customers](https://docs.codat.io/accounting-api#/schemas/Customer) for a given company's connection.

[Customers](https://docs.codat.io/accounting-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

Before using this endpoint, you must have [retrieved data for the company](https://docs.codat.io/codat-api#/operations/refresh-company-data).
    

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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
        Query: codataccounting.String("vel"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Customers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListCustomersRequest](../../models/operations/listcustomersrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.ListCustomersResponse](../../models/operations/listcustomersresponse.md), error**


## ListAttachments

The *List customer attachments* endpoint returns a list of attachments avialable to download for given `customerId`.

[Customers](https://docs.codat.io/accounting-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support listing customer attachments.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
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
        CustomerID: "impedit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachmentsDataset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListCustomerAttachmentsRequest](../../models/operations/listcustomerattachmentsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.ListCustomerAttachmentsResponse](../../models/operations/listcustomerattachmentsresponse.md), error**


## Update

The *Update customer* endpoint updates an existing [customer](https://docs.codat.io/accounting-api#/schemas/Customer) for a given company's connection.

[Customers](https://docs.codat.io/accounting-api#/schemas/Customer) are people or organizations that buy goods or services from the SMB.

**Integration-specific behaviour**

Required data may vary by integration. To see what data to post, first call [Get create/update customer model](https://docs.codat.io/accounting-api#/operations/get-create-update-customers-model).

Check out our [coverage explorer](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support creating an account.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/accounting"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/accounting/pkg/models/operations"
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
                    City: codataccounting.String("West Marcelina"),
                    Country: codataccounting.String("Palestinian Territory"),
                    Line1: codataccounting.String("labore"),
                    Line2: codataccounting.String("adipisci"),
                    PostalCode: codataccounting.String("73480"),
                    Region: codataccounting.String("esse"),
                    Type: shared.AddressTypeBilling,
                },
            },
            ContactName: codataccounting.String("amet"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("East Abdul"),
                        Country: codataccounting.String("Libyan Arab Jamahiriya"),
                        Line1: codataccounting.String("eligendi"),
                        Line2: codataccounting.String("qui"),
                        PostalCode: codataccounting.String("95501"),
                        Region: codataccounting.String("repellendus"),
                        Type: shared.AddressTypeBilling,
                    },
                    Email: codataccounting.String("Verona.Abshire34@gmail.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Name: codataccounting.String("Ray Pfannerstill"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeUnknown,
                        },
                    },
                    Status: shared.CustomerStatusActive,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Windlerfort"),
                        Country: codataccounting.String("Brazil"),
                        Line1: codataccounting.String("rem"),
                        Line2: codataccounting.String("itaque"),
                        PostalCode: codataccounting.String("74945-4921"),
                        Region: codataccounting.String("recusandae"),
                        Type: shared.AddressTypeUnknown,
                    },
                    Email: codataccounting.String("May.Donnelly89@gmail.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Name: codataccounting.String("Myra Hermann"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeUnknown,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Lake Roxaneburgh"),
                        Country: codataccounting.String("Sao Tome and Principe"),
                        Line1: codataccounting.String("iste"),
                        Line2: codataccounting.String("earum"),
                        PostalCode: codataccounting.String("77302"),
                        Region: codataccounting.String("repellat"),
                        Type: shared.AddressTypeBilling,
                    },
                    Email: codataccounting.String("Columbus_Moore82@hotmail.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Name: codataccounting.String("Duane Dibbert DVM"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeMobile,
                        },
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Lexifurt"),
                        Country: codataccounting.String("Somalia"),
                        Line1: codataccounting.String("molestiae"),
                        Line2: codataccounting.String("officiis"),
                        PostalCode: codataccounting.String("50466-9029"),
                        Region: codataccounting.String("dolorem"),
                        Type: shared.AddressTypeUnknown,
                    },
                    Email: codataccounting.String("Lizeth_Kuhlman@hotmail.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
                    Name: codataccounting.String("Nicolas Walker I"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeFax,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
            },
            CustomerName: codataccounting.String("commodi"),
            DefaultCurrency: codataccounting.String("EUR"),
            EmailAddress: codataccounting.String("temporibus"),
            ID: codataccounting.String("368ba921-6bcb-4415-835c-73641723133e"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Phone: codataccounting.String("224-783-1427 x78625"),
            RegistrationNumber: codataccounting.String("dolores"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusActive,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "non": map[string]interface{}{
                        "maxime": "aspernatur",
                    },
                    "magni": map[string]interface{}{
                        "minima": "ipsam",
                        "sequi": "quaerat",
                        "accusantium": "incidunt",
                        "cupiditate": "minima",
                    },
                    "quo": map[string]interface{}{
                        "facere": "quidem",
                        "harum": "adipisci",
                    },
                    "optio": map[string]interface{}{
                        "reprehenderit": "quo",
                        "vitae": "voluptates",
                    },
                },
            },
            TaxNumber: codataccounting.String("tempora"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "iste",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(560736),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateCustomerRequest](../../models/operations/updatecustomerrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.UpdateCustomerResponse](../../models/operations/updatecustomerresponse.md), error**

