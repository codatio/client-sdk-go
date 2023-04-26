# Customers

## Overview

Customers

### Available Operations

* [CreateCustomer](#createcustomer) - Create customer
* [DownloadCustomerAttachment](#downloadcustomerattachment) - Download customer attachment
* [GetCreateUpdateCustomersModel](#getcreateupdatecustomersmodel) - Get create/update customer model
* [GetCustomer](#getcustomer) - Get customer
* [GetCustomerAttachment](#getcustomerattachment) - Get customer attachment
* [GetCustomerAttachments](#getcustomerattachments) - List customer attachments
* [GetCustomers](#getcustomers) - List customers
* [UpdateCustomer](#updatecustomer) - Update customer

## CreateCustomer

Posts an individual customer for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update customer model](https://docs.codat.io/accounting-api#/operations/get-create-update-customers-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support creating customers.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CreateCustomerRequest{
        Customer: &shared.Customer{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("Jerrodfield"),
                    Country: codataccounting.String("Morocco"),
                    Line1: codataccounting.String("maxime"),
                    Line2: codataccounting.String("culpa"),
                    PostalCode: codataccounting.String("90218-8486"),
                    Region: codataccounting.String("est"),
                    Type: shared.AddressTypeEnumDelivery,
                },
            },
            ContactName: codataccounting.String("occaecati"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Port Malachi"),
                        Country: codataccounting.String("New Zealand"),
                        Line1: codataccounting.String("beatae"),
                        Line2: codataccounting.String("quod"),
                        PostalCode: codataccounting.String("95903-8950"),
                        Region: codataccounting.String("amet"),
                        Type: shared.AddressTypeEnumBilling,
                    },
                    Email: codataccounting.String("Jerome49@gmail.com"),
                    ModifiedDate: codataccounting.String("deleniti"),
                    Name: codataccounting.String("Anna Mayer"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "eius",
                            Type: shared.PhoneNumberTypeEnumPrimary,
                        },
                        shared.PhoneNumbersitems{
                            Number: "recusandae",
                            Type: shared.PhoneNumberTypeEnumLandline,
                        },
                        shared.PhoneNumbersitems{
                            Number: "aliquam",
                            Type: shared.PhoneNumberTypeEnumFax,
                        },
                        shared.PhoneNumbersitems{
                            Number: "voluptatum",
                            Type: shared.PhoneNumberTypeEnumLandline,
                        },
                    },
                    Status: shared.CustomerStatusEnumUnknown,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Hayleeborough"),
                        Country: codataccounting.String("Montenegro"),
                        Line1: codataccounting.String("iusto"),
                        Line2: codataccounting.String("optio"),
                        PostalCode: codataccounting.String("01218"),
                        Region: codataccounting.String("inventore"),
                        Type: shared.AddressTypeEnumUnknown,
                    },
                    Email: codataccounting.String("Bertha9@yahoo.com"),
                    ModifiedDate: codataccounting.String("dolorum"),
                    Name: codataccounting.String("Shaun Johnston"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "nemo",
                            Type: shared.PhoneNumberTypeEnumFax,
                        },
                        shared.PhoneNumbersitems{
                            Number: "quidem",
                            Type: shared.PhoneNumberTypeEnumUnknown,
                        },
                        shared.PhoneNumbersitems{
                            Number: "aliquid",
                            Type: shared.PhoneNumberTypeEnumLandline,
                        },
                        shared.PhoneNumbersitems{
                            Number: "atque",
                            Type: shared.PhoneNumberTypeEnumMobile,
                        },
                    },
                    Status: shared.CustomerStatusEnumArchived,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Fort Odie"),
                        Country: codataccounting.String("Bulgaria"),
                        Line1: codataccounting.String("nobis"),
                        Line2: codataccounting.String("voluptatum"),
                        PostalCode: codataccounting.String("16007"),
                        Region: codataccounting.String("saepe"),
                        Type: shared.AddressTypeEnumDelivery,
                    },
                    Email: codataccounting.String("Laurie4@yahoo.com"),
                    ModifiedDate: codataccounting.String("quae"),
                    Name: codataccounting.String("Veronica Kutch"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "libero",
                            Type: shared.PhoneNumberTypeEnumUnknown,
                        },
                    },
                    Status: shared.CustomerStatusEnumArchived,
                },
            },
            CustomerName: codataccounting.String("nihil"),
            DefaultCurrency: codataccounting.String("similique"),
            EmailAddress: codataccounting.String("repellat"),
            ID: codataccounting.String("ded84a35-a412-438e-9a73-5ac26ae33bef"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("provident"),
            Phone: codataccounting.String("265.924.7861 x004"),
            RegistrationNumber: codataccounting.String("delectus"),
            SourceModifiedDate: codataccounting.String("officiis"),
            Status: shared.CustomerStatusEnumActive,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "ipsam": map[string]interface{}{
                        "esse": "vitae",
                        "beatae": "pariatur",
                        "voluptatem": "blanditiis",
                    },
                    "eligendi": map[string]interface{}{
                        "deleniti": "deleniti",
                        "necessitatibus": "cumque",
                        "iste": "reiciendis",
                        "nihil": "libero",
                    },
                },
            },
            TaxNumber: codataccounting.String("perspiciatis"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(577273),
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

## DownloadCustomerAttachment

Download customer attachment

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.DownloadCustomerAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "officia",
    }

    res, err := s.Customers.DownloadCustomerAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```

## GetCreateUpdateCustomersModel

Get create/update customer model. Returns the expected data for the request payload.

See the examples for integration-specific indicative models.

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support creating and updating customers.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetCreateUpdateCustomersModelRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Customers.GetCreateUpdateCustomersModel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOption != nil {
        // handle response
    }
}
```

