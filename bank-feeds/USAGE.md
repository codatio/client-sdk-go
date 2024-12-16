<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v8"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := bankfeeds.New(
		bankfeeds.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: bankfeeds.String("Requested early access to the new financing scheme."),
		Name:        "Technicalium",
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