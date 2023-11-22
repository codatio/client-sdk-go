# Banking

<!-- Start Codat Library Description -->
﻿Use Codat's API to connect to your SMB customer's banks and pull up-to-date standardized account and transaction data from their bank accounts via our partner providers.
<!-- End Codat Library Description -->

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/previous-versions/banking
```
<!-- End SDK Installation -->

## Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/banking"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/shared"
	"log"
)

func main() {
	s := banking.New(
		banking.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.AccountBalances.List(ctx, operations.ListAccountBalancesRequest{
		CompanyID:    "8a210b68-6988-11ed-a1eb-0242ac120002",
		ConnectionID: "2e9d2c44-f675-40ba-8049-353bfcb5e171",
		OrderBy:      banking.String("-modifiedDate"),
		Page:         banking.Int(1),
		PageSize:     banking.Int(100),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountBalances != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [AccountBalances](docs/sdks/accountbalances/README.md)

* [List](docs/sdks/accountbalances/README.md#list) - List account balances

### [Accounts](docs/sdks/accounts/README.md)

* [Get](docs/sdks/accounts/README.md#get) - Get account
* [List](docs/sdks/accounts/README.md#list) - List accounts

### [TransactionCategories](docs/sdks/transactioncategories/README.md)

* [Get](docs/sdks/transactioncategories/README.md#get) - Get transaction category
* [List](docs/sdks/transactioncategories/README.md#list) - List transaction categories

### [Transactions](docs/sdks/transactions/README.md)

* [Get](docs/sdks/transactions/README.md#get) - Get bank transaction
* [List](docs/sdks/transactions/README.md#list) - List transactions
* [~~ListBankTransactions~~](docs/sdks/transactions/README.md#listbanktransactions) - List banking transactions :warning: **Deprecated** Use `List` instead.
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->



<!-- End Dev Containers -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->