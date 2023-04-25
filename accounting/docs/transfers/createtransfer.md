# CreateTransfer
Available in: `Transfers`

Posts a new transfer to the accounting package for a given company.

Required data may vary by integration. To see what data to post, first call [Get create transfer model](https://docs.codat.io/accounting-api#/operations/get-create-transfers-model).

> **Supported Integrations**
> 
> Check out our [Knowledge UI](https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers) for integrations that support creating transfers.

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
    req := operations.CreateTransferRequest{
        Transfer: &shared.Transfer{
            ContactRef: &shared.TransferContactRef{
                DataType: codataccounting.String("quae"),
                ID: "8b38f1b6-1a33-41a5-8dc1-0294f92fed93",
            },
            Date: codataccounting.String("2022-10-23T00:00:00Z"),
            DepositedRecordRefs: []string{
                "expedita",
                "mollitia",
                "quas",
            },
            Description: codataccounting.String("asperiores"),
            From: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("71e2992c-20ee-4122-8ac3-adc647d240bc"),
                    Name: codataccounting.String("Pamela Watsica"),
                },
                Amount: codataccounting.Float64(5375.31),
                Currency: codataccounting.String("explicabo"),
            },
            ID: codataccounting.String("824ccc6a-2f0f-45b9-93cb-11a7687d3100"),
            Metadata: &shared.Metadata{
                IsDeleted: codataccounting.Bool(false),
            },
            ModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            SourceModifiedDate: codataccounting.String("2022-10-23T00:00:00Z"),
            SupplementalData: &shared.SupplementalData{
                Content: map[string]map[string]interface{}{
                    "deleniti": map[string]interface{}{
                        "sed": "cum",
                        "sint": "soluta",
                        "voluptatem": "repellendus",
                        "dignissimos": "quaerat",
                    },
                    "nisi": map[string]interface{}{
                        "quia": "dolorum",
                        "nihil": "quisquam",
                        "molestiae": "fugiat",
                        "ab": "debitis",
                    },
                    "dolorum": map[string]interface{}{
                        "voluptates": "quam",
                    },
                    "iste": map[string]interface{}{
                        "similique": "sint",
                        "nobis": "distinctio",
                        "earum": "veniam",
                        "maiores": "et",
                    },
                },
            },
            To: &shared.TransferAccount{
                AccountRef: &shared.AccountRef{
                    ID: codataccounting.String("79f650b1-e707-4e7e-8396-713bacce072a"),
                    Name: codataccounting.String("Ms. Taylor Jacobson IV"),
                },
                Amount: codataccounting.Float64(8177.85),
                Currency: codataccounting.String("magni"),
            },
            TrackingCategoryRefs: []shared.TrackingCategoryRef{
                shared.TrackingCategoryRef{
                    ID: "9c10c185-16fd-478b-a262-1272628fa503",
                    Name: codataccounting.String("Chester Cole"),
                },
                shared.TrackingCategoryRef{
                    ID: "7e72b3a6-5024-4b15-bf9b-b6ef72a50871",
                    Name: codataccounting.String("Wade Medhurst"),
                },
            },
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Transfers.CreateTransfer(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateTransferResponse != nil {
        // handle response
    }
}
```