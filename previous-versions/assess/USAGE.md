<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/codatio/client-sdk-go/previous-versions/assess"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/models/operations"
)

func main() {
    s := assess.New(
        assess.WithSecurity(shared.Security{
            AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
        }),
    )

    ctx := context.Background()
    res, err := s.DataIntegrity.Details(ctx, operations.ListDataTypeDataIntegrityDetailsRequest{
        CompanyID: "8a210b68-6988-11ed-a1eb-0242ac120002",
        DataType: shared.DataIntegrityDataTypeBankingAccounts,
        OrderBy: assess.String("-modifiedDate"),
        Page: assess.Int(1),
        PageSize: assess.Int(100),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Details != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->