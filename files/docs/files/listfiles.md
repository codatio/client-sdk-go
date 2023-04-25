# ListFiles
Available in: `Files`

Returns an array of files that have been uploaded for a given company.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/files"
	"github.com/codatio/client-sdk-go/files/pkg/models/operations"
)

func main() {
    s := codatfiles.New(
        codatfiles.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListFilesRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
    }

    res, err := s.Files.ListFiles(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Files != nil {
        // handle response
    }
}
```