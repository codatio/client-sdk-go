# Files

<!-- Start Codat Library Description -->
﻿Use Codat's Files API to upload your SMB customers' files.
<!-- End Codat Library Description -->

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/previous-versions/files
```
<!-- End SDK Installation -->

## Example Usage
<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/files"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/operations"
)

func main() {
    s := codatfiles.New(
        codatfiles.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.Files.DownloadFiles(ctx, operations.DownloadFilesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        Date: codatfiles.String("2022-10-23T00:00:00.000Z"),
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Files](docs/sdks/files/README.md)
