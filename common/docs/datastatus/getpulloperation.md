# GetPullOperation
Available in: `DataStatus`

Retrieve information about a single dataset or pull operation.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/common"
	"github.com/codatio/client-sdk-go/common/pkg/models/operations"
)

func main() {
    s := codatcommon.New(
        codatcommon.WithSecurity(shared.Security{
            AuthHeader: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetPullOperationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        DatasetID: "eaed9f0f-e77b-4bc9-a58f-ab8b4b99ab18",
    }

    res, err := s.DataStatus.GetPullOperation(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PullOperation != nil {
        // handle response
    }
}
```