## GetCustomer

Gets a single customer corresponding to the given ID.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetCustomerRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CustomerID: "nemo",
    }

    res, err := s.Customers.GetCustomer(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Customer != nil {
        // handle response
    }
}
```

## GetCustomerAttachment

Get  customer attachment

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetCustomerAttachmentRequest{
        AttachmentID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "quis",
    }

    res, err := s.Customers.GetCustomerAttachment(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Attachment != nil {
        // handle response
    }
}
```

## GetCustomerAttachments

Get customer attachments

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetCustomerAttachmentsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "doloremque",
    }

    res, err := s.Customers.GetCustomerAttachments(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachmentsDataset != nil {
        // handle response
    }
}
```

## GetCustomers

Gets the latest customers for a company, with pagination

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetCustomersRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        OrderBy: codataccounting.String("-modifiedDate"),
        Page: 1,
        PageSize: codataccounting.Int(100),
        Query: codataccounting.String("similique"),
    }

    res, err := s.Customers.GetCustomers(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Customers != nil {
        // handle response
    }
}
```

## UpdateCustomer

Posts an updated customer for a given company.

Required data may vary by integration. To see what data to post, first call [Get create/update customer model](https://docs.codat.io/accounting-api#/operations/get-create-update-customers-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers) for integrations that support updating customers.

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
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.UpdateCustomerRequest{
        Customer: &shared.Customer{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("South Sydni"),
                    Country: codataccounting.String("Somalia"),
                    Line1: codataccounting.String("dolor"),
                    Line2: codataccounting.String("ratione"),
                    PostalCode: codataccounting.String("77078"),
                    Region: codataccounting.String("laudantium"),
                    Type: shared.AddressTypeEnumDelivery,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Jonesborough"),
                    Country: codataccounting.String("Eritrea"),
                    Line1: codataccounting.String("consectetur"),
                    Line2: codataccounting.String("qui"),
                    PostalCode: codataccounting.String("55487-4912"),
                    Region: codataccounting.String("quisquam"),
                    Type: shared.AddressTypeEnumBilling,
                },
            },
            ContactName: codataccounting.String("ipsam"),
            Contacts: []shared.Contact{
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("North Kamron"),
                        Country: codataccounting.String("Nauru"),
                        Line1: codataccounting.String("beatae"),
                        Line2: codataccounting.String("nemo"),
                        PostalCode: codataccounting.String("06400"),
                        Region: codataccounting.String("unde"),
                        Type: shared.AddressTypeEnumBilling,
                    },
                    Email: codataccounting.String("Julianne_Wintheiser@gmail.com"),
                    ModifiedDate: codataccounting.String("earum"),
                    Name: codataccounting.String("Tomas Kiehn"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "perferendis",
                            Type: shared.PhoneNumberTypeEnumLandline,
                        },
                        shared.PhoneNumbersitems{
                            Number: "saepe",
                            Type: shared.PhoneNumberTypeEnumPrimary,
                        },
                        shared.PhoneNumbersitems{
                            Number: "architecto",
                            Type: shared.PhoneNumberTypeEnumPrimary,
                        },
                    },
                    Status: shared.CustomerStatusEnumUnknown,
                },
                shared.Contact{
                    Address: &shared.Addressesitems{
                        City: codataccounting.String("Port Clifford"),
                        Country: codataccounting.String("Lithuania"),
                        Line1: codataccounting.String("excepturi"),
                        Line2: codataccounting.String("alias"),
                        PostalCode: codataccounting.String("03183-7963"),
                        Region: codataccounting.String("blanditiis"),
                        Type: shared.AddressTypeEnumBilling,
                    },
                    Email: codataccounting.String("Winona49@gmail.com"),
                    ModifiedDate: codataccounting.String("quidem"),
                    Name: codataccounting.String("Teri Abshire"),
                    Phone: []shared.PhoneNumbersitems{
                        shared.PhoneNumbersitems{
                            Number: "numquam",
                            Type: shared.PhoneNumberTypeEnumMobile,
                        },
                    },
                    Status: shared.CustomerStatusEnumUnknown,
                },
            },
            CustomerName: codataccounting.String("hic"),
            DefaultCurrency: codataccounting.String("blanditiis"),
            EmailAddress: codataccounting.String("at"),
            ID: codataccounting.String("e30f069d-8106-418d-97e1-52297510da80"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("amet"),
            Phone: codataccounting.String("(315) 377-3071"),
            RegistrationNumber: codataccounting.String("laborum"),
            SourceModifiedDate: codataccounting.String("in"),
            Status: shared.CustomerStatusEnumUnknown,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "distinctio": map[string]interface{}{
                        "sint": "odio",
                        "repudiandae": "accusamus",
                        "quasi": "accusantium",
                    },
                },
            },
            TaxNumber: codataccounting.String("dolores"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        CustomerID: "fugiat",
        ForceUpdate: codataccounting.Bool(false),
        TimeoutInMinutes: codataccounting.Int(664312),
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
