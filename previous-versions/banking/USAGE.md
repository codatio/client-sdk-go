<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/banking"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/operations"
)

func main() {
    s := codatbanking.New(
        codatbanking.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.AccountBalances.List(ctx, operations.ListAccountBalancesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
        OrderBy: codatbanking.String("-modifiedDate"),
        Page: codatbanking.Int(1),
        PageSize: codatbanking.Int(100),
        Query: codatbanking.String("corrupti"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountBalances != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->