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
    s := codatio.New(codatio.WithSecurity(
        shared.Security{
            AuthHeader: shared.SchemeAuthHeader{
                APIKey: "YOUR_API_KEY_HERE",
            },
        },
    ))
    
    req := operations.DownloadFilesRequest{
        PathParams: operations.DownloadFilesPathParams{
            CompanyID: "unde",
        },
        QueryParams: operations.DownloadFilesQueryParams{
            Date: "2022-08-06T14:51:16.891Z",
        },
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