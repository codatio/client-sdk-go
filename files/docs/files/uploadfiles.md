# UploadFiles
Available in: `Files`

Upload files

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
    req := operations.UploadFilesRequest{
        RequestBody: &operations.UploadFilesRequestBody{
            Content: []byte("corrupti"),
            RequestBody: "provident",
        },
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
    }

    res, err := s.Files.UploadFiles(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```