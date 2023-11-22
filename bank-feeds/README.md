# Bank Feeds

<!-- Start Codat Library Description -->
﻿Bank Feeds API enables your SMB users to set up bank feeds from accounts in your application to supported accounting platforms.
<!-- End Codat Library Description -->

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/bank-feeds
```
<!-- End SDK Installation -->

## Example Usage
<!-- Start SDK Example Usage -->
### Example

```go
package main

import (
	"context"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v4"
	"github.com/codatio/client-sdk-go/bank-feeds/v4/pkg/models/shared"
	"log"
)

func main() {
	s := bankfeeds.New(
		bankfeeds.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: bankfeeds.String("Requested early access to the new financing scheme."),
		Name:        "Bank of Dave",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Company != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Companies](docs/sdks/companies/README.md)

* [Create](docs/sdks/companies/README.md#create) - Create company
* [Delete](docs/sdks/companies/README.md#delete) - Delete a company
* [Get](docs/sdks/companies/README.md#get) - Get company
* [List](docs/sdks/companies/README.md#list) - List companies
* [Update](docs/sdks/companies/README.md#update) - Update company

### [Connections](docs/sdks/connections/README.md)

* [Create](docs/sdks/connections/README.md#create) - Create connection
* [Delete](docs/sdks/connections/README.md#delete) - Delete connection
* [Get](docs/sdks/connections/README.md#get) - Get connection
* [List](docs/sdks/connections/README.md#list) - List connections
* [Unlink](docs/sdks/connections/README.md#unlink) - Unlink connection

### [AccountMapping](docs/sdks/accountmapping/README.md)

* [Create](docs/sdks/accountmapping/README.md#create) - Create bank feed account mapping
* [Get](docs/sdks/accountmapping/README.md#get) - List bank feed account mappings

### [SourceAccounts](docs/sdks/sourceaccounts/README.md)

* [Create](docs/sdks/sourceaccounts/README.md#create) - Create source account
* [Delete](docs/sdks/sourceaccounts/README.md#delete) - Delete source account
* [DeleteCredentials](docs/sdks/sourceaccounts/README.md#deletecredentials) - Delete all source account credentials
* [GenerateCredentials](docs/sdks/sourceaccounts/README.md#generatecredentials) - Generate source account credentials
* [List](docs/sdks/sourceaccounts/README.md#list) - List source accounts
* [Update](docs/sdks/sourceaccounts/README.md#update) - Update source account

### [Transactions](docs/sdks/transactions/README.md)

* [Create](docs/sdks/transactions/README.md#create) - Create bank transactions
* [GetCreateOperation](docs/sdks/transactions/README.md#getcreateoperation) - Get create operation
* [ListCreateOperations](docs/sdks/transactions/README.md#listcreateoperations) - List create operations
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->



<!-- End Dev Containers -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Retries -->
## Retries

Some of the endpoints in this SDK support retries.  If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API.  However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a retryConfig object to the call:
```go
package main

import (
	"context"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v4"
	"github.com/codatio/client-sdk-go/bank-feeds/v4/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v4/pkg/utils"
	"log"
	"pkg/models/operations"
)

func main() {
	s := bankfeeds.New(
		bankfeeds.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: bankfeeds.String("Requested early access to the new financing scheme."),
		Name:        "Bank of Dave",
	}, operations.WithRetries(
		utils.RetryConfig{
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

	if res.Company != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can provide a retryConfig at SDK initialization:
```go
package main

import (
	"context"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v4"
	"github.com/codatio/client-sdk-go/bank-feeds/v4/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v4/pkg/utils"
	"log"
)

func main() {
	s := bankfeeds.New(
		bankfeeds.WithRetryConfig(
			utils.RetryConfig{
				Strategy: "backoff",
				Backoff: &utils.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		bankfeeds.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: bankfeeds.String("Requested early access to the new financing scheme."),
		Name:        "Bank of Dave",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Company != nil {
		// handle response
	}
}

```


<!-- End Retries -->



<!-- Start Error Handling -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ErrorMessage      | 400,401,402,403,429,500,503 | application/json            |
| sdkerrors.SDKError          | 400-600                     | */*                         |

### Example

```go
package main

import (
	"context"
	"errors"
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v4"
	"github.com/codatio/client-sdk-go/bank-feeds/v4/pkg/models/sdkerrors"
	"github.com/codatio/client-sdk-go/bank-feeds/v4/pkg/models/shared"
	"log"
)

func main() {
	s := bankfeeds.New(
		bankfeeds.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: bankfeeds.String("Requested early access to the new financing scheme."),
		Name:        "Bank of Dave",
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
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v4"
	"github.com/codatio/client-sdk-go/bank-feeds/v4/pkg/models/shared"
	"log"
)

func main() {
	s := bankfeeds.New(
		bankfeeds.WithServerIndex(0),
		bankfeeds.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: bankfeeds.String("Requested early access to the new financing scheme."),
		Name:        "Bank of Dave",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Company != nil {
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
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v4"
	"github.com/codatio/client-sdk-go/bank-feeds/v4/pkg/models/shared"
	"log"
)

func main() {
	s := bankfeeds.New(
		bankfeeds.WithServerURL("https://api.codat.io"),
		bankfeeds.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: bankfeeds.String("Requested early access to the new financing scheme."),
		Name:        "Bank of Dave",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Company != nil {
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
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v4"
	"github.com/codatio/client-sdk-go/bank-feeds/v4/pkg/models/shared"
	"log"
)

func main() {
	s := bankfeeds.New(
		bankfeeds.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: bankfeeds.String("Requested early access to the new financing scheme."),
		Name:        "Bank of Dave",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Company != nil {
		// handle response
	}
}

```
<!-- End Authentication -->

<!-- Placeholder for Future Speakeasy SDK Sections -->