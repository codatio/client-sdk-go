<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	commerce "github.com/codatio/client-sdk-go/previous-versions/commerce/v3"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/shared"
	"log"
)

func main() {
	s := commerce.New(
		commerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Customers.Get(ctx, operations.GetCustomerRequest{
		CompanyID:    "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
		CustomerID:   "7110701885",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Customer != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->