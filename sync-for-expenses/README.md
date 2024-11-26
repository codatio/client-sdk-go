# Sync for Expenses

<!-- Start Codat Library Description -->
﻿Embedded accounting integrations for corporate card providers.
<!-- End Codat Library Description -->

<!-- Start Summary [summary] -->
## Summary

Sync for Expenses: The API for Sync for Expenses.

Sync for Expenses is an API and a set of supporting tools. It has been built to
enable corporate card and expense management platforms to provide high-quality
integrations with multiple accounting software through a standardized API.

[Explore product](https://docs.codat.io/sync-for-expenses/overview) | [See our OpenAPI spec](https://github.com/codatio/oas)

Not seeing the endpoints you're expecting? We've [reorganized our products](https://docs.codat.io/updates/230901-new-products), and you may be using a [different version of Sync for Expenses](https://docs.codat.io/sync-for-expenses-v1-api#/).

---
<!-- Start Codat Tags Table -->
## Endpoints

| Endpoints | Description |
| :- |:- |
| Companies | Create and manage your SMB users' companies. |
| Connections | Create new and manage existing data connections for a company. |
| Configuration | View and manage mapping configuration and defaults for expense transactions. |
| Sync | Monitor the status of data syncs. |
| Expenses | Create and update transactions that represent your customers' spend. |
| Transfers | Create and update transactions that represent the movement of your customers' money. |
| Reimbursements | Create and update transactions that represent your customers' repayable spend. |
| Attachments | Attach receipts to a transaction for a complete audit trail. |
| Transaction status | Monitor the status of individual transactions in data syncs. |
| Manage data | Control and monitor the retrieval of data from an integration. |
| Push operations | View historic push operations. |
| Accounts | Create accounts and view account schemas. |
| Customers | Get, create, and update customers. |
| Suppliers | Get, create, and update suppliers. |
<!-- End Codat Tags Table -->
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
<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/codatio/client-sdk-go/sync-for-expenses
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
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v5"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v5/pkg/models/shared"
	"log"
)

func main() {
	s := syncforexpenses.New(
		syncforexpenses.WithSecurity("Basic BASE_64_ENCODED(API_KEY)"),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforexpenses.String("Requested early access to the new financing scheme."),
		Name:        "Technicalium",
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

### [Accounts](docs/sdks/accounts/README.md)

* [Create](docs/sdks/accounts/README.md#create) - Create account
* [GetCreateModel](docs/sdks/accounts/README.md#getcreatemodel) - Get create account model

### [Adjustments](docs/sdks/adjustments/README.md)

* [Create](docs/sdks/adjustments/README.md#create) - Create adjustment transaction

### [Attachments](docs/sdks/attachments/README.md)

* [Upload](docs/sdks/attachments/README.md#upload) - Upload attachment

### [BankAccounts](docs/sdks/bankaccounts/README.md)

* [Create](docs/sdks/bankaccounts/README.md#create) - Create bank account
* [GetCreateModel](docs/sdks/bankaccounts/README.md#getcreatemodel) - Get create bank account model


### [Companies](docs/sdks/companies/README.md)

* [Create](docs/sdks/companies/README.md#create) - Create company
* [Delete](docs/sdks/companies/README.md#delete) - Delete a company
* [Get](docs/sdks/companies/README.md#get) - Get company
* [List](docs/sdks/companies/README.md#list) - List companies
* [Update](docs/sdks/companies/README.md#update) - Update company

### [CompanyInfo](docs/sdks/companyinfo/README.md)

* [Get](docs/sdks/companyinfo/README.md#get) - Get company info

### [Configuration](docs/sdks/configuration/README.md)

* [Get](docs/sdks/configuration/README.md#get) - Get company configuration
* [Set](docs/sdks/configuration/README.md#set) - Set company configuration

### [Connections](docs/sdks/connections/README.md)

* [Create](docs/sdks/connections/README.md#create) - Create connection
* [CreatePartnerExpenseConnection](docs/sdks/connections/README.md#createpartnerexpenseconnection) - Create partner expense connection
* [Delete](docs/sdks/connections/README.md#delete) - Delete connection
* [Get](docs/sdks/connections/README.md#get) - Get connection
* [List](docs/sdks/connections/README.md#list) - List connections
* [Unlink](docs/sdks/connections/README.md#unlink) - Unlink connection

### [Customers](docs/sdks/customers/README.md)

* [Create](docs/sdks/customers/README.md#create) - Create customer
* [Get](docs/sdks/customers/README.md#get) - Get customer
* [List](docs/sdks/customers/README.md#list) - List customers
* [Update](docs/sdks/customers/README.md#update) - Update customer

### [Expenses](docs/sdks/expenses/README.md)

* [Create](docs/sdks/expenses/README.md#create) - Create expense transaction
* [Update](docs/sdks/expenses/README.md#update) - Update expense transactions

### [ManageData](docs/sdks/managedata/README.md)

* [Get](docs/sdks/managedata/README.md#get) - Get data status
* [GetPullOperation](docs/sdks/managedata/README.md#getpulloperation) - Get pull operation
* [ListPullOperations](docs/sdks/managedata/README.md#listpulloperations) - List pull operations
* [RefreshAllDataTypes](docs/sdks/managedata/README.md#refreshalldatatypes) - Refresh all data
* [RefreshDataType](docs/sdks/managedata/README.md#refreshdatatype) - Refresh data type

### [MappingOptions](docs/sdks/mappingoptions/README.md)

* [GetMappingOptions](docs/sdks/mappingoptions/README.md#getmappingoptions) - Mapping options

### [PushOperations](docs/sdks/pushoperations/README.md)

* [Get](docs/sdks/pushoperations/README.md#get) - Get push operation
* [List](docs/sdks/pushoperations/README.md#list) - List push operations

### [Reimbursements](docs/sdks/reimbursements/README.md)

* [Create](docs/sdks/reimbursements/README.md#create) - Create reimbursable expense transaction
* [Update](docs/sdks/reimbursements/README.md#update) - Update reimbursable expense transaction

### [Suppliers](docs/sdks/suppliers/README.md)

* [Create](docs/sdks/suppliers/README.md#create) - Create supplier
* [Get](docs/sdks/suppliers/README.md#get) - Get supplier
* [List](docs/sdks/suppliers/README.md#list) - List suppliers
* [Update](docs/sdks/suppliers/README.md#update) - Update supplier

### [Sync](docs/sdks/sync/README.md)

* [Get](docs/sdks/sync/README.md#get) - Get sync status
* [GetLastSuccessfulSync](docs/sdks/sync/README.md#getlastsuccessfulsync) - Last successful sync
* [GetLatestSync](docs/sdks/sync/README.md#getlatestsync) - Latest sync status
* [List](docs/sdks/sync/README.md#list) - List sync statuses

### [TransactionStatus](docs/sdks/transactionstatus/README.md)

* [Get](docs/sdks/transactionstatus/README.md#get) - Get sync transaction
* [List](docs/sdks/transactionstatus/README.md#list) - List sync transactions

### [Transfers](docs/sdks/transfers/README.md)

* [Create](docs/sdks/transfers/README.md#create) - Create transfer transaction

</details>
<!-- End Available Resources and Operations [operations] -->





<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v5"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v5/pkg/retry"
	"log"
	"pkg/models/operations"
)

func main() {
	s := syncforexpenses.New(
		syncforexpenses.WithSecurity("Basic BASE_64_ENCODED(API_KEY)"),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforexpenses.String("Requested early access to the new financing scheme."),
		Name:        "Technicalium",
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
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v5"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v5/pkg/retry"
	"log"
)

func main() {
	s := syncforexpenses.New(
		syncforexpenses.WithRetryConfig(
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
		syncforexpenses.WithSecurity("Basic BASE_64_ENCODED(API_KEY)"),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforexpenses.String("Requested early access to the new financing scheme."),
		Name:        "Technicalium",
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

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Create` function may return the following errors:

| Error Type             | Status Code                       | Content Type     |
| ---------------------- | --------------------------------- | ---------------- |
| sdkerrors.ErrorMessage | 400, 401, 402, 403, 429, 500, 503 | application/json |
| sdkerrors.SDKError     | 4XX, 5XX                          | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v5"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v5/pkg/models/sdkerrors"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v5/pkg/models/shared"
	"log"
)

func main() {
	s := syncforexpenses.New(
		syncforexpenses.WithSecurity("Basic BASE_64_ENCODED(API_KEY)"),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforexpenses.String("Requested early access to the new financing scheme."),
		Name:        "Technicalium",
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

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v5"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v5/pkg/models/shared"
	"log"
)

func main() {
	s := syncforexpenses.New(
		syncforexpenses.WithServerURL("https://api.codat.io"),
		syncforexpenses.WithSecurity("Basic BASE_64_ENCODED(API_KEY)"),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforexpenses.String("Requested early access to the new financing scheme."),
		Name:        "Technicalium",
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

| Name         | Type   | Scheme  |
| ------------ | ------ | ------- |
| `AuthHeader` | apiKey | API key |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	syncforexpenses "github.com/codatio/client-sdk-go/sync-for-expenses/v5"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v5/pkg/models/shared"
	"log"
)

func main() {
	s := syncforexpenses.New(
		syncforexpenses.WithSecurity("Basic BASE_64_ENCODED(API_KEY)"),
	)

	ctx := context.Background()
	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: syncforexpenses.String("Requested early access to the new financing scheme."),
		Name:        "Technicalium",
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