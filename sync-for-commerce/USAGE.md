<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v2"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"log"
)

func main() {
	s := syncforcommerce.New(
		syncforcommerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.SyncFlowSettings.GetConfigTextSyncFlow(ctx, operations.GetConfigTextSyncFlowRequest{
		Locale: shared.LocaleEnUs,
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.LocalizationInfo != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->