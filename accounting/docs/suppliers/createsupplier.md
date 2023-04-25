# CreateSupplier
Available in: `Suppliers`

Push suppliers

Required data may vary by integration. To see what data to post, first call [Get create/update supplier model](https://docs.codat.io/accounting-api#/operations/get-create-update-suppliers-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers) for integrations that support creating suppliers.

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
    req := operations.CreateSupplierRequest{
        Supplier: &shared.Supplier{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("Corpus Christi"),
                    Country: codataccounting.String("Gambia"),
                    Line1: codataccounting.String("incidunt"),
                    Line2: codataccounting.String("eos"),
                    PostalCode: codataccounting.String("40680-2176"),
                    Region: codataccounting.String("optio"),
                    Type: shared.AddressTypeEnumDelivery,
                },
            },
            ContactName: codataccounting.String("debitis"),
            DefaultCurrency: codataccounting.String("corporis"),
            EmailAddress: codataccounting.String("neque"),
            ID: codataccounting.String("2b6526f8-6285-43fe-a859-ce322231fe66"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Phone: codataccounting.String("1-472-281-9763 x876"),
            RegistrationNumber: codataccounting.String("sit"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Status: shared.SupplierStatusEnumActive,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "tempore": map[string]interface{}{
                        "at": "fugit",
                        "cupiditate": "dicta",
                        "libero": "recusandae",
                    },
                    "libero": map[string]interface{}{
                        "quae": "eaque",
                        "est": "sed",
                        "dolorum": "laborum",
                    },
                    "atque": map[string]interface{}{
                        "aliquam": "perspiciatis",
                        "labore": "esse",
                    },
                },
            },
            SupplierName: codataccounting.String("unde"),
            TaxNumber: codataccounting.String("recusandae"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(842306),
    }

    res, err := s.Suppliers.CreateSupplier(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateSupplierResponse != nil {
        // handle response
    }
}
```