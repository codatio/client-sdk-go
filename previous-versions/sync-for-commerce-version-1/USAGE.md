<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	syncforcommerceversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"log"
)

func main() {
	s := syncforcommerceversion1.New(
		syncforcommerceversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.SyncFlowPreferences.GetConfigTextSyncFlow(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.LocalizationInfo != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->