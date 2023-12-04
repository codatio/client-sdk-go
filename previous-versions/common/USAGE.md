<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"log"
)

func main() {
	s := common.New(
		common.WithSecurity(shared.Security{
			AuthHeader: "",
		}),
	)

	ctx := context.Background()
	res, err := s.Settings.CreateAPIKey(ctx, &shared.CreateAPIKey{
		Name: common.String("azure-invoice-finance-processor"),
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