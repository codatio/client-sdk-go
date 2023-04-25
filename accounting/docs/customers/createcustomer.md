# CreateCustomer
Available in: `Customers`

Posts an individual customer for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update customer model](https://docs.codat.io/accounting-api#/operations/get-create-update-customers-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support creating customers.

## Example Usage
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
    req := operations.CreateCustomerRequest{
        Customer: &shared.Customer{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("Pasco"),
                    Country: codataccounting.String("French Polynesia"),
                    Line1: codataccounting.String("unde"),
                    Line2: codataccounting.String("consequatur"),
                    PostalCode: codataccounting.String("74731"),
                    Region: codataccounting.String("culpa"),
                    Type: shared.AddressTypeEnumDelivery,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Port Marcoschester"),
                    Country: codataccounting.String("Democratic People's Republic of Korea"),
                    Line1: codataccounting.String("totam"),
                    Line2: codataccounting.String("incidunt"),
                    PostalCode: codataccounting.String("76181-2098"),
                    Region: codataccounting.String("tenetur"),
                    Type: shared.AddressTypeEnumBilling,
                },
            },
            ContactName: codataccounting.String("dolor"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("North Margieland"),
                        Country: codataccounting.String("Botswana"),
                        Line1: codataccounting.String("veniam"),
                        Line2: codataccounting.String("non"),
                        PostalCode: codataccounting.String("17849-8333"),
                        Region: codataccounting.String("voluptatem"),
                        Type: shared.AddressTypeEnumBilling,
                    },
                    Email: codataccounting.String("Bridget.Bogan53@yahoo.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
                    Name: codataccounting.String("Clifford Stiedemann"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "voluptas",
                            Type: shared.PhoneNumberTypeEnumLandline,
                        },
                        shared.PhoneNumbersitems{
                            Number: "eum",
                            Type: shared.PhoneNumberTypeEnumLandline,
                        },
                        shared.PhoneNumbersitems{
                            Number: "cumque",
                            Type: shared.PhoneNumberTypeEnumPrimary,
                        },
                    },
                    Status: shared.CustomerStatusEnumUnknown,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Batzport"),
                        Country: codataccounting.String("Bouvet Island (Bouvetoya)"),
                        Line1: codataccounting.String("veritatis"),
                        Line2: codataccounting.String("similique"),
                        PostalCode: codataccounting.String("41812-6487"),
                        Region: codataccounting.String("libero"),
                        Type: shared.AddressTypeEnumBilling,
                    },
                    Email: codataccounting.String("Emil82@hotmail.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
                    Name: codataccounting.String("Freddie Conroy"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "atque",
                            Type: shared.PhoneNumberTypeEnumPrimary,
                        },
                        shared.PhoneNumbersitems{
                            Number: "dolores",
                            Type: shared.PhoneNumberTypeEnumUnknown,
                        },
                    },
                    Status: shared.CustomerStatusEnumArchived,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("North Carrie"),
                        Country: codataccounting.String("Samoa"),
                        Line1: codataccounting.String("deleniti"),
                        Line2: codataccounting.String("esse"),
                        PostalCode: codataccounting.String("12079-3259"),
                        Region: codataccounting.String("corrupti"),
                        Type: shared.AddressTypeEnumBilling,
                    },
                    Email: codataccounting.String("Joan68@hotmail.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
                    Name: codataccounting.String("James Reynolds"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "dicta",
                            Type: shared.PhoneNumberTypeEnumUnknown,
                        },
                        shared.PhoneNumbersitems{
                            Number: "dolores",
                            Type: shared.PhoneNumberTypeEnumPrimary,
                        },
                        shared.PhoneNumbersitems{
                            Number: "eum",
                            Type: shared.PhoneNumberTypeEnumUnknown,
                        },
                        shared.PhoneNumbersitems{
                            Number: "corporis",
                            Type: shared.PhoneNumberTypeEnumUnknown,
                        },
                    },
                    Status: shared.CustomerStatusEnumActive,
                },
            },
            CustomerName: codataccounting.String("sequi"),
            DefaultCurrency: codataccounting.String("illo"),
            EmailAddress: codataccounting.String("temporibus"),
            ID: codataccounting.String("0081090f-6706-4673-b3a6-81c5768dce74"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Phone: codataccounting.String("1-405-711-3805"),
            RegistrationNumber: codataccounting.String("voluptas"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Status: shared.CustomerStatusEnumUnknown,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "tempora": map[string]interface{}{
                        "cupiditate": "officia",
                        "minima": "doloribus",
                        "suscipit": "sequi",
                    },
                },
            },
            TaxNumber: codataccounting.String("debitis"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(207202),
    }

    res, err := s.Customers.CreateCustomer(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateCustomerResponse != nil {
        // handle response
    }
}
```