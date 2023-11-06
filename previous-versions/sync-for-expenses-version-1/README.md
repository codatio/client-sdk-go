# Sync for Expenses version 1

<!-- Start Codat Library Description -->
Push expenses to accounting platforms.
<!-- End Codat Library Description -->

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1
```
<!-- End SDK Installation -->

## Example Usage
<!-- Start SDK Example Usage -->
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
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Error Handling -->
# Error Handling

Handling errors in your SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.


<!-- End Error Handling -->



<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.codat.io` | None |

For example:


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
		syncforexpensesversion1.WithServerIndex(0),
	)

	ctx := context.Background()
	res, err := s.Companies.CreateCompany(ctx, &shared.CompanyRequestBody{
		Description: syncforexpensesversion1.String("Requested early access to the new financing scheme."),
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


## Override Server URL Per-Client

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
		syncforexpensesversion1.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
		syncforexpensesversion1.WithServerURL("https://api.codat.io"),
	)

	ctx := context.Background()
	res, err := s.Companies.CreateCompany(ctx, &shared.CompanyRequestBody{
		Description: syncforexpensesversion1.String("Requested early access to the new financing scheme."),
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
# Custom HTTP Client

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

<!-- Placeholder for Future Speakeasy SDK Sections -->


### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)