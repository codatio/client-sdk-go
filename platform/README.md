# Platform

<!-- Start Codat Library Description -->
Manage the building blocks of Codat, including companies, connections, and more.
<!-- End Codat Library Description -->

<!-- Start Summary [summary] -->
## Summary

Platform API: Platform API

An API for the common components of all of Codat's products.

These end points cover creating and managing your companies, data connections, and integrations.

[Read about the building blocks of Codat...](https://docs.codat.io/core-concepts/companies) | [See our OpenAPI spec](https://github.com/codatio/oas) 

---
<!-- Start Codat Tags Table -->
## Endpoints

| Endpoints | Description |
| :- |:- |
| Companies | Create and manage your SMB users' companies. |
| Connections | Create new and manage existing data connections for a company. |
| Connection management | Configure connection management UI and retrieve access tokens for authentication. |
| Webhooks | Create and manage webhooks that listen to Codat's events. |
| Integrations | Get a list of integrations supported by Codat and their logos. |
| Refresh data | Initiate data refreshes, view pull status and history. |
| Settings | Manage company profile configuration, sync settings, and API keys. |
| Push data | Initiate and monitor Create, Update, and Delete operations. |
| Supplemental data | Configure and pull additional data you can include in Codat's standard data types. |
| Custom data type | Configure and pull additional data types that are not included in Codat's standardized data model. |
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
go get github.com/codatio/client-sdk-go/platform
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
	platform "github.com/codatio/client-sdk-go/platform/v5"
	"github.com/codatio/client-sdk-go/platform/v5/pkg/models/shared"
	"log"
)

