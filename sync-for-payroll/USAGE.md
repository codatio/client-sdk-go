<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll/v2"
	"github.com/codatio/client-sdk-go/sync-for-payroll/v2/pkg/models/shared"
	"log"
)

func main() {
	s := syncforpayroll.New(
		syncforpayroll.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforpayroll.String("Requested early access to the new financing scheme."),
		Name:        "Bank of Dave",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Company != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->