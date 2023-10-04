# AccountingCustomers
(*AccountingCustomers*)

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
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/operations"
)

func main() {
    s := syncforcommerceversion1.New(
        syncforcommerceversion1.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountingCustomers.CreateAccountingCustomer(ctx, operations.CreateAccountingCustomerRequest{
        AccountingCustomer: &shared.AccountingCustomer{
            Addresses: []shared.Items1{
                shared.Items1{
                    City: syncforcommerceversion1.String("Binghamton"),
                    Country: syncforcommerceversion1.String("Saint Pierre and Miquelon"),
                    Line1: syncforcommerceversion1.String("Southeast"),
                    Line2: syncforcommerceversion1.String("Adventure"),
                    PostalCode: syncforcommerceversion1.String("03600"),
                    Region: syncforcommerceversion1.String("gray Bespoke Bahamian"),
                    Type: shared.AccountingAddressTypeUnknown,
                },
            },
            ContactName: syncforcommerceversion1.String("Southeast GB Product"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Items1{
                        City: syncforcommerceversion1.String("South Jackyshire"),
                        Country: syncforcommerceversion1.String("Comoros"),
                        Line1: syncforcommerceversion1.String("Coordinator"),
                        Line2: syncforcommerceversion1.String("Lead"),
                        PostalCode: syncforcommerceversion1.String("69139-7324"),
                        Region: syncforcommerceversion1.String("considering"),
                        Type: shared.AccountingAddressTypeDelivery,
                    },
                    Email: syncforcommerceversion1.String("Lonzo_Von@gmail.com"),
                    ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
                    Name: syncforcommerceversion1.String("North Guarani"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
            },
            CustomerName: syncforcommerceversion1.String("Wilmington Soft"),
            DefaultCurrency: syncforcommerceversion1.String("USD"),
            EmailAddress: syncforcommerceversion1.String("Ronny.Terry91@yahoo.com"),
            ID: syncforcommerceversion1.String("<ID>"),
            Metadata: &shared.Metadata{
                IsDeleted: syncforcommerceversion1.Bool(false),
            },
            ModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Phone: syncforcommerceversion1.String("(577) 959-8272"),
            RegistrationNumber: syncforcommerceversion1.String("Plastic"),
            SourceModifiedDate: syncforcommerceversion1.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusArchived,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "distinctio": map[string]interface{}{
                        "molestiae": "programming",
                    },
                },
            },
            TaxNumber: syncforcommerceversion1.String("Rubber Androgyne Bronze"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: syncforcommerceversion1.Int(359161),
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

