# Bank Feeds

<!-- Start Codat Library Description -->
﻿Bank Feeds API enables your SMB users to set up bank feeds from accounts in your application to supported accounting platforms.
<!-- End Codat Library Description -->

<!-- Start Summary [summary] -->
## Summary

Bank Feeds API: Bank Feeds API enables your SMB users to set up bank feeds from accounts in your application to supported accounting software.

A bank feed is a connection between a source bank account in your application and a target bank account in a supported accounting software.

[Explore product](https://docs.codat.io/bank-feeds-api/overview) | [See OpenAPI spec](https://github.com/codatio/oas)

---
<!-- Start Codat Tags Table -->
## Endpoints

| Endpoints | Description |
| :- |:- |
| Companies | Create and manage your SMB users' companies. |
| Connections | Create new and manage existing data connections for a company. |
| Source accounts | Provide and manage lists of source bank accounts. |
| Account mapping | Extra functionality for building an account management UI. |
| Company information | Get detailed information about a company from the underlying platform. |
| Transactions | Create new bank account transactions for a company's connections, and see previous operations. |
<!-- End Codat Tags Table -->
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [Bank Feeds](#bank-feeds)
  * [Endpoints](#endpoints)
  * [SDK Installation](#sdk-installation)
  * [Example Usage](#example-usage)
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
go get github.com/codatio/client-sdk-go/bank-feeds
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
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v8"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/types"
	"log"
)

func main() {
	ctx := context.Background()

	s := bankfeeds.New()

	res, err := s.BankFeedsSourceAccountConnected(ctx, &shared.SourceAccountWebhook{
		EventType:     bankfeeds.String("bankFeeds.sourceAccount.connected"),
		GeneratedDate: bankfeeds.String("2022-10-23T00:00:00Z"),
		ID:            bankfeeds.String("ba29118f-5406-4e59-b05c-ba307ca38d01"),
		Payload: &shared.SourceAccountWebhookPayload{
			ConnectionID: bankfeeds.String("2e9d2c44-f675-40ba-8049-353bfcb5e171"),
			ReferenceCompany: &shared.CompanyReference{
				Description: bankfeeds.String("Requested early access to the new financing scheme."),
				ID:          bankfeeds.String("0498e921-9b53-4396-a412-4f2f5983b0a2"),
				Links: &shared.CompanyReferenceLinks{
					Portal: bankfeeds.String("https://app.codat.io/companies/0498e921-9b53-4396-a412-4f2f5983b0a2/summary"),
				},
				Name: bankfeeds.String("Toft stores"),
			},
			SourceAccount: bankfeeds.Pointer(shared.CreateSourceAccountWebhookPayloadSourceAccountSourceAccount(
				shared.SourceAccount{
					AccountName:   bankfeeds.String("account-081"),
					AccountNumber: bankfeeds.String("12345678"),
					Balance:       types.MustNewDecimalFromString("99.99"),
					Currency:      bankfeeds.String("GBP"),
					ID:            "acc-002",
					ModifiedDate:  bankfeeds.String("2023-01-09T14:14:14.105Z"),
					SortCode:      bankfeeds.String("040004"),
					Status:        shared.StatusPending.ToPointer(),
				},
			)),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AccountMapping](docs/sdks/accountmapping/README.md)

* [Create](docs/sdks/accountmapping/README.md#create) - Create bank feed account mapping
* [Get](docs/sdks/accountmapping/README.md#get) - List bank feed accounts

### [BankAccounts](docs/sdks/bankaccounts/README.md)

* [Create](docs/sdks/bankaccounts/README.md#create) - Create bank account
* [GetCreateModel](docs/sdks/bankaccounts/README.md#getcreatemodel) - Get create/update bank account model
* [List](docs/sdks/bankaccounts/README.md#list) - List bank accounts


### [Companies](docs/sdks/companies/README.md)

* [Create](docs/sdks/companies/README.md#create) - Create company
* [Delete](docs/sdks/companies/README.md#delete) - Delete a company
* [Get](docs/sdks/companies/README.md#get) - Get company
* [GetAccessToken](docs/sdks/companies/README.md#getaccesstoken) - Get company access token
* [List](docs/sdks/companies/README.md#list) - List companies
* [Update](docs/sdks/companies/README.md#update) - Update company

### [CompanyInformation](docs/sdks/companyinformation/README.md)

* [Get](docs/sdks/companyinformation/README.md#get) - Get company information

### [Configuration](docs/sdks/configuration/README.md)

* [Get](docs/sdks/configuration/README.md#get) - Get configuration
* [Set](docs/sdks/configuration/README.md#set) - Set configuration

### [Connections](docs/sdks/connections/README.md)

* [Create](docs/sdks/connections/README.md#create) - Create connection
* [Delete](docs/sdks/connections/README.md#delete) - Delete connection
* [Get](docs/sdks/connections/README.md#get) - Get connection
* [List](docs/sdks/connections/README.md#list) - List connections
* [Unlink](docs/sdks/connections/README.md#unlink) - Unlink connection

### [SourceAccounts](docs/sdks/sourceaccounts/README.md)

* [Create](docs/sdks/sourceaccounts/README.md#create) - Create single source account
* [CreateBatch](docs/sdks/sourceaccounts/README.md#createbatch) - Create source accounts
* [Delete](docs/sdks/sourceaccounts/README.md#delete) - Delete source account
* [DeleteCredentials](docs/sdks/sourceaccounts/README.md#deletecredentials) - Delete all source account credentials
* [GenerateCredentials](docs/sdks/sourceaccounts/README.md#generatecredentials) - Generate source account credentials
* [List](docs/sdks/sourceaccounts/README.md#list) - List source accounts
* [Update](docs/sdks/sourceaccounts/README.md#update) - Update source account

### [Sync](docs/sdks/sync/README.md)

* [GetLastSuccessfulSync](docs/sdks/sync/README.md#getlastsuccessfulsync) - Get last successful sync

### [Transactions](docs/sdks/transactions/README.md)

* [Create](docs/sdks/transactions/README.md#create) - Create bank transactions
* [GetCreateModel](docs/sdks/transactions/README.md#getcreatemodel) - Get create bank transactions model
* [GetCreateOperation](docs/sdks/transactions/README.md#getcreateoperation) - Get create operation
* [ListCreateOperations](docs/sdks/transactions/README.md#listcreateoperations) - List create operations

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
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v8"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/retry"
	"log"
	"pkg/models/operations"
)

func main() {
	ctx := context.Background()

	s := bankfeeds.New(
		bankfeeds.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: bankfeeds.String("Requested early access to the new financing scheme."),
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
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v8"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/shared"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := bankfeeds.New(
		bankfeeds.WithRetryConfig(
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
		bankfeeds.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: bankfeeds.String("Requested early access to the new financing scheme."),
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
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v8"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/sdkerrors"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := bankfeeds.New(
		bankfeeds.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: bankfeeds.String("Requested early access to the new financing scheme."),
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
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v8"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := bankfeeds.New(
		bankfeeds.WithServerURL("https://api.codat.io"),
		bankfeeds.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: bankfeeds.String("Requested early access to the new financing scheme."),
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
	bankfeeds "github.com/codatio/client-sdk-go/bank-feeds/v8"
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := bankfeeds.New(
		bankfeeds.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	res, err := s.Companies.Create(ctx, &shared.CompanyRequestBody{
		Description: bankfeeds.String("Requested early access to the new financing scheme."),
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