# Sync for Commerce

<!-- Start Codat Library Description -->
﻿Embedded accounting integrations for POS and eCommerce platforms.
<!-- End Codat Library Description -->

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/codatio/client-sdk-go/sync-for-commerce
```
<!-- End SDK Installation -->

## Example Usage
<!-- Start SDK Example Usage -->
### Example

```go
package main

import (
	"context"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v3"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/shared"
	"log"
)

func main() {
	s := syncforcommerce.New(
		syncforcommerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.SyncFlowSettings.GetConfigTextSyncFlow(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.LocalizationInfo != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [SyncFlowSettings](docs/sdks/syncflowsettings/README.md)

* [GetConfigTextSyncFlow](docs/sdks/syncflowsettings/README.md#getconfigtextsyncflow) - Get preferences for text fields
* [GetVisibleAccounts](docs/sdks/syncflowsettings/README.md#getvisibleaccounts) - List visible accounts
* [UpdateConfigTextSyncFlow](docs/sdks/syncflowsettings/README.md#updateconfigtextsyncflow) - Update preferences for text fields
* [UpdateVisibleAccountsSyncFlow](docs/sdks/syncflowsettings/README.md#updatevisibleaccountssyncflow) - Update visible accounts

### [AdvancedControls](docs/sdks/advancedcontrols/README.md)

* [CreateCompany](docs/sdks/advancedcontrols/README.md#createcompany) - Create company
* [GetConfiguration](docs/sdks/advancedcontrols/README.md#getconfiguration) - Get company configuration
* [ListCompanies](docs/sdks/advancedcontrols/README.md#listcompanies) - List companies
* [SetConfiguration](docs/sdks/advancedcontrols/README.md#setconfiguration) - Set configuration

### [Connections](docs/sdks/connections/README.md)

* [Create](docs/sdks/connections/README.md#create) - Create connection
* [GetSyncFlowURL](docs/sdks/connections/README.md#getsyncflowurl) - Start new sync flow
* [List](docs/sdks/connections/README.md#list) - List connections
* [UpdateAuthorization](docs/sdks/connections/README.md#updateauthorization) - Update authorization
* [UpdateConnection](docs/sdks/connections/README.md#updateconnection) - Update connection

### [Sync](docs/sdks/sync/README.md)

* [Get](docs/sdks/sync/README.md#get) - Get sync status
* [GetLastSuccessfulSync](docs/sdks/sync/README.md#getlastsuccessfulsync) - Last successful sync
* [GetLatestSync](docs/sdks/sync/README.md#getlatestsync) - Latest sync status
* [GetStatus](docs/sdks/sync/README.md#getstatus) - Get sync status
* [List](docs/sdks/sync/README.md#list) - List sync statuses
* [Request](docs/sdks/sync/README.md#request) - Initiate new sync
* [RequestForDateRange](docs/sdks/sync/README.md#requestfordaterange) - Initiate sync for specific range

### [Integrations](docs/sdks/integrations/README.md)

* [GetBranding](docs/sdks/integrations/README.md#getbranding) - Get branding for an integration
* [List](docs/sdks/integrations/README.md#list) - List integrations
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
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v3"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/shared"
	"log"
	"pkg/models/operations"
	"pkg/utils"
)

func main() {
	s := syncforcommerce.New(
		syncforcommerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.SyncFlowSettings.GetConfigTextSyncFlow(ctx, operations.WithRetries(utils.RetryConfig{
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

	if res.LocalizationInfo != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can provide a retryConfig at SDK initialization:
```go
package main

import (
	"context"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v3"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/shared"
	"log"
	"pkg/models/operations"
	"pkg/utils"
)

func main() {
	s := syncforcommerce.New(
		syncforcommerce.WithRetryConfig(utils.RetryConfig{
			Strategy: "backoff",
			Backoff: &utils.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}),
		syncforcommerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.SyncFlowSettings.GetConfigTextSyncFlow(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.LocalizationInfo != nil {
		// handle response
	}
}

```


<!-- End Retries -->



<!-- Start Error Handling -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object            | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorMessage  | 401,402,403,429,500,503 | application/json        |
| sdkerrors.SDKError      | 400-600                 | */*                     |

### Example

```go
package main

import (
	"context"
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v3"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/shared"
	"log"
)

func main() {
	s := syncforcommerce.New(
		syncforcommerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.SyncFlowSettings.GetConfigTextSyncFlow(ctx)
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
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v3"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/shared"
	"log"
)

func main() {
	s := syncforcommerce.New(
		syncforcommerce.WithServerIndex(0),
		syncforcommerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.SyncFlowSettings.GetConfigTextSyncFlow(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.LocalizationInfo != nil {
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
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v3"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/shared"
	"log"
)

func main() {
	s := syncforcommerce.New(
		syncforcommerce.WithServerURL("https://api.codat.io"),
		syncforcommerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.SyncFlowSettings.GetConfigTextSyncFlow(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.LocalizationInfo != nil {
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
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce/v3"
	"github.com/codatio/client-sdk-go/sync-for-commerce/v3/pkg/models/shared"
	"log"
)

func main() {
	s := syncforcommerce.New(
		syncforcommerce.WithSecurity(shared.Security{
			AuthHeader: "Basic BASE_64_ENCODED(API_KEY)",
		}),
	)

	ctx := context.Background()
	res, err := s.SyncFlowSettings.GetConfigTextSyncFlow(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.LocalizationInfo != nil {
		// handle response
	}
}

```
<!-- End Authentication -->

<!-- Placeholder for Future Speakeasy SDK Sections -->


### Library generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)