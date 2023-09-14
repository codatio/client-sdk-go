# Configuration

## Overview

Manage mapping options and sync configuration.

### Available Operations

* [Get](#get) - Get company configuration
* [GetMappingOptions](#getmappingoptions) - Mapping options
* [Set](#set) - Set company configuration

## Get

Gets a companies expense sync configuration

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-expenses"
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Configuration.Get(ctx, operations.GetCompanyConfigurationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CompanyConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetCompanyConfigurationRequest](../../models/operations/getcompanyconfigurationrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.GetCompanyConfigurationResponse](../../models/operations/getcompanyconfigurationresponse.md), error**


## GetMappingOptions

Gets the expense mapping options for a companies accounting software

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-expenses"
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Configuration.GetMappingOptions(ctx, operations.GetMappingOptionsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MappingOptions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetMappingOptionsRequest](../../models/operations/getmappingoptionsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.GetMappingOptionsResponse](../../models/operations/getmappingoptionsresponse.md), error**


## Set

Sets a companies expense sync configuration

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/sync-for-expenses"
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/models/operations"
)

func main() {
    s := codatsyncexpenses.New(
        codatsyncexpenses.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Configuration.Set(ctx, operations.SetCompanyConfigurationRequest{
        CompanyConfiguration: &shared.CompanyConfiguration{
            BankAccount: shared.BankAccount{
                ID: codatsyncexpenses.String("32"),
            },
            Customer: shared.Customer{
                Addresses: []shared.Items{
                    shared.Items{
                        City: codatsyncexpenses.String("Fort Donnybury"),
                        Country: codatsyncexpenses.String("Kyrgyz Republic"),
                        Line1: codatsyncexpenses.String("minus"),
                        Line2: codatsyncexpenses.String("placeat"),
                        PostalCode: codatsyncexpenses.String("45398-0306"),
                        Region: codatsyncexpenses.String("perferendis"),
                        Type: shared.AccountingAddressTypeBilling,
                    },
                },
                ContactName: codatsyncexpenses.String("repellendus"),
                Contacts: []shared.Contact{
                    shared.Contact{
                        Address: &shared.Items{
                            City: codatsyncexpenses.String("San Antonio"),
                            Country: codatsyncexpenses.String("Burundi"),
                            Line1: codatsyncexpenses.String("at"),
                            Line2: codatsyncexpenses.String("at"),
                            PostalCode: codatsyncexpenses.String("47845-7617"),
                            Region: codatsyncexpenses.String("officia"),
                            Type: shared.AccountingAddressTypeBilling,
                        },
                        Email: codatsyncexpenses.String("Kale_Welch10@gmail.com"),
                        ModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
                        Name: codatsyncexpenses.String("Pauline Dibbert"),
                        Phone: []shared.ContactPhone{
                            shared.ContactPhone{
                                Number: "(877) 492-8687",
                                Type: shared.PhoneNumberTypeLandline,
                            },
                        },
                        Status: shared.CustomerStatusActive,
                    },
                },
                CustomerName: codatsyncexpenses.String("aspernatur"),
                DefaultCurrency: codatsyncexpenses.String("GBP"),
                EmailAddress: codatsyncexpenses.String("ad"),
                ID: codatsyncexpenses.String("929396fe-a759-46eb-90fa-aa2352c59559"),
                Metadata: &shared.Metadata{
                    IsDeleted: codatsyncexpenses.Bool(false),
                },
                ModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
                Phone: codatsyncexpenses.String("799.262.6196 x524"),
                RegistrationNumber: codatsyncexpenses.String("quam"),
                SourceModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
                Status: shared.CustomerStatusUnknown,
                SupplementalData: &shared.SupplementalData{
                    Content: map[string]map[string]interface{}{
                        "error": map[string]interface{}{
                            "quia": "quis",
                        },
                    },
                },
                TaxNumber: codatsyncexpenses.String("vitae"),
            },
            Supplier: shared.Supplier{
                Addresses: []shared.SupplierAccountingAddress{
                    shared.SupplierAccountingAddress{
                        City: codatsyncexpenses.String("O'Konborough"),
                        Country: codatsyncexpenses.String("Burkina Faso"),
                        Line1: codatsyncexpenses.String("quo"),
                        Line2: codatsyncexpenses.String("sequi"),
                        PostalCode: codatsyncexpenses.String("36800-6860"),
                        Region: codatsyncexpenses.String("reiciendis"),
                        Type: shared.AccountingAddressTypeDelivery,
                    },
                },
                ContactName: codatsyncexpenses.String("vero"),
                DefaultCurrency: codatsyncexpenses.String("nihil"),
                EmailAddress: codatsyncexpenses.String("praesentium"),
                ID: codatsyncexpenses.String("f097b007-4f15-4471-b5e6-e13b99d488e1"),
                Metadata: &shared.Metadata{
                    IsDeleted: codatsyncexpenses.Bool(false),
                },
                ModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
                Phone: codatsyncexpenses.String("(877) 492-8687"),
                RegistrationNumber: codatsyncexpenses.String("veritatis"),
                SourceModifiedDate: codatsyncexpenses.String("2022-10-23T00:00:00.000Z"),
                Status: shared.SupplierStatusUnknown,
                SupplementalData: &shared.SupplierSupplementalData{
                    Content: map[string]map[string]interface{}{
                        "enim": map[string]interface{}{
                            "consequatur": "est",
                        },
                    },
                },
                SupplierName: codatsyncexpenses.String("quibusdam"),
                TaxNumber: codatsyncexpenses.String("explicabo"),
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CompanyConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.SetCompanyConfigurationRequest](../../models/operations/setcompanyconfigurationrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.SetCompanyConfigurationResponse](../../models/operations/setcompanyconfigurationresponse.md), error**

