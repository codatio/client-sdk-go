# github.com/codatio/client-sdk-go/sync-for-payables

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/codatio/client-sdk-go/sync-for-payables* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/codatio/client-sdk-go/sync-for-payables&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasy.com/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasy.com/docs/advanced-setup/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README

<!-- Start Summary [summary] -->
## Summary

Bill pay kit: The API reference for the Bill Pay kit. 

The bill pay kit is an API and a set of supporting tools designed to integrate a bill pay flow into your app as quickly as possible. It's ideal for facilitating essential bill payment processes within your SMB's accounting software.

[Explore product](https://docs.codat.io/payables/bill-pay-kit) | [See OpenAPI spec](https://github.com/codatio/oas)

---
## Supported Integrations

| Integration                   | Supported |
|-------------------------------|-----------|
| FreeAgent                     | Yes       |
| QuickBooks Online             | Yes       |
| Oracle NetSuite               | Yes       |
| Xero                          | Yes       |

---
<!-- Start Codat Tags Table -->
## Endpoints

| Endpoints | Description |
| :- |:- |
| Companies | Create and manage your SMB users' companies. |
| Connections | Create new and manage existing data connections for a company. |
| Company information | View company profile from the source platform. |
| Bills | Get, create, and update Bills. |
| Bill payments | Get, create, and update Bill payments. |
| Suppliers | Get, create, and update Suppliers. |
| Bank accounts | Create a bank account for a given company's connection. |
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
go get github.com/codatio/client-sdk-go/sync-for-payables
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v5"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/shared"
	"log"
)

func main() {
	s := syncforpayables.New(
		syncforpayables.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.List(ctx, operations.ListCompaniesRequest{
		Page:     syncforpayables.Int(1),
		PageSize: syncforpayables.Int(100),
		Query:    syncforpayables.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
		OrderBy:  syncforpayables.String("-modifiedDate"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Companies != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [BankAccounts](docs/sdks/bankaccounts/README.md)

* [Create](docs/sdks/bankaccounts/README.md#create) - Create bank account

### [BillPayments](docs/sdks/billpayments/README.md)

* [GetPaymentOptions](docs/sdks/billpayments/README.md#getpaymentoptions) - Get payment mapping options
* [Create](docs/sdks/billpayments/README.md#create) - Create bill payment

### [Bills](docs/sdks/bills/README.md)

* [GetBillOptions](docs/sdks/bills/README.md#getbilloptions) - Get bill mapping options
* [List](docs/sdks/bills/README.md#list) - List bills
* [Create](docs/sdks/bills/README.md#create) - Create bill
* [UploadAttachment](docs/sdks/bills/README.md#uploadattachment) - Upload bill attachment
* [ListAttachments](docs/sdks/bills/README.md#listattachments) - List bill attachments
* [DownloadAttachment](docs/sdks/bills/README.md#downloadattachment) - Download bill attachment


### [Companies](docs/sdks/companies/README.md)

* [List](docs/sdks/companies/README.md#list) - List companies
* [Create](docs/sdks/companies/README.md#create) - Create company
* [Update](docs/sdks/companies/README.md#update) - Update company
* [Delete](docs/sdks/companies/README.md#delete) - Delete a company
* [Get](docs/sdks/companies/README.md#get) - Get company

### [CompanyInformation](docs/sdks/companyinformation/README.md)

* [Get](docs/sdks/companyinformation/README.md#get) - Get company information

### [Connections](docs/sdks/connections/README.md)

* [List](docs/sdks/connections/README.md#list) - List connections
* [Create](docs/sdks/connections/README.md#create) - Create connection
* [Get](docs/sdks/connections/README.md#get) - Get connection
* [Delete](docs/sdks/connections/README.md#delete) - Delete connection
* [Unlink](docs/sdks/connections/README.md#unlink) - Unlink connection

### [Suppliers](docs/sdks/suppliers/README.md)

* [List](docs/sdks/suppliers/README.md#list) - List suppliers
* [Create](docs/sdks/suppliers/README.md#create) - Create supplier

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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v5"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/retry"
	"log"
	"pkg/models/operations"
)

func main() {
	s := syncforpayables.New(
		syncforpayables.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.List(ctx, operations.ListCompaniesRequest{
		Page:     syncforpayables.Int(1),
		PageSize: syncforpayables.Int(100),
		Query:    syncforpayables.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
		OrderBy:  syncforpayables.String("-modifiedDate"),
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
	if res.Companies != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v5"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/retry"
	"log"
)

func main() {
	s := syncforpayables.New(
		syncforpayables.WithRetryConfig(
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
		syncforpayables.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.List(ctx, operations.ListCompaniesRequest{
		Page:     syncforpayables.Int(1),
		PageSize: syncforpayables.Int(100),
		Query:    syncforpayables.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
		OrderBy:  syncforpayables.String("-modifiedDate"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Companies != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `List` function may return the following errors:

| Error Type             | Status Code                            | Content Type     |
| ---------------------- | -------------------------------------- | ---------------- |
| sdkerrors.ErrorMessage | 400, 401, 402, 403, 404, 429, 500, 503 | application/json |
| sdkerrors.SDKError     | 4XX, 5XX                               | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v5"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/sdkerrors"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/shared"
	"log"
)

func main() {
	s := syncforpayables.New(
		syncforpayables.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.List(ctx, operations.ListCompaniesRequest{
		Page:     syncforpayables.Int(1),
		PageSize: syncforpayables.Int(100),
		Query:    syncforpayables.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
		OrderBy:  syncforpayables.String("-modifiedDate"),
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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v5"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/shared"
	"log"
)

func main() {
	s := syncforpayables.New(
		syncforpayables.WithServerURL("https://api.codat.io"),
		syncforpayables.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.List(ctx, operations.ListCompaniesRequest{
		Page:     syncforpayables.Int(1),
		PageSize: syncforpayables.Int(100),
		Query:    syncforpayables.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
		OrderBy:  syncforpayables.String("-modifiedDate"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Companies != nil {
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
	syncforpayables "github.com/codatio/client-sdk-go/sync-for-payables/v5"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/operations"
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/shared"
	"log"
)

func main() {
	s := syncforpayables.New(
		syncforpayables.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Companies.List(ctx, operations.ListCompaniesRequest{
		Page:     syncforpayables.Int(1),
		PageSize: syncforpayables.Int(100),
		Query:    syncforpayables.String("id=e3334455-1aed-4e71-ab43-6bccf12092ee"),
		OrderBy:  syncforpayables.String("-modifiedDate"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Companies != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/codatio/client-sdk-go/sync-for-payables&utm_campaign=go)
