# CreateItem
Available in: `Items`

Posts a new item to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create item model](https://docs.codat.io/accounting-api#/operations/get-create-items-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items) for integrations that support creating items.

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
    req := operations.CreateItemRequest{
        Item: &shared.Item{
            BillItem: &shared.BillItem{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("2743fd6c-2a10-4e6c-a978-ec256a5b0922"),
                    Name: codataccounting.String("Melba Schmeler"),
                },
                Description: codataccounting.String("dignissimos"),
                TaxRateRef: &shared.TaxRateRef{
                    EffectiveTaxRate: codataccounting.Float64(6172.71),
                    ID: codataccounting.String("96c977bb-c57f-4389-a8a8-600c58d67d63"),
                    Name: codataccounting.String("Eddie Murphy"),
                },
                UnitPrice: codataccounting.Float64(3900.58),
            },
            Code: codataccounting.String("quos"),
            ID: codataccounting.String("464579cf-c6c0-4e50-bf56-831f1d8ed87b"),
            InvoiceItem: &shared.InvoiceItem{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("28e8afab-c986-4e24-9e43-b2342417d13e"),
                    Name: codataccounting.String("Miss Jeannie Hudson"),
                },
                Description: codataccounting.String("cupiditate"),
                TaxRateRef: &shared.TaxRateRef{
                    EffectiveTaxRate: codataccounting.Float64(6338.87),
                    ID: codataccounting.String("e4ae8ab4-a9c4-492c-9e8b-a5d4aa4a508b"),
                    Name: codataccounting.String("Miss Rodney Ledner"),
                },
                UnitPrice: codataccounting.Float64(5704.27),
            },
            IsBillItem: false,
            IsInvoiceItem: false,
            ItemStatus: shared.ItemStatusEnumActive,
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Name: codataccounting.String("Jordan Stiedemann"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            Type: "beatae",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        TimeoutInMinutes: codataccounting.Int(697543),
    }

    res, err := s.Items.CreateItem(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateItemResponse != nil {
        // handle response
    }
}
```