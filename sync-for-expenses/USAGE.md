<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v5"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v5/pkg/models/shared"
	"log"
)

func main() {
	s := syncforexpenses.New(
		syncforexpenses.WithSecurity("Basic BASE_64_ENCODED(API_KEY)"),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforexpenses.String("Requested early access to the new financing scheme."),
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