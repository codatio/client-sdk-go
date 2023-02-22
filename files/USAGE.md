<!-- Start SDK Example Usage -->
```go
package main

import (
    "log"
    "github.com/codatio/client-sdk-go/files"
    "github.com/codatio/client-sdk-go/files/pkg/models/shared"
    "github.com/codatio/client-sdk-go/files/pkg/models/operations"
)

func main() {
    opts := []codatio.SDKOption{
        codatio.WithSecurity(
            shared.Security{
                APIKey: shared.SchemeAPIKey{
                    APIKey: "YOUR_API_KEY_HERE",
                },
            }
        ),
    }

    s := codatio.New(opts...)
    
    req := operations.DownloadFilesRequest{
        Security: operations.DownloadFilesSecurity{
            APIKey: shared.SchemeAPIKey{
                APIKey: "YOUR_API_KEY_HERE",
            },
        },
        PathParams: operations.DownloadFilesPathParams{
            CompanyID: "unde",
        },
        QueryParams: operations.DownloadFilesQueryParams{
            Date: "2022-07-21T05:02:33.025Z",
        },
    }
    
    res, err := s.Files.DownloadFiles(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
```
<!-- End SDK Example Usage -->