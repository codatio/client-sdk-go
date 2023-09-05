<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/bank-feeds"
	"github.com/codatio/client-sdk-go/bank-feeds/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/pkg/models/operations"
)

func main() {
    s := codatbankfeeds.New(
        codatbankfeeds.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountMapping.Create(ctx, operations.CreateBankAccountMappingRequest{
        BankFeedAccountMapping: &shared.BankFeedAccountMapping{
            FeedStartDate: codatbankfeeds.String("2022-10-23T00:00:00.000Z"),
            SourceAccountID: codatbankfeeds.String("provident"),
            TargetAccountID: codatbankfeeds.String("distinctio"),
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankFeedAccountMappingResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->