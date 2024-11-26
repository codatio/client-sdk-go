<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v5"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/shared"
	"log"
)

func main() {
	s := syncforpayables.New(
		syncforpayables.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.List(ctx, operations.ListCompaniesRequest{
		Page:     syncforpayables.Int(1),
		PageSize: syncforpayables.Int(100),
		Query:    syncforpayables.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
		OrderBy:  syncforpayables.String("-modifiedDate"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Companies != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->