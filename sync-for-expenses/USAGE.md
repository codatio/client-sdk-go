<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v4"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/models/shared"
	"log"
)

func main() {
	s := syncforexpenses.New(
		syncforexpenses.WithSecurity("Basic BASE_64_ENCODED(API_KEY)"),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforexpenses.String("Requested early access to the new financing scheme."),
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
<!-- End SDK Example Usage -->