# Common

<!-- Start Codat Library Description -->
Manage the building blocks of Codat, including companies, connections, and more.
<!-- End Codat Library Description -->

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/previous-versions/common
```
<!-- End SDK Installation -->

## Example Usage
<!-- Start SDK Example Usage -->
### Example

```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"log"
)

func main() {
	s := common.New(
		common.WithSecurity(shared.Security{
			AuthHeader: "",
		}),
	)

	ctx := context.Background()
	res, err := s.Settings.CreateAPIKey(ctx, &shared.CreateAPIKey{
		Name: common.String("azure-invoice-finance-processor"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.APIKeyDetails != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
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

### [Connections](docs/sdks/connections/README.md)

* [Create](docs/sdks/connections/README.md#create) - Create connection
* [Delete](docs/sdks/connections/README.md#delete) - Delete connection
* [Get](docs/sdks/connections/README.md#get) - Get connection
* [List](docs/sdks/connections/README.md#list) - List connections
* [Unlink](docs/sdks/connections/README.md#unlink) - Unlink connection
* [UpdateAuthorization](docs/sdks/connections/README.md#updateauthorization) - Update authorization

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

### [Integrations](docs/sdks/integrations/README.md)

* [Get](docs/sdks/integrations/README.md#get) - Get integration
* [GetBranding](docs/sdks/integrations/README.md#getbranding) - Get branding
* [List](docs/sdks/integrations/README.md#list) - List integrations

### [SupplementalData](docs/sdks/supplementaldata/README.md)

* [Configure](docs/sdks/supplementaldata/README.md#configure) - Configure
* [GetConfiguration](docs/sdks/supplementaldata/README.md#getconfiguration) - Get configuration

### [Webhooks](docs/sdks/webhooks/README.md)

* [Create](docs/sdks/webhooks/README.md#create) - Create webhook
* [Get](docs/sdks/webhooks/README.md#get) - Get webhook
* [List](docs/sdks/webhooks/README.md#list) - List webhooks
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Error Handling -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorMessage          | 400,401,402,403,409,429,500,503 | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |

### Example

```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"log"
)

func main() {
	s := common.New(
		common.WithSecurity(shared.Security{
			AuthHeader: "",
		}),
	)

	ctx := context.Background()
	res, err := s.Settings.CreateAPIKey(ctx, &shared.CreateAPIKey{
		Name: common.String("azure-invoice-finance-processor"),
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
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"log"
)

func main() {
	s := common.New(
		common.WithServerIndex(0),
		common.WithSecurity(shared.Security{
			AuthHeader: "",
		}),
	)

	ctx := context.Background()
	res, err := s.Settings.CreateAPIKey(ctx, &shared.CreateAPIKey{
		Name: common.String("azure-invoice-finance-processor"),
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
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"log"
)

func main() {
	s := common.New(
		common.WithServerURL("https://api.codat.io"),
		common.WithSecurity(shared.Security{
			AuthHeader: "",
		}),
	)

	ctx := context.Background()
	res, err := s.Settings.CreateAPIKey(ctx, &shared.CreateAPIKey{
		Name: common.String("azure-invoice-finance-processor"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.APIKeyDetails != nil {
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
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"log"
	"pkg/models/operations"
	"pkg/utils"
)

func main() {
	s := common.New(
		common.WithSecurity(shared.Security{
			AuthHeader: "",
		}),
	)

	ctx := context.Background()
	res, err := s.Settings.CreateAPIKey(ctx, &shared.CreateAPIKey{
		Name: common.String("azure-invoice-finance-processor"),
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

	if res.APIKeyDetails != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can provide a retryConfig at SDK initialization:
```go
package main

import (
	"context"
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"log"
	"pkg/models/operations"
	"pkg/utils"
)

func main() {
	s := common.New(
		common.WithRetryConfig(utils.RetryConfig{
			Strategy: "backoff",
			Backoff: &utils.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}),
		common.WithSecurity(shared.Security{
			AuthHeader: "",
		}),
	)

	ctx := context.Background()
	res, err := s.Settings.CreateAPIKey(ctx, &shared.CreateAPIKey{
		Name: common.String("azure-invoice-finance-processor"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.APIKeyDetails != nil {
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
	"github.com/codatio/client-sdk-go/previous-versions/common"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"log"
)

func main() {
	s := common.New(
		common.WithSecurity(shared.Security{
			AuthHeader: "",
		}),
	)

	ctx := context.Background()
	res, err := s.Settings.CreateAPIKey(ctx, &shared.CreateAPIKey{
		Name: common.String("azure-invoice-finance-processor"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.APIKeyDetails != nil {
		// handle response
	}
}

```
<!-- End Authentication -->

<!-- Placeholder for Future Speakeasy SDK Sections -->


### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)