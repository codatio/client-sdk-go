# PutSupplier
Available in: `Suppliers`

Push supplier

Required data may vary by integration. To see what data to post, first call [Get create/update supplier model](https://docs.codat.io/accounting-api#/operations/get-create-update-suppliers-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers) for integrations that support updating suppliers.

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
    req := operations.PutSupplierRequest{
        Supplier: &shared.Supplier{
            Addresses: []shared.Addressesitems{
                shared.Addressesitems{
                    City: codataccounting.String("Roswell"),
                    Country: codataccounting.String("United Kingdom"),
                    Line1: codataccounting.String("ducimus"),
                    Line2: codataccounting.String("libero"),
                    PostalCode: codataccounting.String("07954"),
                    Region: codataccounting.String("a"),
                    Type: shared.AddressTypeEnumUnknown,
                },
                shared.Addressesitems{
                    City: codataccounting.String("Yvonneburgh"),
                    Country: codataccounting.String("Mexico"),
                    Line1: codataccounting.String("fugit"),
                    Line2: codataccounting.String("esse"),
                    PostalCode: codataccounting.String("04461"),
                    Region: codataccounting.String("dolore"),
                    Type: shared.AddressTypeEnumDelivery,
                },
            },
            ContactName: codataccounting.String("incidunt"),
            DefaultCurrency: codataccounting.String("consequatur"),
            EmailAddress: codataccounting.String("porro"),
            ID: codataccounting.String("8f08bff1-081e-488f-8699-6c8e22be0a3c"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Phone: codataccounting.String("445.627.8129 x54400"),
            RegistrationNumber: codataccounting.String("minus"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Status: shared.SupplierStatusEnumActive,
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "porro": map[string]interface{}{
                        "voluptatum": "consectetur",
                        "aliquam": "magni",
                    },
                },
            },
            SupplierName: codataccounting.String("in"),
            TaxNumber: codataccounting.String("ipsum"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        ForceUpdate: codataccounting.Bool(false),
        SupplierID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        TimeoutInMinutes: codataccounting.Int(774266),
    }

    res, err := s.Suppliers.PutSupplier(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PutSupplier200ApplicationJSONObject != nil {
        // handle response
    }
}
```