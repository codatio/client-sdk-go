# Customers
(*Customers*)

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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.Create(ctx, operations.CreateCustomerRequest{
        Customer: &shared.Customer{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: accounting.String("Jenafurt"),
                    Country: accounting.String("Sweden"),
                    Line1: accounting.String("innovative blue"),
                    Line2: accounting.String("grey technology East"),
                    PostalCode: accounting.String("30778"),
                    Region: accounting.String("quantify Polestar mobile"),
                    Type: shared.AccountingAddressTypeBilling,
                },
            },
            ContactName: accounting.String("Durham after"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: accounting.String("Darenberg"),
                        Country: accounting.String("Cote d'Ivoire"),
                        Line1: accounting.String("Buckinghamshire functionalities Grocery"),
                        Line2: accounting.String("Metal"),
                        PostalCode: accounting.String("61380"),
                        Region: accounting.String("Interactions Senior Mouse"),
                        Type: shared.AccountingAddressTypeUnknown,
                    },
                    Email: accounting.String("Judd27@hotmail.com"),
                    ModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
                    Name: accounting.String("transmit likewise"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeFax,
                        },
                    },
                    Status: shared.CustomerStatusActive,
                },
            },
            CustomerName: accounting.String("Rubber silver Indiana"),
            DefaultCurrency: accounting.String("EUR"),
            EmailAddress: accounting.String("Thea_Ritchie76@hotmail.com"),
            ID: accounting.String("<ID>"),
            Metadata: &shared.Metadata{
                IsDeleted: accounting.Bool(false),
            },
            ModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Phone: accounting.String("948.595.2034"),
            RegistrationNumber: accounting.String("digital"),
            SourceModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusUnknown,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "recusandae": map[string]interface{}{
                        "maiores": "Mongolia",
                    },
                },
            },
            TaxNumber: accounting.String("discrete"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: accounting.Int(522311),
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.DownloadAttachment(ctx, operations.DownloadCustomerAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "Dakota Avon specifically",
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.Get(ctx, operations.GetCustomerRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CustomerID: "Northeast Hatchback Kia",
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.GetAttachment(ctx, operations.GetCustomerAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "array East along",
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.List(ctx, operations.ListCustomersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: accounting.String("-modifiedDate"),
        Page: accounting.Int(1),
        PageSize: accounting.Int(100),
        Query: accounting.String("Northeast Metal Canada"),
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.ListAttachments(ctx, operations.ListCustomerAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "intuitive Frozen ouch",
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
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
)

func main() {
    s := accounting.New(
        accounting.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Customers.Update(ctx, operations.UpdateCustomerRequest{
        Customer: &shared.Customer{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: accounting.String("Ann Arbor"),
                    Country: accounting.String("Montserrat"),
                    Line1: accounting.String("Reactive"),
                    Line2: accounting.String("Metal cheater Islands"),
                    PostalCode: accounting.String("43372"),
                    Region: accounting.String("Carolina syndicate"),
                    Type: shared.AccountingAddressTypeBilling,
                },
            },
            ContactName: accounting.String("East"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: accounting.String("Olenfurt"),
                        Country: accounting.String("Paraguay"),
                        Line1: accounting.String("Home users Sharable"),
                        Line2: accounting.String("Lev Wooden"),
                        PostalCode: accounting.String("36848"),
                        Region: accounting.String("brightly"),
                        Type: shared.AccountingAddressTypeDelivery,
                    },
                    Email: accounting.String("Josie49@yahoo.com"),
                    ModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
                    Name: accounting.String("possimus navigating Diesel"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
            },
            CustomerName: accounting.String("Reactive Global Northeast"),
            DefaultCurrency: accounting.String("USD"),
            EmailAddress: accounting.String("Abe.Bogan@hotmail.com"),
            ID: accounting.String("<ID>"),
            Metadata: &shared.Metadata{
                IsDeleted: accounting.Bool(false),
            },
            ModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Phone: accounting.String("(883) 732-4217 x6499"),
            RegistrationNumber: accounting.String("redundant ew"),
            SourceModifiedDate: accounting.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusActive,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "asperiores": map[string]interface{}{
                        "quibusdam": "Omnigender",
                    },
                },
            },
            TaxNumber: accounting.String("Volkswagen Specialist Bacon"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "Copper port East",
        ForceUpdate: accounting.Bool(false),
        TimeoutInMinutes: accounting.Int(373959),
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

