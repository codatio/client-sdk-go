# Sync for Expenses version 1

<!-- Start Codat Library Description -->
Push expenses to accounting platforms.
<!-- End Codat Library Description -->

<!-- Start Summary [summary] -->
## Summary

Sync for Expenses (v1): The API for Sync for Expenses.

Sync for Expenses is an API and a set of supporting tools. It has been built to
enable corporate card and expense management platforms to provide high-quality
integrations with multiple accounting software through a standardized API.

<!-- Start Codat Tags Table -->
## Endpoints

| Endpoints | Description |
| :- |:- |
| Connections | Create and manage partner expense connection. |
| Configuration | Companies sync configuration. |
| Mapping options | Mapping options for a companies expenses. |
| Sync | Triggering a new sync of expenses to accounting software. |
| Sync status | Check the status of ongoing or previous expense syncs. |
| Expenses | Create expense datasets and upload receipts. |
| Transaction status | Retrieve the status of transactions within a sync. |
| Companies | Create and manage your Codat companies. |
<!-- End Codat Tags Table -->

[Read more...](https://docs.codat.io/sync-for-expenses/overview)

[See our OpenAPI spec](https://github.com/codatio/oas)

Not seeing what you expect? [See the main Sync for Expenses API](https://docs.codat.io/sync-for-expenses-api).
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [SDK Installation](#sdk-installation)
* [SDK Example Usage](#sdk-example-usage)
* [Available Resources and Operations](#available-resources-and-operations)
* [Retries](#retries)
* [Error Handling](#error-handling)
* [Server Selection](#server-selection)
* [Custom HTTP Client](#custom-http-client)
* [Authentication](#authentication)
* [Special Types](#special-types)
<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1
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
	syncforexpensesversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"log"
)

func main() {
	s := syncforexpensesversion1.New(
		syncforexpensesversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.CreateCompany(ctx, &shared.CompanyRequestBody{
		Description: syncforexpensesversion1.String("Requested early access to the new financing scheme."),
		Groups: []shared.GroupReference{
			shared.GroupReference{
				ID: syncforexpensesversion1.String("60d2fa12-8a04-11ee-b9d1-0242ac120002"),
			},
		},
		Name: "Technicalium",
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

<details open>
<summary>Available methods</summary>


### [Companies](docs/sdks/companies/README.md)

* [CreateCompany](docs/sdks/companies/README.md#createcompany) - Create company
* [DeleteCompany](docs/sdks/companies/README.md#deletecompany) - Delete a company
* [GetCompany](docs/sdks/companies/README.md#getcompany) - Get company
* [ListCompanies](docs/sdks/companies/README.md#listcompanies) - List companies
* [UpdateCompany](docs/sdks/companies/README.md#updatecompany) - Update company

### [Configuration](docs/sdks/configuration/README.md)

* [GetCompanyConfiguration](docs/sdks/configuration/README.md#getcompanyconfiguration) - Get company configuration
* [SaveCompanyConfiguration](docs/sdks/configuration/README.md#savecompanyconfiguration) - Set company configuration

### [Connections](docs/sdks/connections/README.md)

* [CreateConnection](docs/sdks/connections/README.md#createconnection) - Create connection
* [CreatePartnerExpenseConnection](docs/sdks/connections/README.md#createpartnerexpenseconnection) - Create partner expense connection
* [DeleteConnection](docs/sdks/connections/README.md#deleteconnection) - Delete connection
* [GetConnection](docs/sdks/connections/README.md#getconnection) - Get connection
* [ListConnections](docs/sdks/connections/README.md#listconnections) - List connections
* [Unlink](docs/sdks/connections/README.md#unlink) - Unlink connection

### [Expenses](docs/sdks/expenses/README.md)

* [CreateExpenseDataset](docs/sdks/expenses/README.md#createexpensedataset) - Create expense-transactions
* [UpdateExpenseDataset](docs/sdks/expenses/README.md#updateexpensedataset) - Update expense transactions
* [UploadAttachment](docs/sdks/expenses/README.md#uploadattachment) - Upload attachment

### [MappingOptions](docs/sdks/mappingoptions/README.md)

* [GetMappingOptions](docs/sdks/mappingoptions/README.md#getmappingoptions) - Mapping options

### [Sync](docs/sdks/sync/README.md)

* [InitiateSync](docs/sdks/sync/README.md#initiatesync) - Initiate sync

### [SyncStatus](docs/sdks/syncstatus/README.md)

* [GetLastSuccessfulSync](docs/sdks/syncstatus/README.md#getlastsuccessfulsync) - Last successful sync
* [GetLatestSync](docs/sdks/syncstatus/README.md#getlatestsync) - Latest sync status
* [GetSyncByID](docs/sdks/syncstatus/README.md#getsyncbyid) - Get sync status
* [ListSyncs](docs/sdks/syncstatus/README.md#listsyncs) - List sync statuses

### [TransactionStatus](docs/sdks/transactionstatus/README.md)

* [GetSyncTransaction](docs/sdks/transactionstatus/README.md#getsynctransaction) - Get sync transaction
* [ListSyncTransactions](docs/sdks/transactionstatus/README.md#listsynctransactions) - Get sync transactions

</details>
<!-- End Available Resources and Operations [operations] -->





<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	syncforexpensesversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/retry"
	"log"
	"pkg/models/operations"
)

func main() {
	s := syncforexpensesversion1.New(
		syncforexpensesversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.CreateCompany(ctx, &shared.CompanyRequestBody{
		Description: syncforexpensesversion1.String("Requested early access to the new financing scheme."),
		Groups: []shared.GroupReference{
			shared.GroupReference{
				ID: syncforexpensesversion1.String("60d2fa12-8a04-11ee-b9d1-0242ac120002"),
			},
		},
		Name: "Technicalium",
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
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

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	syncforexpensesversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/retry"
	"log"
)

func main() {
	s := syncforexpensesversion1.New(
		syncforexpensesversion1.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		syncforexpensesversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.CreateCompany(ctx, &shared.CompanyRequestBody{
		Description: syncforexpensesversion1.String("Requested early access to the new financing scheme."),
		Groups: []shared.GroupReference{
			shared.GroupReference{
				ID: syncforexpensesversion1.String("60d2fa12-8a04-11ee-b9d1-0242ac120002"),
			},
		},
		Name: "Technicalium",
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
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

### Example

```go
package main

import (
	"context"
	"errors"
	syncforexpensesversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/sdkerrors"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"log"
)

func main() {
	s := syncforexpensesversion1.New(
		syncforexpensesversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.CreateCompany(ctx, &shared.CompanyRequestBody{
		Description: syncforexpensesversion1.String("Requested early access to the new financing scheme."),
		Groups: []shared.GroupReference{
			shared.GroupReference{
				ID: syncforexpensesversion1.String("60d2fa12-8a04-11ee-b9d1-0242ac120002"),
			},
		},
		Name: "Technicalium",
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
	syncforexpensesversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"log"
)

func main() {
	s := syncforexpensesversion1.New(
		syncforexpensesversion1.WithServerIndex(0),
		syncforexpensesversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.CreateCompany(ctx, &shared.CompanyRequestBody{
		Description: syncforexpensesversion1.String("Requested early access to the new financing scheme."),
		Groups: []shared.GroupReference{
			shared.GroupReference{
				ID: syncforexpensesversion1.String("60d2fa12-8a04-11ee-b9d1-0242ac120002"),
			},
		},
		Name: "Technicalium",
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
	syncforexpensesversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"log"
)

func main() {
	s := syncforexpensesversion1.New(
		syncforexpensesversion1.WithServerURL("https://api.codat.io"),
		syncforexpensesversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.CreateCompany(ctx, &shared.CompanyRequestBody{
		Description: syncforexpensesversion1.String("Requested early access to the new financing scheme."),
		Groups: []shared.GroupReference{
			shared.GroupReference{
				ID: syncforexpensesversion1.String("60d2fa12-8a04-11ee-b9d1-0242ac120002"),
			},
		},
		Name: "Technicalium",
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
	syncforexpensesversion1 "github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"log"
)

func main() {
	s := syncforexpensesversion1.New(
		syncforexpensesversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.CreateCompany(ctx, &shared.CompanyRequestBody{
		Description: syncforexpensesversion1.String("Requested early access to the new financing scheme."),
		Groups: []shared.GroupReference{
			shared.GroupReference{
				ID: syncforexpensesversion1.String("60d2fa12-8a04-11ee-b9d1-0242ac120002"),
			},
		},
		Name: "Technicalium",
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

<!-- Start Codat Support Notes -->
### Support

If you encounter any challenges while utilizing our SDKs, please don't hesitate to reach out for assistance. 
You can raise any issues by contacting your dedicated Codat representative or reaching out to our [support team](mailto:support@codat.io).
We're here to help ensure a smooth experience for you.
<!-- End Codat Support Notes -->

<!-- Start Codat Generated By -->
### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
<!-- End Codat Generated By -->