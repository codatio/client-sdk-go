<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/codatio/client-sdk-go/files"
    "github.com/codatio/client-sdk-go/files/pkg/models/shared"
    "github.com/codatio/client-sdk-go/files/pkg/models/operations"
)

func main() {
    s := codatfiles.New(
        codatfiles.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    req := operations.DownloadFilesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Date: "2022-10-23T00:00:00Z",
    }

    ctx := context.Background()
    res, err := s.Files.DownloadFiles(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Data != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->