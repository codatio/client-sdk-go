<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/banking"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/shared"
	"log"
)

func main() {
	s := banking.New(
		banking.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.AccountBalances.List(ctx, operations.ListAccountBalancesRequest{
		CompanyID:    "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
		OrderBy:      banking.String("-modifiedDate"),
		Page:         banking.Int(1),
		PageSize:     banking.Int(100),
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