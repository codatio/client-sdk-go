<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v3"
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/shared"
	"log"
)

func main() {
	s := syncforpayables.New(
		syncforpayables.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforpayables.String("Requested early access to the new financing scheme."),
		Groups: []shared.GroupReference{
			shared.GroupReference{
				ID: syncforpayables.String("60d2fa12-8a04-11ee-b9d1-0242ac120002"),
			},
		},
		Name: "Technicalium",
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