func main() {
	s := platform.New(
		platform.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Settings.CreateAPIKey(ctx, &shared.CreateAPIKey{
		Name: platform.String("azure-invoice-finance-processor"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.APIKeyDetails != nil {
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

* [AddProduct](docs/sdks/companies/README.md#addproduct) - Add product
* [Create](docs/sdks/companies/README.md#create) - Create company
* [Delete](docs/sdks/companies/README.md#delete) - Delete a company
* [Get](docs/sdks/companies/README.md#get) - Get company
* [GetAccessToken](docs/sdks/companies/README.md#getaccesstoken) - Get company access token
* [List](docs/sdks/companies/README.md#list) - List companies
* [RemoveProduct](docs/sdks/companies/README.md#removeproduct) - Remove product
* [Update](docs/sdks/companies/README.md#update) - Update company

### [ConnectionManagement](docs/sdks/connectionmanagement/README.md)

* [GetAccessToken](docs/sdks/connectionmanagement/README.md#getaccesstoken) - Get access token

#### [ConnectionManagement.CorsSettings](docs/sdks/corssettings/README.md)

* [Get](docs/sdks/corssettings/README.md#get) - Get CORS settings
* [Set](docs/sdks/corssettings/README.md#set) - Set CORS settings

### [Connections](docs/sdks/connections/README.md)

* [Create](docs/sdks/connections/README.md#create) - Create connection
* [Delete](docs/sdks/connections/README.md#delete) - Delete connection
* [Get](docs/sdks/connections/README.md#get) - Get connection
* [List](docs/sdks/connections/README.md#list) - List connections
* [Unlink](docs/sdks/connections/README.md#unlink) - Unlink connection
* [UpdateAuthorization](docs/sdks/connections/README.md#updateauthorization) - Update authorization

### [CustomDataType](docs/sdks/customdatatype/README.md)

* [Configure](docs/sdks/customdatatype/README.md#configure) - Configure custom data type
* [GetConfiguration](docs/sdks/customdatatype/README.md#getconfiguration) - Get custom data configuration
* [List](docs/sdks/customdatatype/README.md#list) - List custom data type records
* [Refresh](docs/sdks/customdatatype/README.md#refresh) - Refresh custom data type

### [Integrations](docs/sdks/integrations/README.md)

* [Get](docs/sdks/integrations/README.md#get) - Get integration
* [GetBranding](docs/sdks/integrations/README.md#getbranding) - Get branding
* [List](docs/sdks/integrations/README.md#list) - List integrations

### [PushData](docs/sdks/pushdata/README.md)

* [GetModelOptions](docs/sdks/pushdata/README.md#getmodeloptions) - Get push options
* [GetOperation](docs/sdks/pushdata/README.md#getoperation) - Get push operation
* [ListOperations](docs/sdks/pushdata/README.md#listoperations) - List push operations

### [RefreshData](docs/sdks/refreshdata/README.md)

* [All](docs/sdks/refreshdata/README.md#all) - Refresh all data
* [ByDataType](docs/sdks/refreshdata/README.md#bydatatype) - Refresh data type
* [Get](docs/sdks/refreshdata/README.md#get) - Get data status
* [GetPullOperation](docs/sdks/refreshdata/README.md#getpulloperation) - Get pull operation
* [ListPullOperations](docs/sdks/refreshdata/README.md#listpulloperations) - List pull operations

### [Settings](docs/sdks/settings/README.md)

* [CreateAPIKey](docs/sdks/settings/README.md#createapikey) - Create API key
* [DeleteAPIKey](docs/sdks/settings/README.md#deleteapikey) - Delete API key
* [GetProfile](docs/sdks/settings/README.md#getprofile) - Get profile
* [GetSyncSettings](docs/sdks/settings/README.md#getsyncsettings) - Get sync settings
* [ListAPIKeys](docs/sdks/settings/README.md#listapikeys) - List API keys
* [UpdateProfile](docs/sdks/settings/README.md#updateprofile) - Update profile
* [UpdateSyncSettings](docs/sdks/settings/README.md#updatesyncsettings) - Update all sync settings

### [SupplementalData](docs/sdks/supplementaldata/README.md)

* [Configure](docs/sdks/supplementaldata/README.md#configure) - Configure
* [GetConfiguration](docs/sdks/supplementaldata/README.md#getconfiguration) - Get configuration

### [Webhooks](docs/sdks/webhooks/README.md)

* [~~Create~~](docs/sdks/webhooks/README.md#create) - Create webhook (legacy) :warning: **Deprecated**
* [CreateConsumer](docs/sdks/webhooks/README.md#createconsumer) - Create webhook consumer
* [DeleteConsumer](docs/sdks/webhooks/README.md#deleteconsumer) - Delete webhook consumer
* [~~Get~~](docs/sdks/webhooks/README.md#get) - Get webhook (legacy) :warning: **Deprecated**
* [~~List~~](docs/sdks/webhooks/README.md#list) - List webhooks (legacy) :warning: **Deprecated**
* [ListConsumers](docs/sdks/webhooks/README.md#listconsumers) - List webhook consumers

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
	platform "github.com/codatio/client-sdk-go/platform/v5"
	"github.com/codatio/client-sdk-go/platform/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/platform/v5/pkg/retry"
	"log"
	"pkg/models/operations"
)

func main() {
	s := platform.New(
		platform.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Settings.CreateAPIKey(ctx, &shared.CreateAPIKey{
		Name: platform.String("azure-invoice-finance-processor"),
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
	if res.APIKeyDetails != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	platform "github.com/codatio/client-sdk-go/platform/v5"
	"github.com/codatio/client-sdk-go/platform/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/platform/v5/pkg/retry"
	"log"
)

func main() {
	s := platform.New(
		platform.WithRetryConfig(
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
		platform.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Settings.CreateAPIKey(ctx, &shared.CreateAPIKey{
		Name: platform.String("azure-invoice-finance-processor"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.APIKeyDetails != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `CreateAPIKey` function may return the following errors:

| Error Type             | Status Code                            | Content Type     |
| ---------------------- | -------------------------------------- | ---------------- |
| sdkerrors.ErrorMessage | 400, 401, 402, 403, 409, 429, 500, 503 | application/json |
| sdkerrors.SDKError     | 4XX, 5XX                               | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	platform "github.com/codatio/client-sdk-go/platform/v5"
	"github.com/codatio/client-sdk-go/platform/v5/pkg/models/sdkerrors"
	"github.com/codatio/client-sdk-go/platform/v5/pkg/models/shared"
	"log"
)

func main() {
	s := platform.New(
		platform.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Settings.CreateAPIKey(ctx, &shared.CreateAPIKey{
		Name: platform.String("azure-invoice-finance-processor"),
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
	platform "github.com/codatio/client-sdk-go/platform/v5"
	"github.com/codatio/client-sdk-go/platform/v5/pkg/models/shared"
	"log"
)

func main() {
	s := platform.New(
		platform.WithServerURL("https://api.codat.io"),
		platform.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Settings.CreateAPIKey(ctx, &shared.CreateAPIKey{
		Name: platform.String("azure-invoice-finance-processor"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.APIKeyDetails != nil {
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
	platform "github.com/codatio/client-sdk-go/platform/v5"
	"github.com/codatio/client-sdk-go/platform/v5/pkg/models/shared"
	"log"
)

func main() {
	s := platform.New(
		platform.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.Settings.CreateAPIKey(ctx, &shared.CreateAPIKey{
		Name: platform.String("azure-invoice-finance-processor"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.APIKeyDetails != nil {
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