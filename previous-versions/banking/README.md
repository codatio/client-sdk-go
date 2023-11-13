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
### Example

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



<!-- Start Error Handling -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.ErrorMessage              | 400,401,402,403,404,409,429,500,503 | application/json                    |
| sdkerrors.SDKError                  | 400-600                             | */*                                 |

### Example

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

		var e *sdkerrors.ErrorMessage
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling -->



<!-- Start Server Selection -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.codat.io` | None |

#### Example

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
		banking.WithServerIndex(0),
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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
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
		banking.WithServerURL("https://api.codat.io"),
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
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->



<!-- Start Retries -->
## Retries

Some of the endpoints in this SDK support retries.  If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API.  However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a retryConfig object to the call:
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/banking"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/shared"
	"log"
	"pkg/models/operations"
	"pkg/utils"
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
	}, operations.WithRetries(utils.RetryConfig{
		Strategy: "backoff",
		Backoff: &utils.BackoffStrategy{
			InitialInterval: 1,
			MaxInterval:     50,
			Exponent:        1.1,
			MaxElapsedTime:  100,
		},
		RetryConnectionErrors: false,
	}))
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountBalances != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can provide a retryConfig at SDK initialization:
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/banking"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/operations"
	"github.com/codatio/client-sdk-go/previous-versions/banking/pkg/models/shared"
	"log"
	"pkg/models/operations"
	"pkg/utils"
)

func main() {
	s := banking.New(
		banking.WithRetryConfig(utils.RetryConfig{
			Strategy: "backoff",
			Backoff: &utils.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}),
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


<!-- End Retries -->



<!-- Start Authentication -->

## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `AuthHeader` | apiKey       | API key      |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
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
<!-- End Authentication -->

<!-- Placeholder for Future Speakeasy SDK Sections -->


### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)