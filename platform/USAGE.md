<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	platform "github.com/codatio/client-sdk-go/platform/v4"
	"github.com/codatio/client-sdk-go/platform/v4/pkg/models/shared"
	"log"
)

func main() {
	s := platform.New(
		platform.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Settings.CreateAPIKey(ctx, &shared.CreateAPIKey{
		Name: platform.String("azure-invoice-finance-processor"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.APIKeyDetails != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->