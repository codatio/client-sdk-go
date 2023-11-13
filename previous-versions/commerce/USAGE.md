<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/commerce"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/commerce/pkg/models/shared"
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
		CustomerID:   "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Customer != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->