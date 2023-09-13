# AccountingCustomers

## Overview

Customers

### Available Operations

* [CreateAccountingCustomer](#createaccountingcustomer) - Create customer

## CreateAccountingCustomer

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
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
)

func main() {
    s := codatsynccommerce.New(
        codatsynccommerce.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountingCustomers.CreateAccountingCustomer(ctx, operations.CreateAccountingCustomerRequest{
        AccountingCustomer: &shared.AccountingCustomer{
            Addresses: []shared.Items1{
                shared.Items1{
                    City: codatsynccommerce.String("East Kylie"),
                    Country: codatsynccommerce.String("Slovakia (Slovak Republic)"),
                    Line1: codatsynccommerce.String("pariatur"),
                    Line2: codatsynccommerce.String("soluta"),
                    PostalCode: codatsynccommerce.String("65211"),
                    Region: codatsynccommerce.String("distinctio"),
                    Type: shared.AccountingAddressTypeDelivery,
                },
            },
            ContactName: codatsynccommerce.String("aliquid"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Items1{
                        City: codatsynccommerce.String("Kennedyhaven"),
                        Country: codatsynccommerce.String("Christmas Island"),
                        Line1: codatsynccommerce.String("neque"),
                        Line2: codatsynccommerce.String("fugit"),
                        PostalCode: codatsynccommerce.String("41379"),
                        Region: codatsynccommerce.String("voluptatem"),
                        Type: shared.AccountingAddressTypeDelivery,
                    },
                    Email: codatsynccommerce.String("Nella.Bosco8@hotmail.com"),
                    ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                    Name: codatsynccommerce.String("Dr. Randolph McDermott"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
            },
            CustomerName: codatsynccommerce.String("dolorum"),
            DefaultCurrency: codatsynccommerce.String("GBP"),
            EmailAddress: codatsynccommerce.String("quae"),
            ID: codatsynccommerce.String("08e0adcf-4b92-4187-9fce-953f73ef7fbc"),
            Metadata: &shared.Metadata{
                IsDeleted: codatsynccommerce.Bool(false),
            },
            ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Phone: codatsynccommerce.String("1-784-488-1670 x9381"),
            RegistrationNumber: codatsynccommerce.String("porro"),
            SourceModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusArchived,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "iusto": map[string]interface{}{
                        "eligendi": "ducimus",
                    },
                },
            },
            TaxNumber: codatsynccommerce.String("alias"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsynccommerce.Int(639473),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountingCreateCustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateAccountingCustomerRequest](../../models/operations/createaccountingcustomerrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.CreateAccountingCustomerResponse](../../models/operations/createaccountingcustomerresponse.md), error**

