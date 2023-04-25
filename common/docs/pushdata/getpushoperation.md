# GetPushOperation
Available in: `PushData`

Retrieve push operation.

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
    req := operations.GetPushOperationRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        PushOperationKey: "05dfc2dd-f7cc-478c-a1ba-928fc816742c",
    }

    res, err := s.PushData.GetPushOperation(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushOperation != nil {
        // handle response
    }
}
```