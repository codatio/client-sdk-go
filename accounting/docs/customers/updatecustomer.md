# UpdateCustomer
Available in: `Customers`

Posts an updated customer for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update customer model](https://docs.codat.io/accounting-api#/operations/get-create-update-customers-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support updating customers.

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
    req := operations.UpdateCustomerRequest{
        Customer: &shared.Customer{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("Southaven"),
                    Country: codataccounting.String("Nicaragua"),
                    Line1: codataccounting.String("dolorem"),
                    Line2: codataccounting.String("velit"),
                    PostalCode: codataccounting.String("88422-5292"),
                    Region: codataccounting.String("dolorum"),
                    Type: shared.AddressTypeEnumBilling,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Morissettestad"),
                    Country: codataccounting.String("Trinidad and Tobago"),
                    Line1: codataccounting.String("modi"),
                    Line2: codataccounting.String("assumenda"),
                    PostalCode: codataccounting.String("24833-7563"),
                    Region: codataccounting.String("nostrum"),
                    Type: shared.AddressTypeEnumDelivery,
                },
                shared.Addressesitems{
                    City: codataccounting.String("North Cristina"),
                    Country: codataccounting.String("Tuvalu"),
                    Line1: codataccounting.String("inventore"),
                    Line2: codataccounting.String("ipsum"),
                    PostalCode: codataccounting.String("25121-0504"),
                    Region: codataccounting.String("cum"),
                    Type: shared.AddressTypeEnumDelivery,
                },
            },
            ContactName: codataccounting.String("ratione"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Port Adriana"),
                        Country: codataccounting.String("Mongolia"),
                        Line1: codataccounting.String("sed"),
                        Line2: codataccounting.String("harum"),
                        PostalCode: codataccounting.String("42290-1225"),
                        Region: codataccounting.String("pariatur"),
                        Type: shared.AddressTypeEnumBilling,
                    },
                    Email: codataccounting.String("Virgie_Gleichner15@yahoo.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
                    Name: codataccounting.String("Tabitha Toy"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "veritatis",
                            Type: shared.PhoneNumberTypeEnumPrimary,
                        },
                        shared.PhoneNumbersitems{
                            Number: "quod",
                            Type: shared.PhoneNumberTypeEnumFax,
                        },
                        shared.PhoneNumbersitems{
                            Number: "sapiente",
                            Type: shared.PhoneNumberTypeEnumUnknown,
                        },
                    },
                    Status: shared.CustomerStatusEnumActive,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("New Raleigh"),
                        Country: codataccounting.String("Palau"),
                        Line1: codataccounting.String("voluptatibus"),
                        Line2: codataccounting.String("modi"),
                        PostalCode: codataccounting.String("61494"),
                        Region: codataccounting.String("error"),
                        Type: shared.AddressTypeEnumDelivery,
                    },
                    Email: codataccounting.String("Pierce.McCullough83@gmail.com"),
                    ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
                    Name: codataccounting.String("Michelle Trantow"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "neque",
                            Type: shared.PhoneNumberTypeEnumFax,
                        },
                        shared.PhoneNumbersitems{
                            Number: "repellendus",
                            Type: shared.PhoneNumberTypeEnumLandline,
                        },
                        shared.PhoneNumbersitems{
                            Number: "quisquam",
                            Type: shared.PhoneNumberTypeEnumMobile,
                        },
                    },
                    Status: shared.CustomerStatusEnumArchived,
                },
            },
            CustomerName: codataccounting.String("doloremque"),
            DefaultCurrency: codataccounting.String("adipisci"),
            EmailAddress: codataccounting.String("quasi"),
            ID: codataccounting.String("08d9c337-4730-482b-94f2-ab1fd5671e9c"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Phone: codataccounting.String("342.406.2441"),
            RegistrationNumber: codataccounting.String("aliquam"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Status: shared.CustomerStatusEnumUnknown,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "blanditiis": map[string]interface{}{
                        "quisquam": "itaque",
                        "consequatur": "recusandae",
                        "iste": "error",
                    },
                    "dicta": map[string]interface{}{
                        "unde": "numquam",
                        "temporibus": "omnis",
                    },
                },
            },
            TaxNumber: codataccounting.String("amet"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "deserunt",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(495515),
    }

    res, err := s.Customers.UpdateCustomer(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCustomerResponse != nil {
        // handle response
    }
}
```