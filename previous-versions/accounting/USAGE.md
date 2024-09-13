<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/accounting"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"log"
)

func main() {
	s := accounting.New(
		accounting.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.AccountTransactions.Get(ctx, operations.GetAccountTransactionRequest{
		AccountTransactionID: "<value>",
		CompanyID:            "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID:         "2e9d2c44-f675-40ba-8049-353bfcb5e171",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountTransaction != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->