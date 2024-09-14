<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	syncforexpensesversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"log"
)

func main() {
	s := syncforexpensesversion1.New(
		syncforexpensesversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.CreateCompany(ctx, &shared.CompanyRequestBody{
		Description: syncforexpensesversion1.String("Requested early access to the new financing scheme."),
		Groups: []shared.GroupReference{
			shared.GroupReference{
				ID: syncforexpensesversion1.String("60d2fa12-8a04-11ee-b9d1-0242ac120002"),
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