<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"log"
	"net/http"
)

func main() {
	s := assess.New(
		assess.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Reports.GenerateLoanSummary(ctx, operations.GenerateLoanSummaryRequest{
		CompanyID:  "8a210b68-6988-11ed-a1eb-0242ac120002",
		SourceType: operations.SourceTypeAccounting,
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->