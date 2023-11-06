<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v2"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/types"
	"log"
)

func main() {
	s := syncforpayables.New(
		syncforpayables.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Accounts.Create(ctx, operations.CreateAccountRequest{
		Account: &shared.Account{
			Currency:               syncforpayables.String("USD"),
			CurrentBalance:         types.MustNewDecimalFromString("0"),
			Description:            syncforpayables.String("Invoices the business has issued but has not yet collected payment on."),
			FullyQualifiedCategory: syncforpayables.String("Asset.Current"),
			FullyQualifiedName:     syncforpayables.String("Cash On Hand"),
			ID:                     syncforpayables.String("1b6266d1-1e44-46c5-8eb5-a8f98e03124e"),
			Metadata:               &shared.Metadata{},
			ModifiedDate:           syncforpayables.String("2022-10-23T00:00:00.000Z"),
			Name:                   syncforpayables.String("Accounts Receivable"),
			NominalCode:            syncforpayables.String("610"),
			SourceModifiedDate:     syncforpayables.String("2022-10-23T00:00:00.000Z"),
			Status:                 shared.AccountStatusActive.ToPointer(),
			SupplementalData: &shared.SupplementalData{
				Content: map[string]map[string]interface{}{
					"key": map[string]interface{}{
						"key": "string",
					},
				},
			},
			Type: shared.AccountTypeAsset.ToPointer(),
			ValidDatatypeLinks: []shared.AccountValidDataTypeLinks{
				shared.AccountValidDataTypeLinks{
					Links: []string{
						"string",
					},
				},
			},
		},
		CompanyID:    "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateAccountResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->