<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v3"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/types"
	"log"
)

func main() {
	s := syncforexpenses.New(
		syncforexpenses.WithSecurity("Basic BASE_64_ENCODED(API_KEY)"),
	)

	ctx := context.Background()
	res, err := s.Accounts.Create(ctx, operations.CreateAccountRequest{
		Account: &shared.Account{
			Currency:               syncforexpenses.String("USD"),
			CurrentBalance:         types.MustNewDecimalFromString("0"),
			Description:            syncforexpenses.String("Invoices the business has issued but has not yet collected payment on."),
			FullyQualifiedCategory: syncforexpenses.String("Asset.Current"),
			FullyQualifiedName:     syncforexpenses.String("Cash On Hand"),
			ID:                     syncforexpenses.String("1b6266d1-1e44-46c5-8eb5-a8f98e03124e"),
			Metadata:               &shared.AccountMetadata{},
			ModifiedDate:           syncforexpenses.String("2022-10-23T00:00:00.000Z"),
			Name:                   syncforexpenses.String("Accounts Receivable"),
			NominalCode:            syncforexpenses.String("610"),
			SourceModifiedDate:     syncforexpenses.String("2022-10-23T00:00:00.000Z"),
			Status:                 shared.AccountStatusActive.ToPointer(),
			Type:                   shared.AccountTypeAsset.ToPointer(),
			ValidDatatypeLinks: []shared.AccountValidDataTypeLinks{
				shared.AccountValidDataTypeLinks{
					Links: []string{
						"Money",
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