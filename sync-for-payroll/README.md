# Sync for Payroll

<!-- Start Codat Library Description -->
Push payroll to accounting platforms.
<!-- End Codat Library Description -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/sync-for-payroll
```
<!-- End SDK Installation [installation] -->

## Example Usage
<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll/v2"
	"github.com/codatio/client-sdk-go/sync-for-payroll/v2/pkg/models/shared"
	"log"
)

func main() {
	s := syncforpayroll.New(
		syncforpayroll.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforpayroll.String("Requested early access to the new financing scheme."),
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
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
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

### [Accounts](docs/sdks/accounts/README.md)

* [Create](docs/sdks/accounts/README.md#create) - Create account
* [Get](docs/sdks/accounts/README.md#get) - Get account
* [GetCreateModel](docs/sdks/accounts/README.md#getcreatemodel) - Get create account model
* [List](docs/sdks/accounts/README.md#list) - List accounts

### [JournalEntries](docs/sdks/journalentries/README.md)

* [Create](docs/sdks/journalentries/README.md#create) - Create journal entry
* [Delete](docs/sdks/journalentries/README.md#delete) - Delete journal entry
* [Get](docs/sdks/journalentries/README.md#get) - Get journal entry
* [GetCreateModel](docs/sdks/journalentries/README.md#getcreatemodel) - Get create journal entry model
* [List](docs/sdks/journalentries/README.md#list) - List journal entries

### [Journals](docs/sdks/journals/README.md)

* [Create](docs/sdks/journals/README.md#create) - Create journal
* [Get](docs/sdks/journals/README.md#get) - Get journal
* [GetCreateModel](docs/sdks/journals/README.md#getcreatemodel) - Get create journal model
* [List](docs/sdks/journals/README.md#list) - List journals

### [ManageData](docs/sdks/managedata/README.md)

* [GetDataStatus](docs/sdks/managedata/README.md#getdatastatus) - Get data status
* [GetPullOperation](docs/sdks/managedata/README.md#getpulloperation) - Get pull operation
* [GetPushOperation](docs/sdks/managedata/README.md#getpushoperation) - Get push operation
* [List](docs/sdks/managedata/README.md#list) - List push operations
* [ListPullOperations](docs/sdks/managedata/README.md#listpulloperations) - List pull operations
* [RefreshAllDataTypes](docs/sdks/managedata/README.md#refreshalldatatypes) - Refresh all data
* [RefreshDataType](docs/sdks/managedata/README.md#refreshdatatype) - Refresh data type

### [CompanyInfo](docs/sdks/companyinfo/README.md)

* [GetAccountingProfile](docs/sdks/companyinfo/README.md#getaccountingprofile) - Get company accounting profile

### [TrackingCategories](docs/sdks/trackingcategories/README.md)

* [Get](docs/sdks/trackingcategories/README.md#get) - Get tracking categories
* [List](docs/sdks/trackingcategories/README.md#list) - List tracking categories
<!-- End Available Resources and Operations [operations] -->





<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries.  If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API.  However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a retryConfig object to the call:
```go
package main

import (
	"context"
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll/v2"
	"github.com/codatio/client-sdk-go/sync-for-payroll/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/v2/pkg/utils"
	"log"
	"pkg/models/operations"
)

func main() {
	s := syncforpayroll.New(
		syncforpayroll.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforpayroll.String("Requested early access to the new financing scheme."),
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
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll/v2"
	"github.com/codatio/client-sdk-go/sync-for-payroll/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payroll/v2/pkg/utils"
	"log"
)

func main() {
	s := syncforpayroll.New(
		syncforpayroll.WithRetryConfig(
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
		syncforpayroll.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforpayroll.String("Requested early access to the new financing scheme."),
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
<!-- End Retries [retries] -->



<!-- Start Error Handling [errors] -->
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
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll/v2"
	"github.com/codatio/client-sdk-go/sync-for-payroll/v2/pkg/models/sdkerrors"
	"github.com/codatio/client-sdk-go/sync-for-payroll/v2/pkg/models/shared"
	"log"
)

func main() {
	s := syncforpayroll.New(
		syncforpayroll.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforpayroll.String("Requested early access to the new financing scheme."),
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
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
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
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll/v2"
	"github.com/codatio/client-sdk-go/sync-for-payroll/v2/pkg/models/shared"
	"log"
)

func main() {
	s := syncforpayroll.New(
		syncforpayroll.WithServerIndex(0),
		syncforpayroll.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforpayroll.String("Requested early access to the new financing scheme."),
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
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll/v2"
	"github.com/codatio/client-sdk-go/sync-for-payroll/v2/pkg/models/shared"
	"log"
)

func main() {
	s := syncforpayroll.New(
		syncforpayroll.WithServerURL("https://api.codat.io"),
		syncforpayroll.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforpayroll.String("Requested early access to the new financing scheme."),
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
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
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
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Authentication [security] -->
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
	syncforpayroll "github.com/codatio/client-sdk-go/sync-for-payroll/v2"
	"github.com/codatio/client-sdk-go/sync-for-payroll/v2/pkg/models/shared"
	"log"
)

func main() {
	s := syncforpayroll.New(
		syncforpayroll.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforpayroll.String("Requested early access to the new financing scheme."),
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
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->


### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)