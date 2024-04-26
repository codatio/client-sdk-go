# Platform

<!-- Start Codat Library Description -->
Manage the building blocks of Codat, including companies, connections, and more.
<!-- End Codat Library Description -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

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
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
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

### [Settings](docs/sdks/settings/README.md)

* [CreateAPIKey](docs/sdks/settings/README.md#createapikey) - Create API key
* [DeleteAPIKey](docs/sdks/settings/README.md#deleteapikey) - Delete API key
* [GetProfile](docs/sdks/settings/README.md#getprofile) - Get profile
* [GetSyncSettings](docs/sdks/settings/README.md#getsyncsettings) - Get sync settings
* [ListAPIKeys](docs/sdks/settings/README.md#listapikeys) - List API keys
* [UpdateProfile](docs/sdks/settings/README.md#updateprofile) - Update profile
* [UpdateSyncSettings](docs/sdks/settings/README.md#updatesyncsettings) - Update all sync settings

### [Companies](docs/sdks/companies/README.md)

* [Create](docs/sdks/companies/README.md#create) - Create company
* [Delete](docs/sdks/companies/README.md#delete) - Delete a company
* [Get](docs/sdks/companies/README.md#get) - Get company
* [List](docs/sdks/companies/README.md#list) - List companies
* [Update](docs/sdks/companies/README.md#update) - Update company

### [ConnectionManagement](docs/sdks/connectionmanagement/README.md)

* [GetAccessToken](docs/sdks/connectionmanagement/README.md#getaccesstoken) - Get access token

### [ConnectionManagement.CorsSettings](docs/sdks/corssettings/README.md)

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

### [Groups](docs/sdks/groups/README.md)

* [AddCompany](docs/sdks/groups/README.md#addcompany) - Add company
* [Create](docs/sdks/groups/README.md#create) - Create group
* [List](docs/sdks/groups/README.md#list) - List groups
* [RemoveCompany](docs/sdks/groups/README.md#removecompany) - Remove company

### [Integrations](docs/sdks/integrations/README.md)

* [Get](docs/sdks/integrations/README.md#get) - Get integration
* [GetBranding](docs/sdks/integrations/README.md#getbranding) - Get branding
* [List](docs/sdks/integrations/README.md#list) - List integrations

### [SupplementalData](docs/sdks/supplementaldata/README.md)

* [Configure](docs/sdks/supplementaldata/README.md#configure) - Configure
* [GetConfiguration](docs/sdks/supplementaldata/README.md#getconfiguration) - Get configuration

### [Webhooks](docs/sdks/webhooks/README.md)

* [~~Create~~](docs/sdks/webhooks/README.md#create) - Create webhook :warning: **Deprecated**
* [CreateConsumer](docs/sdks/webhooks/README.md#createconsumer) - Create webhook consumer
* [DeleteConsumer](docs/sdks/webhooks/README.md#deleteconsumer) - Delete webhook consumer
* [~~Get~~](docs/sdks/webhooks/README.md#get) - Get webhook :warning: **Deprecated**
* [~~List~~](docs/sdks/webhooks/README.md#list) - List webhooks :warning: **Deprecated**
* [ListConsumers](docs/sdks/webhooks/README.md#listconsumers) - List webhook consumers
<!-- End Available Resources and Operations [operations] -->





<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `RetryConfig` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/utils"
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
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/utils"
	"log"
)

func main() {
	s := platform.New(
		platform.WithRetryConfig(
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

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,409,429,500,503 | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |

### Example

```go
package main

import (
	"context"
	"errors"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/sdkerrors"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
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
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	"log"
)

func main() {
	s := platform.New(
		platform.WithServerIndex(0),
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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
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

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `AuthHeader` | apiKey       | API key      |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	platform "github.com/codatio/client-sdk-go/platform/v3"
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
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