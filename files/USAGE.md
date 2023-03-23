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
    s := codatio.New(
        codatio.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    req := operations.DownloadFilesRequest{
        CompanyID: "89bd9d8d-69a6-474e-8f46-7cc8796ed151",
        Date: "2022-07-29T10:12:25.518Z",
    }

    ctx := context.Background()
    res, err := s.Files.DownloadFiles(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->