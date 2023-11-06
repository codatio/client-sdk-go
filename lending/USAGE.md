<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	lending "github.com/codatio/client-sdk-go/lending/v4"
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/operations"
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/shared"
	"log"
)

func main() {
	s := lending.New(
		lending.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.AccountingBankData.ListTransactions(ctx, operations.ListAccountingBankAccountTransactionsRequest{
		AccountID:    "string",
		CompanyID:    "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
		OrderBy:      lending.String("-modifiedDate"),
		Page:         lending.Int(1),
		PageSize:     lending.Int(100),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountingBankTransactions != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->