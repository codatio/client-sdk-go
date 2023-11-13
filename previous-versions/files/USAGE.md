<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/files"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/shared"
	"log"
)

func main() {
	s := files.New(
		files.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Files.DownloadFiles(ctx, operations.DownloadFilesRequest{
		CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
		Date:      files.String("2022-10-23T00:00:00.000Z"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Data != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->