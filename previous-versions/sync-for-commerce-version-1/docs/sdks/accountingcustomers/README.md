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
                    City: codatsynccommerce.String("Reingerstad"),
                    Country: codatsynccommerce.String("Palau"),
                    Line1: codatsynccommerce.String("odio"),
                    Line2: codatsynccommerce.String("fugit"),
                    PostalCode: codatsynccommerce.String("14002"),
                    Region: codatsynccommerce.String("neque"),
                    Type: shared.AccountingAddressTypeBilling,
                },
                shared.Items1{
                    City: codatsynccommerce.String("Blaine"),
                    Country: codatsynccommerce.String("Croatia"),
                    Line1: codatsynccommerce.String("unde"),
                    Line2: codatsynccommerce.String("nulla"),
                    PostalCode: codatsynccommerce.String("81136-7167"),
                    Region: codatsynccommerce.String("fugiat"),
                    Type: shared.AccountingAddressTypeBilling,
                },
            },
            ContactName: codatsynccommerce.String("quos"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Items1{
                        City: codatsynccommerce.String("South Anastacioshire"),
                        Country: codatsynccommerce.String("Belarus"),
                        Line1: codatsynccommerce.String("aperiam"),
                        Line2: codatsynccommerce.String("totam"),
                        PostalCode: codatsynccommerce.String("77044"),
                        Region: codatsynccommerce.String("dolores"),
                        Type: shared.AccountingAddressTypeDelivery,
                    },
                    Email: codatsynccommerce.String("Marcella.Schumm@gmail.com"),
                    ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                    Name: codatsynccommerce.String("Terence Reynolds"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeFax,
                        },
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                    },
                    Status: shared.CustomerStatusArchived,
                },
                shared.Contact{
                    Address: &shared.Items1{
                        City: codatsynccommerce.String("Reno"),
                        Country: codatsynccommerce.String("Taiwan"),
                        Line1: codatsynccommerce.String("alias"),
                        Line2: codatsynccommerce.String("quia"),
                        PostalCode: codatsynccommerce.String("69078-1845"),
                        Region: codatsynccommerce.String("odit"),
                        Type: shared.AccountingAddressTypeUnknown,
                    },
                    Email: codatsynccommerce.String("Lillie92@yahoo.com"),
                    ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                    Name: codatsynccommerce.String("Miss Alison Hayes"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                        shared.PhoneNumbersitems{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                        shared.PhoneNumbersitems{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypeUnknown,
                        },
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeMobile,
                        },
                    },
                    Status: shared.CustomerStatusArchived,
                },
                shared.Contact{
                    Address: &shared.Items1{
                        City: codatsynccommerce.String("Grand Prairie"),
                        Country: codatsynccommerce.String("Cocos (Keeling) Islands"),
                        Line1: codatsynccommerce.String("expedita"),
                        Line2: codatsynccommerce.String("voluptatum"),
                        PostalCode: codatsynccommerce.String("38324-7423"),
                        Region: codatsynccommerce.String("magnam"),
                        Type: shared.AccountingAddressTypeDelivery,
                    },
                    Email: codatsynccommerce.String("Mose.Bayer28@hotmail.com"),
                    ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                    Name: codatsynccommerce.String("Glen Satterfield"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeFax,
                        },
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeUnknown,
                        },
                        shared.PhoneNumbersitems{
                            Number: "01224 658 999",
                            Type: shared.PhoneNumberTypeUnknown,
                        },
                    },
                    Status: shared.CustomerStatusUnknown,
                },
                shared.Contact{
                    Address: &shared.Items1{
                        City: codatsynccommerce.String("Stillwater"),
                        Country: codatsynccommerce.String("Guinea"),
                        Line1: codatsynccommerce.String("reprehenderit"),
                        Line2: codatsynccommerce.String("aperiam"),
                        PostalCode: codatsynccommerce.String("34451"),
                        Region: codatsynccommerce.String("error"),
                        Type: shared.AccountingAddressTypeUnknown,
                    },
                    Email: codatsynccommerce.String("Jabari_Streich76@yahoo.com"),
                    ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
                    Name: codatsynccommerce.String("Tonya Towne"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "(877) 492-8687",
                            Type: shared.PhoneNumberTypeLandline,
                        },
                        shared.PhoneNumbersitems{
                            Number: "+44 25691 154789",
                            Type: shared.PhoneNumberTypePrimary,
                        },
                    },
                    Status: shared.CustomerStatusActive,
                },
            },
            CustomerName: codatsynccommerce.String("vero"),
            DefaultCurrency: codatsynccommerce.String("GBP"),
            EmailAddress: codatsynccommerce.String("repudiandae"),
            ID: codatsynccommerce.String("b1e5a2b1-2eb0-47f1-96db-99545fc95fa8"),
            Metadata: &shared.Metadata{
                IsDeleted: codatsynccommerce.Bool(false),
            },
            ModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Phone: codatsynccommerce.String("(509) 256-8772 x0977"),
            RegistrationNumber: codatsynccommerce.String("ipsum"),
            SourceModifiedDate: codatsynccommerce.String("2022-10-23T00:00:00.000Z"),
            Status: shared.CustomerStatusArchived,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "doloremque": map[string]interface{}{
                        "veniam": "libero",
                        "architecto": "cupiditate",
                    },
                    "molestiae": map[string]interface{}{
                        "possimus": "non",
                        "magnam": "itaque",
                        "sed": "asperiores",
                        "veniam": "consequuntur",
                    },
                    "facere": map[string]interface{}{
                        "odit": "pariatur",
                        "amet": "exercitationem",
                        "ab": "velit",
                    },
                },
            },
            TaxNumber: codatsynccommerce.String("facilis"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codatsynccommerce.Int(731065),
